package main

import (
	"encoding/json"
	"io"
	"log"
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

const saveQueueLength = 100

func NewURLStore(name string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}

	if err := s.load(name); err != nil {
		log.Fatal("Error loading store file: ", err)
	}
	go s.saveLoop(name)
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
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.urls[key]; ok {
		return false
	}

	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)

}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			s.save <- record{key, url}
			return key
		}
	}
}
