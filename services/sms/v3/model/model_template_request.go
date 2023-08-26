package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type TemplateRequest struct {
	Name string `json:"name"`

	IsTemplate bool `json:"is_template"`

	Region string `json:"region"`

	Projectid string `json:"projectid"`

	TargetServerName *string `json:"target_server_name,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Volumetype *TemplateRequestVolumetype `json:"volumetype,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	Vpc *VpcObject `json:"vpc,omitempty"`

	Nics *[]Nics `json:"nics,omitempty"`

	SecurityGroups *[]SgObject `json:"security_groups,omitempty"`

	Publicip *PublicIp `json:"publicip,omitempty"`

	Disk *[]TemplateDisk `json:"disk,omitempty"`

	DataVolumeType *TemplateRequestDataVolumeType `json:"data_volume_type,omitempty"`

	TargetPassword *string `json:"target_password,omitempty"`

	ImageId *string `json:"image_id,omitempty"`
}

func (o TemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRequest struct{}"
	}

	return strings.Join([]string{"TemplateRequest", string(data)}, " ")
}

type TemplateRequestVolumetype struct {
	value string
}

type TemplateRequestVolumetypeEnum struct {
	SAS  TemplateRequestVolumetype
	SSD  TemplateRequestVolumetype
	SATA TemplateRequestVolumetype
}

func GetTemplateRequestVolumetypeEnum() TemplateRequestVolumetypeEnum {
	return TemplateRequestVolumetypeEnum{
		SAS: TemplateRequestVolumetype{
			value: "SAS",
		},
		SSD: TemplateRequestVolumetype{
			value: "SSD",
		},
		SATA: TemplateRequestVolumetype{
			value: "SATA",
		},
	}
}

func (c TemplateRequestVolumetype) Value() string {
	return c.value
}

func (c TemplateRequestVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateRequestVolumetype) UnmarshalJSON(b []byte) error {
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

type TemplateRequestDataVolumeType struct {
	value string
}

type TemplateRequestDataVolumeTypeEnum struct {
	SAS  TemplateRequestDataVolumeType
	SSD  TemplateRequestDataVolumeType
	SATA TemplateRequestDataVolumeType
}

func GetTemplateRequestDataVolumeTypeEnum() TemplateRequestDataVolumeTypeEnum {
	return TemplateRequestDataVolumeTypeEnum{
		SAS: TemplateRequestDataVolumeType{
			value: "SAS",
		},
		SSD: TemplateRequestDataVolumeType{
			value: "SSD",
		},
		SATA: TemplateRequestDataVolumeType{
			value: "SATA",
		},
	}
}

func (c TemplateRequestDataVolumeType) Value() string {
	return c.value
}

func (c TemplateRequestDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateRequestDataVolumeType) UnmarshalJSON(b []byte) error {
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
