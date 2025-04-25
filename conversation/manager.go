package conversation

import (
	"database/sql"
	"log"
	"minichat/constant"
	"minichat/util"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type ConversationManager struct {
	Rooms          map[string]*Room
	Register       chan *Client
	unregister     chan *Client
	broadcast      chan Message
	registerLock   *sync.RWMutex
	unregisterLock *sync.RWMutex
	broadcastLock  *sync.RWMutex
}

type Message struct {
	ID         string    `json:"id"`
	RoomNumber string    `json:"room_number"`
	UserName   string    `json:"username"`
	Cmd        string    `json:"cmd"`
	Payload    string    `json:"payload"`
	Timestamp  time.Time `json:"timestamp"`
}

type Client struct {
	//cmd    string
	RoomNumber string
	UserName   string
	Password   string
	Send       chan Message
	Conn       *websocket.Conn
	Stop       chan bool
}

type Room struct {
	Clients  map[*Client]*Client
	RoomName string
	Password string
}

var Manager = ConversationManager{
	broadcast:      make(chan Message),
	Register:       make(chan *Client),
	unregister:     make(chan *Client),
	Rooms:          make(map[string]*Room),
	registerLock:   new(sync.RWMutex),
	unregisterLock: new(sync.RWMutex),
	broadcastLock:  new(sync.RWMutex),
}

func (manager *ConversationManager) Start() {
	for {
		select {
		case client := <-manager.Register:
			// 新客户端链接
			manager.registerLock.Lock()
			if _, ok := manager.Rooms[client.RoomNumber]; !ok {
				manager.Rooms[client.RoomNumber] = &Room{
					Clients:  make(map[*Client]*Client),
					Password: client.Password,
				}
			}

			// 检查用户是否已经存在，如果不存在则插入到 users 表
			var userId int64
			err := util.DB.QueryRow("SELECT id FROM users WHERE username = ?", client.UserName).Scan(&userId)
			if err == sql.ErrNoRows {
				// 用户不存在，插入新用户
				res, err := util.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", client.UserName, client.Password)
				if err != nil {
					log.Println("Error inserting user into database:", err)
					continue // 继续处理下一个客户端
				}
				userId, err = res.LastInsertId()
			} else if err != nil {
				log.Println("Error querying user from database:", err)
				continue // 继续处理下一个客户端
			}

			// 检查聊天室是否已经存在，如果不存在则插入到 chat_rooms 表
			var chatRoomId int64
			err = util.DB.QueryRow("SELECT id FROM chat_rooms WHERE name = ?", client.RoomNumber).Scan(&chatRoomId)
			if err == sql.ErrNoRows {
				// 聊天室不存在，插入新聊天室
				res, err := util.DB.Exec("INSERT INTO chat_rooms (name) VALUES (?)", client.RoomNumber)
				if err != nil {
					log.Println("Error inserting chat room into database:", err)
					continue // 继续处理下一个客户端
				}
				chatRoomId, err = res.LastInsertId()
			} else if err != nil {
				log.Println("Error querying chat room from database:", err)
				continue // 继续处理下一个客户端
			}

			// 将用户和聊天室的关系插入到 user_chat_rooms 表
			_, err = util.DB.Exec("INSERT INTO user_chat_rooms (user_id, chat_room_id) VALUES (?, ?)", userId, chatRoomId)
			if err != nil {
				log.Println("Error inserting user-chat room relationship into database:", err)
				continue // 继续处理下一个客户端
			}

			// 塞入房间初次数据
			manager.Rooms[client.RoomNumber].Clients[client] = client
			go func() {
				names := ""
				for key, _ := range manager.Rooms[client.RoomNumber].Clients {
					names += "[ " + key.UserName + " ], "
				}
				names = "<span class='is-inline-block'>" + strings.TrimSuffix(names, ", ") + "</span>"
				manager.broadcast <- Message{
					UserName:   client.UserName,
					Payload:    constant.JoinSuccess + constant.Online + names,
					RoomNumber: client.RoomNumber,
					Cmd:        constant.CmdJoin,
				}
			}()
			manager.registerLock.Unlock()

		case client := <-manager.unregister:
			// 客户端断开链接
			manager.unregisterLock.Lock()
			err := client.Conn.Close()
			if err != nil {
				return
			}
			if _, ok := manager.Rooms[client.RoomNumber]; ok {
				delete(manager.Rooms[client.RoomNumber].Clients, client)
				if len(manager.Rooms[client.RoomNumber].Clients) == 0 {
					delete(manager.Rooms, client.RoomNumber)
				}
				//client.stop <- true
				safeClose(client.Send)

				if manager.Rooms != nil && len(manager.Rooms) != 0 && manager.Rooms[client.RoomNumber] != nil && client.RoomNumber != "" {
					for c, _ := range manager.Rooms[client.RoomNumber].Clients {
						names := ""
						for key, _ := range manager.Rooms[client.RoomNumber].Clients {
							names += "[ " + key.UserName + " ], "
						}
						names = strings.TrimSuffix(names, ", ")
						names = "<span class='is-inline-block'>" + strings.TrimSuffix(names, ", ") + "</span>"
						c.Send <- Message{
							UserName:   client.UserName,
							Payload:    constant.ExitSuccess + constant.Online + names,
							RoomNumber: client.RoomNumber,
							Cmd:        constant.CmdExit,
						}
					}
				}
			}

			manager.unregisterLock.Unlock()

		case message := <-manager.broadcast:
			// 设置消息的时间戳
			message.Timestamp = time.Now()

			// 保存消息到数据库
			if message.Cmd != constant.CmdJoin && message.Cmd != constant.CmdExit {
				err := SaveMessageToDB(message)
				if err != nil {
					log.Println("Error saving message to database:", err)

				}
			}

			// 广播消息
			manager.broadcastLock.RLock()
			if manager.Rooms != nil && len(manager.Rooms) != 0 && manager.Rooms[message.RoomNumber] != nil && message.RoomNumber != "" {
				for c, _ := range manager.Rooms[message.RoomNumber].Clients {
					if c != nil && c.Conn != nil && c.Send != nil {
						c.Send <- message
					}
				}
			}
			manager.broadcastLock.RUnlock()
		}

	}
}

func safeClose(ch chan Message) {
	defer func() {
		if recover() != nil {
			log.Println("ch is closed")
		}
	}()
	close(ch)
	log.Println("ch closed successfully")
}

func SaveMessageToDB(message Message) error {
	stmt := "INSERT INTO chat_messages (room_number, username, cmd, payload, timestamp) VALUES (?, ?, ?, ?, ?)"
	_, err := util.DB.Exec(stmt,
		message.RoomNumber,
		message.UserName,
		message.Cmd,
		message.Payload,
		message.Timestamp.Format(time.RFC3339), // 格式化时间
	)
	return err
}
