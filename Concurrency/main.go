package main

import (
	"sync"
	"time"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	wg := new(sync.WaitGroup)
	for number := 2; number <= 12; number++ {
		wg.Add(1)
		go printTable(number, wg)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Wait()
	pl := new(sync.Pool)
	pl.Put(1)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
