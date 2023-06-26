package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ThumbnailPara struct {
	Type *ThumbnailParaType `json:"type,omitempty"`

	Time *int32 `json:"time,omitempty"`

	StartTime *int32 `json:"start_time,omitempty"`

	Duration *int32 `json:"duration,omitempty"`

	Dots *[]int32 `json:"dots,omitempty"`

	OutputFilename *string `json:"output_filename,omitempty"`

	Format *int32 `json:"format,omitempty"`

	Width *int32 `json:"width,omitempty"`

	Height *int32 `json:"height,omitempty"`

	MaxLength *int32 `json:"max_length,omitempty"`
}

func (o ThumbnailPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThumbnailPara struct{}"
	}

	return strings.Join([]string{"ThumbnailPara", string(data)}, " ")
}

type ThumbnailParaType struct {
	value string
}

type ThumbnailParaTypeEnum struct {
	PERCENT ThumbnailParaType
	TIME    ThumbnailParaType
	DOTS    ThumbnailParaType
}

func GetThumbnailParaTypeEnum() ThumbnailParaTypeEnum {
	return ThumbnailParaTypeEnum{
		PERCENT: ThumbnailParaType{
			value: "PERCENT",
		},
		TIME: ThumbnailParaType{
			value: "TIME",
		},
		DOTS: ThumbnailParaType{
			value: "DOTS",
		},
	}
}

func (c ThumbnailParaType) Value() string {
	return c.value
}

func (c ThumbnailParaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThumbnailParaType) UnmarshalJSON(b []byte) error {
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