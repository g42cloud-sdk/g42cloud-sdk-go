package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListMergeChannelsTaskRequest struct {
	TaskId *[]string `json:"task_id,omitempty"`

	Status *ListMergeChannelsTaskRequestStatus `json:"status,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Size *int32 `json:"size,omitempty"`
}

func (o ListMergeChannelsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeChannelsTaskRequest struct{}"
	}

	return strings.Join([]string{"ListMergeChannelsTaskRequest", string(data)}, " ")
}

type ListMergeChannelsTaskRequestStatus struct {
	value string
}

type ListMergeChannelsTaskRequestStatusEnum struct {
	WAITING    ListMergeChannelsTaskRequestStatus
	PROCESSING ListMergeChannelsTaskRequestStatus
	SUCCEEDED  ListMergeChannelsTaskRequestStatus
	FAILED     ListMergeChannelsTaskRequestStatus
	CANCELED   ListMergeChannelsTaskRequestStatus
}

func GetListMergeChannelsTaskRequestStatusEnum() ListMergeChannelsTaskRequestStatusEnum {
	return ListMergeChannelsTaskRequestStatusEnum{
		WAITING: ListMergeChannelsTaskRequestStatus{
			value: "WAITING",
		},
		PROCESSING: ListMergeChannelsTaskRequestStatus{
			value: "PROCESSING",
		},
		SUCCEEDED: ListMergeChannelsTaskRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListMergeChannelsTaskRequestStatus{
			value: "FAILED",
		},
		CANCELED: ListMergeChannelsTaskRequestStatus{
			value: "CANCELED",
		},
	}
}

func (c ListMergeChannelsTaskRequestStatus) Value() string {
	return c.value
}

func (c ListMergeChannelsTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeChannelsTaskRequestStatus) UnmarshalJSON(b []byte) error {
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
