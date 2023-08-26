package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PolicyCreateReq struct {
	Policy *PolicyCreate `json:"policy"`
}

func (o PolicyCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyCreateReq struct{}"
	}

	return strings.Join([]string{"PolicyCreateReq", string(data)}, " ")
}
