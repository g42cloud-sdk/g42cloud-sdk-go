package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddMembersReq struct {
	Members []string `json:"members"`
}

func (o AddMembersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMembersReq struct{}"
	}

	return strings.Join([]string{"AddMembersReq", string(data)}, " ")
}
