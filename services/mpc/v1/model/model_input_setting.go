package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type InputSetting struct {
	Input *ObsObjInfo `json:"input"`

	PaneId string `json:"pane_id"`

	AudioPolicy *InputSettingAudioPolicy `json:"audio_policy,omitempty"`
}

func (o InputSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputSetting struct{}"
	}

	return strings.Join([]string{"InputSetting", string(data)}, " ")
}

type InputSettingAudioPolicy struct {
	value string
}

type InputSettingAudioPolicyEnum struct {
	DISCARD InputSettingAudioPolicy
	RESERVE InputSettingAudioPolicy
}

func GetInputSettingAudioPolicyEnum() InputSettingAudioPolicyEnum {
	return InputSettingAudioPolicyEnum{
		DISCARD: InputSettingAudioPolicy{
			value: "DISCARD",
		},
		RESERVE: InputSettingAudioPolicy{
			value: "RESERVE",
		},
	}
}

func (c InputSettingAudioPolicy) Value() string {
	return c.value
}

func (c InputSettingAudioPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InputSettingAudioPolicy) UnmarshalJSON(b []byte) error {
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
