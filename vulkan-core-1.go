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
	"fmt"
	"math"
	"unsafe"
)

type Version uint32

const (
	API_VERSION_1_0         = 1 << 22
	API_VERSION_1_1         = 1<<22 | 1<<12
	API_VERSION_1_2         = 1<<22 | 2<<12
	HEADER_VERSION_COMPLETE = API_VERSION_1_2 | HEADER_VERSION
)

func MakeVersion(major, minor, patch uint32) Version {
	return Version(((major) << 22) | ((minor) << 12) | (patch))
}

func (ver Version) Major() uint32 { return ((uint32)(ver) >> 22) }
func (ver Version) Minor() uint32 { return (((uint32)(ver) >> 12) & 0x3ff) }
func (ver Version) Patch() uint32 { return ((uint32)(ver) & 0xfff) }

func (ver Version) String() string {
	return fmt.Sprintf("%d.%d.%d", ver.Major(), ver.Minor(), ver.Patch())
}

type (
	VkDeviceAddress       = uint64
	DispatchableHandle    = uintptr // pointer
	NonDispatchableHandle = uint64  // 64-bits, pointer or not
	Flags64               = uint64

	PfnVoidFunction = uintptr
)

const (
	NULL_HANDLE                   = 0
	REMAINING_MIP_LEVELS   uint32 = 0xFFFFFFFF         // ^0
	REMAINING_ARRAY_LAYERS uint32 = 0xFFFFFFFF         // ^0
	WHOLE_SIZE             uint64 = 0xFFFFFFFFFFFFFFFF // ^0
	ATTACHMENT_UNUSED      uint32 = 0xFFFFFFFF         // ^0
	QUEUE_FAMILY_IGNORED   uint32 = 0xFFFFFFFF         // ^0
	SUBPASS_EXTERNAL       uint32 = 0xFFFFFFFF         // ^0
	SHADER_UNUSED_NV       uint32 = 0xFFFFFFFF         // ^0
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

func CreateInstance(pCreateInfo *InstanceCreateInfo, pAllocator *AllocationCallbacks, pInstance *Instance) Result {
	fp := PfnCreateInstance(GetInstanceProcAddr(0, "vkCreateInstance"))
	if fp == 0 {
		return ERROR_UNKNOWN
	}
	return fp.Call(pCreateInfo, pAllocator, pInstance)
}

func EnumerateInstanceVersion(pApiVersion *Version) Result {
	fp := PfnEnumerateInstanceVersion(GetInstanceProcAddr(0, "vkEnumerateInstanceVersion"))
	if fp == 0 {
		return ERROR_UNKNOWN
	}
	return fp.Call((*uint32)(pApiVersion))
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
	*((*[8]byte)(unsafe.Pointer(v))) = *((*[8]byte)(unsafe.Pointer(&x)))
}
func (v *ClearValue) DepthStencil() ClearDepthStencilValue {
	return *((*ClearDepthStencilValue)(unsafe.Pointer(v)))
}

func (v *ClearValue) SetDepth(x float32) {
	*((*[4]byte)(unsafe.Pointer(v))) = *((*[4]byte)(unsafe.Pointer(&x)))
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

// ypedef struct VkAccelerationStructureInstanceKHR {
//     VkTransformMatrixKHR          transform;
//     uint32_t                      instanceCustomIndex:24;
//     uint32_t                      mask:8;
//     uint32_t                      instanceShaderBindingTableRecordOffset:24;
//     VkGeometryInstanceFlagsKHR    flags:8;
//     uint64_t                      accelerationStructureReference;
// } VkAccelerationStructureInstanceKHR;

type AccelerationStructureInstanceKHR struct {
	Transform                      TransformMatrixKHR
	bitfield0                      uint32 // instanceCustomIndex:24 mask:8
	bitfield1                      uint32 // instanceShaderBindingTableRecordOffset:24 flags:8
	AccelerationStructureReference uint64
}

func (a *AccelerationStructureInstanceKHR) SetInstanceCustomIndex(x uint32) {
	a.bitfield0 = (a.bitfield0 & 0xFF000000) | (x & 0x00FFFFFF)
}

func (a AccelerationStructureInstanceKHR) InstanceCustomIndex() uint32 {
	return a.bitfield0 & 0x00FFFFFF
}

func (a *AccelerationStructureInstanceKHR) SetMask(x uint8) {
	a.bitfield0 = (a.bitfield0 & 0x00FFFFFF) | (uint32(x) << 24)
}

func (a AccelerationStructureInstanceKHR) Mask() uint8 {
	return uint8((a.bitfield0 & 0xFF000000) >> 24)
}

func (a *AccelerationStructureInstanceKHR) SetInstanceShaderBindingTableRecordOffset(x uint32) {
	a.bitfield1 = (a.bitfield1 & 0xFF000000) | (x & 0x00FFFFFF)
}

func (a AccelerationStructureInstanceKHR) InstanceShaderBindingTableRecordOffset() uint32 {
	return a.bitfield1 & 0x00FFFFFF
}

func (a *AccelerationStructureInstanceKHR) SetFlags(x uint8) {
	a.bitfield1 = (a.bitfield1 & 0x00FFFFFF) | (uint32(x) << 24)
}

func (a AccelerationStructureInstanceKHR) Flags() uint8 {
	return uint8((a.bitfield1 & 0xFF000000) >> 24)
}

type DeviceOrHostAddressKHR uint64

type DeviceOrHostAddressConstKHR uint64

// typedef union VkAccelerationStructureGeometryDataKHR {
//     VkAccelerationStructureGeometryTrianglesDataKHR    triangles;
//     VkAccelerationStructureGeometryAabbsDataKHR        aabbs;
//     VkAccelerationStructureGeometryInstancesDataKHR    instances;
// } VkAccelerationStructureGeometryDataKHR;
/*
type AccelerationStructureGeometryTrianglesDataKHR struct {
	SType         StructureType   32
	PNext         unsafe.Pointer  64
	VertexFormat  Format          32
	VertexData    DeviceOrHostAddressConstKHR 64
	VertexStride  DeviceSize 64
	MaxVertex     uint32   32
	IndexType     IndexType 32
	IndexData     DeviceOrHostAddressConstKHR  64
	TransformData DeviceOrHostAddressConstKHR  64
}

// AccelerationStructureGeometryAabbsDataKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureGeometryAabbsDataKHR.html
type AccelerationStructureGeometryAabbsDataKHR struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Data   DeviceOrHostAddressConstKHR
	Stride DeviceSize
}

// AccelerationStructureGeometryInstancesDataKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html
type AccelerationStructureGeometryInstancesDataKHR struct {
	SType           StructureType
	PNext           unsafe.Pointer
	ArrayOfPointers Bool32
	Data            DeviceOrHostAddressConstKHR
}
*/
type AccelerationStructureGeometryDataKHR [448]byte // TODO: 验证其尺寸
