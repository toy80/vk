// +build !xlib

package vk

// #cgo linux LDFLAGS: -lvulkan
// #include <stdint.h>
// #include <xcb/xcb.h>
// #include "./vulkan/vulkan.h"
// #include "./vulkan/vulkan_xcb.h"
//
// VkResult bridge_vkCreateXcbSurfaceKHR(uintptr_t fp,VkInstance instance,const VkXcbSurfaceCreateInfoKHR* pCreateInfo,const VkAllocationCallbacks* pAllocator,VkSurfaceKHR* pSurface){
//   return ((PFN_vkCreateXcbSurfaceKHR)fp)(instance,pCreateInfo,pAllocator,pSurface);
// }
// VkBool32 bridge_vkGetPhysicalDeviceXcbPresentationSupportKHR(uintptr_t fp,VkPhysicalDevice physicalDevice,uint32_t queueFamilyIndex,xcb_connection_t* connection,xcb_visualid_t visual_id){
//   return ((PFN_vkGetPhysicalDeviceXcbPresentationSupportKHR)fp)(physicalDevice,queueFamilyIndex,connection,visual_id);
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

type (
	XcbConnection = C.xcb_connection_t
	XcbWindow     = C.xcb_window_t
	XcbVisualID   = C.xcb_visualid_t
)

const KHR_xcb_surface = 1
const KHR_XCB_SURFACE_SPEC_VERSION = 6

var KHR_XCB_SURFACE_EXTENSION_NAME = "VK_KHR_xcb_surface"

type XcbSurfaceCreateFlagsKHR uint32 // reserved
// XcbSurfaceCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkXcbSurfaceCreateInfoKHR.html
type XcbSurfaceCreateInfoKHR struct {
	SType      StructureType
	PNext      unsafe.Pointer
	Flags      XcbSurfaceCreateFlagsKHR
	Connection *XcbConnection
	Window     XcbWindow
}

func NewXcbSurfaceCreateInfoKHR() *XcbSurfaceCreateInfoKHR {
	p := (*XcbSurfaceCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*XcbSurfaceCreateInfoKHR)(nil))))
	p.SType = STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR
	return p
}
func (p *XcbSurfaceCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateXcbSurfaceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateXcbSurfaceKHR.html
type PfnCreateXcbSurfaceKHR uintptr

func (fn PfnCreateXcbSurfaceKHR) Call(instance Instance, pCreateInfo *XcbSurfaceCreateInfoKHR, pAllocator *AllocationCallbacks, pSurface *SurfaceKHR) Result {
	ret := C.bridge_vkCreateXcbSurfaceKHR(C.uintptr_t(fn), (C.VkInstance)(unsafe.Pointer(uintptr(instance))), (*C.VkXcbSurfaceCreateInfoKHR)(unsafe.Pointer(pCreateInfo)), (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)), (*C.VkSurfaceKHR)(unsafe.Pointer(pSurface)))
	debugCheckAndBreak()
	return Result(ret)
}
func (fn PfnCreateXcbSurfaceKHR) String() string { return "vkCreateXcbSurfaceKHR" }

//  PfnGetPhysicalDeviceXcbPresentationSupportKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceXcbPresentationSupportKHR.html
type PfnGetPhysicalDeviceXcbPresentationSupportKHR uintptr

func (fn PfnGetPhysicalDeviceXcbPresentationSupportKHR) Call(physicalDevice PhysicalDevice, queueFamilyIndex uint32, connection *XcbConnection, visual_id XcbVisualID) Bool32 {
	ret := C.bridge_vkGetPhysicalDeviceXcbPresentationSupportKHR(C.uintptr_t(fn), (C.VkPhysicalDevice)(unsafe.Pointer(uintptr(physicalDevice))), (C.uint32_t)(queueFamilyIndex), (*C.xcb_connection_t)(unsafe.Pointer(connection)), (C.xcb_visualid_t)(visual_id))
	debugCheckAndBreak()
	return Bool32(ret)
}
func (fn PfnGetPhysicalDeviceXcbPresentationSupportKHR) String() string {
	return "vkGetPhysicalDeviceXcbPresentationSupportKHR"
}
