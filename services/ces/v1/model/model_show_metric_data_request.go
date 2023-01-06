package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowMetricDataRequest struct {
	ContentType string `json:"Content-Type"`

	Namespace string `json:"namespace"`

	MetricName string `json:"metric_name"`

	Dim0 string `json:"dim.0"`

	Dim1 *string `json:"dim.1,omitempty"`

	Dim2 *string `json:"dim.2,omitempty"`

	Dim3 *string `json:"dim.3,omitempty"`

	Filter ShowMetricDataRequestFilter `json:"filter"`

	Period int32 `json:"period"`

	From int64 `json:"from"`

	To int64 `json:"to"`
}

func (o ShowMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricDataRequest", string(data)}, " ")
}

type ShowMetricDataRequestFilter struct {
	value string
}

type ShowMetricDataRequestFilterEnum struct {
	MAX      ShowMetricDataRequestFilter
	MIN      ShowMetricDataRequestFilter
	AVERAGE  ShowMetricDataRequestFilter
	SUM      ShowMetricDataRequestFilter
	VARIANCE ShowMetricDataRequestFilter
}

func GetShowMetricDataRequestFilterEnum() ShowMetricDataRequestFilterEnum {
	return ShowMetricDataRequestFilterEnum{
		MAX: ShowMetricDataRequestFilter{
			value: "max",
		},
		MIN: ShowMetricDataRequestFilter{
			value: "min",
		},
		AVERAGE: ShowMetricDataRequestFilter{
			value: "average",
		},
		SUM: ShowMetricDataRequestFilter{
			value: "sum",
		},
		VARIANCE: ShowMetricDataRequestFilter{
			value: "variance",
		},
	}
}

func (c ShowMetricDataRequestFilter) Value() string {
	return c.value
}

func (c ShowMetricDataRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricDataRequestFilter) UnmarshalJSON(b []byte) error {
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
