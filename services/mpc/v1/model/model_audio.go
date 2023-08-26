package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Audio struct {
	OutputPolicy *AudioOutputPolicy `json:"output_policy,omitempty"`

	Codec *int32 `json:"codec,omitempty"`

	SampleRate *int32 `json:"sample_rate,omitempty"`

	Bitrate *int32 `json:"bitrate,omitempty"`

	Channels int32 `json:"channels"`
}

func (o Audio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Audio struct{}"
	}

	return strings.Join([]string{"Audio", string(data)}, " ")
}

type AudioOutputPolicy struct {
	value string
}

type AudioOutputPolicyEnum struct {
	TRANSCODE AudioOutputPolicy
	DISCARD   AudioOutputPolicy
	COPY      AudioOutputPolicy
}

func GetAudioOutputPolicyEnum() AudioOutputPolicyEnum {
	return AudioOutputPolicyEnum{
		TRANSCODE: AudioOutputPolicy{
			value: "transcode",
		},
		DISCARD: AudioOutputPolicy{
			value: "discard",
		},
		COPY: AudioOutputPolicy{
			value: "copy",
		},
	}
}

func (c AudioOutputPolicy) Value() string {
	return c.value
}

func (c AudioOutputPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioOutputPolicy) UnmarshalJSON(b []byte) error {
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
