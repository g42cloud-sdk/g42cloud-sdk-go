package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
