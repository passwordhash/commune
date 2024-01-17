package websocket

type Message struct {
	data []byte
	room string
}

type Subscription struct {
	client *Client
	room   string
}

type Hub struct {
	// Зарегистрированные клиенты, распределенные по комантам
	rooms map[string]map[*Client]bool

	// Входящие сообщения от клиентов
	broadcast chan Message

	// Запросы на регистрацию
	register chan Subscription

	// Запросы на снятие c регистрации
	unregister chan Subscription
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan Subscription),
		unregister: make(chan Subscription),
		rooms:      make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case sub := <-h.register:
			roomClients := h.rooms[sub.room]
			// Если комнаты еще не было, создаем и добавляем в нее нового клиента
			if roomClients == nil {
				roomClients = make(map[*Client]bool)
				h.rooms[sub.room] = roomClients
			}
			// Устанавливаем, что автор сообщения подключен и находится в комнате
			h.rooms[sub.room][sub.client] = true
		case sub := <-h.unregister:
			roomClients := h.rooms[sub.room]
			// Если, клиент, отправиший сообщение, есть в комнате, отключаем его
			if roomClients != nil {
				if _, ok := roomClients[sub.client]; ok {
					delete(roomClients, sub.client)
					close(sub.client.send)
					if len(roomClients) == 0 {
						delete(h.rooms, sub.room)
					}
				}
			}
		case msg := <-h.broadcast:
			roomClients := h.rooms[msg.room]
			for client := range roomClients {
				select {
				case client.send <- msg.data:
				default:
					close(client.send)
					delete(roomClients, client)
					if len(roomClients) == 0 {
						delete(h.rooms, msg.room)
					}
				}
			}
		}

	}
}
