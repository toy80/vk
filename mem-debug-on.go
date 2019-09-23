// +build memdbg

package vk

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var (
	dbgMemBlocks = make(map[uintptr]string)
	dbgMemMutex  sync.Mutex
)

func dbgMemAlloc(p uintptr) {
	dbgMemMutex.Lock()
	defer dbgMemMutex.Unlock()

	var loc string
	for i := 2; i < 5; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !strings.Contains(file, "/vk/") {
			loc = fmt.Sprintf("%s:%d", file, line)
			break
		}
	}
	if loc == "" {
		loc = "unkown caller"
	}
	if p == 0 {
		panic(fmt.Sprintf("%s: MemAlloc() returns nil", loc))
	}
	dbgMemBlocks[p] = loc
}

func dbgMemFree(p uintptr) {
	dbgMemMutex.Lock()
	defer dbgMemMutex.Unlock()

	if loc, ok := dbgMemBlocks[p]; ok {
		if loc == "" {
			panic(fmt.Sprintf("%s: MemFree(0x%X) double free", loc, p))
		}
		dbgMemBlocks[p] = ""
	}
}

func DumpMemoryLeaks() {
	dbgMemMutex.Lock()
	defer dbgMemMutex.Unlock()
	any := false
	for _, loc := range dbgMemBlocks {
		if loc != "" {
			if !any {
				any = true
				fmt.Println("------- MEMORY LEAKS DETECTED -------")
			}
			fmt.Println(loc)
		}
	}
	if any {
		fmt.Println("------- END DUMP MEMORY LEAKS -------")
	}

}
