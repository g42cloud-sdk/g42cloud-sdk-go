package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EnterpriseProjectId struct {
}

func (o EnterpriseProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectId struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectId", string(data)}, " ")
}
