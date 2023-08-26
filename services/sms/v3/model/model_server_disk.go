package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ServerDisk struct {
	Name string `json:"name"`

	PartitionStyle *ServerDiskPartitionStyle `json:"partition_style,omitempty"`

	DeviceUse ServerDiskDeviceUse `json:"device_use"`

	Size int64 `json:"size"`

	UsedSize int64 `json:"used_size"`

	PhysicalVolumes []PhysicalVolume `json:"physical_volumes"`

	OsDisk *bool `json:"os_disk,omitempty"`

	RelationName *string `json:"relation_name,omitempty"`

	InodeSize *int32 `json:"inode_size,omitempty"`
}

func (o ServerDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerDisk struct{}"
	}

	return strings.Join([]string{"ServerDisk", string(data)}, " ")
}

type ServerDiskPartitionStyle struct {
	value string
}

type ServerDiskPartitionStyleEnum struct {
	MBR ServerDiskPartitionStyle
	GPT ServerDiskPartitionStyle
}

func GetServerDiskPartitionStyleEnum() ServerDiskPartitionStyleEnum {
	return ServerDiskPartitionStyleEnum{
		MBR: ServerDiskPartitionStyle{
			value: "MBR",
		},
		GPT: ServerDiskPartitionStyle{
			value: "GPT",
		},
	}
}

func (c ServerDiskPartitionStyle) Value() string {
	return c.value
}

func (c ServerDiskPartitionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerDiskPartitionStyle) UnmarshalJSON(b []byte) error {
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

type ServerDiskDeviceUse struct {
	value string
}

type ServerDiskDeviceUseEnum struct {
	BOOT ServerDiskDeviceUse
	OS   ServerDiskDeviceUse
}

func GetServerDiskDeviceUseEnum() ServerDiskDeviceUseEnum {
	return ServerDiskDeviceUseEnum{
		BOOT: ServerDiskDeviceUse{
			value: "BOOT",
		},
		OS: ServerDiskDeviceUse{
			value: "OS",
		},
	}
}

func (c ServerDiskDeviceUse) Value() string {
	return c.value
}

func (c ServerDiskDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerDiskDeviceUse) UnmarshalJSON(b []byte) error {
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
