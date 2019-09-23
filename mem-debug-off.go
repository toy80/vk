// +build !memdbg

package vk

func dbgMemAlloc(p uintptr) {
	// do nothing, optimize out
}

func dbgMemFree(p uintptr) {
	// do nothing, optimize out
}

func DumpMemoryLeaks() {
	// do nothing
}
