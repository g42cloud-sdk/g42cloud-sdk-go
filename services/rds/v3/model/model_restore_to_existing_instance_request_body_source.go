package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type RestoreToExistingInstanceRequestBodySource struct {
	InstanceId string `json:"instance_id"`

	Type *RestoreToExistingInstanceRequestBodySourceType `json:"type,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	RestoreTime *int32 `json:"restore_time,omitempty"`

	DatabaseName map[string]string `json:"database_name,omitempty"`
}

func (o RestoreToExistingInstanceRequestBodySource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceRequestBodySource struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceRequestBodySource", string(data)}, " ")
}

type RestoreToExistingInstanceRequestBodySourceType struct {
	value string
}

type RestoreToExistingInstanceRequestBodySourceTypeEnum struct {
	BACKUP    RestoreToExistingInstanceRequestBodySourceType
	TIMESTAMP RestoreToExistingInstanceRequestBodySourceType
}

func GetRestoreToExistingInstanceRequestBodySourceTypeEnum() RestoreToExistingInstanceRequestBodySourceTypeEnum {
	return RestoreToExistingInstanceRequestBodySourceTypeEnum{
		BACKUP: RestoreToExistingInstanceRequestBodySourceType{
			value: "backup",
		},
		TIMESTAMP: RestoreToExistingInstanceRequestBodySourceType{
			value: "timestamp",
		},
	}
}

func (c RestoreToExistingInstanceRequestBodySourceType) Value() string {
	return c.value
}

func (c RestoreToExistingInstanceRequestBodySourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestoreToExistingInstanceRequestBodySourceType) UnmarshalJSON(b []byte) error {
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
