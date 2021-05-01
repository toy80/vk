package vk

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"unsafe"
)

type sliceHeader = reflect.SliceHeader

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

	//a := (*((*[0x7FFFFFFF]byte)(p)))[:n]
	var a []byte
	h := (*sliceHeader)(unsafe.Pointer(&a))
	h.Data, h.Len, h.Cap = uintptr(p), int(n), int(n)

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
		// a := (*((*[0x7FFFFFFF]*int8)(p)))[:n]
		var a []*int8
		h := (*sliceHeader)(unsafe.Pointer(&a))
		h.Data, h.Len, h.Cap = uintptr(p), int(n), int(n)

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
	return ptrUint8ToString((*uint8)(unsafe.Pointer(p)))
}

func ptrUint8ToString(p *uint8) string {
	if p == nil || *p == 0 {
		return ""
	}
	var n uintptr
	for *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + n)) != 0 {
		n++
	}
	var s []byte
	h := (*sliceHeader)(unsafe.Pointer(&s))
	h.Data, h.Len, h.Cap = uintptr(unsafe.Pointer(p)), int(n), int(n)
	return string(s)
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
	// a := (*((*[0x7FFFFFFF]*int8)(unsafe.Pointer(pp))))[:n]
	var a []*int8
	h := (*sliceHeader)(unsafe.Pointer(&a))
	h.Data, h.Len, h.Cap = uintptr(unsafe.Pointer(pp)), int(n), int(n)

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

	// copy((*((*[0x7FFFFFFF]uint32)(p)))[:n], s)
	var ds []uint32
	hd := (*sliceHeader)(unsafe.Pointer(&ds))
	hd.Data, hd.Len, hd.Cap = uintptr(p), int(n), int(n)
	if len(ds) != int(n) {
		panic("internal error")
	}
	copy(ds, s)

	return (*uint32)(p), n, func() {
		MemFree(p)
	}
}

func CUint16ArrayOrNil(s []uint16) (c *uint16, n uint32, free func()) {
	if len(s) == 0 {
		return nil, 0, func() {}
	}
	return CUint16Array(s)
}

func CUint16Array(s []uint16) (c *uint16, n uint32, free func()) {
	n = uint32(len(s))
	p := MemAlloc(uintptr(n) * 2)
	if p == nil {
		panic("failed to allocate unmanaged memory")
	}
	// copy((*((*[0x7FFFFFFF]uint16)(p)))[:n], s)
	var ds []uint16
	hd := (*sliceHeader)(unsafe.Pointer(&ds))
	hd.Data, hd.Len, hd.Cap = uintptr(p), int(n), int(n)
	if len(ds) != int(n) {
		panic("internal error")
	}
	copy(ds, s)
	return (*uint16)(p), n, func() {
		MemFree(p)
	}
}

func CFloat32ArrayOrNil(s []float32) (c *float32, n uint32, free func()) {
	if len(s) == 0 {
		return nil, 0, func() {}
	}
	return CFloat32Array(s)
}

func CFloat32Array(s []float32) (c *float32, n uint32, free func()) {
	n = uint32(len(s))
	p := MemAlloc(uintptr(n) * 4)
	if p == nil {
		panic("failed to allocate unmanaged memory")
	}
	// copy((*((*[0x7FFFFFFF]float32)(p)))[:n], s)
	var ds []float32
	hd := (*sliceHeader)(unsafe.Pointer(&ds))
	hd.Data, hd.Len, hd.Cap = uintptr(p), int(n), int(n)
	if len(ds) != int(n) {
		panic("internal error")
	}
	copy(ds, s)
	return (*float32)(p), n, func() {
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
	// copy((*((*[0x7FFFFFFF]byte)(p)))[:n], s)
	var ds []byte
	hd := (*sliceHeader)(unsafe.Pointer(&ds))
	hd.Data, hd.Len, hd.Cap = uintptr(p), int(n), int(n)
	copy(ds, s)
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

func MemCopy(dst unsafe.Pointer, dstOffset uint64, src unsafe.Pointer, srcOffset uint64, size uint64) int {
	// *[0x7FFFFFFF]byte 这种模式通不过checkptr:
	// var a [10]byte
	// ((*[0x7FFFFFFF]byte)(unsafe.Pointer(&a)))[1] = 1
	//
	// 我们换一种(不安全的)方式构造slice:

	var ds []byte
	hd := (*sliceHeader)(unsafe.Pointer(&ds))
	hd.Data, hd.Len, hd.Cap = uintptr(dst), int(dstOffset+size), int(dstOffset+size)

	var ss []byte
	hs := (*sliceHeader)(unsafe.Pointer(&ss))
	hs.Data, hs.Len, hs.Cap = uintptr(src), int(srcOffset+size), int(srcOffset+size)
	return copy(ds[dstOffset:], ss[srcOffset:])
}

func MemZero(dst unsafe.Pointer, size uint64) {
	var ds []byte
	hd := (*sliceHeader)(unsafe.Pointer(&ds))
	hd.Data, hd.Len, hd.Cap = uintptr(dst), int(size), int(size)
	for i := range ds {
		ds[i] = 0
	}
}

var ShaderFileTypes = map[string]ShaderStageFlags{
	`.vert`:  SHADER_STAGE_VERTEX_BIT,
	`.tesc`:  SHADER_STAGE_TESSELLATION_CONTROL_BIT,
	`.tese`:  SHADER_STAGE_TESSELLATION_EVALUATION_BIT,
	`.geom`:  SHADER_STAGE_GEOMETRY_BIT,
	`.frag`:  SHADER_STAGE_FRAGMENT_BIT,
	`.comp`:  SHADER_STAGE_COMPUTE_BIT,
	`.mesh`:  SHADER_STAGE_MESH_BIT_NV,
	`.task`:  SHADER_STAGE_TASK_BIT_NV,
	`.rgen`:  SHADER_STAGE_RAYGEN_BIT_NV,
	`.rint`:  SHADER_STAGE_INTERSECTION_BIT_NV,
	`.rahit`: SHADER_STAGE_ANY_HIT_BIT_NV,
	`.rchit`: SHADER_STAGE_CLOSEST_HIT_BIT_NV,
	`.rmiss`: SHADER_STAGE_MISS_BIT_NV,
	`.rcall`: SHADER_STAGE_CALLABLE_BIT_NV,
	`.glsl`:  SHADER_STAGE_ALL,
	`.hlsl`:  SHADER_STAGE_ALL,
}

func ShaderStageByFileName(name string) (stage ShaderStageFlags, ok bool) {
	ext := filepath.Ext(name)
	stage, ok = ShaderFileTypes[ext]
	if !ok {
		return SHADER_STAGE_ALL, false
	}
	if stage == SHADER_STAGE_ALL {
		// foo.vert.glsl  bar.frag.hlsl
		return ShaderStageByFileName(strings.TrimSuffix(name, ext))
	}
	return
}

func FindInNextChain(p unsafe.Pointer, sType StructureType) unsafe.Pointer {
	for p1 := (*BaseOutStructure)(p); p1 != nil; p1 = p1.PNext {
		if p1.SType == sType {
			return unsafe.Pointer(p1)
		}
	}
	return nil
}
