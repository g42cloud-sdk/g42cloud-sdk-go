package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTrackerRequestBody struct {
	TrackerType CreateTrackerRequestBodyTrackerType `json:"tracker_type"`

	TrackerName string `json:"tracker_name"`

	IsLtsEnabled *bool `json:"is_lts_enabled,omitempty"`

	ObsInfo *TrackerObsInfo `json:"obs_info,omitempty"`

	IsSupportTraceFilesEncryption *bool `json:"is_support_trace_files_encryption,omitempty"`

	KmsId *string `json:"kms_id,omitempty"`

	IsSupportValidate *bool `json:"is_support_validate,omitempty"`

	DataBucket *DataBucket `json:"data_bucket,omitempty"`
}

func (o CreateTrackerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTrackerRequestBody", string(data)}, " ")
}

type CreateTrackerRequestBodyTrackerType struct {
	value string
}

type CreateTrackerRequestBodyTrackerTypeEnum struct {
	SYSTEM CreateTrackerRequestBodyTrackerType
	DATA   CreateTrackerRequestBodyTrackerType
}

func GetCreateTrackerRequestBodyTrackerTypeEnum() CreateTrackerRequestBodyTrackerTypeEnum {
	return CreateTrackerRequestBodyTrackerTypeEnum{
		SYSTEM: CreateTrackerRequestBodyTrackerType{
			value: "system",
		},
		DATA: CreateTrackerRequestBodyTrackerType{
			value: "data",
		},
	}
}

func (c CreateTrackerRequestBodyTrackerType) Value() string {
	return c.value
}

func (c CreateTrackerRequestBodyTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerRequestBodyTrackerType) UnmarshalJSON(b []byte) error {
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
