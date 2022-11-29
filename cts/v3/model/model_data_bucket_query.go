package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataBucketQuery struct {
	DataBucketName *string `json:"data_bucket_name,omitempty"`

	SearchEnabled *bool `json:"search_enabled,omitempty"`

	DataEvent *[]DataBucketQueryDataEvent `json:"data_event,omitempty"`
}

func (o DataBucketQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBucketQuery struct{}"
	}

	return strings.Join([]string{"DataBucketQuery", string(data)}, " ")
}

type DataBucketQueryDataEvent struct {
	value string
}

type DataBucketQueryDataEventEnum struct {
	WRITE DataBucketQueryDataEvent
	READ  DataBucketQueryDataEvent
}

func GetDataBucketQueryDataEventEnum() DataBucketQueryDataEventEnum {
	return DataBucketQueryDataEventEnum{
		WRITE: DataBucketQueryDataEvent{
			value: "WRITE",
		},
		READ: DataBucketQueryDataEvent{
			value: "READ",
		},
	}
}

func (c DataBucketQueryDataEvent) Value() string {
	return c.value
}

func (c DataBucketQueryDataEvent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataBucketQueryDataEvent) UnmarshalJSON(b []byte) error {
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
