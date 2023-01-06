package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type DeleteTrackerRequest struct {
	TrackerName *string `json:"tracker_name,omitempty"`

	TrackerType *DeleteTrackerRequestTrackerType `json:"tracker_type,omitempty"`
}

func (o DeleteTrackerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrackerRequest", string(data)}, " ")
}

type DeleteTrackerRequestTrackerType struct {
	value string
}

type DeleteTrackerRequestTrackerTypeEnum struct {
	DATA DeleteTrackerRequestTrackerType
}

func GetDeleteTrackerRequestTrackerTypeEnum() DeleteTrackerRequestTrackerTypeEnum {
	return DeleteTrackerRequestTrackerTypeEnum{
		DATA: DeleteTrackerRequestTrackerType{
			value: "data",
		},
	}
}

func (c DeleteTrackerRequestTrackerType) Value() string {
	return c.value
}

func (c DeleteTrackerRequestTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTrackerRequestTrackerType) UnmarshalJSON(b []byte) error {
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
