package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Member struct {
	Status MemberStatus `json:"status"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	ImageId *string `json:"image_id,omitempty"`

	DestProjectId *string `json:"dest_project_id,omitempty"`

	VaultId *string `json:"vault_id,omitempty"`

	Id *string `json:"id,omitempty"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}

type MemberStatus struct {
	value string
}

type MemberStatusEnum struct {
	PENDING  MemberStatus
	ACCEPTED MemberStatus
	REJECTED MemberStatus
}

func GetMemberStatusEnum() MemberStatusEnum {
	return MemberStatusEnum{
		PENDING: MemberStatus{
			value: "pending",
		},
		ACCEPTED: MemberStatus{
			value: "accepted",
		},
		REJECTED: MemberStatus{
			value: "rejected",
		},
	}
}

func (c MemberStatus) Value() string {
	return c.value
}

func (c MemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberStatus) UnmarshalJSON(b []byte) error {
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
