package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OperationLog struct {
	CheckpointId *string `json:"checkpoint_id,omitempty"`

	CreatedAt string `json:"created_at"`

	EndedAt *string `json:"ended_at,omitempty"`

	ErrorInfo *OpErrorInfo `json:"error_info"`

	ExtraInfo *OpExtraInfo `json:"extra_info,omitempty"`

	Id string `json:"id"`

	OperationType *OperationLogOperationType `json:"operation_type,omitempty"`

	PolicyId *string `json:"policy_id,omitempty"`

	ProjectId string `json:"project_id"`

	ProviderId *string `json:"provider_id,omitempty"`

	StartedAt string `json:"started_at"`

	Status OperationLogStatus `json:"status"`

	UpdatedAt string `json:"updated_at"`

	VaultId *string `json:"vault_id,omitempty"`

	VaultName *string `json:"vault_name,omitempty"`
}

func (o OperationLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationLog struct{}"
	}

	return strings.Join([]string{"OperationLog", string(data)}, " ")
}

type OperationLogOperationType struct {
	value string
}

type OperationLogOperationTypeEnum struct {
	BACKUP          OperationLogOperationType
	COPY            OperationLogOperationType
	REPLICATION     OperationLogOperationType
	RESTORE         OperationLogOperationType
	DELETE          OperationLogOperationType
	SYNC            OperationLogOperationType
	VAULT_DELETE    OperationLogOperationType
	REMOVE_RESOURCE OperationLogOperationType
}

func GetOperationLogOperationTypeEnum() OperationLogOperationTypeEnum {
	return OperationLogOperationTypeEnum{
		BACKUP: OperationLogOperationType{
			value: "backup",
		},
		COPY: OperationLogOperationType{
			value: "copy",
		},
		REPLICATION: OperationLogOperationType{
			value: "replication",
		},
		RESTORE: OperationLogOperationType{
			value: "restore",
		},
		DELETE: OperationLogOperationType{
			value: "delete",
		},
		SYNC: OperationLogOperationType{
			value: "sync",
		},
		VAULT_DELETE: OperationLogOperationType{
			value: "vault_delete",
		},
		REMOVE_RESOURCE: OperationLogOperationType{
			value: "remove_resource",
		},
	}
}

func (c OperationLogOperationType) Value() string {
	return c.value
}

func (c OperationLogOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationLogOperationType) UnmarshalJSON(b []byte) error {
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

type OperationLogStatus struct {
	value string
}

type OperationLogStatusEnum struct {
	SUCCESS OperationLogStatus
	SKIPPED OperationLogStatus
	FAILED  OperationLogStatus
	RUNNING OperationLogStatus
	TIMEOUT OperationLogStatus
	WAITING OperationLogStatus
}

func GetOperationLogStatusEnum() OperationLogStatusEnum {
	return OperationLogStatusEnum{
		SUCCESS: OperationLogStatus{
			value: "success",
		},
		SKIPPED: OperationLogStatus{
			value: "skipped",
		},
		FAILED: OperationLogStatus{
			value: "failed",
		},
		RUNNING: OperationLogStatus{
			value: "running",
		},
		TIMEOUT: OperationLogStatus{
			value: "timeout",
		},
		WAITING: OperationLogStatus{
			value: "waiting",
		},
	}
}

func (c OperationLogStatus) Value() string {
	return c.value
}

func (c OperationLogStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationLogStatus) UnmarshalJSON(b []byte) error {
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
