package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type DiskIntargetServer struct {
	Name string `json:"name"`

	Size int64 `json:"size"`

	DeviceUse *DiskIntargetServerDeviceUse `json:"device_use,omitempty"`

	UsedSize *int64 `json:"used_size,omitempty"`

	PhysicalVolumes *[]PhysicalVolumes `json:"physical_volumes,omitempty"`
}

func (o DiskIntargetServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskIntargetServer struct{}"
	}

	return strings.Join([]string{"DiskIntargetServer", string(data)}, " ")
}

type DiskIntargetServerDeviceUse struct {
	value string
}

type DiskIntargetServerDeviceUseEnum struct {
	BOOT   DiskIntargetServerDeviceUse
	OS     DiskIntargetServerDeviceUse
	NORMAL DiskIntargetServerDeviceUse
}

func GetDiskIntargetServerDeviceUseEnum() DiskIntargetServerDeviceUseEnum {
	return DiskIntargetServerDeviceUseEnum{
		BOOT: DiskIntargetServerDeviceUse{
			value: "BOOT",
		},
		OS: DiskIntargetServerDeviceUse{
			value: "OS",
		},
		NORMAL: DiskIntargetServerDeviceUse{
			value: "NORMAL",
		},
	}
}

func (c DiskIntargetServerDeviceUse) Value() string {
	return c.value
}

func (c DiskIntargetServerDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskIntargetServerDeviceUse) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
