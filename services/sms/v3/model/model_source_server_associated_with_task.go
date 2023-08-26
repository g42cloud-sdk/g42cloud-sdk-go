package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type SourceServerAssociatedWithTask struct {
	Id *string `json:"id,omitempty"`

	Ip string `json:"ip"`

	Name string `json:"name"`

	OsType SourceServerAssociatedWithTaskOsType `json:"os_type"`

	OsVersion *string `json:"os_version,omitempty"`

	OemSystem *bool `json:"oem_system,omitempty"`

	State *SourceServerAssociatedWithTaskState `json:"state,omitempty"`
}

func (o SourceServerAssociatedWithTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceServerAssociatedWithTask struct{}"
	}

	return strings.Join([]string{"SourceServerAssociatedWithTask", string(data)}, " ")
}

type SourceServerAssociatedWithTaskOsType struct {
	value string
}

type SourceServerAssociatedWithTaskOsTypeEnum struct {
	WINDOWS SourceServerAssociatedWithTaskOsType
	LINUX   SourceServerAssociatedWithTaskOsType
}

func GetSourceServerAssociatedWithTaskOsTypeEnum() SourceServerAssociatedWithTaskOsTypeEnum {
	return SourceServerAssociatedWithTaskOsTypeEnum{
		WINDOWS: SourceServerAssociatedWithTaskOsType{
			value: "WINDOWS",
		},
		LINUX: SourceServerAssociatedWithTaskOsType{
			value: "LINUX",
		},
	}
}

func (c SourceServerAssociatedWithTaskOsType) Value() string {
	return c.value
}

func (c SourceServerAssociatedWithTaskOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerAssociatedWithTaskOsType) UnmarshalJSON(b []byte) error {
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

type SourceServerAssociatedWithTaskState struct {
	value string
}

type SourceServerAssociatedWithTaskStateEnum struct {
	UNAVAILABLE SourceServerAssociatedWithTaskState
	WAITING     SourceServerAssociatedWithTaskState
	INITIALIZE  SourceServerAssociatedWithTaskState
	REPLICATE   SourceServerAssociatedWithTaskState
	SYNCING     SourceServerAssociatedWithTaskState
	STOPPING    SourceServerAssociatedWithTaskState
	STOPPED     SourceServerAssociatedWithTaskState
	DELETING    SourceServerAssociatedWithTaskState
	ERROR       SourceServerAssociatedWithTaskState
	CLONING     SourceServerAssociatedWithTaskState
	TESTING     SourceServerAssociatedWithTaskState
	FINISHED    SourceServerAssociatedWithTaskState
}

func GetSourceServerAssociatedWithTaskStateEnum() SourceServerAssociatedWithTaskStateEnum {
	return SourceServerAssociatedWithTaskStateEnum{
		UNAVAILABLE: SourceServerAssociatedWithTaskState{
			value: "unavailable",
		},
		WAITING: SourceServerAssociatedWithTaskState{
			value: "waiting",
		},
		INITIALIZE: SourceServerAssociatedWithTaskState{
			value: "initialize",
		},
		REPLICATE: SourceServerAssociatedWithTaskState{
			value: "replicate",
		},
		SYNCING: SourceServerAssociatedWithTaskState{
			value: "syncing",
		},
		STOPPING: SourceServerAssociatedWithTaskState{
			value: "stopping",
		},
		STOPPED: SourceServerAssociatedWithTaskState{
			value: "stopped",
		},
		DELETING: SourceServerAssociatedWithTaskState{
			value: "deleting",
		},
		ERROR: SourceServerAssociatedWithTaskState{
			value: "error",
		},
		CLONING: SourceServerAssociatedWithTaskState{
			value: "cloning",
		},
		TESTING: SourceServerAssociatedWithTaskState{
			value: "testing",
		},
		FINISHED: SourceServerAssociatedWithTaskState{
			value: "finished",
		},
	}
}

func (c SourceServerAssociatedWithTaskState) Value() string {
	return c.value
}

func (c SourceServerAssociatedWithTaskState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceServerAssociatedWithTaskState) UnmarshalJSON(b []byte) error {
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
