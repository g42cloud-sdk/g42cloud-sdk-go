package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteIds struct {
	Ids []string `json:"ids"`
}

func (o DeleteIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIds struct{}"
	}

	return strings.Join([]string{"DeleteIds", string(data)}, " ")
}
