package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowEngineResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	AuthType *ShowEngineResponseAuthType `json:"auth_type,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	Payment *string `json:"payment,omitempty"`

	Version *string `json:"version,omitempty"`

	LatestVersion *string `json:"latest_version,omitempty"`

	Status *ShowEngineResponseStatus `json:"status,omitempty"`

	BeDefault *bool `json:"be_default,omitempty"`

	CreateUser *string `json:"create_user,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	CceSpec *Spec `json:"cce_spec,omitempty"`

	ExternalEntrypoint *EngineExternalEntrypoint `json:"external_entrypoint,omitempty"`

	Reference *EngineReference `json:"reference,omitempty"`

	LatestJobId *int32 `json:"latest_job_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	EngineAdditionalActions *[]ShowEngineResponseEngineAdditionalActions `json:"engine_additional_actions,omitempty"`

	SpecType *ShowEngineResponseSpecType `json:"spec_type,omitempty"`

	Type *ShowEngineResponseType `json:"type,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	VmIds          *[]string `json:"vm_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineResponse struct{}"
	}

	return strings.Join([]string{"ShowEngineResponse", string(data)}, " ")
}

type ShowEngineResponseAuthType struct {
	value string
}

type ShowEngineResponseAuthTypeEnum struct {
	RBAC ShowEngineResponseAuthType
	NONE ShowEngineResponseAuthType
}

func GetShowEngineResponseAuthTypeEnum() ShowEngineResponseAuthTypeEnum {
	return ShowEngineResponseAuthTypeEnum{
		RBAC: ShowEngineResponseAuthType{
			value: "RBAC",
		},
		NONE: ShowEngineResponseAuthType{
			value: "NONE",
		},
	}
}

func (c ShowEngineResponseAuthType) Value() string {
	return c.value
}

func (c ShowEngineResponseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseAuthType) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseStatus struct {
	value string
}

type ShowEngineResponseStatusEnum struct {
	CREATING       ShowEngineResponseStatus
	AVAILABLE      ShowEngineResponseStatus
	UNAVAILABLE    ShowEngineResponseStatus
	DELETING       ShowEngineResponseStatus
	DELETED        ShowEngineResponseStatus
	UPGRADING      ShowEngineResponseStatus
	MODIFYING      ShowEngineResponseStatus
	CREATE_FAILED  ShowEngineResponseStatus
	DELETE_FAILED  ShowEngineResponseStatus
	UPGRADE_FAILED ShowEngineResponseStatus
	MODIFY_FAILED  ShowEngineResponseStatus
	FREEZED        ShowEngineResponseStatus
}

func GetShowEngineResponseStatusEnum() ShowEngineResponseStatusEnum {
	return ShowEngineResponseStatusEnum{
		CREATING: ShowEngineResponseStatus{
			value: "Creating",
		},
		AVAILABLE: ShowEngineResponseStatus{
			value: "Available",
		},
		UNAVAILABLE: ShowEngineResponseStatus{
			value: "Unavailable",
		},
		DELETING: ShowEngineResponseStatus{
			value: "Deleting",
		},
		DELETED: ShowEngineResponseStatus{
			value: "Deleted",
		},
		UPGRADING: ShowEngineResponseStatus{
			value: "Upgrading",
		},
		MODIFYING: ShowEngineResponseStatus{
			value: "Modifying",
		},
		CREATE_FAILED: ShowEngineResponseStatus{
			value: "CreateFailed",
		},
		DELETE_FAILED: ShowEngineResponseStatus{
			value: "DeleteFailed",
		},
		UPGRADE_FAILED: ShowEngineResponseStatus{
			value: "UpgradeFailed",
		},
		MODIFY_FAILED: ShowEngineResponseStatus{
			value: "ModifyFailed",
		},
		FREEZED: ShowEngineResponseStatus{
			value: "Freezed",
		},
	}
}

func (c ShowEngineResponseStatus) Value() string {
	return c.value
}

func (c ShowEngineResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseEngineAdditionalActions struct {
	value string
}

type ShowEngineResponseEngineAdditionalActionsEnum struct {
	NOTING       ShowEngineResponseEngineAdditionalActions
	FORCE_DELETE ShowEngineResponseEngineAdditionalActions
	ROLLBACK     ShowEngineResponseEngineAdditionalActions
	RETRY        ShowEngineResponseEngineAdditionalActions
}

func GetShowEngineResponseEngineAdditionalActionsEnum() ShowEngineResponseEngineAdditionalActionsEnum {
	return ShowEngineResponseEngineAdditionalActionsEnum{
		NOTING: ShowEngineResponseEngineAdditionalActions{
			value: "Noting",
		},
		FORCE_DELETE: ShowEngineResponseEngineAdditionalActions{
			value: "ForceDelete",
		},
		ROLLBACK: ShowEngineResponseEngineAdditionalActions{
			value: "Rollback",
		},
		RETRY: ShowEngineResponseEngineAdditionalActions{
			value: "Retry",
		},
	}
}

func (c ShowEngineResponseEngineAdditionalActions) Value() string {
	return c.value
}

func (c ShowEngineResponseEngineAdditionalActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseEngineAdditionalActions) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseSpecType struct {
	value string
}

type ShowEngineResponseSpecTypeEnum struct {
	CCE          ShowEngineResponseSpecType
	CSE          ShowEngineResponseSpecType
	SPRING_CLOUD ShowEngineResponseSpecType
}

func GetShowEngineResponseSpecTypeEnum() ShowEngineResponseSpecTypeEnum {
	return ShowEngineResponseSpecTypeEnum{
		CCE: ShowEngineResponseSpecType{
			value: "CCE",
		},
		CSE: ShowEngineResponseSpecType{
			value: "CSE",
		},
		SPRING_CLOUD: ShowEngineResponseSpecType{
			value: "SpringCloud",
		},
	}
}

func (c ShowEngineResponseSpecType) Value() string {
	return c.value
}

func (c ShowEngineResponseSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseSpecType) UnmarshalJSON(b []byte) error {
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

type ShowEngineResponseType struct {
	value string
}

type ShowEngineResponseTypeEnum struct {
	CSE       ShowEngineResponseType
	CSE_SHARE ShowEngineResponseType
}

func GetShowEngineResponseTypeEnum() ShowEngineResponseTypeEnum {
	return ShowEngineResponseTypeEnum{
		CSE: ShowEngineResponseType{
			value: "CSE",
		},
		CSE_SHARE: ShowEngineResponseType{
			value: "CSE_Share",
		},
	}
}

func (c ShowEngineResponseType) Value() string {
	return c.value
}

func (c ShowEngineResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineResponseType) UnmarshalJSON(b []byte) error {
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
