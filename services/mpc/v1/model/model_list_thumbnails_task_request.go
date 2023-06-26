package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListThumbnailsTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	XLanguage *string `json:"x-language,omitempty"`

	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListThumbnailsTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListThumbnailsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThumbnailsTaskRequest struct{}"
	}

	return strings.Join([]string{"ListThumbnailsTaskRequest", string(data)}, " ")
}

type ListThumbnailsTaskRequestStatus struct {
	value string
}

type ListThumbnailsTaskRequestStatusEnum struct {
	WAITING    ListThumbnailsTaskRequestStatus
	PROCESSING ListThumbnailsTaskRequestStatus
	SUCCEEDED  ListThumbnailsTaskRequestStatus
	FAILED     ListThumbnailsTaskRequestStatus
	CANCELED   ListThumbnailsTaskRequestStatus
}

func GetListThumbnailsTaskRequestStatusEnum() ListThumbnailsTaskRequestStatusEnum {
	return ListThumbnailsTaskRequestStatusEnum{
		WAITING: ListThumbnailsTaskRequestStatus{
			value: "WAITING",
		},
		PROCESSING: ListThumbnailsTaskRequestStatus{
			value: "PROCESSING",
		},
		SUCCEEDED: ListThumbnailsTaskRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListThumbnailsTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListThumbnailsTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListThumbnailsTaskRequestStatus) Value() string {
	return c.value
}

func (c ListThumbnailsTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListThumbnailsTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
