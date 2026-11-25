package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// белый список разрешённых IP
var whitelist = map[string]bool{
	"127.0.0.1": true,
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Server started on 127.0.0.1:9000")

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()

	host, _, err := net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		log.Print(err)
		return
	}

	if !whitelist[host] {
		fmt.Println("Отказано в доступе для IP:", host)
		fmt.Fprintln(c, "Доступ запрещён")
		return
	}

	fmt.Println("Подключён клиент:", host)

	s := bufio.NewScanner(c)
	for s.Scan() {
		fmt.Println("Новый запрос с адреса - " + s.Text())

		if _, err := fmt.Fprintln(c, "Запрос на сервер: "+c.LocalAddr().String()+" отправлен"); err != nil {
			log.Print(err)
			return
		}
	}

	if err := s.Err(); err != nil {
		log.Print(err)
	}
}
