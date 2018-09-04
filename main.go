package main

import (
	"bytes"
	"log"
	"net"
	"net/mail"

	"github.com/mhale/smtpd"
)

func mailHandler(origin net.Addr, from string, to []string, data []byte) {
	log.Println("msg received")
	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")

	buf := new(bytes.Buffer)
	buf.ReadFrom(msg.Body)

	log.Printf("Received mail from '%s' for '%s' with subject '%s'", from, to[0], subject)
	log.Printf("Body: '%v'", buf.String())
}

func main() {
	smtpd.ListenAndServe("127.0.0.1:2525", mailHandler, "MyServerApp", "")
}