package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreatePostgresqlDatabaseSchemaResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostgresqlDatabaseSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseSchemaResponse struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseSchemaResponse", string(data)}, " ")
}
