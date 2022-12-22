package main

import (
	"fmt"
	"sync"
	"time"
) // Импортирует пакет sync

var mu sync.Mutex // Объявляет мьютекс

type Visited struct {
	// mu охраняет карту с посещенными страницами
	mu      sync.Mutex     // Объявление мьютекса
	visited map[string]int // Объявление карты из URL (строки) ключей к значениям integer
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()         // Блокирует мьютекс
	defer v.mu.Unlock() // Убедитесь, что мьютекс разблокирован
	count := v.visited[url]
	count++
	v.visited[url] = count // Обновляет карту
	return count
}

func main() {
	web := &Visited{
		visited: map[string]int{},
	}

	go web.VisitLink("4pda.ru")
	go web.VisitLink("rain.tv")
	fmt.Println(web.visited)
	go web.VisitLink("4pda.ru")
	go web.VisitLink("rain.tv")
	fmt.Println(web.visited)
	go web.VisitLink("youtube.com")
	time.Sleep(1 * time.Second)
	fmt.Println(web.visited)

}
