package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowMemberDetailRequest struct {
	BackupId string `json:"backup_id"`

	MemberId string `json:"member_id"`
}

func (o ShowMemberDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailRequest", string(data)}, " ")
}
