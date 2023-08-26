package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TasksObject struct {
	Id *string `json:"id,omitempty"`

	TaskType *TasksObjectTaskType `json:"task_type,omitempty"`

	Status *string `json:"status,omitempty"`

	Processing *int32 `json:"processing,omitempty"`

	Succeed *int32 `json:"succeed,omitempty"`

	Failed *int32 `json:"failed,omitempty"`

	Total *int32 `json:"total,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	FileType *TasksObjectFileType `json:"file_type,omitempty"`
}

func (o TasksObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TasksObject struct{}"
	}

	return strings.Join([]string{"TasksObject", string(data)}, " ")
}

type TasksObjectTaskType struct {
	value string
}

type TasksObjectTaskTypeEnum struct {
	REFRESH    TasksObjectTaskType
	PREHEATING TasksObjectTaskType
}

func GetTasksObjectTaskTypeEnum() TasksObjectTaskTypeEnum {
	return TasksObjectTaskTypeEnum{
		REFRESH: TasksObjectTaskType{
			value: "refresh",
		},
		PREHEATING: TasksObjectTaskType{
			value: "preheating",
		},
	}
}

func (c TasksObjectTaskType) Value() string {
	return c.value
}

func (c TasksObjectTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksObjectTaskType) UnmarshalJSON(b []byte) error {
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

type TasksObjectFileType struct {
	value string
}

type TasksObjectFileTypeEnum struct {
	FILE      TasksObjectFileType
	DIRECTORY TasksObjectFileType
}

func GetTasksObjectFileTypeEnum() TasksObjectFileTypeEnum {
	return TasksObjectFileTypeEnum{
		FILE: TasksObjectFileType{
			value: "file",
		},
		DIRECTORY: TasksObjectFileType{
			value: "directory",
		},
	}
}

func (c TasksObjectFileType) Value() string {
	return c.value
}

func (c TasksObjectFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksObjectFileType) UnmarshalJSON(b []byte) error {
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
