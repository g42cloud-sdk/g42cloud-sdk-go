package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListResetTracksTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListResetTracksTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListResetTracksTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"ListResetTracksTaskRequest", string(data)}, " ")
}

type ListResetTracksTaskRequestStatus struct {
	value string
}

type ListResetTracksTaskRequestStatusEnum struct {
	WAITING    ListResetTracksTaskRequestStatus
	PROCESSING ListResetTracksTaskRequestStatus
	SUCCEEDED  ListResetTracksTaskRequestStatus
	FAILED     ListResetTracksTaskRequestStatus
	CANCELED   ListResetTracksTaskRequestStatus
}

func GetListResetTracksTaskRequestStatusEnum() ListResetTracksTaskRequestStatusEnum {
	return ListResetTracksTaskRequestStatusEnum{
		WAITING: ListResetTracksTaskRequestStatus{
			value: "WAITING",
		},
		PROCESSING: ListResetTracksTaskRequestStatus{
			value: "PROCESSING",
		},
		SUCCEEDED: ListResetTracksTaskRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListResetTracksTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListResetTracksTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListResetTracksTaskRequestStatus) Value() string {
	return c.value
}

func (c ListResetTracksTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResetTracksTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
