package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListRemuxTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListRemuxTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	InputBucket *string `json:"input_bucket,omitempty"`

	InputObject *string `json:"input_object,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"ListRemuxTaskRequest", string(data)}, " ")
}

type ListRemuxTaskRequestStatus struct {
	value string
}

type ListRemuxTaskRequestStatusEnum struct {
	INIT       ListRemuxTaskRequestStatus
	WAITING    ListRemuxTaskRequestStatus
	PROCESSING ListRemuxTaskRequestStatus
	SUCCEED    ListRemuxTaskRequestStatus
	FAILED     ListRemuxTaskRequestStatus
	CANCELED   ListRemuxTaskRequestStatus
}

func GetListRemuxTaskRequestStatusEnum() ListRemuxTaskRequestStatusEnum {
	return ListRemuxTaskRequestStatusEnum{
		INIT: ListRemuxTaskRequestStatus{
			value: "INIT",
		},
		WAITING: ListRemuxTaskRequestStatus{
			value: "WAITING",
		},
		PROCESSING: ListRemuxTaskRequestStatus{
			value: "PROCESSING",
		},
		SUCCEED: ListRemuxTaskRequestStatus{
			value: "SUCCEED",
		},
		FAILED: ListRemuxTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListRemuxTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListRemuxTaskRequestStatus) Value() string {
	return c.value
}

func (c ListRemuxTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRemuxTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
