package youtube

import "github.com/rock8526652/DDBOT/lsp/buntdb"

type KeySet struct {
}

func (k *KeySet) GroupAtAllMarkKey(keys ...interface{}) string {
	return buntdb.YoutubeGroupAtAllMarkKey(keys...)
}

func (k *KeySet) GroupConcernConfigKey(keys ...interface{}) string {
	return buntdb.YoutubeGroupConcernConfigKey(keys...)
}

func (k *KeySet) GroupConcernStateKey(keys ...interface{}) string {
	return buntdb.YoutubeGroupConcernStateKey(keys...)
}

func (k *KeySet) FreshKey(keys ...interface{}) string {
	return buntdb.YoutubeFreshKey(keys...)
}

func (k *KeySet) ParseGroupConcernStateKey(key string) (int64, interface{}, error) {
	return buntdb.ParseConcernStateKeyWithString(key)
}

func NewKeySet() *KeySet {
	return new(KeySet)
}

type extraKey struct {
}

func (e *extraKey) UserInfoKey(keys ...interface{}) string {
	return buntdb.YoutubeUserInfoKey(keys...)
}

func (e *extraKey) InfoKey(keys ...interface{}) string {
	return buntdb.YoutubeInfoKey(keys...)
}

func (e *extraKey) VideoKey(keys ...interface{}) string {
	return buntdb.YoutubeVideoKey(keys...)
}

func NewExtraKey() *extraKey {
	return &extraKey{}
}
