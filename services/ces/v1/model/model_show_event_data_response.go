package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowEventDataResponse struct {
	Datapoints     *[]EventDataInfo `json:"datapoints,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowEventDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventDataResponse struct{}"
	}

	return strings.Join([]string{"ShowEventDataResponse", string(data)}, " ")
}
