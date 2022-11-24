package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ParaGroupDatastore struct {
	Type ParaGroupDatastoreType `json:"type"`

	Version string `json:"version"`
}

func (o ParaGroupDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParaGroupDatastore struct{}"
	}

	return strings.Join([]string{"ParaGroupDatastore", string(data)}, " ")
}

type ParaGroupDatastoreType struct {
	value string
}

type ParaGroupDatastoreTypeEnum struct {
	MY_SQL      ParaGroupDatastoreType
	POSTGRE_SQL ParaGroupDatastoreType
	SQL_SERVER  ParaGroupDatastoreType
}

func GetParaGroupDatastoreTypeEnum() ParaGroupDatastoreTypeEnum {
	return ParaGroupDatastoreTypeEnum{
		MY_SQL: ParaGroupDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: ParaGroupDatastoreType{
			value: "PostgreSQL",
		},
		SQL_SERVER: ParaGroupDatastoreType{
			value: "SQLServer",
		},
	}
}

func (c ParaGroupDatastoreType) Value() string {
	return c.value
}

func (c ParaGroupDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParaGroupDatastoreType) UnmarshalJSON(b []byte) error {
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
