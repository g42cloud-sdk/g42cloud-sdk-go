package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowEngineJobResponse struct {
	Id *int32 `json:"id,omitempty"`

	EngineId *string `json:"engine_id,omitempty"`

	Type *ShowEngineJobResponseType `json:"type,omitempty"`

	Description *string `json:"description,omitempty"`

	Status *ShowEngineJobResponseStatus `json:"status,omitempty"`

	Scheduling *int32 `json:"scheduling,omitempty"`

	CreateUser *string `json:"create_user,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	Context *string `json:"context,omitempty"`

	Tasks          *[]TaskSteps `json:"tasks,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowEngineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineJobResponse struct{}"
	}

	return strings.Join([]string{"ShowEngineJobResponse", string(data)}, " ")
}

type ShowEngineJobResponseType struct {
	value string
}

type ShowEngineJobResponseTypeEnum struct {
	CREATE  ShowEngineJobResponseType
	DELETE  ShowEngineJobResponseType
	UPGRADE ShowEngineJobResponseType
	MODIFY  ShowEngineJobResponseType
}

func GetShowEngineJobResponseTypeEnum() ShowEngineJobResponseTypeEnum {
	return ShowEngineJobResponseTypeEnum{
		CREATE: ShowEngineJobResponseType{
			value: "Create",
		},
		DELETE: ShowEngineJobResponseType{
			value: "Delete",
		},
		UPGRADE: ShowEngineJobResponseType{
			value: "Upgrade",
		},
		MODIFY: ShowEngineJobResponseType{
			value: "Modify",
		},
	}
}

func (c ShowEngineJobResponseType) Value() string {
	return c.value
}

func (c ShowEngineJobResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineJobResponseType) UnmarshalJSON(b []byte) error {
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

type ShowEngineJobResponseStatus struct {
	value string
}

type ShowEngineJobResponseStatusEnum struct {
	INIT      ShowEngineJobResponseStatus
	EXECUTING ShowEngineJobResponseStatus
	ERROR     ShowEngineJobResponseStatus
	TIMEOUT   ShowEngineJobResponseStatus
	FINISHED  ShowEngineJobResponseStatus
}

func GetShowEngineJobResponseStatusEnum() ShowEngineJobResponseStatusEnum {
	return ShowEngineJobResponseStatusEnum{
		INIT: ShowEngineJobResponseStatus{
			value: "Init",
		},
		EXECUTING: ShowEngineJobResponseStatus{
			value: "Executing",
		},
		ERROR: ShowEngineJobResponseStatus{
			value: "Error",
		},
		TIMEOUT: ShowEngineJobResponseStatus{
			value: "Timeout",
		},
		FINISHED: ShowEngineJobResponseStatus{
			value: "Finished",
		},
	}
}

func (c ShowEngineJobResponseStatus) Value() string {
	return c.value
}

func (c ShowEngineJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineJobResponseStatus) UnmarshalJSON(b []byte) error {
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
