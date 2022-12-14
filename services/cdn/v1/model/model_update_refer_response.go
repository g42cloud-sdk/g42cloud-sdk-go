package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateReferResponse struct {
	Referer        *RefererRsp `json:"referer,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateReferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReferResponse struct{}"
	}

	return strings.Join([]string{"UpdateReferResponse", string(data)}, " ")
}
