package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type AnimatedGraphicsOutputParam struct {
	Format *AnimatedGraphicsOutputParamFormat `json:"format,omitempty"`

	Width int32 `json:"width"`

	Height int32 `json:"height"`

	Start int32 `json:"start"`

	End int32 `json:"end"`

	FrameRate *int32 `json:"frame_rate,omitempty"`
}

func (o AnimatedGraphicsOutputParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnimatedGraphicsOutputParam struct{}"
	}

	return strings.Join([]string{"AnimatedGraphicsOutputParam", string(data)}, " ")
}

type AnimatedGraphicsOutputParamFormat struct {
	value string
}

type AnimatedGraphicsOutputParamFormatEnum struct {
	GIF AnimatedGraphicsOutputParamFormat
}

func GetAnimatedGraphicsOutputParamFormatEnum() AnimatedGraphicsOutputParamFormatEnum {
	return AnimatedGraphicsOutputParamFormatEnum{
		GIF: AnimatedGraphicsOutputParamFormat{
			value: "gif",
		},
	}
}

func (c AnimatedGraphicsOutputParamFormat) Value() string {
	return c.value
}

func (c AnimatedGraphicsOutputParamFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnimatedGraphicsOutputParamFormat) UnmarshalJSON(b []byte) error {
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
