package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{}
var todoList []string

func getCmd(input string) string {
	inputArr := strings.Split(input, " ")
	return inputArr[0]
}

func getMessage(input string) string {
	inputArr := strings.Split(input, " ")
	var result string
	for i := 1; i < len(inputArr); i++ {
		result += inputArr[i]
	}
	return result
}

func updateTodoList(input string) {
	tmpList := todoList
	todoList = []string{}

	for _, val := range tmpList {
		if val == input {
			continue
		}
		todoList = append(todoList, val)
	}
}

func main() {
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read failed:", err)
				break
			}

			input := string(message)
			cmd := getCmd(input)
			msg := getMessage(input)

			if cmd == "finish" {
				todoList = append(todoList, msg)
			} else if cmd == "succeed" {
				updateTodoList(msg)
			}

			output := "Current Todos: \n"
			for _, todo := range todoList {
				output += "\n - " + todo + "\n"
			}
			output += "\n----------------------------------------"

			message = []byte(output)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
