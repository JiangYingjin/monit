package request

import (
	"my-server/model/common/request"
	"my-server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
