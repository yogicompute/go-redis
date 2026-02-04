package store

import (
	"maps"
	"encoding/json"
	"os"
	"sync"
	"time"
)



type Store struct{
	data map[string]Item
	mu sync.Mutex
}

type Item struct{
	Value string
	ExpiresAt time.Time
}

type Snapshot struct {
	Data map[string]Item `json:"data"`
}

// <---------PERSISTENCE-----------> 

func (s *Store) SaveToDisk(filename string) error{
	s.mu.Lock()
	defer s.mu.Unlock()

	snap := Snapshot{
		Data: s.data,
	}

	bytes, err := json.MarshalIndent(snap, "", "  ")
	if err != nil{
		return err
	}
	return os.WriteFile(filename, bytes, 0644) //Read+Write
}

func (s *Store) LoadFromDisk(filename string) error{
	bytes, err := os.ReadFile(filename)

	if err != nil{
		return err
	}

	var snap Snapshot
	if err := json.Unmarshal(bytes, &snap); err != nil{
		return err;
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = snap.Data
	return nil
}

// <---------CLEANUP-----------> 

func (s *Store) StartCleanup(interval time.Duration){
	go func(){
		for{
			time.Sleep(interval)

			now := time.Now()

			s.mu.Lock()
			for key, item := range s.data{
				if !item.ExpiresAt.IsZero() && now.After(item.ExpiresAt){
					delete(s.data, key)
				}
			}
			s.mu.Unlock()
		}
	}()
}

func (s *Store) Stats() map[string]int{
	s.mu.Lock()
	defer s.mu.Unlock()

	return map[string]int {
		"keys":len(s.data),
	}
}

func NewStore() *Store{
	return &Store{
		data: make(map[string]Item),
	}
}

func (s *Store) Set(key, val string, ttlSeconds int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var expiresAt time.Time
	if ttlSeconds > 0 {
		expiresAt = time.Now().Add(time.Duration(ttlSeconds) * time.Second)
	}

	s.data[key] = Item{
		val,
		expiresAt,
	}
}

func (s *Store) Get(key string)(string, bool){
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.data[key]

	if !ok {
		return "", false
	}

	if !item.ExpiresAt.IsZero() && time.Now().After(item.ExpiresAt){
		delete(s.data, key)
		return "", false
	}
	return item.Value, true
}

func (s* Store) Del(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; ok{
		delete(s.data, key)
		return true
	}
	return false
}

func (s *Store) Exists(key string) bool{
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]
	return ok
}

func (s *Store) Export() map[string]Item{
	s.mu.Lock()
	defer s.mu.Unlock()

	copy := make(map[string]Item)
	maps.Copy(copy, s.data)
	return copy
}

func (s *Store) Import(data map[string]Item){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = data
}