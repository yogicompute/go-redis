package main

import (
	"bufio"
	"fmt"
	"go-redis/store"
	"go-redis/util"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){
	store := store.NewStore()

	err := store.LoadFromDisk("dump.json")
	if err == nil{
		fmt.Println("Data loaded from disk.")
	}
	store.StartCleanup(1 * time.Second)

	go func(){
		for {
			time.Sleep(5 * time.Second)
			err := store.SaveToDisk("dump.json")
			if err != nil{
				fmt.Println("SnapShot error : ", err)
			}
		}
	}()

	fmt.Println("MiniRedis started 🚀")
	fmt.Println("Supported commands: SET key value | GET key | DEL key | EXISTS key")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		command, err := util.CommandParser(input)

		switch command.Name{
		case "EXISTS":
			if len(command.Args) != 1{
				fmt.Println("ERR: EXISTS require exactly one key")
				continue
			}
			key := command.Args[0]
			
			exists := store.Exists(key)

			if !exists{
				fmt.Println("0")
			}else{
				fmt.Println("1")
			}
		case "SET":
			if len(command.Args) < 2{
				fmt.Println("ERR: SET requires key and value")
				continue
			}

			key := command.Args[0]
			ttl := 0
			
			if len(command.Args) > 2{
				// SET key value EX seconds
				if len(command.Args) < 4 {
					fmt.Println("ERR: invalid SET syntax")
					continue
				}

				option := strings.ToUpper(command.Args[len(command.Args) - 2])

				if option != "EX"{
					fmt.Println("ERR: only EX option is supported")
					continue
				}

				seconds, err := strconv.Atoi(command.Args[len(command.Args) - 1])
				if err != nil || seconds <= 0 {
					fmt.Println("ERR: invalid TTL value")
					continue
				}

				ttl = seconds
				
				val := strings.Join(command.Args[1:len(command.Args) - 2], " ")
				store.Set(key, val, ttl)
			}else{
				val := strings.Join(command.Args[1:], " ")
				store.Set(key, val, 0)
			}

			fmt.Println("OK")
		case "GET":
			if len(command.Args) != 1{
				fmt.Println("ERR: GET requires exactly one key")
				continue
			}

			key := command.Args[0]

			val, exists := store.Get(key)

			if !exists {
				fmt.Println("(nil)")
			}else{
				fmt.Println(val)
			}
		case "DEL":
			if len(command.Args) != 1{
				fmt.Println("ERR: DEL requires exactly one key")
				continue
			}
			
			key := command.Args[0]

			exists := store.Del(key)

			if !exists {
				fmt.Println("0")
			}else{
				fmt.Println("1")
			}
		default:
			fmt.Println("Unkown Command")
		}
	}
}