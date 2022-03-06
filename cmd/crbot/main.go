package main

import (
	"sync"

	"github.com/daniel-hutao/crbot/internal/pkg/ferry"
	"github.com/daniel-hutao/crbot/internal/pkg/fsbot"
	"github.com/daniel-hutao/crbot/internal/pkg/ghbot"
)

func main() {
	var wg sync.WaitGroup

	fsBot := fsbot.NewBot()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fsBot.Run()
	}()

	err := ghbot.GetMsg(ferry.GlobalMessageChan)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
