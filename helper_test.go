package vk

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type aStringerTypeForTesting int

func (x aStringerTypeForTesting) String() string {
	return fmt.Sprint(int(x))
}

func TestGoStr(t *testing.T) {
	type args struct {
		x interface{}
	}
	qwerty := [...]byte{'Q', 'W', 'E', 'R', 'T', 'Y', 0}
	stringer := aStringerTypeForTesting(-123)
	tests := []struct {
		name string
		args args
		want string
	}{
		{"a string", args{"asdf"}, "asdf"},
		{"a byte array", args{qwerty}, "QWERTY"},
		{"pointer to a byte array", args{&qwerty}, "QWERTY"},
		{"a byte pointer", args{(*uint8)(unsafe.Pointer(&qwerty))}, "QWERTY"},
		{"a stringer", args{stringer}, "-123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoStr(tt.args.x); got != tt.want {
				t.Errorf("GoStr() = %q, want %q", got, tt.want)
			}
		})
	}

}

func TestGoStrSlice(t *testing.T) {
	ss := []string{
		"12345",
		"67890",
	}
	pp, n, free := CStrSlice(ss)
	if pp == nil || n != uint32(len(ss)) {
		t.Fatal("CStrSlice failure")
	}
	ss1 := GoStrSlice(pp, n)
	if len(ss1) != len(ss) || !reflect.DeepEqual(ss, ss1) {
		t.Fatal("GoStrSlice failure, ss1=", ss1)
	}
	free()
}

func TestMemCopy(t *testing.T) {
	var dst [100]byte
	var src [100]byte
	for i := 0; i < 100; i++ {
		src[i] = byte(i)
	}
	n := MemCopy(unsafe.Pointer(&dst), unsafe.Pointer(&src), 100)
	if n != 100 {
		t.Fatal("n != 100")
	}
	if src != dst {
		t.Fatalf("src!=dst: src=%v, dst=%v", src, dst)
	}
}
