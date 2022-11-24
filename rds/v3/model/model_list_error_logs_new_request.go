package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListErrorLogsNewRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	Offset *int64 `json:"offset,omitempty"`

	Limit *int64 `json:"limit,omitempty"`

	Level *ListErrorLogsNewRequestLevel `json:"level,omitempty"`
}

func (o ListErrorLogsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorLogsNewRequest struct{}"
	}

	return strings.Join([]string{"ListErrorLogsNewRequest", string(data)}, " ")
}

type ListErrorLogsNewRequestLevel struct {
	value string
}

type ListErrorLogsNewRequestLevelEnum struct {
	ALL     ListErrorLogsNewRequestLevel
	INFO    ListErrorLogsNewRequestLevel
	LOG     ListErrorLogsNewRequestLevel
	WARNING ListErrorLogsNewRequestLevel
	ERROR   ListErrorLogsNewRequestLevel
	FATAL   ListErrorLogsNewRequestLevel
	PANIC   ListErrorLogsNewRequestLevel
	NOTE    ListErrorLogsNewRequestLevel
}

func GetListErrorLogsNewRequestLevelEnum() ListErrorLogsNewRequestLevelEnum {
	return ListErrorLogsNewRequestLevelEnum{
		ALL: ListErrorLogsNewRequestLevel{
			value: "ALL",
		},
		INFO: ListErrorLogsNewRequestLevel{
			value: "INFO",
		},
		LOG: ListErrorLogsNewRequestLevel{
			value: "LOG",
		},
		WARNING: ListErrorLogsNewRequestLevel{
			value: "WARNING",
		},
		ERROR: ListErrorLogsNewRequestLevel{
			value: "ERROR",
		},
		FATAL: ListErrorLogsNewRequestLevel{
			value: "FATAL",
		},
		PANIC: ListErrorLogsNewRequestLevel{
			value: "PANIC",
		},
		NOTE: ListErrorLogsNewRequestLevel{
			value: "NOTE",
		},
	}
}

func (c ListErrorLogsNewRequestLevel) Value() string {
	return c.value
}

func (c ListErrorLogsNewRequestLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListErrorLogsNewRequestLevel) UnmarshalJSON(b []byte) error {
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
