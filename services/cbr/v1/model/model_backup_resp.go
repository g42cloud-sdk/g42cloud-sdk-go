package model

import (
	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/sdktime"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"
	"strings"
)

type BackupResp struct {
	CheckpointId string `json:"checkpoint_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	Description string `json:"description"`

	ExpiredAt *sdktime.SdkTime `json:"expired_at"`

	ExtendInfo *BackupExtendInfo `json:"extend_info"`

	Id string `json:"id"`

	ImageType string `json:"image_type"`

	Name string `json:"name"`

	ParentId string `json:"parent_id"`

	ProjectId string `json:"project_id"`

	ProtectedAt string `json:"protected_at"`

	ResourceAz string `json:"resource_az"`

	ResourceId string `json:"resource_id"`

	ResourceName string `json:"resource_name"`

	ResourceSize int32 `json:"resource_size"`

	ResourceType string `json:"resource_type"`

	Status BackupRespStatus `json:"status"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	VaultId string `json:"vault_id"`

	ReplicationRecords *[]ReplicationRecordGet `json:"replication_records,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ProviderId string `json:"provider_id"`

	Children *[]BackupResp `json:"children,omitempty"`
}

func (o BackupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupResp struct{}"
	}

	return strings.Join([]string{"BackupResp", string(data)}, " ")
}

type BackupRespStatus struct {
	value string
}

type BackupRespStatusEnum struct {
	AVAILABLE       BackupRespStatus
	PROTECTING      BackupRespStatus
	DELETING        BackupRespStatus
	RESTORING       BackupRespStatus
	ERROR           BackupRespStatus
	WAITING_PROTECT BackupRespStatus
	WAITING_DELETE  BackupRespStatus
	WAITING_RESTORE BackupRespStatus
}

func GetBackupRespStatusEnum() BackupRespStatusEnum {
	return BackupRespStatusEnum{
		AVAILABLE: BackupRespStatus{
			value: "available",
		},
		PROTECTING: BackupRespStatus{
			value: "protecting",
		},
		DELETING: BackupRespStatus{
			value: "deleting",
		},
		RESTORING: BackupRespStatus{
			value: "restoring",
		},
		ERROR: BackupRespStatus{
			value: "error",
		},
		WAITING_PROTECT: BackupRespStatus{
			value: "waiting_protect",
		},
		WAITING_DELETE: BackupRespStatus{
			value: "waiting_delete",
		},
		WAITING_RESTORE: BackupRespStatus{
			value: "waiting_restore",
		},
	}
}

func (c BackupRespStatus) Value() string {
	return c.value
}

func (c BackupRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupRespStatus) UnmarshalJSON(b []byte) error {
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
