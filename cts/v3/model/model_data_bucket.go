package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataBucket struct {
	DataBucketName *string `json:"data_bucket_name,omitempty"`

	DataEvent *[]DataBucketDataEvent `json:"data_event,omitempty"`
}

func (o DataBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBucket struct{}"
	}

	return strings.Join([]string{"DataBucket", string(data)}, " ")
}

type DataBucketDataEvent struct {
	value string
}

type DataBucketDataEventEnum struct {
	WRITE DataBucketDataEvent
	READ  DataBucketDataEvent
}

func GetDataBucketDataEventEnum() DataBucketDataEventEnum {
	return DataBucketDataEventEnum{
		WRITE: DataBucketDataEvent{
			value: "WRITE",
		},
		READ: DataBucketDataEvent{
			value: "READ",
		},
	}
}

func (c DataBucketDataEvent) Value() string {
	return c.value
}

func (c DataBucketDataEvent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataBucketDataEvent) UnmarshalJSON(b []byte) error {
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
