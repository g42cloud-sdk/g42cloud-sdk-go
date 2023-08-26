package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type OutputSetting struct {
	Format *OutputSettingFormat `json:"format,omitempty"`

	Video *EditVideoInfo `json:"video,omitempty"`

	Audio *EditAudioInfo `json:"audio,omitempty"`

	Hls *EditHlsInfo `json:"hls,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`
}

func (o OutputSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputSetting struct{}"
	}

	return strings.Join([]string{"OutputSetting", string(data)}, " ")
}

type OutputSettingFormat struct {
	value string
}

type OutputSettingFormatEnum struct {
	MP4 OutputSettingFormat
	HLS OutputSettingFormat
	TS  OutputSettingFormat
	FLV OutputSettingFormat
}

func GetOutputSettingFormatEnum() OutputSettingFormatEnum {
	return OutputSettingFormatEnum{
		MP4: OutputSettingFormat{
			value: "MP4",
		},
		HLS: OutputSettingFormat{
			value: "HLS",
		},
		TS: OutputSettingFormat{
			value: "TS",
		},
		FLV: OutputSettingFormat{
			value: "FLV",
		},
	}
}

func (c OutputSettingFormat) Value() string {
	return c.value
}

func (c OutputSettingFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OutputSettingFormat) UnmarshalJSON(b []byte) error {
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
