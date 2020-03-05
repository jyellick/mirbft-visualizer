package main

import (
	"fmt"
	"time"

	"github.com/vugu/vugu"
)

type Client struct {
	MirNodes   []*MirNode
	EventQueue *EventQueue
	TargetNode int
}

func (c *Client) SetTargetNode(event *vugu.DOMEvent) {
	target := event.JSEvent().Get("target").Get("value").String()
	n, err := fmt.Sscanf(target, "%d", &c.TargetNode)
	if n == 0 || err != nil {
		panic("bad selection")
	}
}

func (c *Client) SendRequest(event *vugu.DOMEvent) {
	fmt.Println("Submitting request")
	c.MirNodes[c.TargetNode].Propose(c.EventQueue, []byte(time.Now().String()))
}
