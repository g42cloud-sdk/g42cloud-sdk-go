package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ImageWatermarkSetting struct {
	Dx *string `json:"dx,omitempty"`

	Dy *string `json:"dy,omitempty"`

	Referpos *string `json:"referpos,omitempty"`

	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineDuration *string `json:"timeline_duration,omitempty"`

	OverlayInput *string `json:"overlay_input,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`

	Base *ImageWatermarkSettingBase `json:"base,omitempty"`
}

func (o ImageWatermarkSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWatermarkSetting struct{}"
	}

	return strings.Join([]string{"ImageWatermarkSetting", string(data)}, " ")
}

type ImageWatermarkSettingBase struct {
	value string
}

type ImageWatermarkSettingBaseEnum struct {
	INPUT  ImageWatermarkSettingBase
	OUTPUT ImageWatermarkSettingBase
}

func GetImageWatermarkSettingBaseEnum() ImageWatermarkSettingBaseEnum {
	return ImageWatermarkSettingBaseEnum{
		INPUT: ImageWatermarkSettingBase{
			value: "input",
		},
		OUTPUT: ImageWatermarkSettingBase{
			value: "output",
		},
	}
}

func (c ImageWatermarkSettingBase) Value() string {
	return c.value
}

func (c ImageWatermarkSettingBase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageWatermarkSettingBase) UnmarshalJSON(b []byte) error {
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
