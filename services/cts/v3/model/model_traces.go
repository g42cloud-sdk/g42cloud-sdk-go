package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Traces struct {
	ResourceId *string `json:"resource_id,omitempty"`

	TraceName *string `json:"trace_name,omitempty"`

	TraceRating *TracesTraceRating `json:"trace_rating,omitempty"`

	TraceType *string `json:"trace_type,omitempty"`

	Request *string `json:"request,omitempty"`

	Response *string `json:"response,omitempty"`

	Code *string `json:"code,omitempty"`

	ApiVersion *string `json:"api_version,omitempty"`

	Message *string `json:"message,omitempty"`

	RecordTime *int64 `json:"record_time,omitempty"`

	TraceId *string `json:"trace_id,omitempty"`

	Time *int64 `json:"time,omitempty"`

	User *UserInfo `json:"user,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	ResourceType *string `json:"resource_type,omitempty"`

	SourceIp *string `json:"source_ip,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	LocationInfo *string `json:"location_info,omitempty"`

	Endpoint *string `json:"endpoint,omitempty"`

	ResourceUrl *string `json:"resource_url,omitempty"`
}

func (o Traces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Traces struct{}"
	}

	return strings.Join([]string{"Traces", string(data)}, " ")
}

type TracesTraceRating struct {
	value string
}

type TracesTraceRatingEnum struct {
	NORMAL   TracesTraceRating
	WARNING  TracesTraceRating
	INCIDENT TracesTraceRating
}

func GetTracesTraceRatingEnum() TracesTraceRatingEnum {
	return TracesTraceRatingEnum{
		NORMAL: TracesTraceRating{
			value: "normal",
		},
		WARNING: TracesTraceRating{
			value: "warning",
		},
		INCIDENT: TracesTraceRating{
			value: "incident",
		},
	}
}

func (c TracesTraceRating) Value() string {
	return c.value
}

func (c TracesTraceRating) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TracesTraceRating) UnmarshalJSON(b []byte) error {
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
