package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UnlockTargetEcsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnlockTargetEcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockTargetEcsResponse struct{}"
	}

	return strings.Join([]string{"UnlockTargetEcsResponse", string(data)}, " ")
}
