package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListServerInterfacesRequest struct {
	ServerId string `json:"server_id"`
}

func (o ListServerInterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerInterfacesRequest struct{}"
	}

	return strings.Join([]string{"ListServerInterfacesRequest", string(data)}, " ")
}
