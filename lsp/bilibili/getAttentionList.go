package bilibili

import (
	"time"

	"github.com/rock8526652/DDBOT/proxy_pool"
	"github.com/rock8526652/DDBOT/requests"
	"github.com/rock8526652/DDBOT/utils"
)

const (
	PathGetAttentionList = "/feed/v1/feed/get_attention_list"
)

func GetAttentionList() (*GetAttentionListResponse, error) {
	if !IsVerifyGiven() {
		return nil, ErrVerifyRequired
	}
	st := time.Now()
	defer func() {
		ed := time.Now()
		logger.WithField("FuncName", utils.FuncName()).Tracef("cost %v", ed.Sub(st))
	}()
	url := BPath(PathGetAttentionList)
	var opts []requests.Option
	opts = append(opts,
		requests.ProxyOption(proxy_pool.PreferNone),
		AddUAOption(),
		requests.TimeoutOption(time.Second*10),
		delete412ProxyOption,
	)
	opts = append(opts, GetVerifyOption()...)
	getAttentionListResp := new(GetAttentionListResponse)
	err := requests.Get(url, map[string]interface{}{
		"uid": accountUid.String(),
	}, getAttentionListResp, opts...)
	if err != nil {
		return nil, err
	}
	return getAttentionListResp, nil
}
