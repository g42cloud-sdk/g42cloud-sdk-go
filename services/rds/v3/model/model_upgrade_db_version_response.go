package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpgradeDbVersionResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeDbVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDbVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeDbVersionResponse", string(data)}, " ")
}
