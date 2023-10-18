package Conc

import (
	"fmt"
	"github.com/Pedroxer/wbl_l1/util"
	"sync"
)

type SafeMap struct {
	Rw     sync.RWMutex
	DefMap map[string]string
}

func (sf *SafeMap) Set(key, value string, wg *sync.WaitGroup, i int) {
	sf.Rw.Lock()
	sf.DefMap[key] = value
	sf.Rw.Unlock()
	wg.Done()
}

func ConcInsertMap() {
	var wg sync.WaitGroup
	sf := &SafeMap{
		DefMap: make(map[string]string)}
	for i := 0; i < 5; i++ {
		key := util.RandomString(3)
		value := util.RandomString(10)
		wg.Add(1)
		go sf.Set(key, value, &wg, i)

	}
	wg.Wait()
	fmt.Println(sf.DefMap)
}
