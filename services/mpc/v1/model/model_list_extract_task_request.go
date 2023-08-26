package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListExtractTaskRequest struct {
	XLanguage *string `json:"x-language,omitempty"`

	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListExtractTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListExtractTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtractTaskRequest struct{}"
	}

	return strings.Join([]string{"ListExtractTaskRequest", string(data)}, " ")
}

type ListExtractTaskRequestStatus struct {
	value string
}

type ListExtractTaskRequestStatusEnum struct {
	INIT          ListExtractTaskRequestStatus
	WAITING       ListExtractTaskRequestStatus
	PREPROCESSING ListExtractTaskRequestStatus
	SUCCEED       ListExtractTaskRequestStatus
	FAILED        ListExtractTaskRequestStatus
	CANCELED      ListExtractTaskRequestStatus
}

func GetListExtractTaskRequestStatusEnum() ListExtractTaskRequestStatusEnum {
	return ListExtractTaskRequestStatusEnum{
		INIT: ListExtractTaskRequestStatus{
			value: "INIT",
		},
		WAITING: ListExtractTaskRequestStatus{
			value: "WAITING",
		},
		PREPROCESSING: ListExtractTaskRequestStatus{
			value: "PREPROCESSING",
		},
		SUCCEED: ListExtractTaskRequestStatus{
			value: "SUCCEED",
		},
		FAILED: ListExtractTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListExtractTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListExtractTaskRequestStatus) Value() string {
	return c.value
}

func (c ListExtractTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListExtractTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
