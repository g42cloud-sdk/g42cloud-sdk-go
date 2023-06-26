package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CommonQueryTaskRsp struct {
	Total *int32 `json:"total,omitempty"`
}

func (o CommonQueryTaskRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonQueryTaskRsp struct{}"
	}

	return strings.Join([]string{"CommonQueryTaskRsp", string(data)}, " ")
}
