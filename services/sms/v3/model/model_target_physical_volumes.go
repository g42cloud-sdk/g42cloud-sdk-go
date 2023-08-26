package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TargetPhysicalVolumes struct {
	Id *int64 `json:"id,omitempty"`

	DeviceUse *TargetPhysicalVolumesDeviceUse `json:"device_use,omitempty"`

	FileSystem *string `json:"file_system,omitempty"`

	Index *int32 `json:"index,omitempty"`

	MountPoint *string `json:"mount_point,omitempty"`

	Name *string `json:"name,omitempty"`

	Size *int64 `json:"size,omitempty"`

	UsedSize *int64 `json:"used_size,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	RelationName *string `json:"relation_name,omitempty"`

	FreeSize *int64 `json:"free_size,omitempty"`
}

func (o TargetPhysicalVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetPhysicalVolumes struct{}"
	}

	return strings.Join([]string{"TargetPhysicalVolumes", string(data)}, " ")
}

type TargetPhysicalVolumesDeviceUse struct {
	value string
}

type TargetPhysicalVolumesDeviceUseEnum struct {
	NORMAL TargetPhysicalVolumesDeviceUse
	OS     TargetPhysicalVolumesDeviceUse
	BOOT   TargetPhysicalVolumesDeviceUse
}

func GetTargetPhysicalVolumesDeviceUseEnum() TargetPhysicalVolumesDeviceUseEnum {
	return TargetPhysicalVolumesDeviceUseEnum{
		NORMAL: TargetPhysicalVolumesDeviceUse{
			value: "NORMAL",
		},
		OS: TargetPhysicalVolumesDeviceUse{
			value: "OS",
		},
		BOOT: TargetPhysicalVolumesDeviceUse{
			value: "BOOT",
		},
	}
}

func (c TargetPhysicalVolumesDeviceUse) Value() string {
	return c.value
}

func (c TargetPhysicalVolumesDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetPhysicalVolumesDeviceUse) UnmarshalJSON(b []byte) error {
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
