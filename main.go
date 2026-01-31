package main

import (
	"net/http"
	"time"
	
	"go-redis/httpapi"
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

	go func(){
		http.Handle("/", http.FileServer(http.Dir("./web")))	

		http.HandleFunc("/set", httpapi.SetHandler(store))
		http.HandleFunc("/get", httpapi.GetHandler(store))
		http.HandleFunc("/del", httpapi.DelHandler(store))
		http.HandleFunc("/exists", httpapi.ExistsHandler(store))
		http.HandleFunc("/stats", httpapi.StatsHandler(store))

		http.ListenAndServe(":8080", nil)
	}()

	StartServer(store)
}