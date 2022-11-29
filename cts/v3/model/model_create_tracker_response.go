package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTrackerResponse struct {
	Id *string `json:"id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	KmsId *string `json:"kms_id,omitempty"`

	IsSupportValidate *bool `json:"is_support_validate,omitempty"`

	Lts *Lts `json:"lts,omitempty"`

	TrackerType *CreateTrackerResponseTrackerType `json:"tracker_type,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	TrackerName *string `json:"tracker_name,omitempty"`

	Status *CreateTrackerResponseStatus `json:"status,omitempty"`

	Detail *string `json:"detail,omitempty"`

	IsSupportTraceFilesEncryption *bool `json:"is_support_trace_files_encryption,omitempty"`

	ObsInfo *ObsInfo `json:"obs_info,omitempty"`

	DataBucket     *DataBucketQuery `json:"data_bucket,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerResponse struct{}"
	}

	return strings.Join([]string{"CreateTrackerResponse", string(data)}, " ")
}

type CreateTrackerResponseTrackerType struct {
	value string
}

type CreateTrackerResponseTrackerTypeEnum struct {
	SYSTEM CreateTrackerResponseTrackerType
	DATA   CreateTrackerResponseTrackerType
}

func GetCreateTrackerResponseTrackerTypeEnum() CreateTrackerResponseTrackerTypeEnum {
	return CreateTrackerResponseTrackerTypeEnum{
		SYSTEM: CreateTrackerResponseTrackerType{
			value: "system",
		},
		DATA: CreateTrackerResponseTrackerType{
			value: "data",
		},
	}
}

func (c CreateTrackerResponseTrackerType) Value() string {
	return c.value
}

func (c CreateTrackerResponseTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerResponseTrackerType) UnmarshalJSON(b []byte) error {
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

type CreateTrackerResponseStatus struct {
	value string
}

type CreateTrackerResponseStatusEnum struct {
	ENABLED  CreateTrackerResponseStatus
	DISABLED CreateTrackerResponseStatus
}

func GetCreateTrackerResponseStatusEnum() CreateTrackerResponseStatusEnum {
	return CreateTrackerResponseStatusEnum{
		ENABLED: CreateTrackerResponseStatus{
			value: "enabled",
		},
		DISABLED: CreateTrackerResponseStatus{
			value: "disabled",
		},
	}
}

func (c CreateTrackerResponseStatus) Value() string {
	return c.value
}

func (c CreateTrackerResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerResponseStatus) UnmarshalJSON(b []byte) error {
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
