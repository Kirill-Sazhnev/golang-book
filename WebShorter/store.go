package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}

type record struct {
	Key, URL string
}

type ProxyStore struct {
	client *rpc.Client
	store  *URLStore
}

type Store interface {
	Get(key, url *string) error
	Put(url, key *string) error
}

const saveQueueLength = 100

func NewURLStore(name string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}

	if name != "" {
		s.save = make(chan record, saveQueueLength)
		if err := s.load(name); err != nil {
			log.Fatal("Error loading store file: ", err)
		}
		go s.saveLoop(name)
	}
	return s
}

func (s *URLStore) saveLoop(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error with the store file: ", err)
	}
	defer f.Close()
	e := json.NewEncoder(f)

	for {
		rec := <-s.save
		if err := e.Encode(rec); err != nil {
			log.Println("Error saving to file: ", err)
		}
	}
}

func (s *URLStore) load(name string) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error with the store file: ", err)
	}
	defer f.Close()
	d := json.NewDecoder(f)

	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(&r.Key, &r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *URLStore) Set(key, url *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.urls[*key]; ok {
		return errors.New("the key is already present in map")
	}

	s.urls[*key] = *url
	return nil
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)

}

func (s *URLStore) Get(key, url *string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if u, ok := s.urls[*key]; ok {
		*url = u
		return nil
	}
	return errors.New("the key is not found")
}

func (s *URLStore) Put(url, key *string) error {
	for {
		*key = genKey(s.Count())
		if s.Set(key, url) == nil {
			break
		}
	}
	rec := record{
		Key: *key,
		URL: *url,
	}
	if s.save != nil {
		s.save <- rec
	}
	return nil
}

func NewProxyStore(adr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", adr)
	if err != nil {
		log.Println("Error dialHTTP: ", err)
	}
	return &ProxyStore{client, NewURLStore("")}
}

func (s *ProxyStore) Get(key, url *string) error {
	if err := s.store.Get(key, url); err == nil {
		return nil
	}
	if err := s.client.Call("Store.Get", key, url); err != nil {
		return err
	}
	return s.store.Set(key, url)
}

func (s *ProxyStore) Put(url, key *string) error {
	if err := s.client.Call("Store.Put", url, key); err == nil {
		return s.store.Set(key, url)
	} else {
		return err
	}
}
