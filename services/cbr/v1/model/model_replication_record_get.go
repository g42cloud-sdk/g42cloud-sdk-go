package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ReplicationRecordGet struct {
	CreatedAt *string `json:"created_at,omitempty"`

	DestinationBackupId *string `json:"destination_backup_id,omitempty"`

	DestinationCheckpointId *string `json:"destination_checkpoint_id,omitempty"`

	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	DestinationRegion *string `json:"destination_region,omitempty"`

	DestinationVaultId *string `json:"destination_vault_id,omitempty"`

	ExtraInfo *ReplicationRecordsExtraInfo `json:"extra_info,omitempty"`

	Id string `json:"id"`

	SourceBackupId *string `json:"source_backup_id,omitempty"`

	SourceCheckpointId *string `json:"source_checkpoint_id,omitempty"`

	SourceProjectId *string `json:"source_project_id,omitempty"`

	SourceRegion *string `json:"source_region,omitempty"`

	Status *ReplicationRecordGetStatus `json:"status,omitempty"`

	VaultId *string `json:"vault_id,omitempty"`
}

func (o ReplicationRecordGet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationRecordGet struct{}"
	}

	return strings.Join([]string{"ReplicationRecordGet", string(data)}, " ")
}

type ReplicationRecordGetStatus struct {
	value string
}

type ReplicationRecordGetStatusEnum struct {
	REPLICATING       ReplicationRecordGetStatus
	SUCCESS           ReplicationRecordGetStatus
	FAIL              ReplicationRecordGetStatus
	SKIP              ReplicationRecordGetStatus
	WAITING_REPLICATE ReplicationRecordGetStatus
}

func GetReplicationRecordGetStatusEnum() ReplicationRecordGetStatusEnum {
	return ReplicationRecordGetStatusEnum{
		REPLICATING: ReplicationRecordGetStatus{
			value: "replicating",
		},
		SUCCESS: ReplicationRecordGetStatus{
			value: "success",
		},
		FAIL: ReplicationRecordGetStatus{
			value: "fail",
		},
		SKIP: ReplicationRecordGetStatus{
			value: "skip",
		},
		WAITING_REPLICATE: ReplicationRecordGetStatus{
			value: "waiting_replicate",
		},
	}
}

func (c ReplicationRecordGetStatus) Value() string {
	return c.value
}

func (c ReplicationRecordGetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReplicationRecordGetStatus) UnmarshalJSON(b []byte) error {
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
