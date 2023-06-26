package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListAnimatedGraphicsTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	XLanguage *string `json:"x-language,omitempty"`

	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListAnimatedGraphicsTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListAnimatedGraphicsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnimatedGraphicsTaskRequest struct{}"
	}

	return strings.Join([]string{"ListAnimatedGraphicsTaskRequest", string(data)}, " ")
}

type ListAnimatedGraphicsTaskRequestStatus struct {
	value string
}

type ListAnimatedGraphicsTaskRequestStatusEnum struct {
	INIT          ListAnimatedGraphicsTaskRequestStatus
	WAITING       ListAnimatedGraphicsTaskRequestStatus
	PREPROCESSING ListAnimatedGraphicsTaskRequestStatus
	SUCCEED       ListAnimatedGraphicsTaskRequestStatus
	FAILED        ListAnimatedGraphicsTaskRequestStatus
	CANCELED      ListAnimatedGraphicsTaskRequestStatus
}

func GetListAnimatedGraphicsTaskRequestStatusEnum() ListAnimatedGraphicsTaskRequestStatusEnum {
	return ListAnimatedGraphicsTaskRequestStatusEnum{
		INIT: ListAnimatedGraphicsTaskRequestStatus{
			value: "INIT",
		},
		WAITING: ListAnimatedGraphicsTaskRequestStatus{
			value: "WAITING",
		},
		PREPROCESSING: ListAnimatedGraphicsTaskRequestStatus{
			value: "PREPROCESSING",
		},
		SUCCEED: ListAnimatedGraphicsTaskRequestStatus{
			value: "SUCCEED",
		},
		FAILED: ListAnimatedGraphicsTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListAnimatedGraphicsTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListAnimatedGraphicsTaskRequestStatus) Value() string {
	return c.value
}

func (c ListAnimatedGraphicsTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAnimatedGraphicsTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
