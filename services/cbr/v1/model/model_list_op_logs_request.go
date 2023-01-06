package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListOpLogsRequest struct {
	EndTime *string `json:"end_time,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	OperationType *ListOpLogsRequestOperationType `json:"operation_type,omitempty"`

	ProviderId *string `json:"provider_id,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	Status *ListOpLogsRequestStatus `json:"status,omitempty"`

	VaultId *string `json:"vault_id,omitempty"`

	VaultName *string `json:"vault_name,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListOpLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpLogsRequest struct{}"
	}

	return strings.Join([]string{"ListOpLogsRequest", string(data)}, " ")
}

type ListOpLogsRequestOperationType struct {
	value string
}

type ListOpLogsRequestOperationTypeEnum struct {
	BACKUP          ListOpLogsRequestOperationType
	COPY            ListOpLogsRequestOperationType
	REPLICATION     ListOpLogsRequestOperationType
	DELETE          ListOpLogsRequestOperationType
	RESTORE         ListOpLogsRequestOperationType
	VAULT_DELETE    ListOpLogsRequestOperationType
	REMOVE_RESOURCE ListOpLogsRequestOperationType
	SYNC            ListOpLogsRequestOperationType
}

func GetListOpLogsRequestOperationTypeEnum() ListOpLogsRequestOperationTypeEnum {
	return ListOpLogsRequestOperationTypeEnum{
		BACKUP: ListOpLogsRequestOperationType{
			value: "backup",
		},
		COPY: ListOpLogsRequestOperationType{
			value: "copy",
		},
		REPLICATION: ListOpLogsRequestOperationType{
			value: "replication",
		},
		DELETE: ListOpLogsRequestOperationType{
			value: "delete",
		},
		RESTORE: ListOpLogsRequestOperationType{
			value: "restore",
		},
		VAULT_DELETE: ListOpLogsRequestOperationType{
			value: "vault_delete",
		},
		REMOVE_RESOURCE: ListOpLogsRequestOperationType{
			value: "remove_resource",
		},
		SYNC: ListOpLogsRequestOperationType{
			value: "sync",
		},
	}
}

func (c ListOpLogsRequestOperationType) Value() string {
	return c.value
}

func (c ListOpLogsRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOpLogsRequestOperationType) UnmarshalJSON(b []byte) error {
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

type ListOpLogsRequestStatus struct {
	value string
}

type ListOpLogsRequestStatusEnum struct {
	SUCCESS ListOpLogsRequestStatus
	SKIPPED ListOpLogsRequestStatus
	FAILED  ListOpLogsRequestStatus
	RUNNING ListOpLogsRequestStatus
	TIMEOUT ListOpLogsRequestStatus
	WAITING ListOpLogsRequestStatus
}

func GetListOpLogsRequestStatusEnum() ListOpLogsRequestStatusEnum {
	return ListOpLogsRequestStatusEnum{
		SUCCESS: ListOpLogsRequestStatus{
			value: "success",
		},
		SKIPPED: ListOpLogsRequestStatus{
			value: "skipped",
		},
		FAILED: ListOpLogsRequestStatus{
			value: "failed",
		},
		RUNNING: ListOpLogsRequestStatus{
			value: "running",
		},
		TIMEOUT: ListOpLogsRequestStatus{
			value: "timeout",
		},
		WAITING: ListOpLogsRequestStatus{
			value: "waiting",
		},
	}
}

func (c ListOpLogsRequestStatus) Value() string {
	return c.value
}

func (c ListOpLogsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOpLogsRequestStatus) UnmarshalJSON(b []byte) error {
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
