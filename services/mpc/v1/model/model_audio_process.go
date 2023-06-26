package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type AudioProcess struct {
	Volume *AudioProcessVolume `json:"volume,omitempty"`

	VolumeExpr *int32 `json:"volume_expr,omitempty"`
}

func (o AudioProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioProcess struct{}"
	}

	return strings.Join([]string{"AudioProcess", string(data)}, " ")
}

type AudioProcessVolume struct {
	value string
}

type AudioProcessVolumeEnum struct {
	AUTO     AudioProcessVolume
	DYNAMIC  AudioProcessVolume
	ORIGINAL AudioProcessVolume
}

func GetAudioProcessVolumeEnum() AudioProcessVolumeEnum {
	return AudioProcessVolumeEnum{
		AUTO: AudioProcessVolume{
			value: "auto",
		},
		DYNAMIC: AudioProcessVolume{
			value: "dynamic",
		},
		ORIGINAL: AudioProcessVolume{
			value: "original",
		},
	}
}

func (c AudioProcessVolume) Value() string {
	return c.value
}

func (c AudioProcessVolume) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioProcessVolume) UnmarshalJSON(b []byte) error {
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
