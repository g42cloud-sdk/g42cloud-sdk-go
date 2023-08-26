package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TargetDisk struct {
	Id *int64 `json:"id,omitempty"`

	DeviceUse *TargetDiskDeviceUse `json:"device_use,omitempty"`

	DiskId *string `json:"disk_id,omitempty"`

	Name *string `json:"name,omitempty"`

	PhysicalVolumes *[]TargetPhysicalVolumes `json:"physical_volumes,omitempty"`

	Size *int64 `json:"size,omitempty"`

	UsedSize *int64 `json:"used_size,omitempty"`

	DiskIndex *string `json:"disk_index,omitempty"`

	OsDisk *bool `json:"os_disk,omitempty"`

	PartitionStyle *TargetDiskPartitionStyle `json:"partition_style,omitempty"`

	RelationName *string `json:"relation_name,omitempty"`
}

func (o TargetDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetDisk struct{}"
	}

	return strings.Join([]string{"TargetDisk", string(data)}, " ")
}

type TargetDiskDeviceUse struct {
	value string
}

type TargetDiskDeviceUseEnum struct {
	NORMAL TargetDiskDeviceUse
	OS     TargetDiskDeviceUse
	BOOT   TargetDiskDeviceUse
}

func GetTargetDiskDeviceUseEnum() TargetDiskDeviceUseEnum {
	return TargetDiskDeviceUseEnum{
		NORMAL: TargetDiskDeviceUse{
			value: "NORMAL",
		},
		OS: TargetDiskDeviceUse{
			value: "OS",
		},
		BOOT: TargetDiskDeviceUse{
			value: "BOOT",
		},
	}
}

func (c TargetDiskDeviceUse) Value() string {
	return c.value
}

func (c TargetDiskDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetDiskDeviceUse) UnmarshalJSON(b []byte) error {
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

type TargetDiskPartitionStyle struct {
	value string
}

type TargetDiskPartitionStyleEnum struct {
	MBR TargetDiskPartitionStyle
	GPT TargetDiskPartitionStyle
}

func GetTargetDiskPartitionStyleEnum() TargetDiskPartitionStyleEnum {
	return TargetDiskPartitionStyleEnum{
		MBR: TargetDiskPartitionStyle{
			value: "MBR",
		},
		GPT: TargetDiskPartitionStyle{
			value: "GPT",
		},
	}
}

func (c TargetDiskPartitionStyle) Value() string {
	return c.value
}

func (c TargetDiskPartitionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetDiskPartitionStyle) UnmarshalJSON(b []byte) error {
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
