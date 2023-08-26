package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeletetemplatesReq struct {
	Ids *[]string `json:"ids,omitempty"`
}

func (o DeletetemplatesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletetemplatesReq struct{}"
	}

	return strings.Join([]string{"DeletetemplatesReq", string(data)}, " ")
}
