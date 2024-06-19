package request

import (
	"my-server/model/common/request"
	"my-server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
