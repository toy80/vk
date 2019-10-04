package vk

import (
	"fmt"
	"reflect"
	"unsafe"
)

func CStrOrNil(s string) (c *int8, free func()) {
	if s == "" {
		return nil, func() {}
	}
	return CStr(s)
}

func CStr(s string) (c *int8, free func()) {
	n := len(s) + 1
	p := MemAlloc(uintptr(n))
	if p == nil {
		panic("failed to allocate unmanaged memory")
	}

	a := (*((*[0x7FFFFFFF]byte)(p)))[:n]
	copy(a, []byte(s))
	a[n-1] = 0

	return (*int8)(p), func() {
		MemFree(p)
	}
}

func CStrSliceOrNil(ss []string) (c **int8, n uint32, free func()) {
	if len(ss) == 0 {
		return nil, 0, func() {}
	}
	return CStrSlice(ss)
}

func CStrSlice(ss []string) (c **int8, n uint32, free func()) {
	n = uint32(len(ss))
	var frees []func()
	p := MemAlloc(uintptr(n) * unsafe.Sizeof((*int8)(nil)))
	if n > 0 {
		if p == nil {
			panic("failed to allocate unmanaged memory")
		}
		a := (*((*[0x7FFFFFFF]*int8)(p)))[:n]
		frees = make([]func(), n)
		for i := range a {
			a[i], frees[i] = CStr(ss[i])
		}
	}
	return (**int8)(p), n, func() {
		if p != nil {
			MemFree(p)
			for _, f := range frees {
				f()
			}
		}
	}
}

func ptrInt8ToString(p *int8) string {
	if p == nil || *p == 0 {
		return ""
	}
	usp := unsafe.Pointer(p)
	pp := (*[0x7FFFFFFF]byte)(usp)
	s := pp[:]
	for i, v := range s {
		if v == 0 {
			tmp := make([]byte, i)
			copy(tmp, s)
			return string(tmp)
		}
	}
	return ""
}

func ptrUint8ToString(p *uint8) string {
	if p == nil || *p == 0 {
		return ""
	}
	s := (*(*[0x7FFFFFFF]byte)(unsafe.Pointer(p)))[:]
	for i, v := range s {
		if v == 0 {
			tmp := make([]byte, i)
			copy(tmp, s)
			return string(tmp)
		}
	}
	return ""
}

func GoStr(x interface{}) string {
	if i, ok := x.(interface{ String() string }); ok {
		return i.String()
	}
	switch y := x.(type) {
	case nil:
		return ``
	case string:
		return y
	case *int8:
		return ptrInt8ToString(y)
	case *uint8:
		return ptrUint8ToString(y)
	default:
		v := reflect.ValueOf(x)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return ""
			}
			v = v.Elem()
		}
		if v.Kind() == reflect.Array && v.Type().Elem().Kind() == reflect.Int8 {
			var t []byte
			for i := 0; i < v.Len(); i++ {
				c := byte(int8(v.Index(i).Int()))
				if c == 0 {
					break
				}
				t = append(t, c)
			}
			return string(t)
		} else if v.Kind() == reflect.Array && v.Type().Elem().Kind() == reflect.Uint8 {
			var t []byte
			for i := 0; i < v.Len(); i++ {
				c := byte(v.Index(i).Uint())
				if c == 0 {
					break
				}
				t = append(t, c)
			}
			return string(t)
		}
		panic("unsupported type: " + reflect.TypeOf(x).String())
	}
}

func GoStrSlice(pp **int8, n uint32) (ss []string) {
	if pp == nil {
		return
	}
	ss = make([]string, 0, n)
	a := (*((*[0x7FFFFFFF]*int8)(unsafe.Pointer(pp))))[:n]
	for _, p := range a {
		ss = append(ss, GoStr(p))
	}
	return
}

func CUint32ArrayOrNil(s []uint32) (c *uint32, n uint32, free func()) {
	if len(s) == 0 {
		return nil, 0, func() {}
	}
	return CUint32Array(s)
}

func CUint32Array(s []uint32) (c *uint32, n uint32, free func()) {
	n = uint32(len(s))
	p := MemAlloc(uintptr(n) * 4)
	if p == nil {
		panic("failed to allocate unmanaged memory")
	}
	copy((*((*[0x7FFFFFFF]uint32)(p)))[:n], s)
	return (*uint32)(p), n, func() {
		MemFree(p)
	}
}

func CByteArrayOrNil(s []byte) (c *byte, n uint32, free func()) {
	if len(s) == 0 {
		return nil, 0, func() {}
	}
	return CByteArray(s)
}

func CByteArray(s []byte) (c *byte, n uint32, free func()) {
	n = uint32(len(s))
	p := MemAlloc(uintptr(n))
	if p == nil {
		panic("failed to allocate unmanaged memory")
	}
	copy((*((*[0x7FFFFFFF]byte)(p)))[:n], s)
	return (*byte)(p), n, func() {
		MemFree(p)
	}
}

func CArrayReflect(dstPtrType reflect.Type, srcSlice reflect.Value, tr func(x interface{}) interface{}) (c unsafe.Pointer, n uint32, free func()) {
	if srcSlice.Kind() != reflect.Slice {
		panic("srcSlice must be a slice")
	}
	n = uint32(srcSlice.Len())
	if dstPtrType.Kind() != reflect.Ptr && dstPtrType.Kind() != reflect.Slice {
		panic("dstPtrType must be a type of pointer or slice")
	}
	sz := dstPtrType.Elem().Size() // sz = sizeof(*dst)
	c = MemAlloc(sz * uintptr(n))
	if tr == nil {
		if dstPtrType.Elem() != srcSlice.Type().Elem() {
			panic("dst element type must same as src element type when tr == nil")
		}
		for i := 0; i < srcSlice.Len(); i++ {
			p := unsafe.Pointer(uintptr(c) + sz*uintptr(i)) // p = c + sz*i
			pv := reflect.NewAt(dstPtrType.Elem(), p)       // pv := (*DstType)(p)
			pv.Elem().Set(srcSlice.Index(i))                // *pv = src[i]
		}
	} else {
		for i := 0; i < srcSlice.Len(); i++ {
			p := unsafe.Pointer(uintptr(c) + sz*uintptr(i))                   // p = c + sz*i
			pv := reflect.NewAt(dstPtrType.Elem(), p)                         // pv := (*DstType)(p)
			pv.Elem().Set(reflect.ValueOf(tr(srcSlice.Index(i).Interface()))) // *pv = tr(src[i])).(*DstType)
		}
	}
	free = func() {
		MemFree(c)
	}
	return
}

func CArray(dstPtrType, srcSlice interface{}, tr func(x interface{}) interface{}) (c unsafe.Pointer, n uint32, free func()) {
	if dstPtrType == nil {
		return CArrayReflect(reflect.TypeOf(srcSlice), reflect.ValueOf(srcSlice), tr)
	}
	return CArrayReflect(reflect.TypeOf(dstPtrType), reflect.ValueOf(srcSlice), tr)
}

func LoadInstanceProc(instance Instance, ppfn interface{}) error {
	if str, ok := ppfn.(interface{ String() string }); ok {
		name := str.String()
		v := reflect.ValueOf(ppfn)
		if v.Type().Kind() == reflect.Ptr && v.Type().Elem().Kind() == reflect.Uintptr {
			addr := GetInstanceProcAddr(instance, name)
			if addr == 0 {
				return fmt.Errorf("LoadInstanceProc() failure: %s", name)
			}
			v.Elem().SetUint(uint64(addr))
			return nil
		}
	}
	panic("ppfn must be pointed to a PfnXXXXXXX")
}

func MemCopy(dst, src unsafe.Pointer, size uint64) int {
	return copy(((*[0x7FFFFFFF]byte)(dst))[:size], ((*[0x7FFFFFFF]byte)(src))[:size])
}
