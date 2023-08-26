package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowTaskResponse struct {
	Name *string `json:"name,omitempty"`

	Type *ShowTaskResponseType `json:"type,omitempty"`

	OsType *ShowTaskResponseOsType `json:"os_type,omitempty"`

	Id *string `json:"id,omitempty"`

	Priority *ShowTaskResponsePriority `json:"priority,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	RegionId *string `json:"region_id,omitempty"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	MigrationIp *string `json:"migration_ip,omitempty"`

	RegionName *string `json:"region_name,omitempty"`

	ProjectName *string `json:"project_name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	VmTemplateId *string `json:"vm_template_id,omitempty"`

	SourceServer *SourceServerResponse `json:"source_server,omitempty"`

	TargetServer *TaskTargetServer `json:"target_server,omitempty"`

	State *string `json:"state,omitempty"`

	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	Connected *bool `json:"connected,omitempty"`

	CreateDate *int64 `json:"create_date,omitempty"`

	StartDate *int64 `json:"start_date,omitempty"`

	FinishDate *int64 `json:"finish_date,omitempty"`

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	CompressRate *float64 `json:"compress_rate,omitempty"`

	ErrorJson *string `json:"error_json,omitempty"`

	TotalTime *int64 `json:"total_time,omitempty"`

	FloatIp *string `json:"float_ip,omitempty"`

	RemainSeconds *int64 `json:"remain_seconds,omitempty"`

	TargetSnapshotId *string `json:"target_snapshot_id,omitempty"`

	CloneServer *CloneServer `json:"clone_server,omitempty"`

	SubTasks *[]SubTask `json:"sub_tasks,omitempty"`

	NetworkCheckInfo *NetworkCheckInfoRequestBody `json:"network_check_info,omitempty"`
	HttpStatusCode   int                          `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}

type ShowTaskResponseType struct {
	value string
}

type ShowTaskResponseTypeEnum struct {
	MIGRATE_FILE  ShowTaskResponseType
	MIGRATE_BLOCK ShowTaskResponseType
}

func GetShowTaskResponseTypeEnum() ShowTaskResponseTypeEnum {
	return ShowTaskResponseTypeEnum{
		MIGRATE_FILE: ShowTaskResponseType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: ShowTaskResponseType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c ShowTaskResponseType) Value() string {
	return c.value
}

func (c ShowTaskResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseType) UnmarshalJSON(b []byte) error {
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

type ShowTaskResponseOsType struct {
	value string
}

type ShowTaskResponseOsTypeEnum struct {
	WINDOWS ShowTaskResponseOsType
	LINUX   ShowTaskResponseOsType
}

func GetShowTaskResponseOsTypeEnum() ShowTaskResponseOsTypeEnum {
	return ShowTaskResponseOsTypeEnum{
		WINDOWS: ShowTaskResponseOsType{
			value: "WINDOWS",
		},
		LINUX: ShowTaskResponseOsType{
			value: "LINUX",
		},
	}
}

func (c ShowTaskResponseOsType) Value() string {
	return c.value
}

func (c ShowTaskResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseOsType) UnmarshalJSON(b []byte) error {
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

type ShowTaskResponsePriority struct {
	value int32
}

type ShowTaskResponsePriorityEnum struct {
	E_0 ShowTaskResponsePriority
	E_1 ShowTaskResponsePriority
	E_2 ShowTaskResponsePriority
}

func GetShowTaskResponsePriorityEnum() ShowTaskResponsePriorityEnum {
	return ShowTaskResponsePriorityEnum{
		E_0: ShowTaskResponsePriority{
			value: 0,
		}, E_1: ShowTaskResponsePriority{
			value: 1,
		}, E_2: ShowTaskResponsePriority{
			value: 2,
		},
	}
}

func (c ShowTaskResponsePriority) Value() int32 {
	return c.value
}

func (c ShowTaskResponsePriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponsePriority) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
