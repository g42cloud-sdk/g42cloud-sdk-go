package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type DetachServerVolumeRequest struct {
	ServerId string `json:"server_id"`

	VolumeId string `json:"volume_id"`

	DeleteFlag *DetachServerVolumeRequestDeleteFlag `json:"delete_flag,omitempty"`
}

func (o DetachServerVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachServerVolumeRequest struct{}"
	}

	return strings.Join([]string{"DetachServerVolumeRequest", string(data)}, " ")
}

type DetachServerVolumeRequestDeleteFlag struct {
	value string
}

type DetachServerVolumeRequestDeleteFlagEnum struct {
	E_0 DetachServerVolumeRequestDeleteFlag
	E_1 DetachServerVolumeRequestDeleteFlag
}

func GetDetachServerVolumeRequestDeleteFlagEnum() DetachServerVolumeRequestDeleteFlagEnum {
	return DetachServerVolumeRequestDeleteFlagEnum{
		E_0: DetachServerVolumeRequestDeleteFlag{
			value: "0",
		},
		E_1: DetachServerVolumeRequestDeleteFlag{
			value: "1",
		},
	}
}

func (c DetachServerVolumeRequestDeleteFlag) Value() string {
	return c.value
}

func (c DetachServerVolumeRequestDeleteFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetachServerVolumeRequestDeleteFlag) UnmarshalJSON(b []byte) error {
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
