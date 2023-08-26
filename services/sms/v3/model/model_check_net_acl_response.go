package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CheckNetAclResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckNetAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNetAclResponse struct{}"
	}

	return strings.Join([]string{"CheckNetAclResponse", string(data)}, " ")
}
