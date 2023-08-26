package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreatePostgresqlDatabaseResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostgresqlDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseResponse", string(data)}, " ")
}
