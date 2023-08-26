package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListEncryptTaskRequest struct {
	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListEncryptTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListEncryptTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"ListEncryptTaskRequest", string(data)}, " ")
}

type ListEncryptTaskRequestStatus struct {
	value string
}

type ListEncryptTaskRequestStatusEnum struct {
	WAITING    ListEncryptTaskRequestStatus
	PROCESSING ListEncryptTaskRequestStatus
	SUCCEEDED  ListEncryptTaskRequestStatus
	FAILED     ListEncryptTaskRequestStatus
	CANCELED   ListEncryptTaskRequestStatus
}

func GetListEncryptTaskRequestStatusEnum() ListEncryptTaskRequestStatusEnum {
	return ListEncryptTaskRequestStatusEnum{
		WAITING: ListEncryptTaskRequestStatus{
			value: "WAITING",
		},
		PROCESSING: ListEncryptTaskRequestStatus{
			value: "PROCESSING",
		},
		SUCCEEDED: ListEncryptTaskRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListEncryptTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListEncryptTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListEncryptTaskRequestStatus) Value() string {
	return c.value
}

func (c ListEncryptTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEncryptTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
