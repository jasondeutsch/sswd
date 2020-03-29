package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", ws)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// todo obvi
		bs, err := run(string(msg))
		log.Printf("recv: %s", msg)

		err = c.WriteMessage(mt, bs)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func run(cmdStr string) ([]byte, error) {
	cmdSlc := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmdSlc[0], cmdSlc[1:]...)

	return cmd.CombinedOutput()
}
