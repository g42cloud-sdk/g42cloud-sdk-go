package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PutCopyStateReq struct {
	Copystate *PutCopyStateReqCopystate `json:"copystate,omitempty"`

	Migrationcycle *PutCopyStateReqMigrationcycle `json:"migrationcycle,omitempty"`
}

func (o PutCopyStateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCopyStateReq struct{}"
	}

	return strings.Join([]string{"PutCopyStateReq", string(data)}, " ")
}

type PutCopyStateReqCopystate struct {
	value string
}

type PutCopyStateReqCopystateEnum struct {
	UNAVAILABLE PutCopyStateReqCopystate
	WAITING     PutCopyStateReqCopystate
	INIT        PutCopyStateReqCopystate
	REPLICATE   PutCopyStateReqCopystate
	SYNCING     PutCopyStateReqCopystate
	STOPPING    PutCopyStateReqCopystate
	STOPPED     PutCopyStateReqCopystate
	DELETING    PutCopyStateReqCopystate
	ERROR       PutCopyStateReqCopystate
	CLONING     PutCopyStateReqCopystate
	CUTOVERING  PutCopyStateReqCopystate
}

func GetPutCopyStateReqCopystateEnum() PutCopyStateReqCopystateEnum {
	return PutCopyStateReqCopystateEnum{
		UNAVAILABLE: PutCopyStateReqCopystate{
			value: "UNAVAILABLE",
		},
		WAITING: PutCopyStateReqCopystate{
			value: "WAITING",
		},
		INIT: PutCopyStateReqCopystate{
			value: "INIT",
		},
		REPLICATE: PutCopyStateReqCopystate{
			value: "REPLICATE",
		},
		SYNCING: PutCopyStateReqCopystate{
			value: "SYNCING",
		},
		STOPPING: PutCopyStateReqCopystate{
			value: "STOPPING",
		},
		STOPPED: PutCopyStateReqCopystate{
			value: "STOPPED",
		},
		DELETING: PutCopyStateReqCopystate{
			value: "DELETING",
		},
		ERROR: PutCopyStateReqCopystate{
			value: "ERROR",
		},
		CLONING: PutCopyStateReqCopystate{
			value: "CLONING",
		},
		CUTOVERING: PutCopyStateReqCopystate{
			value: "CUTOVERING",
		},
	}
}

func (c PutCopyStateReqCopystate) Value() string {
	return c.value
}

func (c PutCopyStateReqCopystate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutCopyStateReqCopystate) UnmarshalJSON(b []byte) error {
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

type PutCopyStateReqMigrationcycle struct {
	value string
}

type PutCopyStateReqMigrationcycleEnum struct {
	CUTOVERING  PutCopyStateReqMigrationcycle
	CUTOVERED   PutCopyStateReqMigrationcycle
	CHECKING    PutCopyStateReqMigrationcycle
	SETTING     PutCopyStateReqMigrationcycle
	REPLICATING PutCopyStateReqMigrationcycle
	SYNCING     PutCopyStateReqMigrationcycle
}

func GetPutCopyStateReqMigrationcycleEnum() PutCopyStateReqMigrationcycleEnum {
	return PutCopyStateReqMigrationcycleEnum{
		CUTOVERING: PutCopyStateReqMigrationcycle{
			value: "cutovering",
		},
		CUTOVERED: PutCopyStateReqMigrationcycle{
			value: "cutovered",
		},
		CHECKING: PutCopyStateReqMigrationcycle{
			value: "checking",
		},
		SETTING: PutCopyStateReqMigrationcycle{
			value: "setting",
		},
		REPLICATING: PutCopyStateReqMigrationcycle{
			value: "replicating",
		},
		SYNCING: PutCopyStateReqMigrationcycle{
			value: "syncing",
		},
	}
}

func (c PutCopyStateReqMigrationcycle) Value() string {
	return c.value
}

func (c PutCopyStateReqMigrationcycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutCopyStateReqMigrationcycle) UnmarshalJSON(b []byte) error {
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
