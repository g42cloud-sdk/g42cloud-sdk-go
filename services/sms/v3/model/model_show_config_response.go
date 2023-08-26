package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigResponse struct {
	Config map[string]string `json:"config,omitempty"`

	Regions        *[]map[string]string `json:"regions,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigResponse", string(data)}, " ")
}
