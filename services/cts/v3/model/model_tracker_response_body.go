package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TrackerResponseBody struct {
	Id *string `json:"id,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	KmsId *string `json:"kms_id,omitempty"`

	IsSupportValidate *bool `json:"is_support_validate,omitempty"`

	Lts *Lts `json:"lts,omitempty"`

	TrackerType *TrackerResponseBodyTrackerType `json:"tracker_type,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	TrackerName *string `json:"tracker_name,omitempty"`

	Status *TrackerResponseBodyStatus `json:"status,omitempty"`

	Detail *string `json:"detail,omitempty"`

	IsSupportTraceFilesEncryption *bool `json:"is_support_trace_files_encryption,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	StreamId *string `json:"stream_id,omitempty"`

	ObsInfo *ObsInfo `json:"obs_info,omitempty"`

	DataBucket *DataBucketQuery `json:"data_bucket,omitempty"`
}

func (o TrackerResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerResponseBody struct{}"
	}

	return strings.Join([]string{"TrackerResponseBody", string(data)}, " ")
}

type TrackerResponseBodyTrackerType struct {
	value string
}

type TrackerResponseBodyTrackerTypeEnum struct {
	SYSTEM TrackerResponseBodyTrackerType
	DATA   TrackerResponseBodyTrackerType
}

func GetTrackerResponseBodyTrackerTypeEnum() TrackerResponseBodyTrackerTypeEnum {
	return TrackerResponseBodyTrackerTypeEnum{
		SYSTEM: TrackerResponseBodyTrackerType{
			value: "system",
		},
		DATA: TrackerResponseBodyTrackerType{
			value: "data",
		},
	}
}

func (c TrackerResponseBodyTrackerType) Value() string {
	return c.value
}

func (c TrackerResponseBodyTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrackerResponseBodyTrackerType) UnmarshalJSON(b []byte) error {
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

type TrackerResponseBodyStatus struct {
	value string
}

type TrackerResponseBodyStatusEnum struct {
	ENABLED  TrackerResponseBodyStatus
	DISABLED TrackerResponseBodyStatus
}

func GetTrackerResponseBodyStatusEnum() TrackerResponseBodyStatusEnum {
	return TrackerResponseBodyStatusEnum{
		ENABLED: TrackerResponseBodyStatus{
			value: "enabled",
		},
		DISABLED: TrackerResponseBodyStatus{
			value: "disabled",
		},
	}
}

func (c TrackerResponseBodyStatus) Value() string {
	return c.value
}

func (c TrackerResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrackerResponseBodyStatus) UnmarshalJSON(b []byte) error {
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
