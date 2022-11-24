package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateReadWeightResponse struct {
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReadWeightResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReadWeightResponse struct{}"
	}

	return strings.Join([]string{"UpdateReadWeightResponse", string(data)}, " ")
}
