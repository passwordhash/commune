package service

import (
	"commune/internal/entity"
	"commune/internal/repository"
	"commune/pkg/auth"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"time"
)

var (
	UserNotFound      = errors.New("user with this credentials not found")
	UserCreationError = errors.New("user did not created")
	UserAlreadyExists = errors.New("user with with credentials already exists")
)

var (
	AccessTokenCreationError = errors.New("access token did not created")
)

type UserService struct {
	userRepo repository.User

	passcodeSalt   string
	accessTokenTTL time.Duration
	sigingKey      string
}

func NewUserService(userRepo repository.User, deps Deps) *UserService {
	return &UserService{
		userRepo:       userRepo,
		passcodeSalt:   deps.PasscodeSalt,
		accessTokenTTL: deps.AccessTokenTTL,
		sigingKey:      deps.SigingKey,
	}
}

func (s *UserService) SignUp(u entity.UserCreate) (entity.JWTToken, entity.Passcode, error) {
	condidate, _ := s.userRepo.GetByEmail(u.Email)
	if !condidate.IsEmpty() {
		return "", "", UserAlreadyExists
	}

	passcode, err := s.GeneratePasscode()
	if err != nil {
		return "", "", err
	}
	hashedPasscode := auth.GeneratePasswordHash(string(passcode), s.passcodeSalt)

	user := entity.User{
		ID:        entity.NewObjectId(),
		Nickname:  u.Nickname,
		Email:     u.Email,
		Passcode:  entity.Passcode(hashedPasscode),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = s.userRepo.Create(user)
	if err != nil {
		return "", "", UserCreationError
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return "", "", AccessTokenCreationError
	}

	return token, passcode, nil
}

func (s *UserService) GeneratePasscode() (entity.Passcode, error) {
	b := make([]byte, 3)

	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return entity.Passcode(fmt.Sprintf("%x", b)), nil
}

func (s *UserService) UpdatePasscode(u entity.UserCreate) (entity.Passcode, error) {
	condidate, _ := s.userRepo.GetByEmail(u.Email)
	if condidate.IsEmpty() || condidate.Nickname != u.Nickname {
		return "", UserNotFound
	}

	newPasscode, err := s.GeneratePasscode()
	if err != nil {
		return "", err
	}
	newHashedPasscode := auth.GeneratePasswordHash(string(newPasscode), s.passcodeSalt)

	return newPasscode, s.userRepo.UpdatePasscode(u.Email, newHashedPasscode)
}

func (s *UserService) GetById(id entity.ObjectID) (entity.User, error) {
	return s.userRepo.GetById(id)
}

func (s *UserService) GetAll() ([]entity.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) Authenticate(credentials entity.UserAuth) (entity.AuthData, error) {
	var token entity.JWTToken

	// Получение кандидата
	user, err := s.userRepo.GetByEmail(credentials.Email)
	hashedPasscode := auth.GeneratePasswordHash(string(credentials.Passcode), s.passcodeSalt)
	if hashedPasscode != string(user.Passcode) {
		return entity.AuthData{}, UserNotFound
	}
	if err != nil {
		return entity.AuthData{}, UserNotFound
	}

	token, err = s.newAccessToken(user)
	if err != nil {
		return entity.AuthData{}, err
	}

	return entity.AuthData{ID: user.ID, Token: token}, nil
}

type tokenClaims struct {
	UserId entity.ObjectID
	jwt.RegisteredClaims
}

func (s *UserService) GenerateToken(u entity.User) (entity.JWTToken, error) {
	return s.newAccessToken(u)
}

func (s *UserService) ParseToken(accessToken entity.JWTToken) (entity.ObjectID, error) {
	var id entity.ObjectID

	token, err := jwt.ParseWithClaims(string(accessToken), &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid siging method")
		}
		return []byte(s.sigingKey), nil
	})
	if err != nil {
		return id, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return id, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *UserService) newAccessToken(user entity.User) (entity.JWTToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		tokenClaims{
			user.ID,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTokenTTL)),
			},
		},
	)

	t, err := token.SignedString([]byte(s.sigingKey))
	return entity.JWTToken(t), err
}
