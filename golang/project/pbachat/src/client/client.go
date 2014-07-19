package client

type client struct {
	Id   string
	Name string
}

type message struct {
	Mtype    string
	FromId   string
	FromName string
	ToId     string
	ToName   string
	Content  []byte
}

//func NewClient(connType int, sClient ) (&client Client) {

//}

//关闭用户连接，并且从用户数组中移除该用户
//func (client *Client) Close() {
//	switch client.ConnType {
//	case 1:
//		client.Conn.Close()

//		for entry := ClientList.Front(); entry != nil; entry = entry.Next() {
//			if client.Conn == entry.Value.(Client).Conn {
//				common.Log("remove from clientList", client.Conn)
//				ClientList.Remove(entry)
//			}
//		}
//	case 2:
//		client.WsConn.Close()

//		for entry := ClientList.Front(); entry != nil; entry = entry.Next() {
//			if client.WsConn == entry.Value.(Client).WsConn {
//				common.Log("remove from customer service clientList", client.WsConn)
//				ClientList.Remove(entry)
//			}
//		}
//	}

//	common.Log(ClientList.Len())
//}

//func (client *Client) SendToAll(msg []byte) {
//	for entry := ClientList.Front(); entry != nil; entry = entry.Next() {
//		c := entry.Value.(Client)
//		//common.Log("conntype: ", c.ConnType)
//		//common.Log("conn: ", c.Conn)
//		//common.Log("wsconn: ", c.WsConn)

//		switch c.ConnType {
//		case 1:
//			common.Log("send msg to socket: ", c.Conn)
//			c.Conn.Write(msg)
//		case 2:
//			common.Log("send msg to websocket: ", c.WsConn)
//			c.WsConn.Write(msg)
//		}
//	}
//}
