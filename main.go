package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/api/chat/v1"
	"google.golang.org/appengine" // Required external App Engine library
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

// Message Object in GChat Payload
type Message struct {
	Text string `json:"text"`
}

// Payload send from GChat
type Payload struct {
	Message Message `json:"message"`
	Space   Space   `json:"space"`
}

type Space struct {
	Name string `json:"name"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ctx := appengine.NewContext(r)

	b, e := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if e != nil {
		json.NewEncoder(w).Encode(e)
	}
	log.Infof(ctx, "Body: %+v", string(b))
	var message Payload
	json.Unmarshal(b, &message)

	if message.Space.Name == "spaces/AAAA0c_TyMI" {
		log.Infof(ctx, "Sending to Bot Dev Room")
		msg, err := postToBotDev(ctx, bytes.NewReader(b))
		if err != nil {
			json.NewEncoder(w).Encode(chat.Message{Text: "An error has occurred"})
		}
		log.Infof(ctx, "Returned from Bot Dev Room: %+v", msg)
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Infof(ctx, "Sending to Bot Development")
		msg, err := postToBotDevelopment(ctx, bytes.NewReader(b))
		if err != nil {
			json.NewEncoder(w).Encode(chat.Message{Text: "An error has occurred"})
		}
		log.Infof(ctx, "Returned from Bot Development: %+v", msg)
		json.NewEncoder(w).Encode(msg)
	}

}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Printf("Testing")
	appengine.Main() // Starts the server to receive requests
}

func postToBotDev(ctx context.Context, body io.Reader) (chat.Message, error) {
	msg, err := postToRoom(ctx, "https://us-central1-uplifted-elixir-203119.cloudfunctions.net/helpBot", body)
	if err != nil {
		log.Infof(ctx, "An Error Occurred: %+v \nMessage: &+v", err, msg)
		return msg, err
	}
	return msg, nil

}
func postToBotDevelopment(ctx context.Context, body io.Reader) (chat.Message, error) {
	msg, err := postToRoom(ctx, "https://bitmoji-bot-dot-uplifted-elixir-203119.appspot.com", body)
	if err != nil {
		log.Infof(ctx, "An Error Occurred: %+v", err)
		log.Infof(ctx, "Message: &+v", msg)
		return msg, err
	}
	return msg, nil
}

func postToRoom(ctx context.Context, url string, body io.Reader) (chat.Message, error) {
	var br chat.Message

	client := urlfetch.Client(ctx)
	resp, err := client.Post(url, "application/json; charset=utf-8", body)
	if err != nil {
		log.Infof(ctx, "Error In Post to Room %+v", err)
		return br, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &br)
	if err != nil {
		return br, err
	}
	return br, nil

}
