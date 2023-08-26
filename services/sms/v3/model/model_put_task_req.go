package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PutTaskReq struct {
	Name *string `json:"name,omitempty"`

	Type *PutTaskReqType `json:"type,omitempty"`

	OsType *PutTaskReqOsType `json:"os_type,omitempty"`

	Id *string `json:"id,omitempty"`

	Priority *PutTaskReqPriority `json:"priority,omitempty"`

	RegionId *string `json:"region_id,omitempty"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	MigrationIp *string `json:"migration_ip,omitempty"`

	RegionName *string `json:"region_name,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	ProjectName *string `json:"project_name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	VmTemplateId *string `json:"vm_template_id,omitempty"`

	SourceServer *PostSourceServerBody `json:"source_server,omitempty"`

	TargetServer *TargetServer `json:"target_server,omitempty"`

	State *string `json:"state,omitempty"`

	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	Connected *bool `json:"connected,omitempty"`

	CreateDate *int64 `json:"create_date,omitempty"`

	StartDate *int64 `json:"start_date,omitempty"`

	FinishDate *int64 `json:"finish_date,omitempty"`

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	ErrorJson *string `json:"error_json,omitempty"`

	TotalTime *int64 `json:"total_time,omitempty"`

	FloatIp *string `json:"float_ip,omitempty"`

	RemainSeconds *int64 `json:"remain_seconds,omitempty"`

	TargetSnapshotId *string `json:"target_snapshot_id,omitempty"`

	CloneServer *CloneServer `json:"clone_server,omitempty"`

	SubTasks *[]SubTask `json:"sub_tasks,omitempty"`
}

func (o PutTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutTaskReq struct{}"
	}

	return strings.Join([]string{"PutTaskReq", string(data)}, " ")
}

type PutTaskReqType struct {
	value string
}

type PutTaskReqTypeEnum struct {
	MIGRATE_FILE  PutTaskReqType
	MIGRATE_BLOCK PutTaskReqType
}

func GetPutTaskReqTypeEnum() PutTaskReqTypeEnum {
	return PutTaskReqTypeEnum{
		MIGRATE_FILE: PutTaskReqType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: PutTaskReqType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c PutTaskReqType) Value() string {
	return c.value
}

func (c PutTaskReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutTaskReqType) UnmarshalJSON(b []byte) error {
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

type PutTaskReqOsType struct {
	value string
}

type PutTaskReqOsTypeEnum struct {
	WINDOWS PutTaskReqOsType
	LINUX   PutTaskReqOsType
}

func GetPutTaskReqOsTypeEnum() PutTaskReqOsTypeEnum {
	return PutTaskReqOsTypeEnum{
		WINDOWS: PutTaskReqOsType{
			value: "WINDOWS",
		},
		LINUX: PutTaskReqOsType{
			value: "LINUX",
		},
	}
}

func (c PutTaskReqOsType) Value() string {
	return c.value
}

func (c PutTaskReqOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutTaskReqOsType) UnmarshalJSON(b []byte) error {
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

type PutTaskReqPriority struct {
	value int32
}

type PutTaskReqPriorityEnum struct {
	E_0 PutTaskReqPriority
	E_1 PutTaskReqPriority
	E_2 PutTaskReqPriority
}

func GetPutTaskReqPriorityEnum() PutTaskReqPriorityEnum {
	return PutTaskReqPriorityEnum{
		E_0: PutTaskReqPriority{
			value: 0,
		}, E_1: PutTaskReqPriority{
			value: 1,
		}, E_2: PutTaskReqPriority{
			value: 2,
		},
	}
}

func (c PutTaskReqPriority) Value() int32 {
	return c.value
}

func (c PutTaskReqPriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutTaskReqPriority) UnmarshalJSON(b []byte) error {
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
