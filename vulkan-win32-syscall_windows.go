// +build !forcecgo

package vk

import (
	"fmt"
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

const KHR_win32_surface = 1
const KHR_WIN32_SURFACE_SPEC_VERSION = 6

var KHR_WIN32_SURFACE_EXTENSION_NAME = "VK_KHR_win32_surface"

type Win32SurfaceCreateFlagsKHR uint32 // reserved
// Win32SurfaceCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkWin32SurfaceCreateInfoKHR.html
type Win32SurfaceCreateInfoKHR struct {
	SType     StructureType
	PNext     unsafe.Pointer
	Flags     Win32SurfaceCreateFlagsKHR
	Hinstance HINSTANCE
	Hwnd      HWND
}

func NewWin32SurfaceCreateInfoKHR() *Win32SurfaceCreateInfoKHR {
	p := (*Win32SurfaceCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*Win32SurfaceCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR
	return p
}
func (p *Win32SurfaceCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateWin32SurfaceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateWin32SurfaceKHR.html
type PfnCreateWin32SurfaceKHR uintptr

func (fn PfnCreateWin32SurfaceKHR) Call(instance Instance, pCreateInfo *Win32SurfaceCreateInfoKHR, pAllocator *AllocationCallbacks, pSurface *SurfaceKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(instance), uintptr(unsafe.Pointer(pCreateInfo)), uintptr(unsafe.Pointer(pAllocator)), uintptr(unsafe.Pointer(pSurface)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnCreateWin32SurfaceKHR) String() string { return "vkCreateWin32SurfaceKHR" }

//  PfnGetPhysicalDeviceWin32PresentationSupportKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceWin32PresentationSupportKHR.html
type PfnGetPhysicalDeviceWin32PresentationSupportKHR uintptr

func (fn PfnGetPhysicalDeviceWin32PresentationSupportKHR) Call(physicalDevice PhysicalDevice, queueFamilyIndex uint32) Bool32 {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(queueFamilyIndex))
	debugCheckAndBreak()
	return Bool32(ret)
}
func (fn PfnGetPhysicalDeviceWin32PresentationSupportKHR) String() string {
	return "vkGetPhysicalDeviceWin32PresentationSupportKHR"
}

const KHR_external_memory_win32 = 1
const KHR_EXTERNAL_MEMORY_WIN32_SPEC_VERSION = 1

var KHR_EXTERNAL_MEMORY_WIN32_EXTENSION_NAME = "VK_KHR_external_memory_win32"

// ImportMemoryWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportMemoryWin32HandleInfoKHR.html
type ImportMemoryWin32HandleInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	HandleType ExternalMemoryHandleTypeFlags
	Handle     HANDLE
	Name       LPCWSTR
}

func NewImportMemoryWin32HandleInfoKHR() *ImportMemoryWin32HandleInfoKHR {
	p := (*ImportMemoryWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImportMemoryWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV
	return p
}
func (p *ImportMemoryWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// ExportMemoryWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportMemoryWin32HandleInfoKHR.html
type ExportMemoryWin32HandleInfoKHR struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}

func NewExportMemoryWin32HandleInfoKHR() *ExportMemoryWin32HandleInfoKHR {
	p := (*ExportMemoryWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ExportMemoryWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV
	return p
}
func (p *ExportMemoryWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryWin32HandlePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryWin32HandlePropertiesKHR.html
type MemoryWin32HandlePropertiesKHR struct {
	SType          StructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

func NewMemoryWin32HandlePropertiesKHR() *MemoryWin32HandlePropertiesKHR {
	p := (*MemoryWin32HandlePropertiesKHR)(MemAlloc(unsafe.Sizeof(*(*MemoryWin32HandlePropertiesKHR)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR
	return p
}
func (p *MemoryWin32HandlePropertiesKHR) Free() { MemFree(unsafe.Pointer(p)) }

// MemoryGetWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMemoryGetWin32HandleInfoKHR.html
type MemoryGetWin32HandleInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Memory     DeviceMemory
	HandleType ExternalMemoryHandleTypeFlags
}

func NewMemoryGetWin32HandleInfoKHR() *MemoryGetWin32HandleInfoKHR {
	p := (*MemoryGetWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*MemoryGetWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR
	return p
}
func (p *MemoryGetWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetMemoryWin32HandleKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetMemoryWin32HandleKHR.html
type PfnGetMemoryWin32HandleKHR uintptr

func (fn PfnGetMemoryWin32HandleKHR) Call(device Device, pGetWin32HandleInfo *MemoryGetWin32HandleInfoKHR, pHandle *HANDLE) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(pHandle)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnGetMemoryWin32HandleKHR) String() string { return "vkGetMemoryWin32HandleKHR" }

//  PfnGetMemoryWin32HandlePropertiesKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetMemoryWin32HandlePropertiesKHR.html
type PfnGetMemoryWin32HandlePropertiesKHR uintptr

func (fn PfnGetMemoryWin32HandlePropertiesKHR) Call(device Device, handleType ExternalMemoryHandleTypeFlags, handle HANDLE, pMemoryWin32HandleProperties *MemoryWin32HandlePropertiesKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(handleType), uintptr(handle), uintptr(unsafe.Pointer(pMemoryWin32HandleProperties)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnGetMemoryWin32HandlePropertiesKHR) String() string {
	return "vkGetMemoryWin32HandlePropertiesKHR"
}

const KHR_win32_keyed_mutex = 1
const KHR_WIN32_KEYED_MUTEX_SPEC_VERSION = 1

var KHR_WIN32_KEYED_MUTEX_EXTENSION_NAME = "VK_KHR_win32_keyed_mutex"

// Win32KeyedMutexAcquireReleaseInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html
type Win32KeyedMutexAcquireReleaseInfoKHR struct {
	SType            StructureType
	PNext            unsafe.Pointer
	AcquireCount     uint32
	PAcquireSyncs    *DeviceMemory
	PAcquireKeys     *uint64
	PAcquireTimeouts *uint32
	ReleaseCount     uint32
	PReleaseSyncs    *DeviceMemory
	PReleaseKeys     *uint64
}

func NewWin32KeyedMutexAcquireReleaseInfoKHR() *Win32KeyedMutexAcquireReleaseInfoKHR {
	p := (*Win32KeyedMutexAcquireReleaseInfoKHR)(MemAlloc(unsafe.Sizeof(*(*Win32KeyedMutexAcquireReleaseInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR
	return p
}
func (p *Win32KeyedMutexAcquireReleaseInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

const KHR_external_semaphore_win32 = 1
const KHR_EXTERNAL_SEMAPHORE_WIN32_SPEC_VERSION = 1

var KHR_EXTERNAL_SEMAPHORE_WIN32_EXTENSION_NAME = "VK_KHR_external_semaphore_win32"

// ImportSemaphoreWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html
type ImportSemaphoreWin32HandleInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Semaphore  Semaphore
	Flags      SemaphoreImportFlags
	HandleType ExternalSemaphoreHandleTypeFlags
	Handle     HANDLE
	Name       LPCWSTR
}

func NewImportSemaphoreWin32HandleInfoKHR() *ImportSemaphoreWin32HandleInfoKHR {
	p := (*ImportSemaphoreWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImportSemaphoreWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR
	return p
}
func (p *ImportSemaphoreWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// ExportSemaphoreWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportSemaphoreWin32HandleInfoKHR.html
type ExportSemaphoreWin32HandleInfoKHR struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}

func NewExportSemaphoreWin32HandleInfoKHR() *ExportSemaphoreWin32HandleInfoKHR {
	p := (*ExportSemaphoreWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ExportSemaphoreWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR
	return p
}
func (p *ExportSemaphoreWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// D3D12FenceSubmitInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkD3D12FenceSubmitInfoKHR.html
type D3D12FenceSubmitInfoKHR struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	WaitSemaphoreValuesCount   uint32
	PWaitSemaphoreValues       *uint64
	SignalSemaphoreValuesCount uint32
	PSignalSemaphoreValues     *uint64
}

func NewD3D12FenceSubmitInfoKHR() *D3D12FenceSubmitInfoKHR {
	p := (*D3D12FenceSubmitInfoKHR)(MemAlloc(unsafe.Sizeof(*(*D3D12FenceSubmitInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR
	return p
}
func (p *D3D12FenceSubmitInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// SemaphoreGetWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSemaphoreGetWin32HandleInfoKHR.html
type SemaphoreGetWin32HandleInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Semaphore  Semaphore
	HandleType ExternalSemaphoreHandleTypeFlags
}

func NewSemaphoreGetWin32HandleInfoKHR() *SemaphoreGetWin32HandleInfoKHR {
	p := (*SemaphoreGetWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*SemaphoreGetWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR
	return p
}
func (p *SemaphoreGetWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnImportSemaphoreWin32HandleKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkImportSemaphoreWin32HandleKHR.html
type PfnImportSemaphoreWin32HandleKHR uintptr

func (fn PfnImportSemaphoreWin32HandleKHR) Call(device Device, pImportSemaphoreWin32HandleInfo *ImportSemaphoreWin32HandleInfoKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pImportSemaphoreWin32HandleInfo)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnImportSemaphoreWin32HandleKHR) String() string { return "vkImportSemaphoreWin32HandleKHR" }

//  PfnGetSemaphoreWin32HandleKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetSemaphoreWin32HandleKHR.html
type PfnGetSemaphoreWin32HandleKHR uintptr

func (fn PfnGetSemaphoreWin32HandleKHR) Call(device Device, pGetWin32HandleInfo *SemaphoreGetWin32HandleInfoKHR, pHandle *HANDLE) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(pHandle)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnGetSemaphoreWin32HandleKHR) String() string { return "vkGetSemaphoreWin32HandleKHR" }

const KHR_external_fence_win32 = 1
const KHR_EXTERNAL_FENCE_WIN32_SPEC_VERSION = 1

var KHR_EXTERNAL_FENCE_WIN32_EXTENSION_NAME = "VK_KHR_external_fence_win32"

// ImportFenceWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html
type ImportFenceWin32HandleInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Fence      Fence
	Flags      FenceImportFlags
	HandleType ExternalFenceHandleTypeFlags
	Handle     HANDLE
	Name       LPCWSTR
}

func NewImportFenceWin32HandleInfoKHR() *ImportFenceWin32HandleInfoKHR {
	p := (*ImportFenceWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ImportFenceWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR
	return p
}
func (p *ImportFenceWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// ExportFenceWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportFenceWin32HandleInfoKHR.html
type ExportFenceWin32HandleInfoKHR struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}

func NewExportFenceWin32HandleInfoKHR() *ExportFenceWin32HandleInfoKHR {
	p := (*ExportFenceWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*ExportFenceWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR
	return p
}
func (p *ExportFenceWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

// FenceGetWin32HandleInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFenceGetWin32HandleInfoKHR.html
type FenceGetWin32HandleInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Fence      Fence
	HandleType ExternalFenceHandleTypeFlags
}

func NewFenceGetWin32HandleInfoKHR() *FenceGetWin32HandleInfoKHR {
	p := (*FenceGetWin32HandleInfoKHR)(MemAlloc(unsafe.Sizeof(*(*FenceGetWin32HandleInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR
	return p
}
func (p *FenceGetWin32HandleInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnImportFenceWin32HandleKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkImportFenceWin32HandleKHR.html
type PfnImportFenceWin32HandleKHR uintptr

func (fn PfnImportFenceWin32HandleKHR) Call(device Device, pImportFenceWin32HandleInfo *ImportFenceWin32HandleInfoKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pImportFenceWin32HandleInfo)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnImportFenceWin32HandleKHR) String() string { return "vkImportFenceWin32HandleKHR" }

//  PfnGetFenceWin32HandleKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetFenceWin32HandleKHR.html
type PfnGetFenceWin32HandleKHR uintptr

func (fn PfnGetFenceWin32HandleKHR) Call(device Device, pGetWin32HandleInfo *FenceGetWin32HandleInfoKHR, pHandle *HANDLE) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pGetWin32HandleInfo)), uintptr(unsafe.Pointer(pHandle)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnGetFenceWin32HandleKHR) String() string { return "vkGetFenceWin32HandleKHR" }

const NV_external_memory_win32 = 1
const NV_EXTERNAL_MEMORY_WIN32_SPEC_VERSION = 1

var NV_EXTERNAL_MEMORY_WIN32_EXTENSION_NAME = "VK_NV_external_memory_win32"

// ImportMemoryWin32HandleInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkImportMemoryWin32HandleInfoNV.html
type ImportMemoryWin32HandleInfoNV struct {
	SType      StructureType
	PNext      unsafe.Pointer
	HandleType ExternalMemoryHandleTypeFlagsNV
	Handle     HANDLE
}

func NewImportMemoryWin32HandleInfoNV() *ImportMemoryWin32HandleInfoNV {
	p := (*ImportMemoryWin32HandleInfoNV)(MemAlloc(unsafe.Sizeof(*(*ImportMemoryWin32HandleInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV
	return p
}
func (p *ImportMemoryWin32HandleInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

// ExportMemoryWin32HandleInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkExportMemoryWin32HandleInfoNV.html
type ExportMemoryWin32HandleInfoNV struct {
	SType       StructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
}

func NewExportMemoryWin32HandleInfoNV() *ExportMemoryWin32HandleInfoNV {
	p := (*ExportMemoryWin32HandleInfoNV)(MemAlloc(unsafe.Sizeof(*(*ExportMemoryWin32HandleInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV
	return p
}
func (p *ExportMemoryWin32HandleInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetMemoryWin32HandleNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetMemoryWin32HandleNV.html
type PfnGetMemoryWin32HandleNV uintptr

func (fn PfnGetMemoryWin32HandleNV) Call(device Device, memory DeviceMemory, handleType ExternalMemoryHandleTypeFlagsNV, pHandle *HANDLE) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(memory), uintptr(handleType), uintptr(unsafe.Pointer(pHandle)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnGetMemoryWin32HandleNV) String() string { return "vkGetMemoryWin32HandleNV" }

const NV_win32_keyed_mutex = 1
const NV_WIN32_KEYED_MUTEX_SPEC_VERSION = 2

var NV_WIN32_KEYED_MUTEX_EXTENSION_NAME = "VK_NV_win32_keyed_mutex"

// Win32KeyedMutexAcquireReleaseInfoNV -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html
type Win32KeyedMutexAcquireReleaseInfoNV struct {
	SType                       StructureType
	PNext                       unsafe.Pointer
	AcquireCount                uint32
	PAcquireSyncs               *DeviceMemory
	PAcquireKeys                *uint64
	PAcquireTimeoutMilliseconds *uint32
	ReleaseCount                uint32
	PReleaseSyncs               *DeviceMemory
	PReleaseKeys                *uint64
}

func NewWin32KeyedMutexAcquireReleaseInfoNV() *Win32KeyedMutexAcquireReleaseInfoNV {
	p := (*Win32KeyedMutexAcquireReleaseInfoNV)(MemAlloc(unsafe.Sizeof(*(*Win32KeyedMutexAcquireReleaseInfoNV)(nil))))
	p.SType = STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV
	return p
}
func (p *Win32KeyedMutexAcquireReleaseInfoNV) Free() { MemFree(unsafe.Pointer(p)) }

const EXT_full_screen_exclusive = 1
const EXT_FULL_SCREEN_EXCLUSIVE_SPEC_VERSION = 4

var EXT_FULL_SCREEN_EXCLUSIVE_EXTENSION_NAME = "VK_EXT_full_screen_exclusive"

// FullScreenExclusiveEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkFullScreenExclusiveEXT.html
type FullScreenExclusiveEXT int32

const (
	FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT                FullScreenExclusiveEXT = 0
	FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT                FullScreenExclusiveEXT = 1
	FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT             FullScreenExclusiveEXT = 2
	FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT FullScreenExclusiveEXT = 3
	FULL_SCREEN_EXCLUSIVE_MAX_ENUM_EXT               FullScreenExclusiveEXT = 0x7FFFFFFF
)

func (x FullScreenExclusiveEXT) String() string {
	switch x {
	case FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT:
		return "FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT"
	case FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT:
		return "FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT"
	case FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT:
		return "FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT"
	case FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT:
		return "FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT"
	case FULL_SCREEN_EXCLUSIVE_MAX_ENUM_EXT:
		return "FULL_SCREEN_EXCLUSIVE_MAX_ENUM_EXT"
	default:
		return fmt.Sprint(int32(x))
	}
}

// SurfaceFullScreenExclusiveInfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html
type SurfaceFullScreenExclusiveInfoEXT struct {
	SType               StructureType
	PNext               unsafe.Pointer
	FullScreenExclusive FullScreenExclusiveEXT
}

func NewSurfaceFullScreenExclusiveInfoEXT() *SurfaceFullScreenExclusiveInfoEXT {
	p := (*SurfaceFullScreenExclusiveInfoEXT)(MemAlloc(unsafe.Sizeof(*(*SurfaceFullScreenExclusiveInfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT
	return p
}
func (p *SurfaceFullScreenExclusiveInfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

// SurfaceCapabilitiesFullScreenExclusiveEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html
type SurfaceCapabilitiesFullScreenExclusiveEXT struct {
	SType                        StructureType
	PNext                        unsafe.Pointer
	FullScreenExclusiveSupported Bool32
}

func NewSurfaceCapabilitiesFullScreenExclusiveEXT() *SurfaceCapabilitiesFullScreenExclusiveEXT {
	p := (*SurfaceCapabilitiesFullScreenExclusiveEXT)(MemAlloc(unsafe.Sizeof(*(*SurfaceCapabilitiesFullScreenExclusiveEXT)(nil))))
	p.SType = STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT
	return p
}
func (p *SurfaceCapabilitiesFullScreenExclusiveEXT) Free() { MemFree(unsafe.Pointer(p)) }

// SurfaceFullScreenExclusiveWin32InfoEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html
type SurfaceFullScreenExclusiveWin32InfoEXT struct {
	SType    StructureType
	PNext    unsafe.Pointer
	Hmonitor HMONITOR
}

func NewSurfaceFullScreenExclusiveWin32InfoEXT() *SurfaceFullScreenExclusiveWin32InfoEXT {
	p := (*SurfaceFullScreenExclusiveWin32InfoEXT)(MemAlloc(unsafe.Sizeof(*(*SurfaceFullScreenExclusiveWin32InfoEXT)(nil))))
	p.SType = STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT
	return p
}
func (p *SurfaceFullScreenExclusiveWin32InfoEXT) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnGetPhysicalDeviceSurfacePresentModes2EXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html
type PfnGetPhysicalDeviceSurfacePresentModes2EXT uintptr

func (fn PfnGetPhysicalDeviceSurfacePresentModes2EXT) Call(physicalDevice PhysicalDevice, pSurfaceInfo *PhysicalDeviceSurfaceInfo2KHR, pPresentModeCount *uint32, pPresentModes *PresentModeKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(physicalDevice), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pPresentModeCount)), uintptr(unsafe.Pointer(pPresentModes)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnGetPhysicalDeviceSurfacePresentModes2EXT) String() string {
	return "vkGetPhysicalDeviceSurfacePresentModes2EXT"
}

//  PfnAcquireFullScreenExclusiveModeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkAcquireFullScreenExclusiveModeEXT.html
type PfnAcquireFullScreenExclusiveModeEXT uintptr

func (fn PfnAcquireFullScreenExclusiveModeEXT) Call(device Device, swapchain SwapchainKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnAcquireFullScreenExclusiveModeEXT) String() string {
	return "vkAcquireFullScreenExclusiveModeEXT"
}

//  PfnReleaseFullScreenExclusiveModeEXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html
type PfnReleaseFullScreenExclusiveModeEXT uintptr

func (fn PfnReleaseFullScreenExclusiveModeEXT) Call(device Device, swapchain SwapchainKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(swapchain))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnReleaseFullScreenExclusiveModeEXT) String() string {
	return "vkReleaseFullScreenExclusiveModeEXT"
}

//  PfnGetDeviceGroupSurfacePresentModes2EXT -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetDeviceGroupSurfacePresentModes2EXT.html
type PfnGetDeviceGroupSurfacePresentModes2EXT uintptr

func (fn PfnGetDeviceGroupSurfacePresentModes2EXT) Call(device Device, pSurfaceInfo *PhysicalDeviceSurfaceInfo2KHR, pModes *DeviceGroupPresentModeFlagsKHR) Result {
	ret, _, _ := call(uintptr(fn), uintptr(device), uintptr(unsafe.Pointer(pSurfaceInfo)), uintptr(unsafe.Pointer(pModes)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnGetDeviceGroupSurfacePresentModes2EXT) String() string {
	return "vkGetDeviceGroupSurfacePresentModes2EXT"
}
