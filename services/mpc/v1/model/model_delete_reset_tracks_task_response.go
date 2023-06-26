package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteResetTracksTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResetTracksTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResetTracksTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteResetTracksTaskResponse", string(data)}, " ")
}
