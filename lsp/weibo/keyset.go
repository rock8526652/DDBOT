package weibo

import localdb "github.com/rock8526652/DDBOT/lsp/buntdb"

type extraKeySet struct{}

func (*extraKeySet) UserInfoKey(keys ...interface{}) string {
	return localdb.WeiboUserInfoKey(keys...)
}

func (*extraKeySet) NewsInfoKey(keys ...interface{}) string {
	return localdb.WeiboNewsInfoKey(keys...)
}

func (*extraKeySet) MarkMblogIdKey(keys ...interface{}) string {
	return localdb.WeiboMarkMblogIdKey(keys...)
}
