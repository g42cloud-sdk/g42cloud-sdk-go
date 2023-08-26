package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteThumbnailsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteThumbnailsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThumbnailsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteThumbnailsTaskResponse", string(data)}, " ")
}
