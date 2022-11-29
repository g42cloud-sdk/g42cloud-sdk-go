package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultRemoveResourceReq struct {
	ResourceIds []string `json:"resource_ids"`
}

func (o VaultRemoveResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultRemoveResourceReq struct{}"
	}

	return strings.Join([]string{"VaultRemoveResourceReq", string(data)}, " ")
}
