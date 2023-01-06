package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowReferResponse struct {
	Referer        *RefererRsp `json:"referer,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowReferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReferResponse struct{}"
	}

	return strings.Join([]string{"ShowReferResponse", string(data)}, " ")
}
