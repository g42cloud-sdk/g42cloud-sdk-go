package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BssParamForResizeVolume struct {
	IsAutoPay *BssParamForResizeVolumeIsAutoPay `json:"isAutoPay,omitempty"`
}

func (o BssParamForResizeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParamForResizeVolume struct{}"
	}

	return strings.Join([]string{"BssParamForResizeVolume", string(data)}, " ")
}

type BssParamForResizeVolumeIsAutoPay struct {
	value string
}

type BssParamForResizeVolumeIsAutoPayEnum struct {
	FALSE BssParamForResizeVolumeIsAutoPay
	TRUE  BssParamForResizeVolumeIsAutoPay
}

func GetBssParamForResizeVolumeIsAutoPayEnum() BssParamForResizeVolumeIsAutoPayEnum {
	return BssParamForResizeVolumeIsAutoPayEnum{
		FALSE: BssParamForResizeVolumeIsAutoPay{
			value: "false",
		},
		TRUE: BssParamForResizeVolumeIsAutoPay{
			value: "true",
		},
	}
}

func (c BssParamForResizeVolumeIsAutoPay) Value() string {
	return c.value
}

func (c BssParamForResizeVolumeIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForResizeVolumeIsAutoPay) UnmarshalJSON(b []byte) error {
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
