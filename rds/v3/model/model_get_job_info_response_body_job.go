package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GetJobInfoResponseBodyJob struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Status GetJobInfoResponseBodyJobStatus `json:"status"`

	Created string `json:"created"`

	Ended *string `json:"ended,omitempty"`

	Process *string `json:"process,omitempty"`

	Instance *GetTaskDetailListRspJobsInstance `json:"instance"`

	Entities *interface{} `json:"entities,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`
}

func (o GetJobInfoResponseBodyJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetJobInfoResponseBodyJob struct{}"
	}

	return strings.Join([]string{"GetJobInfoResponseBodyJob", string(data)}, " ")
}

type GetJobInfoResponseBodyJobStatus struct {
	value string
}

type GetJobInfoResponseBodyJobStatusEnum struct {
	RUNNING   GetJobInfoResponseBodyJobStatus
	COMPLETED GetJobInfoResponseBodyJobStatus
	FAILED    GetJobInfoResponseBodyJobStatus
}

func GetGetJobInfoResponseBodyJobStatusEnum() GetJobInfoResponseBodyJobStatusEnum {
	return GetJobInfoResponseBodyJobStatusEnum{
		RUNNING: GetJobInfoResponseBodyJobStatus{
			value: "Running",
		},
		COMPLETED: GetJobInfoResponseBodyJobStatus{
			value: "Completed",
		},
		FAILED: GetJobInfoResponseBodyJobStatus{
			value: "Failed",
		},
	}
}

func (c GetJobInfoResponseBodyJobStatus) Value() string {
	return c.value
}

func (c GetJobInfoResponseBodyJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetJobInfoResponseBodyJobStatus) UnmarshalJSON(b []byte) error {
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
