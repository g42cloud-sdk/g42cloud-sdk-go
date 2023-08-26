package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type VideoCommon struct {
	OutputPolicy *VideoCommonOutputPolicy `json:"output_policy,omitempty"`

	Codec *int32 `json:"codec,omitempty"`

	Profile *int32 `json:"profile,omitempty"`

	Level *int32 `json:"level,omitempty"`

	Preset *int32 `json:"preset,omitempty"`

	MaxIframesInterval *int32 `json:"max_iframes_interval,omitempty"`

	BframesCount *int32 `json:"bframes_count,omitempty"`

	FrameRate *int32 `json:"frame_rate,omitempty"`

	BlackCut *int32 `json:"black_cut,omitempty"`
}

func (o VideoCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCommon struct{}"
	}

	return strings.Join([]string{"VideoCommon", string(data)}, " ")
}

type VideoCommonOutputPolicy struct {
	value string
}

type VideoCommonOutputPolicyEnum struct {
	TRANSCODE VideoCommonOutputPolicy
	DISCARD   VideoCommonOutputPolicy
	COPY      VideoCommonOutputPolicy
}

func GetVideoCommonOutputPolicyEnum() VideoCommonOutputPolicyEnum {
	return VideoCommonOutputPolicyEnum{
		TRANSCODE: VideoCommonOutputPolicy{
			value: "transcode",
		},
		DISCARD: VideoCommonOutputPolicy{
			value: "discard",
		},
		COPY: VideoCommonOutputPolicy{
			value: "copy",
		},
	}
}

func (c VideoCommonOutputPolicy) Value() string {
	return c.value
}

func (c VideoCommonOutputPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoCommonOutputPolicy) UnmarshalJSON(b []byte) error {
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
