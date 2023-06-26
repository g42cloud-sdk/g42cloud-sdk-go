package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListEditingJobRequest struct {
	XLanguage *string `json:"x-language,omitempty"`

	JobId *[]string `json:"job_id,omitempty"`

	Status *ListEditingJobRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListEditingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEditingJobRequest struct{}"
	}

	return strings.Join([]string{"ListEditingJobRequest", string(data)}, " ")
}

type ListEditingJobRequestStatus struct {
	value string
}

type ListEditingJobRequestStatusEnum struct {
	INIT          ListEditingJobRequestStatus
	WAITING       ListEditingJobRequestStatus
	PREPROCESSING ListEditingJobRequestStatus
	SUCCEED       ListEditingJobRequestStatus
	FAILED        ListEditingJobRequestStatus
	CANCELED      ListEditingJobRequestStatus
}

func GetListEditingJobRequestStatusEnum() ListEditingJobRequestStatusEnum {
	return ListEditingJobRequestStatusEnum{
		INIT: ListEditingJobRequestStatus{
			value: "INIT",
		},
		WAITING: ListEditingJobRequestStatus{
			value: "WAITING",
		},
		PREPROCESSING: ListEditingJobRequestStatus{
			value: "PREPROCESSING",
		},
		SUCCEED: ListEditingJobRequestStatus{
			value: "SUCCEED",
		},
		FAILED: ListEditingJobRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListEditingJobRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListEditingJobRequestStatus) Value() string {
	return c.value
}

func (c ListEditingJobRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEditingJobRequestStatus) UnmarshalJSON(b []byte) error {
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
