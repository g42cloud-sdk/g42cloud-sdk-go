package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrRsp struct {
	Error *ErrMsg `json:"error"`
}

func (o ErrRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrRsp struct{}"
	}

	return strings.Join([]string{"ErrRsp", string(data)}, " ")
}
