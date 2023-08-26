package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TaskTargetServer struct {
	Id *string `json:"id,omitempty"`

	VmId *string `json:"vm_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Ip *string `json:"ip,omitempty"`

	OsType *TaskTargetServerOsType `json:"os_type,omitempty"`

	OsVersion *string `json:"os_version,omitempty"`

	SystemDir *string `json:"system_dir,omitempty"`

	Disks []TargetDisk `json:"disks"`

	VolumeGroups *[]VolumeGroups `json:"volume_groups,omitempty"`

	BtrfsList *[]string `json:"btrfs_list,omitempty"`

	ImageDiskId *string `json:"image_disk_id,omitempty"`

	CutoveredSnapshotIds *string `json:"cutovered_snapshot_ids,omitempty"`
}

func (o TaskTargetServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTargetServer struct{}"
	}

	return strings.Join([]string{"TaskTargetServer", string(data)}, " ")
}

type TaskTargetServerOsType struct {
	value string
}

type TaskTargetServerOsTypeEnum struct {
	WINDOWS TaskTargetServerOsType
	LINUX   TaskTargetServerOsType
}

func GetTaskTargetServerOsTypeEnum() TaskTargetServerOsTypeEnum {
	return TaskTargetServerOsTypeEnum{
		WINDOWS: TaskTargetServerOsType{
			value: "WINDOWS",
		},
		LINUX: TaskTargetServerOsType{
			value: "LINUX",
		},
	}
}

func (c TaskTargetServerOsType) Value() string {
	return c.value
}

func (c TaskTargetServerOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskTargetServerOsType) UnmarshalJSON(b []byte) error {
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
