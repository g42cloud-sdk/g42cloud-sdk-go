package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TextWatermark struct {
	Dx *string `json:"dx,omitempty"`

	Dy *string `json:"dy,omitempty"`

	Referpos *string `json:"referpos,omitempty"`

	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineDuration *string `json:"timeline_duration,omitempty"`

	FontName *string `json:"font_name,omitempty"`

	FontSize *int32 `json:"font_size,omitempty"`

	FontColor *string `json:"font_color,omitempty"`

	Base *TextWatermarkBase `json:"base,omitempty"`
}

func (o TextWatermark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextWatermark struct{}"
	}

	return strings.Join([]string{"TextWatermark", string(data)}, " ")
}

type TextWatermarkBase struct {
	value string
}

type TextWatermarkBaseEnum struct {
	INPUT  TextWatermarkBase
	OUTPUT TextWatermarkBase
}

func GetTextWatermarkBaseEnum() TextWatermarkBaseEnum {
	return TextWatermarkBaseEnum{
		INPUT: TextWatermarkBase{
			value: "input",
		},
		OUTPUT: TextWatermarkBase{
			value: "output",
		},
	}
}

func (c TextWatermarkBase) Value() string {
	return c.value
}

func (c TextWatermarkBase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextWatermarkBase) UnmarshalJSON(b []byte) error {
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
