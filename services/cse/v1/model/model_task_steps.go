package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TaskSteps struct {
	TaskName *string `json:"task_name,omitempty"`

	TaskNames *[]string `json:"task_names,omitempty"`

	Status *TaskStepsStatus `json:"status,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	TaskExecutorBrief *TaskExecutorBrief `json:"task_executor_brief,omitempty"`

	Tasks *[]Task `json:"tasks,omitempty"`
}

func (o TaskSteps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSteps struct{}"
	}

	return strings.Join([]string{"TaskSteps", string(data)}, " ")
}

type TaskStepsStatus struct {
	value string
}

type TaskStepsStatusEnum struct {
	INIT      TaskStepsStatus
	EXECUTING TaskStepsStatus
	ERROR     TaskStepsStatus
	TIMEOUT   TaskStepsStatus
	FINISHED  TaskStepsStatus
}

func GetTaskStepsStatusEnum() TaskStepsStatusEnum {
	return TaskStepsStatusEnum{
		INIT: TaskStepsStatus{
			value: "Init",
		},
		EXECUTING: TaskStepsStatus{
			value: "Executing",
		},
		ERROR: TaskStepsStatus{
			value: "Error",
		},
		TIMEOUT: TaskStepsStatus{
			value: "Timeout",
		},
		FINISHED: TaskStepsStatus{
			value: "Finished",
		},
	}
}

func (c TaskStepsStatus) Value() string {
	return c.value
}

func (c TaskStepsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskStepsStatus) UnmarshalJSON(b []byte) error {
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
