package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	conn       *net.UDPConn
	gkey       bool
	userID     int
	userName   string
	sendMsg    chan string
	receiveMsg chan string
}

func (c *Client) functosend(sid int, msg string) {

	str := fmt.Sprintf("###%d##%d##%s##%s###", sid, c.userID, c.userName, msg)
	_, err := c.conn.Write([]byte(str))
	checkError(err, "functosend")
}

func (c *Client) sendmsg() {
	for c.gkey {
		msg := <-c.sendMsg
		str := fmt.Sprintf("###2##%d##%s##%s###", c.userID, c.userName, msg)
		_, err := c.conn.Write([]byte(str))
		checkError(err, "sendmsg")
	}
}
func (c *Client) receivemsg() {
	var buf [512]byte
	for c.gkey {
		n, err := c.conn.Read(buf[0:])
		checkError(err, "receivemsg")
		c.receiveMsg <- string(buf[0:n])
	}
}

func (c *Client) getmsg() {
	var msg string
	for c.gkey {
		fmt.Println("msg: ")
		_, err := fmt.Scanln(&msg)
		checkError(err, "getmsg")
		if msg == ": quit" {
			c.gkey = false
		} else {
			c.sendMsg <- encodemsg(msg)
		}
	}
}
func (c *Client) printMessage() {
	//var msg string
	for c.gkey {
		msg := <-c.receiveMsg
		fmt.Println(msg)
	}
}

func encodemsg(msg string) string {
	return strings.Join(strings.Split(strings.Join(strings.Split(msg, "\\"), "\\\\"), "#"), "\\#")

}

func checkError(e error, funcname string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s ----in func %s", e.Error(), funcname)
		os.Exit(1)
	}

}
func nowTime() string {
	return time.Now().String()
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err, "main")
	var c Client
	c.gkey = true
	c.sendMsg = make(chan string)
	c.receiveMsg = make(chan string)

	fmt.Println("input id: ")
	_, err = fmt.Scanln(&c.userID)
	checkError(err, "main")
	fmt.Println("input name: ")
	_, err = fmt.Scanln(&c.userName)
	checkError(err, "main")

	c.conn, err = net.DialUDP("udp", nil, udpAddr)

	checkError(err, "main")
	//fmt.Println(c)
	defer c.conn.Close()

	//发送进入聊天室消息,类型1，###1##uid##uName##进入聊天室###
	//messagestr := fmt.Sprintf("###1##%d##%s###", c.userID, c.userName)
	//_,err = c.conn.Write([]byte(messagestr))
	//checkError(err)
	c.functosend(1, c.userName+"进入聊天室")

	//go c.getMessage()
	go c.printMessage()
	go c.receivemsg()

	go c.sendmsg()
	c.getmsg()

	c.functosend(3, c.userName+"离开聊天室")
	fmt.Println("退出成功!")

	os.Exit(0)
}
