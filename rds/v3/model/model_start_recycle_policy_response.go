package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartRecyclePolicyResponse struct {
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"StartRecyclePolicyResponse", string(data)}, " ")
}
