package study

import (
	"encoding/json"
	"fmt"
	"net/http"
	"dog/command/console"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func init() {
	c := console.Console{Signature: "cmdWebSocket", Description: "this is a template", Handle: cmdWebSocket}
	commandList[c.Signature] = c
}

func cmdWebSocket() {
	fs := http.FileServer(http.Dir("./web"))

	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	router.Handle("/", fs)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
	// Configure websocket route

}

var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.data.IP = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = userList
			dataB, _ := json.Marshal(c.data)
			c.sc <- dataB
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b:
			for c := range h.c {
				select {
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}

// data

type cmdWebSocketData struct {
	IP       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"userList"`
}

// connection

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *cmdWebSocketData
}

var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &cmdWebSocketData{}}
	h.r <- c
	go c.writer()
	c.reader()
	defer func() {
		c.data.Type = "logout"
		userList = del(userList, c.data.User)
		c.data.UserList = userList
		c.data.Content = c.data.User
		dataB, _ := json.Marshal(c.data)
		h.b <- dataB
		h.r <- c
	}()
}

func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var userList = []string{}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			userList = append(userList, c.data.User)
			c.data.UserList = userList
			dataB, _ := json.Marshal(c.data)
			h.b <- dataB
		case "user":
			c.data.Type = "user"
			dataB, _ := json.Marshal(c.data)
			h.b <- dataB
		case "logout":
			c.data.Type = "logout"
			userList = del(userList, c.data.User)
			dataB, _ := json.Marshal(c.data)
			h.b <- dataB
			h.r <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var nSlice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			nSlice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(nSlice)
	return nSlice
}
