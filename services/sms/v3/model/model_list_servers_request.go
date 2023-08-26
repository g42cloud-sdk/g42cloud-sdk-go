package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListServersRequest struct {
	State *ListServersRequestState `json:"state,omitempty"`

	Name *string `json:"name,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Migproject *string `json:"migproject,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	MigrationCycle *ListServersRequestMigrationCycle `json:"migration_cycle,omitempty"`

	Connected *bool `json:"connected,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersRequest struct{}"
	}

	return strings.Join([]string{"ListServersRequest", string(data)}, " ")
}

type ListServersRequestState struct {
	value string
}

type ListServersRequestStateEnum struct {
	UNAVAILABLE ListServersRequestState
	WAITING     ListServersRequestState
	INITIALIZE  ListServersRequestState
	REPLICATE   ListServersRequestState
	SYNCING     ListServersRequestState
	STOPPING    ListServersRequestState
	STOPPED     ListServersRequestState
	DELETING    ListServersRequestState
	ERROR       ListServersRequestState
	CLONING     ListServersRequestState
	CUTOVERING  ListServersRequestState
	FINISHED    ListServersRequestState
}

func GetListServersRequestStateEnum() ListServersRequestStateEnum {
	return ListServersRequestStateEnum{
		UNAVAILABLE: ListServersRequestState{
			value: "unavailable",
		},
		WAITING: ListServersRequestState{
			value: "waiting",
		},
		INITIALIZE: ListServersRequestState{
			value: "initialize",
		},
		REPLICATE: ListServersRequestState{
			value: "replicate",
		},
		SYNCING: ListServersRequestState{
			value: "syncing",
		},
		STOPPING: ListServersRequestState{
			value: "stopping",
		},
		STOPPED: ListServersRequestState{
			value: "stopped",
		},
		DELETING: ListServersRequestState{
			value: "deleting",
		},
		ERROR: ListServersRequestState{
			value: "error",
		},
		CLONING: ListServersRequestState{
			value: "cloning",
		},
		CUTOVERING: ListServersRequestState{
			value: "cutovering",
		},
		FINISHED: ListServersRequestState{
			value: "finished",
		},
	}
}

func (c ListServersRequestState) Value() string {
	return c.value
}

func (c ListServersRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServersRequestState) UnmarshalJSON(b []byte) error {
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

type ListServersRequestMigrationCycle struct {
	value string
}

type ListServersRequestMigrationCycleEnum struct {
	CHECKING    ListServersRequestMigrationCycle
	SETTING     ListServersRequestMigrationCycle
	REPLICATING ListServersRequestMigrationCycle
	SYNCING     ListServersRequestMigrationCycle
	CUTOVERING  ListServersRequestMigrationCycle
	CUTOVERED   ListServersRequestMigrationCycle
}

func GetListServersRequestMigrationCycleEnum() ListServersRequestMigrationCycleEnum {
	return ListServersRequestMigrationCycleEnum{
		CHECKING: ListServersRequestMigrationCycle{
			value: "checking",
		},
		SETTING: ListServersRequestMigrationCycle{
			value: "setting",
		},
		REPLICATING: ListServersRequestMigrationCycle{
			value: "replicating",
		},
		SYNCING: ListServersRequestMigrationCycle{
			value: "syncing",
		},
		CUTOVERING: ListServersRequestMigrationCycle{
			value: "cutovering",
		},
		CUTOVERED: ListServersRequestMigrationCycle{
			value: "cutovered",
		},
	}
}

func (c ListServersRequestMigrationCycle) Value() string {
	return c.value
}

func (c ListServersRequestMigrationCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServersRequestMigrationCycle) UnmarshalJSON(b []byte) error {
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
