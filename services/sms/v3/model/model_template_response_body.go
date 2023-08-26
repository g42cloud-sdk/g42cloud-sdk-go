package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TemplateResponseBody struct {
	Id *string `json:"id,omitempty"`

	Name string `json:"name"`

	IsTemplate *string `json:"is_template,omitempty"`

	Region string `json:"region"`

	Projectid string `json:"projectid"`

	TargetServerName string `json:"target_server_name"`

	AvailabilityZone string `json:"availability_zone"`

	Volumetype TemplateResponseBodyVolumetype `json:"volumetype"`

	Flavor string `json:"flavor"`

	Vpc *VpcObject `json:"vpc"`

	Nics []Nics `json:"nics"`

	SecurityGroups []SgObject `json:"security_groups"`

	Publicip *PublicIp `json:"publicip"`

	Disk []TemplateDisk `json:"disk"`

	DataVolumeType TemplateResponseBodyDataVolumeType `json:"data_volume_type"`

	TargetPassword string `json:"target_password"`

	ImageId *string `json:"image_id,omitempty"`
}

func (o TemplateResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateResponseBody struct{}"
	}

	return strings.Join([]string{"TemplateResponseBody", string(data)}, " ")
}

type TemplateResponseBodyVolumetype struct {
	value string
}

type TemplateResponseBodyVolumetypeEnum struct {
	SAS  TemplateResponseBodyVolumetype
	SSD  TemplateResponseBodyVolumetype
	SATA TemplateResponseBodyVolumetype
}

func GetTemplateResponseBodyVolumetypeEnum() TemplateResponseBodyVolumetypeEnum {
	return TemplateResponseBodyVolumetypeEnum{
		SAS: TemplateResponseBodyVolumetype{
			value: "SAS",
		},
		SSD: TemplateResponseBodyVolumetype{
			value: "SSD",
		},
		SATA: TemplateResponseBodyVolumetype{
			value: "SATA",
		},
	}
}

func (c TemplateResponseBodyVolumetype) Value() string {
	return c.value
}

func (c TemplateResponseBodyVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateResponseBodyVolumetype) UnmarshalJSON(b []byte) error {
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

type TemplateResponseBodyDataVolumeType struct {
	value string
}

type TemplateResponseBodyDataVolumeTypeEnum struct {
	SAS  TemplateResponseBodyDataVolumeType
	SSD  TemplateResponseBodyDataVolumeType
	SATA TemplateResponseBodyDataVolumeType
}

func GetTemplateResponseBodyDataVolumeTypeEnum() TemplateResponseBodyDataVolumeTypeEnum {
	return TemplateResponseBodyDataVolumeTypeEnum{
		SAS: TemplateResponseBodyDataVolumeType{
			value: "SAS",
		},
		SSD: TemplateResponseBodyDataVolumeType{
			value: "SSD",
		},
		SATA: TemplateResponseBodyDataVolumeType{
			value: "SATA",
		},
	}
}

func (c TemplateResponseBodyDataVolumeType) Value() string {
	return c.value
}

func (c TemplateResponseBodyDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateResponseBodyDataVolumeType) UnmarshalJSON(b []byte) error {
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
