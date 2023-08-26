package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

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
