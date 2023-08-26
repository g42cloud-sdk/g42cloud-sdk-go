package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Server struct {
	Id *string `json:"id,omitempty"`

	Ip string `json:"ip"`

	Name string `json:"name"`

	Hostname *string `json:"hostname,omitempty"`

	OsType ServerOsType `json:"os_type"`

	OsVersion *string `json:"os_version,omitempty"`

	Firmware *ServerFirmware `json:"firmware,omitempty"`

	CpuQuantity *int32 `json:"cpu_quantity,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	Disks *[]Disk `json:"disks,omitempty"`

	BtrfsList *[]BtrfsFileSystem `json:"btrfs_list,omitempty"`

	Networks *[]NetWork `json:"networks,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	HasRsync *bool `json:"has_rsync,omitempty"`

	Paravirtualization *bool `json:"paravirtualization,omitempty"`

	RawDevices *string `json:"raw_devices,omitempty"`

	DriverFiles *bool `json:"driver_files,omitempty"`

	SystemServices *bool `json:"system_services,omitempty"`

	AccountRights *bool `json:"account_rights,omitempty"`

	BootLoader *ServerBootLoader `json:"boot_loader,omitempty"`

	SystemDir *string `json:"system_dir,omitempty"`

	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`
}

func (o Server) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Server struct{}"
	}

	return strings.Join([]string{"Server", string(data)}, " ")
}

type ServerOsType struct {
	value string
}

type ServerOsTypeEnum struct {
	WINDOWS ServerOsType
	LINUX   ServerOsType
}

func GetServerOsTypeEnum() ServerOsTypeEnum {
	return ServerOsTypeEnum{
		WINDOWS: ServerOsType{
			value: "WINDOWS",
		},
		LINUX: ServerOsType{
			value: "LINUX",
		},
	}
}

func (c ServerOsType) Value() string {
	return c.value
}

func (c ServerOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerOsType) UnmarshalJSON(b []byte) error {
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

type ServerFirmware struct {
	value string
}

type ServerFirmwareEnum struct {
	BIOS ServerFirmware
	UEFI ServerFirmware
}

func GetServerFirmwareEnum() ServerFirmwareEnum {
	return ServerFirmwareEnum{
		BIOS: ServerFirmware{
			value: "BIOS",
		},
		UEFI: ServerFirmware{
			value: "UEFI",
		},
	}
}

func (c ServerFirmware) Value() string {
	return c.value
}

func (c ServerFirmware) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerFirmware) UnmarshalJSON(b []byte) error {
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

type ServerBootLoader struct {
	value string
}

type ServerBootLoaderEnum struct {
	GRUB ServerBootLoader
	LILO ServerBootLoader
}

func GetServerBootLoaderEnum() ServerBootLoaderEnum {
	return ServerBootLoaderEnum{
		GRUB: ServerBootLoader{
			value: "GRUB",
		},
		LILO: ServerBootLoader{
			value: "LILO",
		},
	}
}

func (c ServerBootLoader) Value() string {
	return c.value
}

func (c ServerBootLoader) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerBootLoader) UnmarshalJSON(b []byte) error {
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
