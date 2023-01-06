package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Volume struct {
	Type VolumeType `json:"type"`

	Size int32 `json:"size"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}

type VolumeType struct {
	value string
}

type VolumeTypeEnum struct {
	ULTRAHIGH    VolumeType
	HIGH         VolumeType
	COMMON       VolumeType
	NVMESSD      VolumeType
	ULTRAHIGHPRO VolumeType
	CLOUDSSD     VolumeType
	LOCALSSD     VolumeType
}

func GetVolumeTypeEnum() VolumeTypeEnum {
	return VolumeTypeEnum{
		ULTRAHIGH: VolumeType{
			value: "ULTRAHIGH",
		},
		HIGH: VolumeType{
			value: "HIGH",
		},
		COMMON: VolumeType{
			value: "COMMON",
		},
		NVMESSD: VolumeType{
			value: "NVMESSD",
		},
		ULTRAHIGHPRO: VolumeType{
			value: "ULTRAHIGHPRO",
		},
		CLOUDSSD: VolumeType{
			value: "CLOUDSSD",
		},
		LOCALSSD: VolumeType{
			value: "LOCALSSD",
		},
	}
}

func (c VolumeType) Value() string {
	return c.value
}

func (c VolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeType) UnmarshalJSON(b []byte) error {
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
