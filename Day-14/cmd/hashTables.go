package main

import "fmt"

func HashTables() {
	fmt.Println("\nHash Map Basics: Built-in Hapmaps")

	hashmap := make(map[string]int)

	hashmap["apple"] = 1
	hashmap["banana"] = 2

	fmt.Println("Apple:", hashmap["apple"])
	fmt.Println("Banana:", hashmap["banana"])
}

/*
Implement a custom hash map by creating a struct with a hash function.
*/
type HashMap struct {
	data map[int][]KeyValue
}

type KeyValue struct {
	Key   string
	Value int
}

func hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = 31*hash + int(char)
	}
	return hash
}

// Put method to add a key-value pair to the hashmap
func (h *HashMap) Put(key string, value int) {
	if h.data == nil {
		h.data = make(map[int][]KeyValue)
	}
	index := hash(key) % 100
	for i, kv := range h.data[index] {
		if kv.Key == key {
			h.data[index][i].Value = value // Update existing key's value
			return
		}
	}
	h.data[index] = append(h.data[index], KeyValue{Key: key, Value: value})
}

// get method to retrieve a value by key
func (h *HashMap) Get(key string) (int, bool) {
	index := hash(key) % 100
	for _, kv := range h.data[index] {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return 0, false
}

// Delete method to remove a key-value pair from the hashmap
func (h *HashMap) Delete(key string) bool {
	index := hash(key) % 100
	for i, kv := range h.data[index] {
		if kv.Key == key {
			h.data[index] = append(h.data[index][:i], h.data[index][i+1:]...)
			return true // Key found and deleted
		}
	}
	return false // Key not found
}

func CustomHashTable() {
	fmt.Println("\nCreating A Custom Hash Table")

	// create empty hashmap
	hashmap := HashMap{}

	// insert key-value pairs
	hashmap.Put("apple", 1)
	hashmap.Put("banana", 2)
	hashmap.Put("grape", 3)

	// get values
	val, exists := hashmap.Get("banana")
	if exists {
		fmt.Println("Value for 'banana':", val)
	} else {
		fmt.Println("Key 'banana' not found")
	}

	// update value for an existing key
	hashmap.Put("banana", 22)
	val, _ = hashmap.Get("banana")
	fmt.Println("Updated value for 'banana':", val)

	// delete a key
	if hashmap.Delete("apple") {
		fmt.Println("Deleted key 'apple'")
	}

	// check if 'apple' exists after deletion
	_, exists = hashmap.Get("apple")
	if !exists {
		fmt.Println("Key 'apple' not found after deletion")
	}
}
