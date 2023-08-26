package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type AgenciesTaskReq struct {
	ProjectId *string `json:"project_id,omitempty"`

	OperateType AgenciesTaskReqOperateType `json:"operate_type"`
}

func (o AgenciesTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgenciesTaskReq struct{}"
	}

	return strings.Join([]string{"AgenciesTaskReq", string(data)}, " ")
}

type AgenciesTaskReqOperateType struct {
	value string
}

type AgenciesTaskReqOperateTypeEnum struct {
	CREATED  AgenciesTaskReqOperateType
	CANCELED AgenciesTaskReqOperateType
}

func GetAgenciesTaskReqOperateTypeEnum() AgenciesTaskReqOperateTypeEnum {
	return AgenciesTaskReqOperateTypeEnum{
		CREATED: AgenciesTaskReqOperateType{
			value: "CREATED",
		},
		CANCELED: AgenciesTaskReqOperateType{
			value: "CANCELED",
		},
	}
}

func (c AgenciesTaskReqOperateType) Value() string {
	return c.value
}

func (c AgenciesTaskReqOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgenciesTaskReqOperateType) UnmarshalJSON(b []byte) error {
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
