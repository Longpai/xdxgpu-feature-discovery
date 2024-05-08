package resource

import (
	"github.com/Longpai/xdxml-lib/pkg/xdxlib/device"
	"github.com/Longpai/xdxml-lib/pkg/xdxlib/xdxml"
)

type xdxmlLib struct {
	xdxml.Interface
	devicelib device.Interface
}

func NewXDXMLManager() Manager {
	xdxmllib := xdxml.New()
	devicelib := device.New(device.WithXdxml(xdxmllib))

	m := xdxmlLib{
		Interface: xdxmllib,
		devicelib: devicelib,
	}
	return m
}

func (l xdxmlLib) GetDevices() ([]Device, error) {
	libdevices, err := l.devicelib.GetDevices()
	if err != nil {
		return nil, err
	}

	var devices []Device

	for _, d := range libdevices {
		device := xdxmlDevice {
			Device:		d,
			devicelib: 	l.devicelib,
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (l xdxmlLib) Init() error {
	ret := l.Interface.Init()
	if ret != xdxml.SUCCESS {
		return ret
	}
	return nil
}

func(l xdxmlLib) Shutdown() error {
	ret := l.Interface.Shutdown()
	if ret != xdxml.SUCCESS {
		return ret
	}
	return nil
}

