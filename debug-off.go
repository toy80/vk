// +build !debug

package vk

func DebugBreakAfterVkCall() {}

func debugCheckAndBreak()           {}
func debugMarkMemBlock(p uintptr)   {}
func debugUnmarkMemBlock(p uintptr) {}

func DumpMemoryLeaks() {}
