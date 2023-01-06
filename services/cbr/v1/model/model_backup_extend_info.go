package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BackupExtendInfo struct {
	AutoTrigger *bool `json:"auto_trigger,omitempty"`

	Bootable *bool `json:"bootable,omitempty"`

	Incremental *bool `json:"incremental,omitempty"`

	SnapshotId *string `json:"snapshot_id,omitempty"`

	SupportLld *bool `json:"support_lld,omitempty"`

	SupportedRestoreMode *BackupExtendInfoSupportedRestoreMode `json:"supported_restore_mode,omitempty"`

	OsImagesData *[]ImageData `json:"os_images_data,omitempty"`

	ContainSystemDisk *bool `json:"contain_system_disk,omitempty"`

	Encrypted *bool `json:"encrypted,omitempty"`

	SystemDisk *bool `json:"system_disk,omitempty"`
}

func (o BackupExtendInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupExtendInfo struct{}"
	}

	return strings.Join([]string{"BackupExtendInfo", string(data)}, " ")
}

type BackupExtendInfoSupportedRestoreMode struct {
	value string
}

type BackupExtendInfoSupportedRestoreModeEnum struct {
	NA       BackupExtendInfoSupportedRestoreMode
	BACKUP   BackupExtendInfoSupportedRestoreMode
	SNAPSHOT BackupExtendInfoSupportedRestoreMode
}

func GetBackupExtendInfoSupportedRestoreModeEnum() BackupExtendInfoSupportedRestoreModeEnum {
	return BackupExtendInfoSupportedRestoreModeEnum{
		NA: BackupExtendInfoSupportedRestoreMode{
			value: "na",
		},
		BACKUP: BackupExtendInfoSupportedRestoreMode{
			value: " backup",
		},
		SNAPSHOT: BackupExtendInfoSupportedRestoreMode{
			value: " snapshot",
		},
	}
}

func (c BackupExtendInfoSupportedRestoreMode) Value() string {
	return c.value
}

func (c BackupExtendInfoSupportedRestoreMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupExtendInfoSupportedRestoreMode) UnmarshalJSON(b []byte) error {
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
