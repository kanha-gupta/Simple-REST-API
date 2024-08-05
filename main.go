package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items  = make(map[int]Item)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/api/cmd", handler)
	fmt.Print("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addItem(w, r)
	case http.MethodGet:
		getItems(w)
	case http.MethodDelete:
		deleteItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	item.ID = nextID
	nextID++
	items[item.ID] = item
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func getItems(w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()

	itemList := make([]Item, 0, len(items))
	for _, item := range items {
		itemList = append(itemList, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemList)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if _, exists := items[id]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	delete(items, id)
	w.WriteHeader(http.StatusNoContent)
}
