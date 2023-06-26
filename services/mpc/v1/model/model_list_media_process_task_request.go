package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListMediaProcessTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListMediaProcessTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListMediaProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMediaProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"ListMediaProcessTaskRequest", string(data)}, " ")
}

type ListMediaProcessTaskRequestStatus struct {
	value string
}

type ListMediaProcessTaskRequestStatusEnum struct {
	WAITING    ListMediaProcessTaskRequestStatus
	PROCESSING ListMediaProcessTaskRequestStatus
	SUCCEEDED  ListMediaProcessTaskRequestStatus
	FAILED     ListMediaProcessTaskRequestStatus
	CANCELED   ListMediaProcessTaskRequestStatus
}

func GetListMediaProcessTaskRequestStatusEnum() ListMediaProcessTaskRequestStatusEnum {
	return ListMediaProcessTaskRequestStatusEnum{
		WAITING: ListMediaProcessTaskRequestStatus{
			value: "WAITING",
		},
		PROCESSING: ListMediaProcessTaskRequestStatus{
			value: "PROCESSING",
		},
		SUCCEEDED: ListMediaProcessTaskRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListMediaProcessTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListMediaProcessTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListMediaProcessTaskRequestStatus) Value() string {
	return c.value
}

func (c ListMediaProcessTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMediaProcessTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
