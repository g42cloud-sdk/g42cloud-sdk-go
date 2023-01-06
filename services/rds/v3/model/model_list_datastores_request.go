package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListDatastoresRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	DatabaseName ListDatastoresRequestDatabaseName `json:"database_name"`
}

func (o ListDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoresRequest", string(data)}, " ")
}

type ListDatastoresRequestDatabaseName struct {
	value string
}

type ListDatastoresRequestDatabaseNameEnum struct {
	MY_SQL      ListDatastoresRequestDatabaseName
	POSTGRE_SQL ListDatastoresRequestDatabaseName
	SQL_SERVER  ListDatastoresRequestDatabaseName
}

func GetListDatastoresRequestDatabaseNameEnum() ListDatastoresRequestDatabaseNameEnum {
	return ListDatastoresRequestDatabaseNameEnum{
		MY_SQL: ListDatastoresRequestDatabaseName{
			value: "MySQL",
		},
		POSTGRE_SQL: ListDatastoresRequestDatabaseName{
			value: "PostgreSQL",
		},
		SQL_SERVER: ListDatastoresRequestDatabaseName{
			value: "SQLServer",
		},
	}
}

func (c ListDatastoresRequestDatabaseName) Value() string {
	return c.value
}

func (c ListDatastoresRequestDatabaseName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatastoresRequestDatabaseName) UnmarshalJSON(b []byte) error {
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
