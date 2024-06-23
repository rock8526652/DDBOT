package youtube

import (
	"testing"

	"github.com/rock8526652/DDBOT/internal/test"
	"github.com/rock8526652/DDBOT/lsp/concern"
	"github.com/stretchr/testify/assert"
)

func newLiveInfo(channelId string, live bool, liveStatusChanged bool, liveTitleChanged bool) *ConcernNotify {
	li := &ConcernNotify{
		VideoInfo: &VideoInfo{
			UserInfo: UserInfo{
				ChannelId: channelId,
			},
			liveStatusChanged: liveStatusChanged,
			liveTitleChanged:  liveTitleChanged,
			VideoType:         VideoType_Live,
		},
	}
	if live {
		li.VideoStatus = VideoStatus_Living
	} else {
		li.VideoStatus = VideoStatus_Waiting
	}
	return li
}

func newNewsInfo(channelId string) *ConcernNotify {
	li := &ConcernNotify{
		VideoInfo: &VideoInfo{
			UserInfo: UserInfo{
				ChannelId: channelId,
			},
			VideoType: VideoType_Video,
		},
	}
	return li
}

func TestNewGroupConcernConfig(t *testing.T) {
	g := NewGroupConcernConfig(new(concern.GroupConcernConfig))
	assert.NotNil(t, g)
}

func TestGroupConcernConfig_ShouldSendHook(t *testing.T) {
	var notify = []concern.Notify{
		// 直播预告 推
		newLiveInfo(test.NAME1, false, false, false),
		// 直播预告 标题变了 推
		newLiveInfo(test.NAME1, false, false, true),
		// 可能吗
		newLiveInfo(test.NAME1, false, true, false),
		// 可能吗
		newLiveInfo(test.NAME1, false, true, true),
		// 直播状态 什么也没变 不推
		newLiveInfo(test.NAME1, true, false, false),
		// 直播状态 改了标题 检查配置
		newLiveInfo(test.NAME1, true, false, true),
		// 开播了 推
		newLiveInfo(test.NAME1, true, true, false),
		// 开播了改了标题 推
		newLiveInfo(test.NAME1, true, true, true),
	}

	var testCase = []*GroupConcernConfig{
		{
			IConfig: &concern.GroupConcernConfig{},
		},
		{
			IConfig: &concern.GroupConcernConfig{
				GroupConcernNotify: concern.GroupConcernNotifyConfig{
					TitleChangeNotify: Live,
				},
			},
		},
	}
	var expected = [][]bool{
		{
			true, true, true, true,
			false, false, true, true,
		},
		{
			true, true, true, true,
			false, true, true, true,
		},
	}
	assert.Equal(t, len(expected), len(testCase))
	for index1, g := range testCase {
		assert.Equal(t, len(expected[index1]), len(notify))
		for index2, liveInfo := range notify {
			result := g.ShouldSendHook(liveInfo)
			assert.NotNil(t, result)
			assert.Equal(t, expected[index1][index2], result.Pass)
		}
	}
}

func TestGroupConcernConfig_AtBeforeHook(t *testing.T) {
	var notify = []concern.Notify{
		// 下播状态 什么也没变 不推
		newLiveInfo(test.NAME1, false, false, false),
		// 下播状态 标题变了 不推
		newLiveInfo(test.NAME1, false, false, true),
		// 下播了 检查配置
		newLiveInfo(test.NAME1, false, true, false),
		// 下播了 检查配置
		newLiveInfo(test.NAME1, false, true, true),
		// 直播状态 什么也没变 不推
		newLiveInfo(test.NAME1, true, false, false),
		// 直播状态 改了标题 检查配置
		newLiveInfo(test.NAME1, true, false, true),
		// 开播了 推
		newLiveInfo(test.NAME1, true, true, false),
		// 开播了改了标题 推
		newLiveInfo(test.NAME1, true, true, true),
	}
	var expected = []bool{
		false, false, false, false, false, false, true, true,
	}
	var config = &GroupConcernConfig{IConfig: &concern.GroupConcernConfig{}}
	for idx, n := range notify {
		hook := config.AtBeforeHook(n)
		assert.EqualValues(t, expected[idx], hook.Pass)
	}

}
