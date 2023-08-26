package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Disk struct {
	Name string `json:"name"`

	PartitionStyle *DiskPartitionStyle `json:"partition_style,omitempty"`

	DeviceUse DiskDeviceUse `json:"device_use"`

	Size int64 `json:"size"`

	UsedSize int64 `json:"used_size"`

	PhysicalVolumes []PhysicalVolumes `json:"physical_volumes"`

	DiskId *string `json:"disk_id,omitempty"`

	OsDisk *bool `json:"os_disk,omitempty"`

	RelationName *string `json:"relation_name,omitempty"`
}

func (o Disk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Disk struct{}"
	}

	return strings.Join([]string{"Disk", string(data)}, " ")
}

type DiskPartitionStyle struct {
	value string
}

type DiskPartitionStyleEnum struct {
	MBR DiskPartitionStyle
	GPT DiskPartitionStyle
}

func GetDiskPartitionStyleEnum() DiskPartitionStyleEnum {
	return DiskPartitionStyleEnum{
		MBR: DiskPartitionStyle{
			value: "MBR",
		},
		GPT: DiskPartitionStyle{
			value: "GPT",
		},
	}
}

func (c DiskPartitionStyle) Value() string {
	return c.value
}

func (c DiskPartitionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskPartitionStyle) UnmarshalJSON(b []byte) error {
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

type DiskDeviceUse struct {
	value string
}

type DiskDeviceUseEnum struct {
	BOOT DiskDeviceUse
	OS   DiskDeviceUse
}

func GetDiskDeviceUseEnum() DiskDeviceUseEnum {
	return DiskDeviceUseEnum{
		BOOT: DiskDeviceUse{
			value: "BOOT",
		},
		OS: DiskDeviceUse{
			value: "OS",
		},
	}
}

func (c DiskDeviceUse) Value() string {
	return c.value
}

func (c DiskDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskDeviceUse) UnmarshalJSON(b []byte) error {
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
