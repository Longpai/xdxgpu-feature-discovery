package resource

import (
	"github.com/Longpai/xdxml-lib/pkg/xdxlib/device"
	"github.com/Longpai/xdxml-lib/pkg/xdxlib/xdxml"
)

type xdxmlDevice struct {
	device.Device
	devicelib device.Interface
}

var _ Device = (*xdxmlDevice)(nil)

// GetName returns the device name / model.
func (d xdxmlDevice) GetName() (string, error) {
	name, ret := d.Device.GetArchitecture()
	if ret != xdxml.SUCCESS {
		return "", ret
	}
	return name, nil
}

// GetTotalMemoryMB returns the total memory on a device in MB
func (d xdxmlDevice) GetTotalMemoryMB() (uint64, error) {
	Memory, ret := d.Device.GetMemoryInfo()
	if ret != xdxml.SUCCESS {
		return 0, ret
	}
	return Memory.Fb_total / (1024 * 1024), nil
}
