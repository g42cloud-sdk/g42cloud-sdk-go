package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BulkCreateAndDeleteVaultTagsReq struct {
	Tags *[]Tag `json:"tags,omitempty"`

	SysTags *[]SysTag `json:"sys_tags,omitempty"`

	Action BulkCreateAndDeleteVaultTagsReqAction `json:"action"`
}

func (o BulkCreateAndDeleteVaultTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BulkCreateAndDeleteVaultTagsReq struct{}"
	}

	return strings.Join([]string{"BulkCreateAndDeleteVaultTagsReq", string(data)}, " ")
}

type BulkCreateAndDeleteVaultTagsReqAction struct {
	value string
}

type BulkCreateAndDeleteVaultTagsReqActionEnum struct {
	CREATE BulkCreateAndDeleteVaultTagsReqAction
	DELETE BulkCreateAndDeleteVaultTagsReqAction
}

func GetBulkCreateAndDeleteVaultTagsReqActionEnum() BulkCreateAndDeleteVaultTagsReqActionEnum {
	return BulkCreateAndDeleteVaultTagsReqActionEnum{
		CREATE: BulkCreateAndDeleteVaultTagsReqAction{
			value: "create",
		},
		DELETE: BulkCreateAndDeleteVaultTagsReqAction{
			value: " delete",
		},
	}
}

func (c BulkCreateAndDeleteVaultTagsReqAction) Value() string {
	return c.value
}

func (c BulkCreateAndDeleteVaultTagsReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BulkCreateAndDeleteVaultTagsReqAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
