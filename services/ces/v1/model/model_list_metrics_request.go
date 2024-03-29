package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListMetricsRequest struct {
	ContentType string `json:"Content-Type"`

	Dim0 *string `json:"dim.0,omitempty"`

	Dim1 *string `json:"dim.1,omitempty"`

	Dim2 *string `json:"dim.2,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	MetricName *string `json:"metric_name,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	Order *ListMetricsRequestOrder `json:"order,omitempty"`

	Start *string `json:"start,omitempty"`
}

func (o ListMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsRequest", string(data)}, " ")
}

type ListMetricsRequestOrder struct {
	value string
}

type ListMetricsRequestOrderEnum struct {
	ASC  ListMetricsRequestOrder
	DESC ListMetricsRequestOrder
}

func GetListMetricsRequestOrderEnum() ListMetricsRequestOrderEnum {
	return ListMetricsRequestOrderEnum{
		ASC: ListMetricsRequestOrder{
			value: "asc",
		},
		DESC: ListMetricsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListMetricsRequestOrder) Value() string {
	return c.value
}

func (c ListMetricsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricsRequestOrder) UnmarshalJSON(b []byte) error {
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
