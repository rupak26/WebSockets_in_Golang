package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	//"golang.org/x/text/message"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
} 

func reader(conn *websocket.Conn) {
	for {
		messageType , p , err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p)) 

		if err := conn.WriteMessage(messageType , p) ; err != nil {
           log.Println(err)
		   return
		}
	}
}


func homepage(w http.ResponseWriter , r *http.Request) {
    fmt.Fprintf(w , "Home Page")
}

func wsEndpoints(w http.ResponseWriter , r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws , err := upgrader.Upgrade(w , r , nil) 
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Successfully Connected")
	reader(ws)
}

func setUpRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/ws", wsEndpoints)
}

func main() {
	
	mux := http.NewServeMux() 
    
    setUpRoutes(mux)
	
	port := ":8080" 
    
	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(port , mux) 

	if err != nil {
		fmt.Println("Server is Facing Issues")
	}

}
