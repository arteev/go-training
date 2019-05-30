package main

import (
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	servers := []string{"127.0.0.1:11211"}
	m := memcache.New(servers...)

	err := m.Add(&memcache.Item{
		Key:        "key_1",
		Expiration: 5,
		Value:      []byte("value1"),
	})
	if err != nil {
		log.Fatal(err)
	}

	item, err := m.Get("key_1")
	if err != nil {
		log.Fatalf("get: %v", err)
	}
	fmt.Println("get item: ", string(item.Value))

	m.Set(&memcache.Item{
		Key:   "inc",
		Value: []byte("100"),
	})

	value, err := m.Increment("inc", 1)
	if err != nil {
		log.Fatalf("inc: %v", err)
	}
	fmt.Printf("inc: %v \n", value)

	err = m.Delete("inc")
	if err != nil {
		log.Fatalf("delete: %v", err)
	}

	_, err = m.Get("inc")
	if err != memcache.ErrCacheMiss {
		log.Fatal("Excpected ErrCacheMiss")
	}
}
