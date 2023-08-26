package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TargetDisks struct {
	DeviceUse *TargetDisksDeviceUse `json:"device_use,omitempty"`

	DiskId *string `json:"disk_id,omitempty"`

	Name string `json:"name"`

	PhysicalVolumes []PhysicalVolumes `json:"physical_volumes"`

	Size int64 `json:"size"`

	UsedSize int64 `json:"used_size"`
}

func (o TargetDisks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetDisks struct{}"
	}

	return strings.Join([]string{"TargetDisks", string(data)}, " ")
}

type TargetDisksDeviceUse struct {
	value string
}

type TargetDisksDeviceUseEnum struct {
	NORMAL TargetDisksDeviceUse
	OS     TargetDisksDeviceUse
	BOOT   TargetDisksDeviceUse
}

func GetTargetDisksDeviceUseEnum() TargetDisksDeviceUseEnum {
	return TargetDisksDeviceUseEnum{
		NORMAL: TargetDisksDeviceUse{
			value: "NORMAL",
		},
		OS: TargetDisksDeviceUse{
			value: "OS",
		},
		BOOT: TargetDisksDeviceUse{
			value: "BOOT",
		},
	}
}

func (c TargetDisksDeviceUse) Value() string {
	return c.value
}

func (c TargetDisksDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetDisksDeviceUse) UnmarshalJSON(b []byte) error {
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
