package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type OpExtendInfoBckup struct {
	AppConsistencyErrorCode *string `json:"app_consistency_error_code,omitempty"`

	AppConsistencyErrorMessage *string `json:"app_consistency_error_message,omitempty"`

	AppConsistencyStatus *OpExtendInfoBckupAppConsistencyStatus `json:"app_consistency_status,omitempty"`

	BackupId string `json:"backup_id"`

	BackupName *string `json:"backup_name,omitempty"`

	Incremental *OpExtendInfoBckupIncremental `json:"incremental,omitempty"`
}

func (o OpExtendInfoBckup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoBckup struct{}"
	}

	return strings.Join([]string{"OpExtendInfoBckup", string(data)}, " ")
}

type OpExtendInfoBckupAppConsistencyStatus struct {
	value string
}

type OpExtendInfoBckupAppConsistencyStatusEnum struct {
	E_0 OpExtendInfoBckupAppConsistencyStatus
	E_1 OpExtendInfoBckupAppConsistencyStatus
}

func GetOpExtendInfoBckupAppConsistencyStatusEnum() OpExtendInfoBckupAppConsistencyStatusEnum {
	return OpExtendInfoBckupAppConsistencyStatusEnum{
		E_0: OpExtendInfoBckupAppConsistencyStatus{
			value: "0",
		},
		E_1: OpExtendInfoBckupAppConsistencyStatus{
			value: "1",
		},
	}
}

func (c OpExtendInfoBckupAppConsistencyStatus) Value() string {
	return c.value
}

func (c OpExtendInfoBckupAppConsistencyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpExtendInfoBckupAppConsistencyStatus) UnmarshalJSON(b []byte) error {
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

type OpExtendInfoBckupIncremental struct {
	value string
}

type OpExtendInfoBckupIncrementalEnum struct {
	TRUE  OpExtendInfoBckupIncremental
	FALSE OpExtendInfoBckupIncremental
}

func GetOpExtendInfoBckupIncrementalEnum() OpExtendInfoBckupIncrementalEnum {
	return OpExtendInfoBckupIncrementalEnum{
		TRUE: OpExtendInfoBckupIncremental{
			value: "\"true\"",
		},
		FALSE: OpExtendInfoBckupIncremental{
			value: "\"false\"",
		},
	}
}

func (c OpExtendInfoBckupIncremental) Value() string {
	return c.value
}

func (c OpExtendInfoBckupIncremental) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpExtendInfoBckupIncremental) UnmarshalJSON(b []byte) error {
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
