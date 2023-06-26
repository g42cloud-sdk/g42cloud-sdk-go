package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type WatermarkTemplate struct {
	Dx *string `json:"dx,omitempty"`

	Dy *string `json:"dy,omitempty"`

	Referpos *string `json:"referpos,omitempty"`

	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineDuration *string `json:"timeline_duration,omitempty"`

	ImageProcess *string `json:"image_process,omitempty"`

	Width *string `json:"width,omitempty"`

	Height *string `json:"height,omitempty"`

	Base *WatermarkTemplateBase `json:"base,omitempty"`

	TemplateId *int32 `json:"template_id,omitempty"`

	TemplateName *string `json:"template_name,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o WatermarkTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WatermarkTemplate struct{}"
	}

	return strings.Join([]string{"WatermarkTemplate", string(data)}, " ")
}

type WatermarkTemplateBase struct {
	value string
}

type WatermarkTemplateBaseEnum struct {
	INPUT  WatermarkTemplateBase
	OUTPUT WatermarkTemplateBase
}

func GetWatermarkTemplateBaseEnum() WatermarkTemplateBaseEnum {
	return WatermarkTemplateBaseEnum{
		INPUT: WatermarkTemplateBase{
			value: "input",
		},
		OUTPUT: WatermarkTemplateBase{
			value: "output",
		},
	}
}

func (c WatermarkTemplateBase) Value() string {
	return c.value
}

func (c WatermarkTemplateBase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WatermarkTemplateBase) UnmarshalJSON(b []byte) error {
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
