package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CinderAcceptVolumeTransferResponse struct {
	Transfer       *VolumeTransferSummary `json:"transfer,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CinderAcceptVolumeTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderAcceptVolumeTransferResponse struct{}"
	}

	return strings.Join([]string{"CinderAcceptVolumeTransferResponse", string(data)}, " ")
}
