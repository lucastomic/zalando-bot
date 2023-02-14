package main

import (
	"sync"

	"github.com/lucastomic/zalando-bot/internals/logger"
	"github.com/lucastomic/zalando-bot/internals/proxy"
)

func main() {

	it, _ := proxy.NewIterator()
	it.RefreshProxies()
	for i := 0; i < 5; i++ {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				logger.SignIn("lucastomic17@gmail.com", "94039155")

				wg.Done()
			}()
		}
		// time.Sleep(time.Hour)

		wg.Wait()
	}
}
