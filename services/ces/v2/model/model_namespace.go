package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Namespace struct {
}

func (o Namespace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Namespace struct{}"
	}

	return strings.Join([]string{"Namespace", string(data)}, " ")
}
