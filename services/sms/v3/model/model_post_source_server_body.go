package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PostSourceServerBody struct {
	Id *string `json:"id,omitempty"`

	Ip string `json:"ip"`

	Name string `json:"name"`

	Hostname *string `json:"hostname,omitempty"`

	OsType PostSourceServerBodyOsType `json:"os_type"`

	OsVersion *string `json:"os_version,omitempty"`

	VirtualizationType *string `json:"virtualization_type,omitempty"`

	LinuxBlockCheck *string `json:"linux_block_check,omitempty"`

	Firmware *PostSourceServerBodyFirmware `json:"firmware,omitempty"`

	CpuQuantity *int32 `json:"cpu_quantity,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	Disks *[]ServerDisk `json:"disks,omitempty"`

	BtrfsList *[]BtrfsFileSystem `json:"btrfs_list,omitempty"`

	Networks *[]NetWork `json:"networks,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	HasRsync *bool `json:"has_rsync,omitempty"`

	Paravirtualization *bool `json:"paravirtualization,omitempty"`

	RawDevices *string `json:"raw_devices,omitempty"`

	DriverFiles *bool `json:"driver_files,omitempty"`

	SystemServices *bool `json:"system_services,omitempty"`

	AccountRights *bool `json:"account_rights,omitempty"`

	BootLoader *PostSourceServerBodyBootLoader `json:"boot_loader,omitempty"`

	SystemDir *string `json:"system_dir,omitempty"`

	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`

	AgentVersion string `json:"agent_version"`

	KernelVersion *string `json:"kernel_version,omitempty"`

	MigrationCycle *PostSourceServerBodyMigrationCycle `json:"migration_cycle,omitempty"`

	State *PostSourceServerBodyState `json:"state,omitempty"`

	OemSystem *bool `json:"oem_system,omitempty"`
}

func (o PostSourceServerBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostSourceServerBody struct{}"
	}

	return strings.Join([]string{"PostSourceServerBody", string(data)}, " ")
}

type PostSourceServerBodyOsType struct {
	value string
}

type PostSourceServerBodyOsTypeEnum struct {
	WINDOWS PostSourceServerBodyOsType
	LINUX   PostSourceServerBodyOsType
}

func GetPostSourceServerBodyOsTypeEnum() PostSourceServerBodyOsTypeEnum {
	return PostSourceServerBodyOsTypeEnum{
		WINDOWS: PostSourceServerBodyOsType{
			value: "WINDOWS",
		},
		LINUX: PostSourceServerBodyOsType{
			value: "LINUX",
		},
	}
}

func (c PostSourceServerBodyOsType) Value() string {
	return c.value
}

func (c PostSourceServerBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyOsType) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyFirmware struct {
	value string
}

type PostSourceServerBodyFirmwareEnum struct {
	BIOS PostSourceServerBodyFirmware
	UEFI PostSourceServerBodyFirmware
}

func GetPostSourceServerBodyFirmwareEnum() PostSourceServerBodyFirmwareEnum {
	return PostSourceServerBodyFirmwareEnum{
		BIOS: PostSourceServerBodyFirmware{
			value: "BIOS",
		},
		UEFI: PostSourceServerBodyFirmware{
			value: "UEFI",
		},
	}
}

func (c PostSourceServerBodyFirmware) Value() string {
	return c.value
}

func (c PostSourceServerBodyFirmware) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyFirmware) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyBootLoader struct {
	value string
}

type PostSourceServerBodyBootLoaderEnum struct {
	GRUB PostSourceServerBodyBootLoader
	LILO PostSourceServerBodyBootLoader
}

func GetPostSourceServerBodyBootLoaderEnum() PostSourceServerBodyBootLoaderEnum {
	return PostSourceServerBodyBootLoaderEnum{
		GRUB: PostSourceServerBodyBootLoader{
			value: "GRUB",
		},
		LILO: PostSourceServerBodyBootLoader{
			value: "LILO",
		},
	}
}

func (c PostSourceServerBodyBootLoader) Value() string {
	return c.value
}

func (c PostSourceServerBodyBootLoader) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyBootLoader) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyMigrationCycle struct {
	value string
}

type PostSourceServerBodyMigrationCycleEnum struct {
	CUTOVERING  PostSourceServerBodyMigrationCycle
	CUTOVERED   PostSourceServerBodyMigrationCycle
	CHECKING    PostSourceServerBodyMigrationCycle
	SETTING     PostSourceServerBodyMigrationCycle
	REPLICATING PostSourceServerBodyMigrationCycle
	SYNCING     PostSourceServerBodyMigrationCycle
}

func GetPostSourceServerBodyMigrationCycleEnum() PostSourceServerBodyMigrationCycleEnum {
	return PostSourceServerBodyMigrationCycleEnum{
		CUTOVERING: PostSourceServerBodyMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: PostSourceServerBodyMigrationCycle{
			value: "cutovered",
		},
		CHECKING: PostSourceServerBodyMigrationCycle{
			value: "checking",
		},
		SETTING: PostSourceServerBodyMigrationCycle{
			value: "setting",
		},
		REPLICATING: PostSourceServerBodyMigrationCycle{
			value: "replicating",
		},
		SYNCING: PostSourceServerBodyMigrationCycle{
			value: "syncing",
		},
	}
}

func (c PostSourceServerBodyMigrationCycle) Value() string {
	return c.value
}

func (c PostSourceServerBodyMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyMigrationCycle) UnmarshalJSON(b []byte) error {
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

type PostSourceServerBodyState struct {
	value string
}

type PostSourceServerBodyStateEnum struct {
	UNAVAILABLE PostSourceServerBodyState
	WAITING     PostSourceServerBodyState
	INITIALIZE  PostSourceServerBodyState
	REPLICATE   PostSourceServerBodyState
	SYNCING     PostSourceServerBodyState
	STOPPING    PostSourceServerBodyState
	STOPPED     PostSourceServerBodyState
	DELETING    PostSourceServerBodyState
	ERROR       PostSourceServerBodyState
	CLONING     PostSourceServerBodyState
	CUTOVERING  PostSourceServerBodyState
	FINISHED    PostSourceServerBodyState
}

func GetPostSourceServerBodyStateEnum() PostSourceServerBodyStateEnum {
	return PostSourceServerBodyStateEnum{
		UNAVAILABLE: PostSourceServerBodyState{
			value: "unavailable",
		},
		WAITING: PostSourceServerBodyState{
			value: "waiting",
		},
		INITIALIZE: PostSourceServerBodyState{
			value: "initialize",
		},
		REPLICATE: PostSourceServerBodyState{
			value: "replicate",
		},
		SYNCING: PostSourceServerBodyState{
			value: "syncing",
		},
		STOPPING: PostSourceServerBodyState{
			value: "stopping",
		},
		STOPPED: PostSourceServerBodyState{
			value: "stopped",
		},
		DELETING: PostSourceServerBodyState{
			value: "deleting",
		},
		ERROR: PostSourceServerBodyState{
			value: "error",
		},
		CLONING: PostSourceServerBodyState{
			value: "cloning",
		},
		CUTOVERING: PostSourceServerBodyState{
			value: "cutovering",
		},
		FINISHED: PostSourceServerBodyState{
			value: "finished",
		},
	}
}

func (c PostSourceServerBodyState) Value() string {
	return c.value
}

func (c PostSourceServerBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSourceServerBodyState) UnmarshalJSON(b []byte) error {
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
