package conversation

import (
	"encoding/json"
	"log"
	"minichat/constant"
	"minichat/util"
)

func (c *Client) Read() {
	defer func() {
		Manager.unregister <- c
	}()

	for {
		message, err := util.SocketReceive(c.Conn)
		if err != nil {
			return
		}

		msgStr := string(message)
		// 判断是否为撤回指令，格式：/recall <msgid>
		if len(msgStr) > 8 && msgStr[:8] == "/recall " {
			msgID := msgStr[8:]
			Manager.broadcast <- Message{
				UserName:   c.UserName,
				RoomNumber: c.RoomNumber,
				Payload:    msgID, // 被撤回消息的ID
				Cmd:        constant.CmdRecall,
			}
			continue
		}

		Manager.broadcast <- Message{
			UserName:   c.UserName,
			RoomNumber: c.RoomNumber,
			Payload:    msgStr,
			Cmd:        constant.CmdChat,
		}
	}
}

func (c *Client) Write() {
	for {
		select {
		case message, isOpen := <-c.Send:
			if !isOpen {
				log.Printf("chan is closed")
				return
			}

			byteData, err := json.Marshal(message)
			if err != nil {
				log.Printf("json marshal error, error is %+v", err)
			} else {
				err = util.SocketSend(c.Conn, byteData)
				if err != nil {
					log.Printf("ocket send error, error is %+v", err)
					return
				}
			}
		case makeStop := <-c.Stop:
			if makeStop {
				break
			}
		}
	}
}
