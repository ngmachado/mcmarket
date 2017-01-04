package mcmarket

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/ngmachado/mcmarket/Entities"
	"github.com/ngmachado/mkm"
)

type Message struct {
	cl *mkm.Client
}

//Get the message thread overview for the authenticated user
func (r *Message) All() (*Entities.ThreadPool, error) {
	body, err := r.cl.Request(mkm.Get, "/account/messages", nil)
	if err != nil {
		return nil, err
	}
	messages := Entities.ThreadPool{}
	err = json.Unmarshal(body, &messages)
	return &messages, err
}

//Get a complete message thread for messages between the authenticated user and another user
func (r *Message) From(IDOtherUser int) (*Entities.Thread, error) {
	resource := fmt.Sprintf("/account/messages/%d", IDOtherUser)
	body, err := r.cl.Request(mkm.Get, resource, nil)
	if err != nil {
		return nil, err
	}
	messages := Entities.Thread{}
	err = json.Unmarshal(body, &messages)
	return &messages, err
}

//Returns a specified message with a specified other user
func (r *Message) FromWith(IDOtherUser int, IDMessage int) (*Entities.Thread, error) {
	resource := fmt.Sprintf("/account/messages/%d/%d", IDOtherUser, IDMessage)
	body, err := r.cl.Request(mkm.Get, resource, nil)
	if err != nil {
		return nil, err
	}
	messages := Entities.Thread{}
	err = json.Unmarshal(body, &messages)
	return &messages, err
}

//Write a message to another user
func (r *Message) Send(IDOtherUser int, message string) (*Entities.ThreadView, error) {
	resource := fmt.Sprintf("/account/messages/%d", IDOtherUser)
	postData := []byte(fmt.Sprintf("%s<request><message>%s</message></request>", xml.Header, message))
	messages := Entities.ThreadView{}
	body, err := r.cl.Request(mkm.Post, resource, postData)
	err = json.Unmarshal(body, &messages)
	return &messages, err
}

//Delete all messages between the authenticated and another user
func (r *Message) Delete(IDOtherUser int) error {
	resource := fmt.Sprintf("/account/messages/%d", IDOtherUser)
	_, err := r.cl.Request(mkm.Delete, resource, nil)
	return err
}

//Delete a single message between the authenticated and another user
func (r *Message) DeleteFrom(IDOtherUser int, IDMessage int) error {
	resource := fmt.Sprintf("/account/messages/%d/%d", IDOtherUser, IDMessage)
	_, err := r.cl.Request(mkm.Delete, resource, nil)
	return err
}
