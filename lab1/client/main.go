package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// отправляем свой адрес серверу
	if _, err := fmt.Fprintln(c, c.LocalAddr()); err != nil {
		log.Print(err)
		return
	}

	// читаем ответ сервера
	msg, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Ответ сервера: ", msg)
}
