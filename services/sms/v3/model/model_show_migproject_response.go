package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowMigprojectResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Isdefault *bool `json:"isdefault,omitempty"`

	Template *TemplateResponseBody `json:"template,omitempty"`

	Region *string `json:"region,omitempty"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	ExistServer *bool `json:"exist_server,omitempty"`

	Type *ShowMigprojectResponseType `json:"type,omitempty"`

	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	Syncing *bool `json:"syncing,omitempty"`

	StartNetworkCheck *bool `json:"start_network_check,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigprojectResponse struct{}"
	}

	return strings.Join([]string{"ShowMigprojectResponse", string(data)}, " ")
}

type ShowMigprojectResponseType struct {
	value string
}

type ShowMigprojectResponseTypeEnum struct {
	MIGRATE_BLOCK ShowMigprojectResponseType
	MIGRATE_FILE  ShowMigprojectResponseType
}

func GetShowMigprojectResponseTypeEnum() ShowMigprojectResponseTypeEnum {
	return ShowMigprojectResponseTypeEnum{
		MIGRATE_BLOCK: ShowMigprojectResponseType{
			value: "MIGRATE_BLOCK",
		},
		MIGRATE_FILE: ShowMigprojectResponseType{
			value: "MIGRATE_FILE",
		},
	}
}

func (c ShowMigprojectResponseType) Value() string {
	return c.value
}

func (c ShowMigprojectResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigprojectResponseType) UnmarshalJSON(b []byte) error {
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
