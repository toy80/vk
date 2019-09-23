// +build xlib

package vk

// #cgo linux LDFLAGS: -lvulkan
// #include <stdint.h>
// #include <X11/Xlib.h>
// #include <vulkan/vulkan.h>
// #include <vulkan/vulkan_xlib.h>
//
// VkResult bridge_vkCreateXlibSurfaceKHR(uintptr_t fp,VkInstance instance,const VkXlibSurfaceCreateInfoKHR* pCreateInfo,const VkAllocationCallbacks* pAllocator,VkSurfaceKHR* pSurface){
//   return ((PFN_vkCreateXlibSurfaceKHR)fp)(instance,pCreateInfo,pAllocator,pSurface);
// }
// VkBool32 bridge_vkGetPhysicalDeviceXlibPresentationSupportKHR(uintptr_t fp,VkPhysicalDevice physicalDevice,uint32_t queueFamilyIndex,Display* dpy,VisualID visualID){
//   return ((PFN_vkGetPhysicalDeviceXlibPresentationSupportKHR)fp)(physicalDevice,queueFamilyIndex,dpy,visualID);
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
	Window   = C.Window
	Display  = C.Display
	VisualID = C.VisualID
)

const KHR_xlib_surface = 1
const KHR_XLIB_SURFACE_SPEC_VERSION = 6

var KHR_XLIB_SURFACE_EXTENSION_NAME = "VK_KHR_xlib_surface"

type XlibSurfaceCreateFlagsKHR uint32 // reserved
// XlibSurfaceCreateInfoKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/VkXlibSurfaceCreateInfoKHR.html
type XlibSurfaceCreateInfoKHR struct {
	SType  StructureType
	PNext  unsafe.Pointer
	Flags  XlibSurfaceCreateFlagsKHR
	Dpy    *Display
	Window Window
}

func NewXlibSurfaceCreateInfoKHR() *XlibSurfaceCreateInfoKHR {
	return (*XlibSurfaceCreateInfoKHR)(MemAlloc(unsafe.Sizeof(*(*XlibSurfaceCreateInfoKHR)(nil))))
}
func (p *XlibSurfaceCreateInfoKHR) Free() { MemFree(unsafe.Pointer(p)) }

//  PfnCreateXlibSurfaceKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkCreateXlibSurfaceKHR.html
type PfnCreateXlibSurfaceKHR uintptr

func (fn PfnCreateXlibSurfaceKHR) Call(instance Instance, pCreateInfo *XlibSurfaceCreateInfoKHR, pAllocator *AllocationCallbacks, pSurface *SurfaceKHR) Result {
	ret := C.bridge_vkCreateXlibSurfaceKHR(C.uintptr_t(fn), (C.VkInstance)(unsafe.Pointer(uintptr(instance))), (*C.VkXlibSurfaceCreateInfoKHR)(unsafe.Pointer(pCreateInfo)), (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)), (*C.VkSurfaceKHR)(unsafe.Pointer(pSurface)))
	return Result(ret)
}
func (fn PfnCreateXlibSurfaceKHR) String() string { return "vkCreateXlibSurfaceKHR" }

//  PfnGetPhysicalDeviceXlibPresentationSupportKHR -- https://www.khronos.org/registry/vulkan/specs/1.1-extensions/man/html/vkGetPhysicalDeviceXlibPresentationSupportKHR.html
type PfnGetPhysicalDeviceXlibPresentationSupportKHR uintptr

func (fn PfnGetPhysicalDeviceXlibPresentationSupportKHR) Call(physicalDevice PhysicalDevice, queueFamilyIndex uint32, dpy *Display, visualID VisualID) Bool32 {
	ret := C.bridge_vkGetPhysicalDeviceXlibPresentationSupportKHR(C.uintptr_t(fn), (C.VkPhysicalDevice)(unsafe.Pointer(uintptr(physicalDevice))), (C.uint32_t)(queueFamilyIndex), (*C.Display)(unsafe.Pointer(dpy)), (C.VisualID)(visualID))
	return Bool32(ret)
}
func (fn PfnGetPhysicalDeviceXlibPresentationSupportKHR) String() string {
	return "vkGetPhysicalDeviceXlibPresentationSupportKHR"
}
