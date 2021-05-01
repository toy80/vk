// +build ios

package vk

// #cgo darwin LDFLAGS: -lMoltenVK
// #include <Availability.h>
// #include <stdint.h>
// #include "./vulkan/vulkan.h"
// #include "./vulkan/vulkan_ios.h"
//
// VkResult bridge_vkCreateMacOSSurfaceMVK(uintptr_t fp,VkInstance instance,const VkMacOSSurfaceCreateInfoMVK* pCreateInfo,const VkAllocationCallbacks* pAllocator,VkSurfaceKHR* pSurface){
//   return ((PFN_vkCreateMacOSSurfaceMVK)fp)(instance,pCreateInfo,pAllocator,pSurface);
// }
import "C"

import (
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

const MVK_macos_surface = 1
const MVK_MACOS_SURFACE_SPEC_VERSION = 3

var MVK_MACOS_SURFACE_EXTENSION_NAME = "VK_MVK_macos_surface"

type MacOSSurfaceCreateFlagsMVK uint32 // reserved
// MacOSSurfaceCreateInfoMVK -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkMacOSSurfaceCreateInfoMVK.html
type MacOSSurfaceCreateInfoMVK struct {
	SType StructureType
	PNext unsafe.Pointer
	Flags MacOSSurfaceCreateFlagsMVK
	PView unsafe.Pointer
}

func NewMacOSSurfaceCreateInfoMVK() *MacOSSurfaceCreateInfoMVK {
	p := (*MacOSSurfaceCreateInfoMVK)(MemAlloc(unsafe.Sizeof(*(*MacOSSurfaceCreateInfoMVK)(nil))))
	p.SType = STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK
	return p
}
func (p *MacOSSurfaceCreateInfoMVK) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateMacOSSurfaceMVK -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateMacOSSurfaceMVK.html
type PfnCreateMacOSSurfaceMVK uintptr

func (fn PfnCreateMacOSSurfaceMVK) Call(instance Instance, pCreateInfo *MacOSSurfaceCreateInfoMVK, pAllocator *AllocationCallbacks, pSurface *SurfaceKHR) Result {
	ret := C.bridge_vkCreateMacOSSurfaceMVK(C.uintptr_t(fn), (C.VkInstance)(unsafe.Pointer(uintptr(instance))), (*C.VkMacOSSurfaceCreateInfoMVK)(unsafe.Pointer(pCreateInfo)), (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)), (*C.VkSurfaceKHR)(unsafe.Pointer(pSurface)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnCreateMacOSSurfaceMVK) String() string { return "vkCreateMacOSSurfaceMVK" }
