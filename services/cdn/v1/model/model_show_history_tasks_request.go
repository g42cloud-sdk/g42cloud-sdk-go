package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowHistoryTasksRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNumber *int32 `json:"page_number,omitempty"`

	Status *ShowHistoryTasksRequestStatus `json:"status,omitempty"`

	StartDate *int64 `json:"start_date,omitempty"`

	EndDate *int64 `json:"end_date,omitempty"`

	OrderField *string `json:"order_field,omitempty"`

	OrderType *string `json:"order_type,omitempty"`

	FileType *ShowHistoryTasksRequestFileType `json:"file_type,omitempty"`
}

func (o ShowHistoryTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTasksRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryTasksRequest", string(data)}, " ")
}

type ShowHistoryTasksRequestStatus struct {
	value string
}

type ShowHistoryTasksRequestStatusEnum struct {
	TASK_INPROCESS ShowHistoryTasksRequestStatus
	TASK_DONE      ShowHistoryTasksRequestStatus
}

func GetShowHistoryTasksRequestStatusEnum() ShowHistoryTasksRequestStatusEnum {
	return ShowHistoryTasksRequestStatusEnum{
		TASK_INPROCESS: ShowHistoryTasksRequestStatus{
			value: "task_inprocess",
		},
		TASK_DONE: ShowHistoryTasksRequestStatus{
			value: "task_done",
		},
	}
}

func (c ShowHistoryTasksRequestStatus) Value() string {
	return c.value
}

func (c ShowHistoryTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHistoryTasksRequestStatus) UnmarshalJSON(b []byte) error {
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

type ShowHistoryTasksRequestFileType struct {
	value string
}

type ShowHistoryTasksRequestFileTypeEnum struct {
	FILE      ShowHistoryTasksRequestFileType
	DIRECTORY ShowHistoryTasksRequestFileType
}

func GetShowHistoryTasksRequestFileTypeEnum() ShowHistoryTasksRequestFileTypeEnum {
	return ShowHistoryTasksRequestFileTypeEnum{
		FILE: ShowHistoryTasksRequestFileType{
			value: "file",
		},
		DIRECTORY: ShowHistoryTasksRequestFileType{
			value: "directory",
		},
	}
}

func (c ShowHistoryTasksRequestFileType) Value() string {
	return c.value
}

func (c ShowHistoryTasksRequestFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHistoryTasksRequestFileType) UnmarshalJSON(b []byte) error {
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
