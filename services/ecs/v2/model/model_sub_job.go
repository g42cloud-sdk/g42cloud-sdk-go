package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type SubJob struct {
	Status *SubJobStatus `json:"status,omitempty"`

	Entities *SubJobEntities `json:"entities,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`
}

func (o SubJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJob struct{}"
	}

	return strings.Join([]string{"SubJob", string(data)}, " ")
}

type SubJobStatus struct {
	value string
}

type SubJobStatusEnum struct {
	SUCCESS SubJobStatus
	RUNNING SubJobStatus
	FAIL    SubJobStatus
	INIT    SubJobStatus
}

func GetSubJobStatusEnum() SubJobStatusEnum {
	return SubJobStatusEnum{
		SUCCESS: SubJobStatus{
			value: "SUCCESS",
		},
		RUNNING: SubJobStatus{
			value: "RUNNING",
		},
		FAIL: SubJobStatus{
			value: "FAIL",
		},
		INIT: SubJobStatus{
			value: "INIT",
		},
	}
}

func (c SubJobStatus) Value() string {
	return c.value
}

func (c SubJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobStatus) UnmarshalJSON(b []byte) error {
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
