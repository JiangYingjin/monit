package response

import "my-server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
