package util

import (
	"fmt"
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