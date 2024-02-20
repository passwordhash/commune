package websocket

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const internalServiceErrCode = -1

type SaveResourceError struct {
	Err error
}

func (e SaveResourceError) String() string {
	msg := fmt.Sprintf("CODE -%d: %v", internalServiceErrCode, e.Err)
	logrus.Error(msg)
	return msg
}
