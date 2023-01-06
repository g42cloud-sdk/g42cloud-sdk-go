package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type UpdateMember struct {
	Status UpdateMemberStatus `json:"status"`

	VaultId *string `json:"vault_id,omitempty"`
}

func (o UpdateMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMember struct{}"
	}

	return strings.Join([]string{"UpdateMember", string(data)}, " ")
}

type UpdateMemberStatus struct {
	value string
}

type UpdateMemberStatusEnum struct {
	ACCEPTED UpdateMemberStatus
	PENDING  UpdateMemberStatus
	REJECTED UpdateMemberStatus
}

func GetUpdateMemberStatusEnum() UpdateMemberStatusEnum {
	return UpdateMemberStatusEnum{
		ACCEPTED: UpdateMemberStatus{
			value: "accepted",
		},
		PENDING: UpdateMemberStatus{
			value: "pending",
		},
		REJECTED: UpdateMemberStatus{
			value: "rejected",
		},
	}
}

func (c UpdateMemberStatus) Value() string {
	return c.value
}

func (c UpdateMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMemberStatus) UnmarshalJSON(b []byte) error {
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
