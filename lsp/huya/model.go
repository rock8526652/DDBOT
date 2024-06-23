package huya

import (
	"sync"

	"github.com/rock8526652/DDBOT/lsp/concern_type"
	"github.com/rock8526652/DDBOT/lsp/mmsg"
	"github.com/rock8526652/DDBOT/lsp/template"
	localutils "github.com/rock8526652/DDBOT/utils"
	"github.com/sirupsen/logrus"
)

type LiveInfo struct {
	RoomId   string `json:"room_id"`
	RoomUrl  string `json:"room_url"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	RoomName string `json:"room_name"`
	IsLiving bool   `json:"living"`

	once              sync.Once
	msgCache          *mmsg.MSG
	liveStatusChanged bool
	liveTitleChanged  bool
}

func (m *LiveInfo) TitleChanged() bool {
	return m.liveTitleChanged
}

func (m *LiveInfo) IsLive() bool {
	return true
}

func (m *LiveInfo) Living() bool {
	return m.IsLiving
}

func (m *LiveInfo) LiveStatusChanged() bool {
	return m.liveStatusChanged
}

func (m *LiveInfo) GetUid() interface{} {
	return m.RoomId
}

func (m *LiveInfo) GetName() string {
	if m == nil {
		return ""
	}
	return m.Name
}

func (m *LiveInfo) Type() concern_type.Type {
	return Live
}

func (m *LiveInfo) ToString() string {
	if m == nil {
		return ""
	}
	bin, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(bin)
}

func (m *LiveInfo) Logger() *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"Site":   Site,
		"Name":   m.Name,
		"RoomId": m.RoomId,
		"Title":  m.RoomName,
		"Living": m.IsLiving,
	})
}

func (m *LiveInfo) Site() string {
	return Site
}

func (m *LiveInfo) GetMSG() *mmsg.MSG {
	m.once.Do(func() {
		var data = map[string]interface{}{
			"title":  m.RoomName,
			"name":   m.Name,
			"url":    m.RoomUrl,
			"cover":  m.Avatar,
			"living": m.Living(),
		}
		var err error
		m.msgCache, err = template.LoadAndExec("notify.group.huya.live.tmpl", data)
		if err != nil {
			logger.Errorf("huya: LiveInfo LoadAndExec error %v", err)
		}
		return
	})
	return m.msgCache
}

type ConcernLiveNotify struct {
	*LiveInfo
	GroupCode int64 `json:"group_code"`
}

func (notify *ConcernLiveNotify) GetGroupCode() int64 {
	return notify.GroupCode
}

func (notify *ConcernLiveNotify) ToMessage() (m *mmsg.MSG) {
	return notify.LiveInfo.GetMSG()
}

func (notify *ConcernLiveNotify) Logger() *logrus.Entry {
	if notify == nil {
		return logger
	}
	return notify.LiveInfo.Logger().WithFields(localutils.GroupLogFields(notify.GroupCode))
}

func NewConcernLiveNotify(groupCode int64, l *LiveInfo) *ConcernLiveNotify {
	if l == nil {
		return nil
	}
	return &ConcernLiveNotify{
		l,
		groupCode,
	}
}
