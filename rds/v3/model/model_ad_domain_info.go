package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdDomainInfo struct {
	DomainAdminAccountName string `json:"domain_admin_account_name"`

	DomainAdminPwd string `json:"domain_admin_pwd"`
}

func (o AdDomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdDomainInfo struct{}"
	}

	return strings.Join([]string{"AdDomainInfo", string(data)}, " ")
}
