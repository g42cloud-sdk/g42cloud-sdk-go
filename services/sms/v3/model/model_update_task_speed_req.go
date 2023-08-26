package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type UpdateTaskSpeedReq struct {
	SubtaskName UpdateTaskSpeedReqSubtaskName `json:"subtask_name"`

	Progress int32 `json:"progress"`

	Replicatesize int64 `json:"replicatesize"`

	Totalsize int64 `json:"totalsize"`

	ProcessTrace string `json:"process_trace"`

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	CompressRate *float64 `json:"compress_rate,omitempty"`

	RemainTime *int64 `json:"remain_time,omitempty"`
}

func (o UpdateTaskSpeedReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSpeedReq struct{}"
	}

	return strings.Join([]string{"UpdateTaskSpeedReq", string(data)}, " ")
}

type UpdateTaskSpeedReqSubtaskName struct {
	value string
}

type UpdateTaskSpeedReqSubtaskNameEnum struct {
	CREATE_CLOUD_SERVER     UpdateTaskSpeedReqSubtaskName
	SSL_CONFIG              UpdateTaskSpeedReqSubtaskName
	ATTACH_AGENT_IMAGE      UpdateTaskSpeedReqSubtaskName
	DETTACH_AGENT_IMAGE     UpdateTaskSpeedReqSubtaskName
	FORMAT_DISK_LINUX       UpdateTaskSpeedReqSubtaskName
	FORMAT_DISK_LINUX_FILE  UpdateTaskSpeedReqSubtaskName
	FORMAT_DISK_LINUX_BLOCK UpdateTaskSpeedReqSubtaskName
	FORMAT_DISK_WINDOWS     UpdateTaskSpeedReqSubtaskName
	MIGRATE_LINUX_FILE      UpdateTaskSpeedReqSubtaskName
	MIGRATE_LINUX_BLOCK     UpdateTaskSpeedReqSubtaskName
	MIGRATE_WINDOWS_BLOCK   UpdateTaskSpeedReqSubtaskName
	CLONE_VM                UpdateTaskSpeedReqSubtaskName
	SYNC_LINUX_FILE         UpdateTaskSpeedReqSubtaskName
	SYNC_LINUX_BLOCK        UpdateTaskSpeedReqSubtaskName
	SYNC_WINDOWS_BLOCK      UpdateTaskSpeedReqSubtaskName
	CONFIGURE_LINUX         UpdateTaskSpeedReqSubtaskName
	CONFIGURE_LINUX_BLOCK   UpdateTaskSpeedReqSubtaskName
	CONFIGURE_LINUX_FILE    UpdateTaskSpeedReqSubtaskName
	CONFIGURE_WINDOWS       UpdateTaskSpeedReqSubtaskName
}

func GetUpdateTaskSpeedReqSubtaskNameEnum() UpdateTaskSpeedReqSubtaskNameEnum {
	return UpdateTaskSpeedReqSubtaskNameEnum{
		CREATE_CLOUD_SERVER: UpdateTaskSpeedReqSubtaskName{
			value: "CREATE_CLOUD_SERVER",
		},
		SSL_CONFIG: UpdateTaskSpeedReqSubtaskName{
			value: "SSL_CONFIG",
		},
		ATTACH_AGENT_IMAGE: UpdateTaskSpeedReqSubtaskName{
			value: "ATTACH_AGENT_IMAGE",
		},
		DETTACH_AGENT_IMAGE: UpdateTaskSpeedReqSubtaskName{
			value: "DETTACH_AGENT_IMAGE",
		},
		FORMAT_DISK_LINUX: UpdateTaskSpeedReqSubtaskName{
			value: "FORMAT_DISK_LINUX",
		},
		FORMAT_DISK_LINUX_FILE: UpdateTaskSpeedReqSubtaskName{
			value: "FORMAT_DISK_LINUX_FILE",
		},
		FORMAT_DISK_LINUX_BLOCK: UpdateTaskSpeedReqSubtaskName{
			value: "FORMAT_DISK_LINUX_BLOCK",
		},
		FORMAT_DISK_WINDOWS: UpdateTaskSpeedReqSubtaskName{
			value: "FORMAT_DISK_WINDOWS",
		},
		MIGRATE_LINUX_FILE: UpdateTaskSpeedReqSubtaskName{
			value: "MIGRATE_LINUX_FILE",
		},
		MIGRATE_LINUX_BLOCK: UpdateTaskSpeedReqSubtaskName{
			value: "MIGRATE_LINUX_BLOCK",
		},
		MIGRATE_WINDOWS_BLOCK: UpdateTaskSpeedReqSubtaskName{
			value: "MIGRATE_WINDOWS_BLOCK",
		},
		CLONE_VM: UpdateTaskSpeedReqSubtaskName{
			value: "CLONE_VM",
		},
		SYNC_LINUX_FILE: UpdateTaskSpeedReqSubtaskName{
			value: "SYNC_LINUX_FILE",
		},
		SYNC_LINUX_BLOCK: UpdateTaskSpeedReqSubtaskName{
			value: "SYNC_LINUX_BLOCK",
		},
		SYNC_WINDOWS_BLOCK: UpdateTaskSpeedReqSubtaskName{
			value: "SYNC_WINDOWS_BLOCK",
		},
		CONFIGURE_LINUX: UpdateTaskSpeedReqSubtaskName{
			value: "CONFIGURE_LINUX",
		},
		CONFIGURE_LINUX_BLOCK: UpdateTaskSpeedReqSubtaskName{
			value: "CONFIGURE_LINUX_BLOCK",
		},
		CONFIGURE_LINUX_FILE: UpdateTaskSpeedReqSubtaskName{
			value: "CONFIGURE_LINUX_FILE",
		},
		CONFIGURE_WINDOWS: UpdateTaskSpeedReqSubtaskName{
			value: "CONFIGURE_WINDOWS",
		},
	}
}

func (c UpdateTaskSpeedReqSubtaskName) Value() string {
	return c.value
}

func (c UpdateTaskSpeedReqSubtaskName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTaskSpeedReqSubtaskName) UnmarshalJSON(b []byte) error {
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
