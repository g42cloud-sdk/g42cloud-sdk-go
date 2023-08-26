package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type CheckpointResourceResp struct {
	ExtraInfo *string `json:"extra_info,omitempty"`

	Id string `json:"id"`

	Name string `json:"name"`

	ProtectStatus *CheckpointResourceRespProtectStatus `json:"protect_status,omitempty"`

	ResourceSize *string `json:"resource_size,omitempty"`

	Type string `json:"type"`

	BackupSize *string `json:"backup_size,omitempty"`

	BackupCount *string `json:"backup_count,omitempty"`
}

func (o CheckpointResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointResourceResp struct{}"
	}

	return strings.Join([]string{"CheckpointResourceResp", string(data)}, " ")
}

type CheckpointResourceRespProtectStatus struct {
	value string
}

type CheckpointResourceRespProtectStatusEnum struct {
	AVAILABLE  CheckpointResourceRespProtectStatus
	ERROR      CheckpointResourceRespProtectStatus
	PROTECTING CheckpointResourceRespProtectStatus
	RESTORING  CheckpointResourceRespProtectStatus
	REMOVING   CheckpointResourceRespProtectStatus
}

func GetCheckpointResourceRespProtectStatusEnum() CheckpointResourceRespProtectStatusEnum {
	return CheckpointResourceRespProtectStatusEnum{
		AVAILABLE: CheckpointResourceRespProtectStatus{
			value: "available",
		},
		ERROR: CheckpointResourceRespProtectStatus{
			value: "error",
		},
		PROTECTING: CheckpointResourceRespProtectStatus{
			value: "protecting",
		},
		RESTORING: CheckpointResourceRespProtectStatus{
			value: "restoring",
		},
		REMOVING: CheckpointResourceRespProtectStatus{
			value: "removing",
		},
	}
}

func (c CheckpointResourceRespProtectStatus) Value() string {
	return c.value
}

func (c CheckpointResourceRespProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckpointResourceRespProtectStatus) UnmarshalJSON(b []byte) error {
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
