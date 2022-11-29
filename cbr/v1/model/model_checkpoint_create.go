package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckpointCreate struct {
	CreatedAt string `json:"created_at"`

	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	Status CheckpointCreateStatus `json:"status"`

	Vault *CheckpointPlanCreate `json:"vault,omitempty"`

	ExtraInfo *CheckpointExtraInfoResp `json:"extra_info,omitempty"`
}

func (o CheckpointCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointCreate struct{}"
	}

	return strings.Join([]string{"CheckpointCreate", string(data)}, " ")
}

type CheckpointCreateStatus struct {
	value string
}

type CheckpointCreateStatusEnum struct {
	AVAILABLE      CheckpointCreateStatus
	DELETING       CheckpointCreateStatus
	PROTECTING     CheckpointCreateStatus
	DELETED        CheckpointCreateStatus
	ERROR_DELETING CheckpointCreateStatus
	ERROR          CheckpointCreateStatus
}

func GetCheckpointCreateStatusEnum() CheckpointCreateStatusEnum {
	return CheckpointCreateStatusEnum{
		AVAILABLE: CheckpointCreateStatus{
			value: "available",
		},
		DELETING: CheckpointCreateStatus{
			value: "deleting",
		},
		PROTECTING: CheckpointCreateStatus{
			value: "protecting",
		},
		DELETED: CheckpointCreateStatus{
			value: "deleted",
		},
		ERROR_DELETING: CheckpointCreateStatus{
			value: "error-deleting",
		},
		ERROR: CheckpointCreateStatus{
			value: "error",
		},
	}
}

func (c CheckpointCreateStatus) Value() string {
	return c.value
}

func (c CheckpointCreateStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckpointCreateStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
