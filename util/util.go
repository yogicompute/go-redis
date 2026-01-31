package util

import (
	"fmt"
	"go-redis/store"
	"strconv"
	"strings"
)

type Command struct{
	Name string
	Args []string
}


func CommandParser(input string) (*Command, error){
	input = strings.TrimSpace(input)

	if input == ""{
		return nil, fmt.Errorf("empty command")
	}

	parts := strings.Fields(input)

	cmd := strings.ToUpper(parts[0])
	args := parts[1:]

	return &Command{cmd, args}, nil
}


func ExecuteCommand(cmd *Command, store *store.Store) string {
	switch cmd.Name {
		case "EXISTS":
			if len(cmd.Args) != 1{
				return "ERR: EXISTS require exactly one key"
			}
			key := cmd.Args[0]
			
			exists := store.Exists(key)

			if !exists{
				return "0"
			}else{
				return "1"
		}
		case "SET":
			if len(cmd.Args) < 2{
				return "ERR: SET requires key and value"
			}

			key := cmd.Args[0]
			ttl := 0
			
			if len(cmd.Args) > 2{
				// SET key value EX seconds
				if len(cmd.Args) < 4 {
					return "ERR: invalid SET syntax"
				}

				option := strings.ToUpper(cmd.Args[len(cmd.Args) - 2])

				if option != "EX"{
					return "ERR: only EX option is supported"
				}

				seconds, err := strconv.Atoi(cmd.Args[len(cmd.Args) - 1])
				if err != nil || seconds <= 0 {
					return "ERR: invalid TTL value"
				}

				ttl = seconds
				
				val := strings.Join(cmd.Args[1:len(cmd.Args) - 2], " ")
				store.Set(key, val, ttl)
			}else{
				val := strings.Join(cmd.Args[1:], " ")
				store.Set(key, val, 0)
			}

			return "OK"
		case "GET":
			if len(cmd.Args) != 1{
				return "ERR: GET requires exactly one key"
			}

			key := cmd.Args[0]

			val, exists := store.Get(key)

			if !exists {
				return "(nil)"
			}else{
				return val
			}
		case "DEL":
			if len(cmd.Args) != 1{
				return "ERR: DEL requires exactly one key"
			}
			
			key := cmd.Args[0]

			exists := store.Del(key)

			if !exists {
				return "0"
			}else{
				return "1"
			}
		default:
			return "Unkown Command"
		}
}