package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTasksReq struct {
	Ids []string `json:"ids"`
}

func (o DeleteTasksReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTasksReq struct{}"
	}

	return strings.Join([]string{"DeleteTasksReq", string(data)}, " ")
}
