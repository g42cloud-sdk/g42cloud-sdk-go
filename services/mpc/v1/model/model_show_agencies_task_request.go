package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowAgenciesTaskRequest struct {
}

func (o ShowAgenciesTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgenciesTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowAgenciesTaskRequest", string(data)}, " ")
}
