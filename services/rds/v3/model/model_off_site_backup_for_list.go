package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type OffSiteBackupForList struct {
	Id string `json:"id"`

	InstanceId string `json:"instance_id"`

	Name string `json:"name"`

	Databases *[]BackupDatabase `json:"databases,omitempty"`

	BeginTime string `json:"begin_time"`

	EndTime string `json:"end_time"`

	Status OffSiteBackupForListStatus `json:"status"`

	Type OffSiteBackupForListType `json:"type"`

	Size int64 `json:"size"`

	Datastore *ParaGroupDatastore `json:"datastore"`

	AssociatedWithDdm *bool `json:"associated_with_ddm,omitempty"`
}

func (o OffSiteBackupForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OffSiteBackupForList struct{}"
	}

	return strings.Join([]string{"OffSiteBackupForList", string(data)}, " ")
}

type OffSiteBackupForListStatus struct {
	value string
}

type OffSiteBackupForListStatusEnum struct {
	BUILDING  OffSiteBackupForListStatus
	COMPLETED OffSiteBackupForListStatus
	FAILED    OffSiteBackupForListStatus
	DELETING  OffSiteBackupForListStatus
}

func GetOffSiteBackupForListStatusEnum() OffSiteBackupForListStatusEnum {
	return OffSiteBackupForListStatusEnum{
		BUILDING: OffSiteBackupForListStatus{
			value: "BUILDING",
		},
		COMPLETED: OffSiteBackupForListStatus{
			value: "COMPLETED",
		},
		FAILED: OffSiteBackupForListStatus{
			value: "FAILED",
		},
		DELETING: OffSiteBackupForListStatus{
			value: "DELETING",
		},
	}
}

func (c OffSiteBackupForListStatus) Value() string {
	return c.value
}

func (c OffSiteBackupForListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OffSiteBackupForListStatus) UnmarshalJSON(b []byte) error {
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

type OffSiteBackupForListType struct {
	value string
}

type OffSiteBackupForListTypeEnum struct {
	AUTO        OffSiteBackupForListType
	INCREMENTAL OffSiteBackupForListType
}

func GetOffSiteBackupForListTypeEnum() OffSiteBackupForListTypeEnum {
	return OffSiteBackupForListTypeEnum{
		AUTO: OffSiteBackupForListType{
			value: "auto",
		},
		INCREMENTAL: OffSiteBackupForListType{
			value: "incremental",
		},
	}
}

func (c OffSiteBackupForListType) Value() string {
	return c.value
}

func (c OffSiteBackupForListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OffSiteBackupForListType) UnmarshalJSON(b []byte) error {
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
