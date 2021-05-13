package mysql

import (
	"fmt"
	"sync"
	"testing"
)

func TestD(t *testing.T) {
	var wg = sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			testDB()
		}()
	}
	fmt.Println("dsfs")
	wg.Wait()
}
