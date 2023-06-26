package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ImageWatermark struct {
	Dx *string `json:"dx,omitempty"`

	Dy *string `json:"dy,omitempty"`

	Referpos *string `json:"referpos,omitempty"`

	TimelineStart *string `json:"timeline_start,omitempty"`

	TimelineDuration *string `json:"timeline_duration,omitempty"`

	ImageProcess *string `json:"image_process,omitempty"`

	Width *string `json:"width,omitempty"`

	Height *string `json:"height,omitempty"`

	Base *ImageWatermarkBase `json:"base,omitempty"`
}

func (o ImageWatermark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWatermark struct{}"
	}

	return strings.Join([]string{"ImageWatermark", string(data)}, " ")
}

type ImageWatermarkBase struct {
	value string
}

type ImageWatermarkBaseEnum struct {
	INPUT  ImageWatermarkBase
	OUTPUT ImageWatermarkBase
}

func GetImageWatermarkBaseEnum() ImageWatermarkBaseEnum {
	return ImageWatermarkBaseEnum{
		INPUT: ImageWatermarkBase{
			value: "input",
		},
		OUTPUT: ImageWatermarkBase{
			value: "output",
		},
	}
}

func (c ImageWatermarkBase) Value() string {
	return c.value
}

func (c ImageWatermarkBase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageWatermarkBase) UnmarshalJSON(b []byte) error {
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
