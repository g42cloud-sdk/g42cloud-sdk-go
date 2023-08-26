package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type VideoProcess struct {
	HlsInitCount *int32 `json:"hls_init_count,omitempty"`

	HlsInitInterval *int32 `json:"hls_init_interval,omitempty"`

	Rotate *int32 `json:"rotate,omitempty"`

	Adaptation *VideoProcessAdaptation `json:"adaptation,omitempty"`

	Upsample *int32 `json:"upsample,omitempty"`
}

func (o VideoProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoProcess struct{}"
	}

	return strings.Join([]string{"VideoProcess", string(data)}, " ")
}

type VideoProcessAdaptation struct {
	value string
}

type VideoProcessAdaptationEnum struct {
	SHORT VideoProcessAdaptation
	LONG  VideoProcessAdaptation
	NONE  VideoProcessAdaptation
}

func GetVideoProcessAdaptationEnum() VideoProcessAdaptationEnum {
	return VideoProcessAdaptationEnum{
		SHORT: VideoProcessAdaptation{
			value: "SHORT",
		},
		LONG: VideoProcessAdaptation{
			value: "LONG",
		},
		NONE: VideoProcessAdaptation{
			value: "NONE",
		},
	}
}

func (c VideoProcessAdaptation) Value() string {
	return c.value
}

func (c VideoProcessAdaptation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoProcessAdaptation) UnmarshalJSON(b []byte) error {
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
