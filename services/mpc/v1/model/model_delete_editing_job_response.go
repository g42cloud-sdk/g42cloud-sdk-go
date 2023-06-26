package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteEditingJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEditingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEditingJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteEditingJobResponse", string(data)}, " ")
}
