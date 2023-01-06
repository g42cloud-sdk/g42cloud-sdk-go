package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Match struct {
	Key MatchKey `json:"key"`

	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}

type MatchKey struct {
	value string
}

type MatchKeyEnum struct {
	RESOURCE_NAME MatchKey
	SERVICE_TYPE  MatchKey
}

func GetMatchKeyEnum() MatchKeyEnum {
	return MatchKeyEnum{
		RESOURCE_NAME: MatchKey{
			value: "resource_name",
		},
		SERVICE_TYPE: MatchKey{
			value: "service_type",
		},
	}
}

func (c MatchKey) Value() string {
	return c.value
}

func (c MatchKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MatchKey) UnmarshalJSON(b []byte) error {
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
