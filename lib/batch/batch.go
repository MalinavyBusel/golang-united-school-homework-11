package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	idPasser := make(chan int64)
	var resultMutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(int(n))

	for i := int64(0); i < pool; i++ {
		go func() {
			for { // цикл бесконечнй, тк сам прервется после нужного кол-ва выполнений wg.Done()
				currentUser := <-idPasser
				nextUser := getOne(currentUser)

				resultMutex.Lock()
				res = append(res, nextUser)
				resultMutex.Unlock()

				wg.Done() // считает за выполнение один созданный юзер, а не законченную горутину
			}
		}()
	}

	for j := int64(0); j < n; j++ {
		idPasser <- j
	}

	wg.Wait()
	return res
}
