package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type CreateVolumeOption struct {
	AvailabilityZone string `json:"availability_zone"`

	BackupId *string `json:"backup_id,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ImageRef *string `json:"imageRef,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	Multiattach *bool `json:"multiattach,omitempty"`

	Name *string `json:"name,omitempty"`

	Size int32 `json:"size"`

	SnapshotId *string `json:"snapshot_id,omitempty"`

	VolumeType CreateVolumeOptionVolumeType `json:"volume_type"`

	Tags map[string]string `json:"tags,omitempty"`
}

func (o CreateVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeOption struct{}"
	}

	return strings.Join([]string{"CreateVolumeOption", string(data)}, " ")
}

type CreateVolumeOptionVolumeType struct {
	value string
}

type CreateVolumeOptionVolumeTypeEnum struct {
	SSD   CreateVolumeOptionVolumeType
	GPSSD CreateVolumeOptionVolumeType
	SAS   CreateVolumeOptionVolumeType
	SATA  CreateVolumeOptionVolumeType
	ESSD  CreateVolumeOptionVolumeType
}

func GetCreateVolumeOptionVolumeTypeEnum() CreateVolumeOptionVolumeTypeEnum {
	return CreateVolumeOptionVolumeTypeEnum{
		SSD: CreateVolumeOptionVolumeType{
			value: "SSD",
		},
		GPSSD: CreateVolumeOptionVolumeType{
			value: "GPSSD",
		},
		SAS: CreateVolumeOptionVolumeType{
			value: "SAS",
		},
		SATA: CreateVolumeOptionVolumeType{
			value: "SATA",
		},
		ESSD: CreateVolumeOptionVolumeType{
			value: "ESSD",
		},
	}
}

func (c CreateVolumeOptionVolumeType) Value() string {
	return c.value
}

func (c CreateVolumeOptionVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVolumeOptionVolumeType) UnmarshalJSON(b []byte) error {
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
