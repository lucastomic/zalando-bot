package main

import "github.com/lucastomic/zalando-bot/internals/logger"

func main() {
<<<<<<< HEAD
	logger.SignIn()
=======

	it, _ := proxy.NewIterator()
	it.RefreshProxies()
	for i := 0; i < 5; i++ {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				logger.SignIn("lucastomic17@gmail.com", ".") //removed password to add in github

				wg.Done()
			}()
		}
		// time.Sleep(time.Hour)

		wg.Wait()
	}
>>>>>>> b12de8f446183d8f523770e816a3b84ff2468db5
}
