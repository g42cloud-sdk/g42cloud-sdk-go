package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListTrackersRequest struct {
	TrackerName *string `json:"tracker_name,omitempty"`

	TrackerType *ListTrackersRequestTrackerType `json:"tracker_type,omitempty"`
}

func (o ListTrackersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrackersRequest struct{}"
	}

	return strings.Join([]string{"ListTrackersRequest", string(data)}, " ")
}

type ListTrackersRequestTrackerType struct {
	value string
}

type ListTrackersRequestTrackerTypeEnum struct {
	SYSTEM ListTrackersRequestTrackerType
	DATA   ListTrackersRequestTrackerType
}

func GetListTrackersRequestTrackerTypeEnum() ListTrackersRequestTrackerTypeEnum {
	return ListTrackersRequestTrackerTypeEnum{
		SYSTEM: ListTrackersRequestTrackerType{
			value: "system",
		},
		DATA: ListTrackersRequestTrackerType{
			value: "data",
		},
	}
}

func (c ListTrackersRequestTrackerType) Value() string {
	return c.value
}

func (c ListTrackersRequestTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTrackersRequestTrackerType) UnmarshalJSON(b []byte) error {
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
