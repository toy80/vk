package vk

/*
 ** Copyright (c) 2015-2019 The Khronos Group Inc.
 **
 ** Licensed under the Apache License, Version 2.0 (the "License");
 ** you may not use this file except in compliance with the License.
 ** You may obtain a copy of the License at
 **
 **     http://www.apache.org/licenses/LICENSE-2.0
 **
 ** Unless required by applicable law or agreed to in writing, software
 ** distributed under the License is distributed on an "AS IS" BASIS,
 ** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 ** See the License for the specific language governing permissions and
 ** limitations under the License.
 */

import (
	"math"
	"unsafe"
)

const API_VERSION_1_0 uint32 = 1 << 22
const API_VERSION_1_1 uint32 = 1<<22 | 1<<12

func MakeVersion(major, minor, patch uint32) uint32 {
	return (((major) << 22) | ((minor) << 12) | (patch))
}
func VersionMajor(ver uint32) uint32 { return ((uint32)(ver) >> 22) }
func VersionMinor(ver uint32) uint32 { return (((uint32)(ver) >> 12) & 0x3ff) }
func VersionPatch(ver uint32) uint32 { return ((uint32)(ver) & 0xfff) }

type (
	DispatchableHandle    = uintptr // pointer
	NonDispatchableHandle = uint64  // 64-bits, pointer or not

	PfnVoidFunction = uintptr
)

const (
	REMAINING_MIP_LEVELS     uint32 = 0xFFFFFFFF         // ^0
	REMAINING_ARRAY_LAYERS   uint32 = 0xFFFFFFFF         // ^0
	WHOLE_SIZE               uint64 = 0xFFFFFFFFFFFFFFFF // ^0
	ATTACHMENT_UNUSED        uint32 = 0xFFFFFFFF         // ^0
	QUEUE_FAMILY_IGNORED     uint32 = 0xFFFFFFFF         // ^0
	QUEUE_FAMILY_EXTERNAL    uint32 = 0xFFFFFFFE         // (^0 - 1)
	QUEUE_FAMILY_FOREIGN_EXT uint32 = 0xFFFFFFFD         //(^0 - 2)
	SUBPASS_EXTERNAL         uint32 = 0xFFFFFFFF         // ^0
	SHADER_UNUSED_NV         uint32 = 0xFFFFFFFF         // ^0
)

type ErrorResult Result

func (r ErrorResult) Error() string {
	return "vulkan error: " + Result(r).String()
}

func (r Result) Err() error {
	if r != SUCCESS {
		return ErrorResult(r)
	}
	return nil
}

func AsResult(err error) Result {
	if err == nil {
		return SUCCESS
	}
	if r, ok := err.(ErrorResult); ok {
		return Result(r)
	}
	return RESULT_MAX_ENUM
}

// union VkClearColorValue {
// 	float       float32[4];
// 	int32_t     int32[4];
// 	uint32_t    uint32[4];
// };
// ClearColorValue -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkClearColorValue.html
type ClearColorValue [16]byte

func NewClearColorValue() *ClearColorValue {
	return (*ClearColorValue)(MemAlloc(unsafe.Sizeof(*(*ClearColorValue)(nil))))
}
func (p *ClearColorValue) Free() { MemFree(unsafe.Pointer(p)) }

func (v *ClearColorValue) SetFloat32(x [4]float32) {
	*v = *((*[16]byte)(unsafe.Pointer(&x)))
}
func (v *ClearColorValue) Float32Color() [4]float32 {
	return *((*[4]float32)(unsafe.Pointer(v)))
}
func (v *ClearColorValue) SetInt32(x [4]int32) {
	*v = *((*[16]byte)(unsafe.Pointer(&x)))
}
func (v *ClearColorValue) Int32Color() [4]int32 {
	return *((*[4]int32)(unsafe.Pointer(v)))
}
func (v *ClearColorValue) SetUint32(x [4]uint32) {
	*v = *((*[16]byte)(unsafe.Pointer(&x)))
}
func (v *ClearColorValue) Uint32Color() [4]uint32 {
	return *((*[4]uint32)(unsafe.Pointer(v)))
}

// union VkClearValue {
// 	VkClearColorValue           color;
// 	VkClearDepthStencilValue    depthStencil;
// };
// ClearValue -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkClearValue.html
type ClearValue [16]byte

func NewClearValue() *ClearValue {
	return (*ClearValue)(MemAlloc(unsafe.Sizeof(*(*ClearValue)(nil))))
}
func (p *ClearValue) Free() { MemFree(unsafe.Pointer(p)) }

func (v *ClearValue) SetClearColor(x ClearColorValue) {
	*v = *((*[16]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) ClearColor() ClearColorValue {
	return *((*ClearColorValue)(unsafe.Pointer(v)))
}

func (v *ClearValue) SetFloat32Color(x [4]float32) {
	*v = *((*[16]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) Float32Color() [4]float32 {
	return *((*[4]float32)(unsafe.Pointer(v)))
}
func (v *ClearValue) SetInt32Color(x [4]int32) {
	*v = *((*[16]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) Int32Color() [4]int32 {
	return *((*[4]int32)(unsafe.Pointer(v)))
}
func (v *ClearValue) SetUint32Color(x [4]uint32) {
	*v = *((*[16]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) Uint32Color() [4]uint32 {
	return *((*[4]uint32)(unsafe.Pointer(v)))
}

func (v *ClearValue) SetDepthStencil(x ClearDepthStencilValue) {
	*((*[8]byte)(unsafe.Pointer(&v))) = *((*[8]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) DepthStencil() ClearDepthStencilValue {
	return *((*ClearDepthStencilValue)(unsafe.Pointer(v)))
}

func (v *ClearValue) SetDepth(x float32) {
	*((*[4]byte)(unsafe.Pointer(&v))) = *((*[4]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) Depth() float32 {
	return *((*float32)(unsafe.Pointer(v)))
}
func (v *ClearValue) SetStencil(x int32) {
	// unsafe.Offsetof(ClearDepthStencilValue.Stencil) = 4
	*((*[4]byte)(unsafe.Pointer(&v[4]))) = *((*[4]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) Stencil() int32 {
	// unsafe.Offsetof(ClearDepthStencilValue.Stencil) = 4
	return *((*int32)(unsafe.Pointer(&v[4])))
}

// union VkPerformanceValueDataINTEL {
// 	uint32_t       value32;
// 	uint64_t       value64;
// 	float          valueFloat;
// 	VkBool32       valueBool;
// 	const char*    valueString;
// };
type PerformanceValueDataINTEL uint64

func (v *PerformanceValueDataINTEL) SetValue32(x uint32) {
	*v = PerformanceValueDataINTEL(uint64(x))
}
func (v *PerformanceValueDataINTEL) SetValue64(x uint64) {
	*v = PerformanceValueDataINTEL(x)
}
func (v *PerformanceValueDataINTEL) SetValueFloat(x float32) {
	*v = PerformanceValueDataINTEL(uint64(math.Float32bits(x)))
}
func (v *PerformanceValueDataINTEL) SetValueBool(x Bool32) {
	*v = PerformanceValueDataINTEL(uint64(x))
}
func (v *PerformanceValueDataINTEL) SetValueCString(x *int8) {
	*v = PerformanceValueDataINTEL(uint64(uintptr(unsafe.Pointer(x))))
}

// union VkPipelineExecutableStatisticValueKHR {
// 	VkBool32    b32;
// 	int64_t     i64;
// 	uint64_t    u64;
// 	double      f64;
// } VkPipelineExecutableStatisticValueKHR;

type PipelineExecutableStatisticValueKHR uint64

// TODO:
