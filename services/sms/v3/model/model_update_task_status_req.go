package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type UpdateTaskStatusReq struct {
	Operation UpdateTaskStatusReqOperation `json:"operation"`

	Param map[string]string `json:"param,omitempty"`

	SwitchHce *bool `json:"switch_hce,omitempty"`
}

func (o UpdateTaskStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusReq struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusReq", string(data)}, " ")
}

type UpdateTaskStatusReqOperation struct {
	value string
}

type UpdateTaskStatusReqOperationEnum struct {
	START                UpdateTaskStatusReqOperation
	STOP                 UpdateTaskStatusReqOperation
	COLLECT_LOG          UpdateTaskStatusReqOperation
	TEST                 UpdateTaskStatusReqOperation
	CLONE_TEST           UpdateTaskStatusReqOperation
	RESTART              UpdateTaskStatusReqOperation
	SYNC_FAILED_ROLLBACK UpdateTaskStatusReqOperation
}

func GetUpdateTaskStatusReqOperationEnum() UpdateTaskStatusReqOperationEnum {
	return UpdateTaskStatusReqOperationEnum{
		START: UpdateTaskStatusReqOperation{
			value: "start",
		},
		STOP: UpdateTaskStatusReqOperation{
			value: "stop",
		},
		COLLECT_LOG: UpdateTaskStatusReqOperation{
			value: "collect_log",
		},
		TEST: UpdateTaskStatusReqOperation{
			value: "test",
		},
		CLONE_TEST: UpdateTaskStatusReqOperation{
			value: "clone_test",
		},
		RESTART: UpdateTaskStatusReqOperation{
			value: "restart",
		},
		SYNC_FAILED_ROLLBACK: UpdateTaskStatusReqOperation{
			value: "sync_failed_rollback",
		},
	}
}

func (c UpdateTaskStatusReqOperation) Value() string {
	return c.value
}

func (c UpdateTaskStatusReqOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTaskStatusReqOperation) UnmarshalJSON(b []byte) error {
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
