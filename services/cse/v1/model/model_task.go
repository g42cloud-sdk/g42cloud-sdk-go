package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Task struct {
	JobId *int32 `json:"job_id,omitempty"`

	Id *int64 `json:"id,omitempty"`

	Type *TaskType `json:"type,omitempty"`

	Assigned *string `json:"assigned,omitempty"`

	TaskName *string `json:"task_name,omitempty"`

	EngineName *string `json:"engine_name,omitempty"`

	TaskOrder *int32 `json:"task_order,omitempty"`

	Status *TaskStatus `json:"status,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	Timeout *int32 `json:"timeout,omitempty"`

	Log *string `json:"log,omitempty"`

	Output *string `json:"output,omitempty"`

	TaskExecutorBrief *TaskExecutorBrief `json:"task_executor_brief,omitempty"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}

type TaskType struct {
	value string
}

type TaskTypeEnum struct {
	CREATE  TaskType
	DELETE  TaskType
	UPGRADE TaskType
	MODIFY  TaskType
}

func GetTaskTypeEnum() TaskTypeEnum {
	return TaskTypeEnum{
		CREATE: TaskType{
			value: "Create",
		},
		DELETE: TaskType{
			value: "Delete",
		},
		UPGRADE: TaskType{
			value: "Upgrade",
		},
		MODIFY: TaskType{
			value: "Modify",
		},
	}
}

func (c TaskType) Value() string {
	return c.value
}

func (c TaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskType) UnmarshalJSON(b []byte) error {
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

type TaskStatus struct {
	value string
}

type TaskStatusEnum struct {
	INIT      TaskStatus
	EXECUTING TaskStatus
	ERROR     TaskStatus
	TIMEOUT   TaskStatus
	FINISHED  TaskStatus
}

func GetTaskStatusEnum() TaskStatusEnum {
	return TaskStatusEnum{
		INIT: TaskStatus{
			value: "Init",
		},
		EXECUTING: TaskStatus{
			value: "Executing",
		},
		ERROR: TaskStatus{
			value: "Error",
		},
		TIMEOUT: TaskStatus{
			value: "Timeout",
		},
		FINISHED: TaskStatus{
			value: "Finished",
		},
	}
}

func (c TaskStatus) Value() string {
	return c.value
}

func (c TaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskStatus) UnmarshalJSON(b []byte) error {
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
