package main

import (
	"time"
	"go-redis/store"
)

func main(){
	store := store.NewStore()

	store.LoadFromDisk("dump.json")
	store.StartCleanup(1 * time.Second)

	go func(){
		for {
			time.Sleep(5 * time.Second)
			store.SaveToDisk("dump.json")
		}
	}()

	StartServer(store)
}