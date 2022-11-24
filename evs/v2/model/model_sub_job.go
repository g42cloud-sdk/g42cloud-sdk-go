package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubJob struct {
	Status SubJobStatus `json:"status"`

	Entities *SubJobEntities `json:"entities"`

	JobId string `json:"job_id"`

	JobType string `json:"job_type"`

	BeginTime string `json:"begin_time"`

	EndTime string `json:"end_time"`

	ErrorCode string `json:"error_code"`

	FailReason string `json:"fail_reason"`
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
