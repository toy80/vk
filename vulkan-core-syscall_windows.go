// +build !forcecgo

package vk

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

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

/*
 ** This file is generated from the Vulkan headers.
 */

type (
	HANDLE    = uintptr
	HINSTANCE = uintptr
	HWND      = uintptr
	HMONITOR  = uintptr
	DWORD     = uint32

	LPCWSTR *uint16

	SECURITY_ATTRIBUTES struct {
		Length               DWORD
		LpSecurityDescriptor unsafe.Pointer
		InheritHandle        uint32
	}
)

var (
	kernel32dll    = syscall.NewLazyDLL("kernel32.dll")
	procLocalAlloc = kernel32dll.NewProc("LocalAlloc") // HLOCAL LocalAlloc(UINT uFlags, SIZE_T uBytes);
	procLocalFree  = kernel32dll.NewProc("LocalFree")  // HLOCAL LocalFree(HLOCAL hMem);

	vkdll                   = syscall.NewLazyDLL("vulkan-1.dll")
	procCreateInstance      = vkdll.NewProc("vkCreateInstance")
	procGetInstanceProcAddr = vkdll.NewProc("vkGetInstanceProcAddr")
)

// MemAlloc allocate zeroed C memory block
func MemAlloc(sz uintptr) (p unsafe.Pointer) {
	// Address of a block of C memory is of course "unsafe pointer"
	if sz == 0 {
		sz = 1 // MemAlloc(0) should return a non nil pointer
	}
	*(*uintptr)(unsafe.Pointer(&p)), _, _ = procLocalAlloc.Call(0x0040, sz) // 0x0040 = LMEM_FIXED | LMEM_ZEROINIT.
	dbgMemAlloc(uintptr(p))
	return
}

// MemFree release C memory block that allocated with MemAlloc()
func MemFree(p unsafe.Pointer) {
	dbgMemFree(uintptr(p))
	_, _, _ = procLocalFree.Call(uintptr(p))
}

func CreateInstance(pCreateInfo *InstanceCreateInfo, pAllocator *AllocationCallbacks, pInstance *Instance) Result {
	if err := vkdll.Load(); err != nil {
		fmt.Println(err)
		return ERROR_INITIALIZATION_FAILED
	}
	ret, _, _ := procCreateInstance.Call(
		uintptr(unsafe.Pointer(pCreateInfo)),
		uintptr(unsafe.Pointer(pAllocator)),
		uintptr(unsafe.Pointer(pInstance)))
	return Result(ret)
}

func GetInstanceProcAddr(instance Instance, name string) PfnVoidFunction {
	c := []byte(name)
	c = append(c, 0)
	ret, _, _ := procGetInstanceProcAddr.Call(uintptr(instance), uintptr(unsafe.Pointer(&c[0])))
	return ret
}

func call(addr uintptr, a ...uintptr) (r1, r2 uintptr, lastErr error) {
	switch len(a) {
	case 0:
		return syscall.Syscall(addr, uintptr(len(a)), 0, 0, 0)
	case 1:
		return syscall.Syscall(addr, uintptr(len(a)), a[0], 0, 0)
	case 2:
		return syscall.Syscall(addr, uintptr(len(a)), a[0], a[1], 0)
	case 3:
		return syscall.Syscall(addr, uintptr(len(a)), a[0], a[1], a[2])
	case 4:
		return syscall.Syscall6(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], 0, 0)
	case 5:
		return syscall.Syscall6(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], 0)
	case 6:
		return syscall.Syscall6(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5])
	case 7:
		return syscall.Syscall9(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0)
	case 8:
		return syscall.Syscall9(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0)
	case 9:
		return syscall.Syscall9(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
	case 10:
		return syscall.Syscall12(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0)
	case 11:
		return syscall.Syscall12(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0)
	case 12:
		return syscall.Syscall12(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
	case 13:
		return syscall.Syscall15(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], 0, 0)
	case 14:
		return syscall.Syscall15(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], 0)
	case 15:
		return syscall.Syscall15(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14])
	case 16:
		return syscall.Syscall18(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], 0, 0)
	case 17:
		return syscall.Syscall18(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], a[16], 0)
	case 18:
		return syscall.Syscall18(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], a[16], a[17])
	default:
		panic("Syscall with too many arguments " + fmt.Sprint(len(a)) + ".")
	}
}

const VERSION_1_0 = 1
const HEADER_VERSION = 122
const NULL_HANDLE = 0

type Bool32 = uint32
type DeviceSize = uint64
type SampleMask = uint32

// Instance -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkInstance.html
type Instance DispatchableHandle

// PhysicalDevice -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDevice.html
type PhysicalDevice DispatchableHandle

// Device -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDevice.html
type Device DispatchableHandle

// Queue -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueue.html
type Queue DispatchableHandle

// Semaphore -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSemaphore.html
type Semaphore NonDispatchableHandle

// CommandBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBuffer.html
type CommandBuffer DispatchableHandle

// Fence -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFence.html
type Fence NonDispatchableHandle

// DeviceMemory -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceMemory.html
type DeviceMemory NonDispatchableHandle

// Buffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBuffer.html
type Buffer NonDispatchableHandle

// Image -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImage.html
type Image NonDispatchableHandle

// Event -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkEvent.html
type Event NonDispatchableHandle

// QueryPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryPool.html
type QueryPool NonDispatchableHandle

// BufferView -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferView.html
type BufferView NonDispatchableHandle

// ImageView -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageView.html
type ImageView NonDispatchableHandle

// ShaderModule -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderModule.html
type ShaderModule NonDispatchableHandle

// PipelineCache -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCache.html
type PipelineCache NonDispatchableHandle

// PipelineLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineLayout.html
type PipelineLayout NonDispatchableHandle

// RenderPass -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPass.html
type RenderPass NonDispatchableHandle

// Pipeline -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipeline.html
type Pipeline NonDispatchableHandle

// DescriptorSetLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetLayout.html
type DescriptorSetLayout NonDispatchableHandle

// Sampler -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSampler.html
type Sampler NonDispatchableHandle

// DescriptorPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorPool.html
type DescriptorPool NonDispatchableHandle

// DescriptorSet -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSet.html
type DescriptorSet NonDispatchableHandle

// Framebuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFramebuffer.html
type Framebuffer NonDispatchableHandle

// CommandPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandPool.html
type CommandPool NonDispatchableHandle

const LOD_CLAMP_NONE = 1000.0
const TRUE = 1
const FALSE = 0
const MAX_PHYSICAL_DEVICE_NAME_SIZE = 256
const UUID_SIZE = 16
const MAX_MEMORY_TYPES = 32
const MAX_MEMORY_HEAPS = 16
const MAX_EXTENSION_NAME_SIZE = 256
const MAX_DESCRIPTION_SIZE = 256

// PipelineCacheHeaderVersion -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCacheHeaderVersion.html
type PipelineCacheHeaderVersion int32

const (
	PIPELINE_CACHE_HEADER_VERSION_ONE         PipelineCacheHeaderVersion = 1
	PIPELINE_CACHE_HEADER_VERSION_BEGIN_RANGE PipelineCacheHeaderVersion = PIPELINE_CACHE_HEADER_VERSION_ONE
	PIPELINE_CACHE_HEADER_VERSION_END_RANGE   PipelineCacheHeaderVersion = PIPELINE_CACHE_HEADER_VERSION_ONE
	PIPELINE_CACHE_HEADER_VERSION_RANGE_SIZE  PipelineCacheHeaderVersion = (PIPELINE_CACHE_HEADER_VERSION_ONE - PIPELINE_CACHE_HEADER_VERSION_ONE + 1)
	PIPELINE_CACHE_HEADER_VERSION_MAX_ENUM    PipelineCacheHeaderVersion = 0x7FFFFFFF
)

func (x PipelineCacheHeaderVersion) String() string {
	switch x {
	case PIPELINE_CACHE_HEADER_VERSION_ONE:
		return "PIPELINE_CACHE_HEADER_VERSION_ONE"
	case PIPELINE_CACHE_HEADER_VERSION_MAX_ENUM:
		return "PIPELINE_CACHE_HEADER_VERSION_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// Result -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkResult.html
type Result int32

const (
	SUCCESS                                            Result = 0
	NOT_READY                                          Result = 1
	TIMEOUT                                            Result = 2
	EVENT_SET                                          Result = 3
	EVENT_RESET                                        Result = 4
	INCOMPLETE                                         Result = 5
	ERROR_OUT_OF_HOST_MEMORY                           Result = -1
	ERROR_OUT_OF_DEVICE_MEMORY                         Result = -2
	ERROR_INITIALIZATION_FAILED                        Result = -3
	ERROR_DEVICE_LOST                                  Result = -4
	ERROR_MEMORY_MAP_FAILED                            Result = -5
	ERROR_LAYER_NOT_PRESENT                            Result = -6
	ERROR_EXTENSION_NOT_PRESENT                        Result = -7
	ERROR_FEATURE_NOT_PRESENT                          Result = -8
	ERROR_INCOMPATIBLE_DRIVER                          Result = -9
	ERROR_TOO_MANY_OBJECTS                             Result = -10
	ERROR_FORMAT_NOT_SUPPORTED                         Result = -11
	ERROR_FRAGMENTED_POOL                              Result = -12
	ERROR_OUT_OF_POOL_MEMORY                           Result = -1000069000
	ERROR_INVALID_EXTERNAL_HANDLE                      Result = -1000072003
	ERROR_SURFACE_LOST_KHR                             Result = -1000000000
	ERROR_NATIVE_WINDOW_IN_USE_KHR                     Result = -1000000001
	SUBOPTIMAL_KHR                                     Result = 1000001003
	ERROR_OUT_OF_DATE_KHR                              Result = -1000001004
	ERROR_INCOMPATIBLE_DISPLAY_KHR                     Result = -1000003001
	ERROR_VALIDATION_FAILED_EXT                        Result = -1000011001
	ERROR_INVALID_SHADER_NV                            Result = -1000012000
	ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT Result = -1000158000
	ERROR_FRAGMENTATION_EXT                            Result = -1000161000
	ERROR_NOT_PERMITTED_EXT                            Result = -1000174001
	ERROR_INVALID_DEVICE_ADDRESS_EXT                   Result = -1000244000
	ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT          Result = -1000255000
	ERROR_OUT_OF_POOL_MEMORY_KHR                       Result = ERROR_OUT_OF_POOL_MEMORY
	ERROR_INVALID_EXTERNAL_HANDLE_KHR                  Result = ERROR_INVALID_EXTERNAL_HANDLE
	RESULT_BEGIN_RANGE                                 Result = ERROR_FRAGMENTED_POOL
	RESULT_END_RANGE                                   Result = INCOMPLETE
	RESULT_RANGE_SIZE                                  Result = (INCOMPLETE - ERROR_FRAGMENTED_POOL + 1)
	RESULT_MAX_ENUM                                    Result = 0x7FFFFFFF
)

func (x Result) String() string {
	switch x {
	case SUCCESS:
		return "SUCCESS"
	case NOT_READY:
		return "NOT_READY"
	case TIMEOUT:
		return "TIMEOUT"
	case EVENT_SET:
		return "EVENT_SET"
	case EVENT_RESET:
		return "EVENT_RESET"
	case INCOMPLETE:
		return "INCOMPLETE"
	case ERROR_OUT_OF_HOST_MEMORY:
		return "ERROR_OUT_OF_HOST_MEMORY"
	case ERROR_OUT_OF_DEVICE_MEMORY:
		return "ERROR_OUT_OF_DEVICE_MEMORY"
	case ERROR_INITIALIZATION_FAILED:
		return "ERROR_INITIALIZATION_FAILED"
	case ERROR_DEVICE_LOST:
		return "ERROR_DEVICE_LOST"
	case ERROR_MEMORY_MAP_FAILED:
		return "ERROR_MEMORY_MAP_FAILED"
	case ERROR_LAYER_NOT_PRESENT:
		return "ERROR_LAYER_NOT_PRESENT"
	case ERROR_EXTENSION_NOT_PRESENT:
		return "ERROR_EXTENSION_NOT_PRESENT"
	case ERROR_FEATURE_NOT_PRESENT:
		return "ERROR_FEATURE_NOT_PRESENT"
	case ERROR_INCOMPATIBLE_DRIVER:
		return "ERROR_INCOMPATIBLE_DRIVER"
	case ERROR_TOO_MANY_OBJECTS:
		return "ERROR_TOO_MANY_OBJECTS"
	case ERROR_FORMAT_NOT_SUPPORTED:
		return "ERROR_FORMAT_NOT_SUPPORTED"
	case ERROR_FRAGMENTED_POOL:
		return "ERROR_FRAGMENTED_POOL"
	case ERROR_OUT_OF_POOL_MEMORY:
		return "ERROR_OUT_OF_POOL_MEMORY"
	case ERROR_INVALID_EXTERNAL_HANDLE:
		return "ERROR_INVALID_EXTERNAL_HANDLE"
	case ERROR_SURFACE_LOST_KHR:
		return "ERROR_SURFACE_LOST_KHR"
	case ERROR_NATIVE_WINDOW_IN_USE_KHR:
		return "ERROR_NATIVE_WINDOW_IN_USE_KHR"
	case SUBOPTIMAL_KHR:
		return "SUBOPTIMAL_KHR"
	case ERROR_OUT_OF_DATE_KHR:
		return "ERROR_OUT_OF_DATE_KHR"
	case ERROR_INCOMPATIBLE_DISPLAY_KHR:
		return "ERROR_INCOMPATIBLE_DISPLAY_KHR"
	case ERROR_VALIDATION_FAILED_EXT:
		return "ERROR_VALIDATION_FAILED_EXT"
	case ERROR_INVALID_SHADER_NV:
		return "ERROR_INVALID_SHADER_NV"
	case ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT:
		return "ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT"
	case ERROR_FRAGMENTATION_EXT:
		return "ERROR_FRAGMENTATION_EXT"
	case ERROR_NOT_PERMITTED_EXT:
		return "ERROR_NOT_PERMITTED_EXT"
	case ERROR_INVALID_DEVICE_ADDRESS_EXT:
		return "ERROR_INVALID_DEVICE_ADDRESS_EXT"
	case ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT:
		return "ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT"
	case RESULT_MAX_ENUM:
		return "RESULT_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// StructureType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkStructureType.html
type StructureType int32

const (
	STRUCTURE_TYPE_APPLICATION_INFO                                                StructureType = 0
	STRUCTURE_TYPE_INSTANCE_CREATE_INFO                                            StructureType = 1
	STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO                                        StructureType = 2
	STRUCTURE_TYPE_DEVICE_CREATE_INFO                                              StructureType = 3
	STRUCTURE_TYPE_SUBMIT_INFO                                                     StructureType = 4
	STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO                                            StructureType = 5
	STRUCTURE_TYPE_MAPPED_MEMORY_RANGE                                             StructureType = 6
	STRUCTURE_TYPE_BIND_SPARSE_INFO                                                StructureType = 7
	STRUCTURE_TYPE_FENCE_CREATE_INFO                                               StructureType = 8
	STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO                                           StructureType = 9
	STRUCTURE_TYPE_EVENT_CREATE_INFO                                               StructureType = 10
	STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO                                          StructureType = 11
	STRUCTURE_TYPE_BUFFER_CREATE_INFO                                              StructureType = 12
	STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO                                         StructureType = 13
	STRUCTURE_TYPE_IMAGE_CREATE_INFO                                               StructureType = 14
	STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO                                          StructureType = 15
	STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO                                       StructureType = 16
	STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO                                      StructureType = 17
	STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO                               StructureType = 18
	STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO                         StructureType = 19
	STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO                       StructureType = 20
	STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO                         StructureType = 21
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO                             StructureType = 22
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO                        StructureType = 23
	STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO                          StructureType = 24
	STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO                        StructureType = 25
	STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO                          StructureType = 26
	STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO                              StructureType = 27
	STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO                                   StructureType = 28
	STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO                                    StructureType = 29
	STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO                                     StructureType = 30
	STRUCTURE_TYPE_SAMPLER_CREATE_INFO                                             StructureType = 31
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO                               StructureType = 32
	STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO                                     StructureType = 33
	STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO                                    StructureType = 34
	STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET                                            StructureType = 35
	STRUCTURE_TYPE_COPY_DESCRIPTOR_SET                                             StructureType = 36
	STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO                                         StructureType = 37
	STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO                                         StructureType = 38
	STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO                                        StructureType = 39
	STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO                                    StructureType = 40
	STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO                                 StructureType = 41
	STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO                                       StructureType = 42
	STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO                                          StructureType = 43
	STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER                                           StructureType = 44
	STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER                                            StructureType = 45
	STRUCTURE_TYPE_MEMORY_BARRIER                                                  StructureType = 46
	STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO                                     StructureType = 47
	STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO                                       StructureType = 48
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES                             StructureType = 1000094000
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO                                         StructureType = 1000157000
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO                                          StructureType = 1000157001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES                          StructureType = 1000083000
	STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS                                   StructureType = 1000127000
	STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO                                  StructureType = 1000127001
	STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO                                      StructureType = 1000060000
	STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO                             StructureType = 1000060003
	STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO                          StructureType = 1000060004
	STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO                                        StructureType = 1000060005
	STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO                                   StructureType = 1000060006
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO                            StructureType = 1000060013
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO                             StructureType = 1000060014
	STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES                                StructureType = 1000070000
	STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO                                 StructureType = 1000070001
	STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2                               StructureType = 1000146000
	STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2                                StructureType = 1000146001
	STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2                         StructureType = 1000146002
	STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2                                           StructureType = 1000146003
	STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2                              StructureType = 1000146004
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2                                      StructureType = 1000059000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2                                    StructureType = 1000059001
	STRUCTURE_TYPE_FORMAT_PROPERTIES_2                                             StructureType = 1000059002
	STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2                                       StructureType = 1000059003
	STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2                             StructureType = 1000059004
	STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2                                       StructureType = 1000059005
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2                             StructureType = 1000059006
	STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2                                StructureType = 1000059007
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2                      StructureType = 1000059008
	STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES                       StructureType = 1000117000
	STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO                 StructureType = 1000117001
	STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO                                    StructureType = 1000117002
	STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO           StructureType = 1000117003
	STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO                               StructureType = 1000053000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES                              StructureType = 1000053001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES                            StructureType = 1000053002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES                      StructureType = 1000120000
	STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO                                           StructureType = 1000145000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES                       StructureType = 1000145001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES                     StructureType = 1000145002
	STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2                                             StructureType = 1000145003
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO                            StructureType = 1000156000
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO                                   StructureType = 1000156001
	STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO                                    StructureType = 1000156002
	STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO                            StructureType = 1000156003
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES               StructureType = 1000156004
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES                StructureType = 1000156005
	STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO                          StructureType = 1000085000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO                      StructureType = 1000071000
	STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES                                StructureType = 1000071001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO                            StructureType = 1000071002
	STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES                                      StructureType = 1000071003
	STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES                                   StructureType = 1000071004
	STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO                              StructureType = 1000072000
	STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO                               StructureType = 1000072001
	STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO                                     StructureType = 1000072002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO                             StructureType = 1000112000
	STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES                                       StructureType = 1000112001
	STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO                                        StructureType = 1000113000
	STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO                                    StructureType = 1000077000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO                         StructureType = 1000076000
	STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES                                   StructureType = 1000076001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES                        StructureType = 1000168000
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT                                   StructureType = 1000168001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES                 StructureType = 1000063000
	STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR                                       StructureType = 1000001000
	STRUCTURE_TYPE_PRESENT_INFO_KHR                                                StructureType = 1000001001
	STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR                           StructureType = 1000060007
	STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR                                 StructureType = 1000060008
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR                            StructureType = 1000060009
	STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR                                     StructureType = 1000060010
	STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR                                   StructureType = 1000060011
	STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR                          StructureType = 1000060012
	STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR                                    StructureType = 1000002000
	STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR                                 StructureType = 1000002001
	STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR                                        StructureType = 1000003000
	STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR                                    StructureType = 1000004000
	STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR                                     StructureType = 1000005000
	STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR                                 StructureType = 1000006000
	STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR                                 StructureType = 1000008000
	STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR                                   StructureType = 1000009000
	STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT                           StructureType = 1000011000
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD            StructureType = 1000018000
	STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT                               StructureType = 1000022000
	STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT                                StructureType = 1000022001
	STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT                                    StructureType = 1000022002
	STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV                       StructureType = 1000026000
	STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV                      StructureType = 1000026001
	STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV                    StructureType = 1000026002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_FEATURES_EXT                 StructureType = 1000028000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_PROPERTIES_EXT               StructureType = 1000028001
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT             StructureType = 1000028002
	STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX                                      StructureType = 1000030000
	STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD                        StructureType = 1000041000
	STRUCTURE_TYPE_STREAM_DESCRIPTOR_SURFACE_CREATE_INFO_GGP                       StructureType = 1000049000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV                StructureType = 1000050000
	STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV                            StructureType = 1000056000
	STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV                                  StructureType = 1000056001
	STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV                              StructureType = 1000057000
	STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV                              StructureType = 1000057001
	STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV                       StructureType = 1000058000
	STRUCTURE_TYPE_VALIDATION_FLAGS_EXT                                            StructureType = 1000061000
	STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN                                       StructureType = 1000062000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES_EXT       StructureType = 1000066000
	STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT                                 StructureType = 1000067000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT                        StructureType = 1000067001
	STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR                             StructureType = 1000073000
	STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR                             StructureType = 1000073001
	STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR                              StructureType = 1000073002
	STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR                                StructureType = 1000073003
	STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR                                       StructureType = 1000074000
	STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR                                        StructureType = 1000074001
	STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR                                          StructureType = 1000074002
	STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR                      StructureType = 1000075000
	STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                          StructureType = 1000078000
	STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR                          StructureType = 1000078001
	STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR                                     StructureType = 1000078002
	STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR                             StructureType = 1000078003
	STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR                                    StructureType = 1000079000
	STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR                                       StructureType = 1000079001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR                  StructureType = 1000080000
	STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT       StructureType = 1000081000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT              StructureType = 1000081001
	STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT                            StructureType = 1000081002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR                StructureType = 1000082000
	STRUCTURE_TYPE_PRESENT_REGIONS_KHR                                             StructureType = 1000084000
	STRUCTURE_TYPE_OBJECT_TABLE_CREATE_INFO_NVX                                    StructureType = 1000086000
	STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NVX                        StructureType = 1000086001
	STRUCTURE_TYPE_CMD_PROCESS_COMMANDS_INFO_NVX                                   StructureType = 1000086002
	STRUCTURE_TYPE_CMD_RESERVE_SPACE_FOR_COMMANDS_INFO_NVX                         StructureType = 1000086003
	STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_LIMITS_NVX                            StructureType = 1000086004
	STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_FEATURES_NVX                          StructureType = 1000086005
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV                StructureType = 1000087000
	STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT                                      StructureType = 1000090000
	STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT                                          StructureType = 1000091000
	STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT                                           StructureType = 1000091001
	STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT                                          StructureType = 1000091002
	STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT                               StructureType = 1000091003
	STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE                                       StructureType = 1000092000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX    StructureType = 1000097000
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV                  StructureType = 1000098000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT                StructureType = 1000099000
	STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT                StructureType = 1000099001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT       StructureType = 1000101000
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT       StructureType = 1000101001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_ENABLE_FEATURES_EXT                  StructureType = 1000102000
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT         StructureType = 1000102001
	STRUCTURE_TYPE_HDR_METADATA_EXT                                                StructureType = 1000105000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR              StructureType = 1000108000
	STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR                         StructureType = 1000108001
	STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR                           StructureType = 1000108002
	STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR                           StructureType = 1000108003
	STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR                                    StructureType = 1000109000
	STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR                                      StructureType = 1000109001
	STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR                                       StructureType = 1000109002
	STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR                                        StructureType = 1000109003
	STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR                                   StructureType = 1000109004
	STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR                                          StructureType = 1000109005
	STRUCTURE_TYPE_SUBPASS_END_INFO_KHR                                            StructureType = 1000109006
	STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR                         StructureType = 1000111000
	STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR                              StructureType = 1000114000
	STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR                              StructureType = 1000114001
	STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR                                 StructureType = 1000114002
	STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR                                        StructureType = 1000115000
	STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR                                           StructureType = 1000115001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR                              StructureType = 1000119000
	STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR                                      StructureType = 1000119001
	STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR                                            StructureType = 1000119002
	STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR                                        StructureType = 1000121000
	STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR                                  StructureType = 1000121001
	STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR                                   StructureType = 1000121002
	STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR                                        StructureType = 1000121003
	STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR                                StructureType = 1000121004
	STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK                                     StructureType = 1000122000
	STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK                                   StructureType = 1000123000
	STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT                                StructureType = 1000128000
	STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT                                 StructureType = 1000128001
	STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT                                           StructureType = 1000128002
	STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT                         StructureType = 1000128003
	STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT                           StructureType = 1000128004
	STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID                           StructureType = 1000129000
	STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID                      StructureType = 1000129001
	STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID               StructureType = 1000129002
	STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID                     StructureType = 1000129003
	STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID                 StructureType = 1000129004
	STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID                                         StructureType = 1000129005
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT            StructureType = 1000130000
	STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT                          StructureType = 1000130001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT               StructureType = 1000138000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT             StructureType = 1000138001
	STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT                   StructureType = 1000138002
	STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT            StructureType = 1000138003
	STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT                                       StructureType = 1000143000
	STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT                     StructureType = 1000143001
	STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT                 StructureType = 1000143002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT                 StructureType = 1000143003
	STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT                                      StructureType = 1000143004
	STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR                               StructureType = 1000147000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT           StructureType = 1000148000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT         StructureType = 1000148001
	STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT             StructureType = 1000148002
	STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV                 StructureType = 1000149000
	STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV               StructureType = 1000152000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_FEATURES_NV                  StructureType = 1000154000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV                StructureType = 1000154001
	STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_EXT                         StructureType = 1000158000
	STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT                              StructureType = 1000158001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT              StructureType = 1000158002
	STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT                  StructureType = 1000158003
	STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT              StructureType = 1000158004
	STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT                        StructureType = 1000158005
	STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT                                StructureType = 1000160000
	STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT                  StructureType = 1000160001
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT             StructureType = 1000161000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT                StructureType = 1000161001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT              StructureType = 1000161002
	STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT      StructureType = 1000161003
	STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT     StructureType = 1000161004
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV       StructureType = 1000164000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_FEATURES_NV                  StructureType = 1000164001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV                StructureType = 1000164002
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV      StructureType = 1000164005
	STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_NV                             StructureType = 1000165000
	STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV                           StructureType = 1000165001
	STRUCTURE_TYPE_GEOMETRY_NV                                                     StructureType = 1000165003
	STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV                                           StructureType = 1000165004
	STRUCTURE_TYPE_GEOMETRY_AABB_NV                                                StructureType = 1000165005
	STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV                      StructureType = 1000165006
	STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV                  StructureType = 1000165007
	STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV              StructureType = 1000165008
	STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PROPERTIES_NV                       StructureType = 1000165009
	STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV                         StructureType = 1000165011
	STRUCTURE_TYPE_ACCELERATION_STRUCTURE_INFO_NV                                  StructureType = 1000165012
	STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV        StructureType = 1000166000
	STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV      StructureType = 1000166001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT                StructureType = 1000170000
	STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT             StructureType = 1000170001
	STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT                    StructureType = 1000174000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR                       StructureType = 1000177000
	STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT                             StructureType = 1000178000
	STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT                              StructureType = 1000178001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT             StructureType = 1000178002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR                StructureType = 1000180000
	STRUCTURE_TYPE_PIPELINE_COMPILER_CONTROL_CREATE_INFO_AMD                       StructureType = 1000183000
	STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_EXT                                   StructureType = 1000184000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD                      StructureType = 1000185000
	STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD                    StructureType = 1000189000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT         StructureType = 1000190000
	STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT             StructureType = 1000190001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT           StructureType = 1000190002
	STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP                                         StructureType = 1000191000
	STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO_EXT                      StructureType = 1000192000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR                           StructureType = 1000196000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR                   StructureType = 1000197000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR            StructureType = 1000199000
	STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR                   StructureType = 1000199001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV          StructureType = 1000201000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV                         StructureType = 1000202000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV                       StructureType = 1000202001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV         StructureType = 1000203000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_FOOTPRINT_FEATURES_NV              StructureType = 1000204000
	STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV        StructureType = 1000205000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV                   StructureType = 1000205002
	STRUCTURE_TYPE_CHECKPOINT_DATA_NV                                              StructureType = 1000206000
	STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV                           StructureType = 1000206001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_FUNCTIONS_2_FEATURES_INTEL       StructureType = 1000209000
	STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL                                    StructureType = 1000210000
	STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL                           StructureType = 1000210001
	STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL                                   StructureType = 1000210002
	STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL                            StructureType = 1000210003
	STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL                                 StructureType = 1000210004
	STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL                    StructureType = 1000210005
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR                StructureType = 1000211000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT                     StructureType = 1000212000
	STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD                     StructureType = 1000213000
	STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD                    StructureType = 1000213001
	STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA                           StructureType = 1000214000
	STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT                                   StructureType = 1000217000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT               StructureType = 1000218000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT             StructureType = 1000218001
	STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT                StructureType = 1000218002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT                StructureType = 1000221000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES_EXT            StructureType = 1000225000
	STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT    StructureType = 1000225001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES_EXT              StructureType = 1000225002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD                    StructureType = 1000227000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD                    StructureType = 1000229000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT                    StructureType = 1000237000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT                    StructureType = 1000238000
	STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT                               StructureType = 1000238001
	STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR                              StructureType = 1000239000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_DEDICATED_ALLOCATION_IMAGE_ALIASING_FEATURES_NV StructureType = 1000240000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT              StructureType = 1000244000
	STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_EXT                                  StructureType = 1000244001
	STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT                           StructureType = 1000244002
	STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT                             StructureType = 1000246000
	STRUCTURE_TYPE_VALIDATION_FEATURES_EXT                                         StructureType = 1000247000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV                  StructureType = 1000249000
	STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV                                StructureType = 1000249001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_NV                StructureType = 1000249002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_COVERAGE_REDUCTION_MODE_FEATURES_NV             StructureType = 1000250000
	STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV                StructureType = 1000250001
	STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV                        StructureType = 1000250002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_INTERLOCK_FEATURES_EXT          StructureType = 1000251000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_IMAGE_ARRAYS_FEATURES_EXT                 StructureType = 1000252000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR     StructureType = 1000253000
	STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT                          StructureType = 1000255000
	STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT                  StructureType = 1000255002
	STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT                    StructureType = 1000255001
	STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT                                StructureType = 1000256000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_EXT                 StructureType = 1000259000
	STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_EXT               StructureType = 1000259001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES_EXT               StructureType = 1000259002
	STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT                   StructureType = 1000261000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES_EXT                   StructureType = 1000265000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_EXECUTABLE_PROPERTIES_FEATURES_KHR     StructureType = 1000269000
	STRUCTURE_TYPE_PIPELINE_INFO_KHR                                               StructureType = 1000269001
	STRUCTURE_TYPE_PIPELINE_EXECUTABLE_PROPERTIES_KHR                              StructureType = 1000269002
	STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INFO_KHR                                    StructureType = 1000269003
	STRUCTURE_TYPE_PIPELINE_EXECUTABLE_STATISTIC_KHR                               StructureType = 1000269004
	STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INTERNAL_REPRESENTATION_KHR                 StructureType = 1000269005
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES_EXT StructureType = 1000276000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_FEATURES_EXT             StructureType = 1000281000
	STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES_EXT           StructureType = 1000281001
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES                       StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETER_FEATURES                  StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
	STRUCTURE_TYPE_DEBUG_REPORT_CREATE_INFO_EXT                                    StructureType = STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT
	STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR                           StructureType = STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR                          StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR                        StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR                                  StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR                                StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR                                         StructureType = STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR                                   StructureType = STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR                         StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
	STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR                                   StructureType = STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR                         StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR                            StructureType = STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR                  StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
	STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR                                  StructureType = STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR                         StructureType = STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR                      StructureType = STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR                                    StructureType = STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
	STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR                               StructureType = STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR                        StructureType = STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR                         StructureType = STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR                            StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR                             StructureType = STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR                  StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO
	STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR                            StructureType = STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR                        StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
	STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR                                  StructureType = STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR                               StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR                          StructureType = STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
	STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR                           StructureType = STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
	STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR                                 StructureType = STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR                     StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
	STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR                               StructureType = STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR                                StructureType = STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT16_INT8_FEATURES_KHR                       StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR
	STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR                      StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR                      StructureType = STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
	STRUCTURE_TYPE_SURFACE_CAPABILITIES2_EXT                                       StructureType = STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT
	STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR                         StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR                                   StructureType = STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR                                    StructureType = STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR                   StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR             StructureType = STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR                                StructureType = STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
	STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR       StructureType = STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES_KHR                   StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES
	STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR                  StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES
	STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR                               StructureType = STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
	STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR                              StructureType = STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR                           StructureType = STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR                            StructureType = STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR                     StructureType = STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR                                       StructureType = STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR                          StructureType = STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR                        StructureType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR                               StructureType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
	STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR                                StructureType = STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
	STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR                        StructureType = STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR           StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR            StructureType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
	STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR                                     StructureType = STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
	STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR                                      StructureType = STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
	STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR                    StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR                               StructureType = STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
	STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_ADDRESS_FEATURES_EXT                     StructureType = STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT
	STRUCTURE_TYPE_BEGIN_RANGE                                                     StructureType = STRUCTURE_TYPE_APPLICATION_INFO
	STRUCTURE_TYPE_END_RANGE                                                       StructureType = STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO
	STRUCTURE_TYPE_RANGE_SIZE                                                      StructureType = (STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO - STRUCTURE_TYPE_APPLICATION_INFO + 1)
	STRUCTURE_TYPE_MAX_ENUM                                                        StructureType = 0x7FFFFFFF
)

func (x StructureType) String() string {
	switch x {
	case STRUCTURE_TYPE_APPLICATION_INFO:
		return "STRUCTURE_TYPE_APPLICATION_INFO"
	case STRUCTURE_TYPE_INSTANCE_CREATE_INFO:
		return "STRUCTURE_TYPE_INSTANCE_CREATE_INFO"
	case STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO:
		return "STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO"
	case STRUCTURE_TYPE_DEVICE_CREATE_INFO:
		return "STRUCTURE_TYPE_DEVICE_CREATE_INFO"
	case STRUCTURE_TYPE_SUBMIT_INFO:
		return "STRUCTURE_TYPE_SUBMIT_INFO"
	case STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO:
		return "STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO"
	case STRUCTURE_TYPE_MAPPED_MEMORY_RANGE:
		return "STRUCTURE_TYPE_MAPPED_MEMORY_RANGE"
	case STRUCTURE_TYPE_BIND_SPARSE_INFO:
		return "STRUCTURE_TYPE_BIND_SPARSE_INFO"
	case STRUCTURE_TYPE_FENCE_CREATE_INFO:
		return "STRUCTURE_TYPE_FENCE_CREATE_INFO"
	case STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO:
		return "STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO"
	case STRUCTURE_TYPE_EVENT_CREATE_INFO:
		return "STRUCTURE_TYPE_EVENT_CREATE_INFO"
	case STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO:
		return "STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO"
	case STRUCTURE_TYPE_BUFFER_CREATE_INFO:
		return "STRUCTURE_TYPE_BUFFER_CREATE_INFO"
	case STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO:
		return "STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO"
	case STRUCTURE_TYPE_IMAGE_CREATE_INFO:
		return "STRUCTURE_TYPE_IMAGE_CREATE_INFO"
	case STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO:
		return "STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO"
	case STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO:
		return "STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO:
		return "STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO"
	case STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO:
		return "STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO"
	case STRUCTURE_TYPE_SAMPLER_CREATE_INFO:
		return "STRUCTURE_TYPE_SAMPLER_CREATE_INFO"
	case STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO:
		return "STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO"
	case STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO:
		return "STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO"
	case STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO:
		return "STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO"
	case STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET:
		return "STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET"
	case STRUCTURE_TYPE_COPY_DESCRIPTOR_SET:
		return "STRUCTURE_TYPE_COPY_DESCRIPTOR_SET"
	case STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO:
		return "STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO"
	case STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO:
		return "STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO"
	case STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO:
		return "STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO"
	case STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO:
		return "STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO"
	case STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO:
		return "STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO"
	case STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO:
		return "STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO"
	case STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO:
		return "STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO"
	case STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER:
		return "STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER"
	case STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER:
		return "STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER"
	case STRUCTURE_TYPE_MEMORY_BARRIER:
		return "STRUCTURE_TYPE_MEMORY_BARRIER"
	case STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO:
		return "STRUCTURE_TYPE_LOADER_INSTANCE_CREATE_INFO"
	case STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO:
		return "STRUCTURE_TYPE_LOADER_DEVICE_CREATE_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES"
	case STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO:
		return "STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO"
	case STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO:
		return "STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES"
	case STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS:
		return "STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS"
	case STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO:
		return "STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO"
	case STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO:
		return "STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO"
	case STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO:
		return "STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO"
	case STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO:
		return "STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO"
	case STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO:
		return "STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO"
	case STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO:
		return "STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO"
	case STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO:
		return "STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO"
	case STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO:
		return "STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES"
	case STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO:
		return "STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO"
	case STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2:
		return "STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2"
	case STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2:
		return "STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2"
	case STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2:
		return "STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2"
	case STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2:
		return "STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2"
	case STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2:
		return "STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2"
	case STRUCTURE_TYPE_FORMAT_PROPERTIES_2:
		return "STRUCTURE_TYPE_FORMAT_PROPERTIES_2"
	case STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2:
		return "STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2"
	case STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2:
		return "STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2"
	case STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2:
		return "STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES"
	case STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO:
		return "STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO"
	case STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO:
		return "STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO"
	case STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO:
		return "STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO"
	case STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO:
		return "STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES"
	case STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO:
		return "STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES"
	case STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2:
		return "STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2"
	case STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO:
		return "STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO"
	case STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO:
		return "STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO"
	case STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO:
		return "STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO"
	case STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO:
		return "STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES"
	case STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES:
		return "STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES"
	case STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO:
		return "STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO"
	case STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES:
		return "STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO"
	case STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES:
		return "STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES"
	case STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO:
		return "STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO"
	case STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO:
		return "STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO"
	case STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO:
		return "STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO"
	case STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES:
		return "STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES"
	case STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO:
		return "STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO"
	case STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO:
		return "STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO"
	case STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES:
		return "STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES"
	case STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT:
		return "STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES"
	case STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_PRESENT_INFO_KHR:
		return "STRUCTURE_TYPE_PRESENT_INFO_KHR"
	case STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR:
		return "STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR"
	case STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR:
		return "STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR"
	case STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR:
		return "STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR"
	case STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR:
		return "STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR"
	case STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR:
		return "STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR"
	case STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_ANDROID_SURFACE_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD:
		return "STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD"
	case STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT:
		return "STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT"
	case STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT:
		return "STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT"
	case STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT:
		return "STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT"
	case STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV"
	case STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV:
		return "STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX:
		return "STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX"
	case STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD:
		return "STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD"
	case STRUCTURE_TYPE_STREAM_DESCRIPTOR_SURFACE_CREATE_INFO_GGP:
		return "STRUCTURE_TYPE_STREAM_DESCRIPTOR_SURFACE_CREATE_INFO_GGP"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV"
	case STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV:
		return "STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV"
	case STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV:
		return "STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV"
	case STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV:
		return "STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV"
	case STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV:
		return "STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV"
	case STRUCTURE_TYPE_VALIDATION_FLAGS_EXT:
		return "STRUCTURE_TYPE_VALIDATION_FLAGS_EXT"
	case STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN:
		return "STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES_EXT"
	case STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT:
		return "STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT"
	case STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR:
		return "STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR"
	case STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR:
		return "STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR"
	case STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR:
		return "STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR"
	case STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR:
		return "STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR"
	case STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR:
		return "STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR"
	case STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR:
		return "STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR"
	case STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR:
		return "STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR"
	case STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR:
		return "STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR"
	case STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT:
		return "STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT"
	case STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT:
		return "STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR"
	case STRUCTURE_TYPE_PRESENT_REGIONS_KHR:
		return "STRUCTURE_TYPE_PRESENT_REGIONS_KHR"
	case STRUCTURE_TYPE_OBJECT_TABLE_CREATE_INFO_NVX:
		return "STRUCTURE_TYPE_OBJECT_TABLE_CREATE_INFO_NVX"
	case STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NVX:
		return "STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NVX"
	case STRUCTURE_TYPE_CMD_PROCESS_COMMANDS_INFO_NVX:
		return "STRUCTURE_TYPE_CMD_PROCESS_COMMANDS_INFO_NVX"
	case STRUCTURE_TYPE_CMD_RESERVE_SPACE_FOR_COMMANDS_INFO_NVX:
		return "STRUCTURE_TYPE_CMD_RESERVE_SPACE_FOR_COMMANDS_INFO_NVX"
	case STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_LIMITS_NVX:
		return "STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_LIMITS_NVX"
	case STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_FEATURES_NVX:
		return "STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_FEATURES_NVX"
	case STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT:
		return "STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT"
	case STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT:
		return "STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT"
	case STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT:
		return "STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT"
	case STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT:
		return "STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT"
	case STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE:
		return "STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX"
	case STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_ENABLE_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_ENABLE_FEATURES_EXT"
	case STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_HDR_METADATA_EXT:
		return "STRUCTURE_TYPE_HDR_METADATA_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR"
	case STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR:
		return "STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR"
	case STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR:
		return "STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR"
	case STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR:
		return "STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR"
	case STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR:
		return "STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR"
	case STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR:
		return "STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR"
	case STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR:
		return "STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR"
	case STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR:
		return "STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR"
	case STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR:
		return "STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR"
	case STRUCTURE_TYPE_SUBPASS_END_INFO_KHR:
		return "STRUCTURE_TYPE_SUBPASS_END_INFO_KHR"
	case STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR:
		return "STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR"
	case STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR:
		return "STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR"
	case STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR:
		return "STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR"
	case STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR:
		return "STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR"
	case STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR:
		return "STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR"
	case STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR:
		return "STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR"
	case STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR:
		return "STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR"
	case STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR:
		return "STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR"
	case STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR:
		return "STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR"
	case STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR:
		return "STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR"
	case STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR:
		return "STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR"
	case STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK:
		return "STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK"
	case STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK:
		return "STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK"
	case STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT:
		return "STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT"
	case STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT:
		return "STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT"
	case STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT:
		return "STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT"
	case STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT:
		return "STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT"
	case STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID:
		return "STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID"
	case STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID:
		return "STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID"
	case STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID:
		return "STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID"
	case STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID:
		return "STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID"
	case STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID:
		return "STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID"
	case STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID:
		return "STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT"
	case STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT"
	case STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT:
		return "STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT"
	case STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT:
		return "STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT"
	case STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT:
		return "STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT"
	case STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT"
	case STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT"
	case STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR:
		return "STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_FEATURES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV"
	case STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_EXT:
		return "STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_EXT"
	case STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT"
	case STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT"
	case STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT"
	case STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT:
		return "STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT"
	case STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT:
		return "STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT"
	case STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_FEATURES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV"
	case STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_GEOMETRY_NV:
		return "STRUCTURE_TYPE_GEOMETRY_NV"
	case STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV:
		return "STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV"
	case STRUCTURE_TYPE_GEOMETRY_AABB_NV:
		return "STRUCTURE_TYPE_GEOMETRY_AABB_NV"
	case STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV:
		return "STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV"
	case STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV:
		return "STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV"
	case STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV:
		return "STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PROPERTIES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PROPERTIES_NV"
	case STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV"
	case STRUCTURE_TYPE_ACCELERATION_STRUCTURE_INFO_NV:
		return "STRUCTURE_TYPE_ACCELERATION_STRUCTURE_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV"
	case STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT"
	case STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT"
	case STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR"
	case STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT:
		return "STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT"
	case STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR"
	case STRUCTURE_TYPE_PIPELINE_COMPILER_CONTROL_CREATE_INFO_AMD:
		return "STRUCTURE_TYPE_PIPELINE_COMPILER_CONTROL_CREATE_INFO_AMD"
	case STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_EXT:
		return "STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD"
	case STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD:
		return "STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT"
	case STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP:
		return "STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP"
	case STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR"
	case STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR:
		return "STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_FOOTPRINT_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_FOOTPRINT_FEATURES_NV"
	case STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV"
	case STRUCTURE_TYPE_CHECKPOINT_DATA_NV:
		return "STRUCTURE_TYPE_CHECKPOINT_DATA_NV"
	case STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV:
		return "STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_FUNCTIONS_2_FEATURES_INTEL:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_FUNCTIONS_2_FEATURES_INTEL"
	case STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL:
		return "STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL"
	case STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL:
		return "STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL"
	case STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL:
		return "STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL"
	case STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL:
		return "STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL"
	case STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL:
		return "STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL"
	case STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL:
		return "STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT"
	case STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD:
		return "STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD"
	case STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD:
		return "STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD"
	case STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA:
		return "STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA"
	case STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT"
	case STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT"
	case STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT:
		return "STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT"
	case STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR:
		return "STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_DEDICATED_ALLOCATION_IMAGE_ALIASING_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_DEDICATED_ALLOCATION_IMAGE_ALIASING_FEATURES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT"
	case STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_EXT:
		return "STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_EXT"
	case STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_VALIDATION_FEATURES_EXT:
		return "STRUCTURE_TYPE_VALIDATION_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV"
	case STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV:
		return "STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_COVERAGE_REDUCTION_MODE_FEATURES_NV:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_COVERAGE_REDUCTION_MODE_FEATURES_NV"
	case STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV:
		return "STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV"
	case STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV:
		return "STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_INTERLOCK_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_INTERLOCK_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_IMAGE_ARRAYS_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_IMAGE_ARRAYS_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR"
	case STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT:
		return "STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT"
	case STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT:
		return "STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT"
	case STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT:
		return "STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT"
	case STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_EXT"
	case STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_EXT:
		return "STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_EXECUTABLE_PROPERTIES_FEATURES_KHR:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_EXECUTABLE_PROPERTIES_FEATURES_KHR"
	case STRUCTURE_TYPE_PIPELINE_INFO_KHR:
		return "STRUCTURE_TYPE_PIPELINE_INFO_KHR"
	case STRUCTURE_TYPE_PIPELINE_EXECUTABLE_PROPERTIES_KHR:
		return "STRUCTURE_TYPE_PIPELINE_EXECUTABLE_PROPERTIES_KHR"
	case STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INFO_KHR:
		return "STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INFO_KHR"
	case STRUCTURE_TYPE_PIPELINE_EXECUTABLE_STATISTIC_KHR:
		return "STRUCTURE_TYPE_PIPELINE_EXECUTABLE_STATISTIC_KHR"
	case STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INTERNAL_REPRESENTATION_KHR:
		return "STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INTERNAL_REPRESENTATION_KHR"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_FEATURES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_FEATURES_EXT"
	case STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES_EXT:
		return "STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES_EXT"
	case STRUCTURE_TYPE_MAX_ENUM:
		return "STRUCTURE_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SystemAllocationScope -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSystemAllocationScope.html
type SystemAllocationScope int32

const (
	SYSTEM_ALLOCATION_SCOPE_COMMAND     SystemAllocationScope = 0
	SYSTEM_ALLOCATION_SCOPE_OBJECT      SystemAllocationScope = 1
	SYSTEM_ALLOCATION_SCOPE_CACHE       SystemAllocationScope = 2
	SYSTEM_ALLOCATION_SCOPE_DEVICE      SystemAllocationScope = 3
	SYSTEM_ALLOCATION_SCOPE_INSTANCE    SystemAllocationScope = 4
	SYSTEM_ALLOCATION_SCOPE_BEGIN_RANGE SystemAllocationScope = SYSTEM_ALLOCATION_SCOPE_COMMAND
	SYSTEM_ALLOCATION_SCOPE_END_RANGE   SystemAllocationScope = SYSTEM_ALLOCATION_SCOPE_INSTANCE
	SYSTEM_ALLOCATION_SCOPE_RANGE_SIZE  SystemAllocationScope = (SYSTEM_ALLOCATION_SCOPE_INSTANCE - SYSTEM_ALLOCATION_SCOPE_COMMAND + 1)
	SYSTEM_ALLOCATION_SCOPE_MAX_ENUM    SystemAllocationScope = 0x7FFFFFFF
)

func (x SystemAllocationScope) String() string {
	switch x {
	case SYSTEM_ALLOCATION_SCOPE_COMMAND:
		return "SYSTEM_ALLOCATION_SCOPE_COMMAND"
	case SYSTEM_ALLOCATION_SCOPE_OBJECT:
		return "SYSTEM_ALLOCATION_SCOPE_OBJECT"
	case SYSTEM_ALLOCATION_SCOPE_CACHE:
		return "SYSTEM_ALLOCATION_SCOPE_CACHE"
	case SYSTEM_ALLOCATION_SCOPE_DEVICE:
		return "SYSTEM_ALLOCATION_SCOPE_DEVICE"
	case SYSTEM_ALLOCATION_SCOPE_INSTANCE:
		return "SYSTEM_ALLOCATION_SCOPE_INSTANCE"
	case SYSTEM_ALLOCATION_SCOPE_MAX_ENUM:
		return "SYSTEM_ALLOCATION_SCOPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// InternalAllocationType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkInternalAllocationType.html
type InternalAllocationType int32

const (
	INTERNAL_ALLOCATION_TYPE_EXECUTABLE  InternalAllocationType = 0
	INTERNAL_ALLOCATION_TYPE_BEGIN_RANGE InternalAllocationType = INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	INTERNAL_ALLOCATION_TYPE_END_RANGE   InternalAllocationType = INTERNAL_ALLOCATION_TYPE_EXECUTABLE
	INTERNAL_ALLOCATION_TYPE_RANGE_SIZE  InternalAllocationType = (INTERNAL_ALLOCATION_TYPE_EXECUTABLE - INTERNAL_ALLOCATION_TYPE_EXECUTABLE + 1)
	INTERNAL_ALLOCATION_TYPE_MAX_ENUM    InternalAllocationType = 0x7FFFFFFF
)

func (x InternalAllocationType) String() string {
	switch x {
	case INTERNAL_ALLOCATION_TYPE_EXECUTABLE:
		return "INTERNAL_ALLOCATION_TYPE_EXECUTABLE"
	case INTERNAL_ALLOCATION_TYPE_MAX_ENUM:
		return "INTERNAL_ALLOCATION_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// Format -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFormat.html
type Format int32

const (
	FORMAT_UNDEFINED                                      Format = 0
	FORMAT_R4G4_UNORM_PACK8                               Format = 1
	FORMAT_R4G4B4A4_UNORM_PACK16                          Format = 2
	FORMAT_B4G4R4A4_UNORM_PACK16                          Format = 3
	FORMAT_R5G6B5_UNORM_PACK16                            Format = 4
	FORMAT_B5G6R5_UNORM_PACK16                            Format = 5
	FORMAT_R5G5B5A1_UNORM_PACK16                          Format = 6
	FORMAT_B5G5R5A1_UNORM_PACK16                          Format = 7
	FORMAT_A1R5G5B5_UNORM_PACK16                          Format = 8
	FORMAT_R8_UNORM                                       Format = 9
	FORMAT_R8_SNORM                                       Format = 10
	FORMAT_R8_USCALED                                     Format = 11
	FORMAT_R8_SSCALED                                     Format = 12
	FORMAT_R8_UINT                                        Format = 13
	FORMAT_R8_SINT                                        Format = 14
	FORMAT_R8_SRGB                                        Format = 15
	FORMAT_R8G8_UNORM                                     Format = 16
	FORMAT_R8G8_SNORM                                     Format = 17
	FORMAT_R8G8_USCALED                                   Format = 18
	FORMAT_R8G8_SSCALED                                   Format = 19
	FORMAT_R8G8_UINT                                      Format = 20
	FORMAT_R8G8_SINT                                      Format = 21
	FORMAT_R8G8_SRGB                                      Format = 22
	FORMAT_R8G8B8_UNORM                                   Format = 23
	FORMAT_R8G8B8_SNORM                                   Format = 24
	FORMAT_R8G8B8_USCALED                                 Format = 25
	FORMAT_R8G8B8_SSCALED                                 Format = 26
	FORMAT_R8G8B8_UINT                                    Format = 27
	FORMAT_R8G8B8_SINT                                    Format = 28
	FORMAT_R8G8B8_SRGB                                    Format = 29
	FORMAT_B8G8R8_UNORM                                   Format = 30
	FORMAT_B8G8R8_SNORM                                   Format = 31
	FORMAT_B8G8R8_USCALED                                 Format = 32
	FORMAT_B8G8R8_SSCALED                                 Format = 33
	FORMAT_B8G8R8_UINT                                    Format = 34
	FORMAT_B8G8R8_SINT                                    Format = 35
	FORMAT_B8G8R8_SRGB                                    Format = 36
	FORMAT_R8G8B8A8_UNORM                                 Format = 37
	FORMAT_R8G8B8A8_SNORM                                 Format = 38
	FORMAT_R8G8B8A8_USCALED                               Format = 39
	FORMAT_R8G8B8A8_SSCALED                               Format = 40
	FORMAT_R8G8B8A8_UINT                                  Format = 41
	FORMAT_R8G8B8A8_SINT                                  Format = 42
	FORMAT_R8G8B8A8_SRGB                                  Format = 43
	FORMAT_B8G8R8A8_UNORM                                 Format = 44
	FORMAT_B8G8R8A8_SNORM                                 Format = 45
	FORMAT_B8G8R8A8_USCALED                               Format = 46
	FORMAT_B8G8R8A8_SSCALED                               Format = 47
	FORMAT_B8G8R8A8_UINT                                  Format = 48
	FORMAT_B8G8R8A8_SINT                                  Format = 49
	FORMAT_B8G8R8A8_SRGB                                  Format = 50
	FORMAT_A8B8G8R8_UNORM_PACK32                          Format = 51
	FORMAT_A8B8G8R8_SNORM_PACK32                          Format = 52
	FORMAT_A8B8G8R8_USCALED_PACK32                        Format = 53
	FORMAT_A8B8G8R8_SSCALED_PACK32                        Format = 54
	FORMAT_A8B8G8R8_UINT_PACK32                           Format = 55
	FORMAT_A8B8G8R8_SINT_PACK32                           Format = 56
	FORMAT_A8B8G8R8_SRGB_PACK32                           Format = 57
	FORMAT_A2R10G10B10_UNORM_PACK32                       Format = 58
	FORMAT_A2R10G10B10_SNORM_PACK32                       Format = 59
	FORMAT_A2R10G10B10_USCALED_PACK32                     Format = 60
	FORMAT_A2R10G10B10_SSCALED_PACK32                     Format = 61
	FORMAT_A2R10G10B10_UINT_PACK32                        Format = 62
	FORMAT_A2R10G10B10_SINT_PACK32                        Format = 63
	FORMAT_A2B10G10R10_UNORM_PACK32                       Format = 64
	FORMAT_A2B10G10R10_SNORM_PACK32                       Format = 65
	FORMAT_A2B10G10R10_USCALED_PACK32                     Format = 66
	FORMAT_A2B10G10R10_SSCALED_PACK32                     Format = 67
	FORMAT_A2B10G10R10_UINT_PACK32                        Format = 68
	FORMAT_A2B10G10R10_SINT_PACK32                        Format = 69
	FORMAT_R16_UNORM                                      Format = 70
	FORMAT_R16_SNORM                                      Format = 71
	FORMAT_R16_USCALED                                    Format = 72
	FORMAT_R16_SSCALED                                    Format = 73
	FORMAT_R16_UINT                                       Format = 74
	FORMAT_R16_SINT                                       Format = 75
	FORMAT_R16_SFLOAT                                     Format = 76
	FORMAT_R16G16_UNORM                                   Format = 77
	FORMAT_R16G16_SNORM                                   Format = 78
	FORMAT_R16G16_USCALED                                 Format = 79
	FORMAT_R16G16_SSCALED                                 Format = 80
	FORMAT_R16G16_UINT                                    Format = 81
	FORMAT_R16G16_SINT                                    Format = 82
	FORMAT_R16G16_SFLOAT                                  Format = 83
	FORMAT_R16G16B16_UNORM                                Format = 84
	FORMAT_R16G16B16_SNORM                                Format = 85
	FORMAT_R16G16B16_USCALED                              Format = 86
	FORMAT_R16G16B16_SSCALED                              Format = 87
	FORMAT_R16G16B16_UINT                                 Format = 88
	FORMAT_R16G16B16_SINT                                 Format = 89
	FORMAT_R16G16B16_SFLOAT                               Format = 90
	FORMAT_R16G16B16A16_UNORM                             Format = 91
	FORMAT_R16G16B16A16_SNORM                             Format = 92
	FORMAT_R16G16B16A16_USCALED                           Format = 93
	FORMAT_R16G16B16A16_SSCALED                           Format = 94
	FORMAT_R16G16B16A16_UINT                              Format = 95
	FORMAT_R16G16B16A16_SINT                              Format = 96
	FORMAT_R16G16B16A16_SFLOAT                            Format = 97
	FORMAT_R32_UINT                                       Format = 98
	FORMAT_R32_SINT                                       Format = 99
	FORMAT_R32_SFLOAT                                     Format = 100
	FORMAT_R32G32_UINT                                    Format = 101
	FORMAT_R32G32_SINT                                    Format = 102
	FORMAT_R32G32_SFLOAT                                  Format = 103
	FORMAT_R32G32B32_UINT                                 Format = 104
	FORMAT_R32G32B32_SINT                                 Format = 105
	FORMAT_R32G32B32_SFLOAT                               Format = 106
	FORMAT_R32G32B32A32_UINT                              Format = 107
	FORMAT_R32G32B32A32_SINT                              Format = 108
	FORMAT_R32G32B32A32_SFLOAT                            Format = 109
	FORMAT_R64_UINT                                       Format = 110
	FORMAT_R64_SINT                                       Format = 111
	FORMAT_R64_SFLOAT                                     Format = 112
	FORMAT_R64G64_UINT                                    Format = 113
	FORMAT_R64G64_SINT                                    Format = 114
	FORMAT_R64G64_SFLOAT                                  Format = 115
	FORMAT_R64G64B64_UINT                                 Format = 116
	FORMAT_R64G64B64_SINT                                 Format = 117
	FORMAT_R64G64B64_SFLOAT                               Format = 118
	FORMAT_R64G64B64A64_UINT                              Format = 119
	FORMAT_R64G64B64A64_SINT                              Format = 120
	FORMAT_R64G64B64A64_SFLOAT                            Format = 121
	FORMAT_B10G11R11_UFLOAT_PACK32                        Format = 122
	FORMAT_E5B9G9R9_UFLOAT_PACK32                         Format = 123
	FORMAT_D16_UNORM                                      Format = 124
	FORMAT_X8_D24_UNORM_PACK32                            Format = 125
	FORMAT_D32_SFLOAT                                     Format = 126
	FORMAT_S8_UINT                                        Format = 127
	FORMAT_D16_UNORM_S8_UINT                              Format = 128
	FORMAT_D24_UNORM_S8_UINT                              Format = 129
	FORMAT_D32_SFLOAT_S8_UINT                             Format = 130
	FORMAT_BC1_RGB_UNORM_BLOCK                            Format = 131
	FORMAT_BC1_RGB_SRGB_BLOCK                             Format = 132
	FORMAT_BC1_RGBA_UNORM_BLOCK                           Format = 133
	FORMAT_BC1_RGBA_SRGB_BLOCK                            Format = 134
	FORMAT_BC2_UNORM_BLOCK                                Format = 135
	FORMAT_BC2_SRGB_BLOCK                                 Format = 136
	FORMAT_BC3_UNORM_BLOCK                                Format = 137
	FORMAT_BC3_SRGB_BLOCK                                 Format = 138
	FORMAT_BC4_UNORM_BLOCK                                Format = 139
	FORMAT_BC4_SNORM_BLOCK                                Format = 140
	FORMAT_BC5_UNORM_BLOCK                                Format = 141
	FORMAT_BC5_SNORM_BLOCK                                Format = 142
	FORMAT_BC6H_UFLOAT_BLOCK                              Format = 143
	FORMAT_BC6H_SFLOAT_BLOCK                              Format = 144
	FORMAT_BC7_UNORM_BLOCK                                Format = 145
	FORMAT_BC7_SRGB_BLOCK                                 Format = 146
	FORMAT_ETC2_R8G8B8_UNORM_BLOCK                        Format = 147
	FORMAT_ETC2_R8G8B8_SRGB_BLOCK                         Format = 148
	FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK                      Format = 149
	FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK                       Format = 150
	FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK                      Format = 151
	FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK                       Format = 152
	FORMAT_EAC_R11_UNORM_BLOCK                            Format = 153
	FORMAT_EAC_R11_SNORM_BLOCK                            Format = 154
	FORMAT_EAC_R11G11_UNORM_BLOCK                         Format = 155
	FORMAT_EAC_R11G11_SNORM_BLOCK                         Format = 156
	FORMAT_ASTC_4x4_UNORM_BLOCK                           Format = 157
	FORMAT_ASTC_4x4_SRGB_BLOCK                            Format = 158
	FORMAT_ASTC_5x4_UNORM_BLOCK                           Format = 159
	FORMAT_ASTC_5x4_SRGB_BLOCK                            Format = 160
	FORMAT_ASTC_5x5_UNORM_BLOCK                           Format = 161
	FORMAT_ASTC_5x5_SRGB_BLOCK                            Format = 162
	FORMAT_ASTC_6x5_UNORM_BLOCK                           Format = 163
	FORMAT_ASTC_6x5_SRGB_BLOCK                            Format = 164
	FORMAT_ASTC_6x6_UNORM_BLOCK                           Format = 165
	FORMAT_ASTC_6x6_SRGB_BLOCK                            Format = 166
	FORMAT_ASTC_8x5_UNORM_BLOCK                           Format = 167
	FORMAT_ASTC_8x5_SRGB_BLOCK                            Format = 168
	FORMAT_ASTC_8x6_UNORM_BLOCK                           Format = 169
	FORMAT_ASTC_8x6_SRGB_BLOCK                            Format = 170
	FORMAT_ASTC_8x8_UNORM_BLOCK                           Format = 171
	FORMAT_ASTC_8x8_SRGB_BLOCK                            Format = 172
	FORMAT_ASTC_10x5_UNORM_BLOCK                          Format = 173
	FORMAT_ASTC_10x5_SRGB_BLOCK                           Format = 174
	FORMAT_ASTC_10x6_UNORM_BLOCK                          Format = 175
	FORMAT_ASTC_10x6_SRGB_BLOCK                           Format = 176
	FORMAT_ASTC_10x8_UNORM_BLOCK                          Format = 177
	FORMAT_ASTC_10x8_SRGB_BLOCK                           Format = 178
	FORMAT_ASTC_10x10_UNORM_BLOCK                         Format = 179
	FORMAT_ASTC_10x10_SRGB_BLOCK                          Format = 180
	FORMAT_ASTC_12x10_UNORM_BLOCK                         Format = 181
	FORMAT_ASTC_12x10_SRGB_BLOCK                          Format = 182
	FORMAT_ASTC_12x12_UNORM_BLOCK                         Format = 183
	FORMAT_ASTC_12x12_SRGB_BLOCK                          Format = 184
	FORMAT_G8B8G8R8_422_UNORM                             Format = 1000156000
	FORMAT_B8G8R8G8_422_UNORM                             Format = 1000156001
	FORMAT_G8_B8_R8_3PLANE_420_UNORM                      Format = 1000156002
	FORMAT_G8_B8R8_2PLANE_420_UNORM                       Format = 1000156003
	FORMAT_G8_B8_R8_3PLANE_422_UNORM                      Format = 1000156004
	FORMAT_G8_B8R8_2PLANE_422_UNORM                       Format = 1000156005
	FORMAT_G8_B8_R8_3PLANE_444_UNORM                      Format = 1000156006
	FORMAT_R10X6_UNORM_PACK16                             Format = 1000156007
	FORMAT_R10X6G10X6_UNORM_2PACK16                       Format = 1000156008
	FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16             Format = 1000156009
	FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16         Format = 1000156010
	FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16         Format = 1000156011
	FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16     Format = 1000156012
	FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16      Format = 1000156013
	FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16     Format = 1000156014
	FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16      Format = 1000156015
	FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16     Format = 1000156016
	FORMAT_R12X4_UNORM_PACK16                             Format = 1000156017
	FORMAT_R12X4G12X4_UNORM_2PACK16                       Format = 1000156018
	FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16             Format = 1000156019
	FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16         Format = 1000156020
	FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16         Format = 1000156021
	FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16     Format = 1000156022
	FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16      Format = 1000156023
	FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16     Format = 1000156024
	FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16      Format = 1000156025
	FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16     Format = 1000156026
	FORMAT_G16B16G16R16_422_UNORM                         Format = 1000156027
	FORMAT_B16G16R16G16_422_UNORM                         Format = 1000156028
	FORMAT_G16_B16_R16_3PLANE_420_UNORM                   Format = 1000156029
	FORMAT_G16_B16R16_2PLANE_420_UNORM                    Format = 1000156030
	FORMAT_G16_B16_R16_3PLANE_422_UNORM                   Format = 1000156031
	FORMAT_G16_B16R16_2PLANE_422_UNORM                    Format = 1000156032
	FORMAT_G16_B16_R16_3PLANE_444_UNORM                   Format = 1000156033
	FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG                    Format = 1000054000
	FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG                    Format = 1000054001
	FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG                    Format = 1000054002
	FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG                    Format = 1000054003
	FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG                     Format = 1000054004
	FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG                     Format = 1000054005
	FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG                     Format = 1000054006
	FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG                     Format = 1000054007
	FORMAT_ASTC_4x4_SFLOAT_BLOCK_EXT                      Format = 1000066000
	FORMAT_ASTC_5x4_SFLOAT_BLOCK_EXT                      Format = 1000066001
	FORMAT_ASTC_5x5_SFLOAT_BLOCK_EXT                      Format = 1000066002
	FORMAT_ASTC_6x5_SFLOAT_BLOCK_EXT                      Format = 1000066003
	FORMAT_ASTC_6x6_SFLOAT_BLOCK_EXT                      Format = 1000066004
	FORMAT_ASTC_8x5_SFLOAT_BLOCK_EXT                      Format = 1000066005
	FORMAT_ASTC_8x6_SFLOAT_BLOCK_EXT                      Format = 1000066006
	FORMAT_ASTC_8x8_SFLOAT_BLOCK_EXT                      Format = 1000066007
	FORMAT_ASTC_10x5_SFLOAT_BLOCK_EXT                     Format = 1000066008
	FORMAT_ASTC_10x6_SFLOAT_BLOCK_EXT                     Format = 1000066009
	FORMAT_ASTC_10x8_SFLOAT_BLOCK_EXT                     Format = 1000066010
	FORMAT_ASTC_10x10_SFLOAT_BLOCK_EXT                    Format = 1000066011
	FORMAT_ASTC_12x10_SFLOAT_BLOCK_EXT                    Format = 1000066012
	FORMAT_ASTC_12x12_SFLOAT_BLOCK_EXT                    Format = 1000066013
	FORMAT_G8B8G8R8_422_UNORM_KHR                         Format = FORMAT_G8B8G8R8_422_UNORM
	FORMAT_B8G8R8G8_422_UNORM_KHR                         Format = FORMAT_B8G8R8G8_422_UNORM
	FORMAT_G8_B8_R8_3PLANE_420_UNORM_KHR                  Format = FORMAT_G8_B8_R8_3PLANE_420_UNORM
	FORMAT_G8_B8R8_2PLANE_420_UNORM_KHR                   Format = FORMAT_G8_B8R8_2PLANE_420_UNORM
	FORMAT_G8_B8_R8_3PLANE_422_UNORM_KHR                  Format = FORMAT_G8_B8_R8_3PLANE_422_UNORM
	FORMAT_G8_B8R8_2PLANE_422_UNORM_KHR                   Format = FORMAT_G8_B8R8_2PLANE_422_UNORM
	FORMAT_G8_B8_R8_3PLANE_444_UNORM_KHR                  Format = FORMAT_G8_B8_R8_3PLANE_444_UNORM
	FORMAT_R10X6_UNORM_PACK16_KHR                         Format = FORMAT_R10X6_UNORM_PACK16
	FORMAT_R10X6G10X6_UNORM_2PACK16_KHR                   Format = FORMAT_R10X6G10X6_UNORM_2PACK16
	FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16_KHR         Format = FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16
	FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16_KHR     Format = FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16
	FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16_KHR     Format = FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16
	FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16_KHR Format = FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16
	FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16_KHR  Format = FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16
	FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16_KHR Format = FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16
	FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16_KHR  Format = FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16
	FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16_KHR Format = FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16
	FORMAT_R12X4_UNORM_PACK16_KHR                         Format = FORMAT_R12X4_UNORM_PACK16
	FORMAT_R12X4G12X4_UNORM_2PACK16_KHR                   Format = FORMAT_R12X4G12X4_UNORM_2PACK16
	FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16_KHR         Format = FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16
	FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16_KHR     Format = FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16
	FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR     Format = FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16
	FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16_KHR Format = FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16
	FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16_KHR  Format = FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16
	FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16_KHR Format = FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16
	FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16_KHR  Format = FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16
	FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16_KHR Format = FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16
	FORMAT_G16B16G16R16_422_UNORM_KHR                     Format = FORMAT_G16B16G16R16_422_UNORM
	FORMAT_B16G16R16G16_422_UNORM_KHR                     Format = FORMAT_B16G16R16G16_422_UNORM
	FORMAT_G16_B16_R16_3PLANE_420_UNORM_KHR               Format = FORMAT_G16_B16_R16_3PLANE_420_UNORM
	FORMAT_G16_B16R16_2PLANE_420_UNORM_KHR                Format = FORMAT_G16_B16R16_2PLANE_420_UNORM
	FORMAT_G16_B16_R16_3PLANE_422_UNORM_KHR               Format = FORMAT_G16_B16_R16_3PLANE_422_UNORM
	FORMAT_G16_B16R16_2PLANE_422_UNORM_KHR                Format = FORMAT_G16_B16R16_2PLANE_422_UNORM
	FORMAT_G16_B16_R16_3PLANE_444_UNORM_KHR               Format = FORMAT_G16_B16_R16_3PLANE_444_UNORM
	FORMAT_BEGIN_RANGE                                    Format = FORMAT_UNDEFINED
	FORMAT_END_RANGE                                      Format = FORMAT_ASTC_12x12_SRGB_BLOCK
	FORMAT_RANGE_SIZE                                     Format = (FORMAT_ASTC_12x12_SRGB_BLOCK - FORMAT_UNDEFINED + 1)
	FORMAT_MAX_ENUM                                       Format = 0x7FFFFFFF
)

func (x Format) String() string {
	switch x {
	case FORMAT_UNDEFINED:
		return "FORMAT_UNDEFINED"
	case FORMAT_R4G4_UNORM_PACK8:
		return "FORMAT_R4G4_UNORM_PACK8"
	case FORMAT_R4G4B4A4_UNORM_PACK16:
		return "FORMAT_R4G4B4A4_UNORM_PACK16"
	case FORMAT_B4G4R4A4_UNORM_PACK16:
		return "FORMAT_B4G4R4A4_UNORM_PACK16"
	case FORMAT_R5G6B5_UNORM_PACK16:
		return "FORMAT_R5G6B5_UNORM_PACK16"
	case FORMAT_B5G6R5_UNORM_PACK16:
		return "FORMAT_B5G6R5_UNORM_PACK16"
	case FORMAT_R5G5B5A1_UNORM_PACK16:
		return "FORMAT_R5G5B5A1_UNORM_PACK16"
	case FORMAT_B5G5R5A1_UNORM_PACK16:
		return "FORMAT_B5G5R5A1_UNORM_PACK16"
	case FORMAT_A1R5G5B5_UNORM_PACK16:
		return "FORMAT_A1R5G5B5_UNORM_PACK16"
	case FORMAT_R8_UNORM:
		return "FORMAT_R8_UNORM"
	case FORMAT_R8_SNORM:
		return "FORMAT_R8_SNORM"
	case FORMAT_R8_USCALED:
		return "FORMAT_R8_USCALED"
	case FORMAT_R8_SSCALED:
		return "FORMAT_R8_SSCALED"
	case FORMAT_R8_UINT:
		return "FORMAT_R8_UINT"
	case FORMAT_R8_SINT:
		return "FORMAT_R8_SINT"
	case FORMAT_R8_SRGB:
		return "FORMAT_R8_SRGB"
	case FORMAT_R8G8_UNORM:
		return "FORMAT_R8G8_UNORM"
	case FORMAT_R8G8_SNORM:
		return "FORMAT_R8G8_SNORM"
	case FORMAT_R8G8_USCALED:
		return "FORMAT_R8G8_USCALED"
	case FORMAT_R8G8_SSCALED:
		return "FORMAT_R8G8_SSCALED"
	case FORMAT_R8G8_UINT:
		return "FORMAT_R8G8_UINT"
	case FORMAT_R8G8_SINT:
		return "FORMAT_R8G8_SINT"
	case FORMAT_R8G8_SRGB:
		return "FORMAT_R8G8_SRGB"
	case FORMAT_R8G8B8_UNORM:
		return "FORMAT_R8G8B8_UNORM"
	case FORMAT_R8G8B8_SNORM:
		return "FORMAT_R8G8B8_SNORM"
	case FORMAT_R8G8B8_USCALED:
		return "FORMAT_R8G8B8_USCALED"
	case FORMAT_R8G8B8_SSCALED:
		return "FORMAT_R8G8B8_SSCALED"
	case FORMAT_R8G8B8_UINT:
		return "FORMAT_R8G8B8_UINT"
	case FORMAT_R8G8B8_SINT:
		return "FORMAT_R8G8B8_SINT"
	case FORMAT_R8G8B8_SRGB:
		return "FORMAT_R8G8B8_SRGB"
	case FORMAT_B8G8R8_UNORM:
		return "FORMAT_B8G8R8_UNORM"
	case FORMAT_B8G8R8_SNORM:
		return "FORMAT_B8G8R8_SNORM"
	case FORMAT_B8G8R8_USCALED:
		return "FORMAT_B8G8R8_USCALED"
	case FORMAT_B8G8R8_SSCALED:
		return "FORMAT_B8G8R8_SSCALED"
	case FORMAT_B8G8R8_UINT:
		return "FORMAT_B8G8R8_UINT"
	case FORMAT_B8G8R8_SINT:
		return "FORMAT_B8G8R8_SINT"
	case FORMAT_B8G8R8_SRGB:
		return "FORMAT_B8G8R8_SRGB"
	case FORMAT_R8G8B8A8_UNORM:
		return "FORMAT_R8G8B8A8_UNORM"
	case FORMAT_R8G8B8A8_SNORM:
		return "FORMAT_R8G8B8A8_SNORM"
	case FORMAT_R8G8B8A8_USCALED:
		return "FORMAT_R8G8B8A8_USCALED"
	case FORMAT_R8G8B8A8_SSCALED:
		return "FORMAT_R8G8B8A8_SSCALED"
	case FORMAT_R8G8B8A8_UINT:
		return "FORMAT_R8G8B8A8_UINT"
	case FORMAT_R8G8B8A8_SINT:
		return "FORMAT_R8G8B8A8_SINT"
	case FORMAT_R8G8B8A8_SRGB:
		return "FORMAT_R8G8B8A8_SRGB"
	case FORMAT_B8G8R8A8_UNORM:
		return "FORMAT_B8G8R8A8_UNORM"
	case FORMAT_B8G8R8A8_SNORM:
		return "FORMAT_B8G8R8A8_SNORM"
	case FORMAT_B8G8R8A8_USCALED:
		return "FORMAT_B8G8R8A8_USCALED"
	case FORMAT_B8G8R8A8_SSCALED:
		return "FORMAT_B8G8R8A8_SSCALED"
	case FORMAT_B8G8R8A8_UINT:
		return "FORMAT_B8G8R8A8_UINT"
	case FORMAT_B8G8R8A8_SINT:
		return "FORMAT_B8G8R8A8_SINT"
	case FORMAT_B8G8R8A8_SRGB:
		return "FORMAT_B8G8R8A8_SRGB"
	case FORMAT_A8B8G8R8_UNORM_PACK32:
		return "FORMAT_A8B8G8R8_UNORM_PACK32"
	case FORMAT_A8B8G8R8_SNORM_PACK32:
		return "FORMAT_A8B8G8R8_SNORM_PACK32"
	case FORMAT_A8B8G8R8_USCALED_PACK32:
		return "FORMAT_A8B8G8R8_USCALED_PACK32"
	case FORMAT_A8B8G8R8_SSCALED_PACK32:
		return "FORMAT_A8B8G8R8_SSCALED_PACK32"
	case FORMAT_A8B8G8R8_UINT_PACK32:
		return "FORMAT_A8B8G8R8_UINT_PACK32"
	case FORMAT_A8B8G8R8_SINT_PACK32:
		return "FORMAT_A8B8G8R8_SINT_PACK32"
	case FORMAT_A8B8G8R8_SRGB_PACK32:
		return "FORMAT_A8B8G8R8_SRGB_PACK32"
	case FORMAT_A2R10G10B10_UNORM_PACK32:
		return "FORMAT_A2R10G10B10_UNORM_PACK32"
	case FORMAT_A2R10G10B10_SNORM_PACK32:
		return "FORMAT_A2R10G10B10_SNORM_PACK32"
	case FORMAT_A2R10G10B10_USCALED_PACK32:
		return "FORMAT_A2R10G10B10_USCALED_PACK32"
	case FORMAT_A2R10G10B10_SSCALED_PACK32:
		return "FORMAT_A2R10G10B10_SSCALED_PACK32"
	case FORMAT_A2R10G10B10_UINT_PACK32:
		return "FORMAT_A2R10G10B10_UINT_PACK32"
	case FORMAT_A2R10G10B10_SINT_PACK32:
		return "FORMAT_A2R10G10B10_SINT_PACK32"
	case FORMAT_A2B10G10R10_UNORM_PACK32:
		return "FORMAT_A2B10G10R10_UNORM_PACK32"
	case FORMAT_A2B10G10R10_SNORM_PACK32:
		return "FORMAT_A2B10G10R10_SNORM_PACK32"
	case FORMAT_A2B10G10R10_USCALED_PACK32:
		return "FORMAT_A2B10G10R10_USCALED_PACK32"
	case FORMAT_A2B10G10R10_SSCALED_PACK32:
		return "FORMAT_A2B10G10R10_SSCALED_PACK32"
	case FORMAT_A2B10G10R10_UINT_PACK32:
		return "FORMAT_A2B10G10R10_UINT_PACK32"
	case FORMAT_A2B10G10R10_SINT_PACK32:
		return "FORMAT_A2B10G10R10_SINT_PACK32"
	case FORMAT_R16_UNORM:
		return "FORMAT_R16_UNORM"
	case FORMAT_R16_SNORM:
		return "FORMAT_R16_SNORM"
	case FORMAT_R16_USCALED:
		return "FORMAT_R16_USCALED"
	case FORMAT_R16_SSCALED:
		return "FORMAT_R16_SSCALED"
	case FORMAT_R16_UINT:
		return "FORMAT_R16_UINT"
	case FORMAT_R16_SINT:
		return "FORMAT_R16_SINT"
	case FORMAT_R16_SFLOAT:
		return "FORMAT_R16_SFLOAT"
	case FORMAT_R16G16_UNORM:
		return "FORMAT_R16G16_UNORM"
	case FORMAT_R16G16_SNORM:
		return "FORMAT_R16G16_SNORM"
	case FORMAT_R16G16_USCALED:
		return "FORMAT_R16G16_USCALED"
	case FORMAT_R16G16_SSCALED:
		return "FORMAT_R16G16_SSCALED"
	case FORMAT_R16G16_UINT:
		return "FORMAT_R16G16_UINT"
	case FORMAT_R16G16_SINT:
		return "FORMAT_R16G16_SINT"
	case FORMAT_R16G16_SFLOAT:
		return "FORMAT_R16G16_SFLOAT"
	case FORMAT_R16G16B16_UNORM:
		return "FORMAT_R16G16B16_UNORM"
	case FORMAT_R16G16B16_SNORM:
		return "FORMAT_R16G16B16_SNORM"
	case FORMAT_R16G16B16_USCALED:
		return "FORMAT_R16G16B16_USCALED"
	case FORMAT_R16G16B16_SSCALED:
		return "FORMAT_R16G16B16_SSCALED"
	case FORMAT_R16G16B16_UINT:
		return "FORMAT_R16G16B16_UINT"
	case FORMAT_R16G16B16_SINT:
		return "FORMAT_R16G16B16_SINT"
	case FORMAT_R16G16B16_SFLOAT:
		return "FORMAT_R16G16B16_SFLOAT"
	case FORMAT_R16G16B16A16_UNORM:
		return "FORMAT_R16G16B16A16_UNORM"
	case FORMAT_R16G16B16A16_SNORM:
		return "FORMAT_R16G16B16A16_SNORM"
	case FORMAT_R16G16B16A16_USCALED:
		return "FORMAT_R16G16B16A16_USCALED"
	case FORMAT_R16G16B16A16_SSCALED:
		return "FORMAT_R16G16B16A16_SSCALED"
	case FORMAT_R16G16B16A16_UINT:
		return "FORMAT_R16G16B16A16_UINT"
	case FORMAT_R16G16B16A16_SINT:
		return "FORMAT_R16G16B16A16_SINT"
	case FORMAT_R16G16B16A16_SFLOAT:
		return "FORMAT_R16G16B16A16_SFLOAT"
	case FORMAT_R32_UINT:
		return "FORMAT_R32_UINT"
	case FORMAT_R32_SINT:
		return "FORMAT_R32_SINT"
	case FORMAT_R32_SFLOAT:
		return "FORMAT_R32_SFLOAT"
	case FORMAT_R32G32_UINT:
		return "FORMAT_R32G32_UINT"
	case FORMAT_R32G32_SINT:
		return "FORMAT_R32G32_SINT"
	case FORMAT_R32G32_SFLOAT:
		return "FORMAT_R32G32_SFLOAT"
	case FORMAT_R32G32B32_UINT:
		return "FORMAT_R32G32B32_UINT"
	case FORMAT_R32G32B32_SINT:
		return "FORMAT_R32G32B32_SINT"
	case FORMAT_R32G32B32_SFLOAT:
		return "FORMAT_R32G32B32_SFLOAT"
	case FORMAT_R32G32B32A32_UINT:
		return "FORMAT_R32G32B32A32_UINT"
	case FORMAT_R32G32B32A32_SINT:
		return "FORMAT_R32G32B32A32_SINT"
	case FORMAT_R32G32B32A32_SFLOAT:
		return "FORMAT_R32G32B32A32_SFLOAT"
	case FORMAT_R64_UINT:
		return "FORMAT_R64_UINT"
	case FORMAT_R64_SINT:
		return "FORMAT_R64_SINT"
	case FORMAT_R64_SFLOAT:
		return "FORMAT_R64_SFLOAT"
	case FORMAT_R64G64_UINT:
		return "FORMAT_R64G64_UINT"
	case FORMAT_R64G64_SINT:
		return "FORMAT_R64G64_SINT"
	case FORMAT_R64G64_SFLOAT:
		return "FORMAT_R64G64_SFLOAT"
	case FORMAT_R64G64B64_UINT:
		return "FORMAT_R64G64B64_UINT"
	case FORMAT_R64G64B64_SINT:
		return "FORMAT_R64G64B64_SINT"
	case FORMAT_R64G64B64_SFLOAT:
		return "FORMAT_R64G64B64_SFLOAT"
	case FORMAT_R64G64B64A64_UINT:
		return "FORMAT_R64G64B64A64_UINT"
	case FORMAT_R64G64B64A64_SINT:
		return "FORMAT_R64G64B64A64_SINT"
	case FORMAT_R64G64B64A64_SFLOAT:
		return "FORMAT_R64G64B64A64_SFLOAT"
	case FORMAT_B10G11R11_UFLOAT_PACK32:
		return "FORMAT_B10G11R11_UFLOAT_PACK32"
	case FORMAT_E5B9G9R9_UFLOAT_PACK32:
		return "FORMAT_E5B9G9R9_UFLOAT_PACK32"
	case FORMAT_D16_UNORM:
		return "FORMAT_D16_UNORM"
	case FORMAT_X8_D24_UNORM_PACK32:
		return "FORMAT_X8_D24_UNORM_PACK32"
	case FORMAT_D32_SFLOAT:
		return "FORMAT_D32_SFLOAT"
	case FORMAT_S8_UINT:
		return "FORMAT_S8_UINT"
	case FORMAT_D16_UNORM_S8_UINT:
		return "FORMAT_D16_UNORM_S8_UINT"
	case FORMAT_D24_UNORM_S8_UINT:
		return "FORMAT_D24_UNORM_S8_UINT"
	case FORMAT_D32_SFLOAT_S8_UINT:
		return "FORMAT_D32_SFLOAT_S8_UINT"
	case FORMAT_BC1_RGB_UNORM_BLOCK:
		return "FORMAT_BC1_RGB_UNORM_BLOCK"
	case FORMAT_BC1_RGB_SRGB_BLOCK:
		return "FORMAT_BC1_RGB_SRGB_BLOCK"
	case FORMAT_BC1_RGBA_UNORM_BLOCK:
		return "FORMAT_BC1_RGBA_UNORM_BLOCK"
	case FORMAT_BC1_RGBA_SRGB_BLOCK:
		return "FORMAT_BC1_RGBA_SRGB_BLOCK"
	case FORMAT_BC2_UNORM_BLOCK:
		return "FORMAT_BC2_UNORM_BLOCK"
	case FORMAT_BC2_SRGB_BLOCK:
		return "FORMAT_BC2_SRGB_BLOCK"
	case FORMAT_BC3_UNORM_BLOCK:
		return "FORMAT_BC3_UNORM_BLOCK"
	case FORMAT_BC3_SRGB_BLOCK:
		return "FORMAT_BC3_SRGB_BLOCK"
	case FORMAT_BC4_UNORM_BLOCK:
		return "FORMAT_BC4_UNORM_BLOCK"
	case FORMAT_BC4_SNORM_BLOCK:
		return "FORMAT_BC4_SNORM_BLOCK"
	case FORMAT_BC5_UNORM_BLOCK:
		return "FORMAT_BC5_UNORM_BLOCK"
	case FORMAT_BC5_SNORM_BLOCK:
		return "FORMAT_BC5_SNORM_BLOCK"
	case FORMAT_BC6H_UFLOAT_BLOCK:
		return "FORMAT_BC6H_UFLOAT_BLOCK"
	case FORMAT_BC6H_SFLOAT_BLOCK:
		return "FORMAT_BC6H_SFLOAT_BLOCK"
	case FORMAT_BC7_UNORM_BLOCK:
		return "FORMAT_BC7_UNORM_BLOCK"
	case FORMAT_BC7_SRGB_BLOCK:
		return "FORMAT_BC7_SRGB_BLOCK"
	case FORMAT_ETC2_R8G8B8_UNORM_BLOCK:
		return "FORMAT_ETC2_R8G8B8_UNORM_BLOCK"
	case FORMAT_ETC2_R8G8B8_SRGB_BLOCK:
		return "FORMAT_ETC2_R8G8B8_SRGB_BLOCK"
	case FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK:
		return "FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK"
	case FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK:
		return "FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK"
	case FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK:
		return "FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK"
	case FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK:
		return "FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK"
	case FORMAT_EAC_R11_UNORM_BLOCK:
		return "FORMAT_EAC_R11_UNORM_BLOCK"
	case FORMAT_EAC_R11_SNORM_BLOCK:
		return "FORMAT_EAC_R11_SNORM_BLOCK"
	case FORMAT_EAC_R11G11_UNORM_BLOCK:
		return "FORMAT_EAC_R11G11_UNORM_BLOCK"
	case FORMAT_EAC_R11G11_SNORM_BLOCK:
		return "FORMAT_EAC_R11G11_SNORM_BLOCK"
	case FORMAT_ASTC_4x4_UNORM_BLOCK:
		return "FORMAT_ASTC_4x4_UNORM_BLOCK"
	case FORMAT_ASTC_4x4_SRGB_BLOCK:
		return "FORMAT_ASTC_4x4_SRGB_BLOCK"
	case FORMAT_ASTC_5x4_UNORM_BLOCK:
		return "FORMAT_ASTC_5x4_UNORM_BLOCK"
	case FORMAT_ASTC_5x4_SRGB_BLOCK:
		return "FORMAT_ASTC_5x4_SRGB_BLOCK"
	case FORMAT_ASTC_5x5_UNORM_BLOCK:
		return "FORMAT_ASTC_5x5_UNORM_BLOCK"
	case FORMAT_ASTC_5x5_SRGB_BLOCK:
		return "FORMAT_ASTC_5x5_SRGB_BLOCK"
	case FORMAT_ASTC_6x5_UNORM_BLOCK:
		return "FORMAT_ASTC_6x5_UNORM_BLOCK"
	case FORMAT_ASTC_6x5_SRGB_BLOCK:
		return "FORMAT_ASTC_6x5_SRGB_BLOCK"
	case FORMAT_ASTC_6x6_UNORM_BLOCK:
		return "FORMAT_ASTC_6x6_UNORM_BLOCK"
	case FORMAT_ASTC_6x6_SRGB_BLOCK:
		return "FORMAT_ASTC_6x6_SRGB_BLOCK"
	case FORMAT_ASTC_8x5_UNORM_BLOCK:
		return "FORMAT_ASTC_8x5_UNORM_BLOCK"
	case FORMAT_ASTC_8x5_SRGB_BLOCK:
		return "FORMAT_ASTC_8x5_SRGB_BLOCK"
	case FORMAT_ASTC_8x6_UNORM_BLOCK:
		return "FORMAT_ASTC_8x6_UNORM_BLOCK"
	case FORMAT_ASTC_8x6_SRGB_BLOCK:
		return "FORMAT_ASTC_8x6_SRGB_BLOCK"
	case FORMAT_ASTC_8x8_UNORM_BLOCK:
		return "FORMAT_ASTC_8x8_UNORM_BLOCK"
	case FORMAT_ASTC_8x8_SRGB_BLOCK:
		return "FORMAT_ASTC_8x8_SRGB_BLOCK"
	case FORMAT_ASTC_10x5_UNORM_BLOCK:
		return "FORMAT_ASTC_10x5_UNORM_BLOCK"
	case FORMAT_ASTC_10x5_SRGB_BLOCK:
		return "FORMAT_ASTC_10x5_SRGB_BLOCK"
	case FORMAT_ASTC_10x6_UNORM_BLOCK:
		return "FORMAT_ASTC_10x6_UNORM_BLOCK"
	case FORMAT_ASTC_10x6_SRGB_BLOCK:
		return "FORMAT_ASTC_10x6_SRGB_BLOCK"
	case FORMAT_ASTC_10x8_UNORM_BLOCK:
		return "FORMAT_ASTC_10x8_UNORM_BLOCK"
	case FORMAT_ASTC_10x8_SRGB_BLOCK:
		return "FORMAT_ASTC_10x8_SRGB_BLOCK"
	case FORMAT_ASTC_10x10_UNORM_BLOCK:
		return "FORMAT_ASTC_10x10_UNORM_BLOCK"
	case FORMAT_ASTC_10x10_SRGB_BLOCK:
		return "FORMAT_ASTC_10x10_SRGB_BLOCK"
	case FORMAT_ASTC_12x10_UNORM_BLOCK:
		return "FORMAT_ASTC_12x10_UNORM_BLOCK"
	case FORMAT_ASTC_12x10_SRGB_BLOCK:
		return "FORMAT_ASTC_12x10_SRGB_BLOCK"
	case FORMAT_ASTC_12x12_UNORM_BLOCK:
		return "FORMAT_ASTC_12x12_UNORM_BLOCK"
	case FORMAT_ASTC_12x12_SRGB_BLOCK:
		return "FORMAT_ASTC_12x12_SRGB_BLOCK"
	case FORMAT_G8B8G8R8_422_UNORM:
		return "FORMAT_G8B8G8R8_422_UNORM"
	case FORMAT_B8G8R8G8_422_UNORM:
		return "FORMAT_B8G8R8G8_422_UNORM"
	case FORMAT_G8_B8_R8_3PLANE_420_UNORM:
		return "FORMAT_G8_B8_R8_3PLANE_420_UNORM"
	case FORMAT_G8_B8R8_2PLANE_420_UNORM:
		return "FORMAT_G8_B8R8_2PLANE_420_UNORM"
	case FORMAT_G8_B8_R8_3PLANE_422_UNORM:
		return "FORMAT_G8_B8_R8_3PLANE_422_UNORM"
	case FORMAT_G8_B8R8_2PLANE_422_UNORM:
		return "FORMAT_G8_B8R8_2PLANE_422_UNORM"
	case FORMAT_G8_B8_R8_3PLANE_444_UNORM:
		return "FORMAT_G8_B8_R8_3PLANE_444_UNORM"
	case FORMAT_R10X6_UNORM_PACK16:
		return "FORMAT_R10X6_UNORM_PACK16"
	case FORMAT_R10X6G10X6_UNORM_2PACK16:
		return "FORMAT_R10X6G10X6_UNORM_2PACK16"
	case FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16:
		return "FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16"
	case FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16:
		return "FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16"
	case FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16:
		return "FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16"
	case FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16:
		return "FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16"
	case FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16:
		return "FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16"
	case FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16:
		return "FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16"
	case FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16:
		return "FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16"
	case FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16:
		return "FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16"
	case FORMAT_R12X4_UNORM_PACK16:
		return "FORMAT_R12X4_UNORM_PACK16"
	case FORMAT_R12X4G12X4_UNORM_2PACK16:
		return "FORMAT_R12X4G12X4_UNORM_2PACK16"
	case FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16:
		return "FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16"
	case FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16:
		return "FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16"
	case FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16:
		return "FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16"
	case FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16:
		return "FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16"
	case FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16:
		return "FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16"
	case FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16:
		return "FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16"
	case FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16:
		return "FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16"
	case FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16:
		return "FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16"
	case FORMAT_G16B16G16R16_422_UNORM:
		return "FORMAT_G16B16G16R16_422_UNORM"
	case FORMAT_B16G16R16G16_422_UNORM:
		return "FORMAT_B16G16R16G16_422_UNORM"
	case FORMAT_G16_B16_R16_3PLANE_420_UNORM:
		return "FORMAT_G16_B16_R16_3PLANE_420_UNORM"
	case FORMAT_G16_B16R16_2PLANE_420_UNORM:
		return "FORMAT_G16_B16R16_2PLANE_420_UNORM"
	case FORMAT_G16_B16_R16_3PLANE_422_UNORM:
		return "FORMAT_G16_B16_R16_3PLANE_422_UNORM"
	case FORMAT_G16_B16R16_2PLANE_422_UNORM:
		return "FORMAT_G16_B16R16_2PLANE_422_UNORM"
	case FORMAT_G16_B16_R16_3PLANE_444_UNORM:
		return "FORMAT_G16_B16_R16_3PLANE_444_UNORM"
	case FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG:
		return "FORMAT_PVRTC1_2BPP_UNORM_BLOCK_IMG"
	case FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG:
		return "FORMAT_PVRTC1_4BPP_UNORM_BLOCK_IMG"
	case FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG:
		return "FORMAT_PVRTC2_2BPP_UNORM_BLOCK_IMG"
	case FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG:
		return "FORMAT_PVRTC2_4BPP_UNORM_BLOCK_IMG"
	case FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG:
		return "FORMAT_PVRTC1_2BPP_SRGB_BLOCK_IMG"
	case FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG:
		return "FORMAT_PVRTC1_4BPP_SRGB_BLOCK_IMG"
	case FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG:
		return "FORMAT_PVRTC2_2BPP_SRGB_BLOCK_IMG"
	case FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG:
		return "FORMAT_PVRTC2_4BPP_SRGB_BLOCK_IMG"
	case FORMAT_ASTC_4x4_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_4x4_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_5x4_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_5x4_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_5x5_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_5x5_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_6x5_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_6x5_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_6x6_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_6x6_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_8x5_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_8x5_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_8x6_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_8x6_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_8x8_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_8x8_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_10x5_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_10x5_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_10x6_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_10x6_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_10x8_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_10x8_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_10x10_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_10x10_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_12x10_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_12x10_SFLOAT_BLOCK_EXT"
	case FORMAT_ASTC_12x12_SFLOAT_BLOCK_EXT:
		return "FORMAT_ASTC_12x12_SFLOAT_BLOCK_EXT"
	case FORMAT_MAX_ENUM:
		return "FORMAT_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ImageType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageType.html
type ImageType int32

const (
	IMAGE_TYPE_1D          ImageType = 0
	IMAGE_TYPE_2D          ImageType = 1
	IMAGE_TYPE_3D          ImageType = 2
	IMAGE_TYPE_BEGIN_RANGE ImageType = IMAGE_TYPE_1D
	IMAGE_TYPE_END_RANGE   ImageType = IMAGE_TYPE_3D
	IMAGE_TYPE_RANGE_SIZE  ImageType = (IMAGE_TYPE_3D - IMAGE_TYPE_1D + 1)
	IMAGE_TYPE_MAX_ENUM    ImageType = 0x7FFFFFFF
)

func (x ImageType) String() string {
	switch x {
	case IMAGE_TYPE_1D:
		return "IMAGE_TYPE_1D"
	case IMAGE_TYPE_2D:
		return "IMAGE_TYPE_2D"
	case IMAGE_TYPE_3D:
		return "IMAGE_TYPE_3D"
	case IMAGE_TYPE_MAX_ENUM:
		return "IMAGE_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ImageTiling -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageTiling.html
type ImageTiling int32

const (
	IMAGE_TILING_OPTIMAL                 ImageTiling = 0
	IMAGE_TILING_LINEAR                  ImageTiling = 1
	IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT ImageTiling = 1000158000
	IMAGE_TILING_BEGIN_RANGE             ImageTiling = IMAGE_TILING_OPTIMAL
	IMAGE_TILING_END_RANGE               ImageTiling = IMAGE_TILING_LINEAR
	IMAGE_TILING_RANGE_SIZE              ImageTiling = (IMAGE_TILING_LINEAR - IMAGE_TILING_OPTIMAL + 1)
	IMAGE_TILING_MAX_ENUM                ImageTiling = 0x7FFFFFFF
)

func (x ImageTiling) String() string {
	switch x {
	case IMAGE_TILING_OPTIMAL:
		return "IMAGE_TILING_OPTIMAL"
	case IMAGE_TILING_LINEAR:
		return "IMAGE_TILING_LINEAR"
	case IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT:
		return "IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT"
	case IMAGE_TILING_MAX_ENUM:
		return "IMAGE_TILING_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PhysicalDeviceType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceType.html
type PhysicalDeviceType int32

const (
	PHYSICAL_DEVICE_TYPE_OTHER          PhysicalDeviceType = 0
	PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU PhysicalDeviceType = 1
	PHYSICAL_DEVICE_TYPE_DISCRETE_GPU   PhysicalDeviceType = 2
	PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU    PhysicalDeviceType = 3
	PHYSICAL_DEVICE_TYPE_CPU            PhysicalDeviceType = 4
	PHYSICAL_DEVICE_TYPE_BEGIN_RANGE    PhysicalDeviceType = PHYSICAL_DEVICE_TYPE_OTHER
	PHYSICAL_DEVICE_TYPE_END_RANGE      PhysicalDeviceType = PHYSICAL_DEVICE_TYPE_CPU
	PHYSICAL_DEVICE_TYPE_RANGE_SIZE     PhysicalDeviceType = (PHYSICAL_DEVICE_TYPE_CPU - PHYSICAL_DEVICE_TYPE_OTHER + 1)
	PHYSICAL_DEVICE_TYPE_MAX_ENUM       PhysicalDeviceType = 0x7FFFFFFF
)

func (x PhysicalDeviceType) String() string {
	switch x {
	case PHYSICAL_DEVICE_TYPE_OTHER:
		return "PHYSICAL_DEVICE_TYPE_OTHER"
	case PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU:
		return "PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU"
	case PHYSICAL_DEVICE_TYPE_DISCRETE_GPU:
		return "PHYSICAL_DEVICE_TYPE_DISCRETE_GPU"
	case PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU:
		return "PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU"
	case PHYSICAL_DEVICE_TYPE_CPU:
		return "PHYSICAL_DEVICE_TYPE_CPU"
	case PHYSICAL_DEVICE_TYPE_MAX_ENUM:
		return "PHYSICAL_DEVICE_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// QueryType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryType.html
type QueryType int32

const (
	QUERY_TYPE_OCCLUSION                                QueryType = 0
	QUERY_TYPE_PIPELINE_STATISTICS                      QueryType = 1
	QUERY_TYPE_TIMESTAMP                                QueryType = 2
	QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT            QueryType = 1000028004
	QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV QueryType = 1000165000
	QUERY_TYPE_PERFORMANCE_QUERY_INTEL                  QueryType = 1000210000
	QUERY_TYPE_BEGIN_RANGE                              QueryType = QUERY_TYPE_OCCLUSION
	QUERY_TYPE_END_RANGE                                QueryType = QUERY_TYPE_TIMESTAMP
	QUERY_TYPE_RANGE_SIZE                               QueryType = (QUERY_TYPE_TIMESTAMP - QUERY_TYPE_OCCLUSION + 1)
	QUERY_TYPE_MAX_ENUM                                 QueryType = 0x7FFFFFFF
)

func (x QueryType) String() string {
	switch x {
	case QUERY_TYPE_OCCLUSION:
		return "QUERY_TYPE_OCCLUSION"
	case QUERY_TYPE_PIPELINE_STATISTICS:
		return "QUERY_TYPE_PIPELINE_STATISTICS"
	case QUERY_TYPE_TIMESTAMP:
		return "QUERY_TYPE_TIMESTAMP"
	case QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT:
		return "QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT"
	case QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV:
		return "QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV"
	case QUERY_TYPE_PERFORMANCE_QUERY_INTEL:
		return "QUERY_TYPE_PERFORMANCE_QUERY_INTEL"
	case QUERY_TYPE_MAX_ENUM:
		return "QUERY_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SharingMode -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSharingMode.html
type SharingMode int32

const (
	SHARING_MODE_EXCLUSIVE   SharingMode = 0
	SHARING_MODE_CONCURRENT  SharingMode = 1
	SHARING_MODE_BEGIN_RANGE SharingMode = SHARING_MODE_EXCLUSIVE
	SHARING_MODE_END_RANGE   SharingMode = SHARING_MODE_CONCURRENT
	SHARING_MODE_RANGE_SIZE  SharingMode = (SHARING_MODE_CONCURRENT - SHARING_MODE_EXCLUSIVE + 1)
	SHARING_MODE_MAX_ENUM    SharingMode = 0x7FFFFFFF
)

func (x SharingMode) String() string {
	switch x {
	case SHARING_MODE_EXCLUSIVE:
		return "SHARING_MODE_EXCLUSIVE"
	case SHARING_MODE_CONCURRENT:
		return "SHARING_MODE_CONCURRENT"
	case SHARING_MODE_MAX_ENUM:
		return "SHARING_MODE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ImageLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageLayout.html
type ImageLayout int32

const (
	IMAGE_LAYOUT_UNDEFINED                                      ImageLayout = 0
	IMAGE_LAYOUT_GENERAL                                        ImageLayout = 1
	IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL                       ImageLayout = 2
	IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL               ImageLayout = 3
	IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL                ImageLayout = 4
	IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL                       ImageLayout = 5
	IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL                           ImageLayout = 6
	IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL                           ImageLayout = 7
	IMAGE_LAYOUT_PREINITIALIZED                                 ImageLayout = 8
	IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL     ImageLayout = 1000117000
	IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL     ImageLayout = 1000117001
	IMAGE_LAYOUT_PRESENT_SRC_KHR                                ImageLayout = 1000001002
	IMAGE_LAYOUT_SHARED_PRESENT_KHR                             ImageLayout = 1000111000
	IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV                        ImageLayout = 1000164003
	IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT               ImageLayout = 1000218000
	IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR ImageLayout = IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL
	IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR ImageLayout = IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL
	IMAGE_LAYOUT_BEGIN_RANGE                                    ImageLayout = IMAGE_LAYOUT_UNDEFINED
	IMAGE_LAYOUT_END_RANGE                                      ImageLayout = IMAGE_LAYOUT_PREINITIALIZED
	IMAGE_LAYOUT_RANGE_SIZE                                     ImageLayout = (IMAGE_LAYOUT_PREINITIALIZED - IMAGE_LAYOUT_UNDEFINED + 1)
	IMAGE_LAYOUT_MAX_ENUM                                       ImageLayout = 0x7FFFFFFF
)

func (x ImageLayout) String() string {
	switch x {
	case IMAGE_LAYOUT_UNDEFINED:
		return "IMAGE_LAYOUT_UNDEFINED"
	case IMAGE_LAYOUT_GENERAL:
		return "IMAGE_LAYOUT_GENERAL"
	case IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL:
		return "IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL"
	case IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL:
		return "IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL"
	case IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL:
		return "IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL"
	case IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL:
		return "IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL"
	case IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL:
		return "IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL"
	case IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL:
		return "IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL"
	case IMAGE_LAYOUT_PREINITIALIZED:
		return "IMAGE_LAYOUT_PREINITIALIZED"
	case IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL:
		return "IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL"
	case IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL:
		return "IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL"
	case IMAGE_LAYOUT_PRESENT_SRC_KHR:
		return "IMAGE_LAYOUT_PRESENT_SRC_KHR"
	case IMAGE_LAYOUT_SHARED_PRESENT_KHR:
		return "IMAGE_LAYOUT_SHARED_PRESENT_KHR"
	case IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV:
		return "IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV"
	case IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT:
		return "IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT"
	case IMAGE_LAYOUT_MAX_ENUM:
		return "IMAGE_LAYOUT_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ImageViewType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageViewType.html
type ImageViewType int32

const (
	IMAGE_VIEW_TYPE_1D          ImageViewType = 0
	IMAGE_VIEW_TYPE_2D          ImageViewType = 1
	IMAGE_VIEW_TYPE_3D          ImageViewType = 2
	IMAGE_VIEW_TYPE_CUBE        ImageViewType = 3
	IMAGE_VIEW_TYPE_1D_ARRAY    ImageViewType = 4
	IMAGE_VIEW_TYPE_2D_ARRAY    ImageViewType = 5
	IMAGE_VIEW_TYPE_CUBE_ARRAY  ImageViewType = 6
	IMAGE_VIEW_TYPE_BEGIN_RANGE ImageViewType = IMAGE_VIEW_TYPE_1D
	IMAGE_VIEW_TYPE_END_RANGE   ImageViewType = IMAGE_VIEW_TYPE_CUBE_ARRAY
	IMAGE_VIEW_TYPE_RANGE_SIZE  ImageViewType = (IMAGE_VIEW_TYPE_CUBE_ARRAY - IMAGE_VIEW_TYPE_1D + 1)
	IMAGE_VIEW_TYPE_MAX_ENUM    ImageViewType = 0x7FFFFFFF
)

func (x ImageViewType) String() string {
	switch x {
	case IMAGE_VIEW_TYPE_1D:
		return "IMAGE_VIEW_TYPE_1D"
	case IMAGE_VIEW_TYPE_2D:
		return "IMAGE_VIEW_TYPE_2D"
	case IMAGE_VIEW_TYPE_3D:
		return "IMAGE_VIEW_TYPE_3D"
	case IMAGE_VIEW_TYPE_CUBE:
		return "IMAGE_VIEW_TYPE_CUBE"
	case IMAGE_VIEW_TYPE_1D_ARRAY:
		return "IMAGE_VIEW_TYPE_1D_ARRAY"
	case IMAGE_VIEW_TYPE_2D_ARRAY:
		return "IMAGE_VIEW_TYPE_2D_ARRAY"
	case IMAGE_VIEW_TYPE_CUBE_ARRAY:
		return "IMAGE_VIEW_TYPE_CUBE_ARRAY"
	case IMAGE_VIEW_TYPE_MAX_ENUM:
		return "IMAGE_VIEW_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ComponentSwizzle -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkComponentSwizzle.html
type ComponentSwizzle int32

const (
	COMPONENT_SWIZZLE_IDENTITY    ComponentSwizzle = 0
	COMPONENT_SWIZZLE_ZERO        ComponentSwizzle = 1
	COMPONENT_SWIZZLE_ONE         ComponentSwizzle = 2
	COMPONENT_SWIZZLE_R           ComponentSwizzle = 3
	COMPONENT_SWIZZLE_G           ComponentSwizzle = 4
	COMPONENT_SWIZZLE_B           ComponentSwizzle = 5
	COMPONENT_SWIZZLE_A           ComponentSwizzle = 6
	COMPONENT_SWIZZLE_BEGIN_RANGE ComponentSwizzle = COMPONENT_SWIZZLE_IDENTITY
	COMPONENT_SWIZZLE_END_RANGE   ComponentSwizzle = COMPONENT_SWIZZLE_A
	COMPONENT_SWIZZLE_RANGE_SIZE  ComponentSwizzle = (COMPONENT_SWIZZLE_A - COMPONENT_SWIZZLE_IDENTITY + 1)
	COMPONENT_SWIZZLE_MAX_ENUM    ComponentSwizzle = 0x7FFFFFFF
)

func (x ComponentSwizzle) String() string {
	switch x {
	case COMPONENT_SWIZZLE_IDENTITY:
		return "COMPONENT_SWIZZLE_IDENTITY"
	case COMPONENT_SWIZZLE_ZERO:
		return "COMPONENT_SWIZZLE_ZERO"
	case COMPONENT_SWIZZLE_ONE:
		return "COMPONENT_SWIZZLE_ONE"
	case COMPONENT_SWIZZLE_R:
		return "COMPONENT_SWIZZLE_R"
	case COMPONENT_SWIZZLE_G:
		return "COMPONENT_SWIZZLE_G"
	case COMPONENT_SWIZZLE_B:
		return "COMPONENT_SWIZZLE_B"
	case COMPONENT_SWIZZLE_A:
		return "COMPONENT_SWIZZLE_A"
	case COMPONENT_SWIZZLE_MAX_ENUM:
		return "COMPONENT_SWIZZLE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// VertexInputRate -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkVertexInputRate.html
type VertexInputRate int32

const (
	VERTEX_INPUT_RATE_VERTEX      VertexInputRate = 0
	VERTEX_INPUT_RATE_INSTANCE    VertexInputRate = 1
	VERTEX_INPUT_RATE_BEGIN_RANGE VertexInputRate = VERTEX_INPUT_RATE_VERTEX
	VERTEX_INPUT_RATE_END_RANGE   VertexInputRate = VERTEX_INPUT_RATE_INSTANCE
	VERTEX_INPUT_RATE_RANGE_SIZE  VertexInputRate = (VERTEX_INPUT_RATE_INSTANCE - VERTEX_INPUT_RATE_VERTEX + 1)
	VERTEX_INPUT_RATE_MAX_ENUM    VertexInputRate = 0x7FFFFFFF
)

func (x VertexInputRate) String() string {
	switch x {
	case VERTEX_INPUT_RATE_VERTEX:
		return "VERTEX_INPUT_RATE_VERTEX"
	case VERTEX_INPUT_RATE_INSTANCE:
		return "VERTEX_INPUT_RATE_INSTANCE"
	case VERTEX_INPUT_RATE_MAX_ENUM:
		return "VERTEX_INPUT_RATE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PrimitiveTopology -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPrimitiveTopology.html
type PrimitiveTopology int32

const (
	PRIMITIVE_TOPOLOGY_POINT_LIST                    PrimitiveTopology = 0
	PRIMITIVE_TOPOLOGY_LINE_LIST                     PrimitiveTopology = 1
	PRIMITIVE_TOPOLOGY_LINE_STRIP                    PrimitiveTopology = 2
	PRIMITIVE_TOPOLOGY_TRIANGLE_LIST                 PrimitiveTopology = 3
	PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP                PrimitiveTopology = 4
	PRIMITIVE_TOPOLOGY_TRIANGLE_FAN                  PrimitiveTopology = 5
	PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY      PrimitiveTopology = 6
	PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY     PrimitiveTopology = 7
	PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY  PrimitiveTopology = 8
	PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY PrimitiveTopology = 9
	PRIMITIVE_TOPOLOGY_PATCH_LIST                    PrimitiveTopology = 10
	PRIMITIVE_TOPOLOGY_BEGIN_RANGE                   PrimitiveTopology = PRIMITIVE_TOPOLOGY_POINT_LIST
	PRIMITIVE_TOPOLOGY_END_RANGE                     PrimitiveTopology = PRIMITIVE_TOPOLOGY_PATCH_LIST
	PRIMITIVE_TOPOLOGY_RANGE_SIZE                    PrimitiveTopology = (PRIMITIVE_TOPOLOGY_PATCH_LIST - PRIMITIVE_TOPOLOGY_POINT_LIST + 1)
	PRIMITIVE_TOPOLOGY_MAX_ENUM                      PrimitiveTopology = 0x7FFFFFFF
)

func (x PrimitiveTopology) String() string {
	switch x {
	case PRIMITIVE_TOPOLOGY_POINT_LIST:
		return "PRIMITIVE_TOPOLOGY_POINT_LIST"
	case PRIMITIVE_TOPOLOGY_LINE_LIST:
		return "PRIMITIVE_TOPOLOGY_LINE_LIST"
	case PRIMITIVE_TOPOLOGY_LINE_STRIP:
		return "PRIMITIVE_TOPOLOGY_LINE_STRIP"
	case PRIMITIVE_TOPOLOGY_TRIANGLE_LIST:
		return "PRIMITIVE_TOPOLOGY_TRIANGLE_LIST"
	case PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP:
		return "PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP"
	case PRIMITIVE_TOPOLOGY_TRIANGLE_FAN:
		return "PRIMITIVE_TOPOLOGY_TRIANGLE_FAN"
	case PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY:
		return "PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY"
	case PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY:
		return "PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY"
	case PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY:
		return "PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY"
	case PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY:
		return "PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY"
	case PRIMITIVE_TOPOLOGY_PATCH_LIST:
		return "PRIMITIVE_TOPOLOGY_PATCH_LIST"
	case PRIMITIVE_TOPOLOGY_MAX_ENUM:
		return "PRIMITIVE_TOPOLOGY_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PolygonMode -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPolygonMode.html
type PolygonMode int32

const (
	POLYGON_MODE_FILL              PolygonMode = 0
	POLYGON_MODE_LINE              PolygonMode = 1
	POLYGON_MODE_POINT             PolygonMode = 2
	POLYGON_MODE_FILL_RECTANGLE_NV PolygonMode = 1000153000
	POLYGON_MODE_BEGIN_RANGE       PolygonMode = POLYGON_MODE_FILL
	POLYGON_MODE_END_RANGE         PolygonMode = POLYGON_MODE_POINT
	POLYGON_MODE_RANGE_SIZE        PolygonMode = (POLYGON_MODE_POINT - POLYGON_MODE_FILL + 1)
	POLYGON_MODE_MAX_ENUM          PolygonMode = 0x7FFFFFFF
)

func (x PolygonMode) String() string {
	switch x {
	case POLYGON_MODE_FILL:
		return "POLYGON_MODE_FILL"
	case POLYGON_MODE_LINE:
		return "POLYGON_MODE_LINE"
	case POLYGON_MODE_POINT:
		return "POLYGON_MODE_POINT"
	case POLYGON_MODE_FILL_RECTANGLE_NV:
		return "POLYGON_MODE_FILL_RECTANGLE_NV"
	case POLYGON_MODE_MAX_ENUM:
		return "POLYGON_MODE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// FrontFace -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFrontFace.html
type FrontFace int32

const (
	FRONT_FACE_COUNTER_CLOCKWISE FrontFace = 0
	FRONT_FACE_CLOCKWISE         FrontFace = 1
	FRONT_FACE_BEGIN_RANGE       FrontFace = FRONT_FACE_COUNTER_CLOCKWISE
	FRONT_FACE_END_RANGE         FrontFace = FRONT_FACE_CLOCKWISE
	FRONT_FACE_RANGE_SIZE        FrontFace = (FRONT_FACE_CLOCKWISE - FRONT_FACE_COUNTER_CLOCKWISE + 1)
	FRONT_FACE_MAX_ENUM          FrontFace = 0x7FFFFFFF
)

func (x FrontFace) String() string {
	switch x {
	case FRONT_FACE_COUNTER_CLOCKWISE:
		return "FRONT_FACE_COUNTER_CLOCKWISE"
	case FRONT_FACE_CLOCKWISE:
		return "FRONT_FACE_CLOCKWISE"
	case FRONT_FACE_MAX_ENUM:
		return "FRONT_FACE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// CompareOp -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCompareOp.html
type CompareOp int32

const (
	COMPARE_OP_NEVER            CompareOp = 0
	COMPARE_OP_LESS             CompareOp = 1
	COMPARE_OP_EQUAL            CompareOp = 2
	COMPARE_OP_LESS_OR_EQUAL    CompareOp = 3
	COMPARE_OP_GREATER          CompareOp = 4
	COMPARE_OP_NOT_EQUAL        CompareOp = 5
	COMPARE_OP_GREATER_OR_EQUAL CompareOp = 6
	COMPARE_OP_ALWAYS           CompareOp = 7
	COMPARE_OP_BEGIN_RANGE      CompareOp = COMPARE_OP_NEVER
	COMPARE_OP_END_RANGE        CompareOp = COMPARE_OP_ALWAYS
	COMPARE_OP_RANGE_SIZE       CompareOp = (COMPARE_OP_ALWAYS - COMPARE_OP_NEVER + 1)
	COMPARE_OP_MAX_ENUM         CompareOp = 0x7FFFFFFF
)

func (x CompareOp) String() string {
	switch x {
	case COMPARE_OP_NEVER:
		return "COMPARE_OP_NEVER"
	case COMPARE_OP_LESS:
		return "COMPARE_OP_LESS"
	case COMPARE_OP_EQUAL:
		return "COMPARE_OP_EQUAL"
	case COMPARE_OP_LESS_OR_EQUAL:
		return "COMPARE_OP_LESS_OR_EQUAL"
	case COMPARE_OP_GREATER:
		return "COMPARE_OP_GREATER"
	case COMPARE_OP_NOT_EQUAL:
		return "COMPARE_OP_NOT_EQUAL"
	case COMPARE_OP_GREATER_OR_EQUAL:
		return "COMPARE_OP_GREATER_OR_EQUAL"
	case COMPARE_OP_ALWAYS:
		return "COMPARE_OP_ALWAYS"
	case COMPARE_OP_MAX_ENUM:
		return "COMPARE_OP_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// StencilOp -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkStencilOp.html
type StencilOp int32

const (
	STENCIL_OP_KEEP                StencilOp = 0
	STENCIL_OP_ZERO                StencilOp = 1
	STENCIL_OP_REPLACE             StencilOp = 2
	STENCIL_OP_INCREMENT_AND_CLAMP StencilOp = 3
	STENCIL_OP_DECREMENT_AND_CLAMP StencilOp = 4
	STENCIL_OP_INVERT              StencilOp = 5
	STENCIL_OP_INCREMENT_AND_WRAP  StencilOp = 6
	STENCIL_OP_DECREMENT_AND_WRAP  StencilOp = 7
	STENCIL_OP_BEGIN_RANGE         StencilOp = STENCIL_OP_KEEP
	STENCIL_OP_END_RANGE           StencilOp = STENCIL_OP_DECREMENT_AND_WRAP
	STENCIL_OP_RANGE_SIZE          StencilOp = (STENCIL_OP_DECREMENT_AND_WRAP - STENCIL_OP_KEEP + 1)
	STENCIL_OP_MAX_ENUM            StencilOp = 0x7FFFFFFF
)

func (x StencilOp) String() string {
	switch x {
	case STENCIL_OP_KEEP:
		return "STENCIL_OP_KEEP"
	case STENCIL_OP_ZERO:
		return "STENCIL_OP_ZERO"
	case STENCIL_OP_REPLACE:
		return "STENCIL_OP_REPLACE"
	case STENCIL_OP_INCREMENT_AND_CLAMP:
		return "STENCIL_OP_INCREMENT_AND_CLAMP"
	case STENCIL_OP_DECREMENT_AND_CLAMP:
		return "STENCIL_OP_DECREMENT_AND_CLAMP"
	case STENCIL_OP_INVERT:
		return "STENCIL_OP_INVERT"
	case STENCIL_OP_INCREMENT_AND_WRAP:
		return "STENCIL_OP_INCREMENT_AND_WRAP"
	case STENCIL_OP_DECREMENT_AND_WRAP:
		return "STENCIL_OP_DECREMENT_AND_WRAP"
	case STENCIL_OP_MAX_ENUM:
		return "STENCIL_OP_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// LogicOp -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkLogicOp.html
type LogicOp int32

const (
	LOGIC_OP_CLEAR         LogicOp = 0
	LOGIC_OP_AND           LogicOp = 1
	LOGIC_OP_AND_REVERSE   LogicOp = 2
	LOGIC_OP_COPY          LogicOp = 3
	LOGIC_OP_AND_INVERTED  LogicOp = 4
	LOGIC_OP_NO_OP         LogicOp = 5
	LOGIC_OP_XOR           LogicOp = 6
	LOGIC_OP_OR            LogicOp = 7
	LOGIC_OP_NOR           LogicOp = 8
	LOGIC_OP_EQUIVALENT    LogicOp = 9
	LOGIC_OP_INVERT        LogicOp = 10
	LOGIC_OP_OR_REVERSE    LogicOp = 11
	LOGIC_OP_COPY_INVERTED LogicOp = 12
	LOGIC_OP_OR_INVERTED   LogicOp = 13
	LOGIC_OP_NAND          LogicOp = 14
	LOGIC_OP_SET           LogicOp = 15
	LOGIC_OP_BEGIN_RANGE   LogicOp = LOGIC_OP_CLEAR
	LOGIC_OP_END_RANGE     LogicOp = LOGIC_OP_SET
	LOGIC_OP_RANGE_SIZE    LogicOp = (LOGIC_OP_SET - LOGIC_OP_CLEAR + 1)
	LOGIC_OP_MAX_ENUM      LogicOp = 0x7FFFFFFF
)

func (x LogicOp) String() string {
	switch x {
	case LOGIC_OP_CLEAR:
		return "LOGIC_OP_CLEAR"
	case LOGIC_OP_AND:
		return "LOGIC_OP_AND"
	case LOGIC_OP_AND_REVERSE:
		return "LOGIC_OP_AND_REVERSE"
	case LOGIC_OP_COPY:
		return "LOGIC_OP_COPY"
	case LOGIC_OP_AND_INVERTED:
		return "LOGIC_OP_AND_INVERTED"
	case LOGIC_OP_NO_OP:
		return "LOGIC_OP_NO_OP"
	case LOGIC_OP_XOR:
		return "LOGIC_OP_XOR"
	case LOGIC_OP_OR:
		return "LOGIC_OP_OR"
	case LOGIC_OP_NOR:
		return "LOGIC_OP_NOR"
	case LOGIC_OP_EQUIVALENT:
		return "LOGIC_OP_EQUIVALENT"
	case LOGIC_OP_INVERT:
		return "LOGIC_OP_INVERT"
	case LOGIC_OP_OR_REVERSE:
		return "LOGIC_OP_OR_REVERSE"
	case LOGIC_OP_COPY_INVERTED:
		return "LOGIC_OP_COPY_INVERTED"
	case LOGIC_OP_OR_INVERTED:
		return "LOGIC_OP_OR_INVERTED"
	case LOGIC_OP_NAND:
		return "LOGIC_OP_NAND"
	case LOGIC_OP_SET:
		return "LOGIC_OP_SET"
	case LOGIC_OP_MAX_ENUM:
		return "LOGIC_OP_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// BlendFactor -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBlendFactor.html
type BlendFactor int32

const (
	BLEND_FACTOR_ZERO                     BlendFactor = 0
	BLEND_FACTOR_ONE                      BlendFactor = 1
	BLEND_FACTOR_SRC_COLOR                BlendFactor = 2
	BLEND_FACTOR_ONE_MINUS_SRC_COLOR      BlendFactor = 3
	BLEND_FACTOR_DST_COLOR                BlendFactor = 4
	BLEND_FACTOR_ONE_MINUS_DST_COLOR      BlendFactor = 5
	BLEND_FACTOR_SRC_ALPHA                BlendFactor = 6
	BLEND_FACTOR_ONE_MINUS_SRC_ALPHA      BlendFactor = 7
	BLEND_FACTOR_DST_ALPHA                BlendFactor = 8
	BLEND_FACTOR_ONE_MINUS_DST_ALPHA      BlendFactor = 9
	BLEND_FACTOR_CONSTANT_COLOR           BlendFactor = 10
	BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR BlendFactor = 11
	BLEND_FACTOR_CONSTANT_ALPHA           BlendFactor = 12
	BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA BlendFactor = 13
	BLEND_FACTOR_SRC_ALPHA_SATURATE       BlendFactor = 14
	BLEND_FACTOR_SRC1_COLOR               BlendFactor = 15
	BLEND_FACTOR_ONE_MINUS_SRC1_COLOR     BlendFactor = 16
	BLEND_FACTOR_SRC1_ALPHA               BlendFactor = 17
	BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA     BlendFactor = 18
	BLEND_FACTOR_BEGIN_RANGE              BlendFactor = BLEND_FACTOR_ZERO
	BLEND_FACTOR_END_RANGE                BlendFactor = BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA
	BLEND_FACTOR_RANGE_SIZE               BlendFactor = (BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA - BLEND_FACTOR_ZERO + 1)
	BLEND_FACTOR_MAX_ENUM                 BlendFactor = 0x7FFFFFFF
)

func (x BlendFactor) String() string {
	switch x {
	case BLEND_FACTOR_ZERO:
		return "BLEND_FACTOR_ZERO"
	case BLEND_FACTOR_ONE:
		return "BLEND_FACTOR_ONE"
	case BLEND_FACTOR_SRC_COLOR:
		return "BLEND_FACTOR_SRC_COLOR"
	case BLEND_FACTOR_ONE_MINUS_SRC_COLOR:
		return "BLEND_FACTOR_ONE_MINUS_SRC_COLOR"
	case BLEND_FACTOR_DST_COLOR:
		return "BLEND_FACTOR_DST_COLOR"
	case BLEND_FACTOR_ONE_MINUS_DST_COLOR:
		return "BLEND_FACTOR_ONE_MINUS_DST_COLOR"
	case BLEND_FACTOR_SRC_ALPHA:
		return "BLEND_FACTOR_SRC_ALPHA"
	case BLEND_FACTOR_ONE_MINUS_SRC_ALPHA:
		return "BLEND_FACTOR_ONE_MINUS_SRC_ALPHA"
	case BLEND_FACTOR_DST_ALPHA:
		return "BLEND_FACTOR_DST_ALPHA"
	case BLEND_FACTOR_ONE_MINUS_DST_ALPHA:
		return "BLEND_FACTOR_ONE_MINUS_DST_ALPHA"
	case BLEND_FACTOR_CONSTANT_COLOR:
		return "BLEND_FACTOR_CONSTANT_COLOR"
	case BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR:
		return "BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR"
	case BLEND_FACTOR_CONSTANT_ALPHA:
		return "BLEND_FACTOR_CONSTANT_ALPHA"
	case BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA:
		return "BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA"
	case BLEND_FACTOR_SRC_ALPHA_SATURATE:
		return "BLEND_FACTOR_SRC_ALPHA_SATURATE"
	case BLEND_FACTOR_SRC1_COLOR:
		return "BLEND_FACTOR_SRC1_COLOR"
	case BLEND_FACTOR_ONE_MINUS_SRC1_COLOR:
		return "BLEND_FACTOR_ONE_MINUS_SRC1_COLOR"
	case BLEND_FACTOR_SRC1_ALPHA:
		return "BLEND_FACTOR_SRC1_ALPHA"
	case BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA:
		return "BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA"
	case BLEND_FACTOR_MAX_ENUM:
		return "BLEND_FACTOR_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// BlendOp -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBlendOp.html
type BlendOp int32

const (
	BLEND_OP_ADD                    BlendOp = 0
	BLEND_OP_SUBTRACT               BlendOp = 1
	BLEND_OP_REVERSE_SUBTRACT       BlendOp = 2
	BLEND_OP_MIN                    BlendOp = 3
	BLEND_OP_MAX                    BlendOp = 4
	BLEND_OP_ZERO_EXT               BlendOp = 1000148000
	BLEND_OP_SRC_EXT                BlendOp = 1000148001
	BLEND_OP_DST_EXT                BlendOp = 1000148002
	BLEND_OP_SRC_OVER_EXT           BlendOp = 1000148003
	BLEND_OP_DST_OVER_EXT           BlendOp = 1000148004
	BLEND_OP_SRC_IN_EXT             BlendOp = 1000148005
	BLEND_OP_DST_IN_EXT             BlendOp = 1000148006
	BLEND_OP_SRC_OUT_EXT            BlendOp = 1000148007
	BLEND_OP_DST_OUT_EXT            BlendOp = 1000148008
	BLEND_OP_SRC_ATOP_EXT           BlendOp = 1000148009
	BLEND_OP_DST_ATOP_EXT           BlendOp = 1000148010
	BLEND_OP_XOR_EXT                BlendOp = 1000148011
	BLEND_OP_MULTIPLY_EXT           BlendOp = 1000148012
	BLEND_OP_SCREEN_EXT             BlendOp = 1000148013
	BLEND_OP_OVERLAY_EXT            BlendOp = 1000148014
	BLEND_OP_DARKEN_EXT             BlendOp = 1000148015
	BLEND_OP_LIGHTEN_EXT            BlendOp = 1000148016
	BLEND_OP_COLORDODGE_EXT         BlendOp = 1000148017
	BLEND_OP_COLORBURN_EXT          BlendOp = 1000148018
	BLEND_OP_HARDLIGHT_EXT          BlendOp = 1000148019
	BLEND_OP_SOFTLIGHT_EXT          BlendOp = 1000148020
	BLEND_OP_DIFFERENCE_EXT         BlendOp = 1000148021
	BLEND_OP_EXCLUSION_EXT          BlendOp = 1000148022
	BLEND_OP_INVERT_EXT             BlendOp = 1000148023
	BLEND_OP_INVERT_RGB_EXT         BlendOp = 1000148024
	BLEND_OP_LINEARDODGE_EXT        BlendOp = 1000148025
	BLEND_OP_LINEARBURN_EXT         BlendOp = 1000148026
	BLEND_OP_VIVIDLIGHT_EXT         BlendOp = 1000148027
	BLEND_OP_LINEARLIGHT_EXT        BlendOp = 1000148028
	BLEND_OP_PINLIGHT_EXT           BlendOp = 1000148029
	BLEND_OP_HARDMIX_EXT            BlendOp = 1000148030
	BLEND_OP_HSL_HUE_EXT            BlendOp = 1000148031
	BLEND_OP_HSL_SATURATION_EXT     BlendOp = 1000148032
	BLEND_OP_HSL_COLOR_EXT          BlendOp = 1000148033
	BLEND_OP_HSL_LUMINOSITY_EXT     BlendOp = 1000148034
	BLEND_OP_PLUS_EXT               BlendOp = 1000148035
	BLEND_OP_PLUS_CLAMPED_EXT       BlendOp = 1000148036
	BLEND_OP_PLUS_CLAMPED_ALPHA_EXT BlendOp = 1000148037
	BLEND_OP_PLUS_DARKER_EXT        BlendOp = 1000148038
	BLEND_OP_MINUS_EXT              BlendOp = 1000148039
	BLEND_OP_MINUS_CLAMPED_EXT      BlendOp = 1000148040
	BLEND_OP_CONTRAST_EXT           BlendOp = 1000148041
	BLEND_OP_INVERT_OVG_EXT         BlendOp = 1000148042
	BLEND_OP_RED_EXT                BlendOp = 1000148043
	BLEND_OP_GREEN_EXT              BlendOp = 1000148044
	BLEND_OP_BLUE_EXT               BlendOp = 1000148045
	BLEND_OP_BEGIN_RANGE            BlendOp = BLEND_OP_ADD
	BLEND_OP_END_RANGE              BlendOp = BLEND_OP_MAX
	BLEND_OP_RANGE_SIZE             BlendOp = (BLEND_OP_MAX - BLEND_OP_ADD + 1)
	BLEND_OP_MAX_ENUM               BlendOp = 0x7FFFFFFF
)

func (x BlendOp) String() string {
	switch x {
	case BLEND_OP_ADD:
		return "BLEND_OP_ADD"
	case BLEND_OP_SUBTRACT:
		return "BLEND_OP_SUBTRACT"
	case BLEND_OP_REVERSE_SUBTRACT:
		return "BLEND_OP_REVERSE_SUBTRACT"
	case BLEND_OP_MIN:
		return "BLEND_OP_MIN"
	case BLEND_OP_MAX:
		return "BLEND_OP_MAX"
	case BLEND_OP_ZERO_EXT:
		return "BLEND_OP_ZERO_EXT"
	case BLEND_OP_SRC_EXT:
		return "BLEND_OP_SRC_EXT"
	case BLEND_OP_DST_EXT:
		return "BLEND_OP_DST_EXT"
	case BLEND_OP_SRC_OVER_EXT:
		return "BLEND_OP_SRC_OVER_EXT"
	case BLEND_OP_DST_OVER_EXT:
		return "BLEND_OP_DST_OVER_EXT"
	case BLEND_OP_SRC_IN_EXT:
		return "BLEND_OP_SRC_IN_EXT"
	case BLEND_OP_DST_IN_EXT:
		return "BLEND_OP_DST_IN_EXT"
	case BLEND_OP_SRC_OUT_EXT:
		return "BLEND_OP_SRC_OUT_EXT"
	case BLEND_OP_DST_OUT_EXT:
		return "BLEND_OP_DST_OUT_EXT"
	case BLEND_OP_SRC_ATOP_EXT:
		return "BLEND_OP_SRC_ATOP_EXT"
	case BLEND_OP_DST_ATOP_EXT:
		return "BLEND_OP_DST_ATOP_EXT"
	case BLEND_OP_XOR_EXT:
		return "BLEND_OP_XOR_EXT"
	case BLEND_OP_MULTIPLY_EXT:
		return "BLEND_OP_MULTIPLY_EXT"
	case BLEND_OP_SCREEN_EXT:
		return "BLEND_OP_SCREEN_EXT"
	case BLEND_OP_OVERLAY_EXT:
		return "BLEND_OP_OVERLAY_EXT"
	case BLEND_OP_DARKEN_EXT:
		return "BLEND_OP_DARKEN_EXT"
	case BLEND_OP_LIGHTEN_EXT:
		return "BLEND_OP_LIGHTEN_EXT"
	case BLEND_OP_COLORDODGE_EXT:
		return "BLEND_OP_COLORDODGE_EXT"
	case BLEND_OP_COLORBURN_EXT:
		return "BLEND_OP_COLORBURN_EXT"
	case BLEND_OP_HARDLIGHT_EXT:
		return "BLEND_OP_HARDLIGHT_EXT"
	case BLEND_OP_SOFTLIGHT_EXT:
		return "BLEND_OP_SOFTLIGHT_EXT"
	case BLEND_OP_DIFFERENCE_EXT:
		return "BLEND_OP_DIFFERENCE_EXT"
	case BLEND_OP_EXCLUSION_EXT:
		return "BLEND_OP_EXCLUSION_EXT"
	case BLEND_OP_INVERT_EXT:
		return "BLEND_OP_INVERT_EXT"
	case BLEND_OP_INVERT_RGB_EXT:
		return "BLEND_OP_INVERT_RGB_EXT"
	case BLEND_OP_LINEARDODGE_EXT:
		return "BLEND_OP_LINEARDODGE_EXT"
	case BLEND_OP_LINEARBURN_EXT:
		return "BLEND_OP_LINEARBURN_EXT"
	case BLEND_OP_VIVIDLIGHT_EXT:
		return "BLEND_OP_VIVIDLIGHT_EXT"
	case BLEND_OP_LINEARLIGHT_EXT:
		return "BLEND_OP_LINEARLIGHT_EXT"
	case BLEND_OP_PINLIGHT_EXT:
		return "BLEND_OP_PINLIGHT_EXT"
	case BLEND_OP_HARDMIX_EXT:
		return "BLEND_OP_HARDMIX_EXT"
	case BLEND_OP_HSL_HUE_EXT:
		return "BLEND_OP_HSL_HUE_EXT"
	case BLEND_OP_HSL_SATURATION_EXT:
		return "BLEND_OP_HSL_SATURATION_EXT"
	case BLEND_OP_HSL_COLOR_EXT:
		return "BLEND_OP_HSL_COLOR_EXT"
	case BLEND_OP_HSL_LUMINOSITY_EXT:
		return "BLEND_OP_HSL_LUMINOSITY_EXT"
	case BLEND_OP_PLUS_EXT:
		return "BLEND_OP_PLUS_EXT"
	case BLEND_OP_PLUS_CLAMPED_EXT:
		return "BLEND_OP_PLUS_CLAMPED_EXT"
	case BLEND_OP_PLUS_CLAMPED_ALPHA_EXT:
		return "BLEND_OP_PLUS_CLAMPED_ALPHA_EXT"
	case BLEND_OP_PLUS_DARKER_EXT:
		return "BLEND_OP_PLUS_DARKER_EXT"
	case BLEND_OP_MINUS_EXT:
		return "BLEND_OP_MINUS_EXT"
	case BLEND_OP_MINUS_CLAMPED_EXT:
		return "BLEND_OP_MINUS_CLAMPED_EXT"
	case BLEND_OP_CONTRAST_EXT:
		return "BLEND_OP_CONTRAST_EXT"
	case BLEND_OP_INVERT_OVG_EXT:
		return "BLEND_OP_INVERT_OVG_EXT"
	case BLEND_OP_RED_EXT:
		return "BLEND_OP_RED_EXT"
	case BLEND_OP_GREEN_EXT:
		return "BLEND_OP_GREEN_EXT"
	case BLEND_OP_BLUE_EXT:
		return "BLEND_OP_BLUE_EXT"
	case BLEND_OP_MAX_ENUM:
		return "BLEND_OP_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DynamicState -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDynamicState.html
type DynamicState int32

const (
	DYNAMIC_STATE_VIEWPORT                         DynamicState = 0
	DYNAMIC_STATE_SCISSOR                          DynamicState = 1
	DYNAMIC_STATE_LINE_WIDTH                       DynamicState = 2
	DYNAMIC_STATE_DEPTH_BIAS                       DynamicState = 3
	DYNAMIC_STATE_BLEND_CONSTANTS                  DynamicState = 4
	DYNAMIC_STATE_DEPTH_BOUNDS                     DynamicState = 5
	DYNAMIC_STATE_STENCIL_COMPARE_MASK             DynamicState = 6
	DYNAMIC_STATE_STENCIL_WRITE_MASK               DynamicState = 7
	DYNAMIC_STATE_STENCIL_REFERENCE                DynamicState = 8
	DYNAMIC_STATE_VIEWPORT_W_SCALING_NV            DynamicState = 1000087000
	DYNAMIC_STATE_DISCARD_RECTANGLE_EXT            DynamicState = 1000099000
	DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT             DynamicState = 1000143000
	DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV DynamicState = 1000164004
	DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV  DynamicState = 1000164006
	DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV             DynamicState = 1000205001
	DYNAMIC_STATE_LINE_STIPPLE_EXT                 DynamicState = 1000259000
	DYNAMIC_STATE_BEGIN_RANGE                      DynamicState = DYNAMIC_STATE_VIEWPORT
	DYNAMIC_STATE_END_RANGE                        DynamicState = DYNAMIC_STATE_STENCIL_REFERENCE
	DYNAMIC_STATE_RANGE_SIZE                       DynamicState = (DYNAMIC_STATE_STENCIL_REFERENCE - DYNAMIC_STATE_VIEWPORT + 1)
	DYNAMIC_STATE_MAX_ENUM                         DynamicState = 0x7FFFFFFF
)

func (x DynamicState) String() string {
	switch x {
	case DYNAMIC_STATE_VIEWPORT:
		return "DYNAMIC_STATE_VIEWPORT"
	case DYNAMIC_STATE_SCISSOR:
		return "DYNAMIC_STATE_SCISSOR"
	case DYNAMIC_STATE_LINE_WIDTH:
		return "DYNAMIC_STATE_LINE_WIDTH"
	case DYNAMIC_STATE_DEPTH_BIAS:
		return "DYNAMIC_STATE_DEPTH_BIAS"
	case DYNAMIC_STATE_BLEND_CONSTANTS:
		return "DYNAMIC_STATE_BLEND_CONSTANTS"
	case DYNAMIC_STATE_DEPTH_BOUNDS:
		return "DYNAMIC_STATE_DEPTH_BOUNDS"
	case DYNAMIC_STATE_STENCIL_COMPARE_MASK:
		return "DYNAMIC_STATE_STENCIL_COMPARE_MASK"
	case DYNAMIC_STATE_STENCIL_WRITE_MASK:
		return "DYNAMIC_STATE_STENCIL_WRITE_MASK"
	case DYNAMIC_STATE_STENCIL_REFERENCE:
		return "DYNAMIC_STATE_STENCIL_REFERENCE"
	case DYNAMIC_STATE_VIEWPORT_W_SCALING_NV:
		return "DYNAMIC_STATE_VIEWPORT_W_SCALING_NV"
	case DYNAMIC_STATE_DISCARD_RECTANGLE_EXT:
		return "DYNAMIC_STATE_DISCARD_RECTANGLE_EXT"
	case DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT:
		return "DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT"
	case DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV:
		return "DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV"
	case DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV:
		return "DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV"
	case DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV:
		return "DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV"
	case DYNAMIC_STATE_LINE_STIPPLE_EXT:
		return "DYNAMIC_STATE_LINE_STIPPLE_EXT"
	case DYNAMIC_STATE_MAX_ENUM:
		return "DYNAMIC_STATE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// Filter -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFilter.html
type Filter int32

const (
	FILTER_NEAREST     Filter = 0
	FILTER_LINEAR      Filter = 1
	FILTER_CUBIC_IMG   Filter = 1000015000
	FILTER_CUBIC_EXT   Filter = FILTER_CUBIC_IMG
	FILTER_BEGIN_RANGE Filter = FILTER_NEAREST
	FILTER_END_RANGE   Filter = FILTER_LINEAR
	FILTER_RANGE_SIZE  Filter = (FILTER_LINEAR - FILTER_NEAREST + 1)
	FILTER_MAX_ENUM    Filter = 0x7FFFFFFF
)

func (x Filter) String() string {
	switch x {
	case FILTER_NEAREST:
		return "FILTER_NEAREST"
	case FILTER_LINEAR:
		return "FILTER_LINEAR"
	case FILTER_CUBIC_IMG:
		return "FILTER_CUBIC_IMG"
	case FILTER_MAX_ENUM:
		return "FILTER_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SamplerMipmapMode -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerMipmapMode.html
type SamplerMipmapMode int32

const (
	SAMPLER_MIPMAP_MODE_NEAREST     SamplerMipmapMode = 0
	SAMPLER_MIPMAP_MODE_LINEAR      SamplerMipmapMode = 1
	SAMPLER_MIPMAP_MODE_BEGIN_RANGE SamplerMipmapMode = SAMPLER_MIPMAP_MODE_NEAREST
	SAMPLER_MIPMAP_MODE_END_RANGE   SamplerMipmapMode = SAMPLER_MIPMAP_MODE_LINEAR
	SAMPLER_MIPMAP_MODE_RANGE_SIZE  SamplerMipmapMode = (SAMPLER_MIPMAP_MODE_LINEAR - SAMPLER_MIPMAP_MODE_NEAREST + 1)
	SAMPLER_MIPMAP_MODE_MAX_ENUM    SamplerMipmapMode = 0x7FFFFFFF
)

func (x SamplerMipmapMode) String() string {
	switch x {
	case SAMPLER_MIPMAP_MODE_NEAREST:
		return "SAMPLER_MIPMAP_MODE_NEAREST"
	case SAMPLER_MIPMAP_MODE_LINEAR:
		return "SAMPLER_MIPMAP_MODE_LINEAR"
	case SAMPLER_MIPMAP_MODE_MAX_ENUM:
		return "SAMPLER_MIPMAP_MODE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SamplerAddressMode -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerAddressMode.html
type SamplerAddressMode int32

const (
	SAMPLER_ADDRESS_MODE_REPEAT                   SamplerAddressMode = 0
	SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT          SamplerAddressMode = 1
	SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE            SamplerAddressMode = 2
	SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER          SamplerAddressMode = 3
	SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE     SamplerAddressMode = 4
	SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR SamplerAddressMode = SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE
	SAMPLER_ADDRESS_MODE_BEGIN_RANGE              SamplerAddressMode = SAMPLER_ADDRESS_MODE_REPEAT
	SAMPLER_ADDRESS_MODE_END_RANGE                SamplerAddressMode = SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER
	SAMPLER_ADDRESS_MODE_RANGE_SIZE               SamplerAddressMode = (SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER - SAMPLER_ADDRESS_MODE_REPEAT + 1)
	SAMPLER_ADDRESS_MODE_MAX_ENUM                 SamplerAddressMode = 0x7FFFFFFF
)

func (x SamplerAddressMode) String() string {
	switch x {
	case SAMPLER_ADDRESS_MODE_REPEAT:
		return "SAMPLER_ADDRESS_MODE_REPEAT"
	case SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT:
		return "SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT"
	case SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE:
		return "SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE"
	case SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER:
		return "SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER"
	case SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE:
		return "SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE"
	case SAMPLER_ADDRESS_MODE_MAX_ENUM:
		return "SAMPLER_ADDRESS_MODE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// BorderColor -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBorderColor.html
type BorderColor int32

const (
	BORDER_COLOR_FLOAT_TRANSPARENT_BLACK BorderColor = 0
	BORDER_COLOR_INT_TRANSPARENT_BLACK   BorderColor = 1
	BORDER_COLOR_FLOAT_OPAQUE_BLACK      BorderColor = 2
	BORDER_COLOR_INT_OPAQUE_BLACK        BorderColor = 3
	BORDER_COLOR_FLOAT_OPAQUE_WHITE      BorderColor = 4
	BORDER_COLOR_INT_OPAQUE_WHITE        BorderColor = 5
	BORDER_COLOR_BEGIN_RANGE             BorderColor = BORDER_COLOR_FLOAT_TRANSPARENT_BLACK
	BORDER_COLOR_END_RANGE               BorderColor = BORDER_COLOR_INT_OPAQUE_WHITE
	BORDER_COLOR_RANGE_SIZE              BorderColor = (BORDER_COLOR_INT_OPAQUE_WHITE - BORDER_COLOR_FLOAT_TRANSPARENT_BLACK + 1)
	BORDER_COLOR_MAX_ENUM                BorderColor = 0x7FFFFFFF
)

func (x BorderColor) String() string {
	switch x {
	case BORDER_COLOR_FLOAT_TRANSPARENT_BLACK:
		return "BORDER_COLOR_FLOAT_TRANSPARENT_BLACK"
	case BORDER_COLOR_INT_TRANSPARENT_BLACK:
		return "BORDER_COLOR_INT_TRANSPARENT_BLACK"
	case BORDER_COLOR_FLOAT_OPAQUE_BLACK:
		return "BORDER_COLOR_FLOAT_OPAQUE_BLACK"
	case BORDER_COLOR_INT_OPAQUE_BLACK:
		return "BORDER_COLOR_INT_OPAQUE_BLACK"
	case BORDER_COLOR_FLOAT_OPAQUE_WHITE:
		return "BORDER_COLOR_FLOAT_OPAQUE_WHITE"
	case BORDER_COLOR_INT_OPAQUE_WHITE:
		return "BORDER_COLOR_INT_OPAQUE_WHITE"
	case BORDER_COLOR_MAX_ENUM:
		return "BORDER_COLOR_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DescriptorType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorType.html
type DescriptorType int32

const (
	DESCRIPTOR_TYPE_SAMPLER                   DescriptorType = 0
	DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER    DescriptorType = 1
	DESCRIPTOR_TYPE_SAMPLED_IMAGE             DescriptorType = 2
	DESCRIPTOR_TYPE_STORAGE_IMAGE             DescriptorType = 3
	DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER      DescriptorType = 4
	DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER      DescriptorType = 5
	DESCRIPTOR_TYPE_UNIFORM_BUFFER            DescriptorType = 6
	DESCRIPTOR_TYPE_STORAGE_BUFFER            DescriptorType = 7
	DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC    DescriptorType = 8
	DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC    DescriptorType = 9
	DESCRIPTOR_TYPE_INPUT_ATTACHMENT          DescriptorType = 10
	DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT  DescriptorType = 1000138000
	DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV DescriptorType = 1000165000
	DESCRIPTOR_TYPE_BEGIN_RANGE               DescriptorType = DESCRIPTOR_TYPE_SAMPLER
	DESCRIPTOR_TYPE_END_RANGE                 DescriptorType = DESCRIPTOR_TYPE_INPUT_ATTACHMENT
	DESCRIPTOR_TYPE_RANGE_SIZE                DescriptorType = (DESCRIPTOR_TYPE_INPUT_ATTACHMENT - DESCRIPTOR_TYPE_SAMPLER + 1)
	DESCRIPTOR_TYPE_MAX_ENUM                  DescriptorType = 0x7FFFFFFF
)

func (x DescriptorType) String() string {
	switch x {
	case DESCRIPTOR_TYPE_SAMPLER:
		return "DESCRIPTOR_TYPE_SAMPLER"
	case DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER:
		return "DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER"
	case DESCRIPTOR_TYPE_SAMPLED_IMAGE:
		return "DESCRIPTOR_TYPE_SAMPLED_IMAGE"
	case DESCRIPTOR_TYPE_STORAGE_IMAGE:
		return "DESCRIPTOR_TYPE_STORAGE_IMAGE"
	case DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER:
		return "DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER"
	case DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER:
		return "DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER"
	case DESCRIPTOR_TYPE_UNIFORM_BUFFER:
		return "DESCRIPTOR_TYPE_UNIFORM_BUFFER"
	case DESCRIPTOR_TYPE_STORAGE_BUFFER:
		return "DESCRIPTOR_TYPE_STORAGE_BUFFER"
	case DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC:
		return "DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC"
	case DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC:
		return "DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC"
	case DESCRIPTOR_TYPE_INPUT_ATTACHMENT:
		return "DESCRIPTOR_TYPE_INPUT_ATTACHMENT"
	case DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT:
		return "DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT"
	case DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV:
		return "DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV"
	case DESCRIPTOR_TYPE_MAX_ENUM:
		return "DESCRIPTOR_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// AttachmentLoadOp -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentLoadOp.html
type AttachmentLoadOp int32

const (
	ATTACHMENT_LOAD_OP_LOAD        AttachmentLoadOp = 0
	ATTACHMENT_LOAD_OP_CLEAR       AttachmentLoadOp = 1
	ATTACHMENT_LOAD_OP_DONT_CARE   AttachmentLoadOp = 2
	ATTACHMENT_LOAD_OP_BEGIN_RANGE AttachmentLoadOp = ATTACHMENT_LOAD_OP_LOAD
	ATTACHMENT_LOAD_OP_END_RANGE   AttachmentLoadOp = ATTACHMENT_LOAD_OP_DONT_CARE
	ATTACHMENT_LOAD_OP_RANGE_SIZE  AttachmentLoadOp = (ATTACHMENT_LOAD_OP_DONT_CARE - ATTACHMENT_LOAD_OP_LOAD + 1)
	ATTACHMENT_LOAD_OP_MAX_ENUM    AttachmentLoadOp = 0x7FFFFFFF
)

func (x AttachmentLoadOp) String() string {
	switch x {
	case ATTACHMENT_LOAD_OP_LOAD:
		return "ATTACHMENT_LOAD_OP_LOAD"
	case ATTACHMENT_LOAD_OP_CLEAR:
		return "ATTACHMENT_LOAD_OP_CLEAR"
	case ATTACHMENT_LOAD_OP_DONT_CARE:
		return "ATTACHMENT_LOAD_OP_DONT_CARE"
	case ATTACHMENT_LOAD_OP_MAX_ENUM:
		return "ATTACHMENT_LOAD_OP_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// AttachmentStoreOp -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentStoreOp.html
type AttachmentStoreOp int32

const (
	ATTACHMENT_STORE_OP_STORE       AttachmentStoreOp = 0
	ATTACHMENT_STORE_OP_DONT_CARE   AttachmentStoreOp = 1
	ATTACHMENT_STORE_OP_BEGIN_RANGE AttachmentStoreOp = ATTACHMENT_STORE_OP_STORE
	ATTACHMENT_STORE_OP_END_RANGE   AttachmentStoreOp = ATTACHMENT_STORE_OP_DONT_CARE
	ATTACHMENT_STORE_OP_RANGE_SIZE  AttachmentStoreOp = (ATTACHMENT_STORE_OP_DONT_CARE - ATTACHMENT_STORE_OP_STORE + 1)
	ATTACHMENT_STORE_OP_MAX_ENUM    AttachmentStoreOp = 0x7FFFFFFF
)

func (x AttachmentStoreOp) String() string {
	switch x {
	case ATTACHMENT_STORE_OP_STORE:
		return "ATTACHMENT_STORE_OP_STORE"
	case ATTACHMENT_STORE_OP_DONT_CARE:
		return "ATTACHMENT_STORE_OP_DONT_CARE"
	case ATTACHMENT_STORE_OP_MAX_ENUM:
		return "ATTACHMENT_STORE_OP_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PipelineBindPoint -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineBindPoint.html
type PipelineBindPoint int32

const (
	PIPELINE_BIND_POINT_GRAPHICS       PipelineBindPoint = 0
	PIPELINE_BIND_POINT_COMPUTE        PipelineBindPoint = 1
	PIPELINE_BIND_POINT_RAY_TRACING_NV PipelineBindPoint = 1000165000
	PIPELINE_BIND_POINT_BEGIN_RANGE    PipelineBindPoint = PIPELINE_BIND_POINT_GRAPHICS
	PIPELINE_BIND_POINT_END_RANGE      PipelineBindPoint = PIPELINE_BIND_POINT_COMPUTE
	PIPELINE_BIND_POINT_RANGE_SIZE     PipelineBindPoint = (PIPELINE_BIND_POINT_COMPUTE - PIPELINE_BIND_POINT_GRAPHICS + 1)
	PIPELINE_BIND_POINT_MAX_ENUM       PipelineBindPoint = 0x7FFFFFFF
)

func (x PipelineBindPoint) String() string {
	switch x {
	case PIPELINE_BIND_POINT_GRAPHICS:
		return "PIPELINE_BIND_POINT_GRAPHICS"
	case PIPELINE_BIND_POINT_COMPUTE:
		return "PIPELINE_BIND_POINT_COMPUTE"
	case PIPELINE_BIND_POINT_RAY_TRACING_NV:
		return "PIPELINE_BIND_POINT_RAY_TRACING_NV"
	case PIPELINE_BIND_POINT_MAX_ENUM:
		return "PIPELINE_BIND_POINT_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// CommandBufferLevel -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBufferLevel.html
type CommandBufferLevel int32

const (
	COMMAND_BUFFER_LEVEL_PRIMARY     CommandBufferLevel = 0
	COMMAND_BUFFER_LEVEL_SECONDARY   CommandBufferLevel = 1
	COMMAND_BUFFER_LEVEL_BEGIN_RANGE CommandBufferLevel = COMMAND_BUFFER_LEVEL_PRIMARY
	COMMAND_BUFFER_LEVEL_END_RANGE   CommandBufferLevel = COMMAND_BUFFER_LEVEL_SECONDARY
	COMMAND_BUFFER_LEVEL_RANGE_SIZE  CommandBufferLevel = (COMMAND_BUFFER_LEVEL_SECONDARY - COMMAND_BUFFER_LEVEL_PRIMARY + 1)
	COMMAND_BUFFER_LEVEL_MAX_ENUM    CommandBufferLevel = 0x7FFFFFFF
)

func (x CommandBufferLevel) String() string {
	switch x {
	case COMMAND_BUFFER_LEVEL_PRIMARY:
		return "COMMAND_BUFFER_LEVEL_PRIMARY"
	case COMMAND_BUFFER_LEVEL_SECONDARY:
		return "COMMAND_BUFFER_LEVEL_SECONDARY"
	case COMMAND_BUFFER_LEVEL_MAX_ENUM:
		return "COMMAND_BUFFER_LEVEL_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// IndexType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkIndexType.html
type IndexType int32

const (
	INDEX_TYPE_UINT16      IndexType = 0
	INDEX_TYPE_UINT32      IndexType = 1
	INDEX_TYPE_NONE_NV     IndexType = 1000165000
	INDEX_TYPE_UINT8_EXT   IndexType = 1000265000
	INDEX_TYPE_BEGIN_RANGE IndexType = INDEX_TYPE_UINT16
	INDEX_TYPE_END_RANGE   IndexType = INDEX_TYPE_UINT32
	INDEX_TYPE_RANGE_SIZE  IndexType = (INDEX_TYPE_UINT32 - INDEX_TYPE_UINT16 + 1)
	INDEX_TYPE_MAX_ENUM    IndexType = 0x7FFFFFFF
)

func (x IndexType) String() string {
	switch x {
	case INDEX_TYPE_UINT16:
		return "INDEX_TYPE_UINT16"
	case INDEX_TYPE_UINT32:
		return "INDEX_TYPE_UINT32"
	case INDEX_TYPE_NONE_NV:
		return "INDEX_TYPE_NONE_NV"
	case INDEX_TYPE_UINT8_EXT:
		return "INDEX_TYPE_UINT8_EXT"
	case INDEX_TYPE_MAX_ENUM:
		return "INDEX_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SubpassContents -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassContents.html
type SubpassContents int32

const (
	SUBPASS_CONTENTS_INLINE                    SubpassContents = 0
	SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS SubpassContents = 1
	SUBPASS_CONTENTS_BEGIN_RANGE               SubpassContents = SUBPASS_CONTENTS_INLINE
	SUBPASS_CONTENTS_END_RANGE                 SubpassContents = SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	SUBPASS_CONTENTS_RANGE_SIZE                SubpassContents = (SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS - SUBPASS_CONTENTS_INLINE + 1)
	SUBPASS_CONTENTS_MAX_ENUM                  SubpassContents = 0x7FFFFFFF
)

func (x SubpassContents) String() string {
	switch x {
	case SUBPASS_CONTENTS_INLINE:
		return "SUBPASS_CONTENTS_INLINE"
	case SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS:
		return "SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS"
	case SUBPASS_CONTENTS_MAX_ENUM:
		return "SUBPASS_CONTENTS_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ObjectType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectType.html
type ObjectType int32

const (
	OBJECT_TYPE_UNKNOWN                         ObjectType = 0
	OBJECT_TYPE_INSTANCE                        ObjectType = 1
	OBJECT_TYPE_PHYSICAL_DEVICE                 ObjectType = 2
	OBJECT_TYPE_DEVICE                          ObjectType = 3
	OBJECT_TYPE_QUEUE                           ObjectType = 4
	OBJECT_TYPE_SEMAPHORE                       ObjectType = 5
	OBJECT_TYPE_COMMAND_BUFFER                  ObjectType = 6
	OBJECT_TYPE_FENCE                           ObjectType = 7
	OBJECT_TYPE_DEVICE_MEMORY                   ObjectType = 8
	OBJECT_TYPE_BUFFER                          ObjectType = 9
	OBJECT_TYPE_IMAGE                           ObjectType = 10
	OBJECT_TYPE_EVENT                           ObjectType = 11
	OBJECT_TYPE_QUERY_POOL                      ObjectType = 12
	OBJECT_TYPE_BUFFER_VIEW                     ObjectType = 13
	OBJECT_TYPE_IMAGE_VIEW                      ObjectType = 14
	OBJECT_TYPE_SHADER_MODULE                   ObjectType = 15
	OBJECT_TYPE_PIPELINE_CACHE                  ObjectType = 16
	OBJECT_TYPE_PIPELINE_LAYOUT                 ObjectType = 17
	OBJECT_TYPE_RENDER_PASS                     ObjectType = 18
	OBJECT_TYPE_PIPELINE                        ObjectType = 19
	OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT           ObjectType = 20
	OBJECT_TYPE_SAMPLER                         ObjectType = 21
	OBJECT_TYPE_DESCRIPTOR_POOL                 ObjectType = 22
	OBJECT_TYPE_DESCRIPTOR_SET                  ObjectType = 23
	OBJECT_TYPE_FRAMEBUFFER                     ObjectType = 24
	OBJECT_TYPE_COMMAND_POOL                    ObjectType = 25
	OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION        ObjectType = 1000156000
	OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE      ObjectType = 1000085000
	OBJECT_TYPE_SURFACE_KHR                     ObjectType = 1000000000
	OBJECT_TYPE_SWAPCHAIN_KHR                   ObjectType = 1000001000
	OBJECT_TYPE_DISPLAY_KHR                     ObjectType = 1000002000
	OBJECT_TYPE_DISPLAY_MODE_KHR                ObjectType = 1000002001
	OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT       ObjectType = 1000011000
	OBJECT_TYPE_OBJECT_TABLE_NVX                ObjectType = 1000086000
	OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NVX    ObjectType = 1000086001
	OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT       ObjectType = 1000128000
	OBJECT_TYPE_VALIDATION_CACHE_EXT            ObjectType = 1000160000
	OBJECT_TYPE_ACCELERATION_STRUCTURE_NV       ObjectType = 1000165000
	OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL ObjectType = 1000210000
	OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR  ObjectType = OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE
	OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR    ObjectType = OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION
	OBJECT_TYPE_BEGIN_RANGE                     ObjectType = OBJECT_TYPE_UNKNOWN
	OBJECT_TYPE_END_RANGE                       ObjectType = OBJECT_TYPE_COMMAND_POOL
	OBJECT_TYPE_RANGE_SIZE                      ObjectType = (OBJECT_TYPE_COMMAND_POOL - OBJECT_TYPE_UNKNOWN + 1)
	OBJECT_TYPE_MAX_ENUM                        ObjectType = 0x7FFFFFFF
)

func (x ObjectType) String() string {
	switch x {
	case OBJECT_TYPE_UNKNOWN:
		return "OBJECT_TYPE_UNKNOWN"
	case OBJECT_TYPE_INSTANCE:
		return "OBJECT_TYPE_INSTANCE"
	case OBJECT_TYPE_PHYSICAL_DEVICE:
		return "OBJECT_TYPE_PHYSICAL_DEVICE"
	case OBJECT_TYPE_DEVICE:
		return "OBJECT_TYPE_DEVICE"
	case OBJECT_TYPE_QUEUE:
		return "OBJECT_TYPE_QUEUE"
	case OBJECT_TYPE_SEMAPHORE:
		return "OBJECT_TYPE_SEMAPHORE"
	case OBJECT_TYPE_COMMAND_BUFFER:
		return "OBJECT_TYPE_COMMAND_BUFFER"
	case OBJECT_TYPE_FENCE:
		return "OBJECT_TYPE_FENCE"
	case OBJECT_TYPE_DEVICE_MEMORY:
		return "OBJECT_TYPE_DEVICE_MEMORY"
	case OBJECT_TYPE_BUFFER:
		return "OBJECT_TYPE_BUFFER"
	case OBJECT_TYPE_IMAGE:
		return "OBJECT_TYPE_IMAGE"
	case OBJECT_TYPE_EVENT:
		return "OBJECT_TYPE_EVENT"
	case OBJECT_TYPE_QUERY_POOL:
		return "OBJECT_TYPE_QUERY_POOL"
	case OBJECT_TYPE_BUFFER_VIEW:
		return "OBJECT_TYPE_BUFFER_VIEW"
	case OBJECT_TYPE_IMAGE_VIEW:
		return "OBJECT_TYPE_IMAGE_VIEW"
	case OBJECT_TYPE_SHADER_MODULE:
		return "OBJECT_TYPE_SHADER_MODULE"
	case OBJECT_TYPE_PIPELINE_CACHE:
		return "OBJECT_TYPE_PIPELINE_CACHE"
	case OBJECT_TYPE_PIPELINE_LAYOUT:
		return "OBJECT_TYPE_PIPELINE_LAYOUT"
	case OBJECT_TYPE_RENDER_PASS:
		return "OBJECT_TYPE_RENDER_PASS"
	case OBJECT_TYPE_PIPELINE:
		return "OBJECT_TYPE_PIPELINE"
	case OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT:
		return "OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT"
	case OBJECT_TYPE_SAMPLER:
		return "OBJECT_TYPE_SAMPLER"
	case OBJECT_TYPE_DESCRIPTOR_POOL:
		return "OBJECT_TYPE_DESCRIPTOR_POOL"
	case OBJECT_TYPE_DESCRIPTOR_SET:
		return "OBJECT_TYPE_DESCRIPTOR_SET"
	case OBJECT_TYPE_FRAMEBUFFER:
		return "OBJECT_TYPE_FRAMEBUFFER"
	case OBJECT_TYPE_COMMAND_POOL:
		return "OBJECT_TYPE_COMMAND_POOL"
	case OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION:
		return "OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION"
	case OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE:
		return "OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE"
	case OBJECT_TYPE_SURFACE_KHR:
		return "OBJECT_TYPE_SURFACE_KHR"
	case OBJECT_TYPE_SWAPCHAIN_KHR:
		return "OBJECT_TYPE_SWAPCHAIN_KHR"
	case OBJECT_TYPE_DISPLAY_KHR:
		return "OBJECT_TYPE_DISPLAY_KHR"
	case OBJECT_TYPE_DISPLAY_MODE_KHR:
		return "OBJECT_TYPE_DISPLAY_MODE_KHR"
	case OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT:
		return "OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT"
	case OBJECT_TYPE_OBJECT_TABLE_NVX:
		return "OBJECT_TYPE_OBJECT_TABLE_NVX"
	case OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NVX:
		return "OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NVX"
	case OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT:
		return "OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT"
	case OBJECT_TYPE_VALIDATION_CACHE_EXT:
		return "OBJECT_TYPE_VALIDATION_CACHE_EXT"
	case OBJECT_TYPE_ACCELERATION_STRUCTURE_NV:
		return "OBJECT_TYPE_ACCELERATION_STRUCTURE_NV"
	case OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL:
		return "OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL"
	case OBJECT_TYPE_MAX_ENUM:
		return "OBJECT_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// VendorId -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkVendorId.html
type VendorId int32

const (
	VENDOR_ID_VIV         VendorId = 0x10001
	VENDOR_ID_VSI         VendorId = 0x10002
	VENDOR_ID_KAZAN       VendorId = 0x10003
	VENDOR_ID_BEGIN_RANGE VendorId = VENDOR_ID_VIV
	VENDOR_ID_END_RANGE   VendorId = VENDOR_ID_KAZAN
	VENDOR_ID_RANGE_SIZE  VendorId = (VENDOR_ID_KAZAN - VENDOR_ID_VIV + 1)
	VENDOR_ID_MAX_ENUM    VendorId = 0x7FFFFFFF
)

func (x VendorId) String() string {
	switch x {
	case VENDOR_ID_VIV:
		return "VENDOR_ID_VIV"
	case VENDOR_ID_VSI:
		return "VENDOR_ID_VSI"
	case VENDOR_ID_KAZAN:
		return "VENDOR_ID_KAZAN"
	case VENDOR_ID_MAX_ENUM:
		return "VENDOR_ID_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

type InstanceCreateFlags uint32 // reserved
// FormatFeatureFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFormatFeatureFlags.html
type FormatFeatureFlags uint32

const (
	FORMAT_FEATURE_SAMPLED_IMAGE_BIT                                                               FormatFeatureFlags = 0x00000001
	FORMAT_FEATURE_STORAGE_IMAGE_BIT                                                               FormatFeatureFlags = 0x00000002
	FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT                                                        FormatFeatureFlags = 0x00000004
	FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT                                                        FormatFeatureFlags = 0x00000008
	FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT                                                        FormatFeatureFlags = 0x00000010
	FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT                                                 FormatFeatureFlags = 0x00000020
	FORMAT_FEATURE_VERTEX_BUFFER_BIT                                                               FormatFeatureFlags = 0x00000040
	FORMAT_FEATURE_COLOR_ATTACHMENT_BIT                                                            FormatFeatureFlags = 0x00000080
	FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT                                                      FormatFeatureFlags = 0x00000100
	FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT                                                    FormatFeatureFlags = 0x00000200
	FORMAT_FEATURE_BLIT_SRC_BIT                                                                    FormatFeatureFlags = 0x00000400
	FORMAT_FEATURE_BLIT_DST_BIT                                                                    FormatFeatureFlags = 0x00000800
	FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT                                                 FormatFeatureFlags = 0x00001000
	FORMAT_FEATURE_TRANSFER_SRC_BIT                                                                FormatFeatureFlags = 0x00004000
	FORMAT_FEATURE_TRANSFER_DST_BIT                                                                FormatFeatureFlags = 0x00008000
	FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT                                                     FormatFeatureFlags = 0x00020000
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT                                FormatFeatureFlags = 0x00040000
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT               FormatFeatureFlags = 0x00080000
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT               FormatFeatureFlags = 0x00100000
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT     FormatFeatureFlags = 0x00200000
	FORMAT_FEATURE_DISJOINT_BIT                                                                    FormatFeatureFlags = 0x00400000
	FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT                                                      FormatFeatureFlags = 0x00800000
	FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG                                              FormatFeatureFlags = 0x00002000
	FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT                                             FormatFeatureFlags = 0x00010000
	FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT                                                    FormatFeatureFlags = 0x01000000
	FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR                                                            FormatFeatureFlags = FORMAT_FEATURE_TRANSFER_SRC_BIT
	FORMAT_FEATURE_TRANSFER_DST_BIT_KHR                                                            FormatFeatureFlags = FORMAT_FEATURE_TRANSFER_DST_BIT
	FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT_KHR                                                 FormatFeatureFlags = FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR                            FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR           FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR           FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT
	FORMAT_FEATURE_DISJOINT_BIT_KHR                                                                FormatFeatureFlags = FORMAT_FEATURE_DISJOINT_BIT
	FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT_KHR                                                  FormatFeatureFlags = FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT
	FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT                                              FormatFeatureFlags = FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG
	FORMAT_FEATURE_FLAG_BITS_MAX_ENUM                                                              FormatFeatureFlags = 0x7FFFFFFF
)

func (x FormatFeatureFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch FormatFeatureFlags(1 << i) {
			case FORMAT_FEATURE_SAMPLED_IMAGE_BIT:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_BIT|"
			case FORMAT_FEATURE_STORAGE_IMAGE_BIT:
				s += "FORMAT_FEATURE_STORAGE_IMAGE_BIT|"
			case FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT:
				s += "FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT|"
			case FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT:
				s += "FORMAT_FEATURE_UNIFORM_TEXEL_BUFFER_BIT|"
			case FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT:
				s += "FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT|"
			case FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT:
				s += "FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT|"
			case FORMAT_FEATURE_VERTEX_BUFFER_BIT:
				s += "FORMAT_FEATURE_VERTEX_BUFFER_BIT|"
			case FORMAT_FEATURE_COLOR_ATTACHMENT_BIT:
				s += "FORMAT_FEATURE_COLOR_ATTACHMENT_BIT|"
			case FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT:
				s += "FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT|"
			case FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT:
				s += "FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT|"
			case FORMAT_FEATURE_BLIT_SRC_BIT:
				s += "FORMAT_FEATURE_BLIT_SRC_BIT|"
			case FORMAT_FEATURE_BLIT_DST_BIT:
				s += "FORMAT_FEATURE_BLIT_DST_BIT|"
			case FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT|"
			case FORMAT_FEATURE_TRANSFER_SRC_BIT:
				s += "FORMAT_FEATURE_TRANSFER_SRC_BIT|"
			case FORMAT_FEATURE_TRANSFER_DST_BIT:
				s += "FORMAT_FEATURE_TRANSFER_DST_BIT|"
			case FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT:
				s += "FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT|"
			case FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT|"
			case FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT|"
			case FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT|"
			case FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT|"
			case FORMAT_FEATURE_DISJOINT_BIT:
				s += "FORMAT_FEATURE_DISJOINT_BIT|"
			case FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT:
				s += "FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT|"
			case FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG|"
			case FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT:
				s += "FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT|"
			case FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT:
				s += "FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ImageUsageFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageUsageFlags.html
type ImageUsageFlags uint32

const (
	IMAGE_USAGE_TRANSFER_SRC_BIT             ImageUsageFlags = 0x00000001
	IMAGE_USAGE_TRANSFER_DST_BIT             ImageUsageFlags = 0x00000002
	IMAGE_USAGE_SAMPLED_BIT                  ImageUsageFlags = 0x00000004
	IMAGE_USAGE_STORAGE_BIT                  ImageUsageFlags = 0x00000008
	IMAGE_USAGE_COLOR_ATTACHMENT_BIT         ImageUsageFlags = 0x00000010
	IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT ImageUsageFlags = 0x00000020
	IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT     ImageUsageFlags = 0x00000040
	IMAGE_USAGE_INPUT_ATTACHMENT_BIT         ImageUsageFlags = 0x00000080
	IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV    ImageUsageFlags = 0x00000100
	IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT ImageUsageFlags = 0x00000200
	IMAGE_USAGE_FLAG_BITS_MAX_ENUM           ImageUsageFlags = 0x7FFFFFFF
)

func (x ImageUsageFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ImageUsageFlags(1 << i) {
			case IMAGE_USAGE_TRANSFER_SRC_BIT:
				s += "IMAGE_USAGE_TRANSFER_SRC_BIT|"
			case IMAGE_USAGE_TRANSFER_DST_BIT:
				s += "IMAGE_USAGE_TRANSFER_DST_BIT|"
			case IMAGE_USAGE_SAMPLED_BIT:
				s += "IMAGE_USAGE_SAMPLED_BIT|"
			case IMAGE_USAGE_STORAGE_BIT:
				s += "IMAGE_USAGE_STORAGE_BIT|"
			case IMAGE_USAGE_COLOR_ATTACHMENT_BIT:
				s += "IMAGE_USAGE_COLOR_ATTACHMENT_BIT|"
			case IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT:
				s += "IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT|"
			case IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT:
				s += "IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT|"
			case IMAGE_USAGE_INPUT_ATTACHMENT_BIT:
				s += "IMAGE_USAGE_INPUT_ATTACHMENT_BIT|"
			case IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV:
				s += "IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV|"
			case IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT:
				s += "IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ImageCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageCreateFlags.html
type ImageCreateFlags uint32

const (
	IMAGE_CREATE_SPARSE_BINDING_BIT                        ImageCreateFlags = 0x00000001
	IMAGE_CREATE_SPARSE_RESIDENCY_BIT                      ImageCreateFlags = 0x00000002
	IMAGE_CREATE_SPARSE_ALIASED_BIT                        ImageCreateFlags = 0x00000004
	IMAGE_CREATE_MUTABLE_FORMAT_BIT                        ImageCreateFlags = 0x00000008
	IMAGE_CREATE_CUBE_COMPATIBLE_BIT                       ImageCreateFlags = 0x00000010
	IMAGE_CREATE_ALIAS_BIT                                 ImageCreateFlags = 0x00000400
	IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT           ImageCreateFlags = 0x00000040
	IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT                   ImageCreateFlags = 0x00000020
	IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT           ImageCreateFlags = 0x00000080
	IMAGE_CREATE_EXTENDED_USAGE_BIT                        ImageCreateFlags = 0x00000100
	IMAGE_CREATE_PROTECTED_BIT                             ImageCreateFlags = 0x00000800
	IMAGE_CREATE_DISJOINT_BIT                              ImageCreateFlags = 0x00000200
	IMAGE_CREATE_CORNER_SAMPLED_BIT_NV                     ImageCreateFlags = 0x00002000
	IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT ImageCreateFlags = 0x00001000
	IMAGE_CREATE_SUBSAMPLED_BIT_EXT                        ImageCreateFlags = 0x00004000
	IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR       ImageCreateFlags = IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT
	IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT_KHR               ImageCreateFlags = IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT
	IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT_KHR       ImageCreateFlags = IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT
	IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR                    ImageCreateFlags = IMAGE_CREATE_EXTENDED_USAGE_BIT
	IMAGE_CREATE_DISJOINT_BIT_KHR                          ImageCreateFlags = IMAGE_CREATE_DISJOINT_BIT
	IMAGE_CREATE_ALIAS_BIT_KHR                             ImageCreateFlags = IMAGE_CREATE_ALIAS_BIT
	IMAGE_CREATE_FLAG_BITS_MAX_ENUM                        ImageCreateFlags = 0x7FFFFFFF
)

func (x ImageCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ImageCreateFlags(1 << i) {
			case IMAGE_CREATE_SPARSE_BINDING_BIT:
				s += "IMAGE_CREATE_SPARSE_BINDING_BIT|"
			case IMAGE_CREATE_SPARSE_RESIDENCY_BIT:
				s += "IMAGE_CREATE_SPARSE_RESIDENCY_BIT|"
			case IMAGE_CREATE_SPARSE_ALIASED_BIT:
				s += "IMAGE_CREATE_SPARSE_ALIASED_BIT|"
			case IMAGE_CREATE_MUTABLE_FORMAT_BIT:
				s += "IMAGE_CREATE_MUTABLE_FORMAT_BIT|"
			case IMAGE_CREATE_CUBE_COMPATIBLE_BIT:
				s += "IMAGE_CREATE_CUBE_COMPATIBLE_BIT|"
			case IMAGE_CREATE_ALIAS_BIT:
				s += "IMAGE_CREATE_ALIAS_BIT|"
			case IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT:
				s += "IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT|"
			case IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT:
				s += "IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT|"
			case IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT:
				s += "IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT|"
			case IMAGE_CREATE_EXTENDED_USAGE_BIT:
				s += "IMAGE_CREATE_EXTENDED_USAGE_BIT|"
			case IMAGE_CREATE_PROTECTED_BIT:
				s += "IMAGE_CREATE_PROTECTED_BIT|"
			case IMAGE_CREATE_DISJOINT_BIT:
				s += "IMAGE_CREATE_DISJOINT_BIT|"
			case IMAGE_CREATE_CORNER_SAMPLED_BIT_NV:
				s += "IMAGE_CREATE_CORNER_SAMPLED_BIT_NV|"
			case IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT:
				s += "IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT|"
			case IMAGE_CREATE_SUBSAMPLED_BIT_EXT:
				s += "IMAGE_CREATE_SUBSAMPLED_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SampleCountFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSampleCountFlags.html
type SampleCountFlags uint32

const (
	SAMPLE_COUNT_1_BIT              SampleCountFlags = 0x00000001
	SAMPLE_COUNT_2_BIT              SampleCountFlags = 0x00000002
	SAMPLE_COUNT_4_BIT              SampleCountFlags = 0x00000004
	SAMPLE_COUNT_8_BIT              SampleCountFlags = 0x00000008
	SAMPLE_COUNT_16_BIT             SampleCountFlags = 0x00000010
	SAMPLE_COUNT_32_BIT             SampleCountFlags = 0x00000020
	SAMPLE_COUNT_64_BIT             SampleCountFlags = 0x00000040
	SAMPLE_COUNT_FLAG_BITS_MAX_ENUM SampleCountFlags = 0x7FFFFFFF
)

func (x SampleCountFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SampleCountFlags(1 << i) {
			case SAMPLE_COUNT_1_BIT:
				s += "SAMPLE_COUNT_1_BIT|"
			case SAMPLE_COUNT_2_BIT:
				s += "SAMPLE_COUNT_2_BIT|"
			case SAMPLE_COUNT_4_BIT:
				s += "SAMPLE_COUNT_4_BIT|"
			case SAMPLE_COUNT_8_BIT:
				s += "SAMPLE_COUNT_8_BIT|"
			case SAMPLE_COUNT_16_BIT:
				s += "SAMPLE_COUNT_16_BIT|"
			case SAMPLE_COUNT_32_BIT:
				s += "SAMPLE_COUNT_32_BIT|"
			case SAMPLE_COUNT_64_BIT:
				s += "SAMPLE_COUNT_64_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// QueueFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueueFlags.html
type QueueFlags uint32

const (
	QUEUE_GRAPHICS_BIT       QueueFlags = 0x00000001
	QUEUE_COMPUTE_BIT        QueueFlags = 0x00000002
	QUEUE_TRANSFER_BIT       QueueFlags = 0x00000004
	QUEUE_SPARSE_BINDING_BIT QueueFlags = 0x00000008
	QUEUE_PROTECTED_BIT      QueueFlags = 0x00000010
	QUEUE_FLAG_BITS_MAX_ENUM QueueFlags = 0x7FFFFFFF
)

func (x QueueFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch QueueFlags(1 << i) {
			case QUEUE_GRAPHICS_BIT:
				s += "QUEUE_GRAPHICS_BIT|"
			case QUEUE_COMPUTE_BIT:
				s += "QUEUE_COMPUTE_BIT|"
			case QUEUE_TRANSFER_BIT:
				s += "QUEUE_TRANSFER_BIT|"
			case QUEUE_SPARSE_BINDING_BIT:
				s += "QUEUE_SPARSE_BINDING_BIT|"
			case QUEUE_PROTECTED_BIT:
				s += "QUEUE_PROTECTED_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// MemoryPropertyFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryPropertyFlags.html
type MemoryPropertyFlags uint32

const (
	MEMORY_PROPERTY_DEVICE_LOCAL_BIT        MemoryPropertyFlags = 0x00000001
	MEMORY_PROPERTY_HOST_VISIBLE_BIT        MemoryPropertyFlags = 0x00000002
	MEMORY_PROPERTY_HOST_COHERENT_BIT       MemoryPropertyFlags = 0x00000004
	MEMORY_PROPERTY_HOST_CACHED_BIT         MemoryPropertyFlags = 0x00000008
	MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT    MemoryPropertyFlags = 0x00000010
	MEMORY_PROPERTY_PROTECTED_BIT           MemoryPropertyFlags = 0x00000020
	MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD MemoryPropertyFlags = 0x00000040
	MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD MemoryPropertyFlags = 0x00000080
	MEMORY_PROPERTY_FLAG_BITS_MAX_ENUM      MemoryPropertyFlags = 0x7FFFFFFF
)

func (x MemoryPropertyFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch MemoryPropertyFlags(1 << i) {
			case MEMORY_PROPERTY_DEVICE_LOCAL_BIT:
				s += "MEMORY_PROPERTY_DEVICE_LOCAL_BIT|"
			case MEMORY_PROPERTY_HOST_VISIBLE_BIT:
				s += "MEMORY_PROPERTY_HOST_VISIBLE_BIT|"
			case MEMORY_PROPERTY_HOST_COHERENT_BIT:
				s += "MEMORY_PROPERTY_HOST_COHERENT_BIT|"
			case MEMORY_PROPERTY_HOST_CACHED_BIT:
				s += "MEMORY_PROPERTY_HOST_CACHED_BIT|"
			case MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT:
				s += "MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT|"
			case MEMORY_PROPERTY_PROTECTED_BIT:
				s += "MEMORY_PROPERTY_PROTECTED_BIT|"
			case MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD:
				s += "MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD|"
			case MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD:
				s += "MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// MemoryHeapFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryHeapFlags.html
type MemoryHeapFlags uint32

const (
	MEMORY_HEAP_DEVICE_LOCAL_BIT       MemoryHeapFlags = 0x00000001
	MEMORY_HEAP_MULTI_INSTANCE_BIT     MemoryHeapFlags = 0x00000002
	MEMORY_HEAP_MULTI_INSTANCE_BIT_KHR MemoryHeapFlags = MEMORY_HEAP_MULTI_INSTANCE_BIT
	MEMORY_HEAP_FLAG_BITS_MAX_ENUM     MemoryHeapFlags = 0x7FFFFFFF
)

func (x MemoryHeapFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch MemoryHeapFlags(1 << i) {
			case MEMORY_HEAP_DEVICE_LOCAL_BIT:
				s += "MEMORY_HEAP_DEVICE_LOCAL_BIT|"
			case MEMORY_HEAP_MULTI_INSTANCE_BIT:
				s += "MEMORY_HEAP_MULTI_INSTANCE_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type DeviceCreateFlags uint32 // reserved
// DeviceQueueCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceQueueCreateFlags.html
type DeviceQueueCreateFlags uint32

const (
	DEVICE_QUEUE_CREATE_PROTECTED_BIT      DeviceQueueCreateFlags = 0x00000001
	DEVICE_QUEUE_CREATE_FLAG_BITS_MAX_ENUM DeviceQueueCreateFlags = 0x7FFFFFFF
)

func (x DeviceQueueCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DeviceQueueCreateFlags(1 << i) {
			case DEVICE_QUEUE_CREATE_PROTECTED_BIT:
				s += "DEVICE_QUEUE_CREATE_PROTECTED_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// PipelineStageFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineStageFlags.html
type PipelineStageFlags uint32

const (
	PIPELINE_STAGE_TOP_OF_PIPE_BIT                     PipelineStageFlags = 0x00000001
	PIPELINE_STAGE_DRAW_INDIRECT_BIT                   PipelineStageFlags = 0x00000002
	PIPELINE_STAGE_VERTEX_INPUT_BIT                    PipelineStageFlags = 0x00000004
	PIPELINE_STAGE_VERTEX_SHADER_BIT                   PipelineStageFlags = 0x00000008
	PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT     PipelineStageFlags = 0x00000010
	PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT  PipelineStageFlags = 0x00000020
	PIPELINE_STAGE_GEOMETRY_SHADER_BIT                 PipelineStageFlags = 0x00000040
	PIPELINE_STAGE_FRAGMENT_SHADER_BIT                 PipelineStageFlags = 0x00000080
	PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT            PipelineStageFlags = 0x00000100
	PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT             PipelineStageFlags = 0x00000200
	PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT         PipelineStageFlags = 0x00000400
	PIPELINE_STAGE_COMPUTE_SHADER_BIT                  PipelineStageFlags = 0x00000800
	PIPELINE_STAGE_TRANSFER_BIT                        PipelineStageFlags = 0x00001000
	PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT                  PipelineStageFlags = 0x00002000
	PIPELINE_STAGE_HOST_BIT                            PipelineStageFlags = 0x00004000
	PIPELINE_STAGE_ALL_GRAPHICS_BIT                    PipelineStageFlags = 0x00008000
	PIPELINE_STAGE_ALL_COMMANDS_BIT                    PipelineStageFlags = 0x00010000
	PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT          PipelineStageFlags = 0x01000000
	PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT       PipelineStageFlags = 0x00040000
	PIPELINE_STAGE_COMMAND_PROCESS_BIT_NVX             PipelineStageFlags = 0x00020000
	PIPELINE_STAGE_SHADING_RATE_IMAGE_BIT_NV           PipelineStageFlags = 0x00400000
	PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_NV           PipelineStageFlags = 0x00200000
	PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_NV PipelineStageFlags = 0x02000000
	PIPELINE_STAGE_TASK_SHADER_BIT_NV                  PipelineStageFlags = 0x00080000
	PIPELINE_STAGE_MESH_SHADER_BIT_NV                  PipelineStageFlags = 0x00100000
	PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT    PipelineStageFlags = 0x00800000
	PIPELINE_STAGE_FLAG_BITS_MAX_ENUM                  PipelineStageFlags = 0x7FFFFFFF
)

func (x PipelineStageFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch PipelineStageFlags(1 << i) {
			case PIPELINE_STAGE_TOP_OF_PIPE_BIT:
				s += "PIPELINE_STAGE_TOP_OF_PIPE_BIT|"
			case PIPELINE_STAGE_DRAW_INDIRECT_BIT:
				s += "PIPELINE_STAGE_DRAW_INDIRECT_BIT|"
			case PIPELINE_STAGE_VERTEX_INPUT_BIT:
				s += "PIPELINE_STAGE_VERTEX_INPUT_BIT|"
			case PIPELINE_STAGE_VERTEX_SHADER_BIT:
				s += "PIPELINE_STAGE_VERTEX_SHADER_BIT|"
			case PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT:
				s += "PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT|"
			case PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT:
				s += "PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT|"
			case PIPELINE_STAGE_GEOMETRY_SHADER_BIT:
				s += "PIPELINE_STAGE_GEOMETRY_SHADER_BIT|"
			case PIPELINE_STAGE_FRAGMENT_SHADER_BIT:
				s += "PIPELINE_STAGE_FRAGMENT_SHADER_BIT|"
			case PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT:
				s += "PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT|"
			case PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT:
				s += "PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT|"
			case PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT:
				s += "PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT|"
			case PIPELINE_STAGE_COMPUTE_SHADER_BIT:
				s += "PIPELINE_STAGE_COMPUTE_SHADER_BIT|"
			case PIPELINE_STAGE_TRANSFER_BIT:
				s += "PIPELINE_STAGE_TRANSFER_BIT|"
			case PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT:
				s += "PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT|"
			case PIPELINE_STAGE_HOST_BIT:
				s += "PIPELINE_STAGE_HOST_BIT|"
			case PIPELINE_STAGE_ALL_GRAPHICS_BIT:
				s += "PIPELINE_STAGE_ALL_GRAPHICS_BIT|"
			case PIPELINE_STAGE_ALL_COMMANDS_BIT:
				s += "PIPELINE_STAGE_ALL_COMMANDS_BIT|"
			case PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT:
				s += "PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT|"
			case PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT:
				s += "PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT|"
			case PIPELINE_STAGE_COMMAND_PROCESS_BIT_NVX:
				s += "PIPELINE_STAGE_COMMAND_PROCESS_BIT_NVX|"
			case PIPELINE_STAGE_SHADING_RATE_IMAGE_BIT_NV:
				s += "PIPELINE_STAGE_SHADING_RATE_IMAGE_BIT_NV|"
			case PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_NV:
				s += "PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_NV|"
			case PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_NV:
				s += "PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_NV|"
			case PIPELINE_STAGE_TASK_SHADER_BIT_NV:
				s += "PIPELINE_STAGE_TASK_SHADER_BIT_NV|"
			case PIPELINE_STAGE_MESH_SHADER_BIT_NV:
				s += "PIPELINE_STAGE_MESH_SHADER_BIT_NV|"
			case PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT:
				s += "PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type MemoryMapFlags uint32 // reserved
// ImageAspectFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageAspectFlags.html
type ImageAspectFlags uint32

const (
	IMAGE_ASPECT_COLOR_BIT              ImageAspectFlags = 0x00000001
	IMAGE_ASPECT_DEPTH_BIT              ImageAspectFlags = 0x00000002
	IMAGE_ASPECT_STENCIL_BIT            ImageAspectFlags = 0x00000004
	IMAGE_ASPECT_METADATA_BIT           ImageAspectFlags = 0x00000008
	IMAGE_ASPECT_PLANE_0_BIT            ImageAspectFlags = 0x00000010
	IMAGE_ASPECT_PLANE_1_BIT            ImageAspectFlags = 0x00000020
	IMAGE_ASPECT_PLANE_2_BIT            ImageAspectFlags = 0x00000040
	IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT ImageAspectFlags = 0x00000080
	IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT ImageAspectFlags = 0x00000100
	IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT ImageAspectFlags = 0x00000200
	IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT ImageAspectFlags = 0x00000400
	IMAGE_ASPECT_PLANE_0_BIT_KHR        ImageAspectFlags = IMAGE_ASPECT_PLANE_0_BIT
	IMAGE_ASPECT_PLANE_1_BIT_KHR        ImageAspectFlags = IMAGE_ASPECT_PLANE_1_BIT
	IMAGE_ASPECT_PLANE_2_BIT_KHR        ImageAspectFlags = IMAGE_ASPECT_PLANE_2_BIT
	IMAGE_ASPECT_FLAG_BITS_MAX_ENUM     ImageAspectFlags = 0x7FFFFFFF
)

func (x ImageAspectFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ImageAspectFlags(1 << i) {
			case IMAGE_ASPECT_COLOR_BIT:
				s += "IMAGE_ASPECT_COLOR_BIT|"
			case IMAGE_ASPECT_DEPTH_BIT:
				s += "IMAGE_ASPECT_DEPTH_BIT|"
			case IMAGE_ASPECT_STENCIL_BIT:
				s += "IMAGE_ASPECT_STENCIL_BIT|"
			case IMAGE_ASPECT_METADATA_BIT:
				s += "IMAGE_ASPECT_METADATA_BIT|"
			case IMAGE_ASPECT_PLANE_0_BIT:
				s += "IMAGE_ASPECT_PLANE_0_BIT|"
			case IMAGE_ASPECT_PLANE_1_BIT:
				s += "IMAGE_ASPECT_PLANE_1_BIT|"
			case IMAGE_ASPECT_PLANE_2_BIT:
				s += "IMAGE_ASPECT_PLANE_2_BIT|"
			case IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT:
				s += "IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT|"
			case IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT:
				s += "IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT|"
			case IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT:
				s += "IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT|"
			case IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT:
				s += "IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SparseImageFormatFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageFormatFlags.html
type SparseImageFormatFlags uint32

const (
	SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT         SparseImageFormatFlags = 0x00000001
	SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT       SparseImageFormatFlags = 0x00000002
	SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT SparseImageFormatFlags = 0x00000004
	SPARSE_IMAGE_FORMAT_FLAG_BITS_MAX_ENUM         SparseImageFormatFlags = 0x7FFFFFFF
)

func (x SparseImageFormatFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SparseImageFormatFlags(1 << i) {
			case SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT:
				s += "SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT|"
			case SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT:
				s += "SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT|"
			case SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT:
				s += "SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SparseMemoryBindFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseMemoryBindFlags.html
type SparseMemoryBindFlags uint32

const (
	SPARSE_MEMORY_BIND_METADATA_BIT       SparseMemoryBindFlags = 0x00000001
	SPARSE_MEMORY_BIND_FLAG_BITS_MAX_ENUM SparseMemoryBindFlags = 0x7FFFFFFF
)

func (x SparseMemoryBindFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SparseMemoryBindFlags(1 << i) {
			case SPARSE_MEMORY_BIND_METADATA_BIT:
				s += "SPARSE_MEMORY_BIND_METADATA_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// FenceCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFenceCreateFlags.html
type FenceCreateFlags uint32

const (
	FENCE_CREATE_SIGNALED_BIT       FenceCreateFlags = 0x00000001
	FENCE_CREATE_FLAG_BITS_MAX_ENUM FenceCreateFlags = 0x7FFFFFFF
)

func (x FenceCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch FenceCreateFlags(1 << i) {
			case FENCE_CREATE_SIGNALED_BIT:
				s += "FENCE_CREATE_SIGNALED_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type SemaphoreCreateFlags uint32 // reserved
type EventCreateFlags uint32     // reserved
type QueryPoolCreateFlags uint32 // reserved
// QueryPipelineStatisticFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryPipelineStatisticFlags.html
type QueryPipelineStatisticFlags uint32

const (
	QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT                    QueryPipelineStatisticFlags = 0x00000001
	QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT                  QueryPipelineStatisticFlags = 0x00000002
	QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT                  QueryPipelineStatisticFlags = 0x00000004
	QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT                QueryPipelineStatisticFlags = 0x00000008
	QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT                 QueryPipelineStatisticFlags = 0x00000010
	QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT                       QueryPipelineStatisticFlags = 0x00000020
	QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT                        QueryPipelineStatisticFlags = 0x00000040
	QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT                QueryPipelineStatisticFlags = 0x00000080
	QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT        QueryPipelineStatisticFlags = 0x00000100
	QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT QueryPipelineStatisticFlags = 0x00000200
	QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT                 QueryPipelineStatisticFlags = 0x00000400
	QUERY_PIPELINE_STATISTIC_FLAG_BITS_MAX_ENUM                             QueryPipelineStatisticFlags = 0x7FFFFFFF
)

func (x QueryPipelineStatisticFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch QueryPipelineStatisticFlags(1 << i) {
			case QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT:
				s += "QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT|"
			case QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT:
				s += "QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT|"
			case QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT:
				s += "QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT|"
			case QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT:
				s += "QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT|"
			case QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT:
				s += "QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT|"
			case QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT:
				s += "QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT|"
			case QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT:
				s += "QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT|"
			case QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT:
				s += "QUERY_PIPELINE_STATISTIC_FRAGMENT_SHADER_INVOCATIONS_BIT|"
			case QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT:
				s += "QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT|"
			case QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT:
				s += "QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT|"
			case QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT:
				s += "QUERY_PIPELINE_STATISTIC_COMPUTE_SHADER_INVOCATIONS_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// QueryResultFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryResultFlags.html
type QueryResultFlags uint32

const (
	QUERY_RESULT_64_BIT                QueryResultFlags = 0x00000001
	QUERY_RESULT_WAIT_BIT              QueryResultFlags = 0x00000002
	QUERY_RESULT_WITH_AVAILABILITY_BIT QueryResultFlags = 0x00000004
	QUERY_RESULT_PARTIAL_BIT           QueryResultFlags = 0x00000008
	QUERY_RESULT_FLAG_BITS_MAX_ENUM    QueryResultFlags = 0x7FFFFFFF
)

func (x QueryResultFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch QueryResultFlags(1 << i) {
			case QUERY_RESULT_64_BIT:
				s += "QUERY_RESULT_64_BIT|"
			case QUERY_RESULT_WAIT_BIT:
				s += "QUERY_RESULT_WAIT_BIT|"
			case QUERY_RESULT_WITH_AVAILABILITY_BIT:
				s += "QUERY_RESULT_WITH_AVAILABILITY_BIT|"
			case QUERY_RESULT_PARTIAL_BIT:
				s += "QUERY_RESULT_PARTIAL_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// BufferCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferCreateFlags.html
type BufferCreateFlags uint32

const (
	BUFFER_CREATE_SPARSE_BINDING_BIT                    BufferCreateFlags = 0x00000001
	BUFFER_CREATE_SPARSE_RESIDENCY_BIT                  BufferCreateFlags = 0x00000002
	BUFFER_CREATE_SPARSE_ALIASED_BIT                    BufferCreateFlags = 0x00000004
	BUFFER_CREATE_PROTECTED_BIT                         BufferCreateFlags = 0x00000008
	BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT BufferCreateFlags = 0x00000010
	BUFFER_CREATE_FLAG_BITS_MAX_ENUM                    BufferCreateFlags = 0x7FFFFFFF
)

func (x BufferCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch BufferCreateFlags(1 << i) {
			case BUFFER_CREATE_SPARSE_BINDING_BIT:
				s += "BUFFER_CREATE_SPARSE_BINDING_BIT|"
			case BUFFER_CREATE_SPARSE_RESIDENCY_BIT:
				s += "BUFFER_CREATE_SPARSE_RESIDENCY_BIT|"
			case BUFFER_CREATE_SPARSE_ALIASED_BIT:
				s += "BUFFER_CREATE_SPARSE_ALIASED_BIT|"
			case BUFFER_CREATE_PROTECTED_BIT:
				s += "BUFFER_CREATE_PROTECTED_BIT|"
			case BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT:
				s += "BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// BufferUsageFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferUsageFlags.html
type BufferUsageFlags uint32

const (
	BUFFER_USAGE_TRANSFER_SRC_BIT                          BufferUsageFlags = 0x00000001
	BUFFER_USAGE_TRANSFER_DST_BIT                          BufferUsageFlags = 0x00000002
	BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT                  BufferUsageFlags = 0x00000004
	BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT                  BufferUsageFlags = 0x00000008
	BUFFER_USAGE_UNIFORM_BUFFER_BIT                        BufferUsageFlags = 0x00000010
	BUFFER_USAGE_STORAGE_BUFFER_BIT                        BufferUsageFlags = 0x00000020
	BUFFER_USAGE_INDEX_BUFFER_BIT                          BufferUsageFlags = 0x00000040
	BUFFER_USAGE_VERTEX_BUFFER_BIT                         BufferUsageFlags = 0x00000080
	BUFFER_USAGE_INDIRECT_BUFFER_BIT                       BufferUsageFlags = 0x00000100
	BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT         BufferUsageFlags = 0x00000800
	BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT BufferUsageFlags = 0x00001000
	BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT             BufferUsageFlags = 0x00000200
	BUFFER_USAGE_RAY_TRACING_BIT_NV                        BufferUsageFlags = 0x00000400
	BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_EXT             BufferUsageFlags = 0x00020000
	BUFFER_USAGE_FLAG_BITS_MAX_ENUM                        BufferUsageFlags = 0x7FFFFFFF
)

func (x BufferUsageFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch BufferUsageFlags(1 << i) {
			case BUFFER_USAGE_TRANSFER_SRC_BIT:
				s += "BUFFER_USAGE_TRANSFER_SRC_BIT|"
			case BUFFER_USAGE_TRANSFER_DST_BIT:
				s += "BUFFER_USAGE_TRANSFER_DST_BIT|"
			case BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT:
				s += "BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT|"
			case BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT:
				s += "BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT|"
			case BUFFER_USAGE_UNIFORM_BUFFER_BIT:
				s += "BUFFER_USAGE_UNIFORM_BUFFER_BIT|"
			case BUFFER_USAGE_STORAGE_BUFFER_BIT:
				s += "BUFFER_USAGE_STORAGE_BUFFER_BIT|"
			case BUFFER_USAGE_INDEX_BUFFER_BIT:
				s += "BUFFER_USAGE_INDEX_BUFFER_BIT|"
			case BUFFER_USAGE_VERTEX_BUFFER_BIT:
				s += "BUFFER_USAGE_VERTEX_BUFFER_BIT|"
			case BUFFER_USAGE_INDIRECT_BUFFER_BIT:
				s += "BUFFER_USAGE_INDIRECT_BUFFER_BIT|"
			case BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT:
				s += "BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT|"
			case BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT:
				s += "BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT|"
			case BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT:
				s += "BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT|"
			case BUFFER_USAGE_RAY_TRACING_BIT_NV:
				s += "BUFFER_USAGE_RAY_TRACING_BIT_NV|"
			case BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_EXT:
				s += "BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type BufferViewCreateFlags uint32 // reserved
// ImageViewCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageViewCreateFlags.html
type ImageViewCreateFlags uint32

const (
	IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT ImageViewCreateFlags = 0x00000001
	IMAGE_VIEW_CREATE_FLAG_BITS_MAX_ENUM                   ImageViewCreateFlags = 0x7FFFFFFF
)

func (x ImageViewCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ImageViewCreateFlags(1 << i) {
			case IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT:
				s += "IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ShaderModuleCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderModuleCreateFlags.html
type ShaderModuleCreateFlags uint32

const (
	SHADER_MODULE_CREATE_FLAG_BITS_MAX_ENUM ShaderModuleCreateFlags = 0x7FFFFFFF
)

func (x ShaderModuleCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ShaderModuleCreateFlags(1 << i) {
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type PipelineCacheCreateFlags uint32 // reserved
// PipelineCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCreateFlags.html
type PipelineCreateFlags uint32

const (
	PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT                 PipelineCreateFlags = 0x00000001
	PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT                    PipelineCreateFlags = 0x00000002
	PIPELINE_CREATE_DERIVATIVE_BIT                           PipelineCreateFlags = 0x00000004
	PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT         PipelineCreateFlags = 0x00000008
	PIPELINE_CREATE_DISPATCH_BASE                            PipelineCreateFlags = 0x00000010
	PIPELINE_CREATE_DEFER_COMPILE_BIT_NV                     PipelineCreateFlags = 0x00000020
	PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR               PipelineCreateFlags = 0x00000040
	PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR PipelineCreateFlags = 0x00000080
	PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR     PipelineCreateFlags = PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT
	PIPELINE_CREATE_DISPATCH_BASE_KHR                        PipelineCreateFlags = PIPELINE_CREATE_DISPATCH_BASE
	PIPELINE_CREATE_FLAG_BITS_MAX_ENUM                       PipelineCreateFlags = 0x7FFFFFFF
)

func (x PipelineCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch PipelineCreateFlags(1 << i) {
			case PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT:
				s += "PIPELINE_CREATE_DISABLE_OPTIMIZATION_BIT|"
			case PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT:
				s += "PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT|"
			case PIPELINE_CREATE_DERIVATIVE_BIT:
				s += "PIPELINE_CREATE_DERIVATIVE_BIT|"
			case PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT:
				s += "PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT|"
			case PIPELINE_CREATE_DISPATCH_BASE:
				s += "PIPELINE_CREATE_DISPATCH_BASE|"
			case PIPELINE_CREATE_DEFER_COMPILE_BIT_NV:
				s += "PIPELINE_CREATE_DEFER_COMPILE_BIT_NV|"
			case PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR:
				s += "PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR|"
			case PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR:
				s += "PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// PipelineShaderStageCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineShaderStageCreateFlags.html
type PipelineShaderStageCreateFlags uint32

const (
	PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT PipelineShaderStageCreateFlags = 0x00000001
	PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT      PipelineShaderStageCreateFlags = 0x00000002
	PIPELINE_SHADER_STAGE_CREATE_FLAG_BITS_MAX_ENUM                  PipelineShaderStageCreateFlags = 0x7FFFFFFF
)

func (x PipelineShaderStageCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch PipelineShaderStageCreateFlags(1 << i) {
			case PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT:
				s += "PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT|"
			case PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT:
				s += "PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ShaderStageFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderStageFlags.html
type ShaderStageFlags uint32

const (
	SHADER_STAGE_VERTEX_BIT                  ShaderStageFlags = 0x00000001
	SHADER_STAGE_TESSELLATION_CONTROL_BIT    ShaderStageFlags = 0x00000002
	SHADER_STAGE_TESSELLATION_EVALUATION_BIT ShaderStageFlags = 0x00000004
	SHADER_STAGE_GEOMETRY_BIT                ShaderStageFlags = 0x00000008
	SHADER_STAGE_FRAGMENT_BIT                ShaderStageFlags = 0x00000010
	SHADER_STAGE_COMPUTE_BIT                 ShaderStageFlags = 0x00000020
	SHADER_STAGE_ALL_GRAPHICS                ShaderStageFlags = 0x0000001F
	SHADER_STAGE_ALL                         ShaderStageFlags = 0x7FFFFFFF
	SHADER_STAGE_RAYGEN_BIT_NV               ShaderStageFlags = 0x00000100
	SHADER_STAGE_ANY_HIT_BIT_NV              ShaderStageFlags = 0x00000200
	SHADER_STAGE_CLOSEST_HIT_BIT_NV          ShaderStageFlags = 0x00000400
	SHADER_STAGE_MISS_BIT_NV                 ShaderStageFlags = 0x00000800
	SHADER_STAGE_INTERSECTION_BIT_NV         ShaderStageFlags = 0x00001000
	SHADER_STAGE_CALLABLE_BIT_NV             ShaderStageFlags = 0x00002000
	SHADER_STAGE_TASK_BIT_NV                 ShaderStageFlags = 0x00000040
	SHADER_STAGE_MESH_BIT_NV                 ShaderStageFlags = 0x00000080
	SHADER_STAGE_FLAG_BITS_MAX_ENUM          ShaderStageFlags = 0x7FFFFFFF
)

func (x ShaderStageFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ShaderStageFlags(1 << i) {
			case SHADER_STAGE_VERTEX_BIT:
				s += "SHADER_STAGE_VERTEX_BIT|"
			case SHADER_STAGE_TESSELLATION_CONTROL_BIT:
				s += "SHADER_STAGE_TESSELLATION_CONTROL_BIT|"
			case SHADER_STAGE_TESSELLATION_EVALUATION_BIT:
				s += "SHADER_STAGE_TESSELLATION_EVALUATION_BIT|"
			case SHADER_STAGE_GEOMETRY_BIT:
				s += "SHADER_STAGE_GEOMETRY_BIT|"
			case SHADER_STAGE_FRAGMENT_BIT:
				s += "SHADER_STAGE_FRAGMENT_BIT|"
			case SHADER_STAGE_COMPUTE_BIT:
				s += "SHADER_STAGE_COMPUTE_BIT|"
			case SHADER_STAGE_ALL_GRAPHICS:
				s += "SHADER_STAGE_ALL_GRAPHICS|"
			case SHADER_STAGE_ALL:
				s += "SHADER_STAGE_ALL|"
			case SHADER_STAGE_RAYGEN_BIT_NV:
				s += "SHADER_STAGE_RAYGEN_BIT_NV|"
			case SHADER_STAGE_ANY_HIT_BIT_NV:
				s += "SHADER_STAGE_ANY_HIT_BIT_NV|"
			case SHADER_STAGE_CLOSEST_HIT_BIT_NV:
				s += "SHADER_STAGE_CLOSEST_HIT_BIT_NV|"
			case SHADER_STAGE_MISS_BIT_NV:
				s += "SHADER_STAGE_MISS_BIT_NV|"
			case SHADER_STAGE_INTERSECTION_BIT_NV:
				s += "SHADER_STAGE_INTERSECTION_BIT_NV|"
			case SHADER_STAGE_CALLABLE_BIT_NV:
				s += "SHADER_STAGE_CALLABLE_BIT_NV|"
			case SHADER_STAGE_TASK_BIT_NV:
				s += "SHADER_STAGE_TASK_BIT_NV|"
			case SHADER_STAGE_MESH_BIT_NV:
				s += "SHADER_STAGE_MESH_BIT_NV|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type PipelineVertexInputStateCreateFlags uint32   // reserved
type PipelineInputAssemblyStateCreateFlags uint32 // reserved
type PipelineTessellationStateCreateFlags uint32  // reserved
type PipelineViewportStateCreateFlags uint32      // reserved
type PipelineRasterizationStateCreateFlags uint32 // reserved
// CullModeFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCullModeFlags.html
type CullModeFlags uint32

const (
	CULL_MODE_NONE               CullModeFlags = 0
	CULL_MODE_FRONT_BIT          CullModeFlags = 0x00000001
	CULL_MODE_BACK_BIT           CullModeFlags = 0x00000002
	CULL_MODE_FRONT_AND_BACK     CullModeFlags = 0x00000003
	CULL_MODE_FLAG_BITS_MAX_ENUM CullModeFlags = 0x7FFFFFFF
)

func (x CullModeFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch CullModeFlags(1 << i) {
			case CULL_MODE_NONE:
				s += "CULL_MODE_NONE|"
			case CULL_MODE_FRONT_BIT:
				s += "CULL_MODE_FRONT_BIT|"
			case CULL_MODE_BACK_BIT:
				s += "CULL_MODE_BACK_BIT|"
			case CULL_MODE_FRONT_AND_BACK:
				s += "CULL_MODE_FRONT_AND_BACK|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type PipelineMultisampleStateCreateFlags uint32  // reserved
type PipelineDepthStencilStateCreateFlags uint32 // reserved
type PipelineColorBlendStateCreateFlags uint32   // reserved
// ColorComponentFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkColorComponentFlags.html
type ColorComponentFlags uint32

const (
	COLOR_COMPONENT_R_BIT              ColorComponentFlags = 0x00000001
	COLOR_COMPONENT_G_BIT              ColorComponentFlags = 0x00000002
	COLOR_COMPONENT_B_BIT              ColorComponentFlags = 0x00000004
	COLOR_COMPONENT_A_BIT              ColorComponentFlags = 0x00000008
	COLOR_COMPONENT_FLAG_BITS_MAX_ENUM ColorComponentFlags = 0x7FFFFFFF
)

func (x ColorComponentFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ColorComponentFlags(1 << i) {
			case COLOR_COMPONENT_R_BIT:
				s += "COLOR_COMPONENT_R_BIT|"
			case COLOR_COMPONENT_G_BIT:
				s += "COLOR_COMPONENT_G_BIT|"
			case COLOR_COMPONENT_B_BIT:
				s += "COLOR_COMPONENT_B_BIT|"
			case COLOR_COMPONENT_A_BIT:
				s += "COLOR_COMPONENT_A_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type PipelineDynamicStateCreateFlags uint32 // reserved
type PipelineLayoutCreateFlags uint32       // reserved
// SamplerCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerCreateFlags.html
type SamplerCreateFlags uint32

const (
	SAMPLER_CREATE_SUBSAMPLED_BIT_EXT                       SamplerCreateFlags = 0x00000001
	SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT SamplerCreateFlags = 0x00000002
	SAMPLER_CREATE_FLAG_BITS_MAX_ENUM                       SamplerCreateFlags = 0x7FFFFFFF
)

func (x SamplerCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SamplerCreateFlags(1 << i) {
			case SAMPLER_CREATE_SUBSAMPLED_BIT_EXT:
				s += "SAMPLER_CREATE_SUBSAMPLED_BIT_EXT|"
			case SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT:
				s += "SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DescriptorSetLayoutCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetLayoutCreateFlags.html
type DescriptorSetLayoutCreateFlags uint32

const (
	DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR        DescriptorSetLayoutCreateFlags = 0x00000001
	DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT DescriptorSetLayoutCreateFlags = 0x00000002
	DESCRIPTOR_SET_LAYOUT_CREATE_FLAG_BITS_MAX_ENUM             DescriptorSetLayoutCreateFlags = 0x7FFFFFFF
)

func (x DescriptorSetLayoutCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DescriptorSetLayoutCreateFlags(1 << i) {
			case DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR:
				s += "DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR|"
			case DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT:
				s += "DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DescriptorPoolCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorPoolCreateFlags.html
type DescriptorPoolCreateFlags uint32

const (
	DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT   DescriptorPoolCreateFlags = 0x00000001
	DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT DescriptorPoolCreateFlags = 0x00000002
	DESCRIPTOR_POOL_CREATE_FLAG_BITS_MAX_ENUM        DescriptorPoolCreateFlags = 0x7FFFFFFF
)

func (x DescriptorPoolCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DescriptorPoolCreateFlags(1 << i) {
			case DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT:
				s += "DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT|"
			case DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT:
				s += "DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type DescriptorPoolResetFlags uint32 // reserved
// FramebufferCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFramebufferCreateFlags.html
type FramebufferCreateFlags uint32

const (
	FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR  FramebufferCreateFlags = 0x00000001
	FRAMEBUFFER_CREATE_FLAG_BITS_MAX_ENUM FramebufferCreateFlags = 0x7FFFFFFF
)

func (x FramebufferCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch FramebufferCreateFlags(1 << i) {
			case FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR:
				s += "FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// RenderPassCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassCreateFlags.html
type RenderPassCreateFlags uint32

const (
	RENDER_PASS_CREATE_FLAG_BITS_MAX_ENUM RenderPassCreateFlags = 0x7FFFFFFF
)

func (x RenderPassCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch RenderPassCreateFlags(1 << i) {
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// AttachmentDescriptionFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentDescriptionFlags.html
type AttachmentDescriptionFlags uint32

const (
	ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT      AttachmentDescriptionFlags = 0x00000001
	ATTACHMENT_DESCRIPTION_FLAG_BITS_MAX_ENUM AttachmentDescriptionFlags = 0x7FFFFFFF
)

func (x AttachmentDescriptionFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch AttachmentDescriptionFlags(1 << i) {
			case ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT:
				s += "ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SubpassDescriptionFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassDescriptionFlags.html
type SubpassDescriptionFlags uint32

const (
	SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX      SubpassDescriptionFlags = 0x00000001
	SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX SubpassDescriptionFlags = 0x00000002
	SUBPASS_DESCRIPTION_FLAG_BITS_MAX_ENUM               SubpassDescriptionFlags = 0x7FFFFFFF
)

func (x SubpassDescriptionFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SubpassDescriptionFlags(1 << i) {
			case SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX:
				s += "SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX|"
			case SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX:
				s += "SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// AccessFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccessFlags.html
type AccessFlags uint32

const (
	ACCESS_INDIRECT_COMMAND_READ_BIT                 AccessFlags = 0x00000001
	ACCESS_INDEX_READ_BIT                            AccessFlags = 0x00000002
	ACCESS_VERTEX_ATTRIBUTE_READ_BIT                 AccessFlags = 0x00000004
	ACCESS_UNIFORM_READ_BIT                          AccessFlags = 0x00000008
	ACCESS_INPUT_ATTACHMENT_READ_BIT                 AccessFlags = 0x00000010
	ACCESS_SHADER_READ_BIT                           AccessFlags = 0x00000020
	ACCESS_SHADER_WRITE_BIT                          AccessFlags = 0x00000040
	ACCESS_COLOR_ATTACHMENT_READ_BIT                 AccessFlags = 0x00000080
	ACCESS_COLOR_ATTACHMENT_WRITE_BIT                AccessFlags = 0x00000100
	ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT         AccessFlags = 0x00000200
	ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT        AccessFlags = 0x00000400
	ACCESS_TRANSFER_READ_BIT                         AccessFlags = 0x00000800
	ACCESS_TRANSFER_WRITE_BIT                        AccessFlags = 0x00001000
	ACCESS_HOST_READ_BIT                             AccessFlags = 0x00002000
	ACCESS_HOST_WRITE_BIT                            AccessFlags = 0x00004000
	ACCESS_MEMORY_READ_BIT                           AccessFlags = 0x00008000
	ACCESS_MEMORY_WRITE_BIT                          AccessFlags = 0x00010000
	ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT          AccessFlags = 0x02000000
	ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT   AccessFlags = 0x04000000
	ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT  AccessFlags = 0x08000000
	ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT        AccessFlags = 0x00100000
	ACCESS_COMMAND_PROCESS_READ_BIT_NVX              AccessFlags = 0x00020000
	ACCESS_COMMAND_PROCESS_WRITE_BIT_NVX             AccessFlags = 0x00040000
	ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT AccessFlags = 0x00080000
	ACCESS_SHADING_RATE_IMAGE_READ_BIT_NV            AccessFlags = 0x00800000
	ACCESS_ACCELERATION_STRUCTURE_READ_BIT_NV        AccessFlags = 0x00200000
	ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_NV       AccessFlags = 0x00400000
	ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT         AccessFlags = 0x01000000
	ACCESS_FLAG_BITS_MAX_ENUM                        AccessFlags = 0x7FFFFFFF
)

func (x AccessFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch AccessFlags(1 << i) {
			case ACCESS_INDIRECT_COMMAND_READ_BIT:
				s += "ACCESS_INDIRECT_COMMAND_READ_BIT|"
			case ACCESS_INDEX_READ_BIT:
				s += "ACCESS_INDEX_READ_BIT|"
			case ACCESS_VERTEX_ATTRIBUTE_READ_BIT:
				s += "ACCESS_VERTEX_ATTRIBUTE_READ_BIT|"
			case ACCESS_UNIFORM_READ_BIT:
				s += "ACCESS_UNIFORM_READ_BIT|"
			case ACCESS_INPUT_ATTACHMENT_READ_BIT:
				s += "ACCESS_INPUT_ATTACHMENT_READ_BIT|"
			case ACCESS_SHADER_READ_BIT:
				s += "ACCESS_SHADER_READ_BIT|"
			case ACCESS_SHADER_WRITE_BIT:
				s += "ACCESS_SHADER_WRITE_BIT|"
			case ACCESS_COLOR_ATTACHMENT_READ_BIT:
				s += "ACCESS_COLOR_ATTACHMENT_READ_BIT|"
			case ACCESS_COLOR_ATTACHMENT_WRITE_BIT:
				s += "ACCESS_COLOR_ATTACHMENT_WRITE_BIT|"
			case ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT:
				s += "ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT|"
			case ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT:
				s += "ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT|"
			case ACCESS_TRANSFER_READ_BIT:
				s += "ACCESS_TRANSFER_READ_BIT|"
			case ACCESS_TRANSFER_WRITE_BIT:
				s += "ACCESS_TRANSFER_WRITE_BIT|"
			case ACCESS_HOST_READ_BIT:
				s += "ACCESS_HOST_READ_BIT|"
			case ACCESS_HOST_WRITE_BIT:
				s += "ACCESS_HOST_WRITE_BIT|"
			case ACCESS_MEMORY_READ_BIT:
				s += "ACCESS_MEMORY_READ_BIT|"
			case ACCESS_MEMORY_WRITE_BIT:
				s += "ACCESS_MEMORY_WRITE_BIT|"
			case ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT:
				s += "ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT|"
			case ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT:
				s += "ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT|"
			case ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT:
				s += "ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT|"
			case ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT:
				s += "ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT|"
			case ACCESS_COMMAND_PROCESS_READ_BIT_NVX:
				s += "ACCESS_COMMAND_PROCESS_READ_BIT_NVX|"
			case ACCESS_COMMAND_PROCESS_WRITE_BIT_NVX:
				s += "ACCESS_COMMAND_PROCESS_WRITE_BIT_NVX|"
			case ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT:
				s += "ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT|"
			case ACCESS_SHADING_RATE_IMAGE_READ_BIT_NV:
				s += "ACCESS_SHADING_RATE_IMAGE_READ_BIT_NV|"
			case ACCESS_ACCELERATION_STRUCTURE_READ_BIT_NV:
				s += "ACCESS_ACCELERATION_STRUCTURE_READ_BIT_NV|"
			case ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_NV:
				s += "ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_NV|"
			case ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT:
				s += "ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DependencyFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDependencyFlags.html
type DependencyFlags uint32

const (
	DEPENDENCY_BY_REGION_BIT        DependencyFlags = 0x00000001
	DEPENDENCY_DEVICE_GROUP_BIT     DependencyFlags = 0x00000004
	DEPENDENCY_VIEW_LOCAL_BIT       DependencyFlags = 0x00000002
	DEPENDENCY_VIEW_LOCAL_BIT_KHR   DependencyFlags = DEPENDENCY_VIEW_LOCAL_BIT
	DEPENDENCY_DEVICE_GROUP_BIT_KHR DependencyFlags = DEPENDENCY_DEVICE_GROUP_BIT
	DEPENDENCY_FLAG_BITS_MAX_ENUM   DependencyFlags = 0x7FFFFFFF
)

func (x DependencyFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DependencyFlags(1 << i) {
			case DEPENDENCY_BY_REGION_BIT:
				s += "DEPENDENCY_BY_REGION_BIT|"
			case DEPENDENCY_DEVICE_GROUP_BIT:
				s += "DEPENDENCY_DEVICE_GROUP_BIT|"
			case DEPENDENCY_VIEW_LOCAL_BIT:
				s += "DEPENDENCY_VIEW_LOCAL_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// CommandPoolCreateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandPoolCreateFlags.html
type CommandPoolCreateFlags uint32

const (
	COMMAND_POOL_CREATE_TRANSIENT_BIT            CommandPoolCreateFlags = 0x00000001
	COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT CommandPoolCreateFlags = 0x00000002
	COMMAND_POOL_CREATE_PROTECTED_BIT            CommandPoolCreateFlags = 0x00000004
	COMMAND_POOL_CREATE_FLAG_BITS_MAX_ENUM       CommandPoolCreateFlags = 0x7FFFFFFF
)

func (x CommandPoolCreateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch CommandPoolCreateFlags(1 << i) {
			case COMMAND_POOL_CREATE_TRANSIENT_BIT:
				s += "COMMAND_POOL_CREATE_TRANSIENT_BIT|"
			case COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT:
				s += "COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT|"
			case COMMAND_POOL_CREATE_PROTECTED_BIT:
				s += "COMMAND_POOL_CREATE_PROTECTED_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// CommandPoolResetFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandPoolResetFlags.html
type CommandPoolResetFlags uint32

const (
	COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT CommandPoolResetFlags = 0x00000001
	COMMAND_POOL_RESET_FLAG_BITS_MAX_ENUM    CommandPoolResetFlags = 0x7FFFFFFF
)

func (x CommandPoolResetFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch CommandPoolResetFlags(1 << i) {
			case COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT:
				s += "COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// CommandBufferUsageFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBufferUsageFlags.html
type CommandBufferUsageFlags uint32

const (
	COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT      CommandBufferUsageFlags = 0x00000001
	COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT CommandBufferUsageFlags = 0x00000002
	COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT     CommandBufferUsageFlags = 0x00000004
	COMMAND_BUFFER_USAGE_FLAG_BITS_MAX_ENUM       CommandBufferUsageFlags = 0x7FFFFFFF
)

func (x CommandBufferUsageFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch CommandBufferUsageFlags(1 << i) {
			case COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT:
				s += "COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT|"
			case COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT:
				s += "COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT|"
			case COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT:
				s += "COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// QueryControlFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryControlFlags.html
type QueryControlFlags uint32

const (
	QUERY_CONTROL_PRECISE_BIT        QueryControlFlags = 0x00000001
	QUERY_CONTROL_FLAG_BITS_MAX_ENUM QueryControlFlags = 0x7FFFFFFF
)

func (x QueryControlFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch QueryControlFlags(1 << i) {
			case QUERY_CONTROL_PRECISE_BIT:
				s += "QUERY_CONTROL_PRECISE_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// CommandBufferResetFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBufferResetFlags.html
type CommandBufferResetFlags uint32

const (
	COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT CommandBufferResetFlags = 0x00000001
	COMMAND_BUFFER_RESET_FLAG_BITS_MAX_ENUM    CommandBufferResetFlags = 0x7FFFFFFF
)

func (x CommandBufferResetFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch CommandBufferResetFlags(1 << i) {
			case COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT:
				s += "COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// StencilFaceFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkStencilFaceFlags.html
type StencilFaceFlags uint32

const (
	STENCIL_FACE_FRONT_BIT          StencilFaceFlags = 0x00000001
	STENCIL_FACE_BACK_BIT           StencilFaceFlags = 0x00000002
	STENCIL_FACE_FRONT_AND_BACK     StencilFaceFlags = 0x00000003
	STENCIL_FRONT_AND_BACK          StencilFaceFlags = STENCIL_FACE_FRONT_AND_BACK
	STENCIL_FACE_FLAG_BITS_MAX_ENUM StencilFaceFlags = 0x7FFFFFFF
)

func (x StencilFaceFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch StencilFaceFlags(1 << i) {
			case STENCIL_FACE_FRONT_BIT:
				s += "STENCIL_FACE_FRONT_BIT|"
			case STENCIL_FACE_BACK_BIT:
				s += "STENCIL_FACE_BACK_BIT|"
			case STENCIL_FACE_FRONT_AND_BACK:
				s += "STENCIL_FACE_FRONT_AND_BACK|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ApplicationInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkApplicationInfo.html
type ApplicationInfo struct {
	SType              StructureType
	PNext              unsafe.Pointer
	PApplicationName   *int8
	ApplicationVersion uint32
	PEngineName        *int8
	EngineVersion      uint32
	ApiVersion         uint32
}

func NewApplicationInfo() *ApplicationInfo {
	p := (*ApplicationInfo)(MemAlloc(unsafe.Sizeof(*(*ApplicationInfo)(nil))))
	p.SType = STRUCTURE_TYPE_APPLICATION_INFO
	return p
}
func (p *ApplicationInfo) Free() { MemFree(unsafe.Pointer(p)) }

// InstanceCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkInstanceCreateInfo.html
type InstanceCreateInfo struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	Flags                   InstanceCreateFlags
	PApplicationInfo        *ApplicationInfo
	EnabledLayerCount       uint32
	PpEnabledLayerNames     **int8
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames **int8
}

func NewInstanceCreateInfo() *InstanceCreateInfo {
	p := (*InstanceCreateInfo)(MemAlloc(unsafe.Sizeof(*(*InstanceCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	return p
}
func (p *InstanceCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnAllocationFunction -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAllocationFunction.html
type PfnAllocationFunction uintptr

func (fn PfnAllocationFunction) Call(pUserData unsafe.Pointer, size, alignment uintptr, allocationScope SystemAllocationScope) unsafe.Pointer {
	ret, _, _ := call(uintptr(fn), uintptr(pUserData), uintptr(size), uintptr(alignment), uintptr(allocationScope))
	return unsafe.Pointer(ret)
}
func (fn PfnAllocationFunction) String() string { return "vkAllocationFunction" }

//  PfnReallocationFunction -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkReallocationFunction.html
type PfnReallocationFunction uintptr

func (fn PfnReallocationFunction) Call(pUserData, pOriginal unsafe.Pointer, size, alignment uintptr, allocationScope SystemAllocationScope) unsafe.Pointer {
	ret, _, _ := call(uintptr(fn), uintptr(pUserData), uintptr(pOriginal), uintptr(size), uintptr(alignment), uintptr(allocationScope))
	return unsafe.Pointer(ret)
}
func (fn PfnReallocationFunction) String() string { return "vkReallocationFunction" }

//  PfnFreeFunction -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkFreeFunction.html
type PfnFreeFunction uintptr

func (fn PfnFreeFunction) Call(pUserData, pMemory unsafe.Pointer) {
	_, _, _ = call(uintptr(fn), uintptr(pUserData), uintptr(pMemory))
}
func (fn PfnFreeFunction) String() string { return "vkFreeFunction" }

//  PfnInternalAllocationNotification -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkInternalAllocationNotification.html
type PfnInternalAllocationNotification uintptr

func (fn PfnInternalAllocationNotification) Call(pUserData unsafe.Pointer, size uintptr, allocationType InternalAllocationType, allocationScope SystemAllocationScope) {
	_, _, _ = call(uintptr(fn), uintptr(pUserData), uintptr(size), uintptr(allocationType), uintptr(allocationScope))
}
func (fn PfnInternalAllocationNotification) String() string {
	return "vkInternalAllocationNotification"
}

//  PfnInternalFreeNotification -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkInternalFreeNotification.html
type PfnInternalFreeNotification uintptr

func (fn PfnInternalFreeNotification) Call(pUserData unsafe.Pointer, size uintptr, allocationType InternalAllocationType, allocationScope SystemAllocationScope) {
	_, _, _ = call(uintptr(fn), uintptr(pUserData), uintptr(size), uintptr(allocationType), uintptr(allocationScope))
}
func (fn PfnInternalFreeNotification) String() string { return "vkInternalFreeNotification" }

// AllocationCallbacks -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAllocationCallbacks.html
type AllocationCallbacks struct {
	PUserData             unsafe.Pointer
	PfnAllocation         PfnAllocationFunction
	PfnReallocation       PfnReallocationFunction
	PfnFree               PfnFreeFunction
	PfnInternalAllocation PfnInternalAllocationNotification
	PfnInternalFree       PfnInternalFreeNotification
}

func NewAllocationCallbacks() *AllocationCallbacks {
	return (*AllocationCallbacks)(MemAlloc(unsafe.Sizeof(*(*AllocationCallbacks)(nil))))
}
func (p *AllocationCallbacks) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceFeatures.html
type PhysicalDeviceFeatures struct {
	RobustBufferAccess                      Bool32
	FullDrawIndexUint32                     Bool32
	ImageCubeArray                          Bool32
	IndependentBlend                        Bool32
	GeometryShader                          Bool32
	TessellationShader                      Bool32
	SampleRateShading                       Bool32
	DualSrcBlend                            Bool32
	LogicOp                                 Bool32
	MultiDrawIndirect                       Bool32
	DrawIndirectFirstInstance               Bool32
	DepthClamp                              Bool32
	DepthBiasClamp                          Bool32
	FillModeNonSolid                        Bool32
	DepthBounds                             Bool32
	WideLines                               Bool32
	LargePoints                             Bool32
	AlphaToOne                              Bool32
	MultiViewport                           Bool32
	SamplerAnisotropy                       Bool32
	TextureCompressionETC2                  Bool32
	TextureCompressionASTC_LDR              Bool32
	TextureCompressionBC                    Bool32
	OcclusionQueryPrecise                   Bool32
	PipelineStatisticsQuery                 Bool32
	VertexPipelineStoresAndAtomics          Bool32
	FragmentStoresAndAtomics                Bool32
	ShaderTessellationAndGeometryPointSize  Bool32
	ShaderImageGatherExtended               Bool32
	ShaderStorageImageExtendedFormats       Bool32
	ShaderStorageImageMultisample           Bool32
	ShaderStorageImageReadWithoutFormat     Bool32
	ShaderStorageImageWriteWithoutFormat    Bool32
	ShaderUniformBufferArrayDynamicIndexing Bool32
	ShaderSampledImageArrayDynamicIndexing  Bool32
	ShaderStorageBufferArrayDynamicIndexing Bool32
	ShaderStorageImageArrayDynamicIndexing  Bool32
	ShaderClipDistance                      Bool32
	ShaderCullDistance                      Bool32
	ShaderFloat64                           Bool32
	ShaderInt64                             Bool32
	ShaderInt16                             Bool32
	ShaderResourceResidency                 Bool32
	ShaderResourceMinLod                    Bool32
	SparseBinding                           Bool32
	SparseResidencyBuffer                   Bool32
	SparseResidencyImage2D                  Bool32
	SparseResidencyImage3D                  Bool32
	SparseResidency2Samples                 Bool32
	SparseResidency4Samples                 Bool32
	SparseResidency8Samples                 Bool32
	SparseResidency16Samples                Bool32
	SparseResidencyAliased                  Bool32
	VariableMultisampleRate                 Bool32
	InheritedQueries                        Bool32
}

func NewPhysicalDeviceFeatures() *PhysicalDeviceFeatures {
	return (*PhysicalDeviceFeatures)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceFeatures)(nil))))
}
func (p *PhysicalDeviceFeatures) Free() { MemFree(unsafe.Pointer(p)) }

// FormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFormatProperties.html
type FormatProperties struct {
	LinearTilingFeatures  FormatFeatureFlags
	OptimalTilingFeatures FormatFeatureFlags
	BufferFeatures        FormatFeatureFlags
}

func NewFormatProperties() *FormatProperties {
	return (*FormatProperties)(MemAlloc(unsafe.Sizeof(*(*FormatProperties)(nil))))
}
func (p *FormatProperties) Free() { MemFree(unsafe.Pointer(p)) }

// Extent3D -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExtent3D.html
type Extent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}

func NewExtent3D() *Extent3D { return (*Extent3D)(MemAlloc(unsafe.Sizeof(*(*Extent3D)(nil)))) }
func (p *Extent3D) Free()    { MemFree(unsafe.Pointer(p)) }

// ImageFormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageFormatProperties.html
type ImageFormatProperties struct {
	MaxExtent       Extent3D
	MaxMipLevels    uint32
	MaxArrayLayers  uint32
	SampleCounts    SampleCountFlags
	MaxResourceSize DeviceSize
}

func NewImageFormatProperties() *ImageFormatProperties {
	return (*ImageFormatProperties)(MemAlloc(unsafe.Sizeof(*(*ImageFormatProperties)(nil))))
}
func (p *ImageFormatProperties) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceLimits -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceLimits.html
type PhysicalDeviceLimits struct {
	MaxImageDimension1D                             uint32
	MaxImageDimension2D                             uint32
	MaxImageDimension3D                             uint32
	MaxImageDimensionCube                           uint32
	MaxImageArrayLayers                             uint32
	MaxTexelBufferElements                          uint32
	MaxUniformBufferRange                           uint32
	MaxStorageBufferRange                           uint32
	MaxPushConstantsSize                            uint32
	MaxMemoryAllocationCount                        uint32
	MaxSamplerAllocationCount                       uint32
	BufferImageGranularity                          DeviceSize
	SparseAddressSpaceSize                          DeviceSize
	MaxBoundDescriptorSets                          uint32
	MaxPerStageDescriptorSamplers                   uint32
	MaxPerStageDescriptorUniformBuffers             uint32
	MaxPerStageDescriptorStorageBuffers             uint32
	MaxPerStageDescriptorSampledImages              uint32
	MaxPerStageDescriptorStorageImages              uint32
	MaxPerStageDescriptorInputAttachments           uint32
	MaxPerStageResources                            uint32
	MaxDescriptorSetSamplers                        uint32
	MaxDescriptorSetUniformBuffers                  uint32
	MaxDescriptorSetUniformBuffersDynamic           uint32
	MaxDescriptorSetStorageBuffers                  uint32
	MaxDescriptorSetStorageBuffersDynamic           uint32
	MaxDescriptorSetSampledImages                   uint32
	MaxDescriptorSetStorageImages                   uint32
	MaxDescriptorSetInputAttachments                uint32
	MaxVertexInputAttributes                        uint32
	MaxVertexInputBindings                          uint32
	MaxVertexInputAttributeOffset                   uint32
	MaxVertexInputBindingStride                     uint32
	MaxVertexOutputComponents                       uint32
	MaxTessellationGenerationLevel                  uint32
	MaxTessellationPatchSize                        uint32
	MaxTessellationControlPerVertexInputComponents  uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents  uint32
	MaxTessellationControlTotalOutputComponents     uint32
	MaxTessellationEvaluationInputComponents        uint32
	MaxTessellationEvaluationOutputComponents       uint32
	MaxGeometryShaderInvocations                    uint32
	MaxGeometryInputComponents                      uint32
	MaxGeometryOutputComponents                     uint32
	MaxGeometryOutputVertices                       uint32
	MaxGeometryTotalOutputComponents                uint32
	MaxFragmentInputComponents                      uint32
	MaxFragmentOutputAttachments                    uint32
	MaxFragmentDualSrcAttachments                   uint32
	MaxFragmentCombinedOutputResources              uint32
	MaxComputeSharedMemorySize                      uint32
	MaxComputeWorkGroupCount                        [3]uint32
	MaxComputeWorkGroupInvocations                  uint32
	MaxComputeWorkGroupSize                         [3]uint32
	SubPixelPrecisionBits                           uint32
	SubTexelPrecisionBits                           uint32
	MipmapPrecisionBits                             uint32
	MaxDrawIndexedIndexValue                        uint32
	MaxDrawIndirectCount                            uint32
	MaxSamplerLodBias                               float32
	MaxSamplerAnisotropy                            float32
	MaxViewports                                    uint32
	MaxViewportDimensions                           [2]uint32
	ViewportBoundsRange                             [2]float32
	ViewportSubPixelBits                            uint32
	MinMemoryMapAlignment                           uintptr
	MinTexelBufferOffsetAlignment                   DeviceSize
	MinUniformBufferOffsetAlignment                 DeviceSize
	MinStorageBufferOffsetAlignment                 DeviceSize
	MinTexelOffset                                  int32
	MaxTexelOffset                                  uint32
	MinTexelGatherOffset                            int32
	MaxTexelGatherOffset                            uint32
	MinInterpolationOffset                          float32
	MaxInterpolationOffset                          float32
	SubPixelInterpolationOffsetBits                 uint32
	MaxFramebufferWidth                             uint32
	MaxFramebufferHeight                            uint32
	MaxFramebufferLayers                            uint32
	FramebufferColorSampleCounts                    SampleCountFlags
	FramebufferDepthSampleCounts                    SampleCountFlags
	FramebufferStencilSampleCounts                  SampleCountFlags
	FramebufferNoAttachmentsSampleCounts            SampleCountFlags
	MaxColorAttachments                             uint32
	SampledImageColorSampleCounts                   SampleCountFlags
	SampledImageIntegerSampleCounts                 SampleCountFlags
	SampledImageDepthSampleCounts                   SampleCountFlags
	SampledImageStencilSampleCounts                 SampleCountFlags
	StorageImageSampleCounts                        SampleCountFlags
	MaxSampleMaskWords                              uint32
	TimestampComputeAndGraphics                     Bool32
	TimestampPeriod                                 float32
	MaxClipDistances                                uint32
	MaxCullDistances                                uint32
	MaxCombinedClipAndCullDistances                 uint32
	DiscreteQueuePriorities                         uint32
	PointSizeRange                                  [2]float32
	LineWidthRange                                  [2]float32
	PointSizeGranularity                            float32
	LineWidthGranularity                            float32
	StrictLines                                     Bool32
	StandardSampleLocations                         Bool32
	OptimalBufferCopyOffsetAlignment                DeviceSize
	OptimalBufferCopyRowPitchAlignment              DeviceSize
	NonCoherentAtomSize                             DeviceSize
}

func NewPhysicalDeviceLimits() *PhysicalDeviceLimits {
	return (*PhysicalDeviceLimits)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceLimits)(nil))))
}
func (p *PhysicalDeviceLimits) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceSparseProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSparseProperties.html
type PhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            Bool32
	ResidencyStandard2DMultisampleBlockShape Bool32
	ResidencyStandard3DBlockShape            Bool32
	ResidencyAlignedMipSize                  Bool32
	ResidencyNonResidentStrict               Bool32
}

func NewPhysicalDeviceSparseProperties() *PhysicalDeviceSparseProperties {
	return (*PhysicalDeviceSparseProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSparseProperties)(nil))))
}
func (p *PhysicalDeviceSparseProperties) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceProperties.html
type PhysicalDeviceProperties struct {
	ApiVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        PhysicalDeviceType
	DeviceName        [MAX_PHYSICAL_DEVICE_NAME_SIZE]int8
	PipelineCacheUUID [UUID_SIZE]uint8
	Limits            PhysicalDeviceLimits
	SparseProperties  PhysicalDeviceSparseProperties
}

func NewPhysicalDeviceProperties() *PhysicalDeviceProperties {
	return (*PhysicalDeviceProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceProperties)(nil))))
}
func (p *PhysicalDeviceProperties) Free() { MemFree(unsafe.Pointer(p)) }

// QueueFamilyProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueueFamilyProperties.html
type QueueFamilyProperties struct {
	QueueFlags                  QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}

func NewQueueFamilyProperties() *QueueFamilyProperties {
	return (*QueueFamilyProperties)(MemAlloc(unsafe.Sizeof(*(*QueueFamilyProperties)(nil))))
}
func (p *QueueFamilyProperties) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryType.html
type MemoryType struct {
	PropertyFlags MemoryPropertyFlags
	HeapIndex     uint32
}

func NewMemoryType() *MemoryType { return (*MemoryType)(MemAlloc(unsafe.Sizeof(*(*MemoryType)(nil)))) }
func (p *MemoryType) Free()      { MemFree(unsafe.Pointer(p)) }

// MemoryHeap -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryHeap.html
type MemoryHeap struct {
	Size  DeviceSize
	Flags MemoryHeapFlags
}

func NewMemoryHeap() *MemoryHeap { return (*MemoryHeap)(MemAlloc(unsafe.Sizeof(*(*MemoryHeap)(nil)))) }
func (p *MemoryHeap) Free()      { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceMemoryProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMemoryProperties.html
type PhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [MAX_MEMORY_TYPES]MemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [MAX_MEMORY_HEAPS]MemoryHeap
}

func NewPhysicalDeviceMemoryProperties() *PhysicalDeviceMemoryProperties {
	return (*PhysicalDeviceMemoryProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMemoryProperties)(nil))))
}
func (p *PhysicalDeviceMemoryProperties) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceQueueCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceQueueCreateInfo.html
type DeviceQueueCreateInfo struct {
	SType            StructureType
	PNext            unsafe.Pointer
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueCount       uint32
	PQueuePriorities *float32
}

func NewDeviceQueueCreateInfo() *DeviceQueueCreateInfo {
	p := (*DeviceQueueCreateInfo)(MemAlloc(unsafe.Sizeof(*(*DeviceQueueCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
	return p
}
func (p *DeviceQueueCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceCreateInfo.html
type DeviceCreateInfo struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	Flags                   DeviceCreateFlags
	QueueCreateInfoCount    uint32
	PQueueCreateInfos       *DeviceQueueCreateInfo
	EnabledLayerCount       uint32
	PpEnabledLayerNames     **int8
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames **int8
	PEnabledFeatures        *PhysicalDeviceFeatures
}

func NewDeviceCreateInfo() *DeviceCreateInfo {
	p := (*DeviceCreateInfo)(MemAlloc(unsafe.Sizeof(*(*DeviceCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_CREATE_INFO
	return p
}
func (p *DeviceCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExtensionProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExtensionProperties.html
type ExtensionProperties struct {
	ExtensionName [MAX_EXTENSION_NAME_SIZE]int8
	SpecVersion   uint32
}

func NewExtensionProperties() *ExtensionProperties {
	return (*ExtensionProperties)(MemAlloc(unsafe.Sizeof(*(*ExtensionProperties)(nil))))
}
func (p *ExtensionProperties) Free() { MemFree(unsafe.Pointer(p)) }

// LayerProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkLayerProperties.html
type LayerProperties struct {
	LayerName             [MAX_EXTENSION_NAME_SIZE]int8
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           [MAX_DESCRIPTION_SIZE]int8
}

func NewLayerProperties() *LayerProperties {
	return (*LayerProperties)(MemAlloc(unsafe.Sizeof(*(*LayerProperties)(nil))))
}
func (p *LayerProperties) Free() { MemFree(unsafe.Pointer(p)) }

// SubmitInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubmitInfo.html
type SubmitInfo struct {
	SType                StructureType
	PNext                unsafe.Pointer
	WaitSemaphoreCount   uint32
	PWaitSemaphores      *Semaphore
	PWaitDstStageMask    *PipelineStageFlags
	CommandBufferCount   uint32
	PCommandBuffers      *CommandBuffer
	SignalSemaphoreCount uint32
	PSignalSemaphores    *Semaphore
}

func NewSubmitInfo() *SubmitInfo {
	p := (*SubmitInfo)(MemAlloc(unsafe.Sizeof(*(*SubmitInfo)(nil))))
	p.SType = STRUCTURE_TYPE_SUBMIT_INFO
	return p
}
func (p *SubmitInfo) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryAllocateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryAllocateInfo.html
type MemoryAllocateInfo struct {
	SType           StructureType
	PNext           unsafe.Pointer
	AllocationSize  DeviceSize
	MemoryTypeIndex uint32
}

func NewMemoryAllocateInfo() *MemoryAllocateInfo {
	p := (*MemoryAllocateInfo)(MemAlloc(unsafe.Sizeof(*(*MemoryAllocateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
	return p
}
func (p *MemoryAllocateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// MappedMemoryRange -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMappedMemoryRange.html
type MappedMemoryRange struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Memory DeviceMemory
	Offset DeviceSize
	Size   DeviceSize
}

func NewMappedMemoryRange() *MappedMemoryRange {
	p := (*MappedMemoryRange)(MemAlloc(unsafe.Sizeof(*(*MappedMemoryRange)(nil))))
	p.SType = STRUCTURE_TYPE_MAPPED_MEMORY_RANGE
	return p
}
func (p *MappedMemoryRange) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryRequirements -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryRequirements.html
type MemoryRequirements struct {
	Size           DeviceSize
	Alignment      DeviceSize
	MemoryTypeBits uint32
}

func NewMemoryRequirements() *MemoryRequirements {
	return (*MemoryRequirements)(MemAlloc(unsafe.Sizeof(*(*MemoryRequirements)(nil))))
}
func (p *MemoryRequirements) Free() { MemFree(unsafe.Pointer(p)) }

// SparseImageFormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageFormatProperties.html
type SparseImageFormatProperties struct {
	AspectMask       ImageAspectFlags
	ImageGranularity Extent3D
	Flags            SparseImageFormatFlags
}

func NewSparseImageFormatProperties() *SparseImageFormatProperties {
	return (*SparseImageFormatProperties)(MemAlloc(unsafe.Sizeof(*(*SparseImageFormatProperties)(nil))))
}
func (p *SparseImageFormatProperties) Free() { MemFree(unsafe.Pointer(p)) }

// SparseImageMemoryRequirements -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageMemoryRequirements.html
type SparseImageMemoryRequirements struct {
	FormatProperties     SparseImageFormatProperties
	ImageMipTailFirstLod uint32
	ImageMipTailSize     DeviceSize
	ImageMipTailOffset   DeviceSize
	ImageMipTailStride   DeviceSize
}

func NewSparseImageMemoryRequirements() *SparseImageMemoryRequirements {
	return (*SparseImageMemoryRequirements)(MemAlloc(unsafe.Sizeof(*(*SparseImageMemoryRequirements)(nil))))
}
func (p *SparseImageMemoryRequirements) Free() { MemFree(unsafe.Pointer(p)) }

// SparseMemoryBind -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseMemoryBind.html
type SparseMemoryBind struct {
	ResourceOffset DeviceSize
	Size           DeviceSize
	Memory         DeviceMemory
	MemoryOffset   DeviceSize
	Flags          SparseMemoryBindFlags
}

func NewSparseMemoryBind() *SparseMemoryBind {
	return (*SparseMemoryBind)(MemAlloc(unsafe.Sizeof(*(*SparseMemoryBind)(nil))))
}
func (p *SparseMemoryBind) Free() { MemFree(unsafe.Pointer(p)) }

// SparseBufferMemoryBindInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseBufferMemoryBindInfo.html
type SparseBufferMemoryBindInfo struct {
	Buffer    Buffer
	BindCount uint32
	PBinds    *SparseMemoryBind
}

func NewSparseBufferMemoryBindInfo() *SparseBufferMemoryBindInfo {
	return (*SparseBufferMemoryBindInfo)(MemAlloc(unsafe.Sizeof(*(*SparseBufferMemoryBindInfo)(nil))))
}
func (p *SparseBufferMemoryBindInfo) Free() { MemFree(unsafe.Pointer(p)) }

// SparseImageOpaqueMemoryBindInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageOpaqueMemoryBindInfo.html
type SparseImageOpaqueMemoryBindInfo struct {
	Image     Image
	BindCount uint32
	PBinds    *SparseMemoryBind
}

func NewSparseImageOpaqueMemoryBindInfo() *SparseImageOpaqueMemoryBindInfo {
	return (*SparseImageOpaqueMemoryBindInfo)(MemAlloc(unsafe.Sizeof(*(*SparseImageOpaqueMemoryBindInfo)(nil))))
}
func (p *SparseImageOpaqueMemoryBindInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ImageSubresource -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageSubresource.html
type ImageSubresource struct {
	AspectMask ImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}

func NewImageSubresource() *ImageSubresource {
	return (*ImageSubresource)(MemAlloc(unsafe.Sizeof(*(*ImageSubresource)(nil))))
}
func (p *ImageSubresource) Free() { MemFree(unsafe.Pointer(p)) }

// Offset3D -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkOffset3D.html
type Offset3D struct {
	X int32
	Y int32
	Z int32
}

func NewOffset3D() *Offset3D { return (*Offset3D)(MemAlloc(unsafe.Sizeof(*(*Offset3D)(nil)))) }
func (p *Offset3D) Free()    { MemFree(unsafe.Pointer(p)) }

// SparseImageMemoryBind -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageMemoryBind.html
type SparseImageMemoryBind struct {
	Subresource  ImageSubresource
	Offset       Offset3D
	Extent       Extent3D
	Memory       DeviceMemory
	MemoryOffset DeviceSize
	Flags        SparseMemoryBindFlags
}

func NewSparseImageMemoryBind() *SparseImageMemoryBind {
	return (*SparseImageMemoryBind)(MemAlloc(unsafe.Sizeof(*(*SparseImageMemoryBind)(nil))))
}
func (p *SparseImageMemoryBind) Free() { MemFree(unsafe.Pointer(p)) }

// SparseImageMemoryBindInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageMemoryBindInfo.html
type SparseImageMemoryBindInfo struct {
	Image     Image
	BindCount uint32
	PBinds    *SparseImageMemoryBind
}

func NewSparseImageMemoryBindInfo() *SparseImageMemoryBindInfo {
	return (*SparseImageMemoryBindInfo)(MemAlloc(unsafe.Sizeof(*(*SparseImageMemoryBindInfo)(nil))))
}
func (p *SparseImageMemoryBindInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BindSparseInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindSparseInfo.html
type BindSparseInfo struct {
	SType                StructureType
	PNext                unsafe.Pointer
	WaitSemaphoreCount   uint32
	PWaitSemaphores      *Semaphore
	BufferBindCount      uint32
	PBufferBinds         *SparseBufferMemoryBindInfo
	ImageOpaqueBindCount uint32
	PImageOpaqueBinds    *SparseImageOpaqueMemoryBindInfo
	ImageBindCount       uint32
	PImageBinds          *SparseImageMemoryBindInfo
	SignalSemaphoreCount uint32
	PSignalSemaphores    *Semaphore
}

func NewBindSparseInfo() *BindSparseInfo {
	p := (*BindSparseInfo)(MemAlloc(unsafe.Sizeof(*(*BindSparseInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_SPARSE_INFO
	return p
}
func (p *BindSparseInfo) Free() { MemFree(unsafe.Pointer(p)) }

// FenceCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFenceCreateInfo.html
type FenceCreateInfo struct {
	SType StructureType
	PNext unsafe.Pointer
	Flags FenceCreateFlags
}

func NewFenceCreateInfo() *FenceCreateInfo {
	p := (*FenceCreateInfo)(MemAlloc(unsafe.Sizeof(*(*FenceCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_FENCE_CREATE_INFO
	return p
}
func (p *FenceCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// SemaphoreCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSemaphoreCreateInfo.html
type SemaphoreCreateInfo struct {
	SType StructureType
	PNext unsafe.Pointer
	Flags SemaphoreCreateFlags
}

func NewSemaphoreCreateInfo() *SemaphoreCreateInfo {
	p := (*SemaphoreCreateInfo)(MemAlloc(unsafe.Sizeof(*(*SemaphoreCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
	return p
}
func (p *SemaphoreCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// EventCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkEventCreateInfo.html
type EventCreateInfo struct {
	SType StructureType
	PNext unsafe.Pointer
	Flags EventCreateFlags
}

func NewEventCreateInfo() *EventCreateInfo {
	p := (*EventCreateInfo)(MemAlloc(unsafe.Sizeof(*(*EventCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_EVENT_CREATE_INFO
	return p
}
func (p *EventCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// QueryPoolCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryPoolCreateInfo.html
type QueryPoolCreateInfo struct {
	SType              StructureType
	PNext              unsafe.Pointer
	Flags              QueryPoolCreateFlags
	QueryType          QueryType
	QueryCount         uint32
	PipelineStatistics QueryPipelineStatisticFlags
}

func NewQueryPoolCreateInfo() *QueryPoolCreateInfo {
	p := (*QueryPoolCreateInfo)(MemAlloc(unsafe.Sizeof(*(*QueryPoolCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO
	return p
}
func (p *QueryPoolCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BufferCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferCreateInfo.html
type BufferCreateInfo struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Flags                 BufferCreateFlags
	Size                  DeviceSize
	Usage                 BufferUsageFlags
	SharingMode           SharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}

func NewBufferCreateInfo() *BufferCreateInfo {
	p := (*BufferCreateInfo)(MemAlloc(unsafe.Sizeof(*(*BufferCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BUFFER_CREATE_INFO
	return p
}
func (p *BufferCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BufferViewCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferViewCreateInfo.html
type BufferViewCreateInfo struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Flags  BufferViewCreateFlags
	Buffer Buffer
	Format Format
	Offset DeviceSize
	Range  DeviceSize
}

func NewBufferViewCreateInfo() *BufferViewCreateInfo {
	p := (*BufferViewCreateInfo)(MemAlloc(unsafe.Sizeof(*(*BufferViewCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
	return p
}
func (p *BufferViewCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ImageCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageCreateInfo.html
type ImageCreateInfo struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Flags                 ImageCreateFlags
	ImageType             ImageType
	Format                Format
	Extent                Extent3D
	MipLevels             uint32
	ArrayLayers           uint32
	Samples               SampleCountFlags
	Tiling                ImageTiling
	Usage                 ImageUsageFlags
	SharingMode           SharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
	InitialLayout         ImageLayout
}

func NewImageCreateInfo() *ImageCreateInfo {
	p := (*ImageCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ImageCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_CREATE_INFO
	return p
}
func (p *ImageCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// SubresourceLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubresourceLayout.html
type SubresourceLayout struct {
	Offset     DeviceSize
	Size       DeviceSize
	RowPitch   DeviceSize
	ArrayPitch DeviceSize
	DepthPitch DeviceSize
}

func NewSubresourceLayout() *SubresourceLayout {
	return (*SubresourceLayout)(MemAlloc(unsafe.Sizeof(*(*SubresourceLayout)(nil))))
}
func (p *SubresourceLayout) Free() { MemFree(unsafe.Pointer(p)) }

// ComponentMapping -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkComponentMapping.html
type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

func NewComponentMapping() *ComponentMapping {
	return (*ComponentMapping)(MemAlloc(unsafe.Sizeof(*(*ComponentMapping)(nil))))
}
func (p *ComponentMapping) Free() { MemFree(unsafe.Pointer(p)) }

// ImageSubresourceRange -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageSubresourceRange.html
type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func NewImageSubresourceRange() *ImageSubresourceRange {
	return (*ImageSubresourceRange)(MemAlloc(unsafe.Sizeof(*(*ImageSubresourceRange)(nil))))
}
func (p *ImageSubresourceRange) Free() { MemFree(unsafe.Pointer(p)) }

// ImageViewCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageViewCreateInfo.html
type ImageViewCreateInfo struct {
	SType            StructureType
	PNext            unsafe.Pointer
	Flags            ImageViewCreateFlags
	Image            Image
	ViewType         ImageViewType
	Format           Format
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange
}

func NewImageViewCreateInfo() *ImageViewCreateInfo {
	p := (*ImageViewCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ImageViewCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
	return p
}
func (p *ImageViewCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ShaderModuleCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderModuleCreateInfo.html
type ShaderModuleCreateInfo struct {
	SType    StructureType
	PNext    unsafe.Pointer
	Flags    ShaderModuleCreateFlags
	CodeSize uintptr
	PCode    *uint32
}

func NewShaderModuleCreateInfo() *ShaderModuleCreateInfo {
	p := (*ShaderModuleCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ShaderModuleCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO
	return p
}
func (p *ShaderModuleCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineCacheCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCacheCreateInfo.html
type PipelineCacheCreateInfo struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           PipelineCacheCreateFlags
	InitialDataSize uintptr
	PInitialData    unsafe.Pointer
}

func NewPipelineCacheCreateInfo() *PipelineCacheCreateInfo {
	p := (*PipelineCacheCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineCacheCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_CACHE_CREATE_INFO
	return p
}
func (p *PipelineCacheCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// SpecializationMapEntry -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSpecializationMapEntry.html
type SpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       uintptr
}

func NewSpecializationMapEntry() *SpecializationMapEntry {
	return (*SpecializationMapEntry)(MemAlloc(unsafe.Sizeof(*(*SpecializationMapEntry)(nil))))
}
func (p *SpecializationMapEntry) Free() { MemFree(unsafe.Pointer(p)) }

// SpecializationInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSpecializationInfo.html
type SpecializationInfo struct {
	MapEntryCount uint32
	PMapEntries   *SpecializationMapEntry
	DataSize      uintptr
	PData         unsafe.Pointer
}

func NewSpecializationInfo() *SpecializationInfo {
	return (*SpecializationInfo)(MemAlloc(unsafe.Sizeof(*(*SpecializationInfo)(nil))))
}
func (p *SpecializationInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineShaderStageCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineShaderStageCreateInfo.html
type PipelineShaderStageCreateInfo struct {
	SType               StructureType
	PNext               unsafe.Pointer
	Flags               PipelineShaderStageCreateFlags
	Stage               ShaderStageFlags
	Module              ShaderModule
	PName               *int8
	PSpecializationInfo *SpecializationInfo
}

func NewPipelineShaderStageCreateInfo() *PipelineShaderStageCreateInfo {
	p := (*PipelineShaderStageCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineShaderStageCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO
	return p
}
func (p *PipelineShaderStageCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// VertexInputBindingDescription -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkVertexInputBindingDescription.html
type VertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate
}

func NewVertexInputBindingDescription() *VertexInputBindingDescription {
	return (*VertexInputBindingDescription)(MemAlloc(unsafe.Sizeof(*(*VertexInputBindingDescription)(nil))))
}
func (p *VertexInputBindingDescription) Free() { MemFree(unsafe.Pointer(p)) }

// VertexInputAttributeDescription -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkVertexInputAttributeDescription.html
type VertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32
}

func NewVertexInputAttributeDescription() *VertexInputAttributeDescription {
	return (*VertexInputAttributeDescription)(MemAlloc(unsafe.Sizeof(*(*VertexInputAttributeDescription)(nil))))
}
func (p *VertexInputAttributeDescription) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineVertexInputStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html
type PipelineVertexInputStateCreateInfo struct {
	SType                           StructureType
	PNext                           unsafe.Pointer
	Flags                           PipelineVertexInputStateCreateFlags
	VertexBindingDescriptionCount   uint32
	PVertexBindingDescriptions      *VertexInputBindingDescription
	VertexAttributeDescriptionCount uint32
	PVertexAttributeDescriptions    *VertexInputAttributeDescription
}

func NewPipelineVertexInputStateCreateInfo() *PipelineVertexInputStateCreateInfo {
	p := (*PipelineVertexInputStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineVertexInputStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO
	return p
}
func (p *PipelineVertexInputStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineInputAssemblyStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html
type PipelineInputAssemblyStateCreateInfo struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	Flags                  PipelineInputAssemblyStateCreateFlags
	Topology               PrimitiveTopology
	PrimitiveRestartEnable Bool32
}

func NewPipelineInputAssemblyStateCreateInfo() *PipelineInputAssemblyStateCreateInfo {
	p := (*PipelineInputAssemblyStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineInputAssemblyStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_INPUT_ASSEMBLY_STATE_CREATE_INFO
	return p
}
func (p *PipelineInputAssemblyStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineTessellationStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineTessellationStateCreateInfo.html
type PipelineTessellationStateCreateInfo struct {
	SType              StructureType
	PNext              unsafe.Pointer
	Flags              PipelineTessellationStateCreateFlags
	PatchControlPoints uint32
}

func NewPipelineTessellationStateCreateInfo() *PipelineTessellationStateCreateInfo {
	p := (*PipelineTessellationStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineTessellationStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO
	return p
}
func (p *PipelineTessellationStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// Viewport -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkViewport.html
type Viewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

func NewViewport() *Viewport { return (*Viewport)(MemAlloc(unsafe.Sizeof(*(*Viewport)(nil)))) }
func (p *Viewport) Free()    { MemFree(unsafe.Pointer(p)) }

// Offset2D -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkOffset2D.html
type Offset2D struct {
	X int32
	Y int32
}

func NewOffset2D() *Offset2D { return (*Offset2D)(MemAlloc(unsafe.Sizeof(*(*Offset2D)(nil)))) }
func (p *Offset2D) Free()    { MemFree(unsafe.Pointer(p)) }

// Extent2D -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExtent2D.html
type Extent2D struct {
	Width  uint32
	Height uint32
}

func NewExtent2D() *Extent2D { return (*Extent2D)(MemAlloc(unsafe.Sizeof(*(*Extent2D)(nil)))) }
func (p *Extent2D) Free()    { MemFree(unsafe.Pointer(p)) }

// Rect2D -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRect2D.html
type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}

func NewRect2D() *Rect2D { return (*Rect2D)(MemAlloc(unsafe.Sizeof(*(*Rect2D)(nil)))) }
func (p *Rect2D) Free()  { MemFree(unsafe.Pointer(p)) }

// PipelineViewportStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineViewportStateCreateInfo.html
type PipelineViewportStateCreateInfo struct {
	SType         StructureType
	PNext         unsafe.Pointer
	Flags         PipelineViewportStateCreateFlags
	ViewportCount uint32
	PViewports    *Viewport
	ScissorCount  uint32
	PScissors     *Rect2D
}

func NewPipelineViewportStateCreateInfo() *PipelineViewportStateCreateInfo {
	p := (*PipelineViewportStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineViewportStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO
	return p
}
func (p *PipelineViewportStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineRasterizationStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html
type PipelineRasterizationStateCreateInfo struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	Flags                   PipelineRasterizationStateCreateFlags
	DepthClampEnable        Bool32
	RasterizerDiscardEnable Bool32
	PolygonMode             PolygonMode
	CullMode                CullModeFlags
	FrontFace               FrontFace
	DepthBiasEnable         Bool32
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}

func NewPipelineRasterizationStateCreateInfo() *PipelineRasterizationStateCreateInfo {
	p := (*PipelineRasterizationStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineRasterizationStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO
	return p
}
func (p *PipelineRasterizationStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineMultisampleStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html
type PipelineMultisampleStateCreateInfo struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Flags                 PipelineMultisampleStateCreateFlags
	RasterizationSamples  SampleCountFlags
	SampleShadingEnable   Bool32
	MinSampleShading      float32
	PSampleMask           *SampleMask
	AlphaToCoverageEnable Bool32
	AlphaToOneEnable      Bool32
}

func NewPipelineMultisampleStateCreateInfo() *PipelineMultisampleStateCreateInfo {
	p := (*PipelineMultisampleStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineMultisampleStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO
	return p
}
func (p *PipelineMultisampleStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// StencilOpState -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkStencilOpState.html
type StencilOpState struct {
	FailOp      StencilOp
	PassOp      StencilOp
	DepthFailOp StencilOp
	CompareOp   CompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}

func NewStencilOpState() *StencilOpState {
	return (*StencilOpState)(MemAlloc(unsafe.Sizeof(*(*StencilOpState)(nil))))
}
func (p *StencilOpState) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineDepthStencilStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html
type PipelineDepthStencilStateCreateInfo struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Flags                 PipelineDepthStencilStateCreateFlags
	DepthTestEnable       Bool32
	DepthWriteEnable      Bool32
	DepthCompareOp        CompareOp
	DepthBoundsTestEnable Bool32
	StencilTestEnable     Bool32
	Front                 StencilOpState
	Back                  StencilOpState
	MinDepthBounds        float32
	MaxDepthBounds        float32
}

func NewPipelineDepthStencilStateCreateInfo() *PipelineDepthStencilStateCreateInfo {
	p := (*PipelineDepthStencilStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineDepthStencilStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO
	return p
}
func (p *PipelineDepthStencilStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineColorBlendAttachmentState -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineColorBlendAttachmentState.html
type PipelineColorBlendAttachmentState struct {
	BlendEnable         Bool32
	SrcColorBlendFactor BlendFactor
	DstColorBlendFactor BlendFactor
	ColorBlendOp        BlendOp
	SrcAlphaBlendFactor BlendFactor
	DstAlphaBlendFactor BlendFactor
	AlphaBlendOp        BlendOp
	ColorWriteMask      ColorComponentFlags
}

func NewPipelineColorBlendAttachmentState() *PipelineColorBlendAttachmentState {
	return (*PipelineColorBlendAttachmentState)(MemAlloc(unsafe.Sizeof(*(*PipelineColorBlendAttachmentState)(nil))))
}
func (p *PipelineColorBlendAttachmentState) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineColorBlendStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html
type PipelineColorBlendStateCreateInfo struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           PipelineColorBlendStateCreateFlags
	LogicOpEnable   Bool32
	LogicOp         LogicOp
	AttachmentCount uint32
	PAttachments    *PipelineColorBlendAttachmentState
	BlendConstants  [4]float32
}

func NewPipelineColorBlendStateCreateInfo() *PipelineColorBlendStateCreateInfo {
	p := (*PipelineColorBlendStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineColorBlendStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO
	return p
}
func (p *PipelineColorBlendStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineDynamicStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineDynamicStateCreateInfo.html
type PipelineDynamicStateCreateInfo struct {
	SType             StructureType
	PNext             unsafe.Pointer
	Flags             PipelineDynamicStateCreateFlags
	DynamicStateCount uint32
	PDynamicStates    *DynamicState
}

func NewPipelineDynamicStateCreateInfo() *PipelineDynamicStateCreateInfo {
	p := (*PipelineDynamicStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineDynamicStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO
	return p
}
func (p *PipelineDynamicStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// GraphicsPipelineCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGraphicsPipelineCreateInfo.html
type GraphicsPipelineCreateInfo struct {
	SType               StructureType
	PNext               unsafe.Pointer
	Flags               PipelineCreateFlags
	StageCount          uint32
	PStages             *PipelineShaderStageCreateInfo
	PVertexInputState   *PipelineVertexInputStateCreateInfo
	PInputAssemblyState *PipelineInputAssemblyStateCreateInfo
	PTessellationState  *PipelineTessellationStateCreateInfo
	PViewportState      *PipelineViewportStateCreateInfo
	PRasterizationState *PipelineRasterizationStateCreateInfo
	PMultisampleState   *PipelineMultisampleStateCreateInfo
	PDepthStencilState  *PipelineDepthStencilStateCreateInfo
	PColorBlendState    *PipelineColorBlendStateCreateInfo
	PDynamicState       *PipelineDynamicStateCreateInfo
	Layout              PipelineLayout
	RenderPass          RenderPass
	Subpass             uint32
	BasePipelineHandle  Pipeline
	BasePipelineIndex   int32
}

func NewGraphicsPipelineCreateInfo() *GraphicsPipelineCreateInfo {
	p := (*GraphicsPipelineCreateInfo)(MemAlloc(unsafe.Sizeof(*(*GraphicsPipelineCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
	return p
}
func (p *GraphicsPipelineCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ComputePipelineCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkComputePipelineCreateInfo.html
type ComputePipelineCreateInfo struct {
	SType              StructureType
	PNext              unsafe.Pointer
	Flags              PipelineCreateFlags
	Stage              PipelineShaderStageCreateInfo
	Layout             PipelineLayout
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

func NewComputePipelineCreateInfo() *ComputePipelineCreateInfo {
	p := (*ComputePipelineCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ComputePipelineCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO
	return p
}
func (p *ComputePipelineCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PushConstantRange -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPushConstantRange.html
type PushConstantRange struct {
	StageFlags ShaderStageFlags
	Offset     uint32
	Size       uint32
}

func NewPushConstantRange() *PushConstantRange {
	return (*PushConstantRange)(MemAlloc(unsafe.Sizeof(*(*PushConstantRange)(nil))))
}
func (p *PushConstantRange) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineLayoutCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineLayoutCreateInfo.html
type PipelineLayoutCreateInfo struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	Flags                  PipelineLayoutCreateFlags
	SetLayoutCount         uint32
	PSetLayouts            *DescriptorSetLayout
	PushConstantRangeCount uint32
	PPushConstantRanges    *PushConstantRange
}

func NewPipelineLayoutCreateInfo() *PipelineLayoutCreateInfo {
	p := (*PipelineLayoutCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineLayoutCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO
	return p
}
func (p *PipelineLayoutCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// SamplerCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerCreateInfo.html
type SamplerCreateInfo struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	Flags                   SamplerCreateFlags
	MagFilter               Filter
	MinFilter               Filter
	MipmapMode              SamplerMipmapMode
	AddressModeU            SamplerAddressMode
	AddressModeV            SamplerAddressMode
	AddressModeW            SamplerAddressMode
	MipLodBias              float32
	AnisotropyEnable        Bool32
	MaxAnisotropy           float32
	CompareEnable           Bool32
	CompareOp               CompareOp
	MinLod                  float32
	MaxLod                  float32
	BorderColor             BorderColor
	UnnormalizedCoordinates Bool32
}

func NewSamplerCreateInfo() *SamplerCreateInfo {
	p := (*SamplerCreateInfo)(MemAlloc(unsafe.Sizeof(*(*SamplerCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_SAMPLER_CREATE_INFO
	return p
}
func (p *SamplerCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorSetLayoutBinding -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetLayoutBinding.html
type DescriptorSetLayoutBinding struct {
	Binding            uint32
	DescriptorType     DescriptorType
	DescriptorCount    uint32
	StageFlags         ShaderStageFlags
	PImmutableSamplers *Sampler
}

func NewDescriptorSetLayoutBinding() *DescriptorSetLayoutBinding {
	return (*DescriptorSetLayoutBinding)(MemAlloc(unsafe.Sizeof(*(*DescriptorSetLayoutBinding)(nil))))
}
func (p *DescriptorSetLayoutBinding) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorSetLayoutCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html
type DescriptorSetLayoutCreateInfo struct {
	SType        StructureType
	PNext        unsafe.Pointer
	Flags        DescriptorSetLayoutCreateFlags
	BindingCount uint32
	PBindings    *DescriptorSetLayoutBinding
}

func NewDescriptorSetLayoutCreateInfo() *DescriptorSetLayoutCreateInfo {
	p := (*DescriptorSetLayoutCreateInfo)(MemAlloc(unsafe.Sizeof(*(*DescriptorSetLayoutCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
	return p
}
func (p *DescriptorSetLayoutCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorPoolSize -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorPoolSize.html
type DescriptorPoolSize struct {
	Type            DescriptorType
	DescriptorCount uint32
}

func NewDescriptorPoolSize() *DescriptorPoolSize {
	return (*DescriptorPoolSize)(MemAlloc(unsafe.Sizeof(*(*DescriptorPoolSize)(nil))))
}
func (p *DescriptorPoolSize) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorPoolCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorPoolCreateInfo.html
type DescriptorPoolCreateInfo struct {
	SType         StructureType
	PNext         unsafe.Pointer
	Flags         DescriptorPoolCreateFlags
	MaxSets       uint32
	PoolSizeCount uint32
	PPoolSizes    *DescriptorPoolSize
}

func NewDescriptorPoolCreateInfo() *DescriptorPoolCreateInfo {
	p := (*DescriptorPoolCreateInfo)(MemAlloc(unsafe.Sizeof(*(*DescriptorPoolCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO
	return p
}
func (p *DescriptorPoolCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorSetAllocateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetAllocateInfo.html
type DescriptorSetAllocateInfo struct {
	SType              StructureType
	PNext              unsafe.Pointer
	DescriptorPool     DescriptorPool
	DescriptorSetCount uint32
	PSetLayouts        *DescriptorSetLayout
}

func NewDescriptorSetAllocateInfo() *DescriptorSetAllocateInfo {
	p := (*DescriptorSetAllocateInfo)(MemAlloc(unsafe.Sizeof(*(*DescriptorSetAllocateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO
	return p
}
func (p *DescriptorSetAllocateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorImageInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorImageInfo.html
type DescriptorImageInfo struct {
	Sampler     Sampler
	ImageView   ImageView
	ImageLayout ImageLayout
}

func NewDescriptorImageInfo() *DescriptorImageInfo {
	return (*DescriptorImageInfo)(MemAlloc(unsafe.Sizeof(*(*DescriptorImageInfo)(nil))))
}
func (p *DescriptorImageInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorBufferInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorBufferInfo.html
type DescriptorBufferInfo struct {
	Buffer Buffer
	Offset DeviceSize
	Range  DeviceSize
}

func NewDescriptorBufferInfo() *DescriptorBufferInfo {
	return (*DescriptorBufferInfo)(MemAlloc(unsafe.Sizeof(*(*DescriptorBufferInfo)(nil))))
}
func (p *DescriptorBufferInfo) Free() { MemFree(unsafe.Pointer(p)) }

// WriteDescriptorSet -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkWriteDescriptorSet.html
type WriteDescriptorSet struct {
	SType            StructureType
	PNext            unsafe.Pointer
	DstSet           DescriptorSet
	DstBinding       uint32
	DstArrayElement  uint32
	DescriptorCount  uint32
	DescriptorType   DescriptorType
	PImageInfo       *DescriptorImageInfo
	PBufferInfo      *DescriptorBufferInfo
	PTexelBufferView *BufferView
}

func NewWriteDescriptorSet() *WriteDescriptorSet {
	p := (*WriteDescriptorSet)(MemAlloc(unsafe.Sizeof(*(*WriteDescriptorSet)(nil))))
	p.SType = STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET
	return p
}
func (p *WriteDescriptorSet) Free() { MemFree(unsafe.Pointer(p)) }

// CopyDescriptorSet -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCopyDescriptorSet.html
type CopyDescriptorSet struct {
	SType           StructureType
	PNext           unsafe.Pointer
	SrcSet          DescriptorSet
	SrcBinding      uint32
	SrcArrayElement uint32
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
}

func NewCopyDescriptorSet() *CopyDescriptorSet {
	p := (*CopyDescriptorSet)(MemAlloc(unsafe.Sizeof(*(*CopyDescriptorSet)(nil))))
	p.SType = STRUCTURE_TYPE_COPY_DESCRIPTOR_SET
	return p
}
func (p *CopyDescriptorSet) Free() { MemFree(unsafe.Pointer(p)) }

// FramebufferCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFramebufferCreateInfo.html
type FramebufferCreateInfo struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           FramebufferCreateFlags
	RenderPass      RenderPass
	AttachmentCount uint32
	PAttachments    *ImageView
	Width           uint32
	Height          uint32
	Layers          uint32
}

func NewFramebufferCreateInfo() *FramebufferCreateInfo {
	p := (*FramebufferCreateInfo)(MemAlloc(unsafe.Sizeof(*(*FramebufferCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO
	return p
}
func (p *FramebufferCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// AttachmentDescription -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentDescription.html
type AttachmentDescription struct {
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlags
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}

func NewAttachmentDescription() *AttachmentDescription {
	return (*AttachmentDescription)(MemAlloc(unsafe.Sizeof(*(*AttachmentDescription)(nil))))
}
func (p *AttachmentDescription) Free() { MemFree(unsafe.Pointer(p)) }

// AttachmentReference -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentReference.html
type AttachmentReference struct {
	Attachment uint32
	Layout     ImageLayout
}

func NewAttachmentReference() *AttachmentReference {
	return (*AttachmentReference)(MemAlloc(unsafe.Sizeof(*(*AttachmentReference)(nil))))
}
func (p *AttachmentReference) Free() { MemFree(unsafe.Pointer(p)) }

// SubpassDescription -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassDescription.html
type SubpassDescription struct {
	Flags                   SubpassDescriptionFlags
	PipelineBindPoint       PipelineBindPoint
	InputAttachmentCount    uint32
	PInputAttachments       *AttachmentReference
	ColorAttachmentCount    uint32
	PColorAttachments       *AttachmentReference
	PResolveAttachments     *AttachmentReference
	PDepthStencilAttachment *AttachmentReference
	PreserveAttachmentCount uint32
	PPreserveAttachments    *uint32
}

func NewSubpassDescription() *SubpassDescription {
	return (*SubpassDescription)(MemAlloc(unsafe.Sizeof(*(*SubpassDescription)(nil))))
}
func (p *SubpassDescription) Free() { MemFree(unsafe.Pointer(p)) }

// SubpassDependency -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassDependency.html
type SubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
}

func NewSubpassDependency() *SubpassDependency {
	return (*SubpassDependency)(MemAlloc(unsafe.Sizeof(*(*SubpassDependency)(nil))))
}
func (p *SubpassDependency) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassCreateInfo.html
type RenderPassCreateInfo struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           RenderPassCreateFlags
	AttachmentCount uint32
	PAttachments    *AttachmentDescription
	SubpassCount    uint32
	PSubpasses      *SubpassDescription
	DependencyCount uint32
	PDependencies   *SubpassDependency
}

func NewRenderPassCreateInfo() *RenderPassCreateInfo {
	p := (*RenderPassCreateInfo)(MemAlloc(unsafe.Sizeof(*(*RenderPassCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO
	return p
}
func (p *RenderPassCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// CommandPoolCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandPoolCreateInfo.html
type CommandPoolCreateInfo struct {
	SType            StructureType
	PNext            unsafe.Pointer
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

func NewCommandPoolCreateInfo() *CommandPoolCreateInfo {
	p := (*CommandPoolCreateInfo)(MemAlloc(unsafe.Sizeof(*(*CommandPoolCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
	return p
}
func (p *CommandPoolCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// CommandBufferAllocateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBufferAllocateInfo.html
type CommandBufferAllocateInfo struct {
	SType              StructureType
	PNext              unsafe.Pointer
	CommandPool        CommandPool
	Level              CommandBufferLevel
	CommandBufferCount uint32
}

func NewCommandBufferAllocateInfo() *CommandBufferAllocateInfo {
	p := (*CommandBufferAllocateInfo)(MemAlloc(unsafe.Sizeof(*(*CommandBufferAllocateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	return p
}
func (p *CommandBufferAllocateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// CommandBufferInheritanceInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBufferInheritanceInfo.html
type CommandBufferInheritanceInfo struct {
	SType                StructureType
	PNext                unsafe.Pointer
	RenderPass           RenderPass
	Subpass              uint32
	Framebuffer          Framebuffer
	OcclusionQueryEnable Bool32
	QueryFlags           QueryControlFlags
	PipelineStatistics   QueryPipelineStatisticFlags
}

func NewCommandBufferInheritanceInfo() *CommandBufferInheritanceInfo {
	p := (*CommandBufferInheritanceInfo)(MemAlloc(unsafe.Sizeof(*(*CommandBufferInheritanceInfo)(nil))))
	p.SType = STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
	return p
}
func (p *CommandBufferInheritanceInfo) Free() { MemFree(unsafe.Pointer(p)) }

// CommandBufferBeginInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBufferBeginInfo.html
type CommandBufferBeginInfo struct {
	SType            StructureType
	PNext            unsafe.Pointer
	Flags            CommandBufferUsageFlags
	PInheritanceInfo *CommandBufferInheritanceInfo
}

func NewCommandBufferBeginInfo() *CommandBufferBeginInfo {
	p := (*CommandBufferBeginInfo)(MemAlloc(unsafe.Sizeof(*(*CommandBufferBeginInfo)(nil))))
	p.SType = STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
	return p
}
func (p *CommandBufferBeginInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BufferCopy -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferCopy.html
type BufferCopy struct {
	SrcOffset DeviceSize
	DstOffset DeviceSize
	Size      DeviceSize
}

func NewBufferCopy() *BufferCopy { return (*BufferCopy)(MemAlloc(unsafe.Sizeof(*(*BufferCopy)(nil)))) }
func (p *BufferCopy) Free()      { MemFree(unsafe.Pointer(p)) }

// ImageSubresourceLayers -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageSubresourceLayers.html
type ImageSubresourceLayers struct {
	AspectMask     ImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

func NewImageSubresourceLayers() *ImageSubresourceLayers {
	return (*ImageSubresourceLayers)(MemAlloc(unsafe.Sizeof(*(*ImageSubresourceLayers)(nil))))
}
func (p *ImageSubresourceLayers) Free() { MemFree(unsafe.Pointer(p)) }

// ImageCopy -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageCopy.html
type ImageCopy struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}

func NewImageCopy() *ImageCopy { return (*ImageCopy)(MemAlloc(unsafe.Sizeof(*(*ImageCopy)(nil)))) }
func (p *ImageCopy) Free()     { MemFree(unsafe.Pointer(p)) }

// ImageBlit -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageBlit.html
type ImageBlit struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffsets     [2]Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffsets     [2]Offset3D
}

func NewImageBlit() *ImageBlit { return (*ImageBlit)(MemAlloc(unsafe.Sizeof(*(*ImageBlit)(nil)))) }
func (p *ImageBlit) Free()     { MemFree(unsafe.Pointer(p)) }

// BufferImageCopy -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferImageCopy.html
type BufferImageCopy struct {
	BufferOffset      DeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  ImageSubresourceLayers
	ImageOffset       Offset3D
	ImageExtent       Extent3D
}

func NewBufferImageCopy() *BufferImageCopy {
	return (*BufferImageCopy)(MemAlloc(unsafe.Sizeof(*(*BufferImageCopy)(nil))))
}
func (p *BufferImageCopy) Free() { MemFree(unsafe.Pointer(p)) }

// ClearDepthStencilValue -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkClearDepthStencilValue.html
type ClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}

func NewClearDepthStencilValue() *ClearDepthStencilValue {
	return (*ClearDepthStencilValue)(MemAlloc(unsafe.Sizeof(*(*ClearDepthStencilValue)(nil))))
}
func (p *ClearDepthStencilValue) Free() { MemFree(unsafe.Pointer(p)) }

// ClearAttachment -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkClearAttachment.html
type ClearAttachment struct {
	AspectMask      ImageAspectFlags
	ColorAttachment uint32
	ClearValue      ClearValue
}

func NewClearAttachment() *ClearAttachment {
	return (*ClearAttachment)(MemAlloc(unsafe.Sizeof(*(*ClearAttachment)(nil))))
}
func (p *ClearAttachment) Free() { MemFree(unsafe.Pointer(p)) }

// ClearRect -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkClearRect.html
type ClearRect struct {
	Rect           Rect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}

func NewClearRect() *ClearRect { return (*ClearRect)(MemAlloc(unsafe.Sizeof(*(*ClearRect)(nil)))) }
func (p *ClearRect) Free()     { MemFree(unsafe.Pointer(p)) }

// ImageResolve -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageResolve.html
type ImageResolve struct {
	SrcSubresource ImageSubresourceLayers
	SrcOffset      Offset3D
	DstSubresource ImageSubresourceLayers
	DstOffset      Offset3D
	Extent         Extent3D
}

func NewImageResolve() *ImageResolve {
	return (*ImageResolve)(MemAlloc(unsafe.Sizeof(*(*ImageResolve)(nil))))
}
func (p *ImageResolve) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryBarrier -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryBarrier.html
type MemoryBarrier struct {
	SType         StructureType
	PNext         unsafe.Pointer
	SrcAccessMask AccessFlags
	DstAccessMask AccessFlags
}

func NewMemoryBarrier() *MemoryBarrier {
	p := (*MemoryBarrier)(MemAlloc(unsafe.Sizeof(*(*MemoryBarrier)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_BARRIER
	return p
}
func (p *MemoryBarrier) Free() { MemFree(unsafe.Pointer(p)) }

// BufferMemoryBarrier -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferMemoryBarrier.html
type BufferMemoryBarrier struct {
	SType               StructureType
	PNext               unsafe.Pointer
	SrcAccessMask       AccessFlags
	DstAccessMask       AccessFlags
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              Buffer
	Offset              DeviceSize
	Size                DeviceSize
}

func NewBufferMemoryBarrier() *BufferMemoryBarrier {
	p := (*BufferMemoryBarrier)(MemAlloc(unsafe.Sizeof(*(*BufferMemoryBarrier)(nil))))
	p.SType = STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER
	return p
}
func (p *BufferMemoryBarrier) Free() { MemFree(unsafe.Pointer(p)) }

// ImageMemoryBarrier -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageMemoryBarrier.html
type ImageMemoryBarrier struct {
	SType               StructureType
	PNext               unsafe.Pointer
	SrcAccessMask       AccessFlags
	DstAccessMask       AccessFlags
	OldLayout           ImageLayout
	NewLayout           ImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               Image
	SubresourceRange    ImageSubresourceRange
}

func NewImageMemoryBarrier() *ImageMemoryBarrier {
	p := (*ImageMemoryBarrier)(MemAlloc(unsafe.Sizeof(*(*ImageMemoryBarrier)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER
	return p
}
func (p *ImageMemoryBarrier) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassBeginInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassBeginInfo.html
type RenderPassBeginInfo struct {
	SType           StructureType
	PNext           unsafe.Pointer
	RenderPass      RenderPass
	Framebuffer     Framebuffer
	RenderArea      Rect2D
	ClearValueCount uint32
	PClearValues    *ClearValue
}

func NewRenderPassBeginInfo() *RenderPassBeginInfo {
	p := (*RenderPassBeginInfo)(MemAlloc(unsafe.Sizeof(*(*RenderPassBeginInfo)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
	return p
}
func (p *RenderPassBeginInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DispatchIndirectCommand -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDispatchIndirectCommand.html
type DispatchIndirectCommand struct {
	X uint32
	Y uint32
	Z uint32
}

func NewDispatchIndirectCommand() *DispatchIndirectCommand {
	return (*DispatchIndirectCommand)(MemAlloc(unsafe.Sizeof(*(*DispatchIndirectCommand)(nil))))
}
func (p *DispatchIndirectCommand) Free() { MemFree(unsafe.Pointer(p)) }

// DrawIndexedIndirectCommand -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDrawIndexedIndirectCommand.html
type DrawIndexedIndirectCommand struct {
	IndexCount    uint32
	InstanceCount uint32
	FirstIndex    uint32
	VertexOffset  int32
	FirstInstance uint32
}

func NewDrawIndexedIndirectCommand() *DrawIndexedIndirectCommand {
	return (*DrawIndexedIndirectCommand)(MemAlloc(unsafe.Sizeof(*(*DrawIndexedIndirectCommand)(nil))))
}
func (p *DrawIndexedIndirectCommand) Free() { MemFree(unsafe.Pointer(p)) }

// DrawIndirectCommand -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDrawIndirectCommand.html
type DrawIndirectCommand struct {
	VertexCount   uint32
	InstanceCount uint32
	FirstVertex   uint32
	FirstInstance uint32
}

func NewDrawIndirectCommand() *DrawIndirectCommand {
	return (*DrawIndirectCommand)(MemAlloc(unsafe.Sizeof(*(*DrawIndirectCommand)(nil))))
}
func (p *DrawIndirectCommand) Free() { MemFree(unsafe.Pointer(p)) }

// BaseOutStructure -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBaseOutStructure.html
type BaseOutStructure struct {
	SType StructureType
	PNext *BaseOutStructure
}

func NewBaseOutStructure() *BaseOutStructure {
	return (*BaseOutStructure)(MemAlloc(unsafe.Sizeof(*(*BaseOutStructure)(nil))))
}
func (p *BaseOutStructure) Free() { MemFree(unsafe.Pointer(p)) }

// BaseInStructure -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBaseInStructure.html
type BaseInStructure struct {
	SType StructureType
	PNext *BaseInStructure
}

func NewBaseInStructure() *BaseInStructure {
	return (*BaseInStructure)(MemAlloc(unsafe.Sizeof(*(*BaseInStructure)(nil))))
}
func (p *BaseInStructure) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateInstance -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateInstance.html
type PfnCreateInstance uintptr

func (fn PfnCreateInstance) Call(pCreateInfo *InstanceCreateInfo, pAllocator *AllocationCallbacks, pInstance *Instance) Result {
	ret, _, _ := call(uintptr(fn), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pInstance)))
	return Result(ret)
}
func (fn PfnCreateInstance) String() string { return "vkCreateInstance" }

//  PfnDestroyInstance -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyInstance.html
type PfnDestroyInstance uintptr

func (fn PfnDestroyInstance) Call(instance Instance, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyInstance) String() string { return "vkDestroyInstance" }

//  PfnEnumeratePhysicalDevices -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumeratePhysicalDevices.html
type PfnEnumeratePhysicalDevices uintptr

func (fn PfnEnumeratePhysicalDevices) Call(instance Instance, pPhysicalDeviceCount *uint32, pPhysicalDevices *PhysicalDevice) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pPhysicalDeviceCount)), uintptr(unsafe.Pointer(pPhysicalDevices)))
	return Result(ret)
}
func (fn PfnEnumeratePhysicalDevices) String() string { return "vkEnumeratePhysicalDevices" }

//  PfnGetPhysicalDeviceFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceFeatures.html
type PfnGetPhysicalDeviceFeatures uintptr

func (fn PfnGetPhysicalDeviceFeatures) Call(physicalDevice PhysicalDevice, pFeatures *PhysicalDeviceFeatures) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pFeatures)))
}
func (fn PfnGetPhysicalDeviceFeatures) String() string { return "vkGetPhysicalDeviceFeatures" }

//  PfnGetPhysicalDeviceFormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html
type PfnGetPhysicalDeviceFormatProperties uintptr

func (fn PfnGetPhysicalDeviceFormatProperties) Call(physicalDevice PhysicalDevice, format Format, pFormatProperties *FormatProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(format), uintptr(unsafe.Pointer(pFormatProperties)))
}
func (fn PfnGetPhysicalDeviceFormatProperties) String() string {
	return "vkGetPhysicalDeviceFormatProperties"
}

//  PfnGetPhysicalDeviceImageFormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html
type PfnGetPhysicalDeviceImageFormatProperties uintptr

func (fn PfnGetPhysicalDeviceImageFormatProperties) Call(physicalDevice PhysicalDevice, format Format, type_ ImageType, tiling ImageTiling, usage ImageUsageFlags, flags ImageCreateFlags, pImageFormatProperties *ImageFormatProperties) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(format), uintptr(type_), uintptr(tiling), uintptr(usage), uintptr(flags), uintptr(unsafe.Pointer(pImageFormatProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceImageFormatProperties) String() string {
	return "vkGetPhysicalDeviceImageFormatProperties"
}

//  PfnGetPhysicalDeviceProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceProperties.html
type PfnGetPhysicalDeviceProperties uintptr

func (fn PfnGetPhysicalDeviceProperties) Call(physicalDevice PhysicalDevice, pProperties *PhysicalDeviceProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pProperties)))
}
func (fn PfnGetPhysicalDeviceProperties) String() string { return "vkGetPhysicalDeviceProperties" }

//  PfnGetPhysicalDeviceQueueFamilyProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html
type PfnGetPhysicalDeviceQueueFamilyProperties uintptr

func (fn PfnGetPhysicalDeviceQueueFamilyProperties) Call(physicalDevice PhysicalDevice, pQueueFamilyPropertyCount *uint32, pQueueFamilyProperties *QueueFamilyProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)), uintptr(unsafe.Pointer(pQueueFamilyProperties)))
}
func (fn PfnGetPhysicalDeviceQueueFamilyProperties) String() string {
	return "vkGetPhysicalDeviceQueueFamilyProperties"
}

//  PfnGetPhysicalDeviceMemoryProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html
type PfnGetPhysicalDeviceMemoryProperties uintptr

func (fn PfnGetPhysicalDeviceMemoryProperties) Call(physicalDevice PhysicalDevice, pMemoryProperties *PhysicalDeviceMemoryProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pMemoryProperties)))
}
func (fn PfnGetPhysicalDeviceMemoryProperties) String() string {
	return "vkGetPhysicalDeviceMemoryProperties"
}

//  PfnGetInstanceProcAddr -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetInstanceProcAddr.html
type PfnGetInstanceProcAddr uintptr

func (fn PfnGetInstanceProcAddr) Call(instance Instance, pName *int8) PfnVoidFunction {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pName)))
	return PfnVoidFunction(ret)
}
func (fn PfnGetInstanceProcAddr) String() string { return "vkGetInstanceProcAddr" }

//  PfnGetDeviceProcAddr -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceProcAddr.html
type PfnGetDeviceProcAddr uintptr

func (fn PfnGetDeviceProcAddr) Call(device Device, pName *int8) PfnVoidFunction {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pName)))
	return PfnVoidFunction(ret)
}
func (fn PfnGetDeviceProcAddr) String() string { return "vkGetDeviceProcAddr" }

//  PfnCreateDevice -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDevice.html
type PfnCreateDevice uintptr

func (fn PfnCreateDevice) Call(physicalDevice PhysicalDevice, pCreateInfo *DeviceCreateInfo, pAllocator *AllocationCallbacks, pDevice *Device) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDevice)))
	return Result(ret)
}
func (fn PfnCreateDevice) String() string { return "vkCreateDevice" }

//  PfnDestroyDevice -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyDevice.html
type PfnDestroyDevice uintptr

func (fn PfnDestroyDevice) Call(device Device, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyDevice) String() string { return "vkDestroyDevice" }

//  PfnEnumerateInstanceExtensionProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumerateInstanceExtensionProperties.html
type PfnEnumerateInstanceExtensionProperties uintptr

func (fn PfnEnumerateInstanceExtensionProperties) Call(pLayerName *int8, pPropertyCount *uint32, pProperties *ExtensionProperties) Result {
	ret, _, _ := call(uintptr(fn), uintptr(unsafe.Pointer(pLayerName)), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnEnumerateInstanceExtensionProperties) String() string {
	return "vkEnumerateInstanceExtensionProperties"
}

//  PfnEnumerateDeviceExtensionProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumerateDeviceExtensionProperties.html
type PfnEnumerateDeviceExtensionProperties uintptr

func (fn PfnEnumerateDeviceExtensionProperties) Call(physicalDevice PhysicalDevice, pLayerName *int8, pPropertyCount *uint32, pProperties *ExtensionProperties) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pLayerName)), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnEnumerateDeviceExtensionProperties) String() string {
	return "vkEnumerateDeviceExtensionProperties"
}

//  PfnEnumerateInstanceLayerProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumerateInstanceLayerProperties.html
type PfnEnumerateInstanceLayerProperties uintptr

func (fn PfnEnumerateInstanceLayerProperties) Call(pPropertyCount *uint32, pProperties *LayerProperties) Result {
	ret, _, _ := call(uintptr(fn), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnEnumerateInstanceLayerProperties) String() string {
	return "vkEnumerateInstanceLayerProperties"
}

//  PfnEnumerateDeviceLayerProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumerateDeviceLayerProperties.html
type PfnEnumerateDeviceLayerProperties uintptr

func (fn PfnEnumerateDeviceLayerProperties) Call(physicalDevice PhysicalDevice, pPropertyCount *uint32, pProperties *LayerProperties) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnEnumerateDeviceLayerProperties) String() string {
	return "vkEnumerateDeviceLayerProperties"
}

//  PfnGetDeviceQueue -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceQueue.html
type PfnGetDeviceQueue uintptr

func (fn PfnGetDeviceQueue) Call(device Device, queueFamilyIndex, queueIndex uint32, pQueue *Queue) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(queueFamilyIndex), uintptr(queueIndex), uintptr(unsafe.Pointer(pQueue)))
}
func (fn PfnGetDeviceQueue) String() string { return "vkGetDeviceQueue" }

//  PfnQueueSubmit -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueueSubmit.html
type PfnQueueSubmit uintptr

func (fn PfnQueueSubmit) Call(queue Queue, submitCount uint32, pSubmits *SubmitInfo, fence Fence) Result {
	ret, _, _ := call(uintptr(fn), uintptr(queue), uintptr(submitCount), uintptr(unsafe.Pointer(pSubmits)), uintptr(fence))
	return Result(ret)
}
func (fn PfnQueueSubmit) String() string { return "vkQueueSubmit" }

//  PfnQueueWaitIdle -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueueWaitIdle.html
type PfnQueueWaitIdle uintptr

func (fn PfnQueueWaitIdle) Call(queue Queue) Result {
	ret, _, _ := call(uintptr(fn), uintptr(queue))
	return Result(ret)
}
func (fn PfnQueueWaitIdle) String() string { return "vkQueueWaitIdle" }

//  PfnDeviceWaitIdle -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDeviceWaitIdle.html
type PfnDeviceWaitIdle uintptr

func (fn PfnDeviceWaitIdle) Call(device Device) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device))
	return Result(ret)
}
func (fn PfnDeviceWaitIdle) String() string { return "vkDeviceWaitIdle" }

//  PfnAllocateMemory -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAllocateMemory.html
type PfnAllocateMemory uintptr

func (fn PfnAllocateMemory) Call(device Device, pAllocateInfo *MemoryAllocateInfo, pAllocator *AllocationCallbacks, pMemory *DeviceMemory) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pAllocateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pMemory)))
	return Result(ret)
}
func (fn PfnAllocateMemory) String() string { return "vkAllocateMemory" }

//  PfnFreeMemory -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkFreeMemory.html
type PfnFreeMemory uintptr

func (fn PfnFreeMemory) Call(device Device, memory DeviceMemory, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(memory), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnFreeMemory) String() string { return "vkFreeMemory" }

//  PfnMapMemory -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkMapMemory.html
type PfnMapMemory uintptr

func (fn PfnMapMemory) Call(device Device, memory DeviceMemory, offset, size DeviceSize, flags MemoryMapFlags, ppData *unsafe.Pointer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(memory), uintptr(offset), uintptr(size), uintptr(flags), uintptr(unsafe.Pointer(ppData)))
	return Result(ret)
}
func (fn PfnMapMemory) String() string { return "vkMapMemory" }

//  PfnUnmapMemory -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkUnmapMemory.html
type PfnUnmapMemory uintptr

func (fn PfnUnmapMemory) Call(device Device, memory DeviceMemory) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(memory))
}
func (fn PfnUnmapMemory) String() string { return "vkUnmapMemory" }

//  PfnFlushMappedMemoryRanges -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkFlushMappedMemoryRanges.html
type PfnFlushMappedMemoryRanges uintptr

func (fn PfnFlushMappedMemoryRanges) Call(device Device, memoryRangeCount uint32, pMemoryRanges *MappedMemoryRange) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(memoryRangeCount), uintptr(unsafe.Pointer(pMemoryRanges)))
	return Result(ret)
}
func (fn PfnFlushMappedMemoryRanges) String() string { return "vkFlushMappedMemoryRanges" }

//  PfnInvalidateMappedMemoryRanges -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkInvalidateMappedMemoryRanges.html
type PfnInvalidateMappedMemoryRanges uintptr

func (fn PfnInvalidateMappedMemoryRanges) Call(device Device, memoryRangeCount uint32, pMemoryRanges *MappedMemoryRange) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(memoryRangeCount), uintptr(unsafe.Pointer(pMemoryRanges)))
	return Result(ret)
}
func (fn PfnInvalidateMappedMemoryRanges) String() string { return "vkInvalidateMappedMemoryRanges" }

//  PfnGetDeviceMemoryCommitment -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceMemoryCommitment.html
type PfnGetDeviceMemoryCommitment uintptr

func (fn PfnGetDeviceMemoryCommitment) Call(device Device, memory DeviceMemory, pCommittedMemoryInBytes *DeviceSize) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(memory), uintptr(unsafe.Pointer(pCommittedMemoryInBytes)))
}
func (fn PfnGetDeviceMemoryCommitment) String() string { return "vkGetDeviceMemoryCommitment" }

//  PfnBindBufferMemory -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBindBufferMemory.html
type PfnBindBufferMemory uintptr

func (fn PfnBindBufferMemory) Call(device Device, buffer Buffer, memory DeviceMemory, memoryOffset DeviceSize) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(buffer), uintptr(memory), uintptr(memoryOffset))
	return Result(ret)
}
func (fn PfnBindBufferMemory) String() string { return "vkBindBufferMemory" }

//  PfnBindImageMemory -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBindImageMemory.html
type PfnBindImageMemory uintptr

func (fn PfnBindImageMemory) Call(device Device, image Image, memory DeviceMemory, memoryOffset DeviceSize) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(image), uintptr(memory), uintptr(memoryOffset))
	return Result(ret)
}
func (fn PfnBindImageMemory) String() string { return "vkBindImageMemory" }

//  PfnGetBufferMemoryRequirements -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetBufferMemoryRequirements.html
type PfnGetBufferMemoryRequirements uintptr

func (fn PfnGetBufferMemoryRequirements) Call(device Device, buffer Buffer, pMemoryRequirements *MemoryRequirements) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(buffer), uintptr(unsafe.Pointer(pMemoryRequirements)))
}
func (fn PfnGetBufferMemoryRequirements) String() string { return "vkGetBufferMemoryRequirements" }

//  PfnGetImageMemoryRequirements -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageMemoryRequirements.html
type PfnGetImageMemoryRequirements uintptr

func (fn PfnGetImageMemoryRequirements) Call(device Device, image Image, pMemoryRequirements *MemoryRequirements) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pMemoryRequirements)))
}
func (fn PfnGetImageMemoryRequirements) String() string { return "vkGetImageMemoryRequirements" }

//  PfnGetImageSparseMemoryRequirements -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageSparseMemoryRequirements.html
type PfnGetImageSparseMemoryRequirements uintptr

func (fn PfnGetImageSparseMemoryRequirements) Call(device Device, image Image, pSparseMemoryRequirementCount *uint32, pSparseMemoryRequirements *SparseImageMemoryRequirements) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)), uintptr(unsafe.Pointer(pSparseMemoryRequirements)))
}
func (fn PfnGetImageSparseMemoryRequirements) String() string {
	return "vkGetImageSparseMemoryRequirements"
}

//  PfnGetPhysicalDeviceSparseImageFormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html
type PfnGetPhysicalDeviceSparseImageFormatProperties uintptr

func (fn PfnGetPhysicalDeviceSparseImageFormatProperties) Call(physicalDevice PhysicalDevice, format Format, type_ ImageType, samples SampleCountFlags, usage ImageUsageFlags, tiling ImageTiling, pPropertyCount *uint32, pProperties *SparseImageFormatProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(format), uintptr(type_), uintptr(samples), uintptr(usage), uintptr(tiling), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
}
func (fn PfnGetPhysicalDeviceSparseImageFormatProperties) String() string {
	return "vkGetPhysicalDeviceSparseImageFormatProperties"
}

//  PfnQueueBindSparse -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueueBindSparse.html
type PfnQueueBindSparse uintptr

func (fn PfnQueueBindSparse) Call(queue Queue, bindInfoCount uint32, pBindInfo *BindSparseInfo, fence Fence) Result {
	ret, _, _ := call(uintptr(fn), uintptr(queue), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfo)), uintptr(fence))
	return Result(ret)
}
func (fn PfnQueueBindSparse) String() string { return "vkQueueBindSparse" }

//  PfnCreateFence -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateFence.html
type PfnCreateFence uintptr

func (fn PfnCreateFence) Call(device Device, pCreateInfo *FenceCreateInfo, pAllocator *AllocationCallbacks, pFence *Fence) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFence)))
	return Result(ret)
}
func (fn PfnCreateFence) String() string { return "vkCreateFence" }

//  PfnDestroyFence -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyFence.html
type PfnDestroyFence uintptr

func (fn PfnDestroyFence) Call(device Device, fence Fence, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(fence), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyFence) String() string { return "vkDestroyFence" }

//  PfnResetFences -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkResetFences.html
type PfnResetFences uintptr

func (fn PfnResetFences) Call(device Device, fenceCount uint32, pFences *Fence) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(fenceCount), uintptr(unsafe.Pointer(pFences)))
	return Result(ret)
}
func (fn PfnResetFences) String() string { return "vkResetFences" }

//  PfnGetFenceStatus -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetFenceStatus.html
type PfnGetFenceStatus uintptr

func (fn PfnGetFenceStatus) Call(device Device, fence Fence) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(fence))
	return Result(ret)
}
func (fn PfnGetFenceStatus) String() string { return "vkGetFenceStatus" }

//  PfnWaitForFences -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkWaitForFences.html
type PfnWaitForFences uintptr

func (fn PfnWaitForFences) Call(device Device, fenceCount uint32, pFences *Fence, waitAll Bool32, timeout uint64) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(fenceCount), uintptr(unsafe.Pointer(pFences)), uintptr(waitAll), uintptr(timeout))
	return Result(ret)
}
func (fn PfnWaitForFences) String() string { return "vkWaitForFences" }

//  PfnCreateSemaphore -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateSemaphore.html
type PfnCreateSemaphore uintptr

func (fn PfnCreateSemaphore) Call(device Device, pCreateInfo *SemaphoreCreateInfo, pAllocator *AllocationCallbacks, pSemaphore *Semaphore) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSemaphore)))
	return Result(ret)
}
func (fn PfnCreateSemaphore) String() string { return "vkCreateSemaphore" }

//  PfnDestroySemaphore -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroySemaphore.html
type PfnDestroySemaphore uintptr

func (fn PfnDestroySemaphore) Call(device Device, semaphore Semaphore, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(semaphore), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroySemaphore) String() string { return "vkDestroySemaphore" }

//  PfnCreateEvent -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateEvent.html
type PfnCreateEvent uintptr

func (fn PfnCreateEvent) Call(device Device, pCreateInfo *EventCreateInfo, pAllocator *AllocationCallbacks, pEvent *Event) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pEvent)))
	return Result(ret)
}
func (fn PfnCreateEvent) String() string { return "vkCreateEvent" }

//  PfnDestroyEvent -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyEvent.html
type PfnDestroyEvent uintptr

func (fn PfnDestroyEvent) Call(device Device, event Event, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(event), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyEvent) String() string { return "vkDestroyEvent" }

//  PfnGetEventStatus -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetEventStatus.html
type PfnGetEventStatus uintptr

func (fn PfnGetEventStatus) Call(device Device, event Event) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(event))
	return Result(ret)
}
func (fn PfnGetEventStatus) String() string { return "vkGetEventStatus" }

//  PfnSetEvent -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkSetEvent.html
type PfnSetEvent uintptr

func (fn PfnSetEvent) Call(device Device, event Event) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(event))
	return Result(ret)
}
func (fn PfnSetEvent) String() string { return "vkSetEvent" }

//  PfnResetEvent -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkResetEvent.html
type PfnResetEvent uintptr

func (fn PfnResetEvent) Call(device Device, event Event) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(event))
	return Result(ret)
}
func (fn PfnResetEvent) String() string { return "vkResetEvent" }

//  PfnCreateQueryPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateQueryPool.html
type PfnCreateQueryPool uintptr

func (fn PfnCreateQueryPool) Call(device Device, pCreateInfo *QueryPoolCreateInfo, pAllocator *AllocationCallbacks, pQueryPool *QueryPool) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pQueryPool)))
	return Result(ret)
}
func (fn PfnCreateQueryPool) String() string { return "vkCreateQueryPool" }

//  PfnDestroyQueryPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyQueryPool.html
type PfnDestroyQueryPool uintptr

func (fn PfnDestroyQueryPool) Call(device Device, queryPool QueryPool, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(queryPool), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyQueryPool) String() string { return "vkDestroyQueryPool" }

//  PfnGetQueryPoolResults -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetQueryPoolResults.html
type PfnGetQueryPoolResults uintptr

func (fn PfnGetQueryPoolResults) Call(device Device, queryPool QueryPool, firstQuery, queryCount uint32, dataSize uintptr, pData unsafe.Pointer, stride DeviceSize, flags QueryResultFlags) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount), uintptr(dataSize), uintptr(pData), uintptr(stride), uintptr(flags))
	return Result(ret)
}
func (fn PfnGetQueryPoolResults) String() string { return "vkGetQueryPoolResults" }

//  PfnCreateBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateBuffer.html
type PfnCreateBuffer uintptr

func (fn PfnCreateBuffer) Call(device Device, pCreateInfo *BufferCreateInfo, pAllocator *AllocationCallbacks, pBuffer *Buffer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pBuffer)))
	return Result(ret)
}
func (fn PfnCreateBuffer) String() string { return "vkCreateBuffer" }

//  PfnDestroyBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyBuffer.html
type PfnDestroyBuffer uintptr

func (fn PfnDestroyBuffer) Call(device Device, buffer Buffer, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(buffer), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyBuffer) String() string { return "vkDestroyBuffer" }

//  PfnCreateBufferView -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateBufferView.html
type PfnCreateBufferView uintptr

func (fn PfnCreateBufferView) Call(device Device, pCreateInfo *BufferViewCreateInfo, pAllocator *AllocationCallbacks, pView *BufferView) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pView)))
	return Result(ret)
}
func (fn PfnCreateBufferView) String() string { return "vkCreateBufferView" }

//  PfnDestroyBufferView -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyBufferView.html
type PfnDestroyBufferView uintptr

func (fn PfnDestroyBufferView) Call(device Device, bufferView BufferView, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(bufferView), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyBufferView) String() string { return "vkDestroyBufferView" }

//  PfnCreateImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateImage.html
type PfnCreateImage uintptr

func (fn PfnCreateImage) Call(device Device, pCreateInfo *ImageCreateInfo, pAllocator *AllocationCallbacks, pImage *Image) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pImage)))
	return Result(ret)
}
func (fn PfnCreateImage) String() string { return "vkCreateImage" }

//  PfnDestroyImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyImage.html
type PfnDestroyImage uintptr

func (fn PfnDestroyImage) Call(device Device, image Image, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyImage) String() string { return "vkDestroyImage" }

//  PfnGetImageSubresourceLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageSubresourceLayout.html
type PfnGetImageSubresourceLayout uintptr

func (fn PfnGetImageSubresourceLayout) Call(device Device, image Image, pSubresource *ImageSubresource, pLayout *SubresourceLayout) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pSubresource)), uintptr(unsafe.Pointer(pLayout)))
}
func (fn PfnGetImageSubresourceLayout) String() string { return "vkGetImageSubresourceLayout" }

//  PfnCreateImageView -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateImageView.html
type PfnCreateImageView uintptr

func (fn PfnCreateImageView) Call(device Device, pCreateInfo *ImageViewCreateInfo, pAllocator *AllocationCallbacks, pView *ImageView) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pView)))
	return Result(ret)
}
func (fn PfnCreateImageView) String() string { return "vkCreateImageView" }

//  PfnDestroyImageView -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyImageView.html
type PfnDestroyImageView uintptr

func (fn PfnDestroyImageView) Call(device Device, imageView ImageView, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(imageView), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyImageView) String() string { return "vkDestroyImageView" }

//  PfnCreateShaderModule -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateShaderModule.html
type PfnCreateShaderModule uintptr

func (fn PfnCreateShaderModule) Call(device Device, pCreateInfo *ShaderModuleCreateInfo, pAllocator *AllocationCallbacks, pShaderModule *ShaderModule) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pShaderModule)))
	return Result(ret)
}
func (fn PfnCreateShaderModule) String() string { return "vkCreateShaderModule" }

//  PfnDestroyShaderModule -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyShaderModule.html
type PfnDestroyShaderModule uintptr

func (fn PfnDestroyShaderModule) Call(device Device, shaderModule ShaderModule, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(shaderModule), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyShaderModule) String() string { return "vkDestroyShaderModule" }

//  PfnCreatePipelineCache -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreatePipelineCache.html
type PfnCreatePipelineCache uintptr

func (fn PfnCreatePipelineCache) Call(device Device, pCreateInfo *PipelineCacheCreateInfo, pAllocator *AllocationCallbacks, pPipelineCache *PipelineCache) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelineCache)))
	return Result(ret)
}
func (fn PfnCreatePipelineCache) String() string { return "vkCreatePipelineCache" }

//  PfnDestroyPipelineCache -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyPipelineCache.html
type PfnDestroyPipelineCache uintptr

func (fn PfnDestroyPipelineCache) Call(device Device, pipelineCache PipelineCache, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(pipelineCache), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyPipelineCache) String() string { return "vkDestroyPipelineCache" }

//  PfnGetPipelineCacheData -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPipelineCacheData.html
type PfnGetPipelineCacheData uintptr

func (fn PfnGetPipelineCacheData) Call(device Device, pipelineCache PipelineCache, pDataSize *uintptr, pData unsafe.Pointer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(pipelineCache), uintptr(unsafe.Pointer(pDataSize)), uintptr(pData))
	return Result(ret)
}
func (fn PfnGetPipelineCacheData) String() string { return "vkGetPipelineCacheData" }

//  PfnMergePipelineCaches -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkMergePipelineCaches.html
type PfnMergePipelineCaches uintptr

func (fn PfnMergePipelineCaches) Call(device Device, dstCache PipelineCache, srcCacheCount uint32, pSrcCaches *PipelineCache) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(dstCache), uintptr(srcCacheCount), uintptr(unsafe.Pointer(pSrcCaches)))
	return Result(ret)
}
func (fn PfnMergePipelineCaches) String() string { return "vkMergePipelineCaches" }

//  PfnCreateGraphicsPipelines -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateGraphicsPipelines.html
type PfnCreateGraphicsPipelines uintptr

func (fn PfnCreateGraphicsPipelines) Call(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos *GraphicsPipelineCreateInfo, pAllocator *AllocationCallbacks, pPipelines *Pipeline) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(pipelineCache), uintptr(createInfoCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelines)))
	return Result(ret)
}
func (fn PfnCreateGraphicsPipelines) String() string { return "vkCreateGraphicsPipelines" }

//  PfnCreateComputePipelines -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateComputePipelines.html
type PfnCreateComputePipelines uintptr

func (fn PfnCreateComputePipelines) Call(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos *ComputePipelineCreateInfo, pAllocator *AllocationCallbacks, pPipelines *Pipeline) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(pipelineCache), uintptr(createInfoCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelines)))
	return Result(ret)
}
func (fn PfnCreateComputePipelines) String() string { return "vkCreateComputePipelines" }

//  PfnDestroyPipeline -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyPipeline.html
type PfnDestroyPipeline uintptr

func (fn PfnDestroyPipeline) Call(device Device, pipeline Pipeline, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(pipeline), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyPipeline) String() string { return "vkDestroyPipeline" }

//  PfnCreatePipelineLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreatePipelineLayout.html
type PfnCreatePipelineLayout uintptr

func (fn PfnCreatePipelineLayout) Call(device Device, pCreateInfo *PipelineLayoutCreateInfo, pAllocator *AllocationCallbacks, pPipelineLayout *PipelineLayout) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelineLayout)))
	return Result(ret)
}
func (fn PfnCreatePipelineLayout) String() string { return "vkCreatePipelineLayout" }

//  PfnDestroyPipelineLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyPipelineLayout.html
type PfnDestroyPipelineLayout uintptr

func (fn PfnDestroyPipelineLayout) Call(device Device, pipelineLayout PipelineLayout, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(pipelineLayout), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyPipelineLayout) String() string { return "vkDestroyPipelineLayout" }

//  PfnCreateSampler -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateSampler.html
type PfnCreateSampler uintptr

func (fn PfnCreateSampler) Call(device Device, pCreateInfo *SamplerCreateInfo, pAllocator *AllocationCallbacks, pSampler *Sampler) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSampler)))
	return Result(ret)
}
func (fn PfnCreateSampler) String() string { return "vkCreateSampler" }

//  PfnDestroySampler -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroySampler.html
type PfnDestroySampler uintptr

func (fn PfnDestroySampler) Call(device Device, sampler Sampler, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(sampler), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroySampler) String() string { return "vkDestroySampler" }

//  PfnCreateDescriptorSetLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDescriptorSetLayout.html
type PfnCreateDescriptorSetLayout uintptr

func (fn PfnCreateDescriptorSetLayout) Call(device Device, pCreateInfo *DescriptorSetLayoutCreateInfo, pAllocator *AllocationCallbacks, pSetLayout *DescriptorSetLayout) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSetLayout)))
	return Result(ret)
}
func (fn PfnCreateDescriptorSetLayout) String() string { return "vkCreateDescriptorSetLayout" }

//  PfnDestroyDescriptorSetLayout -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyDescriptorSetLayout.html
type PfnDestroyDescriptorSetLayout uintptr

func (fn PfnDestroyDescriptorSetLayout) Call(device Device, descriptorSetLayout DescriptorSetLayout, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(descriptorSetLayout), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyDescriptorSetLayout) String() string { return "vkDestroyDescriptorSetLayout" }

//  PfnCreateDescriptorPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDescriptorPool.html
type PfnCreateDescriptorPool uintptr

func (fn PfnCreateDescriptorPool) Call(device Device, pCreateInfo *DescriptorPoolCreateInfo, pAllocator *AllocationCallbacks, pDescriptorPool *DescriptorPool) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDescriptorPool)))
	return Result(ret)
}
func (fn PfnCreateDescriptorPool) String() string { return "vkCreateDescriptorPool" }

//  PfnDestroyDescriptorPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyDescriptorPool.html
type PfnDestroyDescriptorPool uintptr

func (fn PfnDestroyDescriptorPool) Call(device Device, descriptorPool DescriptorPool, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(descriptorPool), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyDescriptorPool) String() string { return "vkDestroyDescriptorPool" }

//  PfnResetDescriptorPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkResetDescriptorPool.html
type PfnResetDescriptorPool uintptr

func (fn PfnResetDescriptorPool) Call(device Device, descriptorPool DescriptorPool, flags DescriptorPoolResetFlags) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(descriptorPool), uintptr(flags))
	return Result(ret)
}
func (fn PfnResetDescriptorPool) String() string { return "vkResetDescriptorPool" }

//  PfnAllocateDescriptorSets -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAllocateDescriptorSets.html
type PfnAllocateDescriptorSets uintptr

func (fn PfnAllocateDescriptorSets) Call(device Device, pAllocateInfo *DescriptorSetAllocateInfo, pDescriptorSets *DescriptorSet) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pAllocateInfo)), uintptr(unsafe.Pointer(pDescriptorSets)))
	return Result(ret)
}
func (fn PfnAllocateDescriptorSets) String() string { return "vkAllocateDescriptorSets" }

//  PfnFreeDescriptorSets -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkFreeDescriptorSets.html
type PfnFreeDescriptorSets uintptr

func (fn PfnFreeDescriptorSets) Call(device Device, descriptorPool DescriptorPool, descriptorSetCount uint32, pDescriptorSets *DescriptorSet) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(descriptorPool), uintptr(descriptorSetCount), uintptr(unsafe.Pointer(pDescriptorSets)))
	return Result(ret)
}
func (fn PfnFreeDescriptorSets) String() string { return "vkFreeDescriptorSets" }

//  PfnUpdateDescriptorSets -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkUpdateDescriptorSets.html
type PfnUpdateDescriptorSets uintptr

func (fn PfnUpdateDescriptorSets) Call(device Device, descriptorWriteCount uint32, pDescriptorWrites *WriteDescriptorSet, descriptorCopyCount uint32, pDescriptorCopies *CopyDescriptorSet) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(descriptorWriteCount), uintptr(unsafe.Pointer(pDescriptorWrites)), uintptr(descriptorCopyCount), uintptr(unsafe.Pointer(pDescriptorCopies)))
}
func (fn PfnUpdateDescriptorSets) String() string { return "vkUpdateDescriptorSets" }

//  PfnCreateFramebuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateFramebuffer.html
type PfnCreateFramebuffer uintptr

func (fn PfnCreateFramebuffer) Call(device Device, pCreateInfo *FramebufferCreateInfo, pAllocator *AllocationCallbacks, pFramebuffer *Framebuffer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFramebuffer)))
	return Result(ret)
}
func (fn PfnCreateFramebuffer) String() string { return "vkCreateFramebuffer" }

//  PfnDestroyFramebuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyFramebuffer.html
type PfnDestroyFramebuffer uintptr

func (fn PfnDestroyFramebuffer) Call(device Device, framebuffer Framebuffer, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(framebuffer), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyFramebuffer) String() string { return "vkDestroyFramebuffer" }

//  PfnCreateRenderPass -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateRenderPass.html
type PfnCreateRenderPass uintptr

func (fn PfnCreateRenderPass) Call(device Device, pCreateInfo *RenderPassCreateInfo, pAllocator *AllocationCallbacks, pRenderPass *RenderPass) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pRenderPass)))
	return Result(ret)
}
func (fn PfnCreateRenderPass) String() string { return "vkCreateRenderPass" }

//  PfnDestroyRenderPass -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyRenderPass.html
type PfnDestroyRenderPass uintptr

func (fn PfnDestroyRenderPass) Call(device Device, renderPass RenderPass, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(renderPass), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyRenderPass) String() string { return "vkDestroyRenderPass" }

//  PfnGetRenderAreaGranularity -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetRenderAreaGranularity.html
type PfnGetRenderAreaGranularity uintptr

func (fn PfnGetRenderAreaGranularity) Call(device Device, renderPass RenderPass, pGranularity *Extent2D) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(renderPass), uintptr(unsafe.Pointer(pGranularity)))
}
func (fn PfnGetRenderAreaGranularity) String() string { return "vkGetRenderAreaGranularity" }

//  PfnCreateCommandPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateCommandPool.html
type PfnCreateCommandPool uintptr

func (fn PfnCreateCommandPool) Call(device Device, pCreateInfo *CommandPoolCreateInfo, pAllocator *AllocationCallbacks, pCommandPool *CommandPool) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pCommandPool)))
	return Result(ret)
}
func (fn PfnCreateCommandPool) String() string { return "vkCreateCommandPool" }

//  PfnDestroyCommandPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyCommandPool.html
type PfnDestroyCommandPool uintptr

func (fn PfnDestroyCommandPool) Call(device Device, commandPool CommandPool, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(commandPool), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyCommandPool) String() string { return "vkDestroyCommandPool" }

//  PfnResetCommandPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkResetCommandPool.html
type PfnResetCommandPool uintptr

func (fn PfnResetCommandPool) Call(device Device, commandPool CommandPool, flags CommandPoolResetFlags) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(commandPool), uintptr(flags))
	return Result(ret)
}
func (fn PfnResetCommandPool) String() string { return "vkResetCommandPool" }

//  PfnAllocateCommandBuffers -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAllocateCommandBuffers.html
type PfnAllocateCommandBuffers uintptr

func (fn PfnAllocateCommandBuffers) Call(device Device, pAllocateInfo *CommandBufferAllocateInfo, pCommandBuffers *CommandBuffer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pAllocateInfo)), uintptr(unsafe.Pointer(pCommandBuffers)))
	return Result(ret)
}
func (fn PfnAllocateCommandBuffers) String() string { return "vkAllocateCommandBuffers" }

//  PfnFreeCommandBuffers -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkFreeCommandBuffers.html
type PfnFreeCommandBuffers uintptr

func (fn PfnFreeCommandBuffers) Call(device Device, commandPool CommandPool, commandBufferCount uint32, pCommandBuffers *CommandBuffer) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(commandPool), uintptr(commandBufferCount), uintptr(unsafe.Pointer(pCommandBuffers)))
}
func (fn PfnFreeCommandBuffers) String() string { return "vkFreeCommandBuffers" }

//  PfnBeginCommandBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBeginCommandBuffer.html
type PfnBeginCommandBuffer uintptr

func (fn PfnBeginCommandBuffer) Call(commandBuffer CommandBuffer, pBeginInfo *CommandBufferBeginInfo) Result {
	ret, _, _ := call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pBeginInfo)))
	return Result(ret)
}
func (fn PfnBeginCommandBuffer) String() string { return "vkBeginCommandBuffer" }

//  PfnEndCommandBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEndCommandBuffer.html
type PfnEndCommandBuffer uintptr

func (fn PfnEndCommandBuffer) Call(commandBuffer CommandBuffer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(commandBuffer))
	return Result(ret)
}
func (fn PfnEndCommandBuffer) String() string { return "vkEndCommandBuffer" }

//  PfnResetCommandBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkResetCommandBuffer.html
type PfnResetCommandBuffer uintptr

func (fn PfnResetCommandBuffer) Call(commandBuffer CommandBuffer, flags CommandBufferResetFlags) Result {
	ret, _, _ := call(uintptr(fn), uintptr(commandBuffer), uintptr(flags))
	return Result(ret)
}
func (fn PfnResetCommandBuffer) String() string { return "vkResetCommandBuffer" }

//  PfnCmdBindPipeline -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBindPipeline.html
type PfnCmdBindPipeline uintptr

func (fn PfnCmdBindPipeline) Call(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(pipelineBindPoint), uintptr(pipeline))
}
func (fn PfnCmdBindPipeline) String() string { return "vkCmdBindPipeline" }

//  PfnCmdSetViewport -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetViewport.html
type PfnCmdSetViewport uintptr

func (fn PfnCmdSetViewport) Call(commandBuffer CommandBuffer, firstViewport, viewportCount uint32, pViewports *Viewport) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstViewport), uintptr(viewportCount), uintptr(unsafe.Pointer(pViewports)))
}
func (fn PfnCmdSetViewport) String() string { return "vkCmdSetViewport" }

//  PfnCmdSetScissor -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetScissor.html
type PfnCmdSetScissor uintptr

func (fn PfnCmdSetScissor) Call(commandBuffer CommandBuffer, firstScissor, scissorCount uint32, pScissors *Rect2D) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstScissor), uintptr(scissorCount), uintptr(unsafe.Pointer(pScissors)))
}
func (fn PfnCmdSetScissor) String() string { return "vkCmdSetScissor" }

//  PfnCmdSetLineWidth -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetLineWidth.html
type PfnCmdSetLineWidth uintptr

func (fn PfnCmdSetLineWidth) Call(commandBuffer CommandBuffer, lineWidth float32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(lineWidth))
}
func (fn PfnCmdSetLineWidth) String() string { return "vkCmdSetLineWidth" }

//  PfnCmdSetDepthBias -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetDepthBias.html
type PfnCmdSetDepthBias uintptr

func (fn PfnCmdSetDepthBias) Call(commandBuffer CommandBuffer, depthBiasConstantFactor, depthBiasClamp, depthBiasSlopeFactor float32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(depthBiasConstantFactor), uintptr(depthBiasClamp), uintptr(depthBiasSlopeFactor))
}
func (fn PfnCmdSetDepthBias) String() string { return "vkCmdSetDepthBias" }

//  PfnCmdSetBlendConstants -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetBlendConstants.html
type PfnCmdSetBlendConstants uintptr

func (fn PfnCmdSetBlendConstants) Call(commandBuffer CommandBuffer, blendConstant *[4]float32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(blendConstant)))
}
func (fn PfnCmdSetBlendConstants) String() string { return "vkCmdSetBlendConstants" }

//  PfnCmdSetDepthBounds -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetDepthBounds.html
type PfnCmdSetDepthBounds uintptr

func (fn PfnCmdSetDepthBounds) Call(commandBuffer CommandBuffer, minDepthBounds, maxDepthBounds float32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(minDepthBounds), uintptr(maxDepthBounds))
}
func (fn PfnCmdSetDepthBounds) String() string { return "vkCmdSetDepthBounds" }

//  PfnCmdSetStencilCompareMask -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetStencilCompareMask.html
type PfnCmdSetStencilCompareMask uintptr

func (fn PfnCmdSetStencilCompareMask) Call(commandBuffer CommandBuffer, faceMask StencilFaceFlags, compareMask uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(faceMask), uintptr(compareMask))
}
func (fn PfnCmdSetStencilCompareMask) String() string { return "vkCmdSetStencilCompareMask" }

//  PfnCmdSetStencilWriteMask -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetStencilWriteMask.html
type PfnCmdSetStencilWriteMask uintptr

func (fn PfnCmdSetStencilWriteMask) Call(commandBuffer CommandBuffer, faceMask StencilFaceFlags, writeMask uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(faceMask), uintptr(writeMask))
}
func (fn PfnCmdSetStencilWriteMask) String() string { return "vkCmdSetStencilWriteMask" }

//  PfnCmdSetStencilReference -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetStencilReference.html
type PfnCmdSetStencilReference uintptr

func (fn PfnCmdSetStencilReference) Call(commandBuffer CommandBuffer, faceMask StencilFaceFlags, reference uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(faceMask), uintptr(reference))
}
func (fn PfnCmdSetStencilReference) String() string { return "vkCmdSetStencilReference" }

//  PfnCmdBindDescriptorSets -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBindDescriptorSets.html
type PfnCmdBindDescriptorSets uintptr

func (fn PfnCmdBindDescriptorSets) Call(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, firstSet, descriptorSetCount uint32, pDescriptorSets *DescriptorSet, dynamicOffsetCount uint32, pDynamicOffsets *uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(pipelineBindPoint), uintptr(layout), uintptr(firstSet), uintptr(descriptorSetCount), uintptr(unsafe.Pointer(pDescriptorSets)), uintptr(dynamicOffsetCount), uintptr(unsafe.Pointer(pDynamicOffsets)))
}
func (fn PfnCmdBindDescriptorSets) String() string { return "vkCmdBindDescriptorSets" }

//  PfnCmdBindIndexBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBindIndexBuffer.html
type PfnCmdBindIndexBuffer uintptr

func (fn PfnCmdBindIndexBuffer) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, indexType IndexType) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(indexType))
}
func (fn PfnCmdBindIndexBuffer) String() string { return "vkCmdBindIndexBuffer" }

//  PfnCmdBindVertexBuffers -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBindVertexBuffers.html
type PfnCmdBindVertexBuffers uintptr

func (fn PfnCmdBindVertexBuffers) Call(commandBuffer CommandBuffer, firstBinding, bindingCount uint32, pBuffers *Buffer, pOffsets *DeviceSize) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstBinding), uintptr(bindingCount), uintptr(unsafe.Pointer(pBuffers)), uintptr(unsafe.Pointer(pOffsets)))
}
func (fn PfnCmdBindVertexBuffers) String() string { return "vkCmdBindVertexBuffers" }

//  PfnCmdDraw -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDraw.html
type PfnCmdDraw uintptr

func (fn PfnCmdDraw) Call(commandBuffer CommandBuffer, vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(vertexCount), uintptr(instanceCount), uintptr(firstVertex), uintptr(firstInstance))
}
func (fn PfnCmdDraw) String() string { return "vkCmdDraw" }

//  PfnCmdDrawIndexed -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndexed.html
type PfnCmdDrawIndexed uintptr

func (fn PfnCmdDrawIndexed) Call(commandBuffer CommandBuffer, indexCount, instanceCount, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(indexCount), uintptr(instanceCount), uintptr(firstIndex), uintptr(vertexOffset), uintptr(firstInstance))
}
func (fn PfnCmdDrawIndexed) String() string { return "vkCmdDrawIndexed" }

//  PfnCmdDrawIndirect -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndirect.html
type PfnCmdDrawIndirect uintptr

func (fn PfnCmdDrawIndirect) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(drawCount), uintptr(stride))
}
func (fn PfnCmdDrawIndirect) String() string { return "vkCmdDrawIndirect" }

//  PfnCmdDrawIndexedIndirect -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndexedIndirect.html
type PfnCmdDrawIndexedIndirect uintptr

func (fn PfnCmdDrawIndexedIndirect) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(drawCount), uintptr(stride))
}
func (fn PfnCmdDrawIndexedIndirect) String() string { return "vkCmdDrawIndexedIndirect" }

//  PfnCmdDispatch -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDispatch.html
type PfnCmdDispatch uintptr

func (fn PfnCmdDispatch) Call(commandBuffer CommandBuffer, groupCountX, groupCountY, groupCountZ uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(groupCountX), uintptr(groupCountY), uintptr(groupCountZ))
}
func (fn PfnCmdDispatch) String() string { return "vkCmdDispatch" }

//  PfnCmdDispatchIndirect -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDispatchIndirect.html
type PfnCmdDispatchIndirect uintptr

func (fn PfnCmdDispatchIndirect) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset))
}
func (fn PfnCmdDispatchIndirect) String() string { return "vkCmdDispatchIndirect" }

//  PfnCmdCopyBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdCopyBuffer.html
type PfnCmdCopyBuffer uintptr

func (fn PfnCmdCopyBuffer) Call(commandBuffer CommandBuffer, srcBuffer, dstBuffer Buffer, regionCount uint32, pRegions *BufferCopy) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(srcBuffer), uintptr(dstBuffer), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)))
}
func (fn PfnCmdCopyBuffer) String() string { return "vkCmdCopyBuffer" }

//  PfnCmdCopyImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdCopyImage.html
type PfnCmdCopyImage uintptr

func (fn PfnCmdCopyImage) Call(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions *ImageCopy) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)))
}
func (fn PfnCmdCopyImage) String() string { return "vkCmdCopyImage" }

//  PfnCmdBlitImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBlitImage.html
type PfnCmdBlitImage uintptr

func (fn PfnCmdBlitImage) Call(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions *ImageBlit, filter Filter) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)), uintptr(filter))
}
func (fn PfnCmdBlitImage) String() string { return "vkCmdBlitImage" }

//  PfnCmdCopyBufferToImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdCopyBufferToImage.html
type PfnCmdCopyBufferToImage uintptr

func (fn PfnCmdCopyBufferToImage) Call(commandBuffer CommandBuffer, srcBuffer Buffer, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions *BufferImageCopy) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(srcBuffer), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)))
}
func (fn PfnCmdCopyBufferToImage) String() string { return "vkCmdCopyBufferToImage" }

//  PfnCmdCopyImageToBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdCopyImageToBuffer.html
type PfnCmdCopyImageToBuffer uintptr

func (fn PfnCmdCopyImageToBuffer) Call(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regionCount uint32, pRegions *BufferImageCopy) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstBuffer), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)))
}
func (fn PfnCmdCopyImageToBuffer) String() string { return "vkCmdCopyImageToBuffer" }

//  PfnCmdUpdateBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdUpdateBuffer.html
type PfnCmdUpdateBuffer uintptr

func (fn PfnCmdUpdateBuffer) Call(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset, dataSize DeviceSize, pData unsafe.Pointer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(dstBuffer), uintptr(dstOffset), uintptr(dataSize), uintptr(pData))
}
func (fn PfnCmdUpdateBuffer) String() string { return "vkCmdUpdateBuffer" }

//  PfnCmdFillBuffer -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdFillBuffer.html
type PfnCmdFillBuffer uintptr

func (fn PfnCmdFillBuffer) Call(commandBuffer CommandBuffer, dstBuffer Buffer, dstOffset, size DeviceSize, data uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(dstBuffer), uintptr(dstOffset), uintptr(size), uintptr(data))
}
func (fn PfnCmdFillBuffer) String() string { return "vkCmdFillBuffer" }

//  PfnCmdClearColorImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdClearColorImage.html
type PfnCmdClearColorImage uintptr

func (fn PfnCmdClearColorImage) Call(commandBuffer CommandBuffer, image Image, imageLayout ImageLayout, pColor *ClearColorValue, rangeCount uint32, pRanges *ImageSubresourceRange) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(image), uintptr(imageLayout), uintptr(unsafe.Pointer(pColor)), uintptr(rangeCount), uintptr(unsafe.Pointer(pRanges)))
}
func (fn PfnCmdClearColorImage) String() string { return "vkCmdClearColorImage" }

//  PfnCmdClearDepthStencilImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdClearDepthStencilImage.html
type PfnCmdClearDepthStencilImage uintptr

func (fn PfnCmdClearDepthStencilImage) Call(commandBuffer CommandBuffer, image Image, imageLayout ImageLayout, pDepthStencil *ClearDepthStencilValue, rangeCount uint32, pRanges *ImageSubresourceRange) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(image), uintptr(imageLayout), uintptr(unsafe.Pointer(pDepthStencil)), uintptr(rangeCount), uintptr(unsafe.Pointer(pRanges)))
}
func (fn PfnCmdClearDepthStencilImage) String() string { return "vkCmdClearDepthStencilImage" }

//  PfnCmdClearAttachments -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdClearAttachments.html
type PfnCmdClearAttachments uintptr

func (fn PfnCmdClearAttachments) Call(commandBuffer CommandBuffer, attachmentCount uint32, pAttachments *ClearAttachment, rectCount uint32, pRects *ClearRect) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(attachmentCount), uintptr(unsafe.Pointer(pAttachments)), uintptr(rectCount), uintptr(unsafe.Pointer(pRects)))
}
func (fn PfnCmdClearAttachments) String() string { return "vkCmdClearAttachments" }

//  PfnCmdResolveImage -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdResolveImage.html
type PfnCmdResolveImage uintptr

func (fn PfnCmdResolveImage) Call(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstImage Image, dstImageLayout ImageLayout, regionCount uint32, pRegions *ImageResolve) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(srcImage), uintptr(srcImageLayout), uintptr(dstImage), uintptr(dstImageLayout), uintptr(regionCount), uintptr(unsafe.Pointer(pRegions)))
}
func (fn PfnCmdResolveImage) String() string { return "vkCmdResolveImage" }

//  PfnCmdSetEvent -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetEvent.html
type PfnCmdSetEvent uintptr

func (fn PfnCmdSetEvent) Call(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(event), uintptr(stageMask))
}
func (fn PfnCmdSetEvent) String() string { return "vkCmdSetEvent" }

//  PfnCmdResetEvent -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdResetEvent.html
type PfnCmdResetEvent uintptr

func (fn PfnCmdResetEvent) Call(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(event), uintptr(stageMask))
}
func (fn PfnCmdResetEvent) String() string { return "vkCmdResetEvent" }

//  PfnCmdWaitEvents -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdWaitEvents.html
type PfnCmdWaitEvents uintptr

func (fn PfnCmdWaitEvents) Call(commandBuffer CommandBuffer, eventCount uint32, pEvents *Event, srcStageMask, dstStageMask PipelineStageFlags, memoryBarrierCount uint32, pMemoryBarriers *MemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers *BufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers *ImageMemoryBarrier) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(eventCount), uintptr(unsafe.Pointer(pEvents)), uintptr(srcStageMask), uintptr(dstStageMask), uintptr(memoryBarrierCount), uintptr(unsafe.Pointer(pMemoryBarriers)), uintptr(bufferMemoryBarrierCount), uintptr(unsafe.Pointer(pBufferMemoryBarriers)), uintptr(imageMemoryBarrierCount), uintptr(unsafe.Pointer(pImageMemoryBarriers)))
}
func (fn PfnCmdWaitEvents) String() string { return "vkCmdWaitEvents" }

//  PfnCmdPipelineBarrier -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdPipelineBarrier.html
type PfnCmdPipelineBarrier uintptr

func (fn PfnCmdPipelineBarrier) Call(commandBuffer CommandBuffer, srcStageMask, dstStageMask PipelineStageFlags, dependencyFlags DependencyFlags, memoryBarrierCount uint32, pMemoryBarriers *MemoryBarrier, bufferMemoryBarrierCount uint32, pBufferMemoryBarriers *BufferMemoryBarrier, imageMemoryBarrierCount uint32, pImageMemoryBarriers *ImageMemoryBarrier) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(srcStageMask), uintptr(dstStageMask), uintptr(dependencyFlags), uintptr(memoryBarrierCount), uintptr(unsafe.Pointer(pMemoryBarriers)), uintptr(bufferMemoryBarrierCount), uintptr(unsafe.Pointer(pBufferMemoryBarriers)), uintptr(imageMemoryBarrierCount), uintptr(unsafe.Pointer(pImageMemoryBarriers)))
}
func (fn PfnCmdPipelineBarrier) String() string { return "vkCmdPipelineBarrier" }

//  PfnCmdBeginQuery -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBeginQuery.html
type PfnCmdBeginQuery uintptr

func (fn PfnCmdBeginQuery) Call(commandBuffer CommandBuffer, queryPool QueryPool, query uint32, flags QueryControlFlags) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(queryPool), uintptr(query), uintptr(flags))
}
func (fn PfnCmdBeginQuery) String() string { return "vkCmdBeginQuery" }

//  PfnCmdEndQuery -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdEndQuery.html
type PfnCmdEndQuery uintptr

func (fn PfnCmdEndQuery) Call(commandBuffer CommandBuffer, queryPool QueryPool, query uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(queryPool), uintptr(query))
}
func (fn PfnCmdEndQuery) String() string { return "vkCmdEndQuery" }

//  PfnCmdResetQueryPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdResetQueryPool.html
type PfnCmdResetQueryPool uintptr

func (fn PfnCmdResetQueryPool) Call(commandBuffer CommandBuffer, queryPool QueryPool, firstQuery, queryCount uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount))
}
func (fn PfnCmdResetQueryPool) String() string { return "vkCmdResetQueryPool" }

//  PfnCmdWriteTimestamp -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdWriteTimestamp.html
type PfnCmdWriteTimestamp uintptr

func (fn PfnCmdWriteTimestamp) Call(commandBuffer CommandBuffer, pipelineStage PipelineStageFlags, queryPool QueryPool, query uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(pipelineStage), uintptr(queryPool), uintptr(query))
}
func (fn PfnCmdWriteTimestamp) String() string { return "vkCmdWriteTimestamp" }

//  PfnCmdCopyQueryPoolResults -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdCopyQueryPoolResults.html
type PfnCmdCopyQueryPoolResults uintptr

func (fn PfnCmdCopyQueryPoolResults) Call(commandBuffer CommandBuffer, queryPool QueryPool, firstQuery, queryCount uint32, dstBuffer Buffer, dstOffset, stride DeviceSize, flags QueryResultFlags) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount), uintptr(dstBuffer), uintptr(dstOffset), uintptr(stride), uintptr(flags))
}
func (fn PfnCmdCopyQueryPoolResults) String() string { return "vkCmdCopyQueryPoolResults" }

//  PfnCmdPushConstants -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdPushConstants.html
type PfnCmdPushConstants uintptr

func (fn PfnCmdPushConstants) Call(commandBuffer CommandBuffer, layout PipelineLayout, stageFlags ShaderStageFlags, offset, size uint32, pValues unsafe.Pointer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(layout), uintptr(stageFlags), uintptr(offset), uintptr(size), uintptr(pValues))
}
func (fn PfnCmdPushConstants) String() string { return "vkCmdPushConstants" }

//  PfnCmdBeginRenderPass -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBeginRenderPass.html
type PfnCmdBeginRenderPass uintptr

func (fn PfnCmdBeginRenderPass) Call(commandBuffer CommandBuffer, pRenderPassBegin *RenderPassBeginInfo, contents SubpassContents) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pRenderPassBegin)), uintptr(contents))
}
func (fn PfnCmdBeginRenderPass) String() string { return "vkCmdBeginRenderPass" }

//  PfnCmdNextSubpass -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdNextSubpass.html
type PfnCmdNextSubpass uintptr

func (fn PfnCmdNextSubpass) Call(commandBuffer CommandBuffer, contents SubpassContents) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(contents))
}
func (fn PfnCmdNextSubpass) String() string { return "vkCmdNextSubpass" }

//  PfnCmdEndRenderPass -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdEndRenderPass.html
type PfnCmdEndRenderPass uintptr

func (fn PfnCmdEndRenderPass) Call(commandBuffer CommandBuffer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer))
}
func (fn PfnCmdEndRenderPass) String() string { return "vkCmdEndRenderPass" }

//  PfnCmdExecuteCommands -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdExecuteCommands.html
type PfnCmdExecuteCommands uintptr

func (fn PfnCmdExecuteCommands) Call(commandBuffer CommandBuffer, commandBufferCount uint32, pCommandBuffers *CommandBuffer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(commandBufferCount), uintptr(unsafe.Pointer(pCommandBuffers)))
}
func (fn PfnCmdExecuteCommands) String() string { return "vkCmdExecuteCommands" }

const VERSION_1_1 = 1

// SamplerYcbcrConversion -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerYcbcrConversion.html
type SamplerYcbcrConversion NonDispatchableHandle

// DescriptorUpdateTemplate -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorUpdateTemplate.html
type DescriptorUpdateTemplate NonDispatchableHandle

const MAX_DEVICE_GROUP_SIZE = 32
const LUID_SIZE = 8

// PointClippingBehavior -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPointClippingBehavior.html
type PointClippingBehavior int32

const (
	POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES           PointClippingBehavior = 0
	POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY     PointClippingBehavior = 1
	POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES_KHR       PointClippingBehavior = POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES
	POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY_KHR PointClippingBehavior = POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY
	POINT_CLIPPING_BEHAVIOR_BEGIN_RANGE               PointClippingBehavior = POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES
	POINT_CLIPPING_BEHAVIOR_END_RANGE                 PointClippingBehavior = POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY
	POINT_CLIPPING_BEHAVIOR_RANGE_SIZE                PointClippingBehavior = (POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY - POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES + 1)
	POINT_CLIPPING_BEHAVIOR_MAX_ENUM                  PointClippingBehavior = 0x7FFFFFFF
)

func (x PointClippingBehavior) String() string {
	switch x {
	case POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES:
		return "POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES"
	case POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY:
		return "POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY"
	case POINT_CLIPPING_BEHAVIOR_MAX_ENUM:
		return "POINT_CLIPPING_BEHAVIOR_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// TessellationDomainOrigin -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkTessellationDomainOrigin.html
type TessellationDomainOrigin int32

const (
	TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT     TessellationDomainOrigin = 0
	TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT     TessellationDomainOrigin = 1
	TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT_KHR TessellationDomainOrigin = TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT
	TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT_KHR TessellationDomainOrigin = TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT
	TESSELLATION_DOMAIN_ORIGIN_BEGIN_RANGE    TessellationDomainOrigin = TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT
	TESSELLATION_DOMAIN_ORIGIN_END_RANGE      TessellationDomainOrigin = TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT
	TESSELLATION_DOMAIN_ORIGIN_RANGE_SIZE     TessellationDomainOrigin = (TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT - TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT + 1)
	TESSELLATION_DOMAIN_ORIGIN_MAX_ENUM       TessellationDomainOrigin = 0x7FFFFFFF
)

func (x TessellationDomainOrigin) String() string {
	switch x {
	case TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT:
		return "TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT"
	case TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT:
		return "TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT"
	case TESSELLATION_DOMAIN_ORIGIN_MAX_ENUM:
		return "TESSELLATION_DOMAIN_ORIGIN_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SamplerYcbcrModelConversion -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerYcbcrModelConversion.html
type SamplerYcbcrModelConversion int32

const (
	SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY       SamplerYcbcrModelConversion = 0
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY     SamplerYcbcrModelConversion = 1
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709          SamplerYcbcrModelConversion = 2
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601          SamplerYcbcrModelConversion = 3
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020         SamplerYcbcrModelConversion = 4
	SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY_KHR   SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY_KHR SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR      SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601_KHR      SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601
	SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020_KHR     SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020
	SAMPLER_YCBCR_MODEL_CONVERSION_BEGIN_RANGE        SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY
	SAMPLER_YCBCR_MODEL_CONVERSION_END_RANGE          SamplerYcbcrModelConversion = SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020
	SAMPLER_YCBCR_MODEL_CONVERSION_RANGE_SIZE         SamplerYcbcrModelConversion = (SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020 - SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY + 1)
	SAMPLER_YCBCR_MODEL_CONVERSION_MAX_ENUM           SamplerYcbcrModelConversion = 0x7FFFFFFF
)

func (x SamplerYcbcrModelConversion) String() string {
	switch x {
	case SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY:
		return "SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY"
	case SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY:
		return "SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY"
	case SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709:
		return "SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709"
	case SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601:
		return "SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601"
	case SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020:
		return "SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020"
	case SAMPLER_YCBCR_MODEL_CONVERSION_MAX_ENUM:
		return "SAMPLER_YCBCR_MODEL_CONVERSION_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SamplerYcbcrRange -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerYcbcrRange.html
type SamplerYcbcrRange int32

const (
	SAMPLER_YCBCR_RANGE_ITU_FULL       SamplerYcbcrRange = 0
	SAMPLER_YCBCR_RANGE_ITU_NARROW     SamplerYcbcrRange = 1
	SAMPLER_YCBCR_RANGE_ITU_FULL_KHR   SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_FULL
	SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_NARROW
	SAMPLER_YCBCR_RANGE_BEGIN_RANGE    SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_FULL
	SAMPLER_YCBCR_RANGE_END_RANGE      SamplerYcbcrRange = SAMPLER_YCBCR_RANGE_ITU_NARROW
	SAMPLER_YCBCR_RANGE_RANGE_SIZE     SamplerYcbcrRange = (SAMPLER_YCBCR_RANGE_ITU_NARROW - SAMPLER_YCBCR_RANGE_ITU_FULL + 1)
	SAMPLER_YCBCR_RANGE_MAX_ENUM       SamplerYcbcrRange = 0x7FFFFFFF
)

func (x SamplerYcbcrRange) String() string {
	switch x {
	case SAMPLER_YCBCR_RANGE_ITU_FULL:
		return "SAMPLER_YCBCR_RANGE_ITU_FULL"
	case SAMPLER_YCBCR_RANGE_ITU_NARROW:
		return "SAMPLER_YCBCR_RANGE_ITU_NARROW"
	case SAMPLER_YCBCR_RANGE_MAX_ENUM:
		return "SAMPLER_YCBCR_RANGE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ChromaLocation -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkChromaLocation.html
type ChromaLocation int32

const (
	CHROMA_LOCATION_COSITED_EVEN     ChromaLocation = 0
	CHROMA_LOCATION_MIDPOINT         ChromaLocation = 1
	CHROMA_LOCATION_COSITED_EVEN_KHR ChromaLocation = CHROMA_LOCATION_COSITED_EVEN
	CHROMA_LOCATION_MIDPOINT_KHR     ChromaLocation = CHROMA_LOCATION_MIDPOINT
	CHROMA_LOCATION_BEGIN_RANGE      ChromaLocation = CHROMA_LOCATION_COSITED_EVEN
	CHROMA_LOCATION_END_RANGE        ChromaLocation = CHROMA_LOCATION_MIDPOINT
	CHROMA_LOCATION_RANGE_SIZE       ChromaLocation = (CHROMA_LOCATION_MIDPOINT - CHROMA_LOCATION_COSITED_EVEN + 1)
	CHROMA_LOCATION_MAX_ENUM         ChromaLocation = 0x7FFFFFFF
)

func (x ChromaLocation) String() string {
	switch x {
	case CHROMA_LOCATION_COSITED_EVEN:
		return "CHROMA_LOCATION_COSITED_EVEN"
	case CHROMA_LOCATION_MIDPOINT:
		return "CHROMA_LOCATION_MIDPOINT"
	case CHROMA_LOCATION_MAX_ENUM:
		return "CHROMA_LOCATION_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DescriptorUpdateTemplateType -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorUpdateTemplateType.html
type DescriptorUpdateTemplateType int32

const (
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET       DescriptorUpdateTemplateType = 0
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR DescriptorUpdateTemplateType = 1
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR   DescriptorUpdateTemplateType = DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_BEGIN_RANGE          DescriptorUpdateTemplateType = DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_END_RANGE            DescriptorUpdateTemplateType = DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_RANGE_SIZE           DescriptorUpdateTemplateType = (DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET - DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET + 1)
	DESCRIPTOR_UPDATE_TEMPLATE_TYPE_MAX_ENUM             DescriptorUpdateTemplateType = 0x7FFFFFFF
)

func (x DescriptorUpdateTemplateType) String() string {
	switch x {
	case DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET:
		return "DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET"
	case DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR:
		return "DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR"
	case DESCRIPTOR_UPDATE_TEMPLATE_TYPE_MAX_ENUM:
		return "DESCRIPTOR_UPDATE_TEMPLATE_TYPE_MAX_ENUM"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SubgroupFeatureFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubgroupFeatureFlags.html
type SubgroupFeatureFlags uint32

const (
	SUBGROUP_FEATURE_BASIC_BIT            SubgroupFeatureFlags = 0x00000001
	SUBGROUP_FEATURE_VOTE_BIT             SubgroupFeatureFlags = 0x00000002
	SUBGROUP_FEATURE_ARITHMETIC_BIT       SubgroupFeatureFlags = 0x00000004
	SUBGROUP_FEATURE_BALLOT_BIT           SubgroupFeatureFlags = 0x00000008
	SUBGROUP_FEATURE_SHUFFLE_BIT          SubgroupFeatureFlags = 0x00000010
	SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT SubgroupFeatureFlags = 0x00000020
	SUBGROUP_FEATURE_CLUSTERED_BIT        SubgroupFeatureFlags = 0x00000040
	SUBGROUP_FEATURE_QUAD_BIT             SubgroupFeatureFlags = 0x00000080
	SUBGROUP_FEATURE_PARTITIONED_BIT_NV   SubgroupFeatureFlags = 0x00000100
	SUBGROUP_FEATURE_FLAG_BITS_MAX_ENUM   SubgroupFeatureFlags = 0x7FFFFFFF
)

func (x SubgroupFeatureFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SubgroupFeatureFlags(1 << i) {
			case SUBGROUP_FEATURE_BASIC_BIT:
				s += "SUBGROUP_FEATURE_BASIC_BIT|"
			case SUBGROUP_FEATURE_VOTE_BIT:
				s += "SUBGROUP_FEATURE_VOTE_BIT|"
			case SUBGROUP_FEATURE_ARITHMETIC_BIT:
				s += "SUBGROUP_FEATURE_ARITHMETIC_BIT|"
			case SUBGROUP_FEATURE_BALLOT_BIT:
				s += "SUBGROUP_FEATURE_BALLOT_BIT|"
			case SUBGROUP_FEATURE_SHUFFLE_BIT:
				s += "SUBGROUP_FEATURE_SHUFFLE_BIT|"
			case SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT:
				s += "SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT|"
			case SUBGROUP_FEATURE_CLUSTERED_BIT:
				s += "SUBGROUP_FEATURE_CLUSTERED_BIT|"
			case SUBGROUP_FEATURE_QUAD_BIT:
				s += "SUBGROUP_FEATURE_QUAD_BIT|"
			case SUBGROUP_FEATURE_PARTITIONED_BIT_NV:
				s += "SUBGROUP_FEATURE_PARTITIONED_BIT_NV|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// PeerMemoryFeatureFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPeerMemoryFeatureFlags.html
type PeerMemoryFeatureFlags uint32

const (
	PEER_MEMORY_FEATURE_COPY_SRC_BIT        PeerMemoryFeatureFlags = 0x00000001
	PEER_MEMORY_FEATURE_COPY_DST_BIT        PeerMemoryFeatureFlags = 0x00000002
	PEER_MEMORY_FEATURE_GENERIC_SRC_BIT     PeerMemoryFeatureFlags = 0x00000004
	PEER_MEMORY_FEATURE_GENERIC_DST_BIT     PeerMemoryFeatureFlags = 0x00000008
	PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR    PeerMemoryFeatureFlags = PEER_MEMORY_FEATURE_COPY_SRC_BIT
	PEER_MEMORY_FEATURE_COPY_DST_BIT_KHR    PeerMemoryFeatureFlags = PEER_MEMORY_FEATURE_COPY_DST_BIT
	PEER_MEMORY_FEATURE_GENERIC_SRC_BIT_KHR PeerMemoryFeatureFlags = PEER_MEMORY_FEATURE_GENERIC_SRC_BIT
	PEER_MEMORY_FEATURE_GENERIC_DST_BIT_KHR PeerMemoryFeatureFlags = PEER_MEMORY_FEATURE_GENERIC_DST_BIT
	PEER_MEMORY_FEATURE_FLAG_BITS_MAX_ENUM  PeerMemoryFeatureFlags = 0x7FFFFFFF
)

func (x PeerMemoryFeatureFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch PeerMemoryFeatureFlags(1 << i) {
			case PEER_MEMORY_FEATURE_COPY_SRC_BIT:
				s += "PEER_MEMORY_FEATURE_COPY_SRC_BIT|"
			case PEER_MEMORY_FEATURE_COPY_DST_BIT:
				s += "PEER_MEMORY_FEATURE_COPY_DST_BIT|"
			case PEER_MEMORY_FEATURE_GENERIC_SRC_BIT:
				s += "PEER_MEMORY_FEATURE_GENERIC_SRC_BIT|"
			case PEER_MEMORY_FEATURE_GENERIC_DST_BIT:
				s += "PEER_MEMORY_FEATURE_GENERIC_DST_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// MemoryAllocateFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryAllocateFlags.html
type MemoryAllocateFlags uint32

const (
	MEMORY_ALLOCATE_DEVICE_MASK_BIT     MemoryAllocateFlags = 0x00000001
	MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR MemoryAllocateFlags = MEMORY_ALLOCATE_DEVICE_MASK_BIT
	MEMORY_ALLOCATE_FLAG_BITS_MAX_ENUM  MemoryAllocateFlags = 0x7FFFFFFF
)

func (x MemoryAllocateFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch MemoryAllocateFlags(1 << i) {
			case MEMORY_ALLOCATE_DEVICE_MASK_BIT:
				s += "MEMORY_ALLOCATE_DEVICE_MASK_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type CommandPoolTrimFlags uint32                // reserved
type DescriptorUpdateTemplateCreateFlags uint32 // reserved
// ExternalMemoryHandleTypeFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryHandleTypeFlags.html
type ExternalMemoryHandleTypeFlags uint32

const (
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT                       ExternalMemoryHandleTypeFlags = 0x00000001
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT                    ExternalMemoryHandleTypeFlags = 0x00000002
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT                ExternalMemoryHandleTypeFlags = 0x00000004
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT                   ExternalMemoryHandleTypeFlags = 0x00000008
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT               ExternalMemoryHandleTypeFlags = 0x00000010
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT                      ExternalMemoryHandleTypeFlags = 0x00000020
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT                  ExternalMemoryHandleTypeFlags = 0x00000040
	EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT                     ExternalMemoryHandleTypeFlags = 0x00000200
	EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID ExternalMemoryHandleTypeFlags = 0x00000400
	EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT             ExternalMemoryHandleTypeFlags = 0x00000080
	EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT  ExternalMemoryHandleTypeFlags = 0x00000100
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR                   ExternalMemoryHandleTypeFlags = EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR                ExternalMemoryHandleTypeFlags = EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR            ExternalMemoryHandleTypeFlags = EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR               ExternalMemoryHandleTypeFlags = EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR           ExternalMemoryHandleTypeFlags = EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR                  ExternalMemoryHandleTypeFlags = EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT_KHR              ExternalMemoryHandleTypeFlags = EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT
	EXTERNAL_MEMORY_HANDLE_TYPE_FLAG_BITS_MAX_ENUM                  ExternalMemoryHandleTypeFlags = 0x7FFFFFFF
)

func (x ExternalMemoryHandleTypeFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalMemoryHandleTypeFlags(1 << i) {
			case EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ExternalMemoryFeatureFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryFeatureFlags.html
type ExternalMemoryFeatureFlags uint32

const (
	EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT     ExternalMemoryFeatureFlags = 0x00000001
	EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT         ExternalMemoryFeatureFlags = 0x00000002
	EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT         ExternalMemoryFeatureFlags = 0x00000004
	EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_KHR ExternalMemoryFeatureFlags = EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT
	EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_KHR     ExternalMemoryFeatureFlags = EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT
	EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR     ExternalMemoryFeatureFlags = EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT
	EXTERNAL_MEMORY_FEATURE_FLAG_BITS_MAX_ENUM     ExternalMemoryFeatureFlags = 0x7FFFFFFF
)

func (x ExternalMemoryFeatureFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalMemoryFeatureFlags(1 << i) {
			case EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT:
				s += "EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT|"
			case EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT:
				s += "EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT|"
			case EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT:
				s += "EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ExternalFenceHandleTypeFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalFenceHandleTypeFlags.html
type ExternalFenceHandleTypeFlags uint32

const (
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT            ExternalFenceHandleTypeFlags = 0x00000001
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT         ExternalFenceHandleTypeFlags = 0x00000002
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT     ExternalFenceHandleTypeFlags = 0x00000004
	EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT              ExternalFenceHandleTypeFlags = 0x00000008
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR        ExternalFenceHandleTypeFlags = EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR     ExternalFenceHandleTypeFlags = EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR ExternalFenceHandleTypeFlags = EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR          ExternalFenceHandleTypeFlags = EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT
	EXTERNAL_FENCE_HANDLE_TYPE_FLAG_BITS_MAX_ENUM       ExternalFenceHandleTypeFlags = 0x7FFFFFFF
)

func (x ExternalFenceHandleTypeFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalFenceHandleTypeFlags(1 << i) {
			case EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT:
				s += "EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT|"
			case EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT:
				s += "EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT|"
			case EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT:
				s += "EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT|"
			case EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT:
				s += "EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ExternalFenceFeatureFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalFenceFeatureFlags.html
type ExternalFenceFeatureFlags uint32

const (
	EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT     ExternalFenceFeatureFlags = 0x00000001
	EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT     ExternalFenceFeatureFlags = 0x00000002
	EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT_KHR ExternalFenceFeatureFlags = EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT
	EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT_KHR ExternalFenceFeatureFlags = EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT
	EXTERNAL_FENCE_FEATURE_FLAG_BITS_MAX_ENUM ExternalFenceFeatureFlags = 0x7FFFFFFF
)

func (x ExternalFenceFeatureFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalFenceFeatureFlags(1 << i) {
			case EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT:
				s += "EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT|"
			case EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT:
				s += "EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// FenceImportFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFenceImportFlags.html
type FenceImportFlags uint32

const (
	FENCE_IMPORT_TEMPORARY_BIT      FenceImportFlags = 0x00000001
	FENCE_IMPORT_TEMPORARY_BIT_KHR  FenceImportFlags = FENCE_IMPORT_TEMPORARY_BIT
	FENCE_IMPORT_FLAG_BITS_MAX_ENUM FenceImportFlags = 0x7FFFFFFF
)

func (x FenceImportFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch FenceImportFlags(1 << i) {
			case FENCE_IMPORT_TEMPORARY_BIT:
				s += "FENCE_IMPORT_TEMPORARY_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SemaphoreImportFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSemaphoreImportFlags.html
type SemaphoreImportFlags uint32

const (
	SEMAPHORE_IMPORT_TEMPORARY_BIT      SemaphoreImportFlags = 0x00000001
	SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR  SemaphoreImportFlags = SEMAPHORE_IMPORT_TEMPORARY_BIT
	SEMAPHORE_IMPORT_FLAG_BITS_MAX_ENUM SemaphoreImportFlags = 0x7FFFFFFF
)

func (x SemaphoreImportFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SemaphoreImportFlags(1 << i) {
			case SEMAPHORE_IMPORT_TEMPORARY_BIT:
				s += "SEMAPHORE_IMPORT_TEMPORARY_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ExternalSemaphoreHandleTypeFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalSemaphoreHandleTypeFlags.html
type ExternalSemaphoreHandleTypeFlags uint32

const (
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT            ExternalSemaphoreHandleTypeFlags = 0x00000001
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT         ExternalSemaphoreHandleTypeFlags = 0x00000002
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT     ExternalSemaphoreHandleTypeFlags = 0x00000004
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT          ExternalSemaphoreHandleTypeFlags = 0x00000008
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT              ExternalSemaphoreHandleTypeFlags = 0x00000010
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR        ExternalSemaphoreHandleTypeFlags = EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR     ExternalSemaphoreHandleTypeFlags = EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR ExternalSemaphoreHandleTypeFlags = EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR      ExternalSemaphoreHandleTypeFlags = EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR          ExternalSemaphoreHandleTypeFlags = EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT
	EXTERNAL_SEMAPHORE_HANDLE_TYPE_FLAG_BITS_MAX_ENUM       ExternalSemaphoreHandleTypeFlags = 0x7FFFFFFF
)

func (x ExternalSemaphoreHandleTypeFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalSemaphoreHandleTypeFlags(1 << i) {
			case EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT:
				s += "EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT|"
			case EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT:
				s += "EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT|"
			case EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT:
				s += "EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT|"
			case EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT:
				s += "EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT|"
			case EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT:
				s += "EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ExternalSemaphoreFeatureFlags -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalSemaphoreFeatureFlags.html
type ExternalSemaphoreFeatureFlags uint32

const (
	EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT     ExternalSemaphoreFeatureFlags = 0x00000001
	EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT     ExternalSemaphoreFeatureFlags = 0x00000002
	EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT_KHR ExternalSemaphoreFeatureFlags = EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT
	EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR ExternalSemaphoreFeatureFlags = EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT
	EXTERNAL_SEMAPHORE_FEATURE_FLAG_BITS_MAX_ENUM ExternalSemaphoreFeatureFlags = 0x7FFFFFFF
)

func (x ExternalSemaphoreFeatureFlags) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalSemaphoreFeatureFlags(1 << i) {
			case EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT:
				s += "EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT|"
			case EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT:
				s += "EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// PhysicalDeviceSubgroupProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSubgroupProperties.html
type PhysicalDeviceSubgroupProperties struct {
	SType                     StructureType
	PNext                     unsafe.Pointer
	SubgroupSize              uint32
	SupportedStages           ShaderStageFlags
	SupportedOperations       SubgroupFeatureFlags
	QuadOperationsInAllStages Bool32
}

func NewPhysicalDeviceSubgroupProperties() *PhysicalDeviceSubgroupProperties {
	p := (*PhysicalDeviceSubgroupProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSubgroupProperties)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES
	return p
}
func (p *PhysicalDeviceSubgroupProperties) Free() { MemFree(unsafe.Pointer(p)) }

// BindBufferMemoryInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindBufferMemoryInfo.html
type BindBufferMemoryInfo struct {
	SType        StructureType
	PNext        unsafe.Pointer
	Buffer       Buffer
	Memory       DeviceMemory
	MemoryOffset DeviceSize
}

func NewBindBufferMemoryInfo() *BindBufferMemoryInfo {
	p := (*BindBufferMemoryInfo)(MemAlloc(unsafe.Sizeof(*(*BindBufferMemoryInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO
	return p
}
func (p *BindBufferMemoryInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BindImageMemoryInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindImageMemoryInfo.html
type BindImageMemoryInfo struct {
	SType        StructureType
	PNext        unsafe.Pointer
	Image        Image
	Memory       DeviceMemory
	MemoryOffset DeviceSize
}

func NewBindImageMemoryInfo() *BindImageMemoryInfo {
	p := (*BindImageMemoryInfo)(MemAlloc(unsafe.Sizeof(*(*BindImageMemoryInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO
	return p
}
func (p *BindImageMemoryInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDevice16BitStorageFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDevice16BitStorageFeatures.html
type PhysicalDevice16BitStorageFeatures struct {
	SType                              StructureType
	PNext                              unsafe.Pointer
	StorageBuffer16BitAccess           Bool32
	UniformAndStorageBuffer16BitAccess Bool32
	StoragePushConstant16              Bool32
	StorageInputOutput16               Bool32
}

func NewPhysicalDevice16BitStorageFeatures() *PhysicalDevice16BitStorageFeatures {
	p := (*PhysicalDevice16BitStorageFeatures)(MemAlloc(unsafe.Sizeof(*(*PhysicalDevice16BitStorageFeatures)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES
	return p
}
func (p *PhysicalDevice16BitStorageFeatures) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryDedicatedRequirements -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryDedicatedRequirements.html
type MemoryDedicatedRequirements struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	PrefersDedicatedAllocation  Bool32
	RequiresDedicatedAllocation Bool32
}

func NewMemoryDedicatedRequirements() *MemoryDedicatedRequirements {
	p := (*MemoryDedicatedRequirements)(MemAlloc(unsafe.Sizeof(*(*MemoryDedicatedRequirements)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS
	return p
}
func (p *MemoryDedicatedRequirements) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryDedicatedAllocateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryDedicatedAllocateInfo.html
type MemoryDedicatedAllocateInfo struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Image  Image
	Buffer Buffer
}

func NewMemoryDedicatedAllocateInfo() *MemoryDedicatedAllocateInfo {
	p := (*MemoryDedicatedAllocateInfo)(MemAlloc(unsafe.Sizeof(*(*MemoryDedicatedAllocateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO
	return p
}
func (p *MemoryDedicatedAllocateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryAllocateFlagsInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryAllocateFlagsInfo.html
type MemoryAllocateFlagsInfo struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Flags      MemoryAllocateFlags
	DeviceMask uint32
}

func NewMemoryAllocateFlagsInfo() *MemoryAllocateFlagsInfo {
	p := (*MemoryAllocateFlagsInfo)(MemAlloc(unsafe.Sizeof(*(*MemoryAllocateFlagsInfo)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO
	return p
}
func (p *MemoryAllocateFlagsInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupRenderPassBeginInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html
type DeviceGroupRenderPassBeginInfo struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	DeviceMask            uint32
	DeviceRenderAreaCount uint32
	PDeviceRenderAreas    *Rect2D
}

func NewDeviceGroupRenderPassBeginInfo() *DeviceGroupRenderPassBeginInfo {
	p := (*DeviceGroupRenderPassBeginInfo)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupRenderPassBeginInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO
	return p
}
func (p *DeviceGroupRenderPassBeginInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupCommandBufferBeginInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupCommandBufferBeginInfo.html
type DeviceGroupCommandBufferBeginInfo struct {
	SType      StructureType
	PNext      unsafe.Pointer
	DeviceMask uint32
}

func NewDeviceGroupCommandBufferBeginInfo() *DeviceGroupCommandBufferBeginInfo {
	p := (*DeviceGroupCommandBufferBeginInfo)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupCommandBufferBeginInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO
	return p
}
func (p *DeviceGroupCommandBufferBeginInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupSubmitInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupSubmitInfo.html
type DeviceGroupSubmitInfo struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	WaitSemaphoreCount            uint32
	PWaitSemaphoreDeviceIndices   *uint32
	CommandBufferCount            uint32
	PCommandBufferDeviceMasks     *uint32
	SignalSemaphoreCount          uint32
	PSignalSemaphoreDeviceIndices *uint32
}

func NewDeviceGroupSubmitInfo() *DeviceGroupSubmitInfo {
	p := (*DeviceGroupSubmitInfo)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupSubmitInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO
	return p
}
func (p *DeviceGroupSubmitInfo) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupBindSparseInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupBindSparseInfo.html
type DeviceGroupBindSparseInfo struct {
	SType               StructureType
	PNext               unsafe.Pointer
	ResourceDeviceIndex uint32
	MemoryDeviceIndex   uint32
}

func NewDeviceGroupBindSparseInfo() *DeviceGroupBindSparseInfo {
	p := (*DeviceGroupBindSparseInfo)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupBindSparseInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO
	return p
}
func (p *DeviceGroupBindSparseInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BindBufferMemoryDeviceGroupInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindBufferMemoryDeviceGroupInfo.html
type BindBufferMemoryDeviceGroupInfo struct {
	SType            StructureType
	PNext            unsafe.Pointer
	DeviceIndexCount uint32
	PDeviceIndices   *uint32
}

func NewBindBufferMemoryDeviceGroupInfo() *BindBufferMemoryDeviceGroupInfo {
	p := (*BindBufferMemoryDeviceGroupInfo)(MemAlloc(unsafe.Sizeof(*(*BindBufferMemoryDeviceGroupInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO
	return p
}
func (p *BindBufferMemoryDeviceGroupInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BindImageMemoryDeviceGroupInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html
type BindImageMemoryDeviceGroupInfo struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	DeviceIndexCount             uint32
	PDeviceIndices               *uint32
	SplitInstanceBindRegionCount uint32
	PSplitInstanceBindRegions    *Rect2D
}

func NewBindImageMemoryDeviceGroupInfo() *BindImageMemoryDeviceGroupInfo {
	p := (*BindImageMemoryDeviceGroupInfo)(MemAlloc(unsafe.Sizeof(*(*BindImageMemoryDeviceGroupInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO
	return p
}
func (p *BindImageMemoryDeviceGroupInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceGroupProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceGroupProperties.html
type PhysicalDeviceGroupProperties struct {
	SType               StructureType
	PNext               unsafe.Pointer
	PhysicalDeviceCount uint32
	PhysicalDevices     [MAX_DEVICE_GROUP_SIZE]PhysicalDevice
	SubsetAllocation    Bool32
}

func NewPhysicalDeviceGroupProperties() *PhysicalDeviceGroupProperties {
	p := (*PhysicalDeviceGroupProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceGroupProperties)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES
	return p
}
func (p *PhysicalDeviceGroupProperties) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupDeviceCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html
type DeviceGroupDeviceCreateInfo struct {
	SType               StructureType
	PNext               unsafe.Pointer
	PhysicalDeviceCount uint32
	PPhysicalDevices    *PhysicalDevice
}

func NewDeviceGroupDeviceCreateInfo() *DeviceGroupDeviceCreateInfo {
	p := (*DeviceGroupDeviceCreateInfo)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupDeviceCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO
	return p
}
func (p *DeviceGroupDeviceCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BufferMemoryRequirementsInfo2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferMemoryRequirementsInfo2.html
type BufferMemoryRequirementsInfo2 struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Buffer Buffer
}

func NewBufferMemoryRequirementsInfo2() *BufferMemoryRequirementsInfo2 {
	p := (*BufferMemoryRequirementsInfo2)(MemAlloc(unsafe.Sizeof(*(*BufferMemoryRequirementsInfo2)(nil))))
	p.SType = STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2
	return p
}
func (p *BufferMemoryRequirementsInfo2) Free() { MemFree(unsafe.Pointer(p)) }

// ImageMemoryRequirementsInfo2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageMemoryRequirementsInfo2.html
type ImageMemoryRequirementsInfo2 struct {
	SType StructureType
	PNext unsafe.Pointer
	Image Image
}

func NewImageMemoryRequirementsInfo2() *ImageMemoryRequirementsInfo2 {
	p := (*ImageMemoryRequirementsInfo2)(MemAlloc(unsafe.Sizeof(*(*ImageMemoryRequirementsInfo2)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2
	return p
}
func (p *ImageMemoryRequirementsInfo2) Free() { MemFree(unsafe.Pointer(p)) }

// ImageSparseMemoryRequirementsInfo2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageSparseMemoryRequirementsInfo2.html
type ImageSparseMemoryRequirementsInfo2 struct {
	SType StructureType
	PNext unsafe.Pointer
	Image Image
}

func NewImageSparseMemoryRequirementsInfo2() *ImageSparseMemoryRequirementsInfo2 {
	p := (*ImageSparseMemoryRequirementsInfo2)(MemAlloc(unsafe.Sizeof(*(*ImageSparseMemoryRequirementsInfo2)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2
	return p
}
func (p *ImageSparseMemoryRequirementsInfo2) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryRequirements2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryRequirements2.html
type MemoryRequirements2 struct {
	SType              StructureType
	PNext              unsafe.Pointer
	MemoryRequirements MemoryRequirements
}

func NewMemoryRequirements2() *MemoryRequirements2 {
	p := (*MemoryRequirements2)(MemAlloc(unsafe.Sizeof(*(*MemoryRequirements2)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
	return p
}
func (p *MemoryRequirements2) Free() { MemFree(unsafe.Pointer(p)) }

type MemoryRequirements2KHR = MemoryRequirements2

// SparseImageMemoryRequirements2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageMemoryRequirements2.html
type SparseImageMemoryRequirements2 struct {
	SType              StructureType
	PNext              unsafe.Pointer
	MemoryRequirements SparseImageMemoryRequirements
}

func NewSparseImageMemoryRequirements2() *SparseImageMemoryRequirements2 {
	p := (*SparseImageMemoryRequirements2)(MemAlloc(unsafe.Sizeof(*(*SparseImageMemoryRequirements2)(nil))))
	p.SType = STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2
	return p
}
func (p *SparseImageMemoryRequirements2) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceFeatures2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceFeatures2.html
type PhysicalDeviceFeatures2 struct {
	SType    StructureType
	PNext    unsafe.Pointer
	Features PhysicalDeviceFeatures
}

func NewPhysicalDeviceFeatures2() *PhysicalDeviceFeatures2 {
	p := (*PhysicalDeviceFeatures2)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceFeatures2)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
	return p
}
func (p *PhysicalDeviceFeatures2) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceProperties2.html
type PhysicalDeviceProperties2 struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Properties PhysicalDeviceProperties
}

func NewPhysicalDeviceProperties2() *PhysicalDeviceProperties2 {
	p := (*PhysicalDeviceProperties2)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceProperties2)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
	return p
}
func (p *PhysicalDeviceProperties2) Free() { MemFree(unsafe.Pointer(p)) }

// FormatProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFormatProperties2.html
type FormatProperties2 struct {
	SType            StructureType
	PNext            unsafe.Pointer
	FormatProperties FormatProperties
}

func NewFormatProperties2() *FormatProperties2 {
	p := (*FormatProperties2)(MemAlloc(unsafe.Sizeof(*(*FormatProperties2)(nil))))
	p.SType = STRUCTURE_TYPE_FORMAT_PROPERTIES_2
	return p
}
func (p *FormatProperties2) Free() { MemFree(unsafe.Pointer(p)) }

// ImageFormatProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageFormatProperties2.html
type ImageFormatProperties2 struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	ImageFormatProperties ImageFormatProperties
}

func NewImageFormatProperties2() *ImageFormatProperties2 {
	p := (*ImageFormatProperties2)(MemAlloc(unsafe.Sizeof(*(*ImageFormatProperties2)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2
	return p
}
func (p *ImageFormatProperties2) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceImageFormatInfo2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html
type PhysicalDeviceImageFormatInfo2 struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Format Format
	Type   ImageType
	Tiling ImageTiling
	Usage  ImageUsageFlags
	Flags  ImageCreateFlags
}

func NewPhysicalDeviceImageFormatInfo2() *PhysicalDeviceImageFormatInfo2 {
	p := (*PhysicalDeviceImageFormatInfo2)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceImageFormatInfo2)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2
	return p
}
func (p *PhysicalDeviceImageFormatInfo2) Free() { MemFree(unsafe.Pointer(p)) }

// QueueFamilyProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueueFamilyProperties2.html
type QueueFamilyProperties2 struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	QueueFamilyProperties QueueFamilyProperties
}

func NewQueueFamilyProperties2() *QueueFamilyProperties2 {
	p := (*QueueFamilyProperties2)(MemAlloc(unsafe.Sizeof(*(*QueueFamilyProperties2)(nil))))
	p.SType = STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2
	return p
}
func (p *QueueFamilyProperties2) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceMemoryProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html
type PhysicalDeviceMemoryProperties2 struct {
	SType            StructureType
	PNext            unsafe.Pointer
	MemoryProperties PhysicalDeviceMemoryProperties
}

func NewPhysicalDeviceMemoryProperties2() *PhysicalDeviceMemoryProperties2 {
	p := (*PhysicalDeviceMemoryProperties2)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMemoryProperties2)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2
	return p
}
func (p *PhysicalDeviceMemoryProperties2) Free() { MemFree(unsafe.Pointer(p)) }

// SparseImageFormatProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSparseImageFormatProperties2.html
type SparseImageFormatProperties2 struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Properties SparseImageFormatProperties
}

func NewSparseImageFormatProperties2() *SparseImageFormatProperties2 {
	p := (*SparseImageFormatProperties2)(MemAlloc(unsafe.Sizeof(*(*SparseImageFormatProperties2)(nil))))
	p.SType = STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2
	return p
}
func (p *SparseImageFormatProperties2) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceSparseImageFormatInfo2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html
type PhysicalDeviceSparseImageFormatInfo2 struct {
	SType   StructureType
	PNext   unsafe.Pointer
	Format  Format
	Type    ImageType
	Samples SampleCountFlags
	Usage   ImageUsageFlags
	Tiling  ImageTiling
}

func NewPhysicalDeviceSparseImageFormatInfo2() *PhysicalDeviceSparseImageFormatInfo2 {
	p := (*PhysicalDeviceSparseImageFormatInfo2)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSparseImageFormatInfo2)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2
	return p
}
func (p *PhysicalDeviceSparseImageFormatInfo2) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDevicePointClippingProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDevicePointClippingProperties.html
type PhysicalDevicePointClippingProperties struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	PointClippingBehavior PointClippingBehavior
}

func NewPhysicalDevicePointClippingProperties() *PhysicalDevicePointClippingProperties {
	p := (*PhysicalDevicePointClippingProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDevicePointClippingProperties)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES
	return p
}
func (p *PhysicalDevicePointClippingProperties) Free() { MemFree(unsafe.Pointer(p)) }

// InputAttachmentAspectReference -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkInputAttachmentAspectReference.html
type InputAttachmentAspectReference struct {
	Subpass              uint32
	InputAttachmentIndex uint32
	AspectMask           ImageAspectFlags
}

func NewInputAttachmentAspectReference() *InputAttachmentAspectReference {
	return (*InputAttachmentAspectReference)(MemAlloc(unsafe.Sizeof(*(*InputAttachmentAspectReference)(nil))))
}
func (p *InputAttachmentAspectReference) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassInputAttachmentAspectCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html
type RenderPassInputAttachmentAspectCreateInfo struct {
	SType                StructureType
	PNext                unsafe.Pointer
	AspectReferenceCount uint32
	PAspectReferences    *InputAttachmentAspectReference
}

func NewRenderPassInputAttachmentAspectCreateInfo() *RenderPassInputAttachmentAspectCreateInfo {
	p := (*RenderPassInputAttachmentAspectCreateInfo)(MemAlloc(unsafe.Sizeof(*(*RenderPassInputAttachmentAspectCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO
	return p
}
func (p *RenderPassInputAttachmentAspectCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ImageViewUsageCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageViewUsageCreateInfo.html
type ImageViewUsageCreateInfo struct {
	SType StructureType
	PNext unsafe.Pointer
	Usage ImageUsageFlags
}

func NewImageViewUsageCreateInfo() *ImageViewUsageCreateInfo {
	p := (*ImageViewUsageCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ImageViewUsageCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO
	return p
}
func (p *ImageViewUsageCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineTessellationDomainOriginStateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html
type PipelineTessellationDomainOriginStateCreateInfo struct {
	SType        StructureType
	PNext        unsafe.Pointer
	DomainOrigin TessellationDomainOrigin
}

func NewPipelineTessellationDomainOriginStateCreateInfo() *PipelineTessellationDomainOriginStateCreateInfo {
	p := (*PipelineTessellationDomainOriginStateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*PipelineTessellationDomainOriginStateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO
	return p
}
func (p *PipelineTessellationDomainOriginStateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassMultiviewCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassMultiviewCreateInfo.html
type RenderPassMultiviewCreateInfo struct {
	SType                StructureType
	PNext                unsafe.Pointer
	SubpassCount         uint32
	PViewMasks           *uint32
	DependencyCount      uint32
	PViewOffsets         *int32
	CorrelationMaskCount uint32
	PCorrelationMasks    *uint32
}

func NewRenderPassMultiviewCreateInfo() *RenderPassMultiviewCreateInfo {
	p := (*RenderPassMultiviewCreateInfo)(MemAlloc(unsafe.Sizeof(*(*RenderPassMultiviewCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO
	return p
}
func (p *RenderPassMultiviewCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceMultiviewFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMultiviewFeatures.html
type PhysicalDeviceMultiviewFeatures struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	Multiview                   Bool32
	MultiviewGeometryShader     Bool32
	MultiviewTessellationShader Bool32
}

func NewPhysicalDeviceMultiviewFeatures() *PhysicalDeviceMultiviewFeatures {
	p := (*PhysicalDeviceMultiviewFeatures)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMultiviewFeatures)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES
	return p
}
func (p *PhysicalDeviceMultiviewFeatures) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceMultiviewProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMultiviewProperties.html
type PhysicalDeviceMultiviewProperties struct {
	SType                     StructureType
	PNext                     unsafe.Pointer
	MaxMultiviewViewCount     uint32
	MaxMultiviewInstanceIndex uint32
}

func NewPhysicalDeviceMultiviewProperties() *PhysicalDeviceMultiviewProperties {
	p := (*PhysicalDeviceMultiviewProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMultiviewProperties)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES
	return p
}
func (p *PhysicalDeviceMultiviewProperties) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceVariablePointersFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceVariablePointersFeatures.html
type PhysicalDeviceVariablePointersFeatures struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	VariablePointersStorageBuffer Bool32
	VariablePointers              Bool32
}

func NewPhysicalDeviceVariablePointersFeatures() *PhysicalDeviceVariablePointersFeatures {
	p := (*PhysicalDeviceVariablePointersFeatures)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceVariablePointersFeatures)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES
	return p
}
func (p *PhysicalDeviceVariablePointersFeatures) Free() { MemFree(unsafe.Pointer(p)) }

type PhysicalDeviceVariablePointerFeatures = PhysicalDeviceVariablePointersFeatures

// PhysicalDeviceProtectedMemoryFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceProtectedMemoryFeatures.html
type PhysicalDeviceProtectedMemoryFeatures struct {
	SType           StructureType
	PNext           unsafe.Pointer
	ProtectedMemory Bool32
}

func NewPhysicalDeviceProtectedMemoryFeatures() *PhysicalDeviceProtectedMemoryFeatures {
	p := (*PhysicalDeviceProtectedMemoryFeatures)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceProtectedMemoryFeatures)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES
	return p
}
func (p *PhysicalDeviceProtectedMemoryFeatures) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceProtectedMemoryProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceProtectedMemoryProperties.html
type PhysicalDeviceProtectedMemoryProperties struct {
	SType            StructureType
	PNext            unsafe.Pointer
	ProtectedNoFault Bool32
}

func NewPhysicalDeviceProtectedMemoryProperties() *PhysicalDeviceProtectedMemoryProperties {
	p := (*PhysicalDeviceProtectedMemoryProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceProtectedMemoryProperties)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES
	return p
}
func (p *PhysicalDeviceProtectedMemoryProperties) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceQueueInfo2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceQueueInfo2.html
type DeviceQueueInfo2 struct {
	SType            StructureType
	PNext            unsafe.Pointer
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueIndex       uint32
}

func NewDeviceQueueInfo2() *DeviceQueueInfo2 {
	p := (*DeviceQueueInfo2)(MemAlloc(unsafe.Sizeof(*(*DeviceQueueInfo2)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2
	return p
}
func (p *DeviceQueueInfo2) Free() { MemFree(unsafe.Pointer(p)) }

// ProtectedSubmitInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkProtectedSubmitInfo.html
type ProtectedSubmitInfo struct {
	SType           StructureType
	PNext           unsafe.Pointer
	ProtectedSubmit Bool32
}

func NewProtectedSubmitInfo() *ProtectedSubmitInfo {
	p := (*ProtectedSubmitInfo)(MemAlloc(unsafe.Sizeof(*(*ProtectedSubmitInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO
	return p
}
func (p *ProtectedSubmitInfo) Free() { MemFree(unsafe.Pointer(p)) }

// SamplerYcbcrConversionCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html
type SamplerYcbcrConversionCreateInfo struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	Format                      Format
	YcbcrModel                  SamplerYcbcrModelConversion
	YcbcrRange                  SamplerYcbcrRange
	Components                  ComponentMapping
	XChromaOffset               ChromaLocation
	YChromaOffset               ChromaLocation
	ChromaFilter                Filter
	ForceExplicitReconstruction Bool32
}

func NewSamplerYcbcrConversionCreateInfo() *SamplerYcbcrConversionCreateInfo {
	p := (*SamplerYcbcrConversionCreateInfo)(MemAlloc(unsafe.Sizeof(*(*SamplerYcbcrConversionCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO
	return p
}
func (p *SamplerYcbcrConversionCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// SamplerYcbcrConversionInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerYcbcrConversionInfo.html
type SamplerYcbcrConversionInfo struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Conversion SamplerYcbcrConversion
}

func NewSamplerYcbcrConversionInfo() *SamplerYcbcrConversionInfo {
	p := (*SamplerYcbcrConversionInfo)(MemAlloc(unsafe.Sizeof(*(*SamplerYcbcrConversionInfo)(nil))))
	p.SType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO
	return p
}
func (p *SamplerYcbcrConversionInfo) Free() { MemFree(unsafe.Pointer(p)) }

// BindImagePlaneMemoryInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindImagePlaneMemoryInfo.html
type BindImagePlaneMemoryInfo struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PlaneAspect ImageAspectFlags
}

func NewBindImagePlaneMemoryInfo() *BindImagePlaneMemoryInfo {
	p := (*BindImagePlaneMemoryInfo)(MemAlloc(unsafe.Sizeof(*(*BindImagePlaneMemoryInfo)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO
	return p
}
func (p *BindImagePlaneMemoryInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ImagePlaneMemoryRequirementsInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html
type ImagePlaneMemoryRequirementsInfo struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PlaneAspect ImageAspectFlags
}

func NewImagePlaneMemoryRequirementsInfo() *ImagePlaneMemoryRequirementsInfo {
	p := (*ImagePlaneMemoryRequirementsInfo)(MemAlloc(unsafe.Sizeof(*(*ImagePlaneMemoryRequirementsInfo)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO
	return p
}
func (p *ImagePlaneMemoryRequirementsInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceSamplerYcbcrConversionFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeatures.html
type PhysicalDeviceSamplerYcbcrConversionFeatures struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	SamplerYcbcrConversion Bool32
}

func NewPhysicalDeviceSamplerYcbcrConversionFeatures() *PhysicalDeviceSamplerYcbcrConversionFeatures {
	p := (*PhysicalDeviceSamplerYcbcrConversionFeatures)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSamplerYcbcrConversionFeatures)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES
	return p
}
func (p *PhysicalDeviceSamplerYcbcrConversionFeatures) Free() { MemFree(unsafe.Pointer(p)) }

// SamplerYcbcrConversionImageFormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerYcbcrConversionImageFormatProperties.html
type SamplerYcbcrConversionImageFormatProperties struct {
	SType                               StructureType
	PNext                               unsafe.Pointer
	CombinedImageSamplerDescriptorCount uint32
}

func NewSamplerYcbcrConversionImageFormatProperties() *SamplerYcbcrConversionImageFormatProperties {
	p := (*SamplerYcbcrConversionImageFormatProperties)(MemAlloc(unsafe.Sizeof(*(*SamplerYcbcrConversionImageFormatProperties)(nil))))
	p.SType = STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES
	return p
}
func (p *SamplerYcbcrConversionImageFormatProperties) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorUpdateTemplateEntry -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorUpdateTemplateEntry.html
type DescriptorUpdateTemplateEntry struct {
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  DescriptorType
	Offset          uintptr
	Stride          uintptr
}

func NewDescriptorUpdateTemplateEntry() *DescriptorUpdateTemplateEntry {
	return (*DescriptorUpdateTemplateEntry)(MemAlloc(unsafe.Sizeof(*(*DescriptorUpdateTemplateEntry)(nil))))
}
func (p *DescriptorUpdateTemplateEntry) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorUpdateTemplateCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html
type DescriptorUpdateTemplateCreateInfo struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	Flags                      DescriptorUpdateTemplateCreateFlags
	DescriptorUpdateEntryCount uint32
	PDescriptorUpdateEntries   *DescriptorUpdateTemplateEntry
	TemplateType               DescriptorUpdateTemplateType
	DescriptorSetLayout        DescriptorSetLayout
	PipelineBindPoint          PipelineBindPoint
	PipelineLayout             PipelineLayout
	Set                        uint32
}

func NewDescriptorUpdateTemplateCreateInfo() *DescriptorUpdateTemplateCreateInfo {
	p := (*DescriptorUpdateTemplateCreateInfo)(MemAlloc(unsafe.Sizeof(*(*DescriptorUpdateTemplateCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO
	return p
}
func (p *DescriptorUpdateTemplateCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExternalMemoryProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryProperties.html
type ExternalMemoryProperties struct {
	ExternalMemoryFeatures        ExternalMemoryFeatureFlags
	ExportFromImportedHandleTypes ExternalMemoryHandleTypeFlags
	CompatibleHandleTypes         ExternalMemoryHandleTypeFlags
}

func NewExternalMemoryProperties() *ExternalMemoryProperties {
	return (*ExternalMemoryProperties)(MemAlloc(unsafe.Sizeof(*(*ExternalMemoryProperties)(nil))))
}
func (p *ExternalMemoryProperties) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceExternalImageFormatInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html
type PhysicalDeviceExternalImageFormatInfo struct {
	SType      StructureType
	PNext      unsafe.Pointer
	HandleType ExternalMemoryHandleTypeFlags
}

func NewPhysicalDeviceExternalImageFormatInfo() *PhysicalDeviceExternalImageFormatInfo {
	p := (*PhysicalDeviceExternalImageFormatInfo)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceExternalImageFormatInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO
	return p
}
func (p *PhysicalDeviceExternalImageFormatInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExternalImageFormatProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalImageFormatProperties.html
type ExternalImageFormatProperties struct {
	SType                    StructureType
	PNext                    unsafe.Pointer
	ExternalMemoryProperties ExternalMemoryProperties
}

func NewExternalImageFormatProperties() *ExternalImageFormatProperties {
	p := (*ExternalImageFormatProperties)(MemAlloc(unsafe.Sizeof(*(*ExternalImageFormatProperties)(nil))))
	p.SType = STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES
	return p
}
func (p *ExternalImageFormatProperties) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceExternalBufferInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html
type PhysicalDeviceExternalBufferInfo struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Flags      BufferCreateFlags
	Usage      BufferUsageFlags
	HandleType ExternalMemoryHandleTypeFlags
}

func NewPhysicalDeviceExternalBufferInfo() *PhysicalDeviceExternalBufferInfo {
	p := (*PhysicalDeviceExternalBufferInfo)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceExternalBufferInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO
	return p
}
func (p *PhysicalDeviceExternalBufferInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExternalBufferProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalBufferProperties.html
type ExternalBufferProperties struct {
	SType                    StructureType
	PNext                    unsafe.Pointer
	ExternalMemoryProperties ExternalMemoryProperties
}

func NewExternalBufferProperties() *ExternalBufferProperties {
	p := (*ExternalBufferProperties)(MemAlloc(unsafe.Sizeof(*(*ExternalBufferProperties)(nil))))
	p.SType = STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES
	return p
}
func (p *ExternalBufferProperties) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceIDProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceIDProperties.html
type PhysicalDeviceIDProperties struct {
	SType           StructureType
	PNext           unsafe.Pointer
	DeviceUUID      [UUID_SIZE]uint8
	DriverUUID      [UUID_SIZE]uint8
	DeviceLUID      [LUID_SIZE]uint8
	DeviceNodeMask  uint32
	DeviceLUIDValid Bool32
}

func NewPhysicalDeviceIDProperties() *PhysicalDeviceIDProperties {
	p := (*PhysicalDeviceIDProperties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceIDProperties)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES
	return p
}
func (p *PhysicalDeviceIDProperties) Free() { MemFree(unsafe.Pointer(p)) }

// ExternalMemoryImageCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryImageCreateInfo.html
type ExternalMemoryImageCreateInfo struct {
	SType       StructureType
	PNext       unsafe.Pointer
	HandleTypes ExternalMemoryHandleTypeFlags
}

func NewExternalMemoryImageCreateInfo() *ExternalMemoryImageCreateInfo {
	p := (*ExternalMemoryImageCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ExternalMemoryImageCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO
	return p
}
func (p *ExternalMemoryImageCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExternalMemoryBufferCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryBufferCreateInfo.html
type ExternalMemoryBufferCreateInfo struct {
	SType       StructureType
	PNext       unsafe.Pointer
	HandleTypes ExternalMemoryHandleTypeFlags
}

func NewExternalMemoryBufferCreateInfo() *ExternalMemoryBufferCreateInfo {
	p := (*ExternalMemoryBufferCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ExternalMemoryBufferCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO
	return p
}
func (p *ExternalMemoryBufferCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExportMemoryAllocateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportMemoryAllocateInfo.html
type ExportMemoryAllocateInfo struct {
	SType       StructureType
	PNext       unsafe.Pointer
	HandleTypes ExternalMemoryHandleTypeFlags
}

func NewExportMemoryAllocateInfo() *ExportMemoryAllocateInfo {
	p := (*ExportMemoryAllocateInfo)(MemAlloc(unsafe.Sizeof(*(*ExportMemoryAllocateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO
	return p
}
func (p *ExportMemoryAllocateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceExternalFenceInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html
type PhysicalDeviceExternalFenceInfo struct {
	SType      StructureType
	PNext      unsafe.Pointer
	HandleType ExternalFenceHandleTypeFlags
}

func NewPhysicalDeviceExternalFenceInfo() *PhysicalDeviceExternalFenceInfo {
	p := (*PhysicalDeviceExternalFenceInfo)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceExternalFenceInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO
	return p
}
func (p *PhysicalDeviceExternalFenceInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExternalFenceProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalFenceProperties.html
type ExternalFenceProperties struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	ExportFromImportedHandleTypes ExternalFenceHandleTypeFlags
	CompatibleHandleTypes         ExternalFenceHandleTypeFlags
	ExternalFenceFeatures         ExternalFenceFeatureFlags
}

func NewExternalFenceProperties() *ExternalFenceProperties {
	p := (*ExternalFenceProperties)(MemAlloc(unsafe.Sizeof(*(*ExternalFenceProperties)(nil))))
	p.SType = STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES
	return p
}
func (p *ExternalFenceProperties) Free() { MemFree(unsafe.Pointer(p)) }

// ExportFenceCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportFenceCreateInfo.html
type ExportFenceCreateInfo struct {
	SType       StructureType
	PNext       unsafe.Pointer
	HandleTypes ExternalFenceHandleTypeFlags
}

func NewExportFenceCreateInfo() *ExportFenceCreateInfo {
	p := (*ExportFenceCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ExportFenceCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO
	return p
}
func (p *ExportFenceCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExportSemaphoreCreateInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportSemaphoreCreateInfo.html
type ExportSemaphoreCreateInfo struct {
	SType       StructureType
	PNext       unsafe.Pointer
	HandleTypes ExternalSemaphoreHandleTypeFlags
}

func NewExportSemaphoreCreateInfo() *ExportSemaphoreCreateInfo {
	p := (*ExportSemaphoreCreateInfo)(MemAlloc(unsafe.Sizeof(*(*ExportSemaphoreCreateInfo)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO
	return p
}
func (p *ExportSemaphoreCreateInfo) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceExternalSemaphoreInfo -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html
type PhysicalDeviceExternalSemaphoreInfo struct {
	SType      StructureType
	PNext      unsafe.Pointer
	HandleType ExternalSemaphoreHandleTypeFlags
}

func NewPhysicalDeviceExternalSemaphoreInfo() *PhysicalDeviceExternalSemaphoreInfo {
	p := (*PhysicalDeviceExternalSemaphoreInfo)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceExternalSemaphoreInfo)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO
	return p
}
func (p *PhysicalDeviceExternalSemaphoreInfo) Free() { MemFree(unsafe.Pointer(p)) }

// ExternalSemaphoreProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalSemaphoreProperties.html
type ExternalSemaphoreProperties struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypeFlags
	CompatibleHandleTypes         ExternalSemaphoreHandleTypeFlags
	ExternalSemaphoreFeatures     ExternalSemaphoreFeatureFlags
}

func NewExternalSemaphoreProperties() *ExternalSemaphoreProperties {
	p := (*ExternalSemaphoreProperties)(MemAlloc(unsafe.Sizeof(*(*ExternalSemaphoreProperties)(nil))))
	p.SType = STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES
	return p
}
func (p *ExternalSemaphoreProperties) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceMaintenance3Properties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html
type PhysicalDeviceMaintenance3Properties struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	MaxPerSetDescriptors    uint32
	MaxMemoryAllocationSize DeviceSize
}

func NewPhysicalDeviceMaintenance3Properties() *PhysicalDeviceMaintenance3Properties {
	p := (*PhysicalDeviceMaintenance3Properties)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMaintenance3Properties)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES
	return p
}
func (p *PhysicalDeviceMaintenance3Properties) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorSetLayoutSupport -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetLayoutSupport.html
type DescriptorSetLayoutSupport struct {
	SType     StructureType
	PNext     unsafe.Pointer
	Supported Bool32
}

func NewDescriptorSetLayoutSupport() *DescriptorSetLayoutSupport {
	p := (*DescriptorSetLayoutSupport)(MemAlloc(unsafe.Sizeof(*(*DescriptorSetLayoutSupport)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
	return p
}
func (p *DescriptorSetLayoutSupport) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceShaderDrawParametersFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderDrawParametersFeatures.html
type PhysicalDeviceShaderDrawParametersFeatures struct {
	SType                StructureType
	PNext                unsafe.Pointer
	ShaderDrawParameters Bool32
}

func NewPhysicalDeviceShaderDrawParametersFeatures() *PhysicalDeviceShaderDrawParametersFeatures {
	p := (*PhysicalDeviceShaderDrawParametersFeatures)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderDrawParametersFeatures)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES
	return p
}
func (p *PhysicalDeviceShaderDrawParametersFeatures) Free() { MemFree(unsafe.Pointer(p)) }

type PhysicalDeviceShaderDrawParameterFeatures = PhysicalDeviceShaderDrawParametersFeatures

//  PfnEnumerateInstanceVersion -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumerateInstanceVersion.html
type PfnEnumerateInstanceVersion uintptr

func (fn PfnEnumerateInstanceVersion) Call(pApiVersion *uint32) Result {
	ret, _, _ := call(uintptr(fn), uintptr(unsafe.Pointer(pApiVersion)))
	return Result(ret)
}
func (fn PfnEnumerateInstanceVersion) String() string { return "vkEnumerateInstanceVersion" }

//  PfnBindBufferMemory2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBindBufferMemory2.html
type PfnBindBufferMemory2 uintptr

func (fn PfnBindBufferMemory2) Call(device Device, bindInfoCount uint32, pBindInfos *BindBufferMemoryInfo) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))
	return Result(ret)
}
func (fn PfnBindBufferMemory2) String() string { return "vkBindBufferMemory2" }

//  PfnBindImageMemory2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBindImageMemory2.html
type PfnBindImageMemory2 uintptr

func (fn PfnBindImageMemory2) Call(device Device, bindInfoCount uint32, pBindInfos *BindImageMemoryInfo) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))
	return Result(ret)
}
func (fn PfnBindImageMemory2) String() string { return "vkBindImageMemory2" }

//  PfnGetDeviceGroupPeerMemoryFeatures -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceGroupPeerMemoryFeatures.html
type PfnGetDeviceGroupPeerMemoryFeatures uintptr

func (fn PfnGetDeviceGroupPeerMemoryFeatures) Call(device Device, heapIndex, localDeviceIndex, remoteDeviceIndex uint32, pPeerMemoryFeatures *PeerMemoryFeatureFlags) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(heapIndex), uintptr(localDeviceIndex), uintptr(remoteDeviceIndex), uintptr(unsafe.Pointer(pPeerMemoryFeatures)))
}
func (fn PfnGetDeviceGroupPeerMemoryFeatures) String() string {
	return "vkGetDeviceGroupPeerMemoryFeatures"
}

//  PfnCmdSetDeviceMask -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetDeviceMask.html
type PfnCmdSetDeviceMask uintptr

func (fn PfnCmdSetDeviceMask) Call(commandBuffer CommandBuffer, deviceMask uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(deviceMask))
}
func (fn PfnCmdSetDeviceMask) String() string { return "vkCmdSetDeviceMask" }

//  PfnCmdDispatchBase -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDispatchBase.html
type PfnCmdDispatchBase uintptr

func (fn PfnCmdDispatchBase) Call(commandBuffer CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(baseGroupX), uintptr(baseGroupY), uintptr(baseGroupZ), uintptr(groupCountX), uintptr(groupCountY), uintptr(groupCountZ))
}
func (fn PfnCmdDispatchBase) String() string { return "vkCmdDispatchBase" }

//  PfnEnumeratePhysicalDeviceGroups -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html
type PfnEnumeratePhysicalDeviceGroups uintptr

func (fn PfnEnumeratePhysicalDeviceGroups) Call(instance Instance, pPhysicalDeviceGroupCount *uint32, pPhysicalDeviceGroupProperties *PhysicalDeviceGroupProperties) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pPhysicalDeviceGroupCount)), uintptr(unsafe.Pointer(pPhysicalDeviceGroupProperties)))
	return Result(ret)
}
func (fn PfnEnumeratePhysicalDeviceGroups) String() string { return "vkEnumeratePhysicalDeviceGroups" }

//  PfnGetImageMemoryRequirements2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageMemoryRequirements2.html
type PfnGetImageMemoryRequirements2 uintptr

func (fn PfnGetImageMemoryRequirements2) Call(device Device, pInfo *ImageMemoryRequirementsInfo2, pMemoryRequirements *MemoryRequirements2) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))
}
func (fn PfnGetImageMemoryRequirements2) String() string { return "vkGetImageMemoryRequirements2" }

//  PfnGetBufferMemoryRequirements2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetBufferMemoryRequirements2.html
type PfnGetBufferMemoryRequirements2 uintptr

func (fn PfnGetBufferMemoryRequirements2) Call(device Device, pInfo *BufferMemoryRequirementsInfo2, pMemoryRequirements *MemoryRequirements2) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))
}
func (fn PfnGetBufferMemoryRequirements2) String() string { return "vkGetBufferMemoryRequirements2" }

//  PfnGetImageSparseMemoryRequirements2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageSparseMemoryRequirements2.html
type PfnGetImageSparseMemoryRequirements2 uintptr

func (fn PfnGetImageSparseMemoryRequirements2) Call(device Device, pInfo *ImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount *uint32, pSparseMemoryRequirements *SparseImageMemoryRequirements2) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)), uintptr(unsafe.Pointer(pSparseMemoryRequirements)))
}
func (fn PfnGetImageSparseMemoryRequirements2) String() string {
	return "vkGetImageSparseMemoryRequirements2"
}

//  PfnGetPhysicalDeviceFeatures2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceFeatures2.html
type PfnGetPhysicalDeviceFeatures2 uintptr

func (fn PfnGetPhysicalDeviceFeatures2) Call(physicalDevice PhysicalDevice, pFeatures *PhysicalDeviceFeatures2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pFeatures)))
}
func (fn PfnGetPhysicalDeviceFeatures2) String() string { return "vkGetPhysicalDeviceFeatures2" }

//  PfnGetPhysicalDeviceProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceProperties2.html
type PfnGetPhysicalDeviceProperties2 uintptr

func (fn PfnGetPhysicalDeviceProperties2) Call(physicalDevice PhysicalDevice, pProperties *PhysicalDeviceProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pProperties)))
}
func (fn PfnGetPhysicalDeviceProperties2) String() string { return "vkGetPhysicalDeviceProperties2" }

//  PfnGetPhysicalDeviceFormatProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html
type PfnGetPhysicalDeviceFormatProperties2 uintptr

func (fn PfnGetPhysicalDeviceFormatProperties2) Call(physicalDevice PhysicalDevice, format Format, pFormatProperties *FormatProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(format), uintptr(unsafe.Pointer(pFormatProperties)))
}
func (fn PfnGetPhysicalDeviceFormatProperties2) String() string {
	return "vkGetPhysicalDeviceFormatProperties2"
}

//  PfnGetPhysicalDeviceImageFormatProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html
type PfnGetPhysicalDeviceImageFormatProperties2 uintptr

func (fn PfnGetPhysicalDeviceImageFormatProperties2) Call(physicalDevice PhysicalDevice, pImageFormatInfo *PhysicalDeviceImageFormatInfo2, pImageFormatProperties *ImageFormatProperties2) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pImageFormatInfo)), uintptr(unsafe.Pointer(pImageFormatProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceImageFormatProperties2) String() string {
	return "vkGetPhysicalDeviceImageFormatProperties2"
}

//  PfnGetPhysicalDeviceQueueFamilyProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html
type PfnGetPhysicalDeviceQueueFamilyProperties2 uintptr

func (fn PfnGetPhysicalDeviceQueueFamilyProperties2) Call(physicalDevice PhysicalDevice, pQueueFamilyPropertyCount *uint32, pQueueFamilyProperties *QueueFamilyProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)), uintptr(unsafe.Pointer(pQueueFamilyProperties)))
}
func (fn PfnGetPhysicalDeviceQueueFamilyProperties2) String() string {
	return "vkGetPhysicalDeviceQueueFamilyProperties2"
}

//  PfnGetPhysicalDeviceMemoryProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html
type PfnGetPhysicalDeviceMemoryProperties2 uintptr

func (fn PfnGetPhysicalDeviceMemoryProperties2) Call(physicalDevice PhysicalDevice, pMemoryProperties *PhysicalDeviceMemoryProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pMemoryProperties)))
}
func (fn PfnGetPhysicalDeviceMemoryProperties2) String() string {
	return "vkGetPhysicalDeviceMemoryProperties2"
}

//  PfnGetPhysicalDeviceSparseImageFormatProperties2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2.html
type PfnGetPhysicalDeviceSparseImageFormatProperties2 uintptr

func (fn PfnGetPhysicalDeviceSparseImageFormatProperties2) Call(physicalDevice PhysicalDevice, pFormatInfo *PhysicalDeviceSparseImageFormatInfo2, pPropertyCount *uint32, pProperties *SparseImageFormatProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pFormatInfo)), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
}
func (fn PfnGetPhysicalDeviceSparseImageFormatProperties2) String() string {
	return "vkGetPhysicalDeviceSparseImageFormatProperties2"
}

//  PfnTrimCommandPool -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkTrimCommandPool.html
type PfnTrimCommandPool uintptr

func (fn PfnTrimCommandPool) Call(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(commandPool), uintptr(flags))
}
func (fn PfnTrimCommandPool) String() string { return "vkTrimCommandPool" }

//  PfnGetDeviceQueue2 -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceQueue2.html
type PfnGetDeviceQueue2 uintptr

func (fn PfnGetDeviceQueue2) Call(device Device, pQueueInfo *DeviceQueueInfo2, pQueue *Queue) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pQueueInfo)), uintptr(unsafe.Pointer(pQueue)))
}
func (fn PfnGetDeviceQueue2) String() string { return "vkGetDeviceQueue2" }

//  PfnCreateSamplerYcbcrConversion -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateSamplerYcbcrConversion.html
type PfnCreateSamplerYcbcrConversion uintptr

func (fn PfnCreateSamplerYcbcrConversion) Call(device Device, pCreateInfo *SamplerYcbcrConversionCreateInfo, pAllocator *AllocationCallbacks, pYcbcrConversion *SamplerYcbcrConversion) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pYcbcrConversion)))
	return Result(ret)
}
func (fn PfnCreateSamplerYcbcrConversion) String() string { return "vkCreateSamplerYcbcrConversion" }

//  PfnDestroySamplerYcbcrConversion -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroySamplerYcbcrConversion.html
type PfnDestroySamplerYcbcrConversion uintptr

func (fn PfnDestroySamplerYcbcrConversion) Call(device Device, ycbcrConversion SamplerYcbcrConversion, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(ycbcrConversion), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroySamplerYcbcrConversion) String() string { return "vkDestroySamplerYcbcrConversion" }

//  PfnCreateDescriptorUpdateTemplate -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDescriptorUpdateTemplate.html
type PfnCreateDescriptorUpdateTemplate uintptr

func (fn PfnCreateDescriptorUpdateTemplate) Call(device Device, pCreateInfo *DescriptorUpdateTemplateCreateInfo, pAllocator *AllocationCallbacks, pDescriptorUpdateTemplate *DescriptorUpdateTemplate) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDescriptorUpdateTemplate)))
	return Result(ret)
}
func (fn PfnCreateDescriptorUpdateTemplate) String() string {
	return "vkCreateDescriptorUpdateTemplate"
}

//  PfnDestroyDescriptorUpdateTemplate -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html
type PfnDestroyDescriptorUpdateTemplate uintptr

func (fn PfnDestroyDescriptorUpdateTemplate) Call(device Device, descriptorUpdateTemplate DescriptorUpdateTemplate, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(descriptorUpdateTemplate), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyDescriptorUpdateTemplate) String() string {
	return "vkDestroyDescriptorUpdateTemplate"
}

//  PfnUpdateDescriptorSetWithTemplate -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkUpdateDescriptorSetWithTemplate.html
type PfnUpdateDescriptorSetWithTemplate uintptr

func (fn PfnUpdateDescriptorSetWithTemplate) Call(device Device, descriptorSet DescriptorSet, descriptorUpdateTemplate DescriptorUpdateTemplate, pData unsafe.Pointer) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(descriptorSet), uintptr(descriptorUpdateTemplate), uintptr(pData))
}
func (fn PfnUpdateDescriptorSetWithTemplate) String() string {
	return "vkUpdateDescriptorSetWithTemplate"
}

//  PfnGetPhysicalDeviceExternalBufferProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html
type PfnGetPhysicalDeviceExternalBufferProperties uintptr

func (fn PfnGetPhysicalDeviceExternalBufferProperties) Call(physicalDevice PhysicalDevice, pExternalBufferInfo *PhysicalDeviceExternalBufferInfo, pExternalBufferProperties *ExternalBufferProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalBufferInfo)), uintptr(unsafe.Pointer(pExternalBufferProperties)))
}
func (fn PfnGetPhysicalDeviceExternalBufferProperties) String() string {
	return "vkGetPhysicalDeviceExternalBufferProperties"
}

//  PfnGetPhysicalDeviceExternalFenceProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html
type PfnGetPhysicalDeviceExternalFenceProperties uintptr

func (fn PfnGetPhysicalDeviceExternalFenceProperties) Call(physicalDevice PhysicalDevice, pExternalFenceInfo *PhysicalDeviceExternalFenceInfo, pExternalFenceProperties *ExternalFenceProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalFenceInfo)), uintptr(unsafe.Pointer(pExternalFenceProperties)))
}
func (fn PfnGetPhysicalDeviceExternalFenceProperties) String() string {
	return "vkGetPhysicalDeviceExternalFenceProperties"
}

//  PfnGetPhysicalDeviceExternalSemaphoreProperties -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html
type PfnGetPhysicalDeviceExternalSemaphoreProperties uintptr

func (fn PfnGetPhysicalDeviceExternalSemaphoreProperties) Call(physicalDevice PhysicalDevice, pExternalSemaphoreInfo *PhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties *ExternalSemaphoreProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalSemaphoreInfo)), uintptr(unsafe.Pointer(pExternalSemaphoreProperties)))
}
func (fn PfnGetPhysicalDeviceExternalSemaphoreProperties) String() string {
	return "vkGetPhysicalDeviceExternalSemaphoreProperties"
}

//  PfnGetDescriptorSetLayoutSupport -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDescriptorSetLayoutSupport.html
type PfnGetDescriptorSetLayoutSupport uintptr

func (fn PfnGetDescriptorSetLayoutSupport) Call(device Device, pCreateInfo *DescriptorSetLayoutCreateInfo, pSupport *DescriptorSetLayoutSupport) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pSupport)))
}
func (fn PfnGetDescriptorSetLayoutSupport) String() string { return "vkGetDescriptorSetLayoutSupport" }

const KHR_surface = 1

// SurfaceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceKHR.html
type SurfaceKHR NonDispatchableHandle

const KHR_SURFACE_SPEC_VERSION = 25

var KHR_SURFACE_EXTENSION_NAME = "VK_KHR_surface"

// ColorSpaceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkColorSpaceKHR.html
type ColorSpaceKHR int32

const (
	COLOR_SPACE_SRGB_NONLINEAR_KHR          ColorSpaceKHR = 0
	COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT    ColorSpaceKHR = 1000104001
	COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT    ColorSpaceKHR = 1000104002
	COLOR_SPACE_DISPLAY_P3_LINEAR_EXT       ColorSpaceKHR = 1000104003
	COLOR_SPACE_DCI_P3_NONLINEAR_EXT        ColorSpaceKHR = 1000104004
	COLOR_SPACE_BT709_LINEAR_EXT            ColorSpaceKHR = 1000104005
	COLOR_SPACE_BT709_NONLINEAR_EXT         ColorSpaceKHR = 1000104006
	COLOR_SPACE_BT2020_LINEAR_EXT           ColorSpaceKHR = 1000104007
	COLOR_SPACE_HDR10_ST2084_EXT            ColorSpaceKHR = 1000104008
	COLOR_SPACE_DOLBYVISION_EXT             ColorSpaceKHR = 1000104009
	COLOR_SPACE_HDR10_HLG_EXT               ColorSpaceKHR = 1000104010
	COLOR_SPACE_ADOBERGB_LINEAR_EXT         ColorSpaceKHR = 1000104011
	COLOR_SPACE_ADOBERGB_NONLINEAR_EXT      ColorSpaceKHR = 1000104012
	COLOR_SPACE_PASS_THROUGH_EXT            ColorSpaceKHR = 1000104013
	COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT ColorSpaceKHR = 1000104014
	COLOR_SPACE_DISPLAY_NATIVE_AMD          ColorSpaceKHR = 1000213000
	COLORSPACE_SRGB_NONLINEAR_KHR           ColorSpaceKHR = COLOR_SPACE_SRGB_NONLINEAR_KHR
	COLOR_SPACE_DCI_P3_LINEAR_EXT           ColorSpaceKHR = COLOR_SPACE_DISPLAY_P3_LINEAR_EXT
	COLOR_SPACE_BEGIN_RANGE_KHR             ColorSpaceKHR = COLOR_SPACE_SRGB_NONLINEAR_KHR
	COLOR_SPACE_END_RANGE_KHR               ColorSpaceKHR = COLOR_SPACE_SRGB_NONLINEAR_KHR
	COLOR_SPACE_RANGE_SIZE_KHR              ColorSpaceKHR = (COLOR_SPACE_SRGB_NONLINEAR_KHR - COLOR_SPACE_SRGB_NONLINEAR_KHR + 1)
	COLOR_SPACE_MAX_ENUM_KHR                ColorSpaceKHR = 0x7FFFFFFF
)

func (x ColorSpaceKHR) String() string {
	switch x {
	case COLOR_SPACE_SRGB_NONLINEAR_KHR:
		return "COLOR_SPACE_SRGB_NONLINEAR_KHR"
	case COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT:
		return "COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT"
	case COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT:
		return "COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT"
	case COLOR_SPACE_DISPLAY_P3_LINEAR_EXT:
		return "COLOR_SPACE_DISPLAY_P3_LINEAR_EXT"
	case COLOR_SPACE_DCI_P3_NONLINEAR_EXT:
		return "COLOR_SPACE_DCI_P3_NONLINEAR_EXT"
	case COLOR_SPACE_BT709_LINEAR_EXT:
		return "COLOR_SPACE_BT709_LINEAR_EXT"
	case COLOR_SPACE_BT709_NONLINEAR_EXT:
		return "COLOR_SPACE_BT709_NONLINEAR_EXT"
	case COLOR_SPACE_BT2020_LINEAR_EXT:
		return "COLOR_SPACE_BT2020_LINEAR_EXT"
	case COLOR_SPACE_HDR10_ST2084_EXT:
		return "COLOR_SPACE_HDR10_ST2084_EXT"
	case COLOR_SPACE_DOLBYVISION_EXT:
		return "COLOR_SPACE_DOLBYVISION_EXT"
	case COLOR_SPACE_HDR10_HLG_EXT:
		return "COLOR_SPACE_HDR10_HLG_EXT"
	case COLOR_SPACE_ADOBERGB_LINEAR_EXT:
		return "COLOR_SPACE_ADOBERGB_LINEAR_EXT"
	case COLOR_SPACE_ADOBERGB_NONLINEAR_EXT:
		return "COLOR_SPACE_ADOBERGB_NONLINEAR_EXT"
	case COLOR_SPACE_PASS_THROUGH_EXT:
		return "COLOR_SPACE_PASS_THROUGH_EXT"
	case COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT:
		return "COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT"
	case COLOR_SPACE_DISPLAY_NATIVE_AMD:
		return "COLOR_SPACE_DISPLAY_NATIVE_AMD"
	case COLOR_SPACE_MAX_ENUM_KHR:
		return "COLOR_SPACE_MAX_ENUM_KHR"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PresentModeKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPresentModeKHR.html
type PresentModeKHR int32

const (
	PRESENT_MODE_IMMEDIATE_KHR                 PresentModeKHR = 0
	PRESENT_MODE_MAILBOX_KHR                   PresentModeKHR = 1
	PRESENT_MODE_FIFO_KHR                      PresentModeKHR = 2
	PRESENT_MODE_FIFO_RELAXED_KHR              PresentModeKHR = 3
	PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR     PresentModeKHR = 1000111000
	PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR PresentModeKHR = 1000111001
	PRESENT_MODE_BEGIN_RANGE_KHR               PresentModeKHR = PRESENT_MODE_IMMEDIATE_KHR
	PRESENT_MODE_END_RANGE_KHR                 PresentModeKHR = PRESENT_MODE_FIFO_RELAXED_KHR
	PRESENT_MODE_RANGE_SIZE_KHR                PresentModeKHR = (PRESENT_MODE_FIFO_RELAXED_KHR - PRESENT_MODE_IMMEDIATE_KHR + 1)
	PRESENT_MODE_MAX_ENUM_KHR                  PresentModeKHR = 0x7FFFFFFF
)

func (x PresentModeKHR) String() string {
	switch x {
	case PRESENT_MODE_IMMEDIATE_KHR:
		return "PRESENT_MODE_IMMEDIATE_KHR"
	case PRESENT_MODE_MAILBOX_KHR:
		return "PRESENT_MODE_MAILBOX_KHR"
	case PRESENT_MODE_FIFO_KHR:
		return "PRESENT_MODE_FIFO_KHR"
	case PRESENT_MODE_FIFO_RELAXED_KHR:
		return "PRESENT_MODE_FIFO_RELAXED_KHR"
	case PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR:
		return "PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR"
	case PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR:
		return "PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR"
	case PRESENT_MODE_MAX_ENUM_KHR:
		return "PRESENT_MODE_MAX_ENUM_KHR"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SurfaceTransformFlagsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceTransformFlagsKHR.html
type SurfaceTransformFlagsKHR uint32

const (
	SURFACE_TRANSFORM_IDENTITY_BIT_KHR                     SurfaceTransformFlagsKHR = 0x00000001
	SURFACE_TRANSFORM_ROTATE_90_BIT_KHR                    SurfaceTransformFlagsKHR = 0x00000002
	SURFACE_TRANSFORM_ROTATE_180_BIT_KHR                   SurfaceTransformFlagsKHR = 0x00000004
	SURFACE_TRANSFORM_ROTATE_270_BIT_KHR                   SurfaceTransformFlagsKHR = 0x00000008
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR            SurfaceTransformFlagsKHR = 0x00000010
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR  SurfaceTransformFlagsKHR = 0x00000020
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR SurfaceTransformFlagsKHR = 0x00000040
	SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR SurfaceTransformFlagsKHR = 0x00000080
	SURFACE_TRANSFORM_INHERIT_BIT_KHR                      SurfaceTransformFlagsKHR = 0x00000100
	SURFACE_TRANSFORM_FLAG_BITS_MAX_ENUM_KHR               SurfaceTransformFlagsKHR = 0x7FFFFFFF
)

func (x SurfaceTransformFlagsKHR) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SurfaceTransformFlagsKHR(1 << i) {
			case SURFACE_TRANSFORM_IDENTITY_BIT_KHR:
				s += "SURFACE_TRANSFORM_IDENTITY_BIT_KHR|"
			case SURFACE_TRANSFORM_ROTATE_90_BIT_KHR:
				s += "SURFACE_TRANSFORM_ROTATE_90_BIT_KHR|"
			case SURFACE_TRANSFORM_ROTATE_180_BIT_KHR:
				s += "SURFACE_TRANSFORM_ROTATE_180_BIT_KHR|"
			case SURFACE_TRANSFORM_ROTATE_270_BIT_KHR:
				s += "SURFACE_TRANSFORM_ROTATE_270_BIT_KHR|"
			case SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR:
				s += "SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR|"
			case SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR:
				s += "SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR|"
			case SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR:
				s += "SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR|"
			case SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR:
				s += "SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR|"
			case SURFACE_TRANSFORM_INHERIT_BIT_KHR:
				s += "SURFACE_TRANSFORM_INHERIT_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// CompositeAlphaFlagsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCompositeAlphaFlagsKHR.html
type CompositeAlphaFlagsKHR uint32

const (
	COMPOSITE_ALPHA_OPAQUE_BIT_KHR          CompositeAlphaFlagsKHR = 0x00000001
	COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR  CompositeAlphaFlagsKHR = 0x00000002
	COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR CompositeAlphaFlagsKHR = 0x00000004
	COMPOSITE_ALPHA_INHERIT_BIT_KHR         CompositeAlphaFlagsKHR = 0x00000008
	COMPOSITE_ALPHA_FLAG_BITS_MAX_ENUM_KHR  CompositeAlphaFlagsKHR = 0x7FFFFFFF
)

func (x CompositeAlphaFlagsKHR) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch CompositeAlphaFlagsKHR(1 << i) {
			case COMPOSITE_ALPHA_OPAQUE_BIT_KHR:
				s += "COMPOSITE_ALPHA_OPAQUE_BIT_KHR|"
			case COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR:
				s += "COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR|"
			case COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR:
				s += "COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR|"
			case COMPOSITE_ALPHA_INHERIT_BIT_KHR:
				s += "COMPOSITE_ALPHA_INHERIT_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SurfaceCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceCapabilitiesKHR.html
type SurfaceCapabilitiesKHR struct {
	MinImageCount           uint32
	MaxImageCount           uint32
	CurrentExtent           Extent2D
	MinImageExtent          Extent2D
	MaxImageExtent          Extent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     SurfaceTransformFlagsKHR
	CurrentTransform        SurfaceTransformFlagsKHR
	SupportedCompositeAlpha CompositeAlphaFlagsKHR
	SupportedUsageFlags     ImageUsageFlags
}

func NewSurfaceCapabilitiesKHR() *SurfaceCapabilitiesKHR {
	return (*SurfaceCapabilitiesKHR)(MemAlloc(unsafe.Sizeof(*(*SurfaceCapabilitiesKHR)(nil))))
}
func (p *SurfaceCapabilitiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// SurfaceFormatKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceFormatKHR.html
type SurfaceFormatKHR struct {
	Format     Format
	ColorSpace ColorSpaceKHR
}

func NewSurfaceFormatKHR() *SurfaceFormatKHR {
	return (*SurfaceFormatKHR)(MemAlloc(unsafe.Sizeof(*(*SurfaceFormatKHR)(nil))))
}
func (p *SurfaceFormatKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnDestroySurfaceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroySurfaceKHR.html
type PfnDestroySurfaceKHR uintptr

func (fn PfnDestroySurfaceKHR) Call(instance Instance, surface SurfaceKHR, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(instance), uintptr(surface), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroySurfaceKHR) String() string { return "vkDestroySurfaceKHR" }

//  PfnGetPhysicalDeviceSurfaceSupportKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html
type PfnGetPhysicalDeviceSurfaceSupportKHR uintptr

func (fn PfnGetPhysicalDeviceSurfaceSupportKHR) Call(physicalDevice PhysicalDevice, queueFamilyIndex uint32, surface SurfaceKHR, pSupported *Bool32) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(queueFamilyIndex), uintptr(surface), uintptr(unsafe.Pointer(pSupported)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfaceSupportKHR) String() string {
	return "vkGetPhysicalDeviceSurfaceSupportKHR"
}

//  PfnGetPhysicalDeviceSurfaceCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html
type PfnGetPhysicalDeviceSurfaceCapabilitiesKHR uintptr

func (fn PfnGetPhysicalDeviceSurfaceCapabilitiesKHR) Call(physicalDevice PhysicalDevice, surface SurfaceKHR, pSurfaceCapabilities *SurfaceCapabilitiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pSurfaceCapabilities)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfaceCapabilitiesKHR) String() string {
	return "vkGetPhysicalDeviceSurfaceCapabilitiesKHR"
}

//  PfnGetPhysicalDeviceSurfaceFormatsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html
type PfnGetPhysicalDeviceSurfaceFormatsKHR uintptr

func (fn PfnGetPhysicalDeviceSurfaceFormatsKHR) Call(physicalDevice PhysicalDevice, surface SurfaceKHR, pSurfaceFormatCount *uint32, pSurfaceFormats *SurfaceFormatKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pSurfaceFormatCount)), uintptr(unsafe.Pointer(pSurfaceFormats)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfaceFormatsKHR) String() string {
	return "vkGetPhysicalDeviceSurfaceFormatsKHR"
}

//  PfnGetPhysicalDeviceSurfacePresentModesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html
type PfnGetPhysicalDeviceSurfacePresentModesKHR uintptr

func (fn PfnGetPhysicalDeviceSurfacePresentModesKHR) Call(physicalDevice PhysicalDevice, surface SurfaceKHR, pPresentModeCount *uint32, pPresentModes *PresentModeKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pPresentModeCount)), uintptr(unsafe.Pointer(pPresentModes)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfacePresentModesKHR) String() string {
	return "vkGetPhysicalDeviceSurfacePresentModesKHR"
}

const KHR_swapchain = 1

// SwapchainKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSwapchainKHR.html
type SwapchainKHR NonDispatchableHandle

const KHR_SWAPCHAIN_SPEC_VERSION = 70

var KHR_SWAPCHAIN_EXTENSION_NAME = "VK_KHR_swapchain"

// SwapchainCreateFlagsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSwapchainCreateFlagsKHR.html
type SwapchainCreateFlagsKHR uint32

const (
	SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR SwapchainCreateFlagsKHR = 0x00000001
	SWAPCHAIN_CREATE_PROTECTED_BIT_KHR                   SwapchainCreateFlagsKHR = 0x00000002
	SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR              SwapchainCreateFlagsKHR = 0x00000004
	SWAPCHAIN_CREATE_FLAG_BITS_MAX_ENUM_KHR              SwapchainCreateFlagsKHR = 0x7FFFFFFF
)

func (x SwapchainCreateFlagsKHR) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SwapchainCreateFlagsKHR(1 << i) {
			case SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR:
				s += "SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR|"
			case SWAPCHAIN_CREATE_PROTECTED_BIT_KHR:
				s += "SWAPCHAIN_CREATE_PROTECTED_BIT_KHR|"
			case SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR:
				s += "SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DeviceGroupPresentModeFlagsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html
type DeviceGroupPresentModeFlagsKHR uint32

const (
	DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR              DeviceGroupPresentModeFlagsKHR = 0x00000001
	DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR             DeviceGroupPresentModeFlagsKHR = 0x00000002
	DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR                DeviceGroupPresentModeFlagsKHR = 0x00000004
	DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR DeviceGroupPresentModeFlagsKHR = 0x00000008
	DEVICE_GROUP_PRESENT_MODE_FLAG_BITS_MAX_ENUM_KHR     DeviceGroupPresentModeFlagsKHR = 0x7FFFFFFF
)

func (x DeviceGroupPresentModeFlagsKHR) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DeviceGroupPresentModeFlagsKHR(1 << i) {
			case DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR:
				s += "DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR|"
			case DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR:
				s += "DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR|"
			case DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR:
				s += "DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR|"
			case DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR:
				s += "DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SwapchainCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSwapchainCreateInfoKHR.html
type SwapchainCreateInfoKHR struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Flags                 SwapchainCreateFlagsKHR
	Surface               SurfaceKHR
	MinImageCount         uint32
	ImageFormat           Format
	ImageColorSpace       ColorSpaceKHR
	ImageExtent           Extent2D
	ImageArrayLayers      uint32
	ImageUsage            ImageUsageFlags
	ImageSharingMode      SharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
	PreTransform          SurfaceTransformFlagsKHR
	CompositeAlpha        CompositeAlphaFlagsKHR
	PresentMode           PresentModeKHR
	Clipped               Bool32
	OldSwapchain          SwapchainKHR
}

func NewSwapchainCreateInfoKHR() *SwapchainCreateInfoKHR {
	p := (*SwapchainCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*SwapchainCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	return p
}
func (p *SwapchainCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PresentInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPresentInfoKHR.html
type PresentInfoKHR struct {
	SType              StructureType
	PNext              unsafe.Pointer
	WaitSemaphoreCount uint32
	PWaitSemaphores    *Semaphore
	SwapchainCount     uint32
	PSwapchains        *SwapchainKHR
	PImageIndices      *uint32
	PResults           *Result
}

func NewPresentInfoKHR() *PresentInfoKHR {
	p := (*PresentInfoKHR)(MemAlloc(unsafe.Sizeof(*(*PresentInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PRESENT_INFO_KHR
	return p
}
func (p *PresentInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// ImageSwapchainCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageSwapchainCreateInfoKHR.html
type ImageSwapchainCreateInfoKHR struct {
	SType     StructureType
	PNext     unsafe.Pointer
	Swapchain SwapchainKHR
}

func NewImageSwapchainCreateInfoKHR() *ImageSwapchainCreateInfoKHR {
	p := (*ImageSwapchainCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImageSwapchainCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR
	return p
}
func (p *ImageSwapchainCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// BindImageMemorySwapchainInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html
type BindImageMemorySwapchainInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Swapchain  SwapchainKHR
	ImageIndex uint32
}

func NewBindImageMemorySwapchainInfoKHR() *BindImageMemorySwapchainInfoKHR {
	p := (*BindImageMemorySwapchainInfoKHR)(MemAlloc(unsafe.Sizeof(*(*BindImageMemorySwapchainInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR
	return p
}
func (p *BindImageMemorySwapchainInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// AcquireNextImageInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAcquireNextImageInfoKHR.html
type AcquireNextImageInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Swapchain  SwapchainKHR
	Timeout    uint64
	Semaphore  Semaphore
	Fence      Fence
	DeviceMask uint32
}

func NewAcquireNextImageInfoKHR() *AcquireNextImageInfoKHR {
	p := (*AcquireNextImageInfoKHR)(MemAlloc(unsafe.Sizeof(*(*AcquireNextImageInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR
	return p
}
func (p *AcquireNextImageInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupPresentCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupPresentCapabilitiesKHR.html
type DeviceGroupPresentCapabilitiesKHR struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PresentMask [MAX_DEVICE_GROUP_SIZE]uint32
	Modes       DeviceGroupPresentModeFlagsKHR
}

func NewDeviceGroupPresentCapabilitiesKHR() *DeviceGroupPresentCapabilitiesKHR {
	p := (*DeviceGroupPresentCapabilitiesKHR)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupPresentCapabilitiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR
	return p
}
func (p *DeviceGroupPresentCapabilitiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupPresentInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupPresentInfoKHR.html
type DeviceGroupPresentInfoKHR struct {
	SType          StructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PDeviceMasks   *uint32
	Mode           DeviceGroupPresentModeFlagsKHR
}

func NewDeviceGroupPresentInfoKHR() *DeviceGroupPresentInfoKHR {
	p := (*DeviceGroupPresentInfoKHR)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupPresentInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR
	return p
}
func (p *DeviceGroupPresentInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGroupSwapchainCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html
type DeviceGroupSwapchainCreateInfoKHR struct {
	SType StructureType
	PNext unsafe.Pointer
	Modes DeviceGroupPresentModeFlagsKHR
}

func NewDeviceGroupSwapchainCreateInfoKHR() *DeviceGroupSwapchainCreateInfoKHR {
	p := (*DeviceGroupSwapchainCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*DeviceGroupSwapchainCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR
	return p
}
func (p *DeviceGroupSwapchainCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateSwapchainKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateSwapchainKHR.html
type PfnCreateSwapchainKHR uintptr

func (fn PfnCreateSwapchainKHR) Call(device Device, pCreateInfo *SwapchainCreateInfoKHR, pAllocator *AllocationCallbacks, pSwapchain *SwapchainKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSwapchain)))
	return Result(ret)
}
func (fn PfnCreateSwapchainKHR) String() string { return "vkCreateSwapchainKHR" }

//  PfnDestroySwapchainKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroySwapchainKHR.html
type PfnDestroySwapchainKHR uintptr

func (fn PfnDestroySwapchainKHR) Call(device Device, swapchain SwapchainKHR, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroySwapchainKHR) String() string { return "vkDestroySwapchainKHR" }

//  PfnGetSwapchainImagesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetSwapchainImagesKHR.html
type PfnGetSwapchainImagesKHR uintptr

func (fn PfnGetSwapchainImagesKHR) Call(device Device, swapchain SwapchainKHR, pSwapchainImageCount *uint32, pSwapchainImages *Image) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pSwapchainImageCount)), uintptr(unsafe.Pointer(pSwapchainImages)))
	return Result(ret)
}
func (fn PfnGetSwapchainImagesKHR) String() string { return "vkGetSwapchainImagesKHR" }

//  PfnAcquireNextImageKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAcquireNextImageKHR.html
type PfnAcquireNextImageKHR uintptr

func (fn PfnAcquireNextImageKHR) Call(device Device, swapchain SwapchainKHR, timeout uint64, semaphore Semaphore, fence Fence, pImageIndex *uint32) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain), uintptr(timeout), uintptr(semaphore), uintptr(fence), uintptr(unsafe.Pointer(pImageIndex)))
	return Result(ret)
}
func (fn PfnAcquireNextImageKHR) String() string { return "vkAcquireNextImageKHR" }

//  PfnQueuePresentKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueuePresentKHR.html
type PfnQueuePresentKHR uintptr

func (fn PfnQueuePresentKHR) Call(queue Queue, pPresentInfo *PresentInfoKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(queue), uintptr(unsafe.Pointer(pPresentInfo)))
	return Result(ret)
}
func (fn PfnQueuePresentKHR) String() string { return "vkQueuePresentKHR" }

//  PfnGetDeviceGroupPresentCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html
type PfnGetDeviceGroupPresentCapabilitiesKHR uintptr

func (fn PfnGetDeviceGroupPresentCapabilitiesKHR) Call(device Device, pDeviceGroupPresentCapabilities *DeviceGroupPresentCapabilitiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pDeviceGroupPresentCapabilities)))
	return Result(ret)
}
func (fn PfnGetDeviceGroupPresentCapabilitiesKHR) String() string {
	return "vkGetDeviceGroupPresentCapabilitiesKHR"
}

//  PfnGetDeviceGroupSurfacePresentModesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html
type PfnGetDeviceGroupSurfacePresentModesKHR uintptr

func (fn PfnGetDeviceGroupSurfacePresentModesKHR) Call(device Device, surface SurfaceKHR, pModes *DeviceGroupPresentModeFlagsKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(surface), uintptr(unsafe.Pointer(pModes)))
	return Result(ret)
}
func (fn PfnGetDeviceGroupSurfacePresentModesKHR) String() string {
	return "vkGetDeviceGroupSurfacePresentModesKHR"
}

//  PfnGetPhysicalDevicePresentRectanglesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html
type PfnGetPhysicalDevicePresentRectanglesKHR uintptr

func (fn PfnGetPhysicalDevicePresentRectanglesKHR) Call(physicalDevice PhysicalDevice, surface SurfaceKHR, pRectCount *uint32, pRects *Rect2D) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pRectCount)), uintptr(unsafe.Pointer(pRects)))
	return Result(ret)
}
func (fn PfnGetPhysicalDevicePresentRectanglesKHR) String() string {
	return "vkGetPhysicalDevicePresentRectanglesKHR"
}

//  PfnAcquireNextImage2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAcquireNextImage2KHR.html
type PfnAcquireNextImage2KHR uintptr

func (fn PfnAcquireNextImage2KHR) Call(device Device, pAcquireInfo *AcquireNextImageInfoKHR, pImageIndex *uint32) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pAcquireInfo)), uintptr(unsafe.Pointer(pImageIndex)))
	return Result(ret)
}
func (fn PfnAcquireNextImage2KHR) String() string { return "vkAcquireNextImage2KHR" }

const KHR_display = 1

// DisplayKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayKHR.html
type DisplayKHR NonDispatchableHandle

// DisplayModeKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayModeKHR.html
type DisplayModeKHR NonDispatchableHandle

const KHR_DISPLAY_SPEC_VERSION = 23

var KHR_DISPLAY_EXTENSION_NAME = "VK_KHR_display"

// DisplayPlaneAlphaFlagsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPlaneAlphaFlagsKHR.html
type DisplayPlaneAlphaFlagsKHR uint32

const (
	DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR                  DisplayPlaneAlphaFlagsKHR = 0x00000001
	DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR                  DisplayPlaneAlphaFlagsKHR = 0x00000002
	DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR               DisplayPlaneAlphaFlagsKHR = 0x00000004
	DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR DisplayPlaneAlphaFlagsKHR = 0x00000008
	DISPLAY_PLANE_ALPHA_FLAG_BITS_MAX_ENUM_KHR          DisplayPlaneAlphaFlagsKHR = 0x7FFFFFFF
)

func (x DisplayPlaneAlphaFlagsKHR) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DisplayPlaneAlphaFlagsKHR(1 << i) {
			case DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR:
				s += "DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR|"
			case DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR:
				s += "DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR|"
			case DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR:
				s += "DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR|"
			case DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR:
				s += "DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

type DisplayModeCreateFlagsKHR uint32    // reserved
type DisplaySurfaceCreateFlagsKHR uint32 // reserved
// DisplayPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPropertiesKHR.html
type DisplayPropertiesKHR struct {
	Display              DisplayKHR
	DisplayName          *int8
	PhysicalDimensions   Extent2D
	PhysicalResolution   Extent2D
	SupportedTransforms  SurfaceTransformFlagsKHR
	PlaneReorderPossible Bool32
	PersistentContent    Bool32
}

func NewDisplayPropertiesKHR() *DisplayPropertiesKHR {
	return (*DisplayPropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*DisplayPropertiesKHR)(nil))))
}
func (p *DisplayPropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayModeParametersKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayModeParametersKHR.html
type DisplayModeParametersKHR struct {
	VisibleRegion Extent2D
	RefreshRate   uint32
}

func NewDisplayModeParametersKHR() *DisplayModeParametersKHR {
	return (*DisplayModeParametersKHR)(MemAlloc(unsafe.Sizeof(*(*DisplayModeParametersKHR)(nil))))
}
func (p *DisplayModeParametersKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayModePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayModePropertiesKHR.html
type DisplayModePropertiesKHR struct {
	DisplayMode DisplayModeKHR
	Parameters  DisplayModeParametersKHR
}

func NewDisplayModePropertiesKHR() *DisplayModePropertiesKHR {
	return (*DisplayModePropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*DisplayModePropertiesKHR)(nil))))
}
func (p *DisplayModePropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayModeCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayModeCreateInfoKHR.html
type DisplayModeCreateInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Flags      DisplayModeCreateFlagsKHR
	Parameters DisplayModeParametersKHR
}

func NewDisplayModeCreateInfoKHR() *DisplayModeCreateInfoKHR {
	p := (*DisplayModeCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*DisplayModeCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR
	return p
}
func (p *DisplayModeCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayPlaneCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPlaneCapabilitiesKHR.html
type DisplayPlaneCapabilitiesKHR struct {
	SupportedAlpha DisplayPlaneAlphaFlagsKHR
	MinSrcPosition Offset2D
	MaxSrcPosition Offset2D
	MinSrcExtent   Extent2D
	MaxSrcExtent   Extent2D
	MinDstPosition Offset2D
	MaxDstPosition Offset2D
	MinDstExtent   Extent2D
	MaxDstExtent   Extent2D
}

func NewDisplayPlaneCapabilitiesKHR() *DisplayPlaneCapabilitiesKHR {
	return (*DisplayPlaneCapabilitiesKHR)(MemAlloc(unsafe.Sizeof(*(*DisplayPlaneCapabilitiesKHR)(nil))))
}
func (p *DisplayPlaneCapabilitiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayPlanePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPlanePropertiesKHR.html
type DisplayPlanePropertiesKHR struct {
	CurrentDisplay    DisplayKHR
	CurrentStackIndex uint32
}

func NewDisplayPlanePropertiesKHR() *DisplayPlanePropertiesKHR {
	return (*DisplayPlanePropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*DisplayPlanePropertiesKHR)(nil))))
}
func (p *DisplayPlanePropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplaySurfaceCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html
type DisplaySurfaceCreateInfoKHR struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           DisplaySurfaceCreateFlagsKHR
	DisplayMode     DisplayModeKHR
	PlaneIndex      uint32
	PlaneStackIndex uint32
	Transform       SurfaceTransformFlagsKHR
	GlobalAlpha     float32
	AlphaMode       DisplayPlaneAlphaFlagsKHR
	ImageExtent     Extent2D
}

func NewDisplaySurfaceCreateInfoKHR() *DisplaySurfaceCreateInfoKHR {
	p := (*DisplaySurfaceCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*DisplaySurfaceCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR
	return p
}
func (p *DisplaySurfaceCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceDisplayPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceDisplayPropertiesKHR.html
type PfnGetPhysicalDeviceDisplayPropertiesKHR uintptr

func (fn PfnGetPhysicalDeviceDisplayPropertiesKHR) Call(physicalDevice PhysicalDevice, pPropertyCount *uint32, pProperties *DisplayPropertiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceDisplayPropertiesKHR) String() string {
	return "vkGetPhysicalDeviceDisplayPropertiesKHR"
}

//  PfnGetPhysicalDeviceDisplayPlanePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceDisplayPlanePropertiesKHR.html
type PfnGetPhysicalDeviceDisplayPlanePropertiesKHR uintptr

func (fn PfnGetPhysicalDeviceDisplayPlanePropertiesKHR) Call(physicalDevice PhysicalDevice, pPropertyCount *uint32, pProperties *DisplayPlanePropertiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceDisplayPlanePropertiesKHR) String() string {
	return "vkGetPhysicalDeviceDisplayPlanePropertiesKHR"
}

//  PfnGetDisplayPlaneSupportedDisplaysKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDisplayPlaneSupportedDisplaysKHR.html
type PfnGetDisplayPlaneSupportedDisplaysKHR uintptr

func (fn PfnGetDisplayPlaneSupportedDisplaysKHR) Call(physicalDevice PhysicalDevice, planeIndex uint32, pDisplayCount *uint32, pDisplays *DisplayKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(planeIndex), uintptr(unsafe.Pointer(pDisplayCount)), uintptr(unsafe.Pointer(pDisplays)))
	return Result(ret)
}
func (fn PfnGetDisplayPlaneSupportedDisplaysKHR) String() string {
	return "vkGetDisplayPlaneSupportedDisplaysKHR"
}

//  PfnGetDisplayModePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDisplayModePropertiesKHR.html
type PfnGetDisplayModePropertiesKHR uintptr

func (fn PfnGetDisplayModePropertiesKHR) Call(physicalDevice PhysicalDevice, display DisplayKHR, pPropertyCount *uint32, pProperties *DisplayModePropertiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(display), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetDisplayModePropertiesKHR) String() string { return "vkGetDisplayModePropertiesKHR" }

//  PfnCreateDisplayModeKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDisplayModeKHR.html
type PfnCreateDisplayModeKHR uintptr

func (fn PfnCreateDisplayModeKHR) Call(physicalDevice PhysicalDevice, display DisplayKHR, pCreateInfo *DisplayModeCreateInfoKHR, pAllocator *AllocationCallbacks, pMode *DisplayModeKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(display), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pMode)))
	return Result(ret)
}
func (fn PfnCreateDisplayModeKHR) String() string { return "vkCreateDisplayModeKHR" }

//  PfnGetDisplayPlaneCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDisplayPlaneCapabilitiesKHR.html
type PfnGetDisplayPlaneCapabilitiesKHR uintptr

func (fn PfnGetDisplayPlaneCapabilitiesKHR) Call(physicalDevice PhysicalDevice, mode DisplayModeKHR, planeIndex uint32, pCapabilities *DisplayPlaneCapabilitiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(mode), uintptr(planeIndex), uintptr(unsafe.Pointer(pCapabilities)))
	return Result(ret)
}
func (fn PfnGetDisplayPlaneCapabilitiesKHR) String() string {
	return "vkGetDisplayPlaneCapabilitiesKHR"
}

//  PfnCreateDisplayPlaneSurfaceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDisplayPlaneSurfaceKHR.html
type PfnCreateDisplayPlaneSurfaceKHR uintptr

func (fn PfnCreateDisplayPlaneSurfaceKHR) Call(instance Instance, pCreateInfo *DisplaySurfaceCreateInfoKHR, pAllocator *AllocationCallbacks, pSurface *SurfaceKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)))
	return Result(ret)
}
func (fn PfnCreateDisplayPlaneSurfaceKHR) String() string { return "vkCreateDisplayPlaneSurfaceKHR" }

const KHR_display_swapchain = 1
const KHR_DISPLAY_SWAPCHAIN_SPEC_VERSION = 10

var KHR_DISPLAY_SWAPCHAIN_EXTENSION_NAME = "VK_KHR_display_swapchain"

// DisplayPresentInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPresentInfoKHR.html
type DisplayPresentInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	SrcRect    Rect2D
	DstRect    Rect2D
	Persistent Bool32
}

func NewDisplayPresentInfoKHR() *DisplayPresentInfoKHR {
	p := (*DisplayPresentInfoKHR)(MemAlloc(unsafe.Sizeof(*(*DisplayPresentInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_PRESENT_INFO_KHR
	return p
}
func (p *DisplayPresentInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateSharedSwapchainsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateSharedSwapchainsKHR.html
type PfnCreateSharedSwapchainsKHR uintptr

func (fn PfnCreateSharedSwapchainsKHR) Call(device Device, swapchainCount uint32, pCreateInfos *SwapchainCreateInfoKHR, pAllocator *AllocationCallbacks, pSwapchains *SwapchainKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchainCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSwapchains)))
	return Result(ret)
}
func (fn PfnCreateSharedSwapchainsKHR) String() string { return "vkCreateSharedSwapchainsKHR" }

const KHR_sampler_mirror_clamp_to_edge = 1
const KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_SPEC_VERSION = 3

var KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_EXTENSION_NAME = "VK_KHR_sampler_mirror_clamp_to_edge"

const KHR_multiview = 1
const KHR_MULTIVIEW_SPEC_VERSION = 1

var KHR_MULTIVIEW_EXTENSION_NAME = "VK_KHR_multiview"

type RenderPassMultiviewCreateInfoKHR = RenderPassMultiviewCreateInfo
type PhysicalDeviceMultiviewFeaturesKHR = PhysicalDeviceMultiviewFeatures
type PhysicalDeviceMultiviewPropertiesKHR = PhysicalDeviceMultiviewProperties

const KHR_get_physical_device_properties2 = 1
const KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_SPEC_VERSION = 2

var KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_EXTENSION_NAME = "VK_KHR_get_physical_device_properties2"

type PhysicalDeviceFeatures2KHR = PhysicalDeviceFeatures2
type PhysicalDeviceProperties2KHR = PhysicalDeviceProperties2
type FormatProperties2KHR = FormatProperties2
type ImageFormatProperties2KHR = ImageFormatProperties2
type PhysicalDeviceImageFormatInfo2KHR = PhysicalDeviceImageFormatInfo2
type QueueFamilyProperties2KHR = QueueFamilyProperties2
type PhysicalDeviceMemoryProperties2KHR = PhysicalDeviceMemoryProperties2
type SparseImageFormatProperties2KHR = SparseImageFormatProperties2
type PhysicalDeviceSparseImageFormatInfo2KHR = PhysicalDeviceSparseImageFormatInfo2

//  PfnGetPhysicalDeviceFeatures2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceFeatures2KHR.html
type PfnGetPhysicalDeviceFeatures2KHR uintptr

func (fn PfnGetPhysicalDeviceFeatures2KHR) Call(physicalDevice PhysicalDevice, pFeatures *PhysicalDeviceFeatures2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pFeatures)))
}
func (fn PfnGetPhysicalDeviceFeatures2KHR) String() string { return "vkGetPhysicalDeviceFeatures2KHR" }

//  PfnGetPhysicalDeviceProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceProperties2KHR.html
type PfnGetPhysicalDeviceProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceProperties2KHR) Call(physicalDevice PhysicalDevice, pProperties *PhysicalDeviceProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pProperties)))
}
func (fn PfnGetPhysicalDeviceProperties2KHR) String() string {
	return "vkGetPhysicalDeviceProperties2KHR"
}

//  PfnGetPhysicalDeviceFormatProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceFormatProperties2KHR.html
type PfnGetPhysicalDeviceFormatProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceFormatProperties2KHR) Call(physicalDevice PhysicalDevice, format Format, pFormatProperties *FormatProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(format), uintptr(unsafe.Pointer(pFormatProperties)))
}
func (fn PfnGetPhysicalDeviceFormatProperties2KHR) String() string {
	return "vkGetPhysicalDeviceFormatProperties2KHR"
}

//  PfnGetPhysicalDeviceImageFormatProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2KHR.html
type PfnGetPhysicalDeviceImageFormatProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceImageFormatProperties2KHR) Call(physicalDevice PhysicalDevice, pImageFormatInfo *PhysicalDeviceImageFormatInfo2, pImageFormatProperties *ImageFormatProperties2) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pImageFormatInfo)), uintptr(unsafe.Pointer(pImageFormatProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceImageFormatProperties2KHR) String() string {
	return "vkGetPhysicalDeviceImageFormatProperties2KHR"
}

//  PfnGetPhysicalDeviceQueueFamilyProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2KHR.html
type PfnGetPhysicalDeviceQueueFamilyProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceQueueFamilyProperties2KHR) Call(physicalDevice PhysicalDevice, pQueueFamilyPropertyCount *uint32, pQueueFamilyProperties *QueueFamilyProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pQueueFamilyPropertyCount)), uintptr(unsafe.Pointer(pQueueFamilyProperties)))
}
func (fn PfnGetPhysicalDeviceQueueFamilyProperties2KHR) String() string {
	return "vkGetPhysicalDeviceQueueFamilyProperties2KHR"
}

//  PfnGetPhysicalDeviceMemoryProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2KHR.html
type PfnGetPhysicalDeviceMemoryProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceMemoryProperties2KHR) Call(physicalDevice PhysicalDevice, pMemoryProperties *PhysicalDeviceMemoryProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pMemoryProperties)))
}
func (fn PfnGetPhysicalDeviceMemoryProperties2KHR) String() string {
	return "vkGetPhysicalDeviceMemoryProperties2KHR"
}

//  PfnGetPhysicalDeviceSparseImageFormatProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html
type PfnGetPhysicalDeviceSparseImageFormatProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceSparseImageFormatProperties2KHR) Call(physicalDevice PhysicalDevice, pFormatInfo *PhysicalDeviceSparseImageFormatInfo2, pPropertyCount *uint32, pProperties *SparseImageFormatProperties2) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pFormatInfo)), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
}
func (fn PfnGetPhysicalDeviceSparseImageFormatProperties2KHR) String() string {
	return "vkGetPhysicalDeviceSparseImageFormatProperties2KHR"
}

const KHR_device_group = 1
const KHR_DEVICE_GROUP_SPEC_VERSION = 4

var KHR_DEVICE_GROUP_EXTENSION_NAME = "VK_KHR_device_group"

type PeerMemoryFeatureFlagsKHR = PeerMemoryFeatureFlags
type MemoryAllocateFlagsKHR = MemoryAllocateFlags
type DeviceGroupRenderPassBeginInfoKHR = DeviceGroupRenderPassBeginInfo
type DeviceGroupCommandBufferBeginInfoKHR = DeviceGroupCommandBufferBeginInfo
type DeviceGroupSubmitInfoKHR = DeviceGroupSubmitInfo
type DeviceGroupBindSparseInfoKHR = DeviceGroupBindSparseInfo
type BindBufferMemoryDeviceGroupInfoKHR = BindBufferMemoryDeviceGroupInfo
type BindImageMemoryDeviceGroupInfoKHR = BindImageMemoryDeviceGroupInfo

//  PfnGetDeviceGroupPeerMemoryFeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceGroupPeerMemoryFeaturesKHR.html
type PfnGetDeviceGroupPeerMemoryFeaturesKHR uintptr

func (fn PfnGetDeviceGroupPeerMemoryFeaturesKHR) Call(device Device, heapIndex, localDeviceIndex, remoteDeviceIndex uint32, pPeerMemoryFeatures *PeerMemoryFeatureFlags) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(heapIndex), uintptr(localDeviceIndex), uintptr(remoteDeviceIndex), uintptr(unsafe.Pointer(pPeerMemoryFeatures)))
}
func (fn PfnGetDeviceGroupPeerMemoryFeaturesKHR) String() string {
	return "vkGetDeviceGroupPeerMemoryFeaturesKHR"
}

//  PfnCmdSetDeviceMaskKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetDeviceMaskKHR.html
type PfnCmdSetDeviceMaskKHR uintptr

func (fn PfnCmdSetDeviceMaskKHR) Call(commandBuffer CommandBuffer, deviceMask uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(deviceMask))
}
func (fn PfnCmdSetDeviceMaskKHR) String() string { return "vkCmdSetDeviceMaskKHR" }

//  PfnCmdDispatchBaseKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDispatchBaseKHR.html
type PfnCmdDispatchBaseKHR uintptr

func (fn PfnCmdDispatchBaseKHR) Call(commandBuffer CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(baseGroupX), uintptr(baseGroupY), uintptr(baseGroupZ), uintptr(groupCountX), uintptr(groupCountY), uintptr(groupCountZ))
}
func (fn PfnCmdDispatchBaseKHR) String() string { return "vkCmdDispatchBaseKHR" }

const KHR_shader_draw_parameters = 1
const KHR_SHADER_DRAW_PARAMETERS_SPEC_VERSION = 1

var KHR_SHADER_DRAW_PARAMETERS_EXTENSION_NAME = "VK_KHR_shader_draw_parameters"

const KHR_maintenance1 = 1
const KHR_MAINTENANCE1_SPEC_VERSION = 2

var KHR_MAINTENANCE1_EXTENSION_NAME = "VK_KHR_maintenance1"

//  PfnTrimCommandPoolKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkTrimCommandPoolKHR.html
type PfnTrimCommandPoolKHR uintptr

func (fn PfnTrimCommandPoolKHR) Call(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(commandPool), uintptr(flags))
}
func (fn PfnTrimCommandPoolKHR) String() string { return "vkTrimCommandPoolKHR" }

const KHR_device_group_creation = 1
const KHR_DEVICE_GROUP_CREATION_SPEC_VERSION = 1

var KHR_DEVICE_GROUP_CREATION_EXTENSION_NAME = "VK_KHR_device_group_creation"

type PhysicalDeviceGroupPropertiesKHR = PhysicalDeviceGroupProperties
type DeviceGroupDeviceCreateInfoKHR = DeviceGroupDeviceCreateInfo

//  PfnEnumeratePhysicalDeviceGroupsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkEnumeratePhysicalDeviceGroupsKHR.html
type PfnEnumeratePhysicalDeviceGroupsKHR uintptr

func (fn PfnEnumeratePhysicalDeviceGroupsKHR) Call(instance Instance, pPhysicalDeviceGroupCount *uint32, pPhysicalDeviceGroupProperties *PhysicalDeviceGroupProperties) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pPhysicalDeviceGroupCount)), uintptr(unsafe.Pointer(pPhysicalDeviceGroupProperties)))
	return Result(ret)
}
func (fn PfnEnumeratePhysicalDeviceGroupsKHR) String() string {
	return "vkEnumeratePhysicalDeviceGroupsKHR"
}

const KHR_external_memory_capabilities = 1
const KHR_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION = 1

var KHR_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME = "VK_KHR_external_memory_capabilities"

type ExternalMemoryHandleTypeFlagsKHR = ExternalMemoryHandleTypeFlags
type ExternalMemoryFeatureFlagsKHR = ExternalMemoryFeatureFlags
type ExternalMemoryPropertiesKHR = ExternalMemoryProperties
type PhysicalDeviceExternalImageFormatInfoKHR = PhysicalDeviceExternalImageFormatInfo
type ExternalImageFormatPropertiesKHR = ExternalImageFormatProperties
type PhysicalDeviceExternalBufferInfoKHR = PhysicalDeviceExternalBufferInfo
type ExternalBufferPropertiesKHR = ExternalBufferProperties
type PhysicalDeviceIDPropertiesKHR = PhysicalDeviceIDProperties

//  PfnGetPhysicalDeviceExternalBufferPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceExternalBufferPropertiesKHR.html
type PfnGetPhysicalDeviceExternalBufferPropertiesKHR uintptr

func (fn PfnGetPhysicalDeviceExternalBufferPropertiesKHR) Call(physicalDevice PhysicalDevice, pExternalBufferInfo *PhysicalDeviceExternalBufferInfo, pExternalBufferProperties *ExternalBufferProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalBufferInfo)), uintptr(unsafe.Pointer(pExternalBufferProperties)))
}
func (fn PfnGetPhysicalDeviceExternalBufferPropertiesKHR) String() string {
	return "vkGetPhysicalDeviceExternalBufferPropertiesKHR"
}

const KHR_external_memory = 1
const KHR_EXTERNAL_MEMORY_SPEC_VERSION = 1

var KHR_EXTERNAL_MEMORY_EXTENSION_NAME = "VK_KHR_external_memory"

type ExternalMemoryImageCreateInfoKHR = ExternalMemoryImageCreateInfo
type ExternalMemoryBufferCreateInfoKHR = ExternalMemoryBufferCreateInfo
type ExportMemoryAllocateInfoKHR = ExportMemoryAllocateInfo

const KHR_external_memory_fd = 1
const KHR_EXTERNAL_MEMORY_FD_SPEC_VERSION = 1

var KHR_EXTERNAL_MEMORY_FD_EXTENSION_NAME = "VK_KHR_external_memory_fd"

// ImportMemoryFdInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportMemoryFdInfoKHR.html
type ImportMemoryFdInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	HandleType ExternalMemoryHandleTypeFlags
	Fd         int
}

func NewImportMemoryFdInfoKHR() *ImportMemoryFdInfoKHR {
	p := (*ImportMemoryFdInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImportMemoryFdInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR
	return p
}
func (p *ImportMemoryFdInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryFdPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryFdPropertiesKHR.html
type MemoryFdPropertiesKHR struct {
	SType          StructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

func NewMemoryFdPropertiesKHR() *MemoryFdPropertiesKHR {
	p := (*MemoryFdPropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*MemoryFdPropertiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR
	return p
}
func (p *MemoryFdPropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryGetFdInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryGetFdInfoKHR.html
type MemoryGetFdInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Memory     DeviceMemory
	HandleType ExternalMemoryHandleTypeFlags
}

func NewMemoryGetFdInfoKHR() *MemoryGetFdInfoKHR {
	p := (*MemoryGetFdInfoKHR)(MemAlloc(unsafe.Sizeof(*(*MemoryGetFdInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR
	return p
}
func (p *MemoryGetFdInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetMemoryFdKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetMemoryFdKHR.html
type PfnGetMemoryFdKHR uintptr

func (fn PfnGetMemoryFdKHR) Call(device Device, pGetFdInfo *MemoryGetFdInfoKHR, pFd *int) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pGetFdInfo)), uintptr(unsafe.Pointer(pFd)))
	return Result(ret)
}
func (fn PfnGetMemoryFdKHR) String() string { return "vkGetMemoryFdKHR" }

//  PfnGetMemoryFdPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetMemoryFdPropertiesKHR.html
type PfnGetMemoryFdPropertiesKHR uintptr

func (fn PfnGetMemoryFdPropertiesKHR) Call(device Device, handleType ExternalMemoryHandleTypeFlags, fd int, pMemoryFdProperties *MemoryFdPropertiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(handleType), uintptr(fd), uintptr(unsafe.Pointer(pMemoryFdProperties)))
	return Result(ret)
}
func (fn PfnGetMemoryFdPropertiesKHR) String() string { return "vkGetMemoryFdPropertiesKHR" }

const KHR_external_semaphore_capabilities = 1
const KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_SPEC_VERSION = 1

var KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_EXTENSION_NAME = "VK_KHR_external_semaphore_capabilities"

type ExternalSemaphoreHandleTypeFlagsKHR = ExternalSemaphoreHandleTypeFlags
type ExternalSemaphoreFeatureFlagsKHR = ExternalSemaphoreFeatureFlags
type PhysicalDeviceExternalSemaphoreInfoKHR = PhysicalDeviceExternalSemaphoreInfo
type ExternalSemaphorePropertiesKHR = ExternalSemaphoreProperties

//  PfnGetPhysicalDeviceExternalSemaphorePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceExternalSemaphorePropertiesKHR.html
type PfnGetPhysicalDeviceExternalSemaphorePropertiesKHR uintptr

func (fn PfnGetPhysicalDeviceExternalSemaphorePropertiesKHR) Call(physicalDevice PhysicalDevice, pExternalSemaphoreInfo *PhysicalDeviceExternalSemaphoreInfo, pExternalSemaphoreProperties *ExternalSemaphoreProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalSemaphoreInfo)), uintptr(unsafe.Pointer(pExternalSemaphoreProperties)))
}
func (fn PfnGetPhysicalDeviceExternalSemaphorePropertiesKHR) String() string {
	return "vkGetPhysicalDeviceExternalSemaphorePropertiesKHR"
}

const KHR_external_semaphore = 1
const KHR_EXTERNAL_SEMAPHORE_SPEC_VERSION = 1

var KHR_EXTERNAL_SEMAPHORE_EXTENSION_NAME = "VK_KHR_external_semaphore"

type SemaphoreImportFlagsKHR = SemaphoreImportFlags
type ExportSemaphoreCreateInfoKHR = ExportSemaphoreCreateInfo

const KHR_external_semaphore_fd = 1
const KHR_EXTERNAL_SEMAPHORE_FD_SPEC_VERSION = 1

var KHR_EXTERNAL_SEMAPHORE_FD_EXTENSION_NAME = "VK_KHR_external_semaphore_fd"

// ImportSemaphoreFdInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportSemaphoreFdInfoKHR.html
type ImportSemaphoreFdInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Semaphore  Semaphore
	Flags      SemaphoreImportFlags
	HandleType ExternalSemaphoreHandleTypeFlags
	Fd         int
}

func NewImportSemaphoreFdInfoKHR() *ImportSemaphoreFdInfoKHR {
	p := (*ImportSemaphoreFdInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImportSemaphoreFdInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR
	return p
}
func (p *ImportSemaphoreFdInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// SemaphoreGetFdInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSemaphoreGetFdInfoKHR.html
type SemaphoreGetFdInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Semaphore  Semaphore
	HandleType ExternalSemaphoreHandleTypeFlags
}

func NewSemaphoreGetFdInfoKHR() *SemaphoreGetFdInfoKHR {
	p := (*SemaphoreGetFdInfoKHR)(MemAlloc(unsafe.Sizeof(*(*SemaphoreGetFdInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR
	return p
}
func (p *SemaphoreGetFdInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnImportSemaphoreFdKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkImportSemaphoreFdKHR.html
type PfnImportSemaphoreFdKHR uintptr

func (fn PfnImportSemaphoreFdKHR) Call(device Device, pImportSemaphoreFdInfo *ImportSemaphoreFdInfoKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pImportSemaphoreFdInfo)))
	return Result(ret)
}
func (fn PfnImportSemaphoreFdKHR) String() string { return "vkImportSemaphoreFdKHR" }

//  PfnGetSemaphoreFdKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetSemaphoreFdKHR.html
type PfnGetSemaphoreFdKHR uintptr

func (fn PfnGetSemaphoreFdKHR) Call(device Device, pGetFdInfo *SemaphoreGetFdInfoKHR, pFd *int) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pGetFdInfo)), uintptr(unsafe.Pointer(pFd)))
	return Result(ret)
}
func (fn PfnGetSemaphoreFdKHR) String() string { return "vkGetSemaphoreFdKHR" }

const KHR_push_descriptor = 1
const KHR_PUSH_DESCRIPTOR_SPEC_VERSION = 2

var KHR_PUSH_DESCRIPTOR_EXTENSION_NAME = "VK_KHR_push_descriptor"

// PhysicalDevicePushDescriptorPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDevicePushDescriptorPropertiesKHR.html
type PhysicalDevicePushDescriptorPropertiesKHR struct {
	SType              StructureType
	PNext              unsafe.Pointer
	MaxPushDescriptors uint32
}

func NewPhysicalDevicePushDescriptorPropertiesKHR() *PhysicalDevicePushDescriptorPropertiesKHR {
	p := (*PhysicalDevicePushDescriptorPropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDevicePushDescriptorPropertiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR
	return p
}
func (p *PhysicalDevicePushDescriptorPropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdPushDescriptorSetKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdPushDescriptorSetKHR.html
type PfnCmdPushDescriptorSetKHR uintptr

func (fn PfnCmdPushDescriptorSetKHR) Call(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, set, descriptorWriteCount uint32, pDescriptorWrites *WriteDescriptorSet) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(pipelineBindPoint), uintptr(layout), uintptr(set), uintptr(descriptorWriteCount), uintptr(unsafe.Pointer(pDescriptorWrites)))
}
func (fn PfnCmdPushDescriptorSetKHR) String() string { return "vkCmdPushDescriptorSetKHR" }

//  PfnCmdPushDescriptorSetWithTemplateKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html
type PfnCmdPushDescriptorSetWithTemplateKHR uintptr

func (fn PfnCmdPushDescriptorSetWithTemplateKHR) Call(commandBuffer CommandBuffer, descriptorUpdateTemplate DescriptorUpdateTemplate, layout PipelineLayout, set uint32, pData unsafe.Pointer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(descriptorUpdateTemplate), uintptr(layout), uintptr(set), uintptr(pData))
}
func (fn PfnCmdPushDescriptorSetWithTemplateKHR) String() string {
	return "vkCmdPushDescriptorSetWithTemplateKHR"
}

const KHR_shader_float16_int8 = 1
const KHR_SHADER_FLOAT16_INT8_SPEC_VERSION = 1

var KHR_SHADER_FLOAT16_INT8_EXTENSION_NAME = "VK_KHR_shader_float16_int8"

// PhysicalDeviceShaderFloat16Int8FeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderFloat16Int8FeaturesKHR.html
type PhysicalDeviceShaderFloat16Int8FeaturesKHR struct {
	SType         StructureType
	PNext         unsafe.Pointer
	ShaderFloat16 Bool32
	ShaderInt8    Bool32
}

func NewPhysicalDeviceShaderFloat16Int8FeaturesKHR() *PhysicalDeviceShaderFloat16Int8FeaturesKHR {
	return (*PhysicalDeviceShaderFloat16Int8FeaturesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderFloat16Int8FeaturesKHR)(nil))))
}
func (p *PhysicalDeviceShaderFloat16Int8FeaturesKHR) Free() { MemFree(unsafe.Pointer(p)) }

type PhysicalDeviceFloat16Int8FeaturesKHR = PhysicalDeviceShaderFloat16Int8FeaturesKHR

const KHR_16bit_storage = 1
const KHR_16BIT_STORAGE_SPEC_VERSION = 1

var KHR_16BIT_STORAGE_EXTENSION_NAME = "VK_KHR_16bit_storage"

type PhysicalDevice16BitStorageFeaturesKHR = PhysicalDevice16BitStorageFeatures

const KHR_incremental_present = 1
const KHR_INCREMENTAL_PRESENT_SPEC_VERSION = 1

var KHR_INCREMENTAL_PRESENT_EXTENSION_NAME = "VK_KHR_incremental_present"

// RectLayerKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRectLayerKHR.html
type RectLayerKHR struct {
	Offset Offset2D
	Extent Extent2D
	Layer  uint32
}

func NewRectLayerKHR() *RectLayerKHR {
	return (*RectLayerKHR)(MemAlloc(unsafe.Sizeof(*(*RectLayerKHR)(nil))))
}
func (p *RectLayerKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PresentRegionKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPresentRegionKHR.html
type PresentRegionKHR struct {
	RectangleCount uint32
	PRectangles    *RectLayerKHR
}

func NewPresentRegionKHR() *PresentRegionKHR {
	return (*PresentRegionKHR)(MemAlloc(unsafe.Sizeof(*(*PresentRegionKHR)(nil))))
}
func (p *PresentRegionKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PresentRegionsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPresentRegionsKHR.html
type PresentRegionsKHR struct {
	SType          StructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PRegions       *PresentRegionKHR
}

func NewPresentRegionsKHR() *PresentRegionsKHR {
	p := (*PresentRegionsKHR)(MemAlloc(unsafe.Sizeof(*(*PresentRegionsKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PRESENT_REGIONS_KHR
	return p
}
func (p *PresentRegionsKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_descriptor_update_template = 1

type DescriptorUpdateTemplateKHR = DescriptorUpdateTemplate

const KHR_DESCRIPTOR_UPDATE_TEMPLATE_SPEC_VERSION = 1

var KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME = "VK_KHR_descriptor_update_template"

type DescriptorUpdateTemplateTypeKHR = DescriptorUpdateTemplateType
type DescriptorUpdateTemplateEntryKHR = DescriptorUpdateTemplateEntry
type DescriptorUpdateTemplateCreateInfoKHR = DescriptorUpdateTemplateCreateInfo

//  PfnCreateDescriptorUpdateTemplateKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDescriptorUpdateTemplateKHR.html
type PfnCreateDescriptorUpdateTemplateKHR uintptr

func (fn PfnCreateDescriptorUpdateTemplateKHR) Call(device Device, pCreateInfo *DescriptorUpdateTemplateCreateInfo, pAllocator *AllocationCallbacks, pDescriptorUpdateTemplate *DescriptorUpdateTemplate) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pDescriptorUpdateTemplate)))
	return Result(ret)
}
func (fn PfnCreateDescriptorUpdateTemplateKHR) String() string {
	return "vkCreateDescriptorUpdateTemplateKHR"
}

//  PfnDestroyDescriptorUpdateTemplateKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyDescriptorUpdateTemplateKHR.html
type PfnDestroyDescriptorUpdateTemplateKHR uintptr

func (fn PfnDestroyDescriptorUpdateTemplateKHR) Call(device Device, descriptorUpdateTemplate DescriptorUpdateTemplate, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(descriptorUpdateTemplate), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyDescriptorUpdateTemplateKHR) String() string {
	return "vkDestroyDescriptorUpdateTemplateKHR"
}

//  PfnUpdateDescriptorSetWithTemplateKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
type PfnUpdateDescriptorSetWithTemplateKHR uintptr

func (fn PfnUpdateDescriptorSetWithTemplateKHR) Call(device Device, descriptorSet DescriptorSet, descriptorUpdateTemplate DescriptorUpdateTemplate, pData unsafe.Pointer) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(descriptorSet), uintptr(descriptorUpdateTemplate), uintptr(pData))
}
func (fn PfnUpdateDescriptorSetWithTemplateKHR) String() string {
	return "vkUpdateDescriptorSetWithTemplateKHR"
}

const KHR_imageless_framebuffer = 1
const KHR_IMAGELESS_FRAMEBUFFER_SPEC_VERSION = 1

var KHR_IMAGELESS_FRAMEBUFFER_EXTENSION_NAME = "VK_KHR_imageless_framebuffer"

// PhysicalDeviceImagelessFramebufferFeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceImagelessFramebufferFeaturesKHR.html
type PhysicalDeviceImagelessFramebufferFeaturesKHR struct {
	SType                StructureType
	PNext                unsafe.Pointer
	ImagelessFramebuffer Bool32
}

func NewPhysicalDeviceImagelessFramebufferFeaturesKHR() *PhysicalDeviceImagelessFramebufferFeaturesKHR {
	p := (*PhysicalDeviceImagelessFramebufferFeaturesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceImagelessFramebufferFeaturesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
	return p
}
func (p *PhysicalDeviceImagelessFramebufferFeaturesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// FramebufferAttachmentImageInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFramebufferAttachmentImageInfoKHR.html
type FramebufferAttachmentImageInfoKHR struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           ImageCreateFlags
	Usage           ImageUsageFlags
	Width           uint32
	Height          uint32
	LayerCount      uint32
	ViewFormatCount uint32
	PViewFormats    *Format
}

func NewFramebufferAttachmentImageInfoKHR() *FramebufferAttachmentImageInfoKHR {
	p := (*FramebufferAttachmentImageInfoKHR)(MemAlloc(unsafe.Sizeof(*(*FramebufferAttachmentImageInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR
	return p
}
func (p *FramebufferAttachmentImageInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// FramebufferAttachmentsCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFramebufferAttachmentsCreateInfoKHR.html
type FramebufferAttachmentsCreateInfoKHR struct {
	SType                    StructureType
	PNext                    unsafe.Pointer
	AttachmentImageInfoCount uint32
	PAttachmentImageInfos    *FramebufferAttachmentImageInfoKHR
}

func NewFramebufferAttachmentsCreateInfoKHR() *FramebufferAttachmentsCreateInfoKHR {
	p := (*FramebufferAttachmentsCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*FramebufferAttachmentsCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR
	return p
}
func (p *FramebufferAttachmentsCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassAttachmentBeginInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassAttachmentBeginInfoKHR.html
type RenderPassAttachmentBeginInfoKHR struct {
	SType           StructureType
	PNext           unsafe.Pointer
	AttachmentCount uint32
	PAttachments    *ImageView
}

func NewRenderPassAttachmentBeginInfoKHR() *RenderPassAttachmentBeginInfoKHR {
	p := (*RenderPassAttachmentBeginInfoKHR)(MemAlloc(unsafe.Sizeof(*(*RenderPassAttachmentBeginInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR
	return p
}
func (p *RenderPassAttachmentBeginInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_create_renderpass2 = 1
const KHR_CREATE_RENDERPASS_2_SPEC_VERSION = 1

var KHR_CREATE_RENDERPASS_2_EXTENSION_NAME = "VK_KHR_create_renderpass2"

// AttachmentDescription2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentDescription2KHR.html
type AttachmentDescription2KHR struct {
	SType          StructureType
	PNext          unsafe.Pointer
	Flags          AttachmentDescriptionFlags
	Format         Format
	Samples        SampleCountFlags
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}

func NewAttachmentDescription2KHR() *AttachmentDescription2KHR {
	p := (*AttachmentDescription2KHR)(MemAlloc(unsafe.Sizeof(*(*AttachmentDescription2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR
	return p
}
func (p *AttachmentDescription2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// AttachmentReference2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentReference2KHR.html
type AttachmentReference2KHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Attachment uint32
	Layout     ImageLayout
	AspectMask ImageAspectFlags
}

func NewAttachmentReference2KHR() *AttachmentReference2KHR {
	p := (*AttachmentReference2KHR)(MemAlloc(unsafe.Sizeof(*(*AttachmentReference2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
	return p
}
func (p *AttachmentReference2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// SubpassDescription2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassDescription2KHR.html
type SubpassDescription2KHR struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	Flags                   SubpassDescriptionFlags
	PipelineBindPoint       PipelineBindPoint
	ViewMask                uint32
	InputAttachmentCount    uint32
	PInputAttachments       *AttachmentReference2KHR
	ColorAttachmentCount    uint32
	PColorAttachments       *AttachmentReference2KHR
	PResolveAttachments     *AttachmentReference2KHR
	PDepthStencilAttachment *AttachmentReference2KHR
	PreserveAttachmentCount uint32
	PPreserveAttachments    *uint32
}

func NewSubpassDescription2KHR() *SubpassDescription2KHR {
	p := (*SubpassDescription2KHR)(MemAlloc(unsafe.Sizeof(*(*SubpassDescription2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR
	return p
}
func (p *SubpassDescription2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// SubpassDependency2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassDependency2KHR.html
type SubpassDependency2KHR struct {
	SType           StructureType
	PNext           unsafe.Pointer
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
	ViewOffset      int32
}

func NewSubpassDependency2KHR() *SubpassDependency2KHR {
	p := (*SubpassDependency2KHR)(MemAlloc(unsafe.Sizeof(*(*SubpassDependency2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR
	return p
}
func (p *SubpassDependency2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassCreateInfo2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassCreateInfo2KHR.html
type RenderPassCreateInfo2KHR struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	Flags                   RenderPassCreateFlags
	AttachmentCount         uint32
	PAttachments            *AttachmentDescription2KHR
	SubpassCount            uint32
	PSubpasses              *SubpassDescription2KHR
	DependencyCount         uint32
	PDependencies           *SubpassDependency2KHR
	CorrelatedViewMaskCount uint32
	PCorrelatedViewMasks    *uint32
}

func NewRenderPassCreateInfo2KHR() *RenderPassCreateInfo2KHR {
	p := (*RenderPassCreateInfo2KHR)(MemAlloc(unsafe.Sizeof(*(*RenderPassCreateInfo2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR
	return p
}
func (p *RenderPassCreateInfo2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// SubpassBeginInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassBeginInfoKHR.html
type SubpassBeginInfoKHR struct {
	SType    StructureType
	PNext    unsafe.Pointer
	Contents SubpassContents
}

func NewSubpassBeginInfoKHR() *SubpassBeginInfoKHR {
	p := (*SubpassBeginInfoKHR)(MemAlloc(unsafe.Sizeof(*(*SubpassBeginInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR
	return p
}
func (p *SubpassBeginInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// SubpassEndInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassEndInfoKHR.html
type SubpassEndInfoKHR struct {
	SType StructureType
	PNext unsafe.Pointer
}

func NewSubpassEndInfoKHR() *SubpassEndInfoKHR {
	p := (*SubpassEndInfoKHR)(MemAlloc(unsafe.Sizeof(*(*SubpassEndInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SUBPASS_END_INFO_KHR
	return p
}
func (p *SubpassEndInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateRenderPass2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateRenderPass2KHR.html
type PfnCreateRenderPass2KHR uintptr

func (fn PfnCreateRenderPass2KHR) Call(device Device, pCreateInfo *RenderPassCreateInfo2KHR, pAllocator *AllocationCallbacks, pRenderPass *RenderPass) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pRenderPass)))
	return Result(ret)
}
func (fn PfnCreateRenderPass2KHR) String() string { return "vkCreateRenderPass2KHR" }

//  PfnCmdBeginRenderPass2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBeginRenderPass2KHR.html
type PfnCmdBeginRenderPass2KHR uintptr

func (fn PfnCmdBeginRenderPass2KHR) Call(commandBuffer CommandBuffer, pRenderPassBegin *RenderPassBeginInfo, pSubpassBeginInfo *SubpassBeginInfoKHR) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pRenderPassBegin)), uintptr(unsafe.Pointer(pSubpassBeginInfo)))
}
func (fn PfnCmdBeginRenderPass2KHR) String() string { return "vkCmdBeginRenderPass2KHR" }

//  PfnCmdNextSubpass2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdNextSubpass2KHR.html
type PfnCmdNextSubpass2KHR uintptr

func (fn PfnCmdNextSubpass2KHR) Call(commandBuffer CommandBuffer, pSubpassBeginInfo *SubpassBeginInfoKHR, pSubpassEndInfo *SubpassEndInfoKHR) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pSubpassBeginInfo)), uintptr(unsafe.Pointer(pSubpassEndInfo)))
}
func (fn PfnCmdNextSubpass2KHR) String() string { return "vkCmdNextSubpass2KHR" }

//  PfnCmdEndRenderPass2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdEndRenderPass2KHR.html
type PfnCmdEndRenderPass2KHR uintptr

func (fn PfnCmdEndRenderPass2KHR) Call(commandBuffer CommandBuffer, pSubpassEndInfo *SubpassEndInfoKHR) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pSubpassEndInfo)))
}
func (fn PfnCmdEndRenderPass2KHR) String() string { return "vkCmdEndRenderPass2KHR" }

const KHR_shared_presentable_image = 1
const KHR_SHARED_PRESENTABLE_IMAGE_SPEC_VERSION = 1

var KHR_SHARED_PRESENTABLE_IMAGE_EXTENSION_NAME = "VK_KHR_shared_presentable_image"

// SharedPresentSurfaceCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html
type SharedPresentSurfaceCapabilitiesKHR struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	SharedPresentSupportedUsageFlags ImageUsageFlags
}

func NewSharedPresentSurfaceCapabilitiesKHR() *SharedPresentSurfaceCapabilitiesKHR {
	p := (*SharedPresentSurfaceCapabilitiesKHR)(MemAlloc(unsafe.Sizeof(*(*SharedPresentSurfaceCapabilitiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR
	return p
}
func (p *SharedPresentSurfaceCapabilitiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetSwapchainStatusKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetSwapchainStatusKHR.html
type PfnGetSwapchainStatusKHR uintptr

func (fn PfnGetSwapchainStatusKHR) Call(device Device, swapchain SwapchainKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain))
	return Result(ret)
}
func (fn PfnGetSwapchainStatusKHR) String() string { return "vkGetSwapchainStatusKHR" }

const KHR_external_fence_capabilities = 1
const KHR_EXTERNAL_FENCE_CAPABILITIES_SPEC_VERSION = 1

var KHR_EXTERNAL_FENCE_CAPABILITIES_EXTENSION_NAME = "VK_KHR_external_fence_capabilities"

type ExternalFenceHandleTypeFlagsKHR = ExternalFenceHandleTypeFlags
type ExternalFenceFeatureFlagsKHR = ExternalFenceFeatureFlags
type PhysicalDeviceExternalFenceInfoKHR = PhysicalDeviceExternalFenceInfo
type ExternalFencePropertiesKHR = ExternalFenceProperties

//  PfnGetPhysicalDeviceExternalFencePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceExternalFencePropertiesKHR.html
type PfnGetPhysicalDeviceExternalFencePropertiesKHR uintptr

func (fn PfnGetPhysicalDeviceExternalFencePropertiesKHR) Call(physicalDevice PhysicalDevice, pExternalFenceInfo *PhysicalDeviceExternalFenceInfo, pExternalFenceProperties *ExternalFenceProperties) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pExternalFenceInfo)), uintptr(unsafe.Pointer(pExternalFenceProperties)))
}
func (fn PfnGetPhysicalDeviceExternalFencePropertiesKHR) String() string {
	return "vkGetPhysicalDeviceExternalFencePropertiesKHR"
}

const KHR_external_fence = 1
const KHR_EXTERNAL_FENCE_SPEC_VERSION = 1

var KHR_EXTERNAL_FENCE_EXTENSION_NAME = "VK_KHR_external_fence"

type FenceImportFlagsKHR = FenceImportFlags
type ExportFenceCreateInfoKHR = ExportFenceCreateInfo

const KHR_external_fence_fd = 1
const KHR_EXTERNAL_FENCE_FD_SPEC_VERSION = 1

var KHR_EXTERNAL_FENCE_FD_EXTENSION_NAME = "VK_KHR_external_fence_fd"

// ImportFenceFdInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportFenceFdInfoKHR.html
type ImportFenceFdInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Fence      Fence
	Flags      FenceImportFlags
	HandleType ExternalFenceHandleTypeFlags
	Fd         int
}

func NewImportFenceFdInfoKHR() *ImportFenceFdInfoKHR {
	p := (*ImportFenceFdInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImportFenceFdInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR
	return p
}
func (p *ImportFenceFdInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// FenceGetFdInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFenceGetFdInfoKHR.html
type FenceGetFdInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Fence      Fence
	HandleType ExternalFenceHandleTypeFlags
}

func NewFenceGetFdInfoKHR() *FenceGetFdInfoKHR {
	p := (*FenceGetFdInfoKHR)(MemAlloc(unsafe.Sizeof(*(*FenceGetFdInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR
	return p
}
func (p *FenceGetFdInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnImportFenceFdKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkImportFenceFdKHR.html
type PfnImportFenceFdKHR uintptr

func (fn PfnImportFenceFdKHR) Call(device Device, pImportFenceFdInfo *ImportFenceFdInfoKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pImportFenceFdInfo)))
	return Result(ret)
}
func (fn PfnImportFenceFdKHR) String() string { return "vkImportFenceFdKHR" }

//  PfnGetFenceFdKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetFenceFdKHR.html
type PfnGetFenceFdKHR uintptr

func (fn PfnGetFenceFdKHR) Call(device Device, pGetFdInfo *FenceGetFdInfoKHR, pFd *int) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pGetFdInfo)), uintptr(unsafe.Pointer(pFd)))
	return Result(ret)
}
func (fn PfnGetFenceFdKHR) String() string { return "vkGetFenceFdKHR" }

const KHR_maintenance2 = 1
const KHR_MAINTENANCE2_SPEC_VERSION = 1

var KHR_MAINTENANCE2_EXTENSION_NAME = "VK_KHR_maintenance2"

type PointClippingBehaviorKHR = PointClippingBehavior
type TessellationDomainOriginKHR = TessellationDomainOrigin
type PhysicalDevicePointClippingPropertiesKHR = PhysicalDevicePointClippingProperties
type RenderPassInputAttachmentAspectCreateInfoKHR = RenderPassInputAttachmentAspectCreateInfo
type InputAttachmentAspectReferenceKHR = InputAttachmentAspectReference
type ImageViewUsageCreateInfoKHR = ImageViewUsageCreateInfo
type PipelineTessellationDomainOriginStateCreateInfoKHR = PipelineTessellationDomainOriginStateCreateInfo

const KHR_get_surface_capabilities2 = 1
const KHR_GET_SURFACE_CAPABILITIES_2_SPEC_VERSION = 1

var KHR_GET_SURFACE_CAPABILITIES_2_EXTENSION_NAME = "VK_KHR_get_surface_capabilities2"

// PhysicalDeviceSurfaceInfo2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html
type PhysicalDeviceSurfaceInfo2KHR struct {
	SType   StructureType
	PNext   unsafe.Pointer
	Surface SurfaceKHR
}

func NewPhysicalDeviceSurfaceInfo2KHR() *PhysicalDeviceSurfaceInfo2KHR {
	p := (*PhysicalDeviceSurfaceInfo2KHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSurfaceInfo2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR
	return p
}
func (p *PhysicalDeviceSurfaceInfo2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// SurfaceCapabilities2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceCapabilities2KHR.html
type SurfaceCapabilities2KHR struct {
	SType               StructureType
	PNext               unsafe.Pointer
	SurfaceCapabilities SurfaceCapabilitiesKHR
}

func NewSurfaceCapabilities2KHR() *SurfaceCapabilities2KHR {
	p := (*SurfaceCapabilities2KHR)(MemAlloc(unsafe.Sizeof(*(*SurfaceCapabilities2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR
	return p
}
func (p *SurfaceCapabilities2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// SurfaceFormat2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceFormat2KHR.html
type SurfaceFormat2KHR struct {
	SType         StructureType
	PNext         unsafe.Pointer
	SurfaceFormat SurfaceFormatKHR
}

func NewSurfaceFormat2KHR() *SurfaceFormat2KHR {
	p := (*SurfaceFormat2KHR)(MemAlloc(unsafe.Sizeof(*(*SurfaceFormat2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR
	return p
}
func (p *SurfaceFormat2KHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceSurfaceCapabilities2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html
type PfnGetPhysicalDeviceSurfaceCapabilities2KHR uintptr

func (fn PfnGetPhysicalDeviceSurfaceCapabilities2KHR) Call(physicalDevice PhysicalDevice, pSurfaceInfo *PhysicalDeviceSurfaceInfo2KHR, pSurfaceCapabilities *SurfaceCapabilities2KHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pSurfaceCapabilities)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfaceCapabilities2KHR) String() string {
	return "vkGetPhysicalDeviceSurfaceCapabilities2KHR"
}

//  PfnGetPhysicalDeviceSurfaceFormats2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html
type PfnGetPhysicalDeviceSurfaceFormats2KHR uintptr

func (fn PfnGetPhysicalDeviceSurfaceFormats2KHR) Call(physicalDevice PhysicalDevice, pSurfaceInfo *PhysicalDeviceSurfaceInfo2KHR, pSurfaceFormatCount *uint32, pSurfaceFormats *SurfaceFormat2KHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pSurfaceFormatCount)), uintptr(unsafe.Pointer(pSurfaceFormats)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfaceFormats2KHR) String() string {
	return "vkGetPhysicalDeviceSurfaceFormats2KHR"
}

const KHR_variable_pointers = 1
const KHR_VARIABLE_POINTERS_SPEC_VERSION = 1

var KHR_VARIABLE_POINTERS_EXTENSION_NAME = "VK_KHR_variable_pointers"

type PhysicalDeviceVariablePointerFeaturesKHR = PhysicalDeviceVariablePointersFeatures
type PhysicalDeviceVariablePointersFeaturesKHR = PhysicalDeviceVariablePointersFeatures

const KHR_get_display_properties2 = 1
const KHR_GET_DISPLAY_PROPERTIES_2_SPEC_VERSION = 1

var KHR_GET_DISPLAY_PROPERTIES_2_EXTENSION_NAME = "VK_KHR_get_display_properties2"

// DisplayProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayProperties2KHR.html
type DisplayProperties2KHR struct {
	SType             StructureType
	PNext             unsafe.Pointer
	DisplayProperties DisplayPropertiesKHR
}

func NewDisplayProperties2KHR() *DisplayProperties2KHR {
	p := (*DisplayProperties2KHR)(MemAlloc(unsafe.Sizeof(*(*DisplayProperties2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR
	return p
}
func (p *DisplayProperties2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayPlaneProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPlaneProperties2KHR.html
type DisplayPlaneProperties2KHR struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	DisplayPlaneProperties DisplayPlanePropertiesKHR
}

func NewDisplayPlaneProperties2KHR() *DisplayPlaneProperties2KHR {
	p := (*DisplayPlaneProperties2KHR)(MemAlloc(unsafe.Sizeof(*(*DisplayPlaneProperties2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR
	return p
}
func (p *DisplayPlaneProperties2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayModeProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayModeProperties2KHR.html
type DisplayModeProperties2KHR struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	DisplayModeProperties DisplayModePropertiesKHR
}

func NewDisplayModeProperties2KHR() *DisplayModeProperties2KHR {
	p := (*DisplayModeProperties2KHR)(MemAlloc(unsafe.Sizeof(*(*DisplayModeProperties2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR
	return p
}
func (p *DisplayModeProperties2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayPlaneInfo2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPlaneInfo2KHR.html
type DisplayPlaneInfo2KHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Mode       DisplayModeKHR
	PlaneIndex uint32
}

func NewDisplayPlaneInfo2KHR() *DisplayPlaneInfo2KHR {
	p := (*DisplayPlaneInfo2KHR)(MemAlloc(unsafe.Sizeof(*(*DisplayPlaneInfo2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR
	return p
}
func (p *DisplayPlaneInfo2KHR) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayPlaneCapabilities2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPlaneCapabilities2KHR.html
type DisplayPlaneCapabilities2KHR struct {
	SType        StructureType
	PNext        unsafe.Pointer
	Capabilities DisplayPlaneCapabilitiesKHR
}

func NewDisplayPlaneCapabilities2KHR() *DisplayPlaneCapabilities2KHR {
	p := (*DisplayPlaneCapabilities2KHR)(MemAlloc(unsafe.Sizeof(*(*DisplayPlaneCapabilities2KHR)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR
	return p
}
func (p *DisplayPlaneCapabilities2KHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceDisplayProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceDisplayProperties2KHR.html
type PfnGetPhysicalDeviceDisplayProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceDisplayProperties2KHR) Call(physicalDevice PhysicalDevice, pPropertyCount *uint32, pProperties *DisplayProperties2KHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceDisplayProperties2KHR) String() string {
	return "vkGetPhysicalDeviceDisplayProperties2KHR"
}

//  PfnGetPhysicalDeviceDisplayPlaneProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html
type PfnGetPhysicalDeviceDisplayPlaneProperties2KHR uintptr

func (fn PfnGetPhysicalDeviceDisplayPlaneProperties2KHR) Call(physicalDevice PhysicalDevice, pPropertyCount *uint32, pProperties *DisplayPlaneProperties2KHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceDisplayPlaneProperties2KHR) String() string {
	return "vkGetPhysicalDeviceDisplayPlaneProperties2KHR"
}

//  PfnGetDisplayModeProperties2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDisplayModeProperties2KHR.html
type PfnGetDisplayModeProperties2KHR uintptr

func (fn PfnGetDisplayModeProperties2KHR) Call(physicalDevice PhysicalDevice, display DisplayKHR, pPropertyCount *uint32, pProperties *DisplayModeProperties2KHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(display), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetDisplayModeProperties2KHR) String() string { return "vkGetDisplayModeProperties2KHR" }

//  PfnGetDisplayPlaneCapabilities2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDisplayPlaneCapabilities2KHR.html
type PfnGetDisplayPlaneCapabilities2KHR uintptr

func (fn PfnGetDisplayPlaneCapabilities2KHR) Call(physicalDevice PhysicalDevice, pDisplayPlaneInfo *DisplayPlaneInfo2KHR, pCapabilities *DisplayPlaneCapabilities2KHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pDisplayPlaneInfo)), uintptr(unsafe.Pointer(pCapabilities)))
	return Result(ret)
}
func (fn PfnGetDisplayPlaneCapabilities2KHR) String() string {
	return "vkGetDisplayPlaneCapabilities2KHR"
}

const KHR_dedicated_allocation = 1
const KHR_DEDICATED_ALLOCATION_SPEC_VERSION = 3

var KHR_DEDICATED_ALLOCATION_EXTENSION_NAME = "VK_KHR_dedicated_allocation"

type MemoryDedicatedRequirementsKHR = MemoryDedicatedRequirements
type MemoryDedicatedAllocateInfoKHR = MemoryDedicatedAllocateInfo

const KHR_storage_buffer_storage_class = 1
const KHR_STORAGE_BUFFER_STORAGE_CLASS_SPEC_VERSION = 1

var KHR_STORAGE_BUFFER_STORAGE_CLASS_EXTENSION_NAME = "VK_KHR_storage_buffer_storage_class"

const KHR_relaxed_block_layout = 1
const KHR_RELAXED_BLOCK_LAYOUT_SPEC_VERSION = 1

var KHR_RELAXED_BLOCK_LAYOUT_EXTENSION_NAME = "VK_KHR_relaxed_block_layout"

const KHR_get_memory_requirements2 = 1
const KHR_GET_MEMORY_REQUIREMENTS_2_SPEC_VERSION = 1

var KHR_GET_MEMORY_REQUIREMENTS_2_EXTENSION_NAME = "VK_KHR_get_memory_requirements2"

type BufferMemoryRequirementsInfo2KHR = BufferMemoryRequirementsInfo2
type ImageMemoryRequirementsInfo2KHR = ImageMemoryRequirementsInfo2
type ImageSparseMemoryRequirementsInfo2KHR = ImageSparseMemoryRequirementsInfo2
type SparseImageMemoryRequirements2KHR = SparseImageMemoryRequirements2

//  PfnGetImageMemoryRequirements2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageMemoryRequirements2KHR.html
type PfnGetImageMemoryRequirements2KHR uintptr

func (fn PfnGetImageMemoryRequirements2KHR) Call(device Device, pInfo *ImageMemoryRequirementsInfo2, pMemoryRequirements *MemoryRequirements2) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))
}
func (fn PfnGetImageMemoryRequirements2KHR) String() string {
	return "vkGetImageMemoryRequirements2KHR"
}

//  PfnGetBufferMemoryRequirements2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetBufferMemoryRequirements2KHR.html
type PfnGetBufferMemoryRequirements2KHR uintptr

func (fn PfnGetBufferMemoryRequirements2KHR) Call(device Device, pInfo *BufferMemoryRequirementsInfo2, pMemoryRequirements *MemoryRequirements2) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))
}
func (fn PfnGetBufferMemoryRequirements2KHR) String() string {
	return "vkGetBufferMemoryRequirements2KHR"
}

//  PfnGetImageSparseMemoryRequirements2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageSparseMemoryRequirements2KHR.html
type PfnGetImageSparseMemoryRequirements2KHR uintptr

func (fn PfnGetImageSparseMemoryRequirements2KHR) Call(device Device, pInfo *ImageSparseMemoryRequirementsInfo2, pSparseMemoryRequirementCount *uint32, pSparseMemoryRequirements *SparseImageMemoryRequirements2) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pSparseMemoryRequirementCount)), uintptr(unsafe.Pointer(pSparseMemoryRequirements)))
}
func (fn PfnGetImageSparseMemoryRequirements2KHR) String() string {
	return "vkGetImageSparseMemoryRequirements2KHR"
}

const KHR_image_format_list = 1
const KHR_IMAGE_FORMAT_LIST_SPEC_VERSION = 1

var KHR_IMAGE_FORMAT_LIST_EXTENSION_NAME = "VK_KHR_image_format_list"

// ImageFormatListCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageFormatListCreateInfoKHR.html
type ImageFormatListCreateInfoKHR struct {
	SType           StructureType
	PNext           unsafe.Pointer
	ViewFormatCount uint32
	PViewFormats    *Format
}

func NewImageFormatListCreateInfoKHR() *ImageFormatListCreateInfoKHR {
	p := (*ImageFormatListCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImageFormatListCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR
	return p
}
func (p *ImageFormatListCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_sampler_ycbcr_conversion = 1

type SamplerYcbcrConversionKHR = SamplerYcbcrConversion

const KHR_SAMPLER_YCBCR_CONVERSION_SPEC_VERSION = 14

var KHR_SAMPLER_YCBCR_CONVERSION_EXTENSION_NAME = "VK_KHR_sampler_ycbcr_conversion"

type SamplerYcbcrModelConversionKHR = SamplerYcbcrModelConversion
type SamplerYcbcrRangeKHR = SamplerYcbcrRange
type ChromaLocationKHR = ChromaLocation
type SamplerYcbcrConversionCreateInfoKHR = SamplerYcbcrConversionCreateInfo
type SamplerYcbcrConversionInfoKHR = SamplerYcbcrConversionInfo
type BindImagePlaneMemoryInfoKHR = BindImagePlaneMemoryInfo
type ImagePlaneMemoryRequirementsInfoKHR = ImagePlaneMemoryRequirementsInfo
type PhysicalDeviceSamplerYcbcrConversionFeaturesKHR = PhysicalDeviceSamplerYcbcrConversionFeatures
type SamplerYcbcrConversionImageFormatPropertiesKHR = SamplerYcbcrConversionImageFormatProperties

//  PfnCreateSamplerYcbcrConversionKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateSamplerYcbcrConversionKHR.html
type PfnCreateSamplerYcbcrConversionKHR uintptr

func (fn PfnCreateSamplerYcbcrConversionKHR) Call(device Device, pCreateInfo *SamplerYcbcrConversionCreateInfo, pAllocator *AllocationCallbacks, pYcbcrConversion *SamplerYcbcrConversion) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pYcbcrConversion)))
	return Result(ret)
}
func (fn PfnCreateSamplerYcbcrConversionKHR) String() string {
	return "vkCreateSamplerYcbcrConversionKHR"
}

//  PfnDestroySamplerYcbcrConversionKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroySamplerYcbcrConversionKHR.html
type PfnDestroySamplerYcbcrConversionKHR uintptr

func (fn PfnDestroySamplerYcbcrConversionKHR) Call(device Device, ycbcrConversion SamplerYcbcrConversion, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(ycbcrConversion), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroySamplerYcbcrConversionKHR) String() string {
	return "vkDestroySamplerYcbcrConversionKHR"
}

const KHR_bind_memory2 = 1
const KHR_BIND_MEMORY_2_SPEC_VERSION = 1

var KHR_BIND_MEMORY_2_EXTENSION_NAME = "VK_KHR_bind_memory2"

type BindBufferMemoryInfoKHR = BindBufferMemoryInfo
type BindImageMemoryInfoKHR = BindImageMemoryInfo

//  PfnBindBufferMemory2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBindBufferMemory2KHR.html
type PfnBindBufferMemory2KHR uintptr

func (fn PfnBindBufferMemory2KHR) Call(device Device, bindInfoCount uint32, pBindInfos *BindBufferMemoryInfo) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))
	return Result(ret)
}
func (fn PfnBindBufferMemory2KHR) String() string { return "vkBindBufferMemory2KHR" }

//  PfnBindImageMemory2KHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBindImageMemory2KHR.html
type PfnBindImageMemory2KHR uintptr

func (fn PfnBindImageMemory2KHR) Call(device Device, bindInfoCount uint32, pBindInfos *BindImageMemoryInfo) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))
	return Result(ret)
}
func (fn PfnBindImageMemory2KHR) String() string { return "vkBindImageMemory2KHR" }

const KHR_maintenance3 = 1
const KHR_MAINTENANCE3_SPEC_VERSION = 1

var KHR_MAINTENANCE3_EXTENSION_NAME = "VK_KHR_maintenance3"

type PhysicalDeviceMaintenance3PropertiesKHR = PhysicalDeviceMaintenance3Properties
type DescriptorSetLayoutSupportKHR = DescriptorSetLayoutSupport

//  PfnGetDescriptorSetLayoutSupportKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDescriptorSetLayoutSupportKHR.html
type PfnGetDescriptorSetLayoutSupportKHR uintptr

func (fn PfnGetDescriptorSetLayoutSupportKHR) Call(device Device, pCreateInfo *DescriptorSetLayoutCreateInfo, pSupport *DescriptorSetLayoutSupport) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pSupport)))
}
func (fn PfnGetDescriptorSetLayoutSupportKHR) String() string {
	return "vkGetDescriptorSetLayoutSupportKHR"
}

const KHR_draw_indirect_count = 1
const KHR_DRAW_INDIRECT_COUNT_SPEC_VERSION = 1

var KHR_DRAW_INDIRECT_COUNT_EXTENSION_NAME = "VK_KHR_draw_indirect_count"

//  PfnCmdDrawIndirectCountKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndirectCountKHR.html
type PfnCmdDrawIndirectCountKHR uintptr

func (fn PfnCmdDrawIndirectCountKHR) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride))
}
func (fn PfnCmdDrawIndirectCountKHR) String() string { return "vkCmdDrawIndirectCountKHR" }

//  PfnCmdDrawIndexedIndirectCountKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndexedIndirectCountKHR.html
type PfnCmdDrawIndexedIndirectCountKHR uintptr

func (fn PfnCmdDrawIndexedIndirectCountKHR) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride))
}
func (fn PfnCmdDrawIndexedIndirectCountKHR) String() string {
	return "vkCmdDrawIndexedIndirectCountKHR"
}

const KHR_8bit_storage = 1
const KHR_8BIT_STORAGE_SPEC_VERSION = 1

var KHR_8BIT_STORAGE_EXTENSION_NAME = "VK_KHR_8bit_storage"

// PhysicalDevice8BitStorageFeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDevice8BitStorageFeaturesKHR.html
type PhysicalDevice8BitStorageFeaturesKHR struct {
	SType                             StructureType
	PNext                             unsafe.Pointer
	StorageBuffer8BitAccess           Bool32
	UniformAndStorageBuffer8BitAccess Bool32
	StoragePushConstant8              Bool32
}

func NewPhysicalDevice8BitStorageFeaturesKHR() *PhysicalDevice8BitStorageFeaturesKHR {
	p := (*PhysicalDevice8BitStorageFeaturesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDevice8BitStorageFeaturesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR
	return p
}
func (p *PhysicalDevice8BitStorageFeaturesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_shader_atomic_int64 = 1
const KHR_SHADER_ATOMIC_INT64_SPEC_VERSION = 1

var KHR_SHADER_ATOMIC_INT64_EXTENSION_NAME = "VK_KHR_shader_atomic_int64"

// PhysicalDeviceShaderAtomicInt64FeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderAtomicInt64FeaturesKHR.html
type PhysicalDeviceShaderAtomicInt64FeaturesKHR struct {
	SType                    StructureType
	PNext                    unsafe.Pointer
	ShaderBufferInt64Atomics Bool32
	ShaderSharedInt64Atomics Bool32
}

func NewPhysicalDeviceShaderAtomicInt64FeaturesKHR() *PhysicalDeviceShaderAtomicInt64FeaturesKHR {
	p := (*PhysicalDeviceShaderAtomicInt64FeaturesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderAtomicInt64FeaturesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR
	return p
}
func (p *PhysicalDeviceShaderAtomicInt64FeaturesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_driver_properties = 1
const MAX_DRIVER_NAME_SIZE_KHR = 256
const MAX_DRIVER_INFO_SIZE_KHR = 256
const KHR_DRIVER_PROPERTIES_SPEC_VERSION = 1

var KHR_DRIVER_PROPERTIES_EXTENSION_NAME = "VK_KHR_driver_properties"

// DriverIdKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDriverIdKHR.html
type DriverIdKHR int32

const (
	DRIVER_ID_AMD_PROPRIETARY_KHR           DriverIdKHR = 1
	DRIVER_ID_AMD_OPEN_SOURCE_KHR           DriverIdKHR = 2
	DRIVER_ID_MESA_RADV_KHR                 DriverIdKHR = 3
	DRIVER_ID_NVIDIA_PROPRIETARY_KHR        DriverIdKHR = 4
	DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR DriverIdKHR = 5
	DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR    DriverIdKHR = 6
	DRIVER_ID_IMAGINATION_PROPRIETARY_KHR   DriverIdKHR = 7
	DRIVER_ID_QUALCOMM_PROPRIETARY_KHR      DriverIdKHR = 8
	DRIVER_ID_ARM_PROPRIETARY_KHR           DriverIdKHR = 9
	DRIVER_ID_GOOGLE_SWIFTSHADER_KHR        DriverIdKHR = 10
	DRIVER_ID_GGP_PROPRIETARY_KHR           DriverIdKHR = 11
	DRIVER_ID_BROADCOM_PROPRIETARY_KHR      DriverIdKHR = 12
	DRIVER_ID_BEGIN_RANGE_KHR               DriverIdKHR = DRIVER_ID_AMD_PROPRIETARY_KHR
	DRIVER_ID_END_RANGE_KHR                 DriverIdKHR = DRIVER_ID_BROADCOM_PROPRIETARY_KHR
	DRIVER_ID_RANGE_SIZE_KHR                DriverIdKHR = (DRIVER_ID_BROADCOM_PROPRIETARY_KHR - DRIVER_ID_AMD_PROPRIETARY_KHR + 1)
	DRIVER_ID_MAX_ENUM_KHR                  DriverIdKHR = 0x7FFFFFFF
)

func (x DriverIdKHR) String() string {
	switch x {
	case DRIVER_ID_AMD_PROPRIETARY_KHR:
		return "DRIVER_ID_AMD_PROPRIETARY_KHR"
	case DRIVER_ID_AMD_OPEN_SOURCE_KHR:
		return "DRIVER_ID_AMD_OPEN_SOURCE_KHR"
	case DRIVER_ID_MESA_RADV_KHR:
		return "DRIVER_ID_MESA_RADV_KHR"
	case DRIVER_ID_NVIDIA_PROPRIETARY_KHR:
		return "DRIVER_ID_NVIDIA_PROPRIETARY_KHR"
	case DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR:
		return "DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR"
	case DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR:
		return "DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR"
	case DRIVER_ID_IMAGINATION_PROPRIETARY_KHR:
		return "DRIVER_ID_IMAGINATION_PROPRIETARY_KHR"
	case DRIVER_ID_QUALCOMM_PROPRIETARY_KHR:
		return "DRIVER_ID_QUALCOMM_PROPRIETARY_KHR"
	case DRIVER_ID_ARM_PROPRIETARY_KHR:
		return "DRIVER_ID_ARM_PROPRIETARY_KHR"
	case DRIVER_ID_GOOGLE_SWIFTSHADER_KHR:
		return "DRIVER_ID_GOOGLE_SWIFTSHADER_KHR"
	case DRIVER_ID_GGP_PROPRIETARY_KHR:
		return "DRIVER_ID_GGP_PROPRIETARY_KHR"
	case DRIVER_ID_BROADCOM_PROPRIETARY_KHR:
		return "DRIVER_ID_BROADCOM_PROPRIETARY_KHR"
	case DRIVER_ID_MAX_ENUM_KHR:
		return "DRIVER_ID_MAX_ENUM_KHR"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ConformanceVersionKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkConformanceVersionKHR.html
type ConformanceVersionKHR struct {
	Major    uint8
	Minor    uint8
	Subminor uint8
	Patch    uint8
}

func NewConformanceVersionKHR() *ConformanceVersionKHR {
	return (*ConformanceVersionKHR)(MemAlloc(unsafe.Sizeof(*(*ConformanceVersionKHR)(nil))))
}
func (p *ConformanceVersionKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceDriverPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceDriverPropertiesKHR.html
type PhysicalDeviceDriverPropertiesKHR struct {
	SType              StructureType
	PNext              unsafe.Pointer
	DriverID           DriverIdKHR
	DriverName         [MAX_DRIVER_NAME_SIZE_KHR]int8
	DriverInfo         [MAX_DRIVER_INFO_SIZE_KHR]int8
	ConformanceVersion ConformanceVersionKHR
}

func NewPhysicalDeviceDriverPropertiesKHR() *PhysicalDeviceDriverPropertiesKHR {
	p := (*PhysicalDeviceDriverPropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceDriverPropertiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR
	return p
}
func (p *PhysicalDeviceDriverPropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_shader_float_controls = 1
const KHR_SHADER_FLOAT_CONTROLS_SPEC_VERSION = 4

var KHR_SHADER_FLOAT_CONTROLS_EXTENSION_NAME = "VK_KHR_shader_float_controls"

// ShaderFloatControlsIndependenceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderFloatControlsIndependenceKHR.html
type ShaderFloatControlsIndependenceKHR int32

const (
	SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR ShaderFloatControlsIndependenceKHR = 0
	SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR         ShaderFloatControlsIndependenceKHR = 1
	SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR        ShaderFloatControlsIndependenceKHR = 2
	SHADER_FLOAT_CONTROLS_INDEPENDENCE_BEGIN_RANGE_KHR ShaderFloatControlsIndependenceKHR = SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR
	SHADER_FLOAT_CONTROLS_INDEPENDENCE_END_RANGE_KHR   ShaderFloatControlsIndependenceKHR = SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR
	SHADER_FLOAT_CONTROLS_INDEPENDENCE_RANGE_SIZE_KHR  ShaderFloatControlsIndependenceKHR = (SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR - SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR + 1)
	SHADER_FLOAT_CONTROLS_INDEPENDENCE_MAX_ENUM_KHR    ShaderFloatControlsIndependenceKHR = 0x7FFFFFFF
)

func (x ShaderFloatControlsIndependenceKHR) String() string {
	switch x {
	case SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR:
		return "SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR"
	case SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR:
		return "SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR"
	case SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR:
		return "SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR"
	case SHADER_FLOAT_CONTROLS_INDEPENDENCE_MAX_ENUM_KHR:
		return "SHADER_FLOAT_CONTROLS_INDEPENDENCE_MAX_ENUM_KHR"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PhysicalDeviceFloatControlsPropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceFloatControlsPropertiesKHR.html
type PhysicalDeviceFloatControlsPropertiesKHR struct {
	SType                                 StructureType
	PNext                                 unsafe.Pointer
	DenormBehaviorIndependence            ShaderFloatControlsIndependenceKHR
	RoundingModeIndependence              ShaderFloatControlsIndependenceKHR
	ShaderSignedZeroInfNanPreserveFloat16 Bool32
	ShaderSignedZeroInfNanPreserveFloat32 Bool32
	ShaderSignedZeroInfNanPreserveFloat64 Bool32
	ShaderDenormPreserveFloat16           Bool32
	ShaderDenormPreserveFloat32           Bool32
	ShaderDenormPreserveFloat64           Bool32
	ShaderDenormFlushToZeroFloat16        Bool32
	ShaderDenormFlushToZeroFloat32        Bool32
	ShaderDenormFlushToZeroFloat64        Bool32
	ShaderRoundingModeRTEFloat16          Bool32
	ShaderRoundingModeRTEFloat32          Bool32
	ShaderRoundingModeRTEFloat64          Bool32
	ShaderRoundingModeRTZFloat16          Bool32
	ShaderRoundingModeRTZFloat32          Bool32
	ShaderRoundingModeRTZFloat64          Bool32
}

func NewPhysicalDeviceFloatControlsPropertiesKHR() *PhysicalDeviceFloatControlsPropertiesKHR {
	p := (*PhysicalDeviceFloatControlsPropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceFloatControlsPropertiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR
	return p
}
func (p *PhysicalDeviceFloatControlsPropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_depth_stencil_resolve = 1
const KHR_DEPTH_STENCIL_RESOLVE_SPEC_VERSION = 1

var KHR_DEPTH_STENCIL_RESOLVE_EXTENSION_NAME = "VK_KHR_depth_stencil_resolve"

// ResolveModeFlagsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkResolveModeFlagsKHR.html
type ResolveModeFlagsKHR uint32

const (
	RESOLVE_MODE_NONE_KHR               ResolveModeFlagsKHR = 0
	RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR    ResolveModeFlagsKHR = 0x00000001
	RESOLVE_MODE_AVERAGE_BIT_KHR        ResolveModeFlagsKHR = 0x00000002
	RESOLVE_MODE_MIN_BIT_KHR            ResolveModeFlagsKHR = 0x00000004
	RESOLVE_MODE_MAX_BIT_KHR            ResolveModeFlagsKHR = 0x00000008
	RESOLVE_MODE_FLAG_BITS_MAX_ENUM_KHR ResolveModeFlagsKHR = 0x7FFFFFFF
)

func (x ResolveModeFlagsKHR) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ResolveModeFlagsKHR(1 << i) {
			case RESOLVE_MODE_NONE_KHR:
				s += "RESOLVE_MODE_NONE_KHR|"
			case RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR:
				s += "RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR|"
			case RESOLVE_MODE_AVERAGE_BIT_KHR:
				s += "RESOLVE_MODE_AVERAGE_BIT_KHR|"
			case RESOLVE_MODE_MIN_BIT_KHR:
				s += "RESOLVE_MODE_MIN_BIT_KHR|"
			case RESOLVE_MODE_MAX_BIT_KHR:
				s += "RESOLVE_MODE_MAX_BIT_KHR|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SubpassDescriptionDepthStencilResolveKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassDescriptionDepthStencilResolveKHR.html
type SubpassDescriptionDepthStencilResolveKHR struct {
	SType                          StructureType
	PNext                          unsafe.Pointer
	DepthResolveMode               ResolveModeFlagsKHR
	StencilResolveMode             ResolveModeFlagsKHR
	PDepthStencilResolveAttachment *AttachmentReference2KHR
}

func NewSubpassDescriptionDepthStencilResolveKHR() *SubpassDescriptionDepthStencilResolveKHR {
	p := (*SubpassDescriptionDepthStencilResolveKHR)(MemAlloc(unsafe.Sizeof(*(*SubpassDescriptionDepthStencilResolveKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR
	return p
}
func (p *SubpassDescriptionDepthStencilResolveKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceDepthStencilResolvePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceDepthStencilResolvePropertiesKHR.html
type PhysicalDeviceDepthStencilResolvePropertiesKHR struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	SupportedDepthResolveModes   ResolveModeFlagsKHR
	SupportedStencilResolveModes ResolveModeFlagsKHR
	IndependentResolveNone       Bool32
	IndependentResolve           Bool32
}

func NewPhysicalDeviceDepthStencilResolvePropertiesKHR() *PhysicalDeviceDepthStencilResolvePropertiesKHR {
	p := (*PhysicalDeviceDepthStencilResolvePropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceDepthStencilResolvePropertiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR
	return p
}
func (p *PhysicalDeviceDepthStencilResolvePropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_swapchain_mutable_format = 1
const KHR_SWAPCHAIN_MUTABLE_FORMAT_SPEC_VERSION = 1

var KHR_SWAPCHAIN_MUTABLE_FORMAT_EXTENSION_NAME = "VK_KHR_swapchain_mutable_format"

const KHR_vulkan_memory_model = 1
const KHR_VULKAN_MEMORY_MODEL_SPEC_VERSION = 3

var KHR_VULKAN_MEMORY_MODEL_EXTENSION_NAME = "VK_KHR_vulkan_memory_model"

// PhysicalDeviceVulkanMemoryModelFeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceVulkanMemoryModelFeaturesKHR.html
type PhysicalDeviceVulkanMemoryModelFeaturesKHR struct {
	SType                                         StructureType
	PNext                                         unsafe.Pointer
	VulkanMemoryModel                             Bool32
	VulkanMemoryModelDeviceScope                  Bool32
	VulkanMemoryModelAvailabilityVisibilityChains Bool32
}

func NewPhysicalDeviceVulkanMemoryModelFeaturesKHR() *PhysicalDeviceVulkanMemoryModelFeaturesKHR {
	p := (*PhysicalDeviceVulkanMemoryModelFeaturesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceVulkanMemoryModelFeaturesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
	return p
}
func (p *PhysicalDeviceVulkanMemoryModelFeaturesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_surface_protected_capabilities = 1
const KHR_SURFACE_PROTECTED_CAPABILITIES_SPEC_VERSION = 1

var KHR_SURFACE_PROTECTED_CAPABILITIES_EXTENSION_NAME = "VK_KHR_surface_protected_capabilities"

// SurfaceProtectedCapabilitiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceProtectedCapabilitiesKHR.html
type SurfaceProtectedCapabilitiesKHR struct {
	SType             StructureType
	PNext             unsafe.Pointer
	SupportsProtected Bool32
}

func NewSurfaceProtectedCapabilitiesKHR() *SurfaceProtectedCapabilitiesKHR {
	p := (*SurfaceProtectedCapabilitiesKHR)(MemAlloc(unsafe.Sizeof(*(*SurfaceProtectedCapabilitiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR
	return p
}
func (p *SurfaceProtectedCapabilitiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_uniform_buffer_standard_layout = 1
const KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_SPEC_VERSION = 1

var KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_EXTENSION_NAME = "VK_KHR_uniform_buffer_standard_layout"

// PhysicalDeviceUniformBufferStandardLayoutFeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR.html
type PhysicalDeviceUniformBufferStandardLayoutFeaturesKHR struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	UniformBufferStandardLayout Bool32
}

func NewPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR() *PhysicalDeviceUniformBufferStandardLayoutFeaturesKHR {
	p := (*PhysicalDeviceUniformBufferStandardLayoutFeaturesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceUniformBufferStandardLayoutFeaturesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR
	return p
}
func (p *PhysicalDeviceUniformBufferStandardLayoutFeaturesKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_pipeline_executable_properties = 1
const KHR_PIPELINE_EXECUTABLE_PROPERTIES_SPEC_VERSION = 1

var KHR_PIPELINE_EXECUTABLE_PROPERTIES_EXTENSION_NAME = "VK_KHR_pipeline_executable_properties"

// PipelineExecutableStatisticFormatKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineExecutableStatisticFormatKHR.html
type PipelineExecutableStatisticFormatKHR int32

const (
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR      PipelineExecutableStatisticFormatKHR = 0
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR       PipelineExecutableStatisticFormatKHR = 1
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR      PipelineExecutableStatisticFormatKHR = 2
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR     PipelineExecutableStatisticFormatKHR = 3
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BEGIN_RANGE_KHR PipelineExecutableStatisticFormatKHR = PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_END_RANGE_KHR   PipelineExecutableStatisticFormatKHR = PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_RANGE_SIZE_KHR  PipelineExecutableStatisticFormatKHR = (PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR - PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR + 1)
	PIPELINE_EXECUTABLE_STATISTIC_FORMAT_MAX_ENUM_KHR    PipelineExecutableStatisticFormatKHR = 0x7FFFFFFF
)

func (x PipelineExecutableStatisticFormatKHR) String() string {
	switch x {
	case PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR:
		return "PIPELINE_EXECUTABLE_STATISTIC_FORMAT_BOOL32_KHR"
	case PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR:
		return "PIPELINE_EXECUTABLE_STATISTIC_FORMAT_INT64_KHR"
	case PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR:
		return "PIPELINE_EXECUTABLE_STATISTIC_FORMAT_UINT64_KHR"
	case PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR:
		return "PIPELINE_EXECUTABLE_STATISTIC_FORMAT_FLOAT64_KHR"
	case PIPELINE_EXECUTABLE_STATISTIC_FORMAT_MAX_ENUM_KHR:
		return "PIPELINE_EXECUTABLE_STATISTIC_FORMAT_MAX_ENUM_KHR"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PhysicalDevicePipelineExecutablePropertiesFeaturesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR.html
type PhysicalDevicePipelineExecutablePropertiesFeaturesKHR struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	PipelineExecutableInfo Bool32
}

func NewPhysicalDevicePipelineExecutablePropertiesFeaturesKHR() *PhysicalDevicePipelineExecutablePropertiesFeaturesKHR {
	return (*PhysicalDevicePipelineExecutablePropertiesFeaturesKHR)(MemAlloc(unsafe.Sizeof(*(*PhysicalDevicePipelineExecutablePropertiesFeaturesKHR)(nil))))
}
func (p *PhysicalDevicePipelineExecutablePropertiesFeaturesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineInfoKHR.html
type PipelineInfoKHR struct {
	SType    StructureType
	PNext    unsafe.Pointer
	Pipeline Pipeline
}

func NewPipelineInfoKHR() *PipelineInfoKHR {
	return (*PipelineInfoKHR)(MemAlloc(unsafe.Sizeof(*(*PipelineInfoKHR)(nil))))
}
func (p *PipelineInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineExecutablePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineExecutablePropertiesKHR.html
type PipelineExecutablePropertiesKHR struct {
	SType        StructureType
	PNext        unsafe.Pointer
	Stages       ShaderStageFlags
	Name         [MAX_DESCRIPTION_SIZE]int8
	Description  [MAX_DESCRIPTION_SIZE]int8
	SubgroupSize uint32
}

func NewPipelineExecutablePropertiesKHR() *PipelineExecutablePropertiesKHR {
	return (*PipelineExecutablePropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*PipelineExecutablePropertiesKHR)(nil))))
}
func (p *PipelineExecutablePropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineExecutableInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineExecutableInfoKHR.html
type PipelineExecutableInfoKHR struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Pipeline        Pipeline
	ExecutableIndex uint32
}

func NewPipelineExecutableInfoKHR() *PipelineExecutableInfoKHR {
	return (*PipelineExecutableInfoKHR)(MemAlloc(unsafe.Sizeof(*(*PipelineExecutableInfoKHR)(nil))))
}
func (p *PipelineExecutableInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineExecutableStatisticKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineExecutableStatisticKHR.html
type PipelineExecutableStatisticKHR struct {
	SType       StructureType
	PNext       unsafe.Pointer
	Name        [MAX_DESCRIPTION_SIZE]int8
	Description [MAX_DESCRIPTION_SIZE]int8
	Format      PipelineExecutableStatisticFormatKHR
	Value       PipelineExecutableStatisticValueKHR
}

func NewPipelineExecutableStatisticKHR() *PipelineExecutableStatisticKHR {
	return (*PipelineExecutableStatisticKHR)(MemAlloc(unsafe.Sizeof(*(*PipelineExecutableStatisticKHR)(nil))))
}
func (p *PipelineExecutableStatisticKHR) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineExecutableInternalRepresentationKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineExecutableInternalRepresentationKHR.html
type PipelineExecutableInternalRepresentationKHR struct {
	SType       StructureType
	PNext       unsafe.Pointer
	Name        [MAX_DESCRIPTION_SIZE]int8
	Description [MAX_DESCRIPTION_SIZE]int8
	IsText      Bool32
	DataSize    uintptr
	PData       unsafe.Pointer
}

func NewPipelineExecutableInternalRepresentationKHR() *PipelineExecutableInternalRepresentationKHR {
	return (*PipelineExecutableInternalRepresentationKHR)(MemAlloc(unsafe.Sizeof(*(*PipelineExecutableInternalRepresentationKHR)(nil))))
}
func (p *PipelineExecutableInternalRepresentationKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPipelineExecutablePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPipelineExecutablePropertiesKHR.html
type PfnGetPipelineExecutablePropertiesKHR uintptr

func (fn PfnGetPipelineExecutablePropertiesKHR) Call(device Device, pPipelineInfo *PipelineInfoKHR, pExecutableCount *uint32, pProperties *PipelineExecutablePropertiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pPipelineInfo)), uintptr(unsafe.Pointer(pExecutableCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetPipelineExecutablePropertiesKHR) String() string {
	return "vkGetPipelineExecutablePropertiesKHR"
}

//  PfnGetPipelineExecutableStatisticsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPipelineExecutableStatisticsKHR.html
type PfnGetPipelineExecutableStatisticsKHR uintptr

func (fn PfnGetPipelineExecutableStatisticsKHR) Call(device Device, pExecutableInfo *PipelineExecutableInfoKHR, pStatisticCount *uint32, pStatistics *PipelineExecutableStatisticKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pExecutableInfo)), uintptr(unsafe.Pointer(pStatisticCount)), uintptr(unsafe.Pointer(pStatistics)))
	return Result(ret)
}
func (fn PfnGetPipelineExecutableStatisticsKHR) String() string {
	return "vkGetPipelineExecutableStatisticsKHR"
}

//  PfnGetPipelineExecutableInternalRepresentationsKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPipelineExecutableInternalRepresentationsKHR.html
type PfnGetPipelineExecutableInternalRepresentationsKHR uintptr

func (fn PfnGetPipelineExecutableInternalRepresentationsKHR) Call(device Device, pExecutableInfo *PipelineExecutableInfoKHR, pInternalRepresentationCount *uint32, pInternalRepresentations *PipelineExecutableInternalRepresentationKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pExecutableInfo)), uintptr(unsafe.Pointer(pInternalRepresentationCount)), uintptr(unsafe.Pointer(pInternalRepresentations)))
	return Result(ret)
}
func (fn PfnGetPipelineExecutableInternalRepresentationsKHR) String() string {
	return "vkGetPipelineExecutableInternalRepresentationsKHR"
}

const EXT_debug_report = 1

// DebugReportCallbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugReportCallbackEXT.html
type DebugReportCallbackEXT NonDispatchableHandle

const EXT_DEBUG_REPORT_SPEC_VERSION = 9

var EXT_DEBUG_REPORT_EXTENSION_NAME = "VK_EXT_debug_report"

// DebugReportObjectTypeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugReportObjectTypeEXT.html
type DebugReportObjectTypeEXT int32

const (
	DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT                        DebugReportObjectTypeEXT = 0
	DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT                       DebugReportObjectTypeEXT = 1
	DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT                DebugReportObjectTypeEXT = 2
	DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT                         DebugReportObjectTypeEXT = 3
	DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT                          DebugReportObjectTypeEXT = 4
	DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT                      DebugReportObjectTypeEXT = 5
	DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT                 DebugReportObjectTypeEXT = 6
	DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT                          DebugReportObjectTypeEXT = 7
	DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT                  DebugReportObjectTypeEXT = 8
	DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT                         DebugReportObjectTypeEXT = 9
	DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT                          DebugReportObjectTypeEXT = 10
	DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT                          DebugReportObjectTypeEXT = 11
	DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT                     DebugReportObjectTypeEXT = 12
	DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT                    DebugReportObjectTypeEXT = 13
	DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT                     DebugReportObjectTypeEXT = 14
	DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT                  DebugReportObjectTypeEXT = 15
	DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT                 DebugReportObjectTypeEXT = 16
	DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT                DebugReportObjectTypeEXT = 17
	DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT                    DebugReportObjectTypeEXT = 18
	DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT                       DebugReportObjectTypeEXT = 19
	DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT          DebugReportObjectTypeEXT = 20
	DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT                        DebugReportObjectTypeEXT = 21
	DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT                DebugReportObjectTypeEXT = 22
	DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT                 DebugReportObjectTypeEXT = 23
	DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT                    DebugReportObjectTypeEXT = 24
	DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT                   DebugReportObjectTypeEXT = 25
	DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT                    DebugReportObjectTypeEXT = 26
	DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT                  DebugReportObjectTypeEXT = 27
	DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT      DebugReportObjectTypeEXT = 28
	DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT                    DebugReportObjectTypeEXT = 29
	DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT               DebugReportObjectTypeEXT = 30
	DEBUG_REPORT_OBJECT_TYPE_OBJECT_TABLE_NVX_EXT               DebugReportObjectTypeEXT = 31
	DEBUG_REPORT_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NVX_EXT   DebugReportObjectTypeEXT = 32
	DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT           DebugReportObjectTypeEXT = 33
	DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT       DebugReportObjectTypeEXT = 1000156000
	DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT     DebugReportObjectTypeEXT = 1000085000
	DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT      DebugReportObjectTypeEXT = 1000165000
	DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_EXT                   DebugReportObjectTypeEXT = DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT
	DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT               DebugReportObjectTypeEXT = DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT
	DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR_EXT DebugReportObjectTypeEXT = DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT
	DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR_EXT   DebugReportObjectTypeEXT = DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT
	DEBUG_REPORT_OBJECT_TYPE_BEGIN_RANGE_EXT                    DebugReportObjectTypeEXT = DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT
	DEBUG_REPORT_OBJECT_TYPE_END_RANGE_EXT                      DebugReportObjectTypeEXT = DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT
	DEBUG_REPORT_OBJECT_TYPE_RANGE_SIZE_EXT                     DebugReportObjectTypeEXT = (DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT - DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT + 1)
	DEBUG_REPORT_OBJECT_TYPE_MAX_ENUM_EXT                       DebugReportObjectTypeEXT = 0x7FFFFFFF
)

func (x DebugReportObjectTypeEXT) String() string {
	switch x {
	case DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_OBJECT_TABLE_NVX_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_OBJECT_TABLE_NVX_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NVX_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NVX_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT"
	case DEBUG_REPORT_OBJECT_TYPE_MAX_ENUM_EXT:
		return "DEBUG_REPORT_OBJECT_TYPE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DebugReportFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugReportFlagsEXT.html
type DebugReportFlagsEXT uint32

const (
	DEBUG_REPORT_INFORMATION_BIT_EXT         DebugReportFlagsEXT = 0x00000001
	DEBUG_REPORT_WARNING_BIT_EXT             DebugReportFlagsEXT = 0x00000002
	DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT DebugReportFlagsEXT = 0x00000004
	DEBUG_REPORT_ERROR_BIT_EXT               DebugReportFlagsEXT = 0x00000008
	DEBUG_REPORT_DEBUG_BIT_EXT               DebugReportFlagsEXT = 0x00000010
	DEBUG_REPORT_FLAG_BITS_MAX_ENUM_EXT      DebugReportFlagsEXT = 0x7FFFFFFF
)

func (x DebugReportFlagsEXT) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DebugReportFlagsEXT(1 << i) {
			case DEBUG_REPORT_INFORMATION_BIT_EXT:
				s += "DEBUG_REPORT_INFORMATION_BIT_EXT|"
			case DEBUG_REPORT_WARNING_BIT_EXT:
				s += "DEBUG_REPORT_WARNING_BIT_EXT|"
			case DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT:
				s += "DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT|"
			case DEBUG_REPORT_ERROR_BIT_EXT:
				s += "DEBUG_REPORT_ERROR_BIT_EXT|"
			case DEBUG_REPORT_DEBUG_BIT_EXT:
				s += "DEBUG_REPORT_DEBUG_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

//  PfnDebugReportCallbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDebugReportCallbackEXT.html
type PfnDebugReportCallbackEXT uintptr

func (fn PfnDebugReportCallbackEXT) Call(flags DebugReportFlagsEXT, objectType DebugReportObjectTypeEXT, object uint64, location uintptr, messageCode int32, pLayerPrefix, pMessage *int8, pUserData unsafe.Pointer) Bool32 {
	ret, _, _ := call(uintptr(fn), uintptr(flags), uintptr(objectType), uintptr(object), uintptr(location), uintptr(messageCode), uintptr(unsafe.Pointer(pLayerPrefix)), uintptr(unsafe.Pointer(pMessage)), uintptr(pUserData))
	return Bool32(ret)
}
func (fn PfnDebugReportCallbackEXT) String() string { return "vkDebugReportCallbackEXT" }

// DebugReportCallbackCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html
type DebugReportCallbackCreateInfoEXT struct {
	SType       StructureType
	PNext       unsafe.Pointer
	Flags       DebugReportFlagsEXT
	PfnCallback PfnDebugReportCallbackEXT
	PUserData   unsafe.Pointer
}

func NewDebugReportCallbackCreateInfoEXT() *DebugReportCallbackCreateInfoEXT {
	p := (*DebugReportCallbackCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DebugReportCallbackCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT
	return p
}
func (p *DebugReportCallbackCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateDebugReportCallbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDebugReportCallbackEXT.html
type PfnCreateDebugReportCallbackEXT uintptr

func (fn PfnCreateDebugReportCallbackEXT) Call(instance Instance, pCreateInfo *DebugReportCallbackCreateInfoEXT, pAllocator *AllocationCallbacks, pCallback *DebugReportCallbackEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pCallback)))
	return Result(ret)
}
func (fn PfnCreateDebugReportCallbackEXT) String() string { return "vkCreateDebugReportCallbackEXT" }

//  PfnDestroyDebugReportCallbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyDebugReportCallbackEXT.html
type PfnDestroyDebugReportCallbackEXT uintptr

func (fn PfnDestroyDebugReportCallbackEXT) Call(instance Instance, callback DebugReportCallbackEXT, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(instance), uintptr(callback), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyDebugReportCallbackEXT) String() string { return "vkDestroyDebugReportCallbackEXT" }

//  PfnDebugReportMessageEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDebugReportMessageEXT.html
type PfnDebugReportMessageEXT uintptr

func (fn PfnDebugReportMessageEXT) Call(instance Instance, flags DebugReportFlagsEXT, objectType DebugReportObjectTypeEXT, object uint64, location uintptr, messageCode int32, pLayerPrefix, pMessage *int8) {
	_, _, _ = call(uintptr(fn), uintptr(instance), uintptr(flags), uintptr(objectType), uintptr(object), uintptr(location), uintptr(messageCode), uintptr(unsafe.Pointer(pLayerPrefix)), uintptr(unsafe.Pointer(pMessage)))
}
func (fn PfnDebugReportMessageEXT) String() string { return "vkDebugReportMessageEXT" }

const NV_glsl_shader = 1
const NV_GLSL_SHADER_SPEC_VERSION = 1

var NV_GLSL_SHADER_EXTENSION_NAME = "VK_NV_glsl_shader"

const EXT_depth_range_unrestricted = 1
const EXT_DEPTH_RANGE_UNRESTRICTED_SPEC_VERSION = 1

var EXT_DEPTH_RANGE_UNRESTRICTED_EXTENSION_NAME = "VK_EXT_depth_range_unrestricted"

const IMG_filter_cubic = 1
const IMG_FILTER_CUBIC_SPEC_VERSION = 1

var IMG_FILTER_CUBIC_EXTENSION_NAME = "VK_IMG_filter_cubic"

const AMD_rasterization_order = 1
const AMD_RASTERIZATION_ORDER_SPEC_VERSION = 1

var AMD_RASTERIZATION_ORDER_EXTENSION_NAME = "VK_AMD_rasterization_order"

// RasterizationOrderAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRasterizationOrderAMD.html
type RasterizationOrderAMD int32

const (
	RASTERIZATION_ORDER_STRICT_AMD      RasterizationOrderAMD = 0
	RASTERIZATION_ORDER_RELAXED_AMD     RasterizationOrderAMD = 1
	RASTERIZATION_ORDER_BEGIN_RANGE_AMD RasterizationOrderAMD = RASTERIZATION_ORDER_STRICT_AMD
	RASTERIZATION_ORDER_END_RANGE_AMD   RasterizationOrderAMD = RASTERIZATION_ORDER_RELAXED_AMD
	RASTERIZATION_ORDER_RANGE_SIZE_AMD  RasterizationOrderAMD = (RASTERIZATION_ORDER_RELAXED_AMD - RASTERIZATION_ORDER_STRICT_AMD + 1)
	RASTERIZATION_ORDER_MAX_ENUM_AMD    RasterizationOrderAMD = 0x7FFFFFFF
)

func (x RasterizationOrderAMD) String() string {
	switch x {
	case RASTERIZATION_ORDER_STRICT_AMD:
		return "RASTERIZATION_ORDER_STRICT_AMD"
	case RASTERIZATION_ORDER_RELAXED_AMD:
		return "RASTERIZATION_ORDER_RELAXED_AMD"
	case RASTERIZATION_ORDER_MAX_ENUM_AMD:
		return "RASTERIZATION_ORDER_MAX_ENUM_AMD"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PipelineRasterizationStateRasterizationOrderAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html
type PipelineRasterizationStateRasterizationOrderAMD struct {
	SType              StructureType
	PNext              unsafe.Pointer
	RasterizationOrder RasterizationOrderAMD
}

func NewPipelineRasterizationStateRasterizationOrderAMD() *PipelineRasterizationStateRasterizationOrderAMD {
	p := (*PipelineRasterizationStateRasterizationOrderAMD)(MemAlloc(unsafe.Sizeof(*(*PipelineRasterizationStateRasterizationOrderAMD)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD
	return p
}
func (p *PipelineRasterizationStateRasterizationOrderAMD) Free() { MemFree(unsafe.Pointer(p)) }

const AMD_shader_trinary_minmax = 1
const AMD_SHADER_TRINARY_MINMAX_SPEC_VERSION = 1

var AMD_SHADER_TRINARY_MINMAX_EXTENSION_NAME = "VK_AMD_shader_trinary_minmax"

const AMD_shader_explicit_vertex_parameter = 1
const AMD_SHADER_EXPLICIT_VERTEX_PARAMETER_SPEC_VERSION = 1

var AMD_SHADER_EXPLICIT_VERTEX_PARAMETER_EXTENSION_NAME = "VK_AMD_shader_explicit_vertex_parameter"

const EXT_debug_marker = 1
const EXT_DEBUG_MARKER_SPEC_VERSION = 4

var EXT_DEBUG_MARKER_EXTENSION_NAME = "VK_EXT_debug_marker"

// DebugMarkerObjectNameInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugMarkerObjectNameInfoEXT.html
type DebugMarkerObjectNameInfoEXT struct {
	SType       StructureType
	PNext       unsafe.Pointer
	ObjectType  DebugReportObjectTypeEXT
	Object      uint64
	PObjectName *int8
}

func NewDebugMarkerObjectNameInfoEXT() *DebugMarkerObjectNameInfoEXT {
	p := (*DebugMarkerObjectNameInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DebugMarkerObjectNameInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT
	return p
}
func (p *DebugMarkerObjectNameInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DebugMarkerObjectTagInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugMarkerObjectTagInfoEXT.html
type DebugMarkerObjectTagInfoEXT struct {
	SType      StructureType
	PNext      unsafe.Pointer
	ObjectType DebugReportObjectTypeEXT
	Object     uint64
	TagName    uint64
	TagSize    uintptr
	PTag       unsafe.Pointer
}

func NewDebugMarkerObjectTagInfoEXT() *DebugMarkerObjectTagInfoEXT {
	p := (*DebugMarkerObjectTagInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DebugMarkerObjectTagInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT
	return p
}
func (p *DebugMarkerObjectTagInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DebugMarkerMarkerInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugMarkerMarkerInfoEXT.html
type DebugMarkerMarkerInfoEXT struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PMarkerName *int8
	Color       [4]float32
}

func NewDebugMarkerMarkerInfoEXT() *DebugMarkerMarkerInfoEXT {
	p := (*DebugMarkerMarkerInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DebugMarkerMarkerInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT
	return p
}
func (p *DebugMarkerMarkerInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnDebugMarkerSetObjectTagEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDebugMarkerSetObjectTagEXT.html
type PfnDebugMarkerSetObjectTagEXT uintptr

func (fn PfnDebugMarkerSetObjectTagEXT) Call(device Device, pTagInfo *DebugMarkerObjectTagInfoEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pTagInfo)))
	return Result(ret)
}
func (fn PfnDebugMarkerSetObjectTagEXT) String() string { return "vkDebugMarkerSetObjectTagEXT" }

//  PfnDebugMarkerSetObjectNameEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDebugMarkerSetObjectNameEXT.html
type PfnDebugMarkerSetObjectNameEXT uintptr

func (fn PfnDebugMarkerSetObjectNameEXT) Call(device Device, pNameInfo *DebugMarkerObjectNameInfoEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pNameInfo)))
	return Result(ret)
}
func (fn PfnDebugMarkerSetObjectNameEXT) String() string { return "vkDebugMarkerSetObjectNameEXT" }

//  PfnCmdDebugMarkerBeginEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDebugMarkerBeginEXT.html
type PfnCmdDebugMarkerBeginEXT uintptr

func (fn PfnCmdDebugMarkerBeginEXT) Call(commandBuffer CommandBuffer, pMarkerInfo *DebugMarkerMarkerInfoEXT) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)))
}
func (fn PfnCmdDebugMarkerBeginEXT) String() string { return "vkCmdDebugMarkerBeginEXT" }

//  PfnCmdDebugMarkerEndEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDebugMarkerEndEXT.html
type PfnCmdDebugMarkerEndEXT uintptr

func (fn PfnCmdDebugMarkerEndEXT) Call(commandBuffer CommandBuffer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer))
}
func (fn PfnCmdDebugMarkerEndEXT) String() string { return "vkCmdDebugMarkerEndEXT" }

//  PfnCmdDebugMarkerInsertEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDebugMarkerInsertEXT.html
type PfnCmdDebugMarkerInsertEXT uintptr

func (fn PfnCmdDebugMarkerInsertEXT) Call(commandBuffer CommandBuffer, pMarkerInfo *DebugMarkerMarkerInfoEXT) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)))
}
func (fn PfnCmdDebugMarkerInsertEXT) String() string { return "vkCmdDebugMarkerInsertEXT" }

const AMD_gcn_shader = 1
const AMD_GCN_SHADER_SPEC_VERSION = 1

var AMD_GCN_SHADER_EXTENSION_NAME = "VK_AMD_gcn_shader"

const NV_dedicated_allocation = 1
const NV_DEDICATED_ALLOCATION_SPEC_VERSION = 1

var NV_DEDICATED_ALLOCATION_EXTENSION_NAME = "VK_NV_dedicated_allocation"

// DedicatedAllocationImageCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDedicatedAllocationImageCreateInfoNV.html
type DedicatedAllocationImageCreateInfoNV struct {
	SType               StructureType
	PNext               unsafe.Pointer
	DedicatedAllocation Bool32
}

func NewDedicatedAllocationImageCreateInfoNV() *DedicatedAllocationImageCreateInfoNV {
	p := (*DedicatedAllocationImageCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*DedicatedAllocationImageCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV
	return p
}
func (p *DedicatedAllocationImageCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// DedicatedAllocationBufferCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html
type DedicatedAllocationBufferCreateInfoNV struct {
	SType               StructureType
	PNext               unsafe.Pointer
	DedicatedAllocation Bool32
}

func NewDedicatedAllocationBufferCreateInfoNV() *DedicatedAllocationBufferCreateInfoNV {
	p := (*DedicatedAllocationBufferCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*DedicatedAllocationBufferCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV
	return p
}
func (p *DedicatedAllocationBufferCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// DedicatedAllocationMemoryAllocateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html
type DedicatedAllocationMemoryAllocateInfoNV struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Image  Image
	Buffer Buffer
}

func NewDedicatedAllocationMemoryAllocateInfoNV() *DedicatedAllocationMemoryAllocateInfoNV {
	p := (*DedicatedAllocationMemoryAllocateInfoNV)(MemAlloc(unsafe.Sizeof(*(*DedicatedAllocationMemoryAllocateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV
	return p
}
func (p *DedicatedAllocationMemoryAllocateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_transform_feedback = 1
const EXT_TRANSFORM_FEEDBACK_SPEC_VERSION = 1

var EXT_TRANSFORM_FEEDBACK_EXTENSION_NAME = "VK_EXT_transform_feedback"

type PipelineRasterizationStateStreamCreateFlagsEXT uint32 // reserved
// PhysicalDeviceTransformFeedbackFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceTransformFeedbackFeaturesEXT.html
type PhysicalDeviceTransformFeedbackFeaturesEXT struct {
	SType             StructureType
	PNext             unsafe.Pointer
	TransformFeedback Bool32
	GeometryStreams   Bool32
}

func NewPhysicalDeviceTransformFeedbackFeaturesEXT() *PhysicalDeviceTransformFeedbackFeaturesEXT {
	p := (*PhysicalDeviceTransformFeedbackFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceTransformFeedbackFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceTransformFeedbackFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceTransformFeedbackPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html
type PhysicalDeviceTransformFeedbackPropertiesEXT struct {
	SType                                      StructureType
	PNext                                      unsafe.Pointer
	MaxTransformFeedbackStreams                uint32
	MaxTransformFeedbackBuffers                uint32
	MaxTransformFeedbackBufferSize             DeviceSize
	MaxTransformFeedbackStreamDataSize         uint32
	MaxTransformFeedbackBufferDataSize         uint32
	MaxTransformFeedbackBufferDataStride       uint32
	TransformFeedbackQueries                   Bool32
	TransformFeedbackStreamsLinesTriangles     Bool32
	TransformFeedbackRasterizationStreamSelect Bool32
	TransformFeedbackDraw                      Bool32
}

func NewPhysicalDeviceTransformFeedbackPropertiesEXT() *PhysicalDeviceTransformFeedbackPropertiesEXT {
	p := (*PhysicalDeviceTransformFeedbackPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceTransformFeedbackPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_TRANSFORM_FEEDBACK_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceTransformFeedbackPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineRasterizationStateStreamCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html
type PipelineRasterizationStateStreamCreateInfoEXT struct {
	SType               StructureType
	PNext               unsafe.Pointer
	Flags               PipelineRasterizationStateStreamCreateFlagsEXT
	RasterizationStream uint32
}

func NewPipelineRasterizationStateStreamCreateInfoEXT() *PipelineRasterizationStateStreamCreateInfoEXT {
	p := (*PipelineRasterizationStateStreamCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineRasterizationStateStreamCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT
	return p
}
func (p *PipelineRasterizationStateStreamCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdBindTransformFeedbackBuffersEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBindTransformFeedbackBuffersEXT.html
type PfnCmdBindTransformFeedbackBuffersEXT uintptr

func (fn PfnCmdBindTransformFeedbackBuffersEXT) Call(commandBuffer CommandBuffer, firstBinding, bindingCount uint32, pBuffers *Buffer, pOffsets, pSizes *DeviceSize) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstBinding), uintptr(bindingCount), uintptr(unsafe.Pointer(pBuffers)), uintptr(unsafe.Pointer(pOffsets)), uintptr(unsafe.Pointer(pSizes)))
}
func (fn PfnCmdBindTransformFeedbackBuffersEXT) String() string {
	return "vkCmdBindTransformFeedbackBuffersEXT"
}

//  PfnCmdBeginTransformFeedbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBeginTransformFeedbackEXT.html
type PfnCmdBeginTransformFeedbackEXT uintptr

func (fn PfnCmdBeginTransformFeedbackEXT) Call(commandBuffer CommandBuffer, firstCounterBuffer, counterBufferCount uint32, pCounterBuffers *Buffer, pCounterBufferOffsets *DeviceSize) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstCounterBuffer), uintptr(counterBufferCount), uintptr(unsafe.Pointer(pCounterBuffers)), uintptr(unsafe.Pointer(pCounterBufferOffsets)))
}
func (fn PfnCmdBeginTransformFeedbackEXT) String() string { return "vkCmdBeginTransformFeedbackEXT" }

//  PfnCmdEndTransformFeedbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdEndTransformFeedbackEXT.html
type PfnCmdEndTransformFeedbackEXT uintptr

func (fn PfnCmdEndTransformFeedbackEXT) Call(commandBuffer CommandBuffer, firstCounterBuffer, counterBufferCount uint32, pCounterBuffers *Buffer, pCounterBufferOffsets *DeviceSize) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstCounterBuffer), uintptr(counterBufferCount), uintptr(unsafe.Pointer(pCounterBuffers)), uintptr(unsafe.Pointer(pCounterBufferOffsets)))
}
func (fn PfnCmdEndTransformFeedbackEXT) String() string { return "vkCmdEndTransformFeedbackEXT" }

//  PfnCmdBeginQueryIndexedEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBeginQueryIndexedEXT.html
type PfnCmdBeginQueryIndexedEXT uintptr

func (fn PfnCmdBeginQueryIndexedEXT) Call(commandBuffer CommandBuffer, queryPool QueryPool, query uint32, flags QueryControlFlags, index uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(queryPool), uintptr(query), uintptr(flags), uintptr(index))
}
func (fn PfnCmdBeginQueryIndexedEXT) String() string { return "vkCmdBeginQueryIndexedEXT" }

//  PfnCmdEndQueryIndexedEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdEndQueryIndexedEXT.html
type PfnCmdEndQueryIndexedEXT uintptr

func (fn PfnCmdEndQueryIndexedEXT) Call(commandBuffer CommandBuffer, queryPool QueryPool, query, index uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(queryPool), uintptr(query), uintptr(index))
}
func (fn PfnCmdEndQueryIndexedEXT) String() string { return "vkCmdEndQueryIndexedEXT" }

//  PfnCmdDrawIndirectByteCountEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndirectByteCountEXT.html
type PfnCmdDrawIndirectByteCountEXT uintptr

func (fn PfnCmdDrawIndirectByteCountEXT) Call(commandBuffer CommandBuffer, instanceCount, firstInstance uint32, counterBuffer Buffer, counterBufferOffset DeviceSize, counterOffset, vertexStride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(instanceCount), uintptr(firstInstance), uintptr(counterBuffer), uintptr(counterBufferOffset), uintptr(counterOffset), uintptr(vertexStride))
}
func (fn PfnCmdDrawIndirectByteCountEXT) String() string { return "vkCmdDrawIndirectByteCountEXT" }

const NVX_image_view_handle = 1
const NVX_IMAGE_VIEW_HANDLE_SPEC_VERSION = 1

var NVX_IMAGE_VIEW_HANDLE_EXTENSION_NAME = "VK_NVX_image_view_handle"

// ImageViewHandleInfoNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageViewHandleInfoNVX.html
type ImageViewHandleInfoNVX struct {
	SType          StructureType
	PNext          unsafe.Pointer
	ImageView      ImageView
	DescriptorType DescriptorType
	Sampler        Sampler
}

func NewImageViewHandleInfoNVX() *ImageViewHandleInfoNVX {
	p := (*ImageViewHandleInfoNVX)(MemAlloc(unsafe.Sizeof(*(*ImageViewHandleInfoNVX)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX
	return p
}
func (p *ImageViewHandleInfoNVX) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetImageViewHandleNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageViewHandleNVX.html
type PfnGetImageViewHandleNVX uintptr

func (fn PfnGetImageViewHandleNVX) Call(device Device, pInfo *ImageViewHandleInfoNVX) uint32 {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)))
	return uint32(ret)
}
func (fn PfnGetImageViewHandleNVX) String() string { return "vkGetImageViewHandleNVX" }

const AMD_draw_indirect_count = 1
const AMD_DRAW_INDIRECT_COUNT_SPEC_VERSION = 2

var AMD_DRAW_INDIRECT_COUNT_EXTENSION_NAME = "VK_AMD_draw_indirect_count"

//  PfnCmdDrawIndirectCountAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndirectCountAMD.html
type PfnCmdDrawIndirectCountAMD uintptr

func (fn PfnCmdDrawIndirectCountAMD) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride))
}
func (fn PfnCmdDrawIndirectCountAMD) String() string { return "vkCmdDrawIndirectCountAMD" }

//  PfnCmdDrawIndexedIndirectCountAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawIndexedIndirectCountAMD.html
type PfnCmdDrawIndexedIndirectCountAMD uintptr

func (fn PfnCmdDrawIndexedIndirectCountAMD) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride))
}
func (fn PfnCmdDrawIndexedIndirectCountAMD) String() string {
	return "vkCmdDrawIndexedIndirectCountAMD"
}

const AMD_negative_viewport_height = 1
const AMD_NEGATIVE_VIEWPORT_HEIGHT_SPEC_VERSION = 1

var AMD_NEGATIVE_VIEWPORT_HEIGHT_EXTENSION_NAME = "VK_AMD_negative_viewport_height"

const AMD_gpu_shader_half_float = 1
const AMD_GPU_SHADER_HALF_FLOAT_SPEC_VERSION = 2

var AMD_GPU_SHADER_HALF_FLOAT_EXTENSION_NAME = "VK_AMD_gpu_shader_half_float"

const AMD_shader_ballot = 1
const AMD_SHADER_BALLOT_SPEC_VERSION = 1

var AMD_SHADER_BALLOT_EXTENSION_NAME = "VK_AMD_shader_ballot"

const AMD_texture_gather_bias_lod = 1
const AMD_TEXTURE_GATHER_BIAS_LOD_SPEC_VERSION = 1

var AMD_TEXTURE_GATHER_BIAS_LOD_EXTENSION_NAME = "VK_AMD_texture_gather_bias_lod"

// TextureLODGatherFormatPropertiesAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkTextureLODGatherFormatPropertiesAMD.html
type TextureLODGatherFormatPropertiesAMD struct {
	SType                           StructureType
	PNext                           unsafe.Pointer
	SupportsTextureGatherLODBiasAMD Bool32
}

func NewTextureLODGatherFormatPropertiesAMD() *TextureLODGatherFormatPropertiesAMD {
	p := (*TextureLODGatherFormatPropertiesAMD)(MemAlloc(unsafe.Sizeof(*(*TextureLODGatherFormatPropertiesAMD)(nil))))
	p.SType = STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD
	return p
}
func (p *TextureLODGatherFormatPropertiesAMD) Free() { MemFree(unsafe.Pointer(p)) }

const AMD_shader_info = 1
const AMD_SHADER_INFO_SPEC_VERSION = 1

var AMD_SHADER_INFO_EXTENSION_NAME = "VK_AMD_shader_info"

// ShaderInfoTypeAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderInfoTypeAMD.html
type ShaderInfoTypeAMD int32

const (
	SHADER_INFO_TYPE_STATISTICS_AMD  ShaderInfoTypeAMD = 0
	SHADER_INFO_TYPE_BINARY_AMD      ShaderInfoTypeAMD = 1
	SHADER_INFO_TYPE_DISASSEMBLY_AMD ShaderInfoTypeAMD = 2
	SHADER_INFO_TYPE_BEGIN_RANGE_AMD ShaderInfoTypeAMD = SHADER_INFO_TYPE_STATISTICS_AMD
	SHADER_INFO_TYPE_END_RANGE_AMD   ShaderInfoTypeAMD = SHADER_INFO_TYPE_DISASSEMBLY_AMD
	SHADER_INFO_TYPE_RANGE_SIZE_AMD  ShaderInfoTypeAMD = (SHADER_INFO_TYPE_DISASSEMBLY_AMD - SHADER_INFO_TYPE_STATISTICS_AMD + 1)
	SHADER_INFO_TYPE_MAX_ENUM_AMD    ShaderInfoTypeAMD = 0x7FFFFFFF
)

func (x ShaderInfoTypeAMD) String() string {
	switch x {
	case SHADER_INFO_TYPE_STATISTICS_AMD:
		return "SHADER_INFO_TYPE_STATISTICS_AMD"
	case SHADER_INFO_TYPE_BINARY_AMD:
		return "SHADER_INFO_TYPE_BINARY_AMD"
	case SHADER_INFO_TYPE_DISASSEMBLY_AMD:
		return "SHADER_INFO_TYPE_DISASSEMBLY_AMD"
	case SHADER_INFO_TYPE_MAX_ENUM_AMD:
		return "SHADER_INFO_TYPE_MAX_ENUM_AMD"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ShaderResourceUsageAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderResourceUsageAMD.html
type ShaderResourceUsageAMD struct {
	NumUsedVgprs             uint32
	NumUsedSgprs             uint32
	LdsSizePerLocalWorkGroup uint32
	LdsUsageSizeInBytes      uintptr
	ScratchMemUsageInBytes   uintptr
}

func NewShaderResourceUsageAMD() *ShaderResourceUsageAMD {
	return (*ShaderResourceUsageAMD)(MemAlloc(unsafe.Sizeof(*(*ShaderResourceUsageAMD)(nil))))
}
func (p *ShaderResourceUsageAMD) Free() { MemFree(unsafe.Pointer(p)) }

// ShaderStatisticsInfoAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderStatisticsInfoAMD.html
type ShaderStatisticsInfoAMD struct {
	ShaderStageMask      ShaderStageFlags
	ResourceUsage        ShaderResourceUsageAMD
	NumPhysicalVgprs     uint32
	NumPhysicalSgprs     uint32
	NumAvailableVgprs    uint32
	NumAvailableSgprs    uint32
	ComputeWorkGroupSize [3]uint32
}

func NewShaderStatisticsInfoAMD() *ShaderStatisticsInfoAMD {
	return (*ShaderStatisticsInfoAMD)(MemAlloc(unsafe.Sizeof(*(*ShaderStatisticsInfoAMD)(nil))))
}
func (p *ShaderStatisticsInfoAMD) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetShaderInfoAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetShaderInfoAMD.html
type PfnGetShaderInfoAMD uintptr

func (fn PfnGetShaderInfoAMD) Call(device Device, pipeline Pipeline, shaderStage ShaderStageFlags, infoType ShaderInfoTypeAMD, pInfoSize *uintptr, pInfo unsafe.Pointer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(pipeline), uintptr(shaderStage), uintptr(infoType), uintptr(unsafe.Pointer(pInfoSize)), uintptr(pInfo))
	return Result(ret)
}
func (fn PfnGetShaderInfoAMD) String() string { return "vkGetShaderInfoAMD" }

const AMD_shader_image_load_store_lod = 1
const AMD_SHADER_IMAGE_LOAD_STORE_LOD_SPEC_VERSION = 1

var AMD_SHADER_IMAGE_LOAD_STORE_LOD_EXTENSION_NAME = "VK_AMD_shader_image_load_store_lod"

const NV_corner_sampled_image = 1
const NV_CORNER_SAMPLED_IMAGE_SPEC_VERSION = 2

var NV_CORNER_SAMPLED_IMAGE_EXTENSION_NAME = "VK_NV_corner_sampled_image"

// PhysicalDeviceCornerSampledImageFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceCornerSampledImageFeaturesNV.html
type PhysicalDeviceCornerSampledImageFeaturesNV struct {
	SType              StructureType
	PNext              unsafe.Pointer
	CornerSampledImage Bool32
}

func NewPhysicalDeviceCornerSampledImageFeaturesNV() *PhysicalDeviceCornerSampledImageFeaturesNV {
	p := (*PhysicalDeviceCornerSampledImageFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceCornerSampledImageFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV
	return p
}
func (p *PhysicalDeviceCornerSampledImageFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

const IMG_format_pvrtc = 1
const IMG_FORMAT_PVRTC_SPEC_VERSION = 1

var IMG_FORMAT_PVRTC_EXTENSION_NAME = "VK_IMG_format_pvrtc"

const NV_external_memory_capabilities = 1
const NV_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION = 1

var NV_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME = "VK_NV_external_memory_capabilities"

// ExternalMemoryHandleTypeFlagsNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html
type ExternalMemoryHandleTypeFlagsNV uint32

const (
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV     ExternalMemoryHandleTypeFlagsNV = 0x00000001
	EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_NV ExternalMemoryHandleTypeFlagsNV = 0x00000002
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_BIT_NV      ExternalMemoryHandleTypeFlagsNV = 0x00000004
	EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_KMT_BIT_NV  ExternalMemoryHandleTypeFlagsNV = 0x00000008
	EXTERNAL_MEMORY_HANDLE_TYPE_FLAG_BITS_MAX_ENUM_NV   ExternalMemoryHandleTypeFlagsNV = 0x7FFFFFFF
)

func (x ExternalMemoryHandleTypeFlagsNV) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalMemoryHandleTypeFlagsNV(1 << i) {
			case EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_NV:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_NV|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_BIT_NV:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_BIT_NV|"
			case EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_KMT_BIT_NV:
				s += "EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_KMT_BIT_NV|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ExternalMemoryFeatureFlagsNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryFeatureFlagsNV.html
type ExternalMemoryFeatureFlagsNV uint32

const (
	EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV ExternalMemoryFeatureFlagsNV = 0x00000001
	EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV     ExternalMemoryFeatureFlagsNV = 0x00000002
	EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV     ExternalMemoryFeatureFlagsNV = 0x00000004
	EXTERNAL_MEMORY_FEATURE_FLAG_BITS_MAX_ENUM_NV ExternalMemoryFeatureFlagsNV = 0x7FFFFFFF
)

func (x ExternalMemoryFeatureFlagsNV) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ExternalMemoryFeatureFlagsNV(1 << i) {
			case EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV:
				s += "EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV|"
			case EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV:
				s += "EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV|"
			case EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV:
				s += "EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ExternalImageFormatPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalImageFormatPropertiesNV.html
type ExternalImageFormatPropertiesNV struct {
	ImageFormatProperties         ImageFormatProperties
	ExternalMemoryFeatures        ExternalMemoryFeatureFlagsNV
	ExportFromImportedHandleTypes ExternalMemoryHandleTypeFlagsNV
	CompatibleHandleTypes         ExternalMemoryHandleTypeFlagsNV
}

func NewExternalImageFormatPropertiesNV() *ExternalImageFormatPropertiesNV {
	return (*ExternalImageFormatPropertiesNV)(MemAlloc(unsafe.Sizeof(*(*ExternalImageFormatPropertiesNV)(nil))))
}
func (p *ExternalImageFormatPropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceExternalImageFormatPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html
type PfnGetPhysicalDeviceExternalImageFormatPropertiesNV uintptr

func (fn PfnGetPhysicalDeviceExternalImageFormatPropertiesNV) Call(physicalDevice PhysicalDevice, format Format, type_ ImageType, tiling ImageTiling, usage ImageUsageFlags, flags ImageCreateFlags, externalHandleType ExternalMemoryHandleTypeFlagsNV, pExternalImageFormatProperties *ExternalImageFormatPropertiesNV) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(format), uintptr(type_), uintptr(tiling), uintptr(usage), uintptr(flags), uintptr(externalHandleType), uintptr(unsafe.Pointer(pExternalImageFormatProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceExternalImageFormatPropertiesNV) String() string {
	return "vkGetPhysicalDeviceExternalImageFormatPropertiesNV"
}

const NV_external_memory = 1
const NV_EXTERNAL_MEMORY_SPEC_VERSION = 1

var NV_EXTERNAL_MEMORY_EXTENSION_NAME = "VK_NV_external_memory"

// ExternalMemoryImageCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExternalMemoryImageCreateInfoNV.html
type ExternalMemoryImageCreateInfoNV struct {
	SType       StructureType
	PNext       unsafe.Pointer
	HandleTypes ExternalMemoryHandleTypeFlagsNV
}

func NewExternalMemoryImageCreateInfoNV() *ExternalMemoryImageCreateInfoNV {
	p := (*ExternalMemoryImageCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*ExternalMemoryImageCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV
	return p
}
func (p *ExternalMemoryImageCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// ExportMemoryAllocateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportMemoryAllocateInfoNV.html
type ExportMemoryAllocateInfoNV struct {
	SType       StructureType
	PNext       unsafe.Pointer
	HandleTypes ExternalMemoryHandleTypeFlagsNV
}

func NewExportMemoryAllocateInfoNV() *ExportMemoryAllocateInfoNV {
	p := (*ExportMemoryAllocateInfoNV)(MemAlloc(unsafe.Sizeof(*(*ExportMemoryAllocateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV
	return p
}
func (p *ExportMemoryAllocateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_validation_flags = 1
const EXT_VALIDATION_FLAGS_SPEC_VERSION = 2

var EXT_VALIDATION_FLAGS_EXTENSION_NAME = "VK_EXT_validation_flags"

// ValidationCheckEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationCheckEXT.html
type ValidationCheckEXT int32

const (
	VALIDATION_CHECK_ALL_EXT         ValidationCheckEXT = 0
	VALIDATION_CHECK_SHADERS_EXT     ValidationCheckEXT = 1
	VALIDATION_CHECK_BEGIN_RANGE_EXT ValidationCheckEXT = VALIDATION_CHECK_ALL_EXT
	VALIDATION_CHECK_END_RANGE_EXT   ValidationCheckEXT = VALIDATION_CHECK_SHADERS_EXT
	VALIDATION_CHECK_RANGE_SIZE_EXT  ValidationCheckEXT = (VALIDATION_CHECK_SHADERS_EXT - VALIDATION_CHECK_ALL_EXT + 1)
	VALIDATION_CHECK_MAX_ENUM_EXT    ValidationCheckEXT = 0x7FFFFFFF
)

func (x ValidationCheckEXT) String() string {
	switch x {
	case VALIDATION_CHECK_ALL_EXT:
		return "VALIDATION_CHECK_ALL_EXT"
	case VALIDATION_CHECK_SHADERS_EXT:
		return "VALIDATION_CHECK_SHADERS_EXT"
	case VALIDATION_CHECK_MAX_ENUM_EXT:
		return "VALIDATION_CHECK_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ValidationFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationFlagsEXT.html
type ValidationFlagsEXT struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	DisabledValidationCheckCount uint32
	PDisabledValidationChecks    *ValidationCheckEXT
}

func NewValidationFlagsEXT() *ValidationFlagsEXT {
	p := (*ValidationFlagsEXT)(MemAlloc(unsafe.Sizeof(*(*ValidationFlagsEXT)(nil))))
	p.SType = STRUCTURE_TYPE_VALIDATION_FLAGS_EXT
	return p
}
func (p *ValidationFlagsEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_shader_subgroup_ballot = 1
const EXT_SHADER_SUBGROUP_BALLOT_SPEC_VERSION = 1

var EXT_SHADER_SUBGROUP_BALLOT_EXTENSION_NAME = "VK_EXT_shader_subgroup_ballot"

const EXT_shader_subgroup_vote = 1
const EXT_SHADER_SUBGROUP_VOTE_SPEC_VERSION = 1

var EXT_SHADER_SUBGROUP_VOTE_EXTENSION_NAME = "VK_EXT_shader_subgroup_vote"

const EXT_texture_compression_astc_hdr = 1
const EXT_TEXTURE_COMPRESSION_ASTC_HDR_SPEC_VERSION = 1

var EXT_TEXTURE_COMPRESSION_ASTC_HDR_EXTENSION_NAME = "VK_EXT_texture_compression_astc_hdr"

// PhysicalDeviceTextureCompressionASTCHDRFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT.html
type PhysicalDeviceTextureCompressionASTCHDRFeaturesEXT struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	TextureCompressionASTC_HDR Bool32
}

func NewPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT() *PhysicalDeviceTextureCompressionASTCHDRFeaturesEXT {
	return (*PhysicalDeviceTextureCompressionASTCHDRFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceTextureCompressionASTCHDRFeaturesEXT)(nil))))
}
func (p *PhysicalDeviceTextureCompressionASTCHDRFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_astc_decode_mode = 1
const EXT_ASTC_DECODE_MODE_SPEC_VERSION = 1

var EXT_ASTC_DECODE_MODE_EXTENSION_NAME = "VK_EXT_astc_decode_mode"

// ImageViewASTCDecodeModeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageViewASTCDecodeModeEXT.html
type ImageViewASTCDecodeModeEXT struct {
	SType      StructureType
	PNext      unsafe.Pointer
	DecodeMode Format
}

func NewImageViewASTCDecodeModeEXT() *ImageViewASTCDecodeModeEXT {
	p := (*ImageViewASTCDecodeModeEXT)(MemAlloc(unsafe.Sizeof(*(*ImageViewASTCDecodeModeEXT)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT
	return p
}
func (p *ImageViewASTCDecodeModeEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceASTCDecodeFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceASTCDecodeFeaturesEXT.html
type PhysicalDeviceASTCDecodeFeaturesEXT struct {
	SType                    StructureType
	PNext                    unsafe.Pointer
	DecodeModeSharedExponent Bool32
}

func NewPhysicalDeviceASTCDecodeFeaturesEXT() *PhysicalDeviceASTCDecodeFeaturesEXT {
	p := (*PhysicalDeviceASTCDecodeFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceASTCDecodeFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceASTCDecodeFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_conditional_rendering = 1
const EXT_CONDITIONAL_RENDERING_SPEC_VERSION = 2

var EXT_CONDITIONAL_RENDERING_EXTENSION_NAME = "VK_EXT_conditional_rendering"

// ConditionalRenderingFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkConditionalRenderingFlagsEXT.html
type ConditionalRenderingFlagsEXT uint32

const (
	CONDITIONAL_RENDERING_INVERTED_BIT_EXT       ConditionalRenderingFlagsEXT = 0x00000001
	CONDITIONAL_RENDERING_FLAG_BITS_MAX_ENUM_EXT ConditionalRenderingFlagsEXT = 0x7FFFFFFF
)

func (x ConditionalRenderingFlagsEXT) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ConditionalRenderingFlagsEXT(1 << i) {
			case CONDITIONAL_RENDERING_INVERTED_BIT_EXT:
				s += "CONDITIONAL_RENDERING_INVERTED_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ConditionalRenderingBeginInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkConditionalRenderingBeginInfoEXT.html
type ConditionalRenderingBeginInfoEXT struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Buffer Buffer
	Offset DeviceSize
	Flags  ConditionalRenderingFlagsEXT
}

func NewConditionalRenderingBeginInfoEXT() *ConditionalRenderingBeginInfoEXT {
	p := (*ConditionalRenderingBeginInfoEXT)(MemAlloc(unsafe.Sizeof(*(*ConditionalRenderingBeginInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT
	return p
}
func (p *ConditionalRenderingBeginInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceConditionalRenderingFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceConditionalRenderingFeaturesEXT.html
type PhysicalDeviceConditionalRenderingFeaturesEXT struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	ConditionalRendering          Bool32
	InheritedConditionalRendering Bool32
}

func NewPhysicalDeviceConditionalRenderingFeaturesEXT() *PhysicalDeviceConditionalRenderingFeaturesEXT {
	p := (*PhysicalDeviceConditionalRenderingFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceConditionalRenderingFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceConditionalRenderingFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// CommandBufferInheritanceConditionalRenderingInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCommandBufferInheritanceConditionalRenderingInfoEXT.html
type CommandBufferInheritanceConditionalRenderingInfoEXT struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	ConditionalRenderingEnable Bool32
}

func NewCommandBufferInheritanceConditionalRenderingInfoEXT() *CommandBufferInheritanceConditionalRenderingInfoEXT {
	p := (*CommandBufferInheritanceConditionalRenderingInfoEXT)(MemAlloc(unsafe.Sizeof(*(*CommandBufferInheritanceConditionalRenderingInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT
	return p
}
func (p *CommandBufferInheritanceConditionalRenderingInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdBeginConditionalRenderingEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBeginConditionalRenderingEXT.html
type PfnCmdBeginConditionalRenderingEXT uintptr

func (fn PfnCmdBeginConditionalRenderingEXT) Call(commandBuffer CommandBuffer, pConditionalRenderingBegin *ConditionalRenderingBeginInfoEXT) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pConditionalRenderingBegin)))
}
func (fn PfnCmdBeginConditionalRenderingEXT) String() string {
	return "vkCmdBeginConditionalRenderingEXT"
}

//  PfnCmdEndConditionalRenderingEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdEndConditionalRenderingEXT.html
type PfnCmdEndConditionalRenderingEXT uintptr

func (fn PfnCmdEndConditionalRenderingEXT) Call(commandBuffer CommandBuffer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer))
}
func (fn PfnCmdEndConditionalRenderingEXT) String() string { return "vkCmdEndConditionalRenderingEXT" }

const NVX_device_generated_commands = 1

// ObjectTableNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTableNVX.html
type ObjectTableNVX NonDispatchableHandle

// IndirectCommandsLayoutNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkIndirectCommandsLayoutNVX.html
type IndirectCommandsLayoutNVX NonDispatchableHandle

const NVX_DEVICE_GENERATED_COMMANDS_SPEC_VERSION = 3

var NVX_DEVICE_GENERATED_COMMANDS_EXTENSION_NAME = "VK_NVX_device_generated_commands"

// IndirectCommandsTokenTypeNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkIndirectCommandsTokenTypeNVX.html
type IndirectCommandsTokenTypeNVX int32

const (
	INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NVX       IndirectCommandsTokenTypeNVX = 0
	INDIRECT_COMMANDS_TOKEN_TYPE_DESCRIPTOR_SET_NVX IndirectCommandsTokenTypeNVX = 1
	INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NVX   IndirectCommandsTokenTypeNVX = 2
	INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NVX  IndirectCommandsTokenTypeNVX = 3
	INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NVX  IndirectCommandsTokenTypeNVX = 4
	INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NVX   IndirectCommandsTokenTypeNVX = 5
	INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NVX           IndirectCommandsTokenTypeNVX = 6
	INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NVX       IndirectCommandsTokenTypeNVX = 7
	INDIRECT_COMMANDS_TOKEN_TYPE_BEGIN_RANGE_NVX    IndirectCommandsTokenTypeNVX = INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NVX
	INDIRECT_COMMANDS_TOKEN_TYPE_END_RANGE_NVX      IndirectCommandsTokenTypeNVX = INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NVX
	INDIRECT_COMMANDS_TOKEN_TYPE_RANGE_SIZE_NVX     IndirectCommandsTokenTypeNVX = (INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NVX - INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NVX + 1)
	INDIRECT_COMMANDS_TOKEN_TYPE_MAX_ENUM_NVX       IndirectCommandsTokenTypeNVX = 0x7FFFFFFF
)

func (x IndirectCommandsTokenTypeNVX) String() string {
	switch x {
	case INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_DESCRIPTOR_SET_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_DESCRIPTOR_SET_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_PUSH_CONSTANT_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_INDEXED_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NVX"
	case INDIRECT_COMMANDS_TOKEN_TYPE_MAX_ENUM_NVX:
		return "INDIRECT_COMMANDS_TOKEN_TYPE_MAX_ENUM_NVX"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ObjectEntryTypeNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectEntryTypeNVX.html
type ObjectEntryTypeNVX int32

const (
	OBJECT_ENTRY_TYPE_DESCRIPTOR_SET_NVX ObjectEntryTypeNVX = 0
	OBJECT_ENTRY_TYPE_PIPELINE_NVX       ObjectEntryTypeNVX = 1
	OBJECT_ENTRY_TYPE_INDEX_BUFFER_NVX   ObjectEntryTypeNVX = 2
	OBJECT_ENTRY_TYPE_VERTEX_BUFFER_NVX  ObjectEntryTypeNVX = 3
	OBJECT_ENTRY_TYPE_PUSH_CONSTANT_NVX  ObjectEntryTypeNVX = 4
	OBJECT_ENTRY_TYPE_BEGIN_RANGE_NVX    ObjectEntryTypeNVX = OBJECT_ENTRY_TYPE_DESCRIPTOR_SET_NVX
	OBJECT_ENTRY_TYPE_END_RANGE_NVX      ObjectEntryTypeNVX = OBJECT_ENTRY_TYPE_PUSH_CONSTANT_NVX
	OBJECT_ENTRY_TYPE_RANGE_SIZE_NVX     ObjectEntryTypeNVX = (OBJECT_ENTRY_TYPE_PUSH_CONSTANT_NVX - OBJECT_ENTRY_TYPE_DESCRIPTOR_SET_NVX + 1)
	OBJECT_ENTRY_TYPE_MAX_ENUM_NVX       ObjectEntryTypeNVX = 0x7FFFFFFF
)

func (x ObjectEntryTypeNVX) String() string {
	switch x {
	case OBJECT_ENTRY_TYPE_DESCRIPTOR_SET_NVX:
		return "OBJECT_ENTRY_TYPE_DESCRIPTOR_SET_NVX"
	case OBJECT_ENTRY_TYPE_PIPELINE_NVX:
		return "OBJECT_ENTRY_TYPE_PIPELINE_NVX"
	case OBJECT_ENTRY_TYPE_INDEX_BUFFER_NVX:
		return "OBJECT_ENTRY_TYPE_INDEX_BUFFER_NVX"
	case OBJECT_ENTRY_TYPE_VERTEX_BUFFER_NVX:
		return "OBJECT_ENTRY_TYPE_VERTEX_BUFFER_NVX"
	case OBJECT_ENTRY_TYPE_PUSH_CONSTANT_NVX:
		return "OBJECT_ENTRY_TYPE_PUSH_CONSTANT_NVX"
	case OBJECT_ENTRY_TYPE_MAX_ENUM_NVX:
		return "OBJECT_ENTRY_TYPE_MAX_ENUM_NVX"
	default:
		return fmt.Sprint(int32(x))
	}
}

// IndirectCommandsLayoutUsageFlagsNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkIndirectCommandsLayoutUsageFlagsNVX.html
type IndirectCommandsLayoutUsageFlagsNVX uint32

const (
	INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NVX IndirectCommandsLayoutUsageFlagsNVX = 0x00000001
	INDIRECT_COMMANDS_LAYOUT_USAGE_SPARSE_SEQUENCES_BIT_NVX    IndirectCommandsLayoutUsageFlagsNVX = 0x00000002
	INDIRECT_COMMANDS_LAYOUT_USAGE_EMPTY_EXECUTIONS_BIT_NVX    IndirectCommandsLayoutUsageFlagsNVX = 0x00000004
	INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NVX   IndirectCommandsLayoutUsageFlagsNVX = 0x00000008
	INDIRECT_COMMANDS_LAYOUT_USAGE_FLAG_BITS_MAX_ENUM_NVX      IndirectCommandsLayoutUsageFlagsNVX = 0x7FFFFFFF
)

func (x IndirectCommandsLayoutUsageFlagsNVX) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch IndirectCommandsLayoutUsageFlagsNVX(1 << i) {
			case INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NVX:
				s += "INDIRECT_COMMANDS_LAYOUT_USAGE_UNORDERED_SEQUENCES_BIT_NVX|"
			case INDIRECT_COMMANDS_LAYOUT_USAGE_SPARSE_SEQUENCES_BIT_NVX:
				s += "INDIRECT_COMMANDS_LAYOUT_USAGE_SPARSE_SEQUENCES_BIT_NVX|"
			case INDIRECT_COMMANDS_LAYOUT_USAGE_EMPTY_EXECUTIONS_BIT_NVX:
				s += "INDIRECT_COMMANDS_LAYOUT_USAGE_EMPTY_EXECUTIONS_BIT_NVX|"
			case INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NVX:
				s += "INDIRECT_COMMANDS_LAYOUT_USAGE_INDEXED_SEQUENCES_BIT_NVX|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// ObjectEntryUsageFlagsNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectEntryUsageFlagsNVX.html
type ObjectEntryUsageFlagsNVX uint32

const (
	OBJECT_ENTRY_USAGE_GRAPHICS_BIT_NVX       ObjectEntryUsageFlagsNVX = 0x00000001
	OBJECT_ENTRY_USAGE_COMPUTE_BIT_NVX        ObjectEntryUsageFlagsNVX = 0x00000002
	OBJECT_ENTRY_USAGE_FLAG_BITS_MAX_ENUM_NVX ObjectEntryUsageFlagsNVX = 0x7FFFFFFF
)

func (x ObjectEntryUsageFlagsNVX) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ObjectEntryUsageFlagsNVX(1 << i) {
			case OBJECT_ENTRY_USAGE_GRAPHICS_BIT_NVX:
				s += "OBJECT_ENTRY_USAGE_GRAPHICS_BIT_NVX|"
			case OBJECT_ENTRY_USAGE_COMPUTE_BIT_NVX:
				s += "OBJECT_ENTRY_USAGE_COMPUTE_BIT_NVX|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DeviceGeneratedCommandsFeaturesNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGeneratedCommandsFeaturesNVX.html
type DeviceGeneratedCommandsFeaturesNVX struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	ComputeBindingPointSupport Bool32
}

func NewDeviceGeneratedCommandsFeaturesNVX() *DeviceGeneratedCommandsFeaturesNVX {
	p := (*DeviceGeneratedCommandsFeaturesNVX)(MemAlloc(unsafe.Sizeof(*(*DeviceGeneratedCommandsFeaturesNVX)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_FEATURES_NVX
	return p
}
func (p *DeviceGeneratedCommandsFeaturesNVX) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceGeneratedCommandsLimitsNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceGeneratedCommandsLimitsNVX.html
type DeviceGeneratedCommandsLimitsNVX struct {
	SType                                 StructureType
	PNext                                 unsafe.Pointer
	MaxIndirectCommandsLayoutTokenCount   uint32
	MaxObjectEntryCounts                  uint32
	MinSequenceCountBufferOffsetAlignment uint32
	MinSequenceIndexBufferOffsetAlignment uint32
	MinCommandsTokenBufferOffsetAlignment uint32
}

func NewDeviceGeneratedCommandsLimitsNVX() *DeviceGeneratedCommandsLimitsNVX {
	p := (*DeviceGeneratedCommandsLimitsNVX)(MemAlloc(unsafe.Sizeof(*(*DeviceGeneratedCommandsLimitsNVX)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_GENERATED_COMMANDS_LIMITS_NVX
	return p
}
func (p *DeviceGeneratedCommandsLimitsNVX) Free() { MemFree(unsafe.Pointer(p)) }

// IndirectCommandsTokenNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkIndirectCommandsTokenNVX.html
type IndirectCommandsTokenNVX struct {
	TokenType IndirectCommandsTokenTypeNVX
	Buffer    Buffer
	Offset    DeviceSize
}

func NewIndirectCommandsTokenNVX() *IndirectCommandsTokenNVX {
	return (*IndirectCommandsTokenNVX)(MemAlloc(unsafe.Sizeof(*(*IndirectCommandsTokenNVX)(nil))))
}
func (p *IndirectCommandsTokenNVX) Free() { MemFree(unsafe.Pointer(p)) }

// IndirectCommandsLayoutTokenNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkIndirectCommandsLayoutTokenNVX.html
type IndirectCommandsLayoutTokenNVX struct {
	TokenType    IndirectCommandsTokenTypeNVX
	BindingUnit  uint32
	DynamicCount uint32
	Divisor      uint32
}

func NewIndirectCommandsLayoutTokenNVX() *IndirectCommandsLayoutTokenNVX {
	return (*IndirectCommandsLayoutTokenNVX)(MemAlloc(unsafe.Sizeof(*(*IndirectCommandsLayoutTokenNVX)(nil))))
}
func (p *IndirectCommandsLayoutTokenNVX) Free() { MemFree(unsafe.Pointer(p)) }

// IndirectCommandsLayoutCreateInfoNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNVX.html
type IndirectCommandsLayoutCreateInfoNVX struct {
	SType             StructureType
	PNext             unsafe.Pointer
	PipelineBindPoint PipelineBindPoint
	Flags             IndirectCommandsLayoutUsageFlagsNVX
	TokenCount        uint32
	PTokens           *IndirectCommandsLayoutTokenNVX
}

func NewIndirectCommandsLayoutCreateInfoNVX() *IndirectCommandsLayoutCreateInfoNVX {
	p := (*IndirectCommandsLayoutCreateInfoNVX)(MemAlloc(unsafe.Sizeof(*(*IndirectCommandsLayoutCreateInfoNVX)(nil))))
	p.SType = STRUCTURE_TYPE_INDIRECT_COMMANDS_LAYOUT_CREATE_INFO_NVX
	return p
}
func (p *IndirectCommandsLayoutCreateInfoNVX) Free() { MemFree(unsafe.Pointer(p)) }

// CmdProcessCommandsInfoNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCmdProcessCommandsInfoNVX.html
type CmdProcessCommandsInfoNVX struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	ObjectTable                ObjectTableNVX
	IndirectCommandsLayout     IndirectCommandsLayoutNVX
	IndirectCommandsTokenCount uint32
	PIndirectCommandsTokens    *IndirectCommandsTokenNVX
	MaxSequencesCount          uint32
	TargetCommandBuffer        CommandBuffer
	SequencesCountBuffer       Buffer
	SequencesCountOffset       DeviceSize
	SequencesIndexBuffer       Buffer
	SequencesIndexOffset       DeviceSize
}

func NewCmdProcessCommandsInfoNVX() *CmdProcessCommandsInfoNVX {
	p := (*CmdProcessCommandsInfoNVX)(MemAlloc(unsafe.Sizeof(*(*CmdProcessCommandsInfoNVX)(nil))))
	p.SType = STRUCTURE_TYPE_CMD_PROCESS_COMMANDS_INFO_NVX
	return p
}
func (p *CmdProcessCommandsInfoNVX) Free() { MemFree(unsafe.Pointer(p)) }

// CmdReserveSpaceForCommandsInfoNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCmdReserveSpaceForCommandsInfoNVX.html
type CmdReserveSpaceForCommandsInfoNVX struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	ObjectTable            ObjectTableNVX
	IndirectCommandsLayout IndirectCommandsLayoutNVX
	MaxSequencesCount      uint32
}

func NewCmdReserveSpaceForCommandsInfoNVX() *CmdReserveSpaceForCommandsInfoNVX {
	p := (*CmdReserveSpaceForCommandsInfoNVX)(MemAlloc(unsafe.Sizeof(*(*CmdReserveSpaceForCommandsInfoNVX)(nil))))
	p.SType = STRUCTURE_TYPE_CMD_RESERVE_SPACE_FOR_COMMANDS_INFO_NVX
	return p
}
func (p *CmdReserveSpaceForCommandsInfoNVX) Free() { MemFree(unsafe.Pointer(p)) }

// ObjectTableCreateInfoNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTableCreateInfoNVX.html
type ObjectTableCreateInfoNVX struct {
	SType                          StructureType
	PNext                          unsafe.Pointer
	ObjectCount                    uint32
	PObjectEntryTypes              *ObjectEntryTypeNVX
	PObjectEntryCounts             *uint32
	PObjectEntryUsageFlags         *ObjectEntryUsageFlagsNVX
	MaxUniformBuffersPerDescriptor uint32
	MaxStorageBuffersPerDescriptor uint32
	MaxStorageImagesPerDescriptor  uint32
	MaxSampledImagesPerDescriptor  uint32
	MaxPipelineLayouts             uint32
}

func NewObjectTableCreateInfoNVX() *ObjectTableCreateInfoNVX {
	p := (*ObjectTableCreateInfoNVX)(MemAlloc(unsafe.Sizeof(*(*ObjectTableCreateInfoNVX)(nil))))
	p.SType = STRUCTURE_TYPE_OBJECT_TABLE_CREATE_INFO_NVX
	return p
}
func (p *ObjectTableCreateInfoNVX) Free() { MemFree(unsafe.Pointer(p)) }

// ObjectTableEntryNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTableEntryNVX.html
type ObjectTableEntryNVX struct {
	Type  ObjectEntryTypeNVX
	Flags ObjectEntryUsageFlagsNVX
}

func NewObjectTableEntryNVX() *ObjectTableEntryNVX {
	return (*ObjectTableEntryNVX)(MemAlloc(unsafe.Sizeof(*(*ObjectTableEntryNVX)(nil))))
}
func (p *ObjectTableEntryNVX) Free() { MemFree(unsafe.Pointer(p)) }

// ObjectTablePipelineEntryNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTablePipelineEntryNVX.html
type ObjectTablePipelineEntryNVX struct {
	Type     ObjectEntryTypeNVX
	Flags    ObjectEntryUsageFlagsNVX
	Pipeline Pipeline
}

func NewObjectTablePipelineEntryNVX() *ObjectTablePipelineEntryNVX {
	return (*ObjectTablePipelineEntryNVX)(MemAlloc(unsafe.Sizeof(*(*ObjectTablePipelineEntryNVX)(nil))))
}
func (p *ObjectTablePipelineEntryNVX) Free() { MemFree(unsafe.Pointer(p)) }

// ObjectTableDescriptorSetEntryNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTableDescriptorSetEntryNVX.html
type ObjectTableDescriptorSetEntryNVX struct {
	Type           ObjectEntryTypeNVX
	Flags          ObjectEntryUsageFlagsNVX
	PipelineLayout PipelineLayout
	DescriptorSet  DescriptorSet
}

func NewObjectTableDescriptorSetEntryNVX() *ObjectTableDescriptorSetEntryNVX {
	return (*ObjectTableDescriptorSetEntryNVX)(MemAlloc(unsafe.Sizeof(*(*ObjectTableDescriptorSetEntryNVX)(nil))))
}
func (p *ObjectTableDescriptorSetEntryNVX) Free() { MemFree(unsafe.Pointer(p)) }

// ObjectTableVertexBufferEntryNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTableVertexBufferEntryNVX.html
type ObjectTableVertexBufferEntryNVX struct {
	Type   ObjectEntryTypeNVX
	Flags  ObjectEntryUsageFlagsNVX
	Buffer Buffer
}

func NewObjectTableVertexBufferEntryNVX() *ObjectTableVertexBufferEntryNVX {
	return (*ObjectTableVertexBufferEntryNVX)(MemAlloc(unsafe.Sizeof(*(*ObjectTableVertexBufferEntryNVX)(nil))))
}
func (p *ObjectTableVertexBufferEntryNVX) Free() { MemFree(unsafe.Pointer(p)) }

// ObjectTableIndexBufferEntryNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTableIndexBufferEntryNVX.html
type ObjectTableIndexBufferEntryNVX struct {
	Type      ObjectEntryTypeNVX
	Flags     ObjectEntryUsageFlagsNVX
	Buffer    Buffer
	IndexType IndexType
}

func NewObjectTableIndexBufferEntryNVX() *ObjectTableIndexBufferEntryNVX {
	return (*ObjectTableIndexBufferEntryNVX)(MemAlloc(unsafe.Sizeof(*(*ObjectTableIndexBufferEntryNVX)(nil))))
}
func (p *ObjectTableIndexBufferEntryNVX) Free() { MemFree(unsafe.Pointer(p)) }

// ObjectTablePushConstantEntryNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkObjectTablePushConstantEntryNVX.html
type ObjectTablePushConstantEntryNVX struct {
	Type           ObjectEntryTypeNVX
	Flags          ObjectEntryUsageFlagsNVX
	PipelineLayout PipelineLayout
	StageFlags     ShaderStageFlags
}

func NewObjectTablePushConstantEntryNVX() *ObjectTablePushConstantEntryNVX {
	return (*ObjectTablePushConstantEntryNVX)(MemAlloc(unsafe.Sizeof(*(*ObjectTablePushConstantEntryNVX)(nil))))
}
func (p *ObjectTablePushConstantEntryNVX) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdProcessCommandsNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdProcessCommandsNVX.html
type PfnCmdProcessCommandsNVX uintptr

func (fn PfnCmdProcessCommandsNVX) Call(commandBuffer CommandBuffer, pProcessCommandsInfo *CmdProcessCommandsInfoNVX) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pProcessCommandsInfo)))
}
func (fn PfnCmdProcessCommandsNVX) String() string { return "vkCmdProcessCommandsNVX" }

//  PfnCmdReserveSpaceForCommandsNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdReserveSpaceForCommandsNVX.html
type PfnCmdReserveSpaceForCommandsNVX uintptr

func (fn PfnCmdReserveSpaceForCommandsNVX) Call(commandBuffer CommandBuffer, pReserveSpaceInfo *CmdReserveSpaceForCommandsInfoNVX) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pReserveSpaceInfo)))
}
func (fn PfnCmdReserveSpaceForCommandsNVX) String() string { return "vkCmdReserveSpaceForCommandsNVX" }

//  PfnCreateIndirectCommandsLayoutNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateIndirectCommandsLayoutNVX.html
type PfnCreateIndirectCommandsLayoutNVX uintptr

func (fn PfnCreateIndirectCommandsLayoutNVX) Call(device Device, pCreateInfo *IndirectCommandsLayoutCreateInfoNVX, pAllocator *AllocationCallbacks, pIndirectCommandsLayout *IndirectCommandsLayoutNVX) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pIndirectCommandsLayout)))
	return Result(ret)
}
func (fn PfnCreateIndirectCommandsLayoutNVX) String() string {
	return "vkCreateIndirectCommandsLayoutNVX"
}

//  PfnDestroyIndirectCommandsLayoutNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyIndirectCommandsLayoutNVX.html
type PfnDestroyIndirectCommandsLayoutNVX uintptr

func (fn PfnDestroyIndirectCommandsLayoutNVX) Call(device Device, indirectCommandsLayout IndirectCommandsLayoutNVX, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(indirectCommandsLayout), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyIndirectCommandsLayoutNVX) String() string {
	return "vkDestroyIndirectCommandsLayoutNVX"
}

//  PfnCreateObjectTableNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateObjectTableNVX.html
type PfnCreateObjectTableNVX uintptr

func (fn PfnCreateObjectTableNVX) Call(device Device, pCreateInfo *ObjectTableCreateInfoNVX, pAllocator *AllocationCallbacks, pObjectTable *ObjectTableNVX) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pObjectTable)))
	return Result(ret)
}
func (fn PfnCreateObjectTableNVX) String() string { return "vkCreateObjectTableNVX" }

//  PfnDestroyObjectTableNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyObjectTableNVX.html
type PfnDestroyObjectTableNVX uintptr

func (fn PfnDestroyObjectTableNVX) Call(device Device, objectTable ObjectTableNVX, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(objectTable), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyObjectTableNVX) String() string { return "vkDestroyObjectTableNVX" }

//  PfnRegisterObjectsNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkRegisterObjectsNVX.html
type PfnRegisterObjectsNVX uintptr

func (fn PfnRegisterObjectsNVX) Call(device Device, objectTable ObjectTableNVX, objectCount uint32, ppObjectTableEntries **ObjectTableEntryNVX, pObjectIndices *uint32) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(objectTable), uintptr(objectCount), uintptr(unsafe.Pointer(ppObjectTableEntries)), uintptr(unsafe.Pointer(pObjectIndices)))
	return Result(ret)
}
func (fn PfnRegisterObjectsNVX) String() string { return "vkRegisterObjectsNVX" }

//  PfnUnregisterObjectsNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkUnregisterObjectsNVX.html
type PfnUnregisterObjectsNVX uintptr

func (fn PfnUnregisterObjectsNVX) Call(device Device, objectTable ObjectTableNVX, objectCount uint32, pObjectEntryTypes *ObjectEntryTypeNVX, pObjectIndices *uint32) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(objectTable), uintptr(objectCount), uintptr(unsafe.Pointer(pObjectEntryTypes)), uintptr(unsafe.Pointer(pObjectIndices)))
	return Result(ret)
}
func (fn PfnUnregisterObjectsNVX) String() string { return "vkUnregisterObjectsNVX" }

//  PfnGetPhysicalDeviceGeneratedCommandsPropertiesNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX.html
type PfnGetPhysicalDeviceGeneratedCommandsPropertiesNVX uintptr

func (fn PfnGetPhysicalDeviceGeneratedCommandsPropertiesNVX) Call(physicalDevice PhysicalDevice, pFeatures *DeviceGeneratedCommandsFeaturesNVX, pLimits *DeviceGeneratedCommandsLimitsNVX) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pFeatures)), uintptr(unsafe.Pointer(pLimits)))
}
func (fn PfnGetPhysicalDeviceGeneratedCommandsPropertiesNVX) String() string {
	return "vkGetPhysicalDeviceGeneratedCommandsPropertiesNVX"
}

const NV_clip_space_w_scaling = 1
const NV_CLIP_SPACE_W_SCALING_SPEC_VERSION = 1

var NV_CLIP_SPACE_W_SCALING_EXTENSION_NAME = "VK_NV_clip_space_w_scaling"

// ViewportWScalingNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkViewportWScalingNV.html
type ViewportWScalingNV struct {
	Xcoeff float32
	Ycoeff float32
}

func NewViewportWScalingNV() *ViewportWScalingNV {
	return (*ViewportWScalingNV)(MemAlloc(unsafe.Sizeof(*(*ViewportWScalingNV)(nil))))
}
func (p *ViewportWScalingNV) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineViewportWScalingStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html
type PipelineViewportWScalingStateCreateInfoNV struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	ViewportWScalingEnable Bool32
	ViewportCount          uint32
	PViewportWScalings     *ViewportWScalingNV
}

func NewPipelineViewportWScalingStateCreateInfoNV() *PipelineViewportWScalingStateCreateInfoNV {
	p := (*PipelineViewportWScalingStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineViewportWScalingStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineViewportWScalingStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdSetViewportWScalingNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetViewportWScalingNV.html
type PfnCmdSetViewportWScalingNV uintptr

func (fn PfnCmdSetViewportWScalingNV) Call(commandBuffer CommandBuffer, firstViewport, viewportCount uint32, pViewportWScalings *ViewportWScalingNV) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstViewport), uintptr(viewportCount), uintptr(unsafe.Pointer(pViewportWScalings)))
}
func (fn PfnCmdSetViewportWScalingNV) String() string { return "vkCmdSetViewportWScalingNV" }

const EXT_direct_mode_display = 1
const EXT_DIRECT_MODE_DISPLAY_SPEC_VERSION = 1

var EXT_DIRECT_MODE_DISPLAY_EXTENSION_NAME = "VK_EXT_direct_mode_display"

//  PfnReleaseDisplayEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkReleaseDisplayEXT.html
type PfnReleaseDisplayEXT uintptr

func (fn PfnReleaseDisplayEXT) Call(physicalDevice PhysicalDevice, display DisplayKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(display))
	return Result(ret)
}
func (fn PfnReleaseDisplayEXT) String() string { return "vkReleaseDisplayEXT" }

const EXT_display_surface_counter = 1
const EXT_DISPLAY_SURFACE_COUNTER_SPEC_VERSION = 1

var EXT_DISPLAY_SURFACE_COUNTER_EXTENSION_NAME = "VK_EXT_display_surface_counter"

// SurfaceCounterFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceCounterFlagsEXT.html
type SurfaceCounterFlagsEXT uint32

const (
	SURFACE_COUNTER_VBLANK_EXT             SurfaceCounterFlagsEXT = 0x00000001
	SURFACE_COUNTER_FLAG_BITS_MAX_ENUM_EXT SurfaceCounterFlagsEXT = 0x7FFFFFFF
)

func (x SurfaceCounterFlagsEXT) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch SurfaceCounterFlagsEXT(1 << i) {
			case SURFACE_COUNTER_VBLANK_EXT:
				s += "SURFACE_COUNTER_VBLANK_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// SurfaceCapabilities2EXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceCapabilities2EXT.html
type SurfaceCapabilities2EXT struct {
	SType                    StructureType
	PNext                    unsafe.Pointer
	MinImageCount            uint32
	MaxImageCount            uint32
	CurrentExtent            Extent2D
	MinImageExtent           Extent2D
	MaxImageExtent           Extent2D
	MaxImageArrayLayers      uint32
	SupportedTransforms      SurfaceTransformFlagsKHR
	CurrentTransform         SurfaceTransformFlagsKHR
	SupportedCompositeAlpha  CompositeAlphaFlagsKHR
	SupportedUsageFlags      ImageUsageFlags
	SupportedSurfaceCounters SurfaceCounterFlagsEXT
}

func NewSurfaceCapabilities2EXT() *SurfaceCapabilities2EXT {
	p := (*SurfaceCapabilities2EXT)(MemAlloc(unsafe.Sizeof(*(*SurfaceCapabilities2EXT)(nil))))
	p.SType = STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT
	return p
}
func (p *SurfaceCapabilities2EXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceSurfaceCapabilities2EXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html
type PfnGetPhysicalDeviceSurfaceCapabilities2EXT uintptr

func (fn PfnGetPhysicalDeviceSurfaceCapabilities2EXT) Call(physicalDevice PhysicalDevice, surface SurfaceKHR, pSurfaceCapabilities *SurfaceCapabilities2EXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(surface), uintptr(unsafe.Pointer(pSurfaceCapabilities)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfaceCapabilities2EXT) String() string {
	return "vkGetPhysicalDeviceSurfaceCapabilities2EXT"
}

const EXT_display_control = 1
const EXT_DISPLAY_CONTROL_SPEC_VERSION = 1

var EXT_DISPLAY_CONTROL_EXTENSION_NAME = "VK_EXT_display_control"

// DisplayPowerStateEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPowerStateEXT.html
type DisplayPowerStateEXT int32

const (
	DISPLAY_POWER_STATE_OFF_EXT         DisplayPowerStateEXT = 0
	DISPLAY_POWER_STATE_SUSPEND_EXT     DisplayPowerStateEXT = 1
	DISPLAY_POWER_STATE_ON_EXT          DisplayPowerStateEXT = 2
	DISPLAY_POWER_STATE_BEGIN_RANGE_EXT DisplayPowerStateEXT = DISPLAY_POWER_STATE_OFF_EXT
	DISPLAY_POWER_STATE_END_RANGE_EXT   DisplayPowerStateEXT = DISPLAY_POWER_STATE_ON_EXT
	DISPLAY_POWER_STATE_RANGE_SIZE_EXT  DisplayPowerStateEXT = (DISPLAY_POWER_STATE_ON_EXT - DISPLAY_POWER_STATE_OFF_EXT + 1)
	DISPLAY_POWER_STATE_MAX_ENUM_EXT    DisplayPowerStateEXT = 0x7FFFFFFF
)

func (x DisplayPowerStateEXT) String() string {
	switch x {
	case DISPLAY_POWER_STATE_OFF_EXT:
		return "DISPLAY_POWER_STATE_OFF_EXT"
	case DISPLAY_POWER_STATE_SUSPEND_EXT:
		return "DISPLAY_POWER_STATE_SUSPEND_EXT"
	case DISPLAY_POWER_STATE_ON_EXT:
		return "DISPLAY_POWER_STATE_ON_EXT"
	case DISPLAY_POWER_STATE_MAX_ENUM_EXT:
		return "DISPLAY_POWER_STATE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DeviceEventTypeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceEventTypeEXT.html
type DeviceEventTypeEXT int32

const (
	DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT DeviceEventTypeEXT = 0
	DEVICE_EVENT_TYPE_BEGIN_RANGE_EXT     DeviceEventTypeEXT = DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT
	DEVICE_EVENT_TYPE_END_RANGE_EXT       DeviceEventTypeEXT = DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT
	DEVICE_EVENT_TYPE_RANGE_SIZE_EXT      DeviceEventTypeEXT = (DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT - DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT + 1)
	DEVICE_EVENT_TYPE_MAX_ENUM_EXT        DeviceEventTypeEXT = 0x7FFFFFFF
)

func (x DeviceEventTypeEXT) String() string {
	switch x {
	case DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT:
		return "DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT"
	case DEVICE_EVENT_TYPE_MAX_ENUM_EXT:
		return "DEVICE_EVENT_TYPE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DisplayEventTypeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayEventTypeEXT.html
type DisplayEventTypeEXT int32

const (
	DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT DisplayEventTypeEXT = 0
	DISPLAY_EVENT_TYPE_BEGIN_RANGE_EXT     DisplayEventTypeEXT = DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT
	DISPLAY_EVENT_TYPE_END_RANGE_EXT       DisplayEventTypeEXT = DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT
	DISPLAY_EVENT_TYPE_RANGE_SIZE_EXT      DisplayEventTypeEXT = (DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT - DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT + 1)
	DISPLAY_EVENT_TYPE_MAX_ENUM_EXT        DisplayEventTypeEXT = 0x7FFFFFFF
)

func (x DisplayEventTypeEXT) String() string {
	switch x {
	case DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT:
		return "DISPLAY_EVENT_TYPE_FIRST_PIXEL_OUT_EXT"
	case DISPLAY_EVENT_TYPE_MAX_ENUM_EXT:
		return "DISPLAY_EVENT_TYPE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DisplayPowerInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayPowerInfoEXT.html
type DisplayPowerInfoEXT struct {
	SType      StructureType
	PNext      unsafe.Pointer
	PowerState DisplayPowerStateEXT
}

func NewDisplayPowerInfoEXT() *DisplayPowerInfoEXT {
	p := (*DisplayPowerInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DisplayPowerInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT
	return p
}
func (p *DisplayPowerInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DeviceEventInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceEventInfoEXT.html
type DeviceEventInfoEXT struct {
	SType       StructureType
	PNext       unsafe.Pointer
	DeviceEvent DeviceEventTypeEXT
}

func NewDeviceEventInfoEXT() *DeviceEventInfoEXT {
	p := (*DeviceEventInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DeviceEventInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT
	return p
}
func (p *DeviceEventInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DisplayEventInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayEventInfoEXT.html
type DisplayEventInfoEXT struct {
	SType        StructureType
	PNext        unsafe.Pointer
	DisplayEvent DisplayEventTypeEXT
}

func NewDisplayEventInfoEXT() *DisplayEventInfoEXT {
	p := (*DisplayEventInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DisplayEventInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT
	return p
}
func (p *DisplayEventInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// SwapchainCounterCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSwapchainCounterCreateInfoEXT.html
type SwapchainCounterCreateInfoEXT struct {
	SType           StructureType
	PNext           unsafe.Pointer
	SurfaceCounters SurfaceCounterFlagsEXT
}

func NewSwapchainCounterCreateInfoEXT() *SwapchainCounterCreateInfoEXT {
	p := (*SwapchainCounterCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*SwapchainCounterCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT
	return p
}
func (p *SwapchainCounterCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnDisplayPowerControlEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDisplayPowerControlEXT.html
type PfnDisplayPowerControlEXT uintptr

func (fn PfnDisplayPowerControlEXT) Call(device Device, display DisplayKHR, pDisplayPowerInfo *DisplayPowerInfoEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(display), uintptr(unsafe.Pointer(pDisplayPowerInfo)))
	return Result(ret)
}
func (fn PfnDisplayPowerControlEXT) String() string { return "vkDisplayPowerControlEXT" }

//  PfnRegisterDeviceEventEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkRegisterDeviceEventEXT.html
type PfnRegisterDeviceEventEXT uintptr

func (fn PfnRegisterDeviceEventEXT) Call(device Device, pDeviceEventInfo *DeviceEventInfoEXT, pAllocator *AllocationCallbacks, pFence *Fence) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pDeviceEventInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFence)))
	return Result(ret)
}
func (fn PfnRegisterDeviceEventEXT) String() string { return "vkRegisterDeviceEventEXT" }

//  PfnRegisterDisplayEventEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkRegisterDisplayEventEXT.html
type PfnRegisterDisplayEventEXT uintptr

func (fn PfnRegisterDisplayEventEXT) Call(device Device, display DisplayKHR, pDisplayEventInfo *DisplayEventInfoEXT, pAllocator *AllocationCallbacks, pFence *Fence) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(display), uintptr(unsafe.Pointer(pDisplayEventInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pFence)))
	return Result(ret)
}
func (fn PfnRegisterDisplayEventEXT) String() string { return "vkRegisterDisplayEventEXT" }

//  PfnGetSwapchainCounterEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetSwapchainCounterEXT.html
type PfnGetSwapchainCounterEXT uintptr

func (fn PfnGetSwapchainCounterEXT) Call(device Device, swapchain SwapchainKHR, counter SurfaceCounterFlagsEXT, pCounterValue *uint64) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain), uintptr(counter), uintptr(unsafe.Pointer(pCounterValue)))
	return Result(ret)
}
func (fn PfnGetSwapchainCounterEXT) String() string { return "vkGetSwapchainCounterEXT" }

const GOOGLE_display_timing = 1
const GOOGLE_DISPLAY_TIMING_SPEC_VERSION = 1

var GOOGLE_DISPLAY_TIMING_EXTENSION_NAME = "VK_GOOGLE_display_timing"

// RefreshCycleDurationGOOGLE -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRefreshCycleDurationGOOGLE.html
type RefreshCycleDurationGOOGLE struct {
	RefreshDuration uint64
}

func NewRefreshCycleDurationGOOGLE() *RefreshCycleDurationGOOGLE {
	return (*RefreshCycleDurationGOOGLE)(MemAlloc(unsafe.Sizeof(*(*RefreshCycleDurationGOOGLE)(nil))))
}
func (p *RefreshCycleDurationGOOGLE) Free() { MemFree(unsafe.Pointer(p)) }

// PastPresentationTimingGOOGLE -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPastPresentationTimingGOOGLE.html
type PastPresentationTimingGOOGLE struct {
	PresentID           uint32
	DesiredPresentTime  uint64
	ActualPresentTime   uint64
	EarliestPresentTime uint64
	PresentMargin       uint64
}

func NewPastPresentationTimingGOOGLE() *PastPresentationTimingGOOGLE {
	return (*PastPresentationTimingGOOGLE)(MemAlloc(unsafe.Sizeof(*(*PastPresentationTimingGOOGLE)(nil))))
}
func (p *PastPresentationTimingGOOGLE) Free() { MemFree(unsafe.Pointer(p)) }

// PresentTimeGOOGLE -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPresentTimeGOOGLE.html
type PresentTimeGOOGLE struct {
	PresentID          uint32
	DesiredPresentTime uint64
}

func NewPresentTimeGOOGLE() *PresentTimeGOOGLE {
	return (*PresentTimeGOOGLE)(MemAlloc(unsafe.Sizeof(*(*PresentTimeGOOGLE)(nil))))
}
func (p *PresentTimeGOOGLE) Free() { MemFree(unsafe.Pointer(p)) }

// PresentTimesInfoGOOGLE -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPresentTimesInfoGOOGLE.html
type PresentTimesInfoGOOGLE struct {
	SType          StructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PTimes         *PresentTimeGOOGLE
}

func NewPresentTimesInfoGOOGLE() *PresentTimesInfoGOOGLE {
	p := (*PresentTimesInfoGOOGLE)(MemAlloc(unsafe.Sizeof(*(*PresentTimesInfoGOOGLE)(nil))))
	p.SType = STRUCTURE_TYPE_PRESENT_TIMES_INFO_GOOGLE
	return p
}
func (p *PresentTimesInfoGOOGLE) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetRefreshCycleDurationGOOGLE -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetRefreshCycleDurationGOOGLE.html
type PfnGetRefreshCycleDurationGOOGLE uintptr

func (fn PfnGetRefreshCycleDurationGOOGLE) Call(device Device, swapchain SwapchainKHR, pDisplayTimingProperties *RefreshCycleDurationGOOGLE) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pDisplayTimingProperties)))
	return Result(ret)
}
func (fn PfnGetRefreshCycleDurationGOOGLE) String() string { return "vkGetRefreshCycleDurationGOOGLE" }

//  PfnGetPastPresentationTimingGOOGLE -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPastPresentationTimingGOOGLE.html
type PfnGetPastPresentationTimingGOOGLE uintptr

func (fn PfnGetPastPresentationTimingGOOGLE) Call(device Device, swapchain SwapchainKHR, pPresentationTimingCount *uint32, pPresentationTimings *PastPresentationTimingGOOGLE) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain), uintptr(unsafe.Pointer(pPresentationTimingCount)), uintptr(unsafe.Pointer(pPresentationTimings)))
	return Result(ret)
}
func (fn PfnGetPastPresentationTimingGOOGLE) String() string {
	return "vkGetPastPresentationTimingGOOGLE"
}

const NV_sample_mask_override_coverage = 1
const NV_SAMPLE_MASK_OVERRIDE_COVERAGE_SPEC_VERSION = 1

var NV_SAMPLE_MASK_OVERRIDE_COVERAGE_EXTENSION_NAME = "VK_NV_sample_mask_override_coverage"

const NV_geometry_shader_passthrough = 1
const NV_GEOMETRY_SHADER_PASSTHROUGH_SPEC_VERSION = 1

var NV_GEOMETRY_SHADER_PASSTHROUGH_EXTENSION_NAME = "VK_NV_geometry_shader_passthrough"

const NV_viewport_array2 = 1
const NV_VIEWPORT_ARRAY2_SPEC_VERSION = 1

var NV_VIEWPORT_ARRAY2_EXTENSION_NAME = "VK_NV_viewport_array2"

const NVX_multiview_per_view_attributes = 1
const NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_SPEC_VERSION = 1

var NVX_MULTIVIEW_PER_VIEW_ATTRIBUTES_EXTENSION_NAME = "VK_NVX_multiview_per_view_attributes"

// PhysicalDeviceMultiviewPerViewAttributesPropertiesNVX -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX.html
type PhysicalDeviceMultiviewPerViewAttributesPropertiesNVX struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	PerViewPositionAllComponents Bool32
}

func NewPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX() *PhysicalDeviceMultiviewPerViewAttributesPropertiesNVX {
	p := (*PhysicalDeviceMultiviewPerViewAttributesPropertiesNVX)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMultiviewPerViewAttributesPropertiesNVX)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_ATTRIBUTES_PROPERTIES_NVX
	return p
}
func (p *PhysicalDeviceMultiviewPerViewAttributesPropertiesNVX) Free() { MemFree(unsafe.Pointer(p)) }

const NV_viewport_swizzle = 1
const NV_VIEWPORT_SWIZZLE_SPEC_VERSION = 1

var NV_VIEWPORT_SWIZZLE_EXTENSION_NAME = "VK_NV_viewport_swizzle"

// ViewportCoordinateSwizzleNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkViewportCoordinateSwizzleNV.html
type ViewportCoordinateSwizzleNV int32

const (
	VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV  ViewportCoordinateSwizzleNV = 0
	VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_X_NV  ViewportCoordinateSwizzleNV = 1
	VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Y_NV  ViewportCoordinateSwizzleNV = 2
	VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Y_NV  ViewportCoordinateSwizzleNV = 3
	VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Z_NV  ViewportCoordinateSwizzleNV = 4
	VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Z_NV  ViewportCoordinateSwizzleNV = 5
	VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_W_NV  ViewportCoordinateSwizzleNV = 6
	VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV  ViewportCoordinateSwizzleNV = 7
	VIEWPORT_COORDINATE_SWIZZLE_BEGIN_RANGE_NV ViewportCoordinateSwizzleNV = VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV
	VIEWPORT_COORDINATE_SWIZZLE_END_RANGE_NV   ViewportCoordinateSwizzleNV = VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV
	VIEWPORT_COORDINATE_SWIZZLE_RANGE_SIZE_NV  ViewportCoordinateSwizzleNV = (VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV - VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV + 1)
	VIEWPORT_COORDINATE_SWIZZLE_MAX_ENUM_NV    ViewportCoordinateSwizzleNV = 0x7FFFFFFF
)

func (x ViewportCoordinateSwizzleNV) String() string {
	switch x {
	case VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_X_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_X_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_X_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Y_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Y_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Y_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Y_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Z_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_Z_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Z_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_Z_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_W_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_POSITIVE_W_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_NEGATIVE_W_NV"
	case VIEWPORT_COORDINATE_SWIZZLE_MAX_ENUM_NV:
		return "VIEWPORT_COORDINATE_SWIZZLE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

type PipelineViewportSwizzleStateCreateFlagsNV uint32 // reserved
// ViewportSwizzleNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkViewportSwizzleNV.html
type ViewportSwizzleNV struct {
	X ViewportCoordinateSwizzleNV
	Y ViewportCoordinateSwizzleNV
	Z ViewportCoordinateSwizzleNV
	W ViewportCoordinateSwizzleNV
}

func NewViewportSwizzleNV() *ViewportSwizzleNV {
	return (*ViewportSwizzleNV)(MemAlloc(unsafe.Sizeof(*(*ViewportSwizzleNV)(nil))))
}
func (p *ViewportSwizzleNV) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineViewportSwizzleStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html
type PipelineViewportSwizzleStateCreateInfoNV struct {
	SType             StructureType
	PNext             unsafe.Pointer
	Flags             PipelineViewportSwizzleStateCreateFlagsNV
	ViewportCount     uint32
	PViewportSwizzles *ViewportSwizzleNV
}

func NewPipelineViewportSwizzleStateCreateInfoNV() *PipelineViewportSwizzleStateCreateInfoNV {
	p := (*PipelineViewportSwizzleStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineViewportSwizzleStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineViewportSwizzleStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_discard_rectangles = 1
const EXT_DISCARD_RECTANGLES_SPEC_VERSION = 1

var EXT_DISCARD_RECTANGLES_EXTENSION_NAME = "VK_EXT_discard_rectangles"

// DiscardRectangleModeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDiscardRectangleModeEXT.html
type DiscardRectangleModeEXT int32

const (
	DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT   DiscardRectangleModeEXT = 0
	DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT   DiscardRectangleModeEXT = 1
	DISCARD_RECTANGLE_MODE_BEGIN_RANGE_EXT DiscardRectangleModeEXT = DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT
	DISCARD_RECTANGLE_MODE_END_RANGE_EXT   DiscardRectangleModeEXT = DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT
	DISCARD_RECTANGLE_MODE_RANGE_SIZE_EXT  DiscardRectangleModeEXT = (DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT - DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT + 1)
	DISCARD_RECTANGLE_MODE_MAX_ENUM_EXT    DiscardRectangleModeEXT = 0x7FFFFFFF
)

func (x DiscardRectangleModeEXT) String() string {
	switch x {
	case DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT:
		return "DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT"
	case DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT:
		return "DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT"
	case DISCARD_RECTANGLE_MODE_MAX_ENUM_EXT:
		return "DISCARD_RECTANGLE_MODE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

type PipelineDiscardRectangleStateCreateFlagsEXT uint32 // reserved
// PhysicalDeviceDiscardRectanglePropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceDiscardRectanglePropertiesEXT.html
type PhysicalDeviceDiscardRectanglePropertiesEXT struct {
	SType                StructureType
	PNext                unsafe.Pointer
	MaxDiscardRectangles uint32
}

func NewPhysicalDeviceDiscardRectanglePropertiesEXT() *PhysicalDeviceDiscardRectanglePropertiesEXT {
	p := (*PhysicalDeviceDiscardRectanglePropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceDiscardRectanglePropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceDiscardRectanglePropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineDiscardRectangleStateCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html
type PipelineDiscardRectangleStateCreateInfoEXT struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Flags                 PipelineDiscardRectangleStateCreateFlagsEXT
	DiscardRectangleMode  DiscardRectangleModeEXT
	DiscardRectangleCount uint32
	PDiscardRectangles    *Rect2D
}

func NewPipelineDiscardRectangleStateCreateInfoEXT() *PipelineDiscardRectangleStateCreateInfoEXT {
	p := (*PipelineDiscardRectangleStateCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineDiscardRectangleStateCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT
	return p
}
func (p *PipelineDiscardRectangleStateCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdSetDiscardRectangleEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetDiscardRectangleEXT.html
type PfnCmdSetDiscardRectangleEXT uintptr

func (fn PfnCmdSetDiscardRectangleEXT) Call(commandBuffer CommandBuffer, firstDiscardRectangle, discardRectangleCount uint32, pDiscardRectangles *Rect2D) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstDiscardRectangle), uintptr(discardRectangleCount), uintptr(unsafe.Pointer(pDiscardRectangles)))
}
func (fn PfnCmdSetDiscardRectangleEXT) String() string { return "vkCmdSetDiscardRectangleEXT" }

const EXT_conservative_rasterization = 1
const EXT_CONSERVATIVE_RASTERIZATION_SPEC_VERSION = 1

var EXT_CONSERVATIVE_RASTERIZATION_EXTENSION_NAME = "VK_EXT_conservative_rasterization"

// ConservativeRasterizationModeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkConservativeRasterizationModeEXT.html
type ConservativeRasterizationModeEXT int32

const (
	CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT      ConservativeRasterizationModeEXT = 0
	CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT  ConservativeRasterizationModeEXT = 1
	CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT ConservativeRasterizationModeEXT = 2
	CONSERVATIVE_RASTERIZATION_MODE_BEGIN_RANGE_EXT   ConservativeRasterizationModeEXT = CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT
	CONSERVATIVE_RASTERIZATION_MODE_END_RANGE_EXT     ConservativeRasterizationModeEXT = CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT
	CONSERVATIVE_RASTERIZATION_MODE_RANGE_SIZE_EXT    ConservativeRasterizationModeEXT = (CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT - CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT + 1)
	CONSERVATIVE_RASTERIZATION_MODE_MAX_ENUM_EXT      ConservativeRasterizationModeEXT = 0x7FFFFFFF
)

func (x ConservativeRasterizationModeEXT) String() string {
	switch x {
	case CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT:
		return "CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT"
	case CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT:
		return "CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT"
	case CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT:
		return "CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT"
	case CONSERVATIVE_RASTERIZATION_MODE_MAX_ENUM_EXT:
		return "CONSERVATIVE_RASTERIZATION_MODE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

type PipelineRasterizationConservativeStateCreateFlagsEXT uint32 // reserved
// PhysicalDeviceConservativeRasterizationPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceConservativeRasterizationPropertiesEXT.html
type PhysicalDeviceConservativeRasterizationPropertiesEXT struct {
	SType                                       StructureType
	PNext                                       unsafe.Pointer
	PrimitiveOverestimationSize                 float32
	MaxExtraPrimitiveOverestimationSize         float32
	ExtraPrimitiveOverestimationSizeGranularity float32
	PrimitiveUnderestimation                    Bool32
	ConservativePointAndLineRasterization       Bool32
	DegenerateTrianglesRasterized               Bool32
	DegenerateLinesRasterized                   Bool32
	FullyCoveredFragmentShaderInputVariable     Bool32
	ConservativeRasterizationPostDepthCoverage  Bool32
}

func NewPhysicalDeviceConservativeRasterizationPropertiesEXT() *PhysicalDeviceConservativeRasterizationPropertiesEXT {
	p := (*PhysicalDeviceConservativeRasterizationPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceConservativeRasterizationPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceConservativeRasterizationPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineRasterizationConservativeStateCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html
type PipelineRasterizationConservativeStateCreateInfoEXT struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	Flags                            PipelineRasterizationConservativeStateCreateFlagsEXT
	ConservativeRasterizationMode    ConservativeRasterizationModeEXT
	ExtraPrimitiveOverestimationSize float32
}

func NewPipelineRasterizationConservativeStateCreateInfoEXT() *PipelineRasterizationConservativeStateCreateInfoEXT {
	p := (*PipelineRasterizationConservativeStateCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineRasterizationConservativeStateCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT
	return p
}
func (p *PipelineRasterizationConservativeStateCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_depth_clip_enable = 1
const EXT_DEPTH_CLIP_ENABLE_SPEC_VERSION = 1

var EXT_DEPTH_CLIP_ENABLE_EXTENSION_NAME = "VK_EXT_depth_clip_enable"

type PipelineRasterizationDepthClipStateCreateFlagsEXT uint32 // reserved
// PhysicalDeviceDepthClipEnableFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceDepthClipEnableFeaturesEXT.html
type PhysicalDeviceDepthClipEnableFeaturesEXT struct {
	SType           StructureType
	PNext           unsafe.Pointer
	DepthClipEnable Bool32
}

func NewPhysicalDeviceDepthClipEnableFeaturesEXT() *PhysicalDeviceDepthClipEnableFeaturesEXT {
	p := (*PhysicalDeviceDepthClipEnableFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceDepthClipEnableFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_ENABLE_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceDepthClipEnableFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineRasterizationDepthClipStateCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html
type PipelineRasterizationDepthClipStateCreateInfoEXT struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           PipelineRasterizationDepthClipStateCreateFlagsEXT
	DepthClipEnable Bool32
}

func NewPipelineRasterizationDepthClipStateCreateInfoEXT() *PipelineRasterizationDepthClipStateCreateInfoEXT {
	p := (*PipelineRasterizationDepthClipStateCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineRasterizationDepthClipStateCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT
	return p
}
func (p *PipelineRasterizationDepthClipStateCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_swapchain_colorspace = 1
const EXT_SWAPCHAIN_COLOR_SPACE_SPEC_VERSION = 4

var EXT_SWAPCHAIN_COLOR_SPACE_EXTENSION_NAME = "VK_EXT_swapchain_colorspace"

const EXT_hdr_metadata = 1
const EXT_HDR_METADATA_SPEC_VERSION = 2

var EXT_HDR_METADATA_EXTENSION_NAME = "VK_EXT_hdr_metadata"

// XYColorEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkXYColorEXT.html
type XYColorEXT struct {
	X float32
	Y float32
}

func NewXYColorEXT() *XYColorEXT { return (*XYColorEXT)(MemAlloc(unsafe.Sizeof(*(*XYColorEXT)(nil)))) }
func (p *XYColorEXT) Free()      { MemFree(unsafe.Pointer(p)) }

// HdrMetadataEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkHdrMetadataEXT.html
type HdrMetadataEXT struct {
	SType                     StructureType
	PNext                     unsafe.Pointer
	DisplayPrimaryRed         XYColorEXT
	DisplayPrimaryGreen       XYColorEXT
	DisplayPrimaryBlue        XYColorEXT
	WhitePoint                XYColorEXT
	MaxLuminance              float32
	MinLuminance              float32
	MaxContentLightLevel      float32
	MaxFrameAverageLightLevel float32
}

func NewHdrMetadataEXT() *HdrMetadataEXT {
	p := (*HdrMetadataEXT)(MemAlloc(unsafe.Sizeof(*(*HdrMetadataEXT)(nil))))
	p.SType = STRUCTURE_TYPE_HDR_METADATA_EXT
	return p
}
func (p *HdrMetadataEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnSetHdrMetadataEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkSetHdrMetadataEXT.html
type PfnSetHdrMetadataEXT uintptr

func (fn PfnSetHdrMetadataEXT) Call(device Device, swapchainCount uint32, pSwapchains *SwapchainKHR, pMetadata *HdrMetadataEXT) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(swapchainCount), uintptr(unsafe.Pointer(pSwapchains)), uintptr(unsafe.Pointer(pMetadata)))
}
func (fn PfnSetHdrMetadataEXT) String() string { return "vkSetHdrMetadataEXT" }

const EXT_external_memory_dma_buf = 1
const EXT_EXTERNAL_MEMORY_DMA_BUF_SPEC_VERSION = 1

var EXT_EXTERNAL_MEMORY_DMA_BUF_EXTENSION_NAME = "VK_EXT_external_memory_dma_buf"

const EXT_queue_family_foreign = 1
const EXT_QUEUE_FAMILY_FOREIGN_SPEC_VERSION = 1

var EXT_QUEUE_FAMILY_FOREIGN_EXTENSION_NAME = "VK_EXT_queue_family_foreign"

const EXT_debug_utils = 1

// DebugUtilsMessengerEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsMessengerEXT.html
type DebugUtilsMessengerEXT NonDispatchableHandle

const EXT_DEBUG_UTILS_SPEC_VERSION = 1

var EXT_DEBUG_UTILS_EXTENSION_NAME = "VK_EXT_debug_utils"

type DebugUtilsMessengerCallbackDataFlagsEXT uint32 // reserved
type DebugUtilsMessengerCreateFlagsEXT uint32       // reserved
// DebugUtilsMessageSeverityFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsMessageSeverityFlagsEXT.html
type DebugUtilsMessageSeverityFlagsEXT uint32

const (
	DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT        DebugUtilsMessageSeverityFlagsEXT = 0x00000001
	DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT           DebugUtilsMessageSeverityFlagsEXT = 0x00000010
	DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT        DebugUtilsMessageSeverityFlagsEXT = 0x00000100
	DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT          DebugUtilsMessageSeverityFlagsEXT = 0x00001000
	DEBUG_UTILS_MESSAGE_SEVERITY_FLAG_BITS_MAX_ENUM_EXT DebugUtilsMessageSeverityFlagsEXT = 0x7FFFFFFF
)

func (x DebugUtilsMessageSeverityFlagsEXT) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DebugUtilsMessageSeverityFlagsEXT(1 << i) {
			case DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT:
				s += "DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT|"
			case DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT:
				s += "DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT|"
			case DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT:
				s += "DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT|"
			case DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT:
				s += "DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DebugUtilsMessageTypeFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsMessageTypeFlagsEXT.html
type DebugUtilsMessageTypeFlagsEXT uint32

const (
	DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT        DebugUtilsMessageTypeFlagsEXT = 0x00000001
	DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT     DebugUtilsMessageTypeFlagsEXT = 0x00000002
	DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT    DebugUtilsMessageTypeFlagsEXT = 0x00000004
	DEBUG_UTILS_MESSAGE_TYPE_FLAG_BITS_MAX_ENUM_EXT DebugUtilsMessageTypeFlagsEXT = 0x7FFFFFFF
)

func (x DebugUtilsMessageTypeFlagsEXT) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DebugUtilsMessageTypeFlagsEXT(1 << i) {
			case DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT:
				s += "DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT|"
			case DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT:
				s += "DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT|"
			case DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT:
				s += "DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DebugUtilsObjectNameInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html
type DebugUtilsObjectNameInfoEXT struct {
	SType        StructureType
	PNext        unsafe.Pointer
	ObjectType   ObjectType
	ObjectHandle uint64
	PObjectName  *int8
}

func NewDebugUtilsObjectNameInfoEXT() *DebugUtilsObjectNameInfoEXT {
	p := (*DebugUtilsObjectNameInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DebugUtilsObjectNameInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT
	return p
}
func (p *DebugUtilsObjectNameInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DebugUtilsObjectTagInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsObjectTagInfoEXT.html
type DebugUtilsObjectTagInfoEXT struct {
	SType        StructureType
	PNext        unsafe.Pointer
	ObjectType   ObjectType
	ObjectHandle uint64
	TagName      uint64
	TagSize      uintptr
	PTag         unsafe.Pointer
}

func NewDebugUtilsObjectTagInfoEXT() *DebugUtilsObjectTagInfoEXT {
	p := (*DebugUtilsObjectTagInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DebugUtilsObjectTagInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT
	return p
}
func (p *DebugUtilsObjectTagInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DebugUtilsLabelEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsLabelEXT.html
type DebugUtilsLabelEXT struct {
	SType      StructureType
	PNext      unsafe.Pointer
	PLabelName *int8
	Color      [4]float32
}

func NewDebugUtilsLabelEXT() *DebugUtilsLabelEXT {
	p := (*DebugUtilsLabelEXT)(MemAlloc(unsafe.Sizeof(*(*DebugUtilsLabelEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
	return p
}
func (p *DebugUtilsLabelEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DebugUtilsMessengerCallbackDataEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html
type DebugUtilsMessengerCallbackDataEXT struct {
	SType            StructureType
	PNext            unsafe.Pointer
	Flags            DebugUtilsMessengerCallbackDataFlagsEXT
	PMessageIdName   *int8
	MessageIdNumber  int32
	PMessage         *int8
	QueueLabelCount  uint32
	PQueueLabels     *DebugUtilsLabelEXT
	CmdBufLabelCount uint32
	PCmdBufLabels    *DebugUtilsLabelEXT
	ObjectCount      uint32
	PObjects         *DebugUtilsObjectNameInfoEXT
}

func NewDebugUtilsMessengerCallbackDataEXT() *DebugUtilsMessengerCallbackDataEXT {
	p := (*DebugUtilsMessengerCallbackDataEXT)(MemAlloc(unsafe.Sizeof(*(*DebugUtilsMessengerCallbackDataEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT
	return p
}
func (p *DebugUtilsMessengerCallbackDataEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnDebugUtilsMessengerCallbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDebugUtilsMessengerCallbackEXT.html
type PfnDebugUtilsMessengerCallbackEXT uintptr

func (fn PfnDebugUtilsMessengerCallbackEXT) Call(messageSeverity DebugUtilsMessageSeverityFlagsEXT, messageTypes DebugUtilsMessageTypeFlagsEXT, pCallbackData *DebugUtilsMessengerCallbackDataEXT, pUserData unsafe.Pointer) Bool32 {
	ret, _, _ := call(uintptr(fn), uintptr(messageSeverity), uintptr(messageTypes), uintptr(unsafe.Pointer(pCallbackData)), uintptr(pUserData))
	return Bool32(ret)
}
func (fn PfnDebugUtilsMessengerCallbackEXT) String() string {
	return "vkDebugUtilsMessengerCallbackEXT"
}

// DebugUtilsMessengerCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html
type DebugUtilsMessengerCreateInfoEXT struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           DebugUtilsMessengerCreateFlagsEXT
	MessageSeverity DebugUtilsMessageSeverityFlagsEXT
	MessageType     DebugUtilsMessageTypeFlagsEXT
	PfnUserCallback PfnDebugUtilsMessengerCallbackEXT
	PUserData       unsafe.Pointer
}

func NewDebugUtilsMessengerCreateInfoEXT() *DebugUtilsMessengerCreateInfoEXT {
	p := (*DebugUtilsMessengerCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DebugUtilsMessengerCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
	return p
}
func (p *DebugUtilsMessengerCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnSetDebugUtilsObjectNameEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkSetDebugUtilsObjectNameEXT.html
type PfnSetDebugUtilsObjectNameEXT uintptr

func (fn PfnSetDebugUtilsObjectNameEXT) Call(device Device, pNameInfo *DebugUtilsObjectNameInfoEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pNameInfo)))
	return Result(ret)
}
func (fn PfnSetDebugUtilsObjectNameEXT) String() string { return "vkSetDebugUtilsObjectNameEXT" }

//  PfnSetDebugUtilsObjectTagEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkSetDebugUtilsObjectTagEXT.html
type PfnSetDebugUtilsObjectTagEXT uintptr

func (fn PfnSetDebugUtilsObjectTagEXT) Call(device Device, pTagInfo *DebugUtilsObjectTagInfoEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pTagInfo)))
	return Result(ret)
}
func (fn PfnSetDebugUtilsObjectTagEXT) String() string { return "vkSetDebugUtilsObjectTagEXT" }

//  PfnQueueBeginDebugUtilsLabelEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueueBeginDebugUtilsLabelEXT.html
type PfnQueueBeginDebugUtilsLabelEXT uintptr

func (fn PfnQueueBeginDebugUtilsLabelEXT) Call(queue Queue, pLabelInfo *DebugUtilsLabelEXT) {
	_, _, _ = call(uintptr(fn), uintptr(queue), uintptr(unsafe.Pointer(pLabelInfo)))
}
func (fn PfnQueueBeginDebugUtilsLabelEXT) String() string { return "vkQueueBeginDebugUtilsLabelEXT" }

//  PfnQueueEndDebugUtilsLabelEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueueEndDebugUtilsLabelEXT.html
type PfnQueueEndDebugUtilsLabelEXT uintptr

func (fn PfnQueueEndDebugUtilsLabelEXT) Call(queue Queue) {
	_, _, _ = call(uintptr(fn), uintptr(queue))
}
func (fn PfnQueueEndDebugUtilsLabelEXT) String() string { return "vkQueueEndDebugUtilsLabelEXT" }

//  PfnQueueInsertDebugUtilsLabelEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueueInsertDebugUtilsLabelEXT.html
type PfnQueueInsertDebugUtilsLabelEXT uintptr

func (fn PfnQueueInsertDebugUtilsLabelEXT) Call(queue Queue, pLabelInfo *DebugUtilsLabelEXT) {
	_, _, _ = call(uintptr(fn), uintptr(queue), uintptr(unsafe.Pointer(pLabelInfo)))
}
func (fn PfnQueueInsertDebugUtilsLabelEXT) String() string { return "vkQueueInsertDebugUtilsLabelEXT" }

//  PfnCmdBeginDebugUtilsLabelEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBeginDebugUtilsLabelEXT.html
type PfnCmdBeginDebugUtilsLabelEXT uintptr

func (fn PfnCmdBeginDebugUtilsLabelEXT) Call(commandBuffer CommandBuffer, pLabelInfo *DebugUtilsLabelEXT) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pLabelInfo)))
}
func (fn PfnCmdBeginDebugUtilsLabelEXT) String() string { return "vkCmdBeginDebugUtilsLabelEXT" }

//  PfnCmdEndDebugUtilsLabelEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdEndDebugUtilsLabelEXT.html
type PfnCmdEndDebugUtilsLabelEXT uintptr

func (fn PfnCmdEndDebugUtilsLabelEXT) Call(commandBuffer CommandBuffer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer))
}
func (fn PfnCmdEndDebugUtilsLabelEXT) String() string { return "vkCmdEndDebugUtilsLabelEXT" }

//  PfnCmdInsertDebugUtilsLabelEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdInsertDebugUtilsLabelEXT.html
type PfnCmdInsertDebugUtilsLabelEXT uintptr

func (fn PfnCmdInsertDebugUtilsLabelEXT) Call(commandBuffer CommandBuffer, pLabelInfo *DebugUtilsLabelEXT) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pLabelInfo)))
}
func (fn PfnCmdInsertDebugUtilsLabelEXT) String() string { return "vkCmdInsertDebugUtilsLabelEXT" }

//  PfnCreateDebugUtilsMessengerEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html
type PfnCreateDebugUtilsMessengerEXT uintptr

func (fn PfnCreateDebugUtilsMessengerEXT) Call(instance Instance, pCreateInfo *DebugUtilsMessengerCreateInfoEXT, pAllocator *AllocationCallbacks, pMessenger *DebugUtilsMessengerEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pMessenger)))
	return Result(ret)
}
func (fn PfnCreateDebugUtilsMessengerEXT) String() string { return "vkCreateDebugUtilsMessengerEXT" }

//  PfnDestroyDebugUtilsMessengerEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyDebugUtilsMessengerEXT.html
type PfnDestroyDebugUtilsMessengerEXT uintptr

func (fn PfnDestroyDebugUtilsMessengerEXT) Call(instance Instance, messenger DebugUtilsMessengerEXT, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(instance), uintptr(messenger), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyDebugUtilsMessengerEXT) String() string { return "vkDestroyDebugUtilsMessengerEXT" }

//  PfnSubmitDebugUtilsMessageEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkSubmitDebugUtilsMessageEXT.html
type PfnSubmitDebugUtilsMessageEXT uintptr

func (fn PfnSubmitDebugUtilsMessageEXT) Call(instance Instance, messageSeverity DebugUtilsMessageSeverityFlagsEXT, messageTypes DebugUtilsMessageTypeFlagsEXT, pCallbackData *DebugUtilsMessengerCallbackDataEXT) {
	_, _, _ = call(uintptr(fn), uintptr(instance), uintptr(messageSeverity), uintptr(messageTypes), uintptr(unsafe.Pointer(pCallbackData)))
}
func (fn PfnSubmitDebugUtilsMessageEXT) String() string { return "vkSubmitDebugUtilsMessageEXT" }

const EXT_sampler_filter_minmax = 1
const EXT_SAMPLER_FILTER_MINMAX_SPEC_VERSION = 2

var EXT_SAMPLER_FILTER_MINMAX_EXTENSION_NAME = "VK_EXT_sampler_filter_minmax"

// SamplerReductionModeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerReductionModeEXT.html
type SamplerReductionModeEXT int32

const (
	SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT SamplerReductionModeEXT = 0
	SAMPLER_REDUCTION_MODE_MIN_EXT              SamplerReductionModeEXT = 1
	SAMPLER_REDUCTION_MODE_MAX_EXT              SamplerReductionModeEXT = 2
	SAMPLER_REDUCTION_MODE_BEGIN_RANGE_EXT      SamplerReductionModeEXT = SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT
	SAMPLER_REDUCTION_MODE_END_RANGE_EXT        SamplerReductionModeEXT = SAMPLER_REDUCTION_MODE_MAX_EXT
	SAMPLER_REDUCTION_MODE_RANGE_SIZE_EXT       SamplerReductionModeEXT = (SAMPLER_REDUCTION_MODE_MAX_EXT - SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT + 1)
	SAMPLER_REDUCTION_MODE_MAX_ENUM_EXT         SamplerReductionModeEXT = 0x7FFFFFFF
)

func (x SamplerReductionModeEXT) String() string {
	switch x {
	case SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT:
		return "SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT"
	case SAMPLER_REDUCTION_MODE_MIN_EXT:
		return "SAMPLER_REDUCTION_MODE_MIN_EXT"
	case SAMPLER_REDUCTION_MODE_MAX_EXT:
		return "SAMPLER_REDUCTION_MODE_MAX_EXT"
	case SAMPLER_REDUCTION_MODE_MAX_ENUM_EXT:
		return "SAMPLER_REDUCTION_MODE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SamplerReductionModeCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSamplerReductionModeCreateInfoEXT.html
type SamplerReductionModeCreateInfoEXT struct {
	SType         StructureType
	PNext         unsafe.Pointer
	ReductionMode SamplerReductionModeEXT
}

func NewSamplerReductionModeCreateInfoEXT() *SamplerReductionModeCreateInfoEXT {
	p := (*SamplerReductionModeCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*SamplerReductionModeCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT
	return p
}
func (p *SamplerReductionModeCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceSamplerFilterMinmaxPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT.html
type PhysicalDeviceSamplerFilterMinmaxPropertiesEXT struct {
	SType                              StructureType
	PNext                              unsafe.Pointer
	FilterMinmaxSingleComponentFormats Bool32
	FilterMinmaxImageComponentMapping  Bool32
}

func NewPhysicalDeviceSamplerFilterMinmaxPropertiesEXT() *PhysicalDeviceSamplerFilterMinmaxPropertiesEXT {
	p := (*PhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceSamplerFilterMinmaxPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const AMD_gpu_shader_int16 = 1
const AMD_GPU_SHADER_INT16_SPEC_VERSION = 2

var AMD_GPU_SHADER_INT16_EXTENSION_NAME = "VK_AMD_gpu_shader_int16"

const AMD_mixed_attachment_samples = 1
const AMD_MIXED_ATTACHMENT_SAMPLES_SPEC_VERSION = 1

var AMD_MIXED_ATTACHMENT_SAMPLES_EXTENSION_NAME = "VK_AMD_mixed_attachment_samples"

const AMD_shader_fragment_mask = 1
const AMD_SHADER_FRAGMENT_MASK_SPEC_VERSION = 1

var AMD_SHADER_FRAGMENT_MASK_EXTENSION_NAME = "VK_AMD_shader_fragment_mask"

const EXT_inline_uniform_block = 1
const EXT_INLINE_UNIFORM_BLOCK_SPEC_VERSION = 1

var EXT_INLINE_UNIFORM_BLOCK_EXTENSION_NAME = "VK_EXT_inline_uniform_block"

// PhysicalDeviceInlineUniformBlockFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceInlineUniformBlockFeaturesEXT.html
type PhysicalDeviceInlineUniformBlockFeaturesEXT struct {
	SType                                              StructureType
	PNext                                              unsafe.Pointer
	InlineUniformBlock                                 Bool32
	DescriptorBindingInlineUniformBlockUpdateAfterBind Bool32
}

func NewPhysicalDeviceInlineUniformBlockFeaturesEXT() *PhysicalDeviceInlineUniformBlockFeaturesEXT {
	p := (*PhysicalDeviceInlineUniformBlockFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceInlineUniformBlockFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceInlineUniformBlockFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceInlineUniformBlockPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceInlineUniformBlockPropertiesEXT.html
type PhysicalDeviceInlineUniformBlockPropertiesEXT struct {
	SType                                                   StructureType
	PNext                                                   unsafe.Pointer
	MaxInlineUniformBlockSize                               uint32
	MaxPerStageDescriptorInlineUniformBlocks                uint32
	MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks uint32
	MaxDescriptorSetInlineUniformBlocks                     uint32
	MaxDescriptorSetUpdateAfterBindInlineUniformBlocks      uint32
}

func NewPhysicalDeviceInlineUniformBlockPropertiesEXT() *PhysicalDeviceInlineUniformBlockPropertiesEXT {
	p := (*PhysicalDeviceInlineUniformBlockPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceInlineUniformBlockPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceInlineUniformBlockPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// WriteDescriptorSetInlineUniformBlockEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkWriteDescriptorSetInlineUniformBlockEXT.html
type WriteDescriptorSetInlineUniformBlockEXT struct {
	SType    StructureType
	PNext    unsafe.Pointer
	DataSize uint32
	PData    unsafe.Pointer
}

func NewWriteDescriptorSetInlineUniformBlockEXT() *WriteDescriptorSetInlineUniformBlockEXT {
	p := (*WriteDescriptorSetInlineUniformBlockEXT)(MemAlloc(unsafe.Sizeof(*(*WriteDescriptorSetInlineUniformBlockEXT)(nil))))
	p.SType = STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT
	return p
}
func (p *WriteDescriptorSetInlineUniformBlockEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorPoolInlineUniformBlockCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorPoolInlineUniformBlockCreateInfoEXT.html
type DescriptorPoolInlineUniformBlockCreateInfoEXT struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	MaxInlineUniformBlockBindings uint32
}

func NewDescriptorPoolInlineUniformBlockCreateInfoEXT() *DescriptorPoolInlineUniformBlockCreateInfoEXT {
	p := (*DescriptorPoolInlineUniformBlockCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DescriptorPoolInlineUniformBlockCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT
	return p
}
func (p *DescriptorPoolInlineUniformBlockCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_shader_stencil_export = 1
const EXT_SHADER_STENCIL_EXPORT_SPEC_VERSION = 1

var EXT_SHADER_STENCIL_EXPORT_EXTENSION_NAME = "VK_EXT_shader_stencil_export"

const EXT_sample_locations = 1
const EXT_SAMPLE_LOCATIONS_SPEC_VERSION = 1

var EXT_SAMPLE_LOCATIONS_EXTENSION_NAME = "VK_EXT_sample_locations"

// SampleLocationEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSampleLocationEXT.html
type SampleLocationEXT struct {
	X float32
	Y float32
}

func NewSampleLocationEXT() *SampleLocationEXT {
	return (*SampleLocationEXT)(MemAlloc(unsafe.Sizeof(*(*SampleLocationEXT)(nil))))
}
func (p *SampleLocationEXT) Free() { MemFree(unsafe.Pointer(p)) }

// SampleLocationsInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSampleLocationsInfoEXT.html
type SampleLocationsInfoEXT struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	SampleLocationsPerPixel SampleCountFlags
	SampleLocationGridSize  Extent2D
	SampleLocationsCount    uint32
	PSampleLocations        *SampleLocationEXT
}

func NewSampleLocationsInfoEXT() *SampleLocationsInfoEXT {
	p := (*SampleLocationsInfoEXT)(MemAlloc(unsafe.Sizeof(*(*SampleLocationsInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT
	return p
}
func (p *SampleLocationsInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// AttachmentSampleLocationsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAttachmentSampleLocationsEXT.html
type AttachmentSampleLocationsEXT struct {
	AttachmentIndex     uint32
	SampleLocationsInfo SampleLocationsInfoEXT
}

func NewAttachmentSampleLocationsEXT() *AttachmentSampleLocationsEXT {
	return (*AttachmentSampleLocationsEXT)(MemAlloc(unsafe.Sizeof(*(*AttachmentSampleLocationsEXT)(nil))))
}
func (p *AttachmentSampleLocationsEXT) Free() { MemFree(unsafe.Pointer(p)) }

// SubpassSampleLocationsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSubpassSampleLocationsEXT.html
type SubpassSampleLocationsEXT struct {
	SubpassIndex        uint32
	SampleLocationsInfo SampleLocationsInfoEXT
}

func NewSubpassSampleLocationsEXT() *SubpassSampleLocationsEXT {
	return (*SubpassSampleLocationsEXT)(MemAlloc(unsafe.Sizeof(*(*SubpassSampleLocationsEXT)(nil))))
}
func (p *SubpassSampleLocationsEXT) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassSampleLocationsBeginInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html
type RenderPassSampleLocationsBeginInfoEXT struct {
	SType                                 StructureType
	PNext                                 unsafe.Pointer
	AttachmentInitialSampleLocationsCount uint32
	PAttachmentInitialSampleLocations     *AttachmentSampleLocationsEXT
	PostSubpassSampleLocationsCount       uint32
	PPostSubpassSampleLocations           *SubpassSampleLocationsEXT
}

func NewRenderPassSampleLocationsBeginInfoEXT() *RenderPassSampleLocationsBeginInfoEXT {
	p := (*RenderPassSampleLocationsBeginInfoEXT)(MemAlloc(unsafe.Sizeof(*(*RenderPassSampleLocationsBeginInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT
	return p
}
func (p *RenderPassSampleLocationsBeginInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineSampleLocationsStateCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html
type PipelineSampleLocationsStateCreateInfoEXT struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	SampleLocationsEnable Bool32
	SampleLocationsInfo   SampleLocationsInfoEXT
}

func NewPipelineSampleLocationsStateCreateInfoEXT() *PipelineSampleLocationsStateCreateInfoEXT {
	p := (*PipelineSampleLocationsStateCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineSampleLocationsStateCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT
	return p
}
func (p *PipelineSampleLocationsStateCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceSampleLocationsPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html
type PhysicalDeviceSampleLocationsPropertiesEXT struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	SampleLocationSampleCounts    SampleCountFlags
	MaxSampleLocationGridSize     Extent2D
	SampleLocationCoordinateRange [2]float32
	SampleLocationSubPixelBits    uint32
	VariableSampleLocations       Bool32
}

func NewPhysicalDeviceSampleLocationsPropertiesEXT() *PhysicalDeviceSampleLocationsPropertiesEXT {
	p := (*PhysicalDeviceSampleLocationsPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSampleLocationsPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceSampleLocationsPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// MultisamplePropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMultisamplePropertiesEXT.html
type MultisamplePropertiesEXT struct {
	SType                     StructureType
	PNext                     unsafe.Pointer
	MaxSampleLocationGridSize Extent2D
}

func NewMultisamplePropertiesEXT() *MultisamplePropertiesEXT {
	p := (*MultisamplePropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*MultisamplePropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT
	return p
}
func (p *MultisamplePropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdSetSampleLocationsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetSampleLocationsEXT.html
type PfnCmdSetSampleLocationsEXT uintptr

func (fn PfnCmdSetSampleLocationsEXT) Call(commandBuffer CommandBuffer, pSampleLocationsInfo *SampleLocationsInfoEXT) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pSampleLocationsInfo)))
}
func (fn PfnCmdSetSampleLocationsEXT) String() string { return "vkCmdSetSampleLocationsEXT" }

//  PfnGetPhysicalDeviceMultisamplePropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html
type PfnGetPhysicalDeviceMultisamplePropertiesEXT uintptr

func (fn PfnGetPhysicalDeviceMultisamplePropertiesEXT) Call(physicalDevice PhysicalDevice, samples SampleCountFlags, pMultisampleProperties *MultisamplePropertiesEXT) {
	_, _, _ = call(uintptr(fn), uintptr(physicalDevice), uintptr(samples), uintptr(unsafe.Pointer(pMultisampleProperties)))
}
func (fn PfnGetPhysicalDeviceMultisamplePropertiesEXT) String() string {
	return "vkGetPhysicalDeviceMultisamplePropertiesEXT"
}

const EXT_blend_operation_advanced = 1
const EXT_BLEND_OPERATION_ADVANCED_SPEC_VERSION = 2

var EXT_BLEND_OPERATION_ADVANCED_EXTENSION_NAME = "VK_EXT_blend_operation_advanced"

// BlendOverlapEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBlendOverlapEXT.html
type BlendOverlapEXT int32

const (
	BLEND_OVERLAP_UNCORRELATED_EXT BlendOverlapEXT = 0
	BLEND_OVERLAP_DISJOINT_EXT     BlendOverlapEXT = 1
	BLEND_OVERLAP_CONJOINT_EXT     BlendOverlapEXT = 2
	BLEND_OVERLAP_BEGIN_RANGE_EXT  BlendOverlapEXT = BLEND_OVERLAP_UNCORRELATED_EXT
	BLEND_OVERLAP_END_RANGE_EXT    BlendOverlapEXT = BLEND_OVERLAP_CONJOINT_EXT
	BLEND_OVERLAP_RANGE_SIZE_EXT   BlendOverlapEXT = (BLEND_OVERLAP_CONJOINT_EXT - BLEND_OVERLAP_UNCORRELATED_EXT + 1)
	BLEND_OVERLAP_MAX_ENUM_EXT     BlendOverlapEXT = 0x7FFFFFFF
)

func (x BlendOverlapEXT) String() string {
	switch x {
	case BLEND_OVERLAP_UNCORRELATED_EXT:
		return "BLEND_OVERLAP_UNCORRELATED_EXT"
	case BLEND_OVERLAP_DISJOINT_EXT:
		return "BLEND_OVERLAP_DISJOINT_EXT"
	case BLEND_OVERLAP_CONJOINT_EXT:
		return "BLEND_OVERLAP_CONJOINT_EXT"
	case BLEND_OVERLAP_MAX_ENUM_EXT:
		return "BLEND_OVERLAP_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PhysicalDeviceBlendOperationAdvancedFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT.html
type PhysicalDeviceBlendOperationAdvancedFeaturesEXT struct {
	SType                           StructureType
	PNext                           unsafe.Pointer
	AdvancedBlendCoherentOperations Bool32
}

func NewPhysicalDeviceBlendOperationAdvancedFeaturesEXT() *PhysicalDeviceBlendOperationAdvancedFeaturesEXT {
	p := (*PhysicalDeviceBlendOperationAdvancedFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceBlendOperationAdvancedFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceBlendOperationAdvancedFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceBlendOperationAdvancedPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html
type PhysicalDeviceBlendOperationAdvancedPropertiesEXT struct {
	SType                                 StructureType
	PNext                                 unsafe.Pointer
	AdvancedBlendMaxColorAttachments      uint32
	AdvancedBlendIndependentBlend         Bool32
	AdvancedBlendNonPremultipliedSrcColor Bool32
	AdvancedBlendNonPremultipliedDstColor Bool32
	AdvancedBlendCorrelatedOverlap        Bool32
	AdvancedBlendAllOperations            Bool32
}

func NewPhysicalDeviceBlendOperationAdvancedPropertiesEXT() *PhysicalDeviceBlendOperationAdvancedPropertiesEXT {
	p := (*PhysicalDeviceBlendOperationAdvancedPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceBlendOperationAdvancedPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceBlendOperationAdvancedPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineColorBlendAdvancedStateCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html
type PipelineColorBlendAdvancedStateCreateInfoEXT struct {
	SType            StructureType
	PNext            unsafe.Pointer
	SrcPremultiplied Bool32
	DstPremultiplied Bool32
	BlendOverlap     BlendOverlapEXT
}

func NewPipelineColorBlendAdvancedStateCreateInfoEXT() *PipelineColorBlendAdvancedStateCreateInfoEXT {
	p := (*PipelineColorBlendAdvancedStateCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineColorBlendAdvancedStateCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT
	return p
}
func (p *PipelineColorBlendAdvancedStateCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const NV_fragment_coverage_to_color = 1
const NV_FRAGMENT_COVERAGE_TO_COLOR_SPEC_VERSION = 1

var NV_FRAGMENT_COVERAGE_TO_COLOR_EXTENSION_NAME = "VK_NV_fragment_coverage_to_color"

type PipelineCoverageToColorStateCreateFlagsNV uint32 // reserved
// PipelineCoverageToColorStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html
type PipelineCoverageToColorStateCreateInfoNV struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	Flags                   PipelineCoverageToColorStateCreateFlagsNV
	CoverageToColorEnable   Bool32
	CoverageToColorLocation uint32
}

func NewPipelineCoverageToColorStateCreateInfoNV() *PipelineCoverageToColorStateCreateInfoNV {
	p := (*PipelineCoverageToColorStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineCoverageToColorStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineCoverageToColorStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

const NV_framebuffer_mixed_samples = 1
const NV_FRAMEBUFFER_MIXED_SAMPLES_SPEC_VERSION = 1

var NV_FRAMEBUFFER_MIXED_SAMPLES_EXTENSION_NAME = "VK_NV_framebuffer_mixed_samples"

// CoverageModulationModeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCoverageModulationModeNV.html
type CoverageModulationModeNV int32

const (
	COVERAGE_MODULATION_MODE_NONE_NV        CoverageModulationModeNV = 0
	COVERAGE_MODULATION_MODE_RGB_NV         CoverageModulationModeNV = 1
	COVERAGE_MODULATION_MODE_ALPHA_NV       CoverageModulationModeNV = 2
	COVERAGE_MODULATION_MODE_RGBA_NV        CoverageModulationModeNV = 3
	COVERAGE_MODULATION_MODE_BEGIN_RANGE_NV CoverageModulationModeNV = COVERAGE_MODULATION_MODE_NONE_NV
	COVERAGE_MODULATION_MODE_END_RANGE_NV   CoverageModulationModeNV = COVERAGE_MODULATION_MODE_RGBA_NV
	COVERAGE_MODULATION_MODE_RANGE_SIZE_NV  CoverageModulationModeNV = (COVERAGE_MODULATION_MODE_RGBA_NV - COVERAGE_MODULATION_MODE_NONE_NV + 1)
	COVERAGE_MODULATION_MODE_MAX_ENUM_NV    CoverageModulationModeNV = 0x7FFFFFFF
)

func (x CoverageModulationModeNV) String() string {
	switch x {
	case COVERAGE_MODULATION_MODE_NONE_NV:
		return "COVERAGE_MODULATION_MODE_NONE_NV"
	case COVERAGE_MODULATION_MODE_RGB_NV:
		return "COVERAGE_MODULATION_MODE_RGB_NV"
	case COVERAGE_MODULATION_MODE_ALPHA_NV:
		return "COVERAGE_MODULATION_MODE_ALPHA_NV"
	case COVERAGE_MODULATION_MODE_RGBA_NV:
		return "COVERAGE_MODULATION_MODE_RGBA_NV"
	case COVERAGE_MODULATION_MODE_MAX_ENUM_NV:
		return "COVERAGE_MODULATION_MODE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

type PipelineCoverageModulationStateCreateFlagsNV uint32 // reserved
// PipelineCoverageModulationStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html
type PipelineCoverageModulationStateCreateInfoNV struct {
	SType                         StructureType
	PNext                         unsafe.Pointer
	Flags                         PipelineCoverageModulationStateCreateFlagsNV
	CoverageModulationMode        CoverageModulationModeNV
	CoverageModulationTableEnable Bool32
	CoverageModulationTableCount  uint32
	PCoverageModulationTable      *float32
}

func NewPipelineCoverageModulationStateCreateInfoNV() *PipelineCoverageModulationStateCreateInfoNV {
	p := (*PipelineCoverageModulationStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineCoverageModulationStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineCoverageModulationStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

const NV_fill_rectangle = 1
const NV_FILL_RECTANGLE_SPEC_VERSION = 1

var NV_FILL_RECTANGLE_EXTENSION_NAME = "VK_NV_fill_rectangle"

const NV_shader_sm_builtins = 1
const NV_SHADER_SM_BUILTINS_SPEC_VERSION = 1

var NV_SHADER_SM_BUILTINS_EXTENSION_NAME = "VK_NV_shader_sm_builtins"

// PhysicalDeviceShaderSMBuiltinsPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderSMBuiltinsPropertiesNV.html
type PhysicalDeviceShaderSMBuiltinsPropertiesNV struct {
	SType            StructureType
	PNext            unsafe.Pointer
	ShaderSMCount    uint32
	ShaderWarpsPerSM uint32
}

func NewPhysicalDeviceShaderSMBuiltinsPropertiesNV() *PhysicalDeviceShaderSMBuiltinsPropertiesNV {
	p := (*PhysicalDeviceShaderSMBuiltinsPropertiesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderSMBuiltinsPropertiesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV
	return p
}
func (p *PhysicalDeviceShaderSMBuiltinsPropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceShaderSMBuiltinsFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderSMBuiltinsFeaturesNV.html
type PhysicalDeviceShaderSMBuiltinsFeaturesNV struct {
	SType            StructureType
	PNext            unsafe.Pointer
	ShaderSMBuiltins Bool32
}

func NewPhysicalDeviceShaderSMBuiltinsFeaturesNV() *PhysicalDeviceShaderSMBuiltinsFeaturesNV {
	p := (*PhysicalDeviceShaderSMBuiltinsFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderSMBuiltinsFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_FEATURES_NV
	return p
}
func (p *PhysicalDeviceShaderSMBuiltinsFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_post_depth_coverage = 1
const EXT_POST_DEPTH_COVERAGE_SPEC_VERSION = 1

var EXT_POST_DEPTH_COVERAGE_EXTENSION_NAME = "VK_EXT_post_depth_coverage"

const EXT_image_drm_format_modifier = 1
const EXT_IMAGE_DRM_FORMAT_MODIFIER_SPEC_VERSION = 1

var EXT_IMAGE_DRM_FORMAT_MODIFIER_EXTENSION_NAME = "VK_EXT_image_drm_format_modifier"

// DrmFormatModifierPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html
type DrmFormatModifierPropertiesEXT struct {
	DrmFormatModifier               uint64
	DrmFormatModifierPlaneCount     uint32
	DrmFormatModifierTilingFeatures FormatFeatureFlags
}

func NewDrmFormatModifierPropertiesEXT() *DrmFormatModifierPropertiesEXT {
	return (*DrmFormatModifierPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*DrmFormatModifierPropertiesEXT)(nil))))
}
func (p *DrmFormatModifierPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DrmFormatModifierPropertiesListEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html
type DrmFormatModifierPropertiesListEXT struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	DrmFormatModifierCount       uint32
	PDrmFormatModifierProperties *DrmFormatModifierPropertiesEXT
}

func NewDrmFormatModifierPropertiesListEXT() *DrmFormatModifierPropertiesListEXT {
	p := (*DrmFormatModifierPropertiesListEXT)(MemAlloc(unsafe.Sizeof(*(*DrmFormatModifierPropertiesListEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_EXT
	return p
}
func (p *DrmFormatModifierPropertiesListEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceImageDrmFormatModifierInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html
type PhysicalDeviceImageDrmFormatModifierInfoEXT struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	DrmFormatModifier     uint64
	SharingMode           SharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}

func NewPhysicalDeviceImageDrmFormatModifierInfoEXT() *PhysicalDeviceImageDrmFormatModifierInfoEXT {
	p := (*PhysicalDeviceImageDrmFormatModifierInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceImageDrmFormatModifierInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT
	return p
}
func (p *PhysicalDeviceImageDrmFormatModifierInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// ImageDrmFormatModifierListCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html
type ImageDrmFormatModifierListCreateInfoEXT struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	DrmFormatModifierCount uint32
	PDrmFormatModifiers    *uint64
}

func NewImageDrmFormatModifierListCreateInfoEXT() *ImageDrmFormatModifierListCreateInfoEXT {
	p := (*ImageDrmFormatModifierListCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*ImageDrmFormatModifierListCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT
	return p
}
func (p *ImageDrmFormatModifierListCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// ImageDrmFormatModifierExplicitCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html
type ImageDrmFormatModifierExplicitCreateInfoEXT struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	DrmFormatModifier           uint64
	DrmFormatModifierPlaneCount uint32
	PPlaneLayouts               *SubresourceLayout
}

func NewImageDrmFormatModifierExplicitCreateInfoEXT() *ImageDrmFormatModifierExplicitCreateInfoEXT {
	p := (*ImageDrmFormatModifierExplicitCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*ImageDrmFormatModifierExplicitCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT
	return p
}
func (p *ImageDrmFormatModifierExplicitCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// ImageDrmFormatModifierPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html
type ImageDrmFormatModifierPropertiesEXT struct {
	SType             StructureType
	PNext             unsafe.Pointer
	DrmFormatModifier uint64
}

func NewImageDrmFormatModifierPropertiesEXT() *ImageDrmFormatModifierPropertiesEXT {
	p := (*ImageDrmFormatModifierPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*ImageDrmFormatModifierPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT
	return p
}
func (p *ImageDrmFormatModifierPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetImageDrmFormatModifierPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetImageDrmFormatModifierPropertiesEXT.html
type PfnGetImageDrmFormatModifierPropertiesEXT uintptr

func (fn PfnGetImageDrmFormatModifierPropertiesEXT) Call(device Device, image Image, pProperties *ImageDrmFormatModifierPropertiesEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(image), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetImageDrmFormatModifierPropertiesEXT) String() string {
	return "vkGetImageDrmFormatModifierPropertiesEXT"
}

const EXT_validation_cache = 1

// ValidationCacheEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationCacheEXT.html
type ValidationCacheEXT NonDispatchableHandle

const EXT_VALIDATION_CACHE_SPEC_VERSION = 1

var EXT_VALIDATION_CACHE_EXTENSION_NAME = "VK_EXT_validation_cache"

// ValidationCacheHeaderVersionEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationCacheHeaderVersionEXT.html
type ValidationCacheHeaderVersionEXT int32

const (
	VALIDATION_CACHE_HEADER_VERSION_ONE_EXT         ValidationCacheHeaderVersionEXT = 1
	VALIDATION_CACHE_HEADER_VERSION_BEGIN_RANGE_EXT ValidationCacheHeaderVersionEXT = VALIDATION_CACHE_HEADER_VERSION_ONE_EXT
	VALIDATION_CACHE_HEADER_VERSION_END_RANGE_EXT   ValidationCacheHeaderVersionEXT = VALIDATION_CACHE_HEADER_VERSION_ONE_EXT
	VALIDATION_CACHE_HEADER_VERSION_RANGE_SIZE_EXT  ValidationCacheHeaderVersionEXT = (VALIDATION_CACHE_HEADER_VERSION_ONE_EXT - VALIDATION_CACHE_HEADER_VERSION_ONE_EXT + 1)
	VALIDATION_CACHE_HEADER_VERSION_MAX_ENUM_EXT    ValidationCacheHeaderVersionEXT = 0x7FFFFFFF
)

func (x ValidationCacheHeaderVersionEXT) String() string {
	switch x {
	case VALIDATION_CACHE_HEADER_VERSION_ONE_EXT:
		return "VALIDATION_CACHE_HEADER_VERSION_ONE_EXT"
	case VALIDATION_CACHE_HEADER_VERSION_MAX_ENUM_EXT:
		return "VALIDATION_CACHE_HEADER_VERSION_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

type ValidationCacheCreateFlagsEXT uint32 // reserved
// ValidationCacheCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationCacheCreateInfoEXT.html
type ValidationCacheCreateInfoEXT struct {
	SType           StructureType
	PNext           unsafe.Pointer
	Flags           ValidationCacheCreateFlagsEXT
	InitialDataSize uintptr
	PInitialData    unsafe.Pointer
}

func NewValidationCacheCreateInfoEXT() *ValidationCacheCreateInfoEXT {
	p := (*ValidationCacheCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*ValidationCacheCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT
	return p
}
func (p *ValidationCacheCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// ShaderModuleValidationCacheCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html
type ShaderModuleValidationCacheCreateInfoEXT struct {
	SType           StructureType
	PNext           unsafe.Pointer
	ValidationCache ValidationCacheEXT
}

func NewShaderModuleValidationCacheCreateInfoEXT() *ShaderModuleValidationCacheCreateInfoEXT {
	p := (*ShaderModuleValidationCacheCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*ShaderModuleValidationCacheCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT
	return p
}
func (p *ShaderModuleValidationCacheCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateValidationCacheEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateValidationCacheEXT.html
type PfnCreateValidationCacheEXT uintptr

func (fn PfnCreateValidationCacheEXT) Call(device Device, pCreateInfo *ValidationCacheCreateInfoEXT, pAllocator *AllocationCallbacks, pValidationCache *ValidationCacheEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pValidationCache)))
	return Result(ret)
}
func (fn PfnCreateValidationCacheEXT) String() string { return "vkCreateValidationCacheEXT" }

//  PfnDestroyValidationCacheEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyValidationCacheEXT.html
type PfnDestroyValidationCacheEXT uintptr

func (fn PfnDestroyValidationCacheEXT) Call(device Device, validationCache ValidationCacheEXT, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(validationCache), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyValidationCacheEXT) String() string { return "vkDestroyValidationCacheEXT" }

//  PfnMergeValidationCachesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkMergeValidationCachesEXT.html
type PfnMergeValidationCachesEXT uintptr

func (fn PfnMergeValidationCachesEXT) Call(device Device, dstCache ValidationCacheEXT, srcCacheCount uint32, pSrcCaches *ValidationCacheEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(dstCache), uintptr(srcCacheCount), uintptr(unsafe.Pointer(pSrcCaches)))
	return Result(ret)
}
func (fn PfnMergeValidationCachesEXT) String() string { return "vkMergeValidationCachesEXT" }

//  PfnGetValidationCacheDataEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetValidationCacheDataEXT.html
type PfnGetValidationCacheDataEXT uintptr

func (fn PfnGetValidationCacheDataEXT) Call(device Device, validationCache ValidationCacheEXT, pDataSize *uintptr, pData unsafe.Pointer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(validationCache), uintptr(unsafe.Pointer(pDataSize)), uintptr(pData))
	return Result(ret)
}
func (fn PfnGetValidationCacheDataEXT) String() string { return "vkGetValidationCacheDataEXT" }

const EXT_descriptor_indexing = 1
const EXT_DESCRIPTOR_INDEXING_SPEC_VERSION = 2

var EXT_DESCRIPTOR_INDEXING_EXTENSION_NAME = "VK_EXT_descriptor_indexing"

// DescriptorBindingFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorBindingFlagsEXT.html
type DescriptorBindingFlagsEXT uint32

const (
	DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT           DescriptorBindingFlagsEXT = 0x00000001
	DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT DescriptorBindingFlagsEXT = 0x00000002
	DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT             DescriptorBindingFlagsEXT = 0x00000004
	DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT   DescriptorBindingFlagsEXT = 0x00000008
	DESCRIPTOR_BINDING_FLAG_BITS_MAX_ENUM_EXT              DescriptorBindingFlagsEXT = 0x7FFFFFFF
)

func (x DescriptorBindingFlagsEXT) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch DescriptorBindingFlagsEXT(1 << i) {
			case DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT:
				s += "DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT|"
			case DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT:
				s += "DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT|"
			case DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT:
				s += "DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT|"
			case DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT:
				s += "DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// DescriptorSetLayoutBindingFlagsCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfoEXT.html
type DescriptorSetLayoutBindingFlagsCreateInfoEXT struct {
	SType         StructureType
	PNext         unsafe.Pointer
	BindingCount  uint32
	PBindingFlags *DescriptorBindingFlagsEXT
}

func NewDescriptorSetLayoutBindingFlagsCreateInfoEXT() *DescriptorSetLayoutBindingFlagsCreateInfoEXT {
	p := (*DescriptorSetLayoutBindingFlagsCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DescriptorSetLayoutBindingFlagsCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT
	return p
}
func (p *DescriptorSetLayoutBindingFlagsCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceDescriptorIndexingFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeaturesEXT.html
type PhysicalDeviceDescriptorIndexingFeaturesEXT struct {
	SType                                              StructureType
	PNext                                              unsafe.Pointer
	ShaderInputAttachmentArrayDynamicIndexing          Bool32
	ShaderUniformTexelBufferArrayDynamicIndexing       Bool32
	ShaderStorageTexelBufferArrayDynamicIndexing       Bool32
	ShaderUniformBufferArrayNonUniformIndexing         Bool32
	ShaderSampledImageArrayNonUniformIndexing          Bool32
	ShaderStorageBufferArrayNonUniformIndexing         Bool32
	ShaderStorageImageArrayNonUniformIndexing          Bool32
	ShaderInputAttachmentArrayNonUniformIndexing       Bool32
	ShaderUniformTexelBufferArrayNonUniformIndexing    Bool32
	ShaderStorageTexelBufferArrayNonUniformIndexing    Bool32
	DescriptorBindingUniformBufferUpdateAfterBind      Bool32
	DescriptorBindingSampledImageUpdateAfterBind       Bool32
	DescriptorBindingStorageImageUpdateAfterBind       Bool32
	DescriptorBindingStorageBufferUpdateAfterBind      Bool32
	DescriptorBindingUniformTexelBufferUpdateAfterBind Bool32
	DescriptorBindingStorageTexelBufferUpdateAfterBind Bool32
	DescriptorBindingUpdateUnusedWhilePending          Bool32
	DescriptorBindingPartiallyBound                    Bool32
	DescriptorBindingVariableDescriptorCount           Bool32
	RuntimeDescriptorArray                             Bool32
}

func NewPhysicalDeviceDescriptorIndexingFeaturesEXT() *PhysicalDeviceDescriptorIndexingFeaturesEXT {
	p := (*PhysicalDeviceDescriptorIndexingFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceDescriptorIndexingFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceDescriptorIndexingFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceDescriptorIndexingPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceDescriptorIndexingPropertiesEXT.html
type PhysicalDeviceDescriptorIndexingPropertiesEXT struct {
	SType                                                StructureType
	PNext                                                unsafe.Pointer
	MaxUpdateAfterBindDescriptorsInAllPools              uint32
	ShaderUniformBufferArrayNonUniformIndexingNative     Bool32
	ShaderSampledImageArrayNonUniformIndexingNative      Bool32
	ShaderStorageBufferArrayNonUniformIndexingNative     Bool32
	ShaderStorageImageArrayNonUniformIndexingNative      Bool32
	ShaderInputAttachmentArrayNonUniformIndexingNative   Bool32
	RobustBufferAccessUpdateAfterBind                    Bool32
	QuadDivergentImplicitLod                             Bool32
	MaxPerStageDescriptorUpdateAfterBindSamplers         uint32
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindSampledImages    uint32
	MaxPerStageDescriptorUpdateAfterBindStorageImages    uint32
	MaxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	MaxPerStageUpdateAfterBindResources                  uint32
	MaxDescriptorSetUpdateAfterBindSamplers              uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffers        uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffers        uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindSampledImages         uint32
	MaxDescriptorSetUpdateAfterBindStorageImages         uint32
	MaxDescriptorSetUpdateAfterBindInputAttachments      uint32
}

func NewPhysicalDeviceDescriptorIndexingPropertiesEXT() *PhysicalDeviceDescriptorIndexingPropertiesEXT {
	p := (*PhysicalDeviceDescriptorIndexingPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceDescriptorIndexingPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceDescriptorIndexingPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorSetVariableDescriptorCountAllocateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfoEXT.html
type DescriptorSetVariableDescriptorCountAllocateInfoEXT struct {
	SType              StructureType
	PNext              unsafe.Pointer
	DescriptorSetCount uint32
	PDescriptorCounts  *uint32
}

func NewDescriptorSetVariableDescriptorCountAllocateInfoEXT() *DescriptorSetVariableDescriptorCountAllocateInfoEXT {
	p := (*DescriptorSetVariableDescriptorCountAllocateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DescriptorSetVariableDescriptorCountAllocateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT
	return p
}
func (p *DescriptorSetVariableDescriptorCountAllocateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// DescriptorSetVariableDescriptorCountLayoutSupportEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupportEXT.html
type DescriptorSetVariableDescriptorCountLayoutSupportEXT struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	MaxVariableDescriptorCount uint32
}

func NewDescriptorSetVariableDescriptorCountLayoutSupportEXT() *DescriptorSetVariableDescriptorCountLayoutSupportEXT {
	p := (*DescriptorSetVariableDescriptorCountLayoutSupportEXT)(MemAlloc(unsafe.Sizeof(*(*DescriptorSetVariableDescriptorCountLayoutSupportEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT
	return p
}
func (p *DescriptorSetVariableDescriptorCountLayoutSupportEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_shader_viewport_index_layer = 1
const EXT_SHADER_VIEWPORT_INDEX_LAYER_SPEC_VERSION = 1

var EXT_SHADER_VIEWPORT_INDEX_LAYER_EXTENSION_NAME = "VK_EXT_shader_viewport_index_layer"

const NV_shading_rate_image = 1
const NV_SHADING_RATE_IMAGE_SPEC_VERSION = 3

var NV_SHADING_RATE_IMAGE_EXTENSION_NAME = "VK_NV_shading_rate_image"

// ShadingRatePaletteEntryNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShadingRatePaletteEntryNV.html
type ShadingRatePaletteEntryNV int32

const (
	SHADING_RATE_PALETTE_ENTRY_NO_INVOCATIONS_NV              ShadingRatePaletteEntryNV = 0
	SHADING_RATE_PALETTE_ENTRY_16_INVOCATIONS_PER_PIXEL_NV    ShadingRatePaletteEntryNV = 1
	SHADING_RATE_PALETTE_ENTRY_8_INVOCATIONS_PER_PIXEL_NV     ShadingRatePaletteEntryNV = 2
	SHADING_RATE_PALETTE_ENTRY_4_INVOCATIONS_PER_PIXEL_NV     ShadingRatePaletteEntryNV = 3
	SHADING_RATE_PALETTE_ENTRY_2_INVOCATIONS_PER_PIXEL_NV     ShadingRatePaletteEntryNV = 4
	SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_PIXEL_NV      ShadingRatePaletteEntryNV = 5
	SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X1_PIXELS_NV ShadingRatePaletteEntryNV = 6
	SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_1X2_PIXELS_NV ShadingRatePaletteEntryNV = 7
	SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X2_PIXELS_NV ShadingRatePaletteEntryNV = 8
	SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X2_PIXELS_NV ShadingRatePaletteEntryNV = 9
	SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X4_PIXELS_NV ShadingRatePaletteEntryNV = 10
	SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X4_PIXELS_NV ShadingRatePaletteEntryNV = 11
	SHADING_RATE_PALETTE_ENTRY_BEGIN_RANGE_NV                 ShadingRatePaletteEntryNV = SHADING_RATE_PALETTE_ENTRY_NO_INVOCATIONS_NV
	SHADING_RATE_PALETTE_ENTRY_END_RANGE_NV                   ShadingRatePaletteEntryNV = SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X4_PIXELS_NV
	SHADING_RATE_PALETTE_ENTRY_RANGE_SIZE_NV                  ShadingRatePaletteEntryNV = (SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X4_PIXELS_NV - SHADING_RATE_PALETTE_ENTRY_NO_INVOCATIONS_NV + 1)
	SHADING_RATE_PALETTE_ENTRY_MAX_ENUM_NV                    ShadingRatePaletteEntryNV = 0x7FFFFFFF
)

func (x ShadingRatePaletteEntryNV) String() string {
	switch x {
	case SHADING_RATE_PALETTE_ENTRY_NO_INVOCATIONS_NV:
		return "SHADING_RATE_PALETTE_ENTRY_NO_INVOCATIONS_NV"
	case SHADING_RATE_PALETTE_ENTRY_16_INVOCATIONS_PER_PIXEL_NV:
		return "SHADING_RATE_PALETTE_ENTRY_16_INVOCATIONS_PER_PIXEL_NV"
	case SHADING_RATE_PALETTE_ENTRY_8_INVOCATIONS_PER_PIXEL_NV:
		return "SHADING_RATE_PALETTE_ENTRY_8_INVOCATIONS_PER_PIXEL_NV"
	case SHADING_RATE_PALETTE_ENTRY_4_INVOCATIONS_PER_PIXEL_NV:
		return "SHADING_RATE_PALETTE_ENTRY_4_INVOCATIONS_PER_PIXEL_NV"
	case SHADING_RATE_PALETTE_ENTRY_2_INVOCATIONS_PER_PIXEL_NV:
		return "SHADING_RATE_PALETTE_ENTRY_2_INVOCATIONS_PER_PIXEL_NV"
	case SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_PIXEL_NV:
		return "SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_PIXEL_NV"
	case SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X1_PIXELS_NV:
		return "SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X1_PIXELS_NV"
	case SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_1X2_PIXELS_NV:
		return "SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_1X2_PIXELS_NV"
	case SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X2_PIXELS_NV:
		return "SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X2_PIXELS_NV"
	case SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X2_PIXELS_NV:
		return "SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X2_PIXELS_NV"
	case SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X4_PIXELS_NV:
		return "SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_2X4_PIXELS_NV"
	case SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X4_PIXELS_NV:
		return "SHADING_RATE_PALETTE_ENTRY_1_INVOCATION_PER_4X4_PIXELS_NV"
	case SHADING_RATE_PALETTE_ENTRY_MAX_ENUM_NV:
		return "SHADING_RATE_PALETTE_ENTRY_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// CoarseSampleOrderTypeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCoarseSampleOrderTypeNV.html
type CoarseSampleOrderTypeNV int32

const (
	COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV      CoarseSampleOrderTypeNV = 0
	COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV       CoarseSampleOrderTypeNV = 1
	COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV  CoarseSampleOrderTypeNV = 2
	COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV CoarseSampleOrderTypeNV = 3
	COARSE_SAMPLE_ORDER_TYPE_BEGIN_RANGE_NV  CoarseSampleOrderTypeNV = COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV
	COARSE_SAMPLE_ORDER_TYPE_END_RANGE_NV    CoarseSampleOrderTypeNV = COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV
	COARSE_SAMPLE_ORDER_TYPE_RANGE_SIZE_NV   CoarseSampleOrderTypeNV = (COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV - COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV + 1)
	COARSE_SAMPLE_ORDER_TYPE_MAX_ENUM_NV     CoarseSampleOrderTypeNV = 0x7FFFFFFF
)

func (x CoarseSampleOrderTypeNV) String() string {
	switch x {
	case COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV:
		return "COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV"
	case COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV:
		return "COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV"
	case COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV:
		return "COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV"
	case COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV:
		return "COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV"
	case COARSE_SAMPLE_ORDER_TYPE_MAX_ENUM_NV:
		return "COARSE_SAMPLE_ORDER_TYPE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ShadingRatePaletteNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShadingRatePaletteNV.html
type ShadingRatePaletteNV struct {
	ShadingRatePaletteEntryCount uint32
	PShadingRatePaletteEntries   *ShadingRatePaletteEntryNV
}

func NewShadingRatePaletteNV() *ShadingRatePaletteNV {
	return (*ShadingRatePaletteNV)(MemAlloc(unsafe.Sizeof(*(*ShadingRatePaletteNV)(nil))))
}
func (p *ShadingRatePaletteNV) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineViewportShadingRateImageStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html
type PipelineViewportShadingRateImageStateCreateInfoNV struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	ShadingRateImageEnable Bool32
	ViewportCount          uint32
	PShadingRatePalettes   *ShadingRatePaletteNV
}

func NewPipelineViewportShadingRateImageStateCreateInfoNV() *PipelineViewportShadingRateImageStateCreateInfoNV {
	p := (*PipelineViewportShadingRateImageStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineViewportShadingRateImageStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineViewportShadingRateImageStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceShadingRateImageFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShadingRateImageFeaturesNV.html
type PhysicalDeviceShadingRateImageFeaturesNV struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	ShadingRateImage             Bool32
	ShadingRateCoarseSampleOrder Bool32
}

func NewPhysicalDeviceShadingRateImageFeaturesNV() *PhysicalDeviceShadingRateImageFeaturesNV {
	p := (*PhysicalDeviceShadingRateImageFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShadingRateImageFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_FEATURES_NV
	return p
}
func (p *PhysicalDeviceShadingRateImageFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceShadingRateImagePropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShadingRateImagePropertiesNV.html
type PhysicalDeviceShadingRateImagePropertiesNV struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	ShadingRateTexelSize        Extent2D
	ShadingRatePaletteSize      uint32
	ShadingRateMaxCoarseSamples uint32
}

func NewPhysicalDeviceShadingRateImagePropertiesNV() *PhysicalDeviceShadingRateImagePropertiesNV {
	p := (*PhysicalDeviceShadingRateImagePropertiesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShadingRateImagePropertiesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV
	return p
}
func (p *PhysicalDeviceShadingRateImagePropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

// CoarseSampleLocationNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCoarseSampleLocationNV.html
type CoarseSampleLocationNV struct {
	PixelX uint32
	PixelY uint32
	Sample uint32
}

func NewCoarseSampleLocationNV() *CoarseSampleLocationNV {
	return (*CoarseSampleLocationNV)(MemAlloc(unsafe.Sizeof(*(*CoarseSampleLocationNV)(nil))))
}
func (p *CoarseSampleLocationNV) Free() { MemFree(unsafe.Pointer(p)) }

// CoarseSampleOrderCustomNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCoarseSampleOrderCustomNV.html
type CoarseSampleOrderCustomNV struct {
	ShadingRate         ShadingRatePaletteEntryNV
	SampleCount         uint32
	SampleLocationCount uint32
	PSampleLocations    *CoarseSampleLocationNV
}

func NewCoarseSampleOrderCustomNV() *CoarseSampleOrderCustomNV {
	return (*CoarseSampleOrderCustomNV)(MemAlloc(unsafe.Sizeof(*(*CoarseSampleOrderCustomNV)(nil))))
}
func (p *CoarseSampleOrderCustomNV) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineViewportCoarseSampleOrderStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html
type PipelineViewportCoarseSampleOrderStateCreateInfoNV struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	SampleOrderType        CoarseSampleOrderTypeNV
	CustomSampleOrderCount uint32
	PCustomSampleOrders    *CoarseSampleOrderCustomNV
}

func NewPipelineViewportCoarseSampleOrderStateCreateInfoNV() *PipelineViewportCoarseSampleOrderStateCreateInfoNV {
	p := (*PipelineViewportCoarseSampleOrderStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineViewportCoarseSampleOrderStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineViewportCoarseSampleOrderStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdBindShadingRateImageNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBindShadingRateImageNV.html
type PfnCmdBindShadingRateImageNV uintptr

func (fn PfnCmdBindShadingRateImageNV) Call(commandBuffer CommandBuffer, imageView ImageView, imageLayout ImageLayout) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(imageView), uintptr(imageLayout))
}
func (fn PfnCmdBindShadingRateImageNV) String() string { return "vkCmdBindShadingRateImageNV" }

//  PfnCmdSetViewportShadingRatePaletteNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetViewportShadingRatePaletteNV.html
type PfnCmdSetViewportShadingRatePaletteNV uintptr

func (fn PfnCmdSetViewportShadingRatePaletteNV) Call(commandBuffer CommandBuffer, firstViewport, viewportCount uint32, pShadingRatePalettes *ShadingRatePaletteNV) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstViewport), uintptr(viewportCount), uintptr(unsafe.Pointer(pShadingRatePalettes)))
}
func (fn PfnCmdSetViewportShadingRatePaletteNV) String() string {
	return "vkCmdSetViewportShadingRatePaletteNV"
}

//  PfnCmdSetCoarseSampleOrderNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetCoarseSampleOrderNV.html
type PfnCmdSetCoarseSampleOrderNV uintptr

func (fn PfnCmdSetCoarseSampleOrderNV) Call(commandBuffer CommandBuffer, sampleOrderType CoarseSampleOrderTypeNV, customSampleOrderCount uint32, pCustomSampleOrders *CoarseSampleOrderCustomNV) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(sampleOrderType), uintptr(customSampleOrderCount), uintptr(unsafe.Pointer(pCustomSampleOrders)))
}
func (fn PfnCmdSetCoarseSampleOrderNV) String() string { return "vkCmdSetCoarseSampleOrderNV" }

const NV_ray_tracing = 1

// AccelerationStructureNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureNV.html
type AccelerationStructureNV NonDispatchableHandle

const NV_RAY_TRACING_SPEC_VERSION = 3

var NV_RAY_TRACING_EXTENSION_NAME = "VK_NV_ray_tracing"

// AccelerationStructureTypeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureTypeNV.html
type AccelerationStructureTypeNV int32

const (
	ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV    AccelerationStructureTypeNV = 0
	ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV AccelerationStructureTypeNV = 1
	ACCELERATION_STRUCTURE_TYPE_BEGIN_RANGE_NV  AccelerationStructureTypeNV = ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV
	ACCELERATION_STRUCTURE_TYPE_END_RANGE_NV    AccelerationStructureTypeNV = ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV
	ACCELERATION_STRUCTURE_TYPE_RANGE_SIZE_NV   AccelerationStructureTypeNV = (ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV - ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV + 1)
	ACCELERATION_STRUCTURE_TYPE_MAX_ENUM_NV     AccelerationStructureTypeNV = 0x7FFFFFFF
)

func (x AccelerationStructureTypeNV) String() string {
	switch x {
	case ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV:
		return "ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_NV"
	case ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV:
		return "ACCELERATION_STRUCTURE_TYPE_BOTTOM_LEVEL_NV"
	case ACCELERATION_STRUCTURE_TYPE_MAX_ENUM_NV:
		return "ACCELERATION_STRUCTURE_TYPE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// RayTracingShaderGroupTypeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRayTracingShaderGroupTypeNV.html
type RayTracingShaderGroupTypeNV int32

const (
	RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV              RayTracingShaderGroupTypeNV = 0
	RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV  RayTracingShaderGroupTypeNV = 1
	RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV RayTracingShaderGroupTypeNV = 2
	RAY_TRACING_SHADER_GROUP_TYPE_BEGIN_RANGE_NV          RayTracingShaderGroupTypeNV = RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV
	RAY_TRACING_SHADER_GROUP_TYPE_END_RANGE_NV            RayTracingShaderGroupTypeNV = RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV
	RAY_TRACING_SHADER_GROUP_TYPE_RANGE_SIZE_NV           RayTracingShaderGroupTypeNV = (RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV - RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV + 1)
	RAY_TRACING_SHADER_GROUP_TYPE_MAX_ENUM_NV             RayTracingShaderGroupTypeNV = 0x7FFFFFFF
)

func (x RayTracingShaderGroupTypeNV) String() string {
	switch x {
	case RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV:
		return "RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV"
	case RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV:
		return "RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV"
	case RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV:
		return "RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV"
	case RAY_TRACING_SHADER_GROUP_TYPE_MAX_ENUM_NV:
		return "RAY_TRACING_SHADER_GROUP_TYPE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// GeometryTypeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGeometryTypeNV.html
type GeometryTypeNV int32

const (
	GEOMETRY_TYPE_TRIANGLES_NV   GeometryTypeNV = 0
	GEOMETRY_TYPE_AABBS_NV       GeometryTypeNV = 1
	GEOMETRY_TYPE_BEGIN_RANGE_NV GeometryTypeNV = GEOMETRY_TYPE_TRIANGLES_NV
	GEOMETRY_TYPE_END_RANGE_NV   GeometryTypeNV = GEOMETRY_TYPE_AABBS_NV
	GEOMETRY_TYPE_RANGE_SIZE_NV  GeometryTypeNV = (GEOMETRY_TYPE_AABBS_NV - GEOMETRY_TYPE_TRIANGLES_NV + 1)
	GEOMETRY_TYPE_MAX_ENUM_NV    GeometryTypeNV = 0x7FFFFFFF
)

func (x GeometryTypeNV) String() string {
	switch x {
	case GEOMETRY_TYPE_TRIANGLES_NV:
		return "GEOMETRY_TYPE_TRIANGLES_NV"
	case GEOMETRY_TYPE_AABBS_NV:
		return "GEOMETRY_TYPE_AABBS_NV"
	case GEOMETRY_TYPE_MAX_ENUM_NV:
		return "GEOMETRY_TYPE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// CopyAccelerationStructureModeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCopyAccelerationStructureModeNV.html
type CopyAccelerationStructureModeNV int32

const (
	COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV       CopyAccelerationStructureModeNV = 0
	COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV     CopyAccelerationStructureModeNV = 1
	COPY_ACCELERATION_STRUCTURE_MODE_BEGIN_RANGE_NV CopyAccelerationStructureModeNV = COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV
	COPY_ACCELERATION_STRUCTURE_MODE_END_RANGE_NV   CopyAccelerationStructureModeNV = COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV
	COPY_ACCELERATION_STRUCTURE_MODE_RANGE_SIZE_NV  CopyAccelerationStructureModeNV = (COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV - COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV + 1)
	COPY_ACCELERATION_STRUCTURE_MODE_MAX_ENUM_NV    CopyAccelerationStructureModeNV = 0x7FFFFFFF
)

func (x CopyAccelerationStructureModeNV) String() string {
	switch x {
	case COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV:
		return "COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV"
	case COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV:
		return "COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV"
	case COPY_ACCELERATION_STRUCTURE_MODE_MAX_ENUM_NV:
		return "COPY_ACCELERATION_STRUCTURE_MODE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// AccelerationStructureMemoryRequirementsTypeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureMemoryRequirementsTypeNV.html
type AccelerationStructureMemoryRequirementsTypeNV int32

const (
	ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV         AccelerationStructureMemoryRequirementsTypeNV = 0
	ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV  AccelerationStructureMemoryRequirementsTypeNV = 1
	ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV AccelerationStructureMemoryRequirementsTypeNV = 2
	ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BEGIN_RANGE_NV    AccelerationStructureMemoryRequirementsTypeNV = ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV
	ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_END_RANGE_NV      AccelerationStructureMemoryRequirementsTypeNV = ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV
	ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_RANGE_SIZE_NV     AccelerationStructureMemoryRequirementsTypeNV = (ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV - ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV + 1)
	ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_MAX_ENUM_NV       AccelerationStructureMemoryRequirementsTypeNV = 0x7FFFFFFF
)

func (x AccelerationStructureMemoryRequirementsTypeNV) String() string {
	switch x {
	case ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV:
		return "ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV"
	case ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV:
		return "ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV"
	case ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV:
		return "ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV"
	case ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_MAX_ENUM_NV:
		return "ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// GeometryFlagsNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGeometryFlagsNV.html
type GeometryFlagsNV uint32

const (
	GEOMETRY_OPAQUE_BIT_NV                          GeometryFlagsNV = 0x00000001
	GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_NV GeometryFlagsNV = 0x00000002
	GEOMETRY_FLAG_BITS_MAX_ENUM_NV                  GeometryFlagsNV = 0x7FFFFFFF
)

func (x GeometryFlagsNV) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch GeometryFlagsNV(1 << i) {
			case GEOMETRY_OPAQUE_BIT_NV:
				s += "GEOMETRY_OPAQUE_BIT_NV|"
			case GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_NV:
				s += "GEOMETRY_NO_DUPLICATE_ANY_HIT_INVOCATION_BIT_NV|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// GeometryInstanceFlagsNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGeometryInstanceFlagsNV.html
type GeometryInstanceFlagsNV uint32

const (
	GEOMETRY_INSTANCE_TRIANGLE_CULL_DISABLE_BIT_NV           GeometryInstanceFlagsNV = 0x00000001
	GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_NV GeometryInstanceFlagsNV = 0x00000002
	GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_NV                    GeometryInstanceFlagsNV = 0x00000004
	GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_NV                 GeometryInstanceFlagsNV = 0x00000008
	GEOMETRY_INSTANCE_FLAG_BITS_MAX_ENUM_NV                  GeometryInstanceFlagsNV = 0x7FFFFFFF
)

func (x GeometryInstanceFlagsNV) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch GeometryInstanceFlagsNV(1 << i) {
			case GEOMETRY_INSTANCE_TRIANGLE_CULL_DISABLE_BIT_NV:
				s += "GEOMETRY_INSTANCE_TRIANGLE_CULL_DISABLE_BIT_NV|"
			case GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_NV:
				s += "GEOMETRY_INSTANCE_TRIANGLE_FRONT_COUNTERCLOCKWISE_BIT_NV|"
			case GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_NV:
				s += "GEOMETRY_INSTANCE_FORCE_OPAQUE_BIT_NV|"
			case GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_NV:
				s += "GEOMETRY_INSTANCE_FORCE_NO_OPAQUE_BIT_NV|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// BuildAccelerationStructureFlagsNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBuildAccelerationStructureFlagsNV.html
type BuildAccelerationStructureFlagsNV uint32

const (
	BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV      BuildAccelerationStructureFlagsNV = 0x00000001
	BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_NV  BuildAccelerationStructureFlagsNV = 0x00000002
	BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV BuildAccelerationStructureFlagsNV = 0x00000004
	BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV BuildAccelerationStructureFlagsNV = 0x00000008
	BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_NV        BuildAccelerationStructureFlagsNV = 0x00000010
	BUILD_ACCELERATION_STRUCTURE_FLAG_BITS_MAX_ENUM_NV    BuildAccelerationStructureFlagsNV = 0x7FFFFFFF
)

func (x BuildAccelerationStructureFlagsNV) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch BuildAccelerationStructureFlagsNV(1 << i) {
			case BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV:
				s += "BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV|"
			case BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_NV:
				s += "BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_NV|"
			case BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV:
				s += "BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV|"
			case BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV:
				s += "BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV|"
			case BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_NV:
				s += "BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_NV|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// RayTracingShaderGroupCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRayTracingShaderGroupCreateInfoNV.html
type RayTracingShaderGroupCreateInfoNV struct {
	SType              StructureType
	PNext              unsafe.Pointer
	Type               RayTracingShaderGroupTypeNV
	GeneralShader      uint32
	ClosestHitShader   uint32
	AnyHitShader       uint32
	IntersectionShader uint32
}

func NewRayTracingShaderGroupCreateInfoNV() *RayTracingShaderGroupCreateInfoNV {
	p := (*RayTracingShaderGroupCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*RayTracingShaderGroupCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV
	return p
}
func (p *RayTracingShaderGroupCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// RayTracingPipelineCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html
type RayTracingPipelineCreateInfoNV struct {
	SType              StructureType
	PNext              unsafe.Pointer
	Flags              PipelineCreateFlags
	StageCount         uint32
	PStages            *PipelineShaderStageCreateInfo
	GroupCount         uint32
	PGroups            *RayTracingShaderGroupCreateInfoNV
	MaxRecursionDepth  uint32
	Layout             PipelineLayout
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

func NewRayTracingPipelineCreateInfoNV() *RayTracingPipelineCreateInfoNV {
	p := (*RayTracingPipelineCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*RayTracingPipelineCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_NV
	return p
}
func (p *RayTracingPipelineCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// GeometryTrianglesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGeometryTrianglesNV.html
type GeometryTrianglesNV struct {
	SType           StructureType
	PNext           unsafe.Pointer
	VertexData      Buffer
	VertexOffset    DeviceSize
	VertexCount     uint32
	VertexStride    DeviceSize
	VertexFormat    Format
	IndexData       Buffer
	IndexOffset     DeviceSize
	IndexCount      uint32
	IndexType       IndexType
	TransformData   Buffer
	TransformOffset DeviceSize
}

func NewGeometryTrianglesNV() *GeometryTrianglesNV {
	p := (*GeometryTrianglesNV)(MemAlloc(unsafe.Sizeof(*(*GeometryTrianglesNV)(nil))))
	p.SType = STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV
	return p
}
func (p *GeometryTrianglesNV) Free() { MemFree(unsafe.Pointer(p)) }

// GeometryAABBNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGeometryAABBNV.html
type GeometryAABBNV struct {
	SType    StructureType
	PNext    unsafe.Pointer
	AabbData Buffer
	NumAABBs uint32
	Stride   uint32
	Offset   DeviceSize
}

func NewGeometryAABBNV() *GeometryAABBNV {
	p := (*GeometryAABBNV)(MemAlloc(unsafe.Sizeof(*(*GeometryAABBNV)(nil))))
	p.SType = STRUCTURE_TYPE_GEOMETRY_AABB_NV
	return p
}
func (p *GeometryAABBNV) Free() { MemFree(unsafe.Pointer(p)) }

// GeometryDataNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGeometryDataNV.html
type GeometryDataNV struct {
	Triangles GeometryTrianglesNV
	Aabbs     GeometryAABBNV
}

func NewGeometryDataNV() *GeometryDataNV {
	return (*GeometryDataNV)(MemAlloc(unsafe.Sizeof(*(*GeometryDataNV)(nil))))
}
func (p *GeometryDataNV) Free() { MemFree(unsafe.Pointer(p)) }

// GeometryNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkGeometryNV.html
type GeometryNV struct {
	SType        StructureType
	PNext        unsafe.Pointer
	GeometryType GeometryTypeNV
	Geometry     GeometryDataNV
	Flags        GeometryFlagsNV
}

func NewGeometryNV() *GeometryNV {
	p := (*GeometryNV)(MemAlloc(unsafe.Sizeof(*(*GeometryNV)(nil))))
	p.SType = STRUCTURE_TYPE_GEOMETRY_NV
	return p
}
func (p *GeometryNV) Free() { MemFree(unsafe.Pointer(p)) }

// AccelerationStructureInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureInfoNV.html
type AccelerationStructureInfoNV struct {
	SType         StructureType
	PNext         unsafe.Pointer
	Type          AccelerationStructureTypeNV
	Flags         BuildAccelerationStructureFlagsNV
	InstanceCount uint32
	GeometryCount uint32
	PGeometries   *GeometryNV
}

func NewAccelerationStructureInfoNV() *AccelerationStructureInfoNV {
	p := (*AccelerationStructureInfoNV)(MemAlloc(unsafe.Sizeof(*(*AccelerationStructureInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_ACCELERATION_STRUCTURE_INFO_NV
	return p
}
func (p *AccelerationStructureInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// AccelerationStructureCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureCreateInfoNV.html
type AccelerationStructureCreateInfoNV struct {
	SType         StructureType
	PNext         unsafe.Pointer
	CompactedSize DeviceSize
	Info          AccelerationStructureInfoNV
}

func NewAccelerationStructureCreateInfoNV() *AccelerationStructureCreateInfoNV {
	p := (*AccelerationStructureCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*AccelerationStructureCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV
	return p
}
func (p *AccelerationStructureCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// BindAccelerationStructureMemoryInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBindAccelerationStructureMemoryInfoNV.html
type BindAccelerationStructureMemoryInfoNV struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	AccelerationStructure AccelerationStructureNV
	Memory                DeviceMemory
	MemoryOffset          DeviceSize
	DeviceIndexCount      uint32
	PDeviceIndices        *uint32
}

func NewBindAccelerationStructureMemoryInfoNV() *BindAccelerationStructureMemoryInfoNV {
	p := (*BindAccelerationStructureMemoryInfoNV)(MemAlloc(unsafe.Sizeof(*(*BindAccelerationStructureMemoryInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV
	return p
}
func (p *BindAccelerationStructureMemoryInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// WriteDescriptorSetAccelerationStructureNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkWriteDescriptorSetAccelerationStructureNV.html
type WriteDescriptorSetAccelerationStructureNV struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	AccelerationStructureCount uint32
	PAccelerationStructures    *AccelerationStructureNV
}

func NewWriteDescriptorSetAccelerationStructureNV() *WriteDescriptorSetAccelerationStructureNV {
	p := (*WriteDescriptorSetAccelerationStructureNV)(MemAlloc(unsafe.Sizeof(*(*WriteDescriptorSetAccelerationStructureNV)(nil))))
	p.SType = STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV
	return p
}
func (p *WriteDescriptorSetAccelerationStructureNV) Free() { MemFree(unsafe.Pointer(p)) }

// AccelerationStructureMemoryRequirementsInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html
type AccelerationStructureMemoryRequirementsInfoNV struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Type                  AccelerationStructureMemoryRequirementsTypeNV
	AccelerationStructure AccelerationStructureNV
}

func NewAccelerationStructureMemoryRequirementsInfoNV() *AccelerationStructureMemoryRequirementsInfoNV {
	p := (*AccelerationStructureMemoryRequirementsInfoNV)(MemAlloc(unsafe.Sizeof(*(*AccelerationStructureMemoryRequirementsInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV
	return p
}
func (p *AccelerationStructureMemoryRequirementsInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceRayTracingPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html
type PhysicalDeviceRayTracingPropertiesNV struct {
	SType                                  StructureType
	PNext                                  unsafe.Pointer
	ShaderGroupHandleSize                  uint32
	MaxRecursionDepth                      uint32
	MaxShaderGroupStride                   uint32
	ShaderGroupBaseAlignment               uint32
	MaxGeometryCount                       uint64
	MaxInstanceCount                       uint64
	MaxTriangleCount                       uint64
	MaxDescriptorSetAccelerationStructures uint32
}

func NewPhysicalDeviceRayTracingPropertiesNV() *PhysicalDeviceRayTracingPropertiesNV {
	p := (*PhysicalDeviceRayTracingPropertiesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceRayTracingPropertiesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PROPERTIES_NV
	return p
}
func (p *PhysicalDeviceRayTracingPropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateAccelerationStructureNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateAccelerationStructureNV.html
type PfnCreateAccelerationStructureNV uintptr

func (fn PfnCreateAccelerationStructureNV) Call(device Device, pCreateInfo *AccelerationStructureCreateInfoNV, pAllocator *AllocationCallbacks, pAccelerationStructure *AccelerationStructureNV) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pAccelerationStructure)))
	return Result(ret)
}
func (fn PfnCreateAccelerationStructureNV) String() string { return "vkCreateAccelerationStructureNV" }

//  PfnDestroyAccelerationStructureNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkDestroyAccelerationStructureNV.html
type PfnDestroyAccelerationStructureNV uintptr

func (fn PfnDestroyAccelerationStructureNV) Call(device Device, accelerationStructure AccelerationStructureNV, pAllocator *AllocationCallbacks) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(accelerationStructure), uintptr(unsafe.Pointer(pAllocator)))
}
func (fn PfnDestroyAccelerationStructureNV) String() string {
	return "vkDestroyAccelerationStructureNV"
}

//  PfnGetAccelerationStructureMemoryRequirementsNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html
type PfnGetAccelerationStructureMemoryRequirementsNV uintptr

func (fn PfnGetAccelerationStructureMemoryRequirementsNV) Call(device Device, pInfo *AccelerationStructureMemoryRequirementsInfoNV, pMemoryRequirements *MemoryRequirements2KHR) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)), uintptr(unsafe.Pointer(pMemoryRequirements)))
}
func (fn PfnGetAccelerationStructureMemoryRequirementsNV) String() string {
	return "vkGetAccelerationStructureMemoryRequirementsNV"
}

//  PfnBindAccelerationStructureMemoryNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkBindAccelerationStructureMemoryNV.html
type PfnBindAccelerationStructureMemoryNV uintptr

func (fn PfnBindAccelerationStructureMemoryNV) Call(device Device, bindInfoCount uint32, pBindInfos *BindAccelerationStructureMemoryInfoNV) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(bindInfoCount), uintptr(unsafe.Pointer(pBindInfos)))
	return Result(ret)
}
func (fn PfnBindAccelerationStructureMemoryNV) String() string {
	return "vkBindAccelerationStructureMemoryNV"
}

//  PfnCmdBuildAccelerationStructureNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdBuildAccelerationStructureNV.html
type PfnCmdBuildAccelerationStructureNV uintptr

func (fn PfnCmdBuildAccelerationStructureNV) Call(commandBuffer CommandBuffer, pInfo *AccelerationStructureInfoNV, instanceData Buffer, instanceOffset DeviceSize, update Bool32, dst, src AccelerationStructureNV, scratch Buffer, scratchOffset DeviceSize) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pInfo)), uintptr(instanceData), uintptr(instanceOffset), uintptr(update), uintptr(dst), uintptr(src), uintptr(scratch), uintptr(scratchOffset))
}
func (fn PfnCmdBuildAccelerationStructureNV) String() string {
	return "vkCmdBuildAccelerationStructureNV"
}

//  PfnCmdCopyAccelerationStructureNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdCopyAccelerationStructureNV.html
type PfnCmdCopyAccelerationStructureNV uintptr

func (fn PfnCmdCopyAccelerationStructureNV) Call(commandBuffer CommandBuffer, dst, src AccelerationStructureNV, mode CopyAccelerationStructureModeNV) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(dst), uintptr(src), uintptr(mode))
}
func (fn PfnCmdCopyAccelerationStructureNV) String() string {
	return "vkCmdCopyAccelerationStructureNV"
}

//  PfnCmdTraceRaysNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdTraceRaysNV.html
type PfnCmdTraceRaysNV uintptr

func (fn PfnCmdTraceRaysNV) Call(commandBuffer CommandBuffer, raygenShaderBindingTableBuffer Buffer, raygenShaderBindingOffset DeviceSize, missShaderBindingTableBuffer Buffer, missShaderBindingOffset, missShaderBindingStride DeviceSize, hitShaderBindingTableBuffer Buffer, hitShaderBindingOffset, hitShaderBindingStride DeviceSize, callableShaderBindingTableBuffer Buffer, callableShaderBindingOffset, callableShaderBindingStride DeviceSize, width, height, depth uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(raygenShaderBindingTableBuffer), uintptr(raygenShaderBindingOffset), uintptr(missShaderBindingTableBuffer), uintptr(missShaderBindingOffset), uintptr(missShaderBindingStride), uintptr(hitShaderBindingTableBuffer), uintptr(hitShaderBindingOffset), uintptr(hitShaderBindingStride), uintptr(callableShaderBindingTableBuffer), uintptr(callableShaderBindingOffset), uintptr(callableShaderBindingStride), uintptr(width), uintptr(height), uintptr(depth))
}
func (fn PfnCmdTraceRaysNV) String() string { return "vkCmdTraceRaysNV" }

//  PfnCreateRayTracingPipelinesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateRayTracingPipelinesNV.html
type PfnCreateRayTracingPipelinesNV uintptr

func (fn PfnCreateRayTracingPipelinesNV) Call(device Device, pipelineCache PipelineCache, createInfoCount uint32, pCreateInfos *RayTracingPipelineCreateInfoNV, pAllocator *AllocationCallbacks, pPipelines *Pipeline) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(pipelineCache), uintptr(createInfoCount), uintptr(unsafe.Pointer(pCreateInfos)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pPipelines)))
	return Result(ret)
}
func (fn PfnCreateRayTracingPipelinesNV) String() string { return "vkCreateRayTracingPipelinesNV" }

//  PfnGetRayTracingShaderGroupHandlesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetRayTracingShaderGroupHandlesNV.html
type PfnGetRayTracingShaderGroupHandlesNV uintptr

func (fn PfnGetRayTracingShaderGroupHandlesNV) Call(device Device, pipeline Pipeline, firstGroup, groupCount uint32, dataSize uintptr, pData unsafe.Pointer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(pipeline), uintptr(firstGroup), uintptr(groupCount), uintptr(dataSize), uintptr(pData))
	return Result(ret)
}
func (fn PfnGetRayTracingShaderGroupHandlesNV) String() string {
	return "vkGetRayTracingShaderGroupHandlesNV"
}

//  PfnGetAccelerationStructureHandleNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetAccelerationStructureHandleNV.html
type PfnGetAccelerationStructureHandleNV uintptr

func (fn PfnGetAccelerationStructureHandleNV) Call(device Device, accelerationStructure AccelerationStructureNV, dataSize uintptr, pData unsafe.Pointer) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(accelerationStructure), uintptr(dataSize), uintptr(pData))
	return Result(ret)
}
func (fn PfnGetAccelerationStructureHandleNV) String() string {
	return "vkGetAccelerationStructureHandleNV"
}

//  PfnCmdWriteAccelerationStructuresPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html
type PfnCmdWriteAccelerationStructuresPropertiesNV uintptr

func (fn PfnCmdWriteAccelerationStructuresPropertiesNV) Call(commandBuffer CommandBuffer, accelerationStructureCount uint32, pAccelerationStructures *AccelerationStructureNV, queryType QueryType, queryPool QueryPool, firstQuery uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(accelerationStructureCount), uintptr(unsafe.Pointer(pAccelerationStructures)), uintptr(queryType), uintptr(queryPool), uintptr(firstQuery))
}
func (fn PfnCmdWriteAccelerationStructuresPropertiesNV) String() string {
	return "vkCmdWriteAccelerationStructuresPropertiesNV"
}

//  PfnCompileDeferredNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCompileDeferredNV.html
type PfnCompileDeferredNV uintptr

func (fn PfnCompileDeferredNV) Call(device Device, pipeline Pipeline, shader uint32) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(pipeline), uintptr(shader))
	return Result(ret)
}
func (fn PfnCompileDeferredNV) String() string { return "vkCompileDeferredNV" }

const NV_representative_fragment_test = 1
const NV_REPRESENTATIVE_FRAGMENT_TEST_SPEC_VERSION = 2

var NV_REPRESENTATIVE_FRAGMENT_TEST_EXTENSION_NAME = "VK_NV_representative_fragment_test"

// PhysicalDeviceRepresentativeFragmentTestFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV.html
type PhysicalDeviceRepresentativeFragmentTestFeaturesNV struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	RepresentativeFragmentTest Bool32
}

func NewPhysicalDeviceRepresentativeFragmentTestFeaturesNV() *PhysicalDeviceRepresentativeFragmentTestFeaturesNV {
	p := (*PhysicalDeviceRepresentativeFragmentTestFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceRepresentativeFragmentTestFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV
	return p
}
func (p *PhysicalDeviceRepresentativeFragmentTestFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineRepresentativeFragmentTestStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html
type PipelineRepresentativeFragmentTestStateCreateInfoNV struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	RepresentativeFragmentTestEnable Bool32
}

func NewPipelineRepresentativeFragmentTestStateCreateInfoNV() *PipelineRepresentativeFragmentTestStateCreateInfoNV {
	p := (*PipelineRepresentativeFragmentTestStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineRepresentativeFragmentTestStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineRepresentativeFragmentTestStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_filter_cubic = 1
const EXT_FILTER_CUBIC_SPEC_VERSION = 2

var EXT_FILTER_CUBIC_EXTENSION_NAME = "VK_EXT_filter_cubic"

// PhysicalDeviceImageViewImageFormatInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceImageViewImageFormatInfoEXT.html
type PhysicalDeviceImageViewImageFormatInfoEXT struct {
	SType         StructureType
	PNext         unsafe.Pointer
	ImageViewType ImageViewType
}

func NewPhysicalDeviceImageViewImageFormatInfoEXT() *PhysicalDeviceImageViewImageFormatInfoEXT {
	p := (*PhysicalDeviceImageViewImageFormatInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceImageViewImageFormatInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_IMAGE_FORMAT_INFO_EXT
	return p
}
func (p *PhysicalDeviceImageViewImageFormatInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// FilterCubicImageViewImageFormatPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html
type FilterCubicImageViewImageFormatPropertiesEXT struct {
	SType             StructureType
	PNext             unsafe.Pointer
	FilterCubic       Bool32
	FilterCubicMinmax Bool32
}

func NewFilterCubicImageViewImageFormatPropertiesEXT() *FilterCubicImageViewImageFormatPropertiesEXT {
	p := (*FilterCubicImageViewImageFormatPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*FilterCubicImageViewImageFormatPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_FILTER_CUBIC_IMAGE_VIEW_IMAGE_FORMAT_PROPERTIES_EXT
	return p
}
func (p *FilterCubicImageViewImageFormatPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_global_priority = 1
const EXT_GLOBAL_PRIORITY_SPEC_VERSION = 2

var EXT_GLOBAL_PRIORITY_EXTENSION_NAME = "VK_EXT_global_priority"

// QueueGlobalPriorityEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueueGlobalPriorityEXT.html
type QueueGlobalPriorityEXT int32

const (
	QUEUE_GLOBAL_PRIORITY_LOW_EXT         QueueGlobalPriorityEXT = 128
	QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT      QueueGlobalPriorityEXT = 256
	QUEUE_GLOBAL_PRIORITY_HIGH_EXT        QueueGlobalPriorityEXT = 512
	QUEUE_GLOBAL_PRIORITY_REALTIME_EXT    QueueGlobalPriorityEXT = 1024
	QUEUE_GLOBAL_PRIORITY_BEGIN_RANGE_EXT QueueGlobalPriorityEXT = QUEUE_GLOBAL_PRIORITY_LOW_EXT
	QUEUE_GLOBAL_PRIORITY_END_RANGE_EXT   QueueGlobalPriorityEXT = QUEUE_GLOBAL_PRIORITY_REALTIME_EXT
	QUEUE_GLOBAL_PRIORITY_RANGE_SIZE_EXT  QueueGlobalPriorityEXT = (QUEUE_GLOBAL_PRIORITY_REALTIME_EXT - QUEUE_GLOBAL_PRIORITY_LOW_EXT + 1)
	QUEUE_GLOBAL_PRIORITY_MAX_ENUM_EXT    QueueGlobalPriorityEXT = 0x7FFFFFFF
)

func (x QueueGlobalPriorityEXT) String() string {
	switch x {
	case QUEUE_GLOBAL_PRIORITY_LOW_EXT:
		return "QUEUE_GLOBAL_PRIORITY_LOW_EXT"
	case QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT:
		return "QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT"
	case QUEUE_GLOBAL_PRIORITY_HIGH_EXT:
		return "QUEUE_GLOBAL_PRIORITY_HIGH_EXT"
	case QUEUE_GLOBAL_PRIORITY_REALTIME_EXT:
		return "QUEUE_GLOBAL_PRIORITY_REALTIME_EXT"
	case QUEUE_GLOBAL_PRIORITY_MAX_ENUM_EXT:
		return "QUEUE_GLOBAL_PRIORITY_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DeviceQueueGlobalPriorityCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceQueueGlobalPriorityCreateInfoEXT.html
type DeviceQueueGlobalPriorityCreateInfoEXT struct {
	SType          StructureType
	PNext          unsafe.Pointer
	GlobalPriority QueueGlobalPriorityEXT
}

func NewDeviceQueueGlobalPriorityCreateInfoEXT() *DeviceQueueGlobalPriorityCreateInfoEXT {
	p := (*DeviceQueueGlobalPriorityCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*DeviceQueueGlobalPriorityCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_EXT
	return p
}
func (p *DeviceQueueGlobalPriorityCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_external_memory_host = 1
const EXT_EXTERNAL_MEMORY_HOST_SPEC_VERSION = 1

var EXT_EXTERNAL_MEMORY_HOST_EXTENSION_NAME = "VK_EXT_external_memory_host"

// ImportMemoryHostPointerInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportMemoryHostPointerInfoEXT.html
type ImportMemoryHostPointerInfoEXT struct {
	SType        StructureType
	PNext        unsafe.Pointer
	HandleType   ExternalMemoryHandleTypeFlags
	PHostPointer unsafe.Pointer
}

func NewImportMemoryHostPointerInfoEXT() *ImportMemoryHostPointerInfoEXT {
	p := (*ImportMemoryHostPointerInfoEXT)(MemAlloc(unsafe.Sizeof(*(*ImportMemoryHostPointerInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT
	return p
}
func (p *ImportMemoryHostPointerInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryHostPointerPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryHostPointerPropertiesEXT.html
type MemoryHostPointerPropertiesEXT struct {
	SType          StructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

func NewMemoryHostPointerPropertiesEXT() *MemoryHostPointerPropertiesEXT {
	p := (*MemoryHostPointerPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*MemoryHostPointerPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT
	return p
}
func (p *MemoryHostPointerPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceExternalMemoryHostPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceExternalMemoryHostPropertiesEXT.html
type PhysicalDeviceExternalMemoryHostPropertiesEXT struct {
	SType                           StructureType
	PNext                           unsafe.Pointer
	MinImportedHostPointerAlignment DeviceSize
}

func NewPhysicalDeviceExternalMemoryHostPropertiesEXT() *PhysicalDeviceExternalMemoryHostPropertiesEXT {
	p := (*PhysicalDeviceExternalMemoryHostPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceExternalMemoryHostPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceExternalMemoryHostPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetMemoryHostPointerPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetMemoryHostPointerPropertiesEXT.html
type PfnGetMemoryHostPointerPropertiesEXT uintptr

func (fn PfnGetMemoryHostPointerPropertiesEXT) Call(device Device, handleType ExternalMemoryHandleTypeFlags, pHostPointer unsafe.Pointer, pMemoryHostPointerProperties *MemoryHostPointerPropertiesEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(handleType), uintptr(pHostPointer), uintptr(unsafe.Pointer(pMemoryHostPointerProperties)))
	return Result(ret)
}
func (fn PfnGetMemoryHostPointerPropertiesEXT) String() string {
	return "vkGetMemoryHostPointerPropertiesEXT"
}

const AMD_buffer_marker = 1
const AMD_BUFFER_MARKER_SPEC_VERSION = 1

var AMD_BUFFER_MARKER_EXTENSION_NAME = "VK_AMD_buffer_marker"

//  PfnCmdWriteBufferMarkerAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdWriteBufferMarkerAMD.html
type PfnCmdWriteBufferMarkerAMD uintptr

func (fn PfnCmdWriteBufferMarkerAMD) Call(commandBuffer CommandBuffer, pipelineStage PipelineStageFlags, dstBuffer Buffer, dstOffset DeviceSize, marker uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(pipelineStage), uintptr(dstBuffer), uintptr(dstOffset), uintptr(marker))
}
func (fn PfnCmdWriteBufferMarkerAMD) String() string { return "vkCmdWriteBufferMarkerAMD" }

const AMD_pipeline_compiler_control = 1
const AMD_PIPELINE_COMPILER_CONTROL_SPEC_VERSION = 1

var AMD_PIPELINE_COMPILER_CONTROL_EXTENSION_NAME = "VK_AMD_pipeline_compiler_control"

// PipelineCompilerControlFlagsAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCompilerControlFlagsAMD.html
type PipelineCompilerControlFlagsAMD uint32

const (
	PIPELINE_COMPILER_CONTROL_FLAG_BITS_MAX_ENUM_AMD PipelineCompilerControlFlagsAMD = 0x7FFFFFFF
)

func (x PipelineCompilerControlFlagsAMD) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch PipelineCompilerControlFlagsAMD(1 << i) {
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// PipelineCompilerControlCreateInfoAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCompilerControlCreateInfoAMD.html
type PipelineCompilerControlCreateInfoAMD struct {
	SType                StructureType
	PNext                unsafe.Pointer
	CompilerControlFlags PipelineCompilerControlFlagsAMD
}

func NewPipelineCompilerControlCreateInfoAMD() *PipelineCompilerControlCreateInfoAMD {
	return (*PipelineCompilerControlCreateInfoAMD)(MemAlloc(unsafe.Sizeof(*(*PipelineCompilerControlCreateInfoAMD)(nil))))
}
func (p *PipelineCompilerControlCreateInfoAMD) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_calibrated_timestamps = 1
const EXT_CALIBRATED_TIMESTAMPS_SPEC_VERSION = 1

var EXT_CALIBRATED_TIMESTAMPS_EXTENSION_NAME = "VK_EXT_calibrated_timestamps"

// TimeDomainEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkTimeDomainEXT.html
type TimeDomainEXT int32

const (
	TIME_DOMAIN_DEVICE_EXT                    TimeDomainEXT = 0
	TIME_DOMAIN_CLOCK_MONOTONIC_EXT           TimeDomainEXT = 1
	TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT       TimeDomainEXT = 2
	TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT TimeDomainEXT = 3
	TIME_DOMAIN_BEGIN_RANGE_EXT               TimeDomainEXT = TIME_DOMAIN_DEVICE_EXT
	TIME_DOMAIN_END_RANGE_EXT                 TimeDomainEXT = TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT
	TIME_DOMAIN_RANGE_SIZE_EXT                TimeDomainEXT = (TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT - TIME_DOMAIN_DEVICE_EXT + 1)
	TIME_DOMAIN_MAX_ENUM_EXT                  TimeDomainEXT = 0x7FFFFFFF
)

func (x TimeDomainEXT) String() string {
	switch x {
	case TIME_DOMAIN_DEVICE_EXT:
		return "TIME_DOMAIN_DEVICE_EXT"
	case TIME_DOMAIN_CLOCK_MONOTONIC_EXT:
		return "TIME_DOMAIN_CLOCK_MONOTONIC_EXT"
	case TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT:
		return "TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT"
	case TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT:
		return "TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT"
	case TIME_DOMAIN_MAX_ENUM_EXT:
		return "TIME_DOMAIN_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// CalibratedTimestampInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCalibratedTimestampInfoEXT.html
type CalibratedTimestampInfoEXT struct {
	SType      StructureType
	PNext      unsafe.Pointer
	TimeDomain TimeDomainEXT
}

func NewCalibratedTimestampInfoEXT() *CalibratedTimestampInfoEXT {
	p := (*CalibratedTimestampInfoEXT)(MemAlloc(unsafe.Sizeof(*(*CalibratedTimestampInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_EXT
	return p
}
func (p *CalibratedTimestampInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceCalibrateableTimeDomainsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsEXT.html
type PfnGetPhysicalDeviceCalibrateableTimeDomainsEXT uintptr

func (fn PfnGetPhysicalDeviceCalibrateableTimeDomainsEXT) Call(physicalDevice PhysicalDevice, pTimeDomainCount *uint32, pTimeDomains *TimeDomainEXT) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pTimeDomainCount)), uintptr(unsafe.Pointer(pTimeDomains)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceCalibrateableTimeDomainsEXT) String() string {
	return "vkGetPhysicalDeviceCalibrateableTimeDomainsEXT"
}

//  PfnGetCalibratedTimestampsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetCalibratedTimestampsEXT.html
type PfnGetCalibratedTimestampsEXT uintptr

func (fn PfnGetCalibratedTimestampsEXT) Call(device Device, timestampCount uint32, pTimestampInfos *CalibratedTimestampInfoEXT, pTimestamps, pMaxDeviation *uint64) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(timestampCount), uintptr(unsafe.Pointer(pTimestampInfos)), uintptr(unsafe.Pointer(pTimestamps)), uintptr(unsafe.Pointer(pMaxDeviation)))
	return Result(ret)
}
func (fn PfnGetCalibratedTimestampsEXT) String() string { return "vkGetCalibratedTimestampsEXT" }

const AMD_shader_core_properties = 1
const AMD_SHADER_CORE_PROPERTIES_SPEC_VERSION = 2

var AMD_SHADER_CORE_PROPERTIES_EXTENSION_NAME = "VK_AMD_shader_core_properties"

// PhysicalDeviceShaderCorePropertiesAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderCorePropertiesAMD.html
type PhysicalDeviceShaderCorePropertiesAMD struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	ShaderEngineCount          uint32
	ShaderArraysPerEngineCount uint32
	ComputeUnitsPerShaderArray uint32
	SimdPerComputeUnit         uint32
	WavefrontsPerSimd          uint32
	WavefrontSize              uint32
	SgprsPerSimd               uint32
	MinSgprAllocation          uint32
	MaxSgprAllocation          uint32
	SgprAllocationGranularity  uint32
	VgprsPerSimd               uint32
	MinVgprAllocation          uint32
	MaxVgprAllocation          uint32
	VgprAllocationGranularity  uint32
}

func NewPhysicalDeviceShaderCorePropertiesAMD() *PhysicalDeviceShaderCorePropertiesAMD {
	p := (*PhysicalDeviceShaderCorePropertiesAMD)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderCorePropertiesAMD)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_AMD
	return p
}
func (p *PhysicalDeviceShaderCorePropertiesAMD) Free() { MemFree(unsafe.Pointer(p)) }

const AMD_memory_overallocation_behavior = 1
const AMD_MEMORY_OVERALLOCATION_BEHAVIOR_SPEC_VERSION = 1

var AMD_MEMORY_OVERALLOCATION_BEHAVIOR_EXTENSION_NAME = "VK_AMD_memory_overallocation_behavior"

// MemoryOverallocationBehaviorAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryOverallocationBehaviorAMD.html
type MemoryOverallocationBehaviorAMD int32

const (
	MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD     MemoryOverallocationBehaviorAMD = 0
	MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD     MemoryOverallocationBehaviorAMD = 1
	MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD  MemoryOverallocationBehaviorAMD = 2
	MEMORY_OVERALLOCATION_BEHAVIOR_BEGIN_RANGE_AMD MemoryOverallocationBehaviorAMD = MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD
	MEMORY_OVERALLOCATION_BEHAVIOR_END_RANGE_AMD   MemoryOverallocationBehaviorAMD = MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD
	MEMORY_OVERALLOCATION_BEHAVIOR_RANGE_SIZE_AMD  MemoryOverallocationBehaviorAMD = (MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD - MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD + 1)
	MEMORY_OVERALLOCATION_BEHAVIOR_MAX_ENUM_AMD    MemoryOverallocationBehaviorAMD = 0x7FFFFFFF
)

func (x MemoryOverallocationBehaviorAMD) String() string {
	switch x {
	case MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD:
		return "MEMORY_OVERALLOCATION_BEHAVIOR_DEFAULT_AMD"
	case MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD:
		return "MEMORY_OVERALLOCATION_BEHAVIOR_ALLOWED_AMD"
	case MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD:
		return "MEMORY_OVERALLOCATION_BEHAVIOR_DISALLOWED_AMD"
	case MEMORY_OVERALLOCATION_BEHAVIOR_MAX_ENUM_AMD:
		return "MEMORY_OVERALLOCATION_BEHAVIOR_MAX_ENUM_AMD"
	default:
		return fmt.Sprint(int32(x))
	}
}

// DeviceMemoryOverallocationCreateInfoAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html
type DeviceMemoryOverallocationCreateInfoAMD struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	OverallocationBehavior MemoryOverallocationBehaviorAMD
}

func NewDeviceMemoryOverallocationCreateInfoAMD() *DeviceMemoryOverallocationCreateInfoAMD {
	p := (*DeviceMemoryOverallocationCreateInfoAMD)(MemAlloc(unsafe.Sizeof(*(*DeviceMemoryOverallocationCreateInfoAMD)(nil))))
	p.SType = STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD
	return p
}
func (p *DeviceMemoryOverallocationCreateInfoAMD) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_vertex_attribute_divisor = 1
const EXT_VERTEX_ATTRIBUTE_DIVISOR_SPEC_VERSION = 3

var EXT_VERTEX_ATTRIBUTE_DIVISOR_EXTENSION_NAME = "VK_EXT_vertex_attribute_divisor"

// PhysicalDeviceVertexAttributeDivisorPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT.html
type PhysicalDeviceVertexAttributeDivisorPropertiesEXT struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	MaxVertexAttribDivisor uint32
}

func NewPhysicalDeviceVertexAttributeDivisorPropertiesEXT() *PhysicalDeviceVertexAttributeDivisorPropertiesEXT {
	p := (*PhysicalDeviceVertexAttributeDivisorPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceVertexAttributeDivisorPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceVertexAttributeDivisorPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// VertexInputBindingDivisorDescriptionEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkVertexInputBindingDivisorDescriptionEXT.html
type VertexInputBindingDivisorDescriptionEXT struct {
	Binding uint32
	Divisor uint32
}

func NewVertexInputBindingDivisorDescriptionEXT() *VertexInputBindingDivisorDescriptionEXT {
	return (*VertexInputBindingDivisorDescriptionEXT)(MemAlloc(unsafe.Sizeof(*(*VertexInputBindingDivisorDescriptionEXT)(nil))))
}
func (p *VertexInputBindingDivisorDescriptionEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineVertexInputDivisorStateCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoEXT.html
type PipelineVertexInputDivisorStateCreateInfoEXT struct {
	SType                     StructureType
	PNext                     unsafe.Pointer
	VertexBindingDivisorCount uint32
	PVertexBindingDivisors    *VertexInputBindingDivisorDescriptionEXT
}

func NewPipelineVertexInputDivisorStateCreateInfoEXT() *PipelineVertexInputDivisorStateCreateInfoEXT {
	p := (*PipelineVertexInputDivisorStateCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineVertexInputDivisorStateCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT
	return p
}
func (p *PipelineVertexInputDivisorStateCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceVertexAttributeDivisorFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT.html
type PhysicalDeviceVertexAttributeDivisorFeaturesEXT struct {
	SType                                  StructureType
	PNext                                  unsafe.Pointer
	VertexAttributeInstanceRateDivisor     Bool32
	VertexAttributeInstanceRateZeroDivisor Bool32
}

func NewPhysicalDeviceVertexAttributeDivisorFeaturesEXT() *PhysicalDeviceVertexAttributeDivisorFeaturesEXT {
	p := (*PhysicalDeviceVertexAttributeDivisorFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceVertexAttributeDivisorFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceVertexAttributeDivisorFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_pipeline_creation_feedback = 1
const EXT_PIPELINE_CREATION_FEEDBACK_SPEC_VERSION = 1

var EXT_PIPELINE_CREATION_FEEDBACK_EXTENSION_NAME = "VK_EXT_pipeline_creation_feedback"

// PipelineCreationFeedbackFlagsEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCreationFeedbackFlagsEXT.html
type PipelineCreationFeedbackFlagsEXT uint32

const (
	PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT                          PipelineCreationFeedbackFlagsEXT = 0x00000001
	PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT PipelineCreationFeedbackFlagsEXT = 0x00000002
	PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT     PipelineCreationFeedbackFlagsEXT = 0x00000004
	PIPELINE_CREATION_FEEDBACK_FLAG_BITS_MAX_ENUM_EXT                 PipelineCreationFeedbackFlagsEXT = 0x7FFFFFFF
)

func (x PipelineCreationFeedbackFlagsEXT) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch PipelineCreationFeedbackFlagsEXT(1 << i) {
			case PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT:
				s += "PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT|"
			case PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT:
				s += "PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT|"
			case PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT:
				s += "PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT|"
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// PipelineCreationFeedbackEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCreationFeedbackEXT.html
type PipelineCreationFeedbackEXT struct {
	Flags    PipelineCreationFeedbackFlagsEXT
	Duration uint64
}

func NewPipelineCreationFeedbackEXT() *PipelineCreationFeedbackEXT {
	return (*PipelineCreationFeedbackEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineCreationFeedbackEXT)(nil))))
}
func (p *PipelineCreationFeedbackEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineCreationFeedbackCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCreationFeedbackCreateInfoEXT.html
type PipelineCreationFeedbackCreateInfoEXT struct {
	SType                              StructureType
	PNext                              unsafe.Pointer
	PPipelineCreationFeedback          *PipelineCreationFeedbackEXT
	PipelineStageCreationFeedbackCount uint32
	PPipelineStageCreationFeedbacks    *PipelineCreationFeedbackEXT
}

func NewPipelineCreationFeedbackCreateInfoEXT() *PipelineCreationFeedbackCreateInfoEXT {
	p := (*PipelineCreationFeedbackCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineCreationFeedbackCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO_EXT
	return p
}
func (p *PipelineCreationFeedbackCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const NV_shader_subgroup_partitioned = 1
const NV_SHADER_SUBGROUP_PARTITIONED_SPEC_VERSION = 1

var NV_SHADER_SUBGROUP_PARTITIONED_EXTENSION_NAME = "VK_NV_shader_subgroup_partitioned"

const NV_compute_shader_derivatives = 1
const NV_COMPUTE_SHADER_DERIVATIVES_SPEC_VERSION = 1

var NV_COMPUTE_SHADER_DERIVATIVES_EXTENSION_NAME = "VK_NV_compute_shader_derivatives"

// PhysicalDeviceComputeShaderDerivativesFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceComputeShaderDerivativesFeaturesNV.html
type PhysicalDeviceComputeShaderDerivativesFeaturesNV struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	ComputeDerivativeGroupQuads  Bool32
	ComputeDerivativeGroupLinear Bool32
}

func NewPhysicalDeviceComputeShaderDerivativesFeaturesNV() *PhysicalDeviceComputeShaderDerivativesFeaturesNV {
	p := (*PhysicalDeviceComputeShaderDerivativesFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceComputeShaderDerivativesFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_COMPUTE_SHADER_DERIVATIVES_FEATURES_NV
	return p
}
func (p *PhysicalDeviceComputeShaderDerivativesFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

const NV_mesh_shader = 1
const NV_MESH_SHADER_SPEC_VERSION = 1

var NV_MESH_SHADER_EXTENSION_NAME = "VK_NV_mesh_shader"

// PhysicalDeviceMeshShaderFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMeshShaderFeaturesNV.html
type PhysicalDeviceMeshShaderFeaturesNV struct {
	SType      StructureType
	PNext      unsafe.Pointer
	TaskShader Bool32
	MeshShader Bool32
}

func NewPhysicalDeviceMeshShaderFeaturesNV() *PhysicalDeviceMeshShaderFeaturesNV {
	p := (*PhysicalDeviceMeshShaderFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMeshShaderFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV
	return p
}
func (p *PhysicalDeviceMeshShaderFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceMeshShaderPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html
type PhysicalDeviceMeshShaderPropertiesNV struct {
	SType                             StructureType
	PNext                             unsafe.Pointer
	MaxDrawMeshTasksCount             uint32
	MaxTaskWorkGroupInvocations       uint32
	MaxTaskWorkGroupSize              [3]uint32
	MaxTaskTotalMemorySize            uint32
	MaxTaskOutputCount                uint32
	MaxMeshWorkGroupInvocations       uint32
	MaxMeshWorkGroupSize              [3]uint32
	MaxMeshTotalMemorySize            uint32
	MaxMeshOutputVertices             uint32
	MaxMeshOutputPrimitives           uint32
	MaxMeshMultiviewViewCount         uint32
	MeshOutputPerVertexGranularity    uint32
	MeshOutputPerPrimitiveGranularity uint32
}

func NewPhysicalDeviceMeshShaderPropertiesNV() *PhysicalDeviceMeshShaderPropertiesNV {
	p := (*PhysicalDeviceMeshShaderPropertiesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMeshShaderPropertiesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV
	return p
}
func (p *PhysicalDeviceMeshShaderPropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

// DrawMeshTasksIndirectCommandNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDrawMeshTasksIndirectCommandNV.html
type DrawMeshTasksIndirectCommandNV struct {
	TaskCount uint32
	FirstTask uint32
}

func NewDrawMeshTasksIndirectCommandNV() *DrawMeshTasksIndirectCommandNV {
	return (*DrawMeshTasksIndirectCommandNV)(MemAlloc(unsafe.Sizeof(*(*DrawMeshTasksIndirectCommandNV)(nil))))
}
func (p *DrawMeshTasksIndirectCommandNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdDrawMeshTasksNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawMeshTasksNV.html
type PfnCmdDrawMeshTasksNV uintptr

func (fn PfnCmdDrawMeshTasksNV) Call(commandBuffer CommandBuffer, taskCount, firstTask uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(taskCount), uintptr(firstTask))
}
func (fn PfnCmdDrawMeshTasksNV) String() string { return "vkCmdDrawMeshTasksNV" }

//  PfnCmdDrawMeshTasksIndirectNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawMeshTasksIndirectNV.html
type PfnCmdDrawMeshTasksIndirectNV uintptr

func (fn PfnCmdDrawMeshTasksIndirectNV) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(drawCount), uintptr(stride))
}
func (fn PfnCmdDrawMeshTasksIndirectNV) String() string { return "vkCmdDrawMeshTasksIndirectNV" }

//  PfnCmdDrawMeshTasksIndirectCountNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdDrawMeshTasksIndirectCountNV.html
type PfnCmdDrawMeshTasksIndirectCountNV uintptr

func (fn PfnCmdDrawMeshTasksIndirectCountNV) Call(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount, stride uint32) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(buffer), uintptr(offset), uintptr(countBuffer), uintptr(countBufferOffset), uintptr(maxDrawCount), uintptr(stride))
}
func (fn PfnCmdDrawMeshTasksIndirectCountNV) String() string {
	return "vkCmdDrawMeshTasksIndirectCountNV"
}

const NV_fragment_shader_barycentric = 1
const NV_FRAGMENT_SHADER_BARYCENTRIC_SPEC_VERSION = 1

var NV_FRAGMENT_SHADER_BARYCENTRIC_EXTENSION_NAME = "VK_NV_fragment_shader_barycentric"

// PhysicalDeviceFragmentShaderBarycentricFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV.html
type PhysicalDeviceFragmentShaderBarycentricFeaturesNV struct {
	SType                     StructureType
	PNext                     unsafe.Pointer
	FragmentShaderBarycentric Bool32
}

func NewPhysicalDeviceFragmentShaderBarycentricFeaturesNV() *PhysicalDeviceFragmentShaderBarycentricFeaturesNV {
	p := (*PhysicalDeviceFragmentShaderBarycentricFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceFragmentShaderBarycentricFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_FEATURES_NV
	return p
}
func (p *PhysicalDeviceFragmentShaderBarycentricFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

const NV_shader_image_footprint = 1
const NV_SHADER_IMAGE_FOOTPRINT_SPEC_VERSION = 2

var NV_SHADER_IMAGE_FOOTPRINT_EXTENSION_NAME = "VK_NV_shader_image_footprint"

// PhysicalDeviceShaderImageFootprintFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderImageFootprintFeaturesNV.html
type PhysicalDeviceShaderImageFootprintFeaturesNV struct {
	SType          StructureType
	PNext          unsafe.Pointer
	ImageFootprint Bool32
}

func NewPhysicalDeviceShaderImageFootprintFeaturesNV() *PhysicalDeviceShaderImageFootprintFeaturesNV {
	p := (*PhysicalDeviceShaderImageFootprintFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderImageFootprintFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_FOOTPRINT_FEATURES_NV
	return p
}
func (p *PhysicalDeviceShaderImageFootprintFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

const NV_scissor_exclusive = 1
const NV_SCISSOR_EXCLUSIVE_SPEC_VERSION = 1

var NV_SCISSOR_EXCLUSIVE_EXTENSION_NAME = "VK_NV_scissor_exclusive"

// PipelineViewportExclusiveScissorStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html
type PipelineViewportExclusiveScissorStateCreateInfoNV struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	ExclusiveScissorCount uint32
	PExclusiveScissors    *Rect2D
}

func NewPipelineViewportExclusiveScissorStateCreateInfoNV() *PipelineViewportExclusiveScissorStateCreateInfoNV {
	p := (*PipelineViewportExclusiveScissorStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineViewportExclusiveScissorStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineViewportExclusiveScissorStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceExclusiveScissorFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceExclusiveScissorFeaturesNV.html
type PhysicalDeviceExclusiveScissorFeaturesNV struct {
	SType            StructureType
	PNext            unsafe.Pointer
	ExclusiveScissor Bool32
}

func NewPhysicalDeviceExclusiveScissorFeaturesNV() *PhysicalDeviceExclusiveScissorFeaturesNV {
	p := (*PhysicalDeviceExclusiveScissorFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceExclusiveScissorFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV
	return p
}
func (p *PhysicalDeviceExclusiveScissorFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdSetExclusiveScissorNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetExclusiveScissorNV.html
type PfnCmdSetExclusiveScissorNV uintptr

func (fn PfnCmdSetExclusiveScissorNV) Call(commandBuffer CommandBuffer, firstExclusiveScissor, exclusiveScissorCount uint32, pExclusiveScissors *Rect2D) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(firstExclusiveScissor), uintptr(exclusiveScissorCount), uintptr(unsafe.Pointer(pExclusiveScissors)))
}
func (fn PfnCmdSetExclusiveScissorNV) String() string { return "vkCmdSetExclusiveScissorNV" }

const NV_device_diagnostic_checkpoints = 1
const NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_SPEC_VERSION = 2

var NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_EXTENSION_NAME = "VK_NV_device_diagnostic_checkpoints"

// QueueFamilyCheckpointPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueueFamilyCheckpointPropertiesNV.html
type QueueFamilyCheckpointPropertiesNV struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	CheckpointExecutionStageMask PipelineStageFlags
}

func NewQueueFamilyCheckpointPropertiesNV() *QueueFamilyCheckpointPropertiesNV {
	p := (*QueueFamilyCheckpointPropertiesNV)(MemAlloc(unsafe.Sizeof(*(*QueueFamilyCheckpointPropertiesNV)(nil))))
	p.SType = STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV
	return p
}
func (p *QueueFamilyCheckpointPropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

// CheckpointDataNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCheckpointDataNV.html
type CheckpointDataNV struct {
	SType             StructureType
	PNext             unsafe.Pointer
	Stage             PipelineStageFlags
	PCheckpointMarker unsafe.Pointer
}

func NewCheckpointDataNV() *CheckpointDataNV {
	p := (*CheckpointDataNV)(MemAlloc(unsafe.Sizeof(*(*CheckpointDataNV)(nil))))
	p.SType = STRUCTURE_TYPE_CHECKPOINT_DATA_NV
	return p
}
func (p *CheckpointDataNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdSetCheckpointNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetCheckpointNV.html
type PfnCmdSetCheckpointNV uintptr

func (fn PfnCmdSetCheckpointNV) Call(commandBuffer CommandBuffer, pCheckpointMarker unsafe.Pointer) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(pCheckpointMarker))
}
func (fn PfnCmdSetCheckpointNV) String() string { return "vkCmdSetCheckpointNV" }

//  PfnGetQueueCheckpointDataNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetQueueCheckpointDataNV.html
type PfnGetQueueCheckpointDataNV uintptr

func (fn PfnGetQueueCheckpointDataNV) Call(queue Queue, pCheckpointDataCount *uint32, pCheckpointData *CheckpointDataNV) {
	_, _, _ = call(uintptr(fn), uintptr(queue), uintptr(unsafe.Pointer(pCheckpointDataCount)), uintptr(unsafe.Pointer(pCheckpointData)))
}
func (fn PfnGetQueueCheckpointDataNV) String() string { return "vkGetQueueCheckpointDataNV" }

const INTEL_shader_integer_functions2 = 1
const INTEL_SHADER_INTEGER_FUNCTIONS_2_SPEC_VERSION = 1

var INTEL_SHADER_INTEGER_FUNCTIONS_2_EXTENSION_NAME = "VK_INTEL_shader_integer_functions2"

// PhysicalDeviceShaderIntegerFunctions2FeaturesINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL.html
type PhysicalDeviceShaderIntegerFunctions2FeaturesINTEL struct {
	SType                   StructureType
	PNext                   unsafe.Pointer
	ShaderIntegerFunctions2 Bool32
}

func NewPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL() *PhysicalDeviceShaderIntegerFunctions2FeaturesINTEL {
	return (*PhysicalDeviceShaderIntegerFunctions2FeaturesINTEL)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderIntegerFunctions2FeaturesINTEL)(nil))))
}
func (p *PhysicalDeviceShaderIntegerFunctions2FeaturesINTEL) Free() { MemFree(unsafe.Pointer(p)) }

const INTEL_performance_query = 1

// PerformanceConfigurationINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceConfigurationINTEL.html
type PerformanceConfigurationINTEL NonDispatchableHandle

const INTEL_PERFORMANCE_QUERY_SPEC_VERSION = 1

var INTEL_PERFORMANCE_QUERY_EXTENSION_NAME = "VK_INTEL_performance_query"

// PerformanceConfigurationTypeINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceConfigurationTypeINTEL.html
type PerformanceConfigurationTypeINTEL int32

const (
	PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL PerformanceConfigurationTypeINTEL = 0
	PERFORMANCE_CONFIGURATION_TYPE_BEGIN_RANGE_INTEL                               PerformanceConfigurationTypeINTEL = PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL
	PERFORMANCE_CONFIGURATION_TYPE_END_RANGE_INTEL                                 PerformanceConfigurationTypeINTEL = PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL
	PERFORMANCE_CONFIGURATION_TYPE_RANGE_SIZE_INTEL                                PerformanceConfigurationTypeINTEL = (PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL - PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL + 1)
	PERFORMANCE_CONFIGURATION_TYPE_MAX_ENUM_INTEL                                  PerformanceConfigurationTypeINTEL = 0x7FFFFFFF
)

func (x PerformanceConfigurationTypeINTEL) String() string {
	switch x {
	case PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL:
		return "PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL"
	case PERFORMANCE_CONFIGURATION_TYPE_MAX_ENUM_INTEL:
		return "PERFORMANCE_CONFIGURATION_TYPE_MAX_ENUM_INTEL"
	default:
		return fmt.Sprint(int32(x))
	}
}

// QueryPoolSamplingModeINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryPoolSamplingModeINTEL.html
type QueryPoolSamplingModeINTEL int32

const (
	QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL      QueryPoolSamplingModeINTEL = 0
	QUERY_POOL_SAMPLING_MODE_BEGIN_RANGE_INTEL QueryPoolSamplingModeINTEL = QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL
	QUERY_POOL_SAMPLING_MODE_END_RANGE_INTEL   QueryPoolSamplingModeINTEL = QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL
	QUERY_POOL_SAMPLING_MODE_RANGE_SIZE_INTEL  QueryPoolSamplingModeINTEL = (QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL - QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL + 1)
	QUERY_POOL_SAMPLING_MODE_MAX_ENUM_INTEL    QueryPoolSamplingModeINTEL = 0x7FFFFFFF
)

func (x QueryPoolSamplingModeINTEL) String() string {
	switch x {
	case QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL:
		return "QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL"
	case QUERY_POOL_SAMPLING_MODE_MAX_ENUM_INTEL:
		return "QUERY_POOL_SAMPLING_MODE_MAX_ENUM_INTEL"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PerformanceOverrideTypeINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceOverrideTypeINTEL.html
type PerformanceOverrideTypeINTEL int32

const (
	PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL    PerformanceOverrideTypeINTEL = 0
	PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL PerformanceOverrideTypeINTEL = 1
	PERFORMANCE_OVERRIDE_TYPE_BEGIN_RANGE_INTEL      PerformanceOverrideTypeINTEL = PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL
	PERFORMANCE_OVERRIDE_TYPE_END_RANGE_INTEL        PerformanceOverrideTypeINTEL = PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL
	PERFORMANCE_OVERRIDE_TYPE_RANGE_SIZE_INTEL       PerformanceOverrideTypeINTEL = (PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL - PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL + 1)
	PERFORMANCE_OVERRIDE_TYPE_MAX_ENUM_INTEL         PerformanceOverrideTypeINTEL = 0x7FFFFFFF
)

func (x PerformanceOverrideTypeINTEL) String() string {
	switch x {
	case PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL:
		return "PERFORMANCE_OVERRIDE_TYPE_NULL_HARDWARE_INTEL"
	case PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL:
		return "PERFORMANCE_OVERRIDE_TYPE_FLUSH_GPU_CACHES_INTEL"
	case PERFORMANCE_OVERRIDE_TYPE_MAX_ENUM_INTEL:
		return "PERFORMANCE_OVERRIDE_TYPE_MAX_ENUM_INTEL"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PerformanceParameterTypeINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceParameterTypeINTEL.html
type PerformanceParameterTypeINTEL int32

const (
	PERFORMANCE_PARAMETER_TYPE_HW_COUNTERS_SUPPORTED_INTEL    PerformanceParameterTypeINTEL = 0
	PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL PerformanceParameterTypeINTEL = 1
	PERFORMANCE_PARAMETER_TYPE_BEGIN_RANGE_INTEL              PerformanceParameterTypeINTEL = PERFORMANCE_PARAMETER_TYPE_HW_COUNTERS_SUPPORTED_INTEL
	PERFORMANCE_PARAMETER_TYPE_END_RANGE_INTEL                PerformanceParameterTypeINTEL = PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL
	PERFORMANCE_PARAMETER_TYPE_RANGE_SIZE_INTEL               PerformanceParameterTypeINTEL = (PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL - PERFORMANCE_PARAMETER_TYPE_HW_COUNTERS_SUPPORTED_INTEL + 1)
	PERFORMANCE_PARAMETER_TYPE_MAX_ENUM_INTEL                 PerformanceParameterTypeINTEL = 0x7FFFFFFF
)

func (x PerformanceParameterTypeINTEL) String() string {
	switch x {
	case PERFORMANCE_PARAMETER_TYPE_HW_COUNTERS_SUPPORTED_INTEL:
		return "PERFORMANCE_PARAMETER_TYPE_HW_COUNTERS_SUPPORTED_INTEL"
	case PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL:
		return "PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL"
	case PERFORMANCE_PARAMETER_TYPE_MAX_ENUM_INTEL:
		return "PERFORMANCE_PARAMETER_TYPE_MAX_ENUM_INTEL"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PerformanceValueTypeINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceValueTypeINTEL.html
type PerformanceValueTypeINTEL int32

const (
	PERFORMANCE_VALUE_TYPE_UINT32_INTEL      PerformanceValueTypeINTEL = 0
	PERFORMANCE_VALUE_TYPE_UINT64_INTEL      PerformanceValueTypeINTEL = 1
	PERFORMANCE_VALUE_TYPE_FLOAT_INTEL       PerformanceValueTypeINTEL = 2
	PERFORMANCE_VALUE_TYPE_BOOL_INTEL        PerformanceValueTypeINTEL = 3
	PERFORMANCE_VALUE_TYPE_STRING_INTEL      PerformanceValueTypeINTEL = 4
	PERFORMANCE_VALUE_TYPE_BEGIN_RANGE_INTEL PerformanceValueTypeINTEL = PERFORMANCE_VALUE_TYPE_UINT32_INTEL
	PERFORMANCE_VALUE_TYPE_END_RANGE_INTEL   PerformanceValueTypeINTEL = PERFORMANCE_VALUE_TYPE_STRING_INTEL
	PERFORMANCE_VALUE_TYPE_RANGE_SIZE_INTEL  PerformanceValueTypeINTEL = (PERFORMANCE_VALUE_TYPE_STRING_INTEL - PERFORMANCE_VALUE_TYPE_UINT32_INTEL + 1)
	PERFORMANCE_VALUE_TYPE_MAX_ENUM_INTEL    PerformanceValueTypeINTEL = 0x7FFFFFFF
)

func (x PerformanceValueTypeINTEL) String() string {
	switch x {
	case PERFORMANCE_VALUE_TYPE_UINT32_INTEL:
		return "PERFORMANCE_VALUE_TYPE_UINT32_INTEL"
	case PERFORMANCE_VALUE_TYPE_UINT64_INTEL:
		return "PERFORMANCE_VALUE_TYPE_UINT64_INTEL"
	case PERFORMANCE_VALUE_TYPE_FLOAT_INTEL:
		return "PERFORMANCE_VALUE_TYPE_FLOAT_INTEL"
	case PERFORMANCE_VALUE_TYPE_BOOL_INTEL:
		return "PERFORMANCE_VALUE_TYPE_BOOL_INTEL"
	case PERFORMANCE_VALUE_TYPE_STRING_INTEL:
		return "PERFORMANCE_VALUE_TYPE_STRING_INTEL"
	case PERFORMANCE_VALUE_TYPE_MAX_ENUM_INTEL:
		return "PERFORMANCE_VALUE_TYPE_MAX_ENUM_INTEL"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PerformanceValueINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceValueINTEL.html
type PerformanceValueINTEL struct {
	Type PerformanceValueTypeINTEL
	Data PerformanceValueDataINTEL
}

func NewPerformanceValueINTEL() *PerformanceValueINTEL {
	return (*PerformanceValueINTEL)(MemAlloc(unsafe.Sizeof(*(*PerformanceValueINTEL)(nil))))
}
func (p *PerformanceValueINTEL) Free() { MemFree(unsafe.Pointer(p)) }

// InitializePerformanceApiInfoINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkInitializePerformanceApiInfoINTEL.html
type InitializePerformanceApiInfoINTEL struct {
	SType     StructureType
	PNext     unsafe.Pointer
	PUserData unsafe.Pointer
}

func NewInitializePerformanceApiInfoINTEL() *InitializePerformanceApiInfoINTEL {
	p := (*InitializePerformanceApiInfoINTEL)(MemAlloc(unsafe.Sizeof(*(*InitializePerformanceApiInfoINTEL)(nil))))
	p.SType = STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL
	return p
}
func (p *InitializePerformanceApiInfoINTEL) Free() { MemFree(unsafe.Pointer(p)) }

// QueryPoolCreateInfoINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkQueryPoolCreateInfoINTEL.html
type QueryPoolCreateInfoINTEL struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	PerformanceCountersSampling QueryPoolSamplingModeINTEL
}

func NewQueryPoolCreateInfoINTEL() *QueryPoolCreateInfoINTEL {
	p := (*QueryPoolCreateInfoINTEL)(MemAlloc(unsafe.Sizeof(*(*QueryPoolCreateInfoINTEL)(nil))))
	p.SType = STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL
	return p
}
func (p *QueryPoolCreateInfoINTEL) Free() { MemFree(unsafe.Pointer(p)) }

// PerformanceMarkerInfoINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceMarkerInfoINTEL.html
type PerformanceMarkerInfoINTEL struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Marker uint64
}

func NewPerformanceMarkerInfoINTEL() *PerformanceMarkerInfoINTEL {
	p := (*PerformanceMarkerInfoINTEL)(MemAlloc(unsafe.Sizeof(*(*PerformanceMarkerInfoINTEL)(nil))))
	p.SType = STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL
	return p
}
func (p *PerformanceMarkerInfoINTEL) Free() { MemFree(unsafe.Pointer(p)) }

// PerformanceStreamMarkerInfoINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceStreamMarkerInfoINTEL.html
type PerformanceStreamMarkerInfoINTEL struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Marker uint32
}

func NewPerformanceStreamMarkerInfoINTEL() *PerformanceStreamMarkerInfoINTEL {
	p := (*PerformanceStreamMarkerInfoINTEL)(MemAlloc(unsafe.Sizeof(*(*PerformanceStreamMarkerInfoINTEL)(nil))))
	p.SType = STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL
	return p
}
func (p *PerformanceStreamMarkerInfoINTEL) Free() { MemFree(unsafe.Pointer(p)) }

// PerformanceOverrideInfoINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceOverrideInfoINTEL.html
type PerformanceOverrideInfoINTEL struct {
	SType     StructureType
	PNext     unsafe.Pointer
	Type      PerformanceOverrideTypeINTEL
	Enable    Bool32
	Parameter uint64
}

func NewPerformanceOverrideInfoINTEL() *PerformanceOverrideInfoINTEL {
	p := (*PerformanceOverrideInfoINTEL)(MemAlloc(unsafe.Sizeof(*(*PerformanceOverrideInfoINTEL)(nil))))
	p.SType = STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL
	return p
}
func (p *PerformanceOverrideInfoINTEL) Free() { MemFree(unsafe.Pointer(p)) }

// PerformanceConfigurationAcquireInfoINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html
type PerformanceConfigurationAcquireInfoINTEL struct {
	SType StructureType
	PNext unsafe.Pointer
	Type  PerformanceConfigurationTypeINTEL
}

func NewPerformanceConfigurationAcquireInfoINTEL() *PerformanceConfigurationAcquireInfoINTEL {
	p := (*PerformanceConfigurationAcquireInfoINTEL)(MemAlloc(unsafe.Sizeof(*(*PerformanceConfigurationAcquireInfoINTEL)(nil))))
	p.SType = STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL
	return p
}
func (p *PerformanceConfigurationAcquireInfoINTEL) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnInitializePerformanceApiINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkInitializePerformanceApiINTEL.html
type PfnInitializePerformanceApiINTEL uintptr

func (fn PfnInitializePerformanceApiINTEL) Call(device Device, pInitializeInfo *InitializePerformanceApiInfoINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInitializeInfo)))
	return Result(ret)
}
func (fn PfnInitializePerformanceApiINTEL) String() string { return "vkInitializePerformanceApiINTEL" }

//  PfnUninitializePerformanceApiINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkUninitializePerformanceApiINTEL.html
type PfnUninitializePerformanceApiINTEL uintptr

func (fn PfnUninitializePerformanceApiINTEL) Call(device Device) {
	_, _, _ = call(uintptr(fn), uintptr(device))
}
func (fn PfnUninitializePerformanceApiINTEL) String() string {
	return "vkUninitializePerformanceApiINTEL"
}

//  PfnCmdSetPerformanceMarkerINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetPerformanceMarkerINTEL.html
type PfnCmdSetPerformanceMarkerINTEL uintptr

func (fn PfnCmdSetPerformanceMarkerINTEL) Call(commandBuffer CommandBuffer, pMarkerInfo *PerformanceMarkerInfoINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)))
	return Result(ret)
}
func (fn PfnCmdSetPerformanceMarkerINTEL) String() string { return "vkCmdSetPerformanceMarkerINTEL" }

//  PfnCmdSetPerformanceStreamMarkerINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetPerformanceStreamMarkerINTEL.html
type PfnCmdSetPerformanceStreamMarkerINTEL uintptr

func (fn PfnCmdSetPerformanceStreamMarkerINTEL) Call(commandBuffer CommandBuffer, pMarkerInfo *PerformanceStreamMarkerInfoINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pMarkerInfo)))
	return Result(ret)
}
func (fn PfnCmdSetPerformanceStreamMarkerINTEL) String() string {
	return "vkCmdSetPerformanceStreamMarkerINTEL"
}

//  PfnCmdSetPerformanceOverrideINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetPerformanceOverrideINTEL.html
type PfnCmdSetPerformanceOverrideINTEL uintptr

func (fn PfnCmdSetPerformanceOverrideINTEL) Call(commandBuffer CommandBuffer, pOverrideInfo *PerformanceOverrideInfoINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(commandBuffer), uintptr(unsafe.Pointer(pOverrideInfo)))
	return Result(ret)
}
func (fn PfnCmdSetPerformanceOverrideINTEL) String() string {
	return "vkCmdSetPerformanceOverrideINTEL"
}

//  PfnAcquirePerformanceConfigurationINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAcquirePerformanceConfigurationINTEL.html
type PfnAcquirePerformanceConfigurationINTEL uintptr

func (fn PfnAcquirePerformanceConfigurationINTEL) Call(device Device, pAcquireInfo *PerformanceConfigurationAcquireInfoINTEL, pConfiguration *PerformanceConfigurationINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pAcquireInfo)), uintptr(unsafe.Pointer(pConfiguration)))
	return Result(ret)
}
func (fn PfnAcquirePerformanceConfigurationINTEL) String() string {
	return "vkAcquirePerformanceConfigurationINTEL"
}

//  PfnReleasePerformanceConfigurationINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkReleasePerformanceConfigurationINTEL.html
type PfnReleasePerformanceConfigurationINTEL uintptr

func (fn PfnReleasePerformanceConfigurationINTEL) Call(device Device, configuration PerformanceConfigurationINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(configuration))
	return Result(ret)
}
func (fn PfnReleasePerformanceConfigurationINTEL) String() string {
	return "vkReleasePerformanceConfigurationINTEL"
}

//  PfnQueueSetPerformanceConfigurationINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkQueueSetPerformanceConfigurationINTEL.html
type PfnQueueSetPerformanceConfigurationINTEL uintptr

func (fn PfnQueueSetPerformanceConfigurationINTEL) Call(queue Queue, configuration PerformanceConfigurationINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(queue), uintptr(configuration))
	return Result(ret)
}
func (fn PfnQueueSetPerformanceConfigurationINTEL) String() string {
	return "vkQueueSetPerformanceConfigurationINTEL"
}

//  PfnGetPerformanceParameterINTEL -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPerformanceParameterINTEL.html
type PfnGetPerformanceParameterINTEL uintptr

func (fn PfnGetPerformanceParameterINTEL) Call(device Device, parameter PerformanceParameterTypeINTEL, pValue *PerformanceValueINTEL) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(parameter), uintptr(unsafe.Pointer(pValue)))
	return Result(ret)
}
func (fn PfnGetPerformanceParameterINTEL) String() string { return "vkGetPerformanceParameterINTEL" }

const EXT_pci_bus_info = 1
const EXT_PCI_BUS_INFO_SPEC_VERSION = 2

var EXT_PCI_BUS_INFO_EXTENSION_NAME = "VK_EXT_pci_bus_info"

// PhysicalDevicePCIBusInfoPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDevicePCIBusInfoPropertiesEXT.html
type PhysicalDevicePCIBusInfoPropertiesEXT struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PciDomain   uint32
	PciBus      uint32
	PciDevice   uint32
	PciFunction uint32
}

func NewPhysicalDevicePCIBusInfoPropertiesEXT() *PhysicalDevicePCIBusInfoPropertiesEXT {
	p := (*PhysicalDevicePCIBusInfoPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDevicePCIBusInfoPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT
	return p
}
func (p *PhysicalDevicePCIBusInfoPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const AMD_display_native_hdr = 1
const AMD_DISPLAY_NATIVE_HDR_SPEC_VERSION = 1

var AMD_DISPLAY_NATIVE_HDR_EXTENSION_NAME = "VK_AMD_display_native_hdr"

// DisplayNativeHdrSurfaceCapabilitiesAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html
type DisplayNativeHdrSurfaceCapabilitiesAMD struct {
	SType               StructureType
	PNext               unsafe.Pointer
	LocalDimmingSupport Bool32
}

func NewDisplayNativeHdrSurfaceCapabilitiesAMD() *DisplayNativeHdrSurfaceCapabilitiesAMD {
	p := (*DisplayNativeHdrSurfaceCapabilitiesAMD)(MemAlloc(unsafe.Sizeof(*(*DisplayNativeHdrSurfaceCapabilitiesAMD)(nil))))
	p.SType = STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD
	return p
}
func (p *DisplayNativeHdrSurfaceCapabilitiesAMD) Free() { MemFree(unsafe.Pointer(p)) }

// SwapchainDisplayNativeHdrCreateInfoAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSwapchainDisplayNativeHdrCreateInfoAMD.html
type SwapchainDisplayNativeHdrCreateInfoAMD struct {
	SType              StructureType
	PNext              unsafe.Pointer
	LocalDimmingEnable Bool32
}

func NewSwapchainDisplayNativeHdrCreateInfoAMD() *SwapchainDisplayNativeHdrCreateInfoAMD {
	p := (*SwapchainDisplayNativeHdrCreateInfoAMD)(MemAlloc(unsafe.Sizeof(*(*SwapchainDisplayNativeHdrCreateInfoAMD)(nil))))
	p.SType = STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD
	return p
}
func (p *SwapchainDisplayNativeHdrCreateInfoAMD) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnSetLocalDimmingAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkSetLocalDimmingAMD.html
type PfnSetLocalDimmingAMD uintptr

func (fn PfnSetLocalDimmingAMD) Call(device Device, swapChain SwapchainKHR, localDimmingEnable Bool32) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(swapChain), uintptr(localDimmingEnable))
}
func (fn PfnSetLocalDimmingAMD) String() string { return "vkSetLocalDimmingAMD" }

const EXT_fragment_density_map = 1
const EXT_FRAGMENT_DENSITY_MAP_SPEC_VERSION = 1

var EXT_FRAGMENT_DENSITY_MAP_EXTENSION_NAME = "VK_EXT_fragment_density_map"

// PhysicalDeviceFragmentDensityMapFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceFragmentDensityMapFeaturesEXT.html
type PhysicalDeviceFragmentDensityMapFeaturesEXT struct {
	SType                                 StructureType
	PNext                                 unsafe.Pointer
	FragmentDensityMap                    Bool32
	FragmentDensityMapDynamic             Bool32
	FragmentDensityMapNonSubsampledImages Bool32
}

func NewPhysicalDeviceFragmentDensityMapFeaturesEXT() *PhysicalDeviceFragmentDensityMapFeaturesEXT {
	p := (*PhysicalDeviceFragmentDensityMapFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceFragmentDensityMapFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceFragmentDensityMapFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceFragmentDensityMapPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceFragmentDensityMapPropertiesEXT.html
type PhysicalDeviceFragmentDensityMapPropertiesEXT struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	MinFragmentDensityTexelSize Extent2D
	MaxFragmentDensityTexelSize Extent2D
	FragmentDensityInvocations  Bool32
}

func NewPhysicalDeviceFragmentDensityMapPropertiesEXT() *PhysicalDeviceFragmentDensityMapPropertiesEXT {
	p := (*PhysicalDeviceFragmentDensityMapPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceFragmentDensityMapPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceFragmentDensityMapPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// RenderPassFragmentDensityMapCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html
type RenderPassFragmentDensityMapCreateInfoEXT struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	FragmentDensityMapAttachment AttachmentReference
}

func NewRenderPassFragmentDensityMapCreateInfoEXT() *RenderPassFragmentDensityMapCreateInfoEXT {
	p := (*RenderPassFragmentDensityMapCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*RenderPassFragmentDensityMapCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT
	return p
}
func (p *RenderPassFragmentDensityMapCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_scalar_block_layout = 1
const EXT_SCALAR_BLOCK_LAYOUT_SPEC_VERSION = 1

var EXT_SCALAR_BLOCK_LAYOUT_EXTENSION_NAME = "VK_EXT_scalar_block_layout"

// PhysicalDeviceScalarBlockLayoutFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceScalarBlockLayoutFeaturesEXT.html
type PhysicalDeviceScalarBlockLayoutFeaturesEXT struct {
	SType             StructureType
	PNext             unsafe.Pointer
	ScalarBlockLayout Bool32
}

func NewPhysicalDeviceScalarBlockLayoutFeaturesEXT() *PhysicalDeviceScalarBlockLayoutFeaturesEXT {
	p := (*PhysicalDeviceScalarBlockLayoutFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceScalarBlockLayoutFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceScalarBlockLayoutFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const GOOGLE_hlsl_functionality1 = 1
const GOOGLE_HLSL_FUNCTIONALITY1_SPEC_VERSION = 1

var GOOGLE_HLSL_FUNCTIONALITY1_EXTENSION_NAME = "VK_GOOGLE_hlsl_functionality1"

const GOOGLE_decorate_string = 1
const GOOGLE_DECORATE_STRING_SPEC_VERSION = 1

var GOOGLE_DECORATE_STRING_EXTENSION_NAME = "VK_GOOGLE_decorate_string"

const EXT_subgroup_size_control = 1
const EXT_SUBGROUP_SIZE_CONTROL_SPEC_VERSION = 2

var EXT_SUBGROUP_SIZE_CONTROL_EXTENSION_NAME = "VK_EXT_subgroup_size_control"

// PhysicalDeviceSubgroupSizeControlFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlFeaturesEXT.html
type PhysicalDeviceSubgroupSizeControlFeaturesEXT struct {
	SType                StructureType
	PNext                unsafe.Pointer
	SubgroupSizeControl  Bool32
	ComputeFullSubgroups Bool32
}

func NewPhysicalDeviceSubgroupSizeControlFeaturesEXT() *PhysicalDeviceSubgroupSizeControlFeaturesEXT {
	return (*PhysicalDeviceSubgroupSizeControlFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSubgroupSizeControlFeaturesEXT)(nil))))
}
func (p *PhysicalDeviceSubgroupSizeControlFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceSubgroupSizeControlPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlPropertiesEXT.html
type PhysicalDeviceSubgroupSizeControlPropertiesEXT struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	MinSubgroupSize              uint32
	MaxSubgroupSize              uint32
	MaxComputeWorkgroupSubgroups uint32
	RequiredSubgroupSizeStages   ShaderStageFlags
}

func NewPhysicalDeviceSubgroupSizeControlPropertiesEXT() *PhysicalDeviceSubgroupSizeControlPropertiesEXT {
	return (*PhysicalDeviceSubgroupSizeControlPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceSubgroupSizeControlPropertiesEXT)(nil))))
}
func (p *PhysicalDeviceSubgroupSizeControlPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineShaderStageRequiredSubgroupSizeCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT.html
type PipelineShaderStageRequiredSubgroupSizeCreateInfoEXT struct {
	SType                StructureType
	PNext                unsafe.Pointer
	RequiredSubgroupSize uint32
}

func NewPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT() *PipelineShaderStageRequiredSubgroupSizeCreateInfoEXT {
	return (*PipelineShaderStageRequiredSubgroupSizeCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineShaderStageRequiredSubgroupSizeCreateInfoEXT)(nil))))
}
func (p *PipelineShaderStageRequiredSubgroupSizeCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const AMD_shader_core_properties2 = 1
const AMD_SHADER_CORE_PROPERTIES_2_SPEC_VERSION = 1

var AMD_SHADER_CORE_PROPERTIES_2_EXTENSION_NAME = "VK_AMD_shader_core_properties2"

// ShaderCorePropertiesFlagsAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkShaderCorePropertiesFlagsAMD.html
type ShaderCorePropertiesFlagsAMD uint32

const (
	SHADER_CORE_PROPERTIES_FLAG_BITS_MAX_ENUM_AMD ShaderCorePropertiesFlagsAMD = 0x7FFFFFFF
)

func (x ShaderCorePropertiesFlagsAMD) String() string {
	var s string
	for i := uint32(0); i < 32; i++ {
		if int32(x)&(1<<i) != 0 {
			switch ShaderCorePropertiesFlagsAMD(1 << i) {
			}
		}
	}
	return strings.TrimSuffix(s, `|`)
}

// PhysicalDeviceShaderCoreProperties2AMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderCoreProperties2AMD.html
type PhysicalDeviceShaderCoreProperties2AMD struct {
	SType                  StructureType
	PNext                  unsafe.Pointer
	ShaderCoreFeatures     ShaderCorePropertiesFlagsAMD
	ActiveComputeUnitCount uint32
}

func NewPhysicalDeviceShaderCoreProperties2AMD() *PhysicalDeviceShaderCoreProperties2AMD {
	return (*PhysicalDeviceShaderCoreProperties2AMD)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderCoreProperties2AMD)(nil))))
}
func (p *PhysicalDeviceShaderCoreProperties2AMD) Free() { MemFree(unsafe.Pointer(p)) }

const AMD_device_coherent_memory = 1
const AMD_DEVICE_COHERENT_MEMORY_SPEC_VERSION = 1

var AMD_DEVICE_COHERENT_MEMORY_EXTENSION_NAME = "VK_AMD_device_coherent_memory"

// PhysicalDeviceCoherentMemoryFeaturesAMD -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceCoherentMemoryFeaturesAMD.html
type PhysicalDeviceCoherentMemoryFeaturesAMD struct {
	SType                StructureType
	PNext                unsafe.Pointer
	DeviceCoherentMemory Bool32
}

func NewPhysicalDeviceCoherentMemoryFeaturesAMD() *PhysicalDeviceCoherentMemoryFeaturesAMD {
	return (*PhysicalDeviceCoherentMemoryFeaturesAMD)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceCoherentMemoryFeaturesAMD)(nil))))
}
func (p *PhysicalDeviceCoherentMemoryFeaturesAMD) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_memory_budget = 1
const EXT_MEMORY_BUDGET_SPEC_VERSION = 1

var EXT_MEMORY_BUDGET_EXTENSION_NAME = "VK_EXT_memory_budget"

// PhysicalDeviceMemoryBudgetPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMemoryBudgetPropertiesEXT.html
type PhysicalDeviceMemoryBudgetPropertiesEXT struct {
	SType      StructureType
	PNext      unsafe.Pointer
	HeapBudget [MAX_MEMORY_HEAPS]DeviceSize
	HeapUsage  [MAX_MEMORY_HEAPS]DeviceSize
}

func NewPhysicalDeviceMemoryBudgetPropertiesEXT() *PhysicalDeviceMemoryBudgetPropertiesEXT {
	p := (*PhysicalDeviceMemoryBudgetPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMemoryBudgetPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceMemoryBudgetPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_memory_priority = 1
const EXT_MEMORY_PRIORITY_SPEC_VERSION = 1

var EXT_MEMORY_PRIORITY_EXTENSION_NAME = "VK_EXT_memory_priority"

// PhysicalDeviceMemoryPriorityFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceMemoryPriorityFeaturesEXT.html
type PhysicalDeviceMemoryPriorityFeaturesEXT struct {
	SType          StructureType
	PNext          unsafe.Pointer
	MemoryPriority Bool32
}

func NewPhysicalDeviceMemoryPriorityFeaturesEXT() *PhysicalDeviceMemoryPriorityFeaturesEXT {
	p := (*PhysicalDeviceMemoryPriorityFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceMemoryPriorityFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceMemoryPriorityFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryPriorityAllocateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryPriorityAllocateInfoEXT.html
type MemoryPriorityAllocateInfoEXT struct {
	SType    StructureType
	PNext    unsafe.Pointer
	Priority float32
}

func NewMemoryPriorityAllocateInfoEXT() *MemoryPriorityAllocateInfoEXT {
	p := (*MemoryPriorityAllocateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*MemoryPriorityAllocateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT
	return p
}
func (p *MemoryPriorityAllocateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const NV_dedicated_allocation_image_aliasing = 1
const NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_SPEC_VERSION = 1

var NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_EXTENSION_NAME = "VK_NV_dedicated_allocation_image_aliasing"

// PhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV.html
type PhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	DedicatedAllocationImageAliasing Bool32
}

func NewPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV() *PhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV {
	p := (*PhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_DEDICATED_ALLOCATION_IMAGE_ALIASING_FEATURES_NV
	return p
}
func (p *PhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_buffer_device_address = 1

type DeviceAddress = uint64

const EXT_BUFFER_DEVICE_ADDRESS_SPEC_VERSION = 2

var EXT_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME = "VK_EXT_buffer_device_address"

// PhysicalDeviceBufferDeviceAddressFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeaturesEXT.html
type PhysicalDeviceBufferDeviceAddressFeaturesEXT struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	BufferDeviceAddress              Bool32
	BufferDeviceAddressCaptureReplay Bool32
	BufferDeviceAddressMultiDevice   Bool32
}

func NewPhysicalDeviceBufferDeviceAddressFeaturesEXT() *PhysicalDeviceBufferDeviceAddressFeaturesEXT {
	p := (*PhysicalDeviceBufferDeviceAddressFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceBufferDeviceAddressFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceBufferDeviceAddressFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

type PhysicalDeviceBufferAddressFeaturesEXT = PhysicalDeviceBufferDeviceAddressFeaturesEXT

// BufferDeviceAddressInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferDeviceAddressInfoEXT.html
type BufferDeviceAddressInfoEXT struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Buffer Buffer
}

func NewBufferDeviceAddressInfoEXT() *BufferDeviceAddressInfoEXT {
	p := (*BufferDeviceAddressInfoEXT)(MemAlloc(unsafe.Sizeof(*(*BufferDeviceAddressInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_EXT
	return p
}
func (p *BufferDeviceAddressInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// BufferDeviceAddressCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkBufferDeviceAddressCreateInfoEXT.html
type BufferDeviceAddressCreateInfoEXT struct {
	SType         StructureType
	PNext         unsafe.Pointer
	DeviceAddress DeviceAddress
}

func NewBufferDeviceAddressCreateInfoEXT() *BufferDeviceAddressCreateInfoEXT {
	p := (*BufferDeviceAddressCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*BufferDeviceAddressCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT
	return p
}
func (p *BufferDeviceAddressCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetBufferDeviceAddressEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetBufferDeviceAddressEXT.html
type PfnGetBufferDeviceAddressEXT uintptr

func (fn PfnGetBufferDeviceAddressEXT) Call(device Device, pInfo *BufferDeviceAddressInfoEXT) DeviceAddress {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pInfo)))
	return DeviceAddress(ret)
}
func (fn PfnGetBufferDeviceAddressEXT) String() string { return "vkGetBufferDeviceAddressEXT" }

const EXT_separate_stencil_usage = 1
const EXT_SEPARATE_STENCIL_USAGE_SPEC_VERSION = 1

var EXT_SEPARATE_STENCIL_USAGE_EXTENSION_NAME = "VK_EXT_separate_stencil_usage"

// ImageStencilUsageCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImageStencilUsageCreateInfoEXT.html
type ImageStencilUsageCreateInfoEXT struct {
	SType        StructureType
	PNext        unsafe.Pointer
	StencilUsage ImageUsageFlags
}

func NewImageStencilUsageCreateInfoEXT() *ImageStencilUsageCreateInfoEXT {
	p := (*ImageStencilUsageCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*ImageStencilUsageCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT
	return p
}
func (p *ImageStencilUsageCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_validation_features = 1
const EXT_VALIDATION_FEATURES_SPEC_VERSION = 2

var EXT_VALIDATION_FEATURES_EXTENSION_NAME = "VK_EXT_validation_features"

// ValidationFeatureEnableEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationFeatureEnableEXT.html
type ValidationFeatureEnableEXT int32

const (
	VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT                      ValidationFeatureEnableEXT = 0
	VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT ValidationFeatureEnableEXT = 1
	VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT                    ValidationFeatureEnableEXT = 2
	VALIDATION_FEATURE_ENABLE_BEGIN_RANGE_EXT                       ValidationFeatureEnableEXT = VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT
	VALIDATION_FEATURE_ENABLE_END_RANGE_EXT                         ValidationFeatureEnableEXT = VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT
	VALIDATION_FEATURE_ENABLE_RANGE_SIZE_EXT                        ValidationFeatureEnableEXT = (VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT - VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT + 1)
	VALIDATION_FEATURE_ENABLE_MAX_ENUM_EXT                          ValidationFeatureEnableEXT = 0x7FFFFFFF
)

func (x ValidationFeatureEnableEXT) String() string {
	switch x {
	case VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT:
		return "VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT"
	case VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT:
		return "VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT"
	case VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT:
		return "VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT"
	case VALIDATION_FEATURE_ENABLE_MAX_ENUM_EXT:
		return "VALIDATION_FEATURE_ENABLE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ValidationFeatureDisableEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationFeatureDisableEXT.html
type ValidationFeatureDisableEXT int32

const (
	VALIDATION_FEATURE_DISABLE_ALL_EXT              ValidationFeatureDisableEXT = 0
	VALIDATION_FEATURE_DISABLE_SHADERS_EXT          ValidationFeatureDisableEXT = 1
	VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT    ValidationFeatureDisableEXT = 2
	VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT   ValidationFeatureDisableEXT = 3
	VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT ValidationFeatureDisableEXT = 4
	VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT      ValidationFeatureDisableEXT = 5
	VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT   ValidationFeatureDisableEXT = 6
	VALIDATION_FEATURE_DISABLE_BEGIN_RANGE_EXT      ValidationFeatureDisableEXT = VALIDATION_FEATURE_DISABLE_ALL_EXT
	VALIDATION_FEATURE_DISABLE_END_RANGE_EXT        ValidationFeatureDisableEXT = VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT
	VALIDATION_FEATURE_DISABLE_RANGE_SIZE_EXT       ValidationFeatureDisableEXT = (VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT - VALIDATION_FEATURE_DISABLE_ALL_EXT + 1)
	VALIDATION_FEATURE_DISABLE_MAX_ENUM_EXT         ValidationFeatureDisableEXT = 0x7FFFFFFF
)

func (x ValidationFeatureDisableEXT) String() string {
	switch x {
	case VALIDATION_FEATURE_DISABLE_ALL_EXT:
		return "VALIDATION_FEATURE_DISABLE_ALL_EXT"
	case VALIDATION_FEATURE_DISABLE_SHADERS_EXT:
		return "VALIDATION_FEATURE_DISABLE_SHADERS_EXT"
	case VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT:
		return "VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT"
	case VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT:
		return "VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT"
	case VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT:
		return "VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT"
	case VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT:
		return "VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT"
	case VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT:
		return "VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT"
	case VALIDATION_FEATURE_DISABLE_MAX_ENUM_EXT:
		return "VALIDATION_FEATURE_DISABLE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ValidationFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkValidationFeaturesEXT.html
type ValidationFeaturesEXT struct {
	SType                          StructureType
	PNext                          unsafe.Pointer
	EnabledValidationFeatureCount  uint32
	PEnabledValidationFeatures     *ValidationFeatureEnableEXT
	DisabledValidationFeatureCount uint32
	PDisabledValidationFeatures    *ValidationFeatureDisableEXT
}

func NewValidationFeaturesEXT() *ValidationFeaturesEXT {
	p := (*ValidationFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*ValidationFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_VALIDATION_FEATURES_EXT
	return p
}
func (p *ValidationFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const NV_cooperative_matrix = 1
const NV_COOPERATIVE_MATRIX_SPEC_VERSION = 1

var NV_COOPERATIVE_MATRIX_EXTENSION_NAME = "VK_NV_cooperative_matrix"

// ComponentTypeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkComponentTypeNV.html
type ComponentTypeNV int32

const (
	COMPONENT_TYPE_FLOAT16_NV     ComponentTypeNV = 0
	COMPONENT_TYPE_FLOAT32_NV     ComponentTypeNV = 1
	COMPONENT_TYPE_FLOAT64_NV     ComponentTypeNV = 2
	COMPONENT_TYPE_SINT8_NV       ComponentTypeNV = 3
	COMPONENT_TYPE_SINT16_NV      ComponentTypeNV = 4
	COMPONENT_TYPE_SINT32_NV      ComponentTypeNV = 5
	COMPONENT_TYPE_SINT64_NV      ComponentTypeNV = 6
	COMPONENT_TYPE_UINT8_NV       ComponentTypeNV = 7
	COMPONENT_TYPE_UINT16_NV      ComponentTypeNV = 8
	COMPONENT_TYPE_UINT32_NV      ComponentTypeNV = 9
	COMPONENT_TYPE_UINT64_NV      ComponentTypeNV = 10
	COMPONENT_TYPE_BEGIN_RANGE_NV ComponentTypeNV = COMPONENT_TYPE_FLOAT16_NV
	COMPONENT_TYPE_END_RANGE_NV   ComponentTypeNV = COMPONENT_TYPE_UINT64_NV
	COMPONENT_TYPE_RANGE_SIZE_NV  ComponentTypeNV = (COMPONENT_TYPE_UINT64_NV - COMPONENT_TYPE_FLOAT16_NV + 1)
	COMPONENT_TYPE_MAX_ENUM_NV    ComponentTypeNV = 0x7FFFFFFF
)

func (x ComponentTypeNV) String() string {
	switch x {
	case COMPONENT_TYPE_FLOAT16_NV:
		return "COMPONENT_TYPE_FLOAT16_NV"
	case COMPONENT_TYPE_FLOAT32_NV:
		return "COMPONENT_TYPE_FLOAT32_NV"
	case COMPONENT_TYPE_FLOAT64_NV:
		return "COMPONENT_TYPE_FLOAT64_NV"
	case COMPONENT_TYPE_SINT8_NV:
		return "COMPONENT_TYPE_SINT8_NV"
	case COMPONENT_TYPE_SINT16_NV:
		return "COMPONENT_TYPE_SINT16_NV"
	case COMPONENT_TYPE_SINT32_NV:
		return "COMPONENT_TYPE_SINT32_NV"
	case COMPONENT_TYPE_SINT64_NV:
		return "COMPONENT_TYPE_SINT64_NV"
	case COMPONENT_TYPE_UINT8_NV:
		return "COMPONENT_TYPE_UINT8_NV"
	case COMPONENT_TYPE_UINT16_NV:
		return "COMPONENT_TYPE_UINT16_NV"
	case COMPONENT_TYPE_UINT32_NV:
		return "COMPONENT_TYPE_UINT32_NV"
	case COMPONENT_TYPE_UINT64_NV:
		return "COMPONENT_TYPE_UINT64_NV"
	case COMPONENT_TYPE_MAX_ENUM_NV:
		return "COMPONENT_TYPE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// ScopeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkScopeNV.html
type ScopeNV int32

const (
	SCOPE_DEVICE_NV       ScopeNV = 1
	SCOPE_WORKGROUP_NV    ScopeNV = 2
	SCOPE_SUBGROUP_NV     ScopeNV = 3
	SCOPE_QUEUE_FAMILY_NV ScopeNV = 5
	SCOPE_BEGIN_RANGE_NV  ScopeNV = SCOPE_DEVICE_NV
	SCOPE_END_RANGE_NV    ScopeNV = SCOPE_QUEUE_FAMILY_NV
	SCOPE_RANGE_SIZE_NV   ScopeNV = (SCOPE_QUEUE_FAMILY_NV - SCOPE_DEVICE_NV + 1)
	SCOPE_MAX_ENUM_NV     ScopeNV = 0x7FFFFFFF
)

func (x ScopeNV) String() string {
	switch x {
	case SCOPE_DEVICE_NV:
		return "SCOPE_DEVICE_NV"
	case SCOPE_WORKGROUP_NV:
		return "SCOPE_WORKGROUP_NV"
	case SCOPE_SUBGROUP_NV:
		return "SCOPE_SUBGROUP_NV"
	case SCOPE_QUEUE_FAMILY_NV:
		return "SCOPE_QUEUE_FAMILY_NV"
	case SCOPE_MAX_ENUM_NV:
		return "SCOPE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

// CooperativeMatrixPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCooperativeMatrixPropertiesNV.html
type CooperativeMatrixPropertiesNV struct {
	SType StructureType
	PNext unsafe.Pointer
	MSize uint32
	NSize uint32
	KSize uint32
	AType ComponentTypeNV
	BType ComponentTypeNV
	CType ComponentTypeNV
	DType ComponentTypeNV
	Scope ScopeNV
}

func NewCooperativeMatrixPropertiesNV() *CooperativeMatrixPropertiesNV {
	p := (*CooperativeMatrixPropertiesNV)(MemAlloc(unsafe.Sizeof(*(*CooperativeMatrixPropertiesNV)(nil))))
	p.SType = STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV
	return p
}
func (p *CooperativeMatrixPropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceCooperativeMatrixFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceCooperativeMatrixFeaturesNV.html
type PhysicalDeviceCooperativeMatrixFeaturesNV struct {
	SType                               StructureType
	PNext                               unsafe.Pointer
	CooperativeMatrix                   Bool32
	CooperativeMatrixRobustBufferAccess Bool32
}

func NewPhysicalDeviceCooperativeMatrixFeaturesNV() *PhysicalDeviceCooperativeMatrixFeaturesNV {
	p := (*PhysicalDeviceCooperativeMatrixFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceCooperativeMatrixFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV
	return p
}
func (p *PhysicalDeviceCooperativeMatrixFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceCooperativeMatrixPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesNV.html
type PhysicalDeviceCooperativeMatrixPropertiesNV struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	CooperativeMatrixSupportedStages ShaderStageFlags
}

func NewPhysicalDeviceCooperativeMatrixPropertiesNV() *PhysicalDeviceCooperativeMatrixPropertiesNV {
	p := (*PhysicalDeviceCooperativeMatrixPropertiesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceCooperativeMatrixPropertiesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_NV
	return p
}
func (p *PhysicalDeviceCooperativeMatrixPropertiesNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceCooperativeMatrixPropertiesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html
type PfnGetPhysicalDeviceCooperativeMatrixPropertiesNV uintptr

func (fn PfnGetPhysicalDeviceCooperativeMatrixPropertiesNV) Call(physicalDevice PhysicalDevice, pPropertyCount *uint32, pProperties *CooperativeMatrixPropertiesNV) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pPropertyCount)), uintptr(unsafe.Pointer(pProperties)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceCooperativeMatrixPropertiesNV) String() string {
	return "vkGetPhysicalDeviceCooperativeMatrixPropertiesNV"
}

const NV_coverage_reduction_mode = 1
const NV_COVERAGE_REDUCTION_MODE_SPEC_VERSION = 1

var NV_COVERAGE_REDUCTION_MODE_EXTENSION_NAME = "VK_NV_coverage_reduction_mode"

// CoverageReductionModeNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkCoverageReductionModeNV.html
type CoverageReductionModeNV int32

const (
	COVERAGE_REDUCTION_MODE_MERGE_NV       CoverageReductionModeNV = 0
	COVERAGE_REDUCTION_MODE_TRUNCATE_NV    CoverageReductionModeNV = 1
	COVERAGE_REDUCTION_MODE_BEGIN_RANGE_NV CoverageReductionModeNV = COVERAGE_REDUCTION_MODE_MERGE_NV
	COVERAGE_REDUCTION_MODE_END_RANGE_NV   CoverageReductionModeNV = COVERAGE_REDUCTION_MODE_TRUNCATE_NV
	COVERAGE_REDUCTION_MODE_RANGE_SIZE_NV  CoverageReductionModeNV = (COVERAGE_REDUCTION_MODE_TRUNCATE_NV - COVERAGE_REDUCTION_MODE_MERGE_NV + 1)
	COVERAGE_REDUCTION_MODE_MAX_ENUM_NV    CoverageReductionModeNV = 0x7FFFFFFF
)

func (x CoverageReductionModeNV) String() string {
	switch x {
	case COVERAGE_REDUCTION_MODE_MERGE_NV:
		return "COVERAGE_REDUCTION_MODE_MERGE_NV"
	case COVERAGE_REDUCTION_MODE_TRUNCATE_NV:
		return "COVERAGE_REDUCTION_MODE_TRUNCATE_NV"
	case COVERAGE_REDUCTION_MODE_MAX_ENUM_NV:
		return "COVERAGE_REDUCTION_MODE_MAX_ENUM_NV"
	default:
		return fmt.Sprint(int32(x))
	}
}

type PipelineCoverageReductionStateCreateFlagsNV uint32 // reserved
// PhysicalDeviceCoverageReductionModeFeaturesNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceCoverageReductionModeFeaturesNV.html
type PhysicalDeviceCoverageReductionModeFeaturesNV struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	CoverageReductionMode Bool32
}

func NewPhysicalDeviceCoverageReductionModeFeaturesNV() *PhysicalDeviceCoverageReductionModeFeaturesNV {
	p := (*PhysicalDeviceCoverageReductionModeFeaturesNV)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceCoverageReductionModeFeaturesNV)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_COVERAGE_REDUCTION_MODE_FEATURES_NV
	return p
}
func (p *PhysicalDeviceCoverageReductionModeFeaturesNV) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineCoverageReductionStateCreateInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html
type PipelineCoverageReductionStateCreateInfoNV struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	Flags                 PipelineCoverageReductionStateCreateFlagsNV
	CoverageReductionMode CoverageReductionModeNV
}

func NewPipelineCoverageReductionStateCreateInfoNV() *PipelineCoverageReductionStateCreateInfoNV {
	p := (*PipelineCoverageReductionStateCreateInfoNV)(MemAlloc(unsafe.Sizeof(*(*PipelineCoverageReductionStateCreateInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV
	return p
}
func (p *PipelineCoverageReductionStateCreateInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// FramebufferMixedSamplesCombinationNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFramebufferMixedSamplesCombinationNV.html
type FramebufferMixedSamplesCombinationNV struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	CoverageReductionMode CoverageReductionModeNV
	RasterizationSamples  SampleCountFlags
	DepthStencilSamples   SampleCountFlags
	ColorSamples          SampleCountFlags
}

func NewFramebufferMixedSamplesCombinationNV() *FramebufferMixedSamplesCombinationNV {
	p := (*FramebufferMixedSamplesCombinationNV)(MemAlloc(unsafe.Sizeof(*(*FramebufferMixedSamplesCombinationNV)(nil))))
	p.SType = STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV
	return p
}
func (p *FramebufferMixedSamplesCombinationNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html
type PfnGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV uintptr

func (fn PfnGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV) Call(physicalDevice PhysicalDevice, pCombinationCount *uint32, pCombinations *FramebufferMixedSamplesCombinationNV) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pCombinationCount)), uintptr(unsafe.Pointer(pCombinations)))
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV) String() string {
	return "vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV"
}

const EXT_fragment_shader_interlock = 1
const EXT_FRAGMENT_SHADER_INTERLOCK_SPEC_VERSION = 1

var EXT_FRAGMENT_SHADER_INTERLOCK_EXTENSION_NAME = "VK_EXT_fragment_shader_interlock"

// PhysicalDeviceFragmentShaderInterlockFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT.html
type PhysicalDeviceFragmentShaderInterlockFeaturesEXT struct {
	SType                              StructureType
	PNext                              unsafe.Pointer
	FragmentShaderSampleInterlock      Bool32
	FragmentShaderPixelInterlock       Bool32
	FragmentShaderShadingRateInterlock Bool32
}

func NewPhysicalDeviceFragmentShaderInterlockFeaturesEXT() *PhysicalDeviceFragmentShaderInterlockFeaturesEXT {
	p := (*PhysicalDeviceFragmentShaderInterlockFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceFragmentShaderInterlockFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_INTERLOCK_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceFragmentShaderInterlockFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_ycbcr_image_arrays = 1
const EXT_YCBCR_IMAGE_ARRAYS_SPEC_VERSION = 1

var EXT_YCBCR_IMAGE_ARRAYS_EXTENSION_NAME = "VK_EXT_ycbcr_image_arrays"

// PhysicalDeviceYcbcrImageArraysFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceYcbcrImageArraysFeaturesEXT.html
type PhysicalDeviceYcbcrImageArraysFeaturesEXT struct {
	SType            StructureType
	PNext            unsafe.Pointer
	YcbcrImageArrays Bool32
}

func NewPhysicalDeviceYcbcrImageArraysFeaturesEXT() *PhysicalDeviceYcbcrImageArraysFeaturesEXT {
	p := (*PhysicalDeviceYcbcrImageArraysFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceYcbcrImageArraysFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_IMAGE_ARRAYS_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceYcbcrImageArraysFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_headless_surface = 1
const EXT_HEADLESS_SURFACE_SPEC_VERSION = 1

var EXT_HEADLESS_SURFACE_EXTENSION_NAME = "VK_EXT_headless_surface"

type HeadlessSurfaceCreateFlagsEXT uint32 // reserved
// HeadlessSurfaceCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkHeadlessSurfaceCreateInfoEXT.html
type HeadlessSurfaceCreateInfoEXT struct {
	SType StructureType
	PNext unsafe.Pointer
	Flags HeadlessSurfaceCreateFlagsEXT
}

func NewHeadlessSurfaceCreateInfoEXT() *HeadlessSurfaceCreateInfoEXT {
	p := (*HeadlessSurfaceCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*HeadlessSurfaceCreateInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT
	return p
}
func (p *HeadlessSurfaceCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateHeadlessSurfaceEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateHeadlessSurfaceEXT.html
type PfnCreateHeadlessSurfaceEXT uintptr

func (fn PfnCreateHeadlessSurfaceEXT) Call(instance Instance, pCreateInfo *HeadlessSurfaceCreateInfoEXT, pAllocator *AllocationCallbacks, pSurface *SurfaceKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)))
	return Result(ret)
}
func (fn PfnCreateHeadlessSurfaceEXT) String() string { return "vkCreateHeadlessSurfaceEXT" }

const EXT_line_rasterization = 1
const EXT_LINE_RASTERIZATION_SPEC_VERSION = 1

var EXT_LINE_RASTERIZATION_EXTENSION_NAME = "VK_EXT_line_rasterization"

// LineRasterizationModeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkLineRasterizationModeEXT.html
type LineRasterizationModeEXT int32

const (
	LINE_RASTERIZATION_MODE_DEFAULT_EXT            LineRasterizationModeEXT = 0
	LINE_RASTERIZATION_MODE_RECTANGULAR_EXT        LineRasterizationModeEXT = 1
	LINE_RASTERIZATION_MODE_BRESENHAM_EXT          LineRasterizationModeEXT = 2
	LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT LineRasterizationModeEXT = 3
	LINE_RASTERIZATION_MODE_BEGIN_RANGE_EXT        LineRasterizationModeEXT = LINE_RASTERIZATION_MODE_DEFAULT_EXT
	LINE_RASTERIZATION_MODE_END_RANGE_EXT          LineRasterizationModeEXT = LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT
	LINE_RASTERIZATION_MODE_RANGE_SIZE_EXT         LineRasterizationModeEXT = (LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT - LINE_RASTERIZATION_MODE_DEFAULT_EXT + 1)
	LINE_RASTERIZATION_MODE_MAX_ENUM_EXT           LineRasterizationModeEXT = 0x7FFFFFFF
)

func (x LineRasterizationModeEXT) String() string {
	switch x {
	case LINE_RASTERIZATION_MODE_DEFAULT_EXT:
		return "LINE_RASTERIZATION_MODE_DEFAULT_EXT"
	case LINE_RASTERIZATION_MODE_RECTANGULAR_EXT:
		return "LINE_RASTERIZATION_MODE_RECTANGULAR_EXT"
	case LINE_RASTERIZATION_MODE_BRESENHAM_EXT:
		return "LINE_RASTERIZATION_MODE_BRESENHAM_EXT"
	case LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT:
		return "LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT"
	case LINE_RASTERIZATION_MODE_MAX_ENUM_EXT:
		return "LINE_RASTERIZATION_MODE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// PhysicalDeviceLineRasterizationFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceLineRasterizationFeaturesEXT.html
type PhysicalDeviceLineRasterizationFeaturesEXT struct {
	SType                    StructureType
	PNext                    unsafe.Pointer
	RectangularLines         Bool32
	BresenhamLines           Bool32
	SmoothLines              Bool32
	StippledRectangularLines Bool32
	StippledBresenhamLines   Bool32
	StippledSmoothLines      Bool32
}

func NewPhysicalDeviceLineRasterizationFeaturesEXT() *PhysicalDeviceLineRasterizationFeaturesEXT {
	return (*PhysicalDeviceLineRasterizationFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceLineRasterizationFeaturesEXT)(nil))))
}
func (p *PhysicalDeviceLineRasterizationFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceLineRasterizationPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceLineRasterizationPropertiesEXT.html
type PhysicalDeviceLineRasterizationPropertiesEXT struct {
	SType                     StructureType
	PNext                     unsafe.Pointer
	LineSubPixelPrecisionBits uint32
}

func NewPhysicalDeviceLineRasterizationPropertiesEXT() *PhysicalDeviceLineRasterizationPropertiesEXT {
	return (*PhysicalDeviceLineRasterizationPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceLineRasterizationPropertiesEXT)(nil))))
}
func (p *PhysicalDeviceLineRasterizationPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PipelineRasterizationLineStateCreateInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoEXT.html
type PipelineRasterizationLineStateCreateInfoEXT struct {
	SType                 StructureType
	PNext                 unsafe.Pointer
	LineRasterizationMode LineRasterizationModeEXT
	StippledLineEnable    Bool32
	LineStippleFactor     uint32
	LineStipplePattern    uint16
}

func NewPipelineRasterizationLineStateCreateInfoEXT() *PipelineRasterizationLineStateCreateInfoEXT {
	return (*PipelineRasterizationLineStateCreateInfoEXT)(MemAlloc(unsafe.Sizeof(*(*PipelineRasterizationLineStateCreateInfoEXT)(nil))))
}
func (p *PipelineRasterizationLineStateCreateInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCmdSetLineStippleEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCmdSetLineStippleEXT.html
type PfnCmdSetLineStippleEXT uintptr

func (fn PfnCmdSetLineStippleEXT) Call(commandBuffer CommandBuffer, lineStippleFactor uint32, lineStipplePattern uint16) {
	_, _, _ = call(uintptr(fn), uintptr(commandBuffer), uintptr(lineStippleFactor), uintptr(lineStipplePattern))
}
func (fn PfnCmdSetLineStippleEXT) String() string { return "vkCmdSetLineStippleEXT" }

const EXT_host_query_reset = 1
const EXT_HOST_QUERY_RESET_SPEC_VERSION = 1

var EXT_HOST_QUERY_RESET_EXTENSION_NAME = "VK_EXT_host_query_reset"

// PhysicalDeviceHostQueryResetFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceHostQueryResetFeaturesEXT.html
type PhysicalDeviceHostQueryResetFeaturesEXT struct {
	SType          StructureType
	PNext          unsafe.Pointer
	HostQueryReset Bool32
}

func NewPhysicalDeviceHostQueryResetFeaturesEXT() *PhysicalDeviceHostQueryResetFeaturesEXT {
	p := (*PhysicalDeviceHostQueryResetFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceHostQueryResetFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceHostQueryResetFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnResetQueryPoolEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkResetQueryPoolEXT.html
type PfnResetQueryPoolEXT uintptr

func (fn PfnResetQueryPoolEXT) Call(device Device, queryPool QueryPool, firstQuery, queryCount uint32) {
	_, _, _ = call(uintptr(fn), uintptr(device), uintptr(queryPool), uintptr(firstQuery), uintptr(queryCount))
}
func (fn PfnResetQueryPoolEXT) String() string { return "vkResetQueryPoolEXT" }

const EXT_index_type_uint8 = 1
const EXT_INDEX_TYPE_UINT8_SPEC_VERSION = 1

var EXT_INDEX_TYPE_UINT8_EXTENSION_NAME = "VK_EXT_index_type_uint8"

// PhysicalDeviceIndexTypeUint8FeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceIndexTypeUint8FeaturesEXT.html
type PhysicalDeviceIndexTypeUint8FeaturesEXT struct {
	SType          StructureType
	PNext          unsafe.Pointer
	IndexTypeUint8 Bool32
}

func NewPhysicalDeviceIndexTypeUint8FeaturesEXT() *PhysicalDeviceIndexTypeUint8FeaturesEXT {
	return (*PhysicalDeviceIndexTypeUint8FeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceIndexTypeUint8FeaturesEXT)(nil))))
}
func (p *PhysicalDeviceIndexTypeUint8FeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_shader_demote_to_helper_invocation = 1
const EXT_SHADER_DEMOTE_TO_HELPER_INVOCATION_SPEC_VERSION = 1

var EXT_SHADER_DEMOTE_TO_HELPER_INVOCATION_EXTENSION_NAME = "VK_EXT_shader_demote_to_helper_invocation"

// PhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT.html
type PhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT struct {
	SType                          StructureType
	PNext                          unsafe.Pointer
	ShaderDemoteToHelperInvocation Bool32
}

func NewPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT() *PhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT {
	p := (*PhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_texel_buffer_alignment = 1
const EXT_TEXEL_BUFFER_ALIGNMENT_SPEC_VERSION = 1

var EXT_TEXEL_BUFFER_ALIGNMENT_EXTENSION_NAME = "VK_EXT_texel_buffer_alignment"

// PhysicalDeviceTexelBufferAlignmentFeaturesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT.html
type PhysicalDeviceTexelBufferAlignmentFeaturesEXT struct {
	SType                StructureType
	PNext                unsafe.Pointer
	TexelBufferAlignment Bool32
}

func NewPhysicalDeviceTexelBufferAlignmentFeaturesEXT() *PhysicalDeviceTexelBufferAlignmentFeaturesEXT {
	p := (*PhysicalDeviceTexelBufferAlignmentFeaturesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceTexelBufferAlignmentFeaturesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_FEATURES_EXT
	return p
}
func (p *PhysicalDeviceTexelBufferAlignmentFeaturesEXT) Free() { MemFree(unsafe.Pointer(p)) }

// PhysicalDeviceTexelBufferAlignmentPropertiesEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentPropertiesEXT.html
type PhysicalDeviceTexelBufferAlignmentPropertiesEXT struct {
	SType                                        StructureType
	PNext                                        unsafe.Pointer
	StorageTexelBufferOffsetAlignmentBytes       DeviceSize
	StorageTexelBufferOffsetSingleTexelAlignment Bool32
	UniformTexelBufferOffsetAlignmentBytes       DeviceSize
	UniformTexelBufferOffsetSingleTexelAlignment Bool32
}

func NewPhysicalDeviceTexelBufferAlignmentPropertiesEXT() *PhysicalDeviceTexelBufferAlignmentPropertiesEXT {
	p := (*PhysicalDeviceTexelBufferAlignmentPropertiesEXT)(MemAlloc(unsafe.Sizeof(*(*PhysicalDeviceTexelBufferAlignmentPropertiesEXT)(nil))))
	p.SType = STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES_EXT
	return p
}
func (p *PhysicalDeviceTexelBufferAlignmentPropertiesEXT) Free() { MemFree(unsafe.Pointer(p)) }
