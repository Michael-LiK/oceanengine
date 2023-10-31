package rta

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rta"
)

// Get 获取可用的RTA策略
func Get(clt *core.SDKClient, accessToken string, req *rta.GetRequest) ([]rta.RtaInfo, error) {
	var resp rta.GetResponse
	if err := clt.Get("2/tools/rta/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
