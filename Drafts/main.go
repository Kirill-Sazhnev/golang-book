package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher()           // Начало горутины
	time.Sleep(4 * time.Second) // Ожидание храпа гофера
} // Здесь все горутины останавливаются

func sleepyGopher() {
	time.Sleep(3 * time.Second) // гофер спит
	fmt.Println("... snore ...")
}
