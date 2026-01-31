package httpapi

import (
	"encoding/json"
	"net/http"

	"go-redis/store"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func SetHandler(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		var req struct {
			Key string `json:"key"`
			Val string `json:"val"`
			TTL int `json:"ttl"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
			writeJSON(w, 400, map[string]string{"error":"invalid Json"})
			return
		}

		if req.Key == ""{
			writeJSON(w, 400, map[string]string{"error":"key is required"})
			return
		}

		store.Set(req.Key, req.Val, req.TTL)
		writeJSON(w, 200, map[string]string{"status":"OK"})
	}
}

func GetHandler(store *store.Store) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		key := req.URL.Query().Get("key")

		if key == ""{
			writeJSON(w, 400, map[string]string{"error":"key is required"})
			return
		}

		val, ok := store.Get(key)

		if !ok {
			writeJSON(w, 400, map[string]string{"error":"key not found"})
			return
		}

		writeJSON(w, 200, map[string]string{
			"key":key,
			"value":val,
		})
	}
}

func DelHandler(store *store.Store) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		key := req.URL.Query().Get("key")

		if key == "" {
			writeJSON(w, 400, map[string]string{"error": "key required"})
			return
		}
		
		ok := store.Del(key)

		writeJSON(w, 200, map[string]any{
			"deleted":ok,
		})
	}
}

func ExistsHandler(store *store.Store) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		key := req.URL.Query().Get("key")
		
		if key == "" {
			writeJSON(w, 400, map[string]string{"error": "key required"})
			return
		}

		writeJSON(w, 200, map[string]bool{
			"exists":store.Exists(key),
		})
	}
}

func StatsHandler(store *store.Store) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		writeJSON(w, 200, store.Stats())
	}
}
