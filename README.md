# vk

[![Build Status](https://travis-ci.org/toy80/vk.svg?branch=master)](https://travis-ci.org/toy80/vk)

Package vk is an experimental Vulkan binding for golang.

## Example

Install the official vulkan runtime and run the [./toy80-example-vk](./toy80-example-vk) package

```
go get github.com/toy80/vk/toy80-example-vk

toy80-example-vk

```

## Example Ouputs

```
$ toy80-example-vk
VK_KHR_device_group_creation : 0.0.1
VK_KHR_display : 0.0.23
VK_KHR_external_fence_capabilities : 0.0.1
VK_KHR_external_memory_capabilities : 0.0.1
VK_KHR_external_semaphore_capabilities : 0.0.1
VK_KHR_get_display_properties2 : 0.0.1
VK_KHR_get_physical_device_properties2 : 0.0.2
VK_KHR_get_surface_capabilities2 : 0.0.1
VK_KHR_surface : 0.0.25
VK_KHR_surface_protected_capabilities : 0.0.1
VK_KHR_win32_surface : 0.0.6
VK_EXT_debug_report : 0.0.9
VK_EXT_debug_utils : 0.0.2
VK_EXT_swapchain_colorspace : 0.0.4
VK_NV_external_memory_capabilities : 0.0.1
```

## Code snippets

```go
package main

import "github.com/toy80/vk"

//...
appInfo := vk.NewApplicationInfo()
defer appInfo.Free()
appInfo.PApplicationName = appName
appInfo.ApplicationVersion = vk.MakeVersion(1, 0, 0)
appInfo.PEngineName = engineName
appInfo.EngineVersion = vk.MakeVersion(1, 0, 0)
appInfo.ApiVersion = vk.API_VERSION_1_0

var createInfo vk.InstanceCreateInfo
createInfo.SType = vk.STRUCTURE_TYPE_INSTANCE_CREATE_INFO
createInfo.PApplicationInfo = appInfo

if ret := vk.CreateInstance(&createInfo, nil, &ex.instance); ret != vk.SUCCESS {
    log.Fatalln("vk.CreateInstance():", ret)
}
if err := ex.LoadProc(&ex.enumerateInstanceExtensionProperties); err != nil {
    log.Fatalln("Example.LoadProc():", err)
}
//...
```