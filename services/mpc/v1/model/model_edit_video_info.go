package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type EditVideoInfo struct {
	Reference *EditVideoInfoReference `json:"reference,omitempty"`

	Width *int32 `json:"width,omitempty"`

	Height *int32 `json:"height,omitempty"`

	Codec *EditVideoInfoCodec `json:"codec,omitempty"`

	Bitrate *int32 `json:"bitrate,omitempty"`

	FrameRate *int32 `json:"frame_rate,omitempty"`
}

func (o EditVideoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditVideoInfo struct{}"
	}

	return strings.Join([]string{"EditVideoInfo", string(data)}, " ")
}

type EditVideoInfoReference struct {
	value string
}

type EditVideoInfoReferenceEnum struct {
	MAX                      EditVideoInfoReference
	MIN                      EditVideoInfoReference
	CUSTOM                   EditVideoInfoReference
	SHORT_HEIGHT_SHORT_WIDTH EditVideoInfoReference
}

func GetEditVideoInfoReferenceEnum() EditVideoInfoReferenceEnum {
	return EditVideoInfoReferenceEnum{
		MAX: EditVideoInfoReference{
			value: "MAX",
		},
		MIN: EditVideoInfoReference{
			value: "MIN",
		},
		CUSTOM: EditVideoInfoReference{
			value: "CUSTOM",
		},
		SHORT_HEIGHT_SHORT_WIDTH: EditVideoInfoReference{
			value: "SHORT_HEIGHT_SHORT_WIDTH",
		},
	}
}

func (c EditVideoInfoReference) Value() string {
	return c.value
}

func (c EditVideoInfoReference) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EditVideoInfoReference) UnmarshalJSON(b []byte) error {
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

type EditVideoInfoCodec struct {
	value string
}

type EditVideoInfoCodecEnum struct {
	H264 EditVideoInfoCodec
	H265 EditVideoInfoCodec
}

func GetEditVideoInfoCodecEnum() EditVideoInfoCodecEnum {
	return EditVideoInfoCodecEnum{
		H264: EditVideoInfoCodec{
			value: "H264",
		},
		H265: EditVideoInfoCodec{
			value: "H265",
		},
	}
}

func (c EditVideoInfoCodec) Value() string {
	return c.value
}

func (c EditVideoInfoCodec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EditVideoInfoCodec) UnmarshalJSON(b []byte) error {
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
