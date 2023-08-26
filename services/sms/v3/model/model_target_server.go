package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TargetServer struct {
	Id *string `json:"id,omitempty"`

	Ip string `json:"ip"`

	Name string `json:"name"`

	Hostname *string `json:"hostname,omitempty"`

	OsType TargetServerOsType `json:"os_type"`

	OsVersion *string `json:"os_version,omitempty"`

	Firmware *TargetServerFirmware `json:"firmware,omitempty"`

	CpuQuantity *int32 `json:"cpu_quantity,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	Disks []TargetDisk `json:"disks"`

	BtrfsList *[]string `json:"btrfs_list,omitempty"`

	Networks *[]NetWork `json:"networks,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	HasRsync *bool `json:"has_rsync,omitempty"`

	Paravirtualization *bool `json:"paravirtualization,omitempty"`

	RawDevices *string `json:"raw_devices,omitempty"`

	DriverFiles *bool `json:"driver_files,omitempty"`

	SystemServices *bool `json:"system_services,omitempty"`

	AccountRights *bool `json:"account_rights,omitempty"`

	BootLoader *TargetServerBootLoader `json:"boot_loader,omitempty"`

	SystemDir *string `json:"system_dir,omitempty"`

	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`

	VmId *string `json:"vm_id,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	ImageDiskId *string `json:"image_disk_id,omitempty"`

	SnapshotIds *string `json:"snapshot_ids,omitempty"`

	CutoveredSnapshotIds *string `json:"cutovered_snapshot_ids,omitempty"`
}

func (o TargetServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetServer struct{}"
	}

	return strings.Join([]string{"TargetServer", string(data)}, " ")
}

type TargetServerOsType struct {
	value string
}

type TargetServerOsTypeEnum struct {
	WINDOWS TargetServerOsType
	LINUX   TargetServerOsType
}

func GetTargetServerOsTypeEnum() TargetServerOsTypeEnum {
	return TargetServerOsTypeEnum{
		WINDOWS: TargetServerOsType{
			value: "WINDOWS",
		},
		LINUX: TargetServerOsType{
			value: "LINUX",
		},
	}
}

func (c TargetServerOsType) Value() string {
	return c.value
}

func (c TargetServerOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetServerOsType) UnmarshalJSON(b []byte) error {
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

type TargetServerFirmware struct {
	value string
}

type TargetServerFirmwareEnum struct {
	BIOS TargetServerFirmware
	UEFI TargetServerFirmware
}

func GetTargetServerFirmwareEnum() TargetServerFirmwareEnum {
	return TargetServerFirmwareEnum{
		BIOS: TargetServerFirmware{
			value: "BIOS",
		},
		UEFI: TargetServerFirmware{
			value: "UEFI",
		},
	}
}

func (c TargetServerFirmware) Value() string {
	return c.value
}

func (c TargetServerFirmware) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetServerFirmware) UnmarshalJSON(b []byte) error {
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

type TargetServerBootLoader struct {
	value string
}

type TargetServerBootLoaderEnum struct {
	GRUB TargetServerBootLoader
	LILO TargetServerBootLoader
}

func GetTargetServerBootLoaderEnum() TargetServerBootLoaderEnum {
	return TargetServerBootLoaderEnum{
		GRUB: TargetServerBootLoader{
			value: "GRUB",
		},
		LILO: TargetServerBootLoader{
			value: "LILO",
		},
	}
}

func (c TargetServerBootLoader) Value() string {
	return c.value
}

func (c TargetServerBootLoader) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetServerBootLoader) UnmarshalJSON(b []byte) error {
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
