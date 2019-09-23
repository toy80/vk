package main

import (
	"fmt"
	"log"

	"github.com/toy80/vk"
)

var ()

type Example struct {
	instance vk.Instance

	enumerateInstanceExtensionProperties vk.PfnEnumerateInstanceExtensionProperties
}

func (ex *Example) LoadProc(ppfn interface{}) error {
	return vk.LoadInstanceProc(ex.instance, ppfn)
}

func (ex *Example) Init() {
	// copy strings into C memory for obey "the rules"
	appName, freeAppName := vk.CStr("Hello World!")
	defer freeAppName()
	engineName, freeEngineName := vk.CStr("Foo")
	defer freeEngineName()

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
	var numExts uint32
	ex.enumerateInstanceExtensionProperties.Call(nil, &numExts, nil)
	props := make([]vk.ExtensionProperties, numExts)
	if numExts > 0 {
		ex.enumerateInstanceExtensionProperties.Call(nil, &numExts, &props[0])
	}
	for _, prop := range props {
		fmt.Println(vk.GoStr(&prop.ExtensionName), ":", prop.SpecVersion)
	}
}

func main() {
	var ex Example
	ex.Init()
}
