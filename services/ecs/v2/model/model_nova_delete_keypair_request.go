package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaDeleteKeypairRequest struct {
	KeypairName string `json:"keypair_name"`
}

func (o NovaDeleteKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDeleteKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaDeleteKeypairRequest", string(data)}, " ")
}
