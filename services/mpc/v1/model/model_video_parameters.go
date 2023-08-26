package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type VideoParameters struct {
	OutputPolicy *VideoParametersOutputPolicy `json:"output_policy,omitempty"`

	Codec *int32 `json:"codec,omitempty"`

	Bitrate *int32 `json:"bitrate,omitempty"`

	Profile *int32 `json:"profile,omitempty"`

	Level *int32 `json:"level,omitempty"`

	Preset *int32 `json:"preset,omitempty"`

	MaxIframesInterval *int32 `json:"max_iframes_interval,omitempty"`

	BframesCount *int32 `json:"bframes_count,omitempty"`

	FrameRate *int32 `json:"frame_rate,omitempty"`

	Width *int32 `json:"width,omitempty"`

	Height *int32 `json:"height,omitempty"`

	BlackCut *int32 `json:"black_cut,omitempty"`
}

func (o VideoParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoParameters struct{}"
	}

	return strings.Join([]string{"VideoParameters", string(data)}, " ")
}

type VideoParametersOutputPolicy struct {
	value string
}

type VideoParametersOutputPolicyEnum struct {
	TRANSCODE VideoParametersOutputPolicy
	DISCARD   VideoParametersOutputPolicy
	COPY      VideoParametersOutputPolicy
}

func GetVideoParametersOutputPolicyEnum() VideoParametersOutputPolicyEnum {
	return VideoParametersOutputPolicyEnum{
		TRANSCODE: VideoParametersOutputPolicy{
			value: "transcode",
		},
		DISCARD: VideoParametersOutputPolicy{
			value: "discard",
		},
		COPY: VideoParametersOutputPolicy{
			value: "copy",
		},
	}
}

func (c VideoParametersOutputPolicy) Value() string {
	return c.value
}

func (c VideoParametersOutputPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoParametersOutputPolicy) UnmarshalJSON(b []byte) error {
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
