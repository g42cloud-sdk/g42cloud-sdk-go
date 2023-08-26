package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateRetryRemuxTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRetryRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRetryRemuxTaskResponse", string(data)}, " ")
}
