package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ThumbTask struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *ThumbTaskStatus `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	OutputFileName *string `json:"output_file_name,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Description *string `json:"description,omitempty"`

	ThumbnailInfo *[]PicInfo `json:"thumbnail_info,omitempty"`
}

func (o ThumbTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThumbTask struct{}"
	}

	return strings.Join([]string{"ThumbTask", string(data)}, " ")
}

type ThumbTaskStatus struct {
	value string
}

type ThumbTaskStatusEnum struct {
	NO_TASK    ThumbTaskStatus
	WAITING    ThumbTaskStatus
	PROCESSING ThumbTaskStatus
	SUCCEEDED  ThumbTaskStatus
	FAILED     ThumbTaskStatus
	CANCELED   ThumbTaskStatus
}

func GetThumbTaskStatusEnum() ThumbTaskStatusEnum {
	return ThumbTaskStatusEnum{
		NO_TASK: ThumbTaskStatus{
			value: "NO_TASK",
		},
		WAITING: ThumbTaskStatus{
			value: "WAITING",
		},
		PROCESSING: ThumbTaskStatus{
			value: "PROCESSING",
		},
		SUCCEEDED: ThumbTaskStatus{
			value: "SUCCEEDED",
		},
		FAILED: ThumbTaskStatus{
			value: "FAILED",
		},
		CANCELED: ThumbTaskStatus{
			value: "CANCELED",
		},
	}
}

func (c ThumbTaskStatus) Value() string {
	return c.value
}

func (c ThumbTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThumbTaskStatus) UnmarshalJSON(b []byte) error {
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
