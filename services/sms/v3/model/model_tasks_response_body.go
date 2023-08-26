package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TasksResponseBody struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *TasksResponseBodyType `json:"type,omitempty"`

	OsType *TasksResponseBodyOsType `json:"os_type,omitempty"`

	State *string `json:"state,omitempty"`

	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	CreateDate *int64 `json:"create_date,omitempty"`

	Priority *TasksResponseBodyPriority `json:"priority,omitempty"`

	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	CompressRate *float64 `json:"compress_rate,omitempty"`

	StartTargetServer *bool `json:"start_target_server,omitempty"`

	ErrorJson *string `json:"error_json,omitempty"`

	TotalTime *int64 `json:"total_time,omitempty"`

	MigrationIp *string `json:"migration_ip,omitempty"`

	SubTasks *[]SubTaskAssociatedWithTask `json:"sub_tasks,omitempty"`

	SourceServer *SourceServerAssociatedWithTask `json:"source_server,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	TargetServer *TargetServerAssociatedWithTask `json:"target_server,omitempty"`

	LogCollectStatus *TasksResponseBodyLogCollectStatus `json:"log_collect_status,omitempty"`

	CloneServer *CloneServerBrief `json:"clone_server,omitempty"`

	Syncing *bool `json:"syncing,omitempty"`

	NetworkCheckInfo *NetworkCheckInfoRequestBody `json:"network_check_info,omitempty"`

	SpecialConfig *[]ConfigBody `json:"special_config,omitempty"`
}

func (o TasksResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TasksResponseBody struct{}"
	}

	return strings.Join([]string{"TasksResponseBody", string(data)}, " ")
}

type TasksResponseBodyType struct {
	value string
}

type TasksResponseBodyTypeEnum struct {
	MIGRATE_FILE  TasksResponseBodyType
	MIGRATE_BLOCK TasksResponseBodyType
}

func GetTasksResponseBodyTypeEnum() TasksResponseBodyTypeEnum {
	return TasksResponseBodyTypeEnum{
		MIGRATE_FILE: TasksResponseBodyType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: TasksResponseBodyType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c TasksResponseBodyType) Value() string {
	return c.value
}

func (c TasksResponseBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyType) UnmarshalJSON(b []byte) error {
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

type TasksResponseBodyOsType struct {
	value string
}

type TasksResponseBodyOsTypeEnum struct {
	WINDOWS TasksResponseBodyOsType
	LINUX   TasksResponseBodyOsType
}

func GetTasksResponseBodyOsTypeEnum() TasksResponseBodyOsTypeEnum {
	return TasksResponseBodyOsTypeEnum{
		WINDOWS: TasksResponseBodyOsType{
			value: "WINDOWS",
		},
		LINUX: TasksResponseBodyOsType{
			value: "LINUX",
		},
	}
}

func (c TasksResponseBodyOsType) Value() string {
	return c.value
}

func (c TasksResponseBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyOsType) UnmarshalJSON(b []byte) error {
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

type TasksResponseBodyPriority struct {
	value int32
}

type TasksResponseBodyPriorityEnum struct {
	E_0 TasksResponseBodyPriority
	E_1 TasksResponseBodyPriority
	E_2 TasksResponseBodyPriority
}

func GetTasksResponseBodyPriorityEnum() TasksResponseBodyPriorityEnum {
	return TasksResponseBodyPriorityEnum{
		E_0: TasksResponseBodyPriority{
			value: 0,
		}, E_1: TasksResponseBodyPriority{
			value: 1,
		}, E_2: TasksResponseBodyPriority{
			value: 2,
		},
	}
}

func (c TasksResponseBodyPriority) Value() int32 {
	return c.value
}

func (c TasksResponseBodyPriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyPriority) UnmarshalJSON(b []byte) error {
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

type TasksResponseBodyLogCollectStatus struct {
	value string
}

type TasksResponseBodyLogCollectStatusEnum struct {
	INIT                     TasksResponseBodyLogCollectStatus
	TELL_AGENT_TO_COLLECT    TasksResponseBodyLogCollectStatus
	WAIT_AGENT_COLLECT_ACK   TasksResponseBodyLogCollectStatus
	AGENT_COLLECT_FAIL       TasksResponseBodyLogCollectStatus
	AGENT_COLLECT_SUCCESS    TasksResponseBodyLogCollectStatus
	WAIT_SERVER_COLLECT      TasksResponseBodyLogCollectStatus
	SERVER_COLLECT_FAIL      TasksResponseBodyLogCollectStatus
	SERVER_COLLECT_SUCCESS   TasksResponseBodyLogCollectStatus
	TELL_AGENT_RESET_ACL     TasksResponseBodyLogCollectStatus
	WAIT_AGENT_RESET_ACL_ACK TasksResponseBodyLogCollectStatus
}

func GetTasksResponseBodyLogCollectStatusEnum() TasksResponseBodyLogCollectStatusEnum {
	return TasksResponseBodyLogCollectStatusEnum{
		INIT: TasksResponseBodyLogCollectStatus{
			value: "INIT",
		},
		TELL_AGENT_TO_COLLECT: TasksResponseBodyLogCollectStatus{
			value: "TELL_AGENT_TO_COLLECT",
		},
		WAIT_AGENT_COLLECT_ACK: TasksResponseBodyLogCollectStatus{
			value: "WAIT_AGENT_COLLECT_ACK",
		},
		AGENT_COLLECT_FAIL: TasksResponseBodyLogCollectStatus{
			value: "AGENT_COLLECT_FAIL",
		},
		AGENT_COLLECT_SUCCESS: TasksResponseBodyLogCollectStatus{
			value: "AGENT_COLLECT_SUCCESS",
		},
		WAIT_SERVER_COLLECT: TasksResponseBodyLogCollectStatus{
			value: "WAIT_SERVER_COLLECT",
		},
		SERVER_COLLECT_FAIL: TasksResponseBodyLogCollectStatus{
			value: "SERVER_COLLECT_FAIL",
		},
		SERVER_COLLECT_SUCCESS: TasksResponseBodyLogCollectStatus{
			value: "SERVER_COLLECT_SUCCESS",
		},
		TELL_AGENT_RESET_ACL: TasksResponseBodyLogCollectStatus{
			value: "TELL_AGENT_RESET_ACL",
		},
		WAIT_AGENT_RESET_ACL_ACK: TasksResponseBodyLogCollectStatus{
			value: "WAIT_AGENT_RESET_ACL_ACK",
		},
	}
}

func (c TasksResponseBodyLogCollectStatus) Value() string {
	return c.value
}

func (c TasksResponseBodyLogCollectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TasksResponseBodyLogCollectStatus) UnmarshalJSON(b []byte) error {
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
