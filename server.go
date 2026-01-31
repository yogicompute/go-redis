package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"go-redis/store"
	"go-redis/util"
)

func handleClient(conn net.Conn, store *store.Store){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writter := bufio.NewWriter(conn)

	for{
		line, err := reader.ReadString('\n')
		if err != nil{
			return //client disconnected
		}

		line = strings.TrimSpace(line)

		cmd, err := util.CommandParser(line)
		if err != nil{
			writter.WriteString("ERR invalid command\n")
			writter.Flush()
			continue
		}

		res := util.ExecuteCommand(cmd, store)
		writter.WriteString(res +"\n")
		writter.Flush()
	}
}

func StartServer(store *store.Store){
	listener, err := net.Listen("tcp", ":6379")

	if err != nil{
		panic(err)
	}

	fmt.Println("Mini redis listening on port 6379")

	for{
		conn, err := listener.Accept()

		if err != nil{
			continue
		}

		go handleClient(conn, store)
	}
}