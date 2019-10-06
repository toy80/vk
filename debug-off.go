// +build !debug

package vk

func debugMarkMemBlock(p uintptr)   {}
func debugUnmarkMemBlock(p uintptr) {}

func DumpMemoryLeaks() {}
