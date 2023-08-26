package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PostMigProjectBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Isdefault *bool `json:"isdefault,omitempty"`

	Region string `json:"region"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	UsePublicIp bool `json:"use_public_ip"`

	ExistServer bool `json:"exist_server"`

	Type PostMigProjectBodyType `json:"type"`

	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	Syncing bool `json:"syncing"`

	StartNetworckCheck *bool `json:"start_networck_check,omitempty"`
}

func (o PostMigProjectBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostMigProjectBody struct{}"
	}

	return strings.Join([]string{"PostMigProjectBody", string(data)}, " ")
}

type PostMigProjectBodyType struct {
	value string
}

type PostMigProjectBodyTypeEnum struct {
	MIGRATE_BLOCK PostMigProjectBodyType
	MIGRATE_FILE  PostMigProjectBodyType
}

func GetPostMigProjectBodyTypeEnum() PostMigProjectBodyTypeEnum {
	return PostMigProjectBodyTypeEnum{
		MIGRATE_BLOCK: PostMigProjectBodyType{
			value: "MIGRATE_BLOCK",
		},
		MIGRATE_FILE: PostMigProjectBodyType{
			value: "MIGRATE_FILE",
		},
	}
}

func (c PostMigProjectBodyType) Value() string {
	return c.value
}

func (c PostMigProjectBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostMigProjectBodyType) UnmarshalJSON(b []byte) error {
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
