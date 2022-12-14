package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ForceRedirect struct {
	Switch int32 `json:"switch"`

	RedirectType *string `json:"redirect_type,omitempty"`
}

func (o ForceRedirect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForceRedirect struct{}"
	}

	return strings.Join([]string{"ForceRedirect", string(data)}, " ")
}
