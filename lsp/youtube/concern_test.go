package youtube

import (
	"context"
	"testing"
	"time"

	"github.com/rock8526652/DDBOT/internal/test"
	"github.com/rock8526652/DDBOT/lsp/concern"
	"github.com/stretchr/testify/assert"
)

func TestConcern(t *testing.T) {
	test.InitBuntdb(t)
	defer test.CloseBuntdb(t)

	testEventChan := make(chan concern.Event, 16)
	testNotifyChan := make(chan concern.Notify)

	c := NewConcern(testNotifyChan)

	assert.NotNil(t, c.GetStateManager())

	c.StateManager.UseNotifyGeneratorFunc(c.notifyGenerator())
	c.StateManager.UseFreshFunc(func(ctx context.Context, eventChan chan<- concern.Event) {
		for {
			select {
			case e := <-testEventChan:
				if e != nil {
					eventChan <- e
				}
			case <-ctx.Done():
				return
			}
		}
	})

	assert.Nil(t, c.StateManager.Start())
	defer c.Stop()
	defer close(testEventChan)

	_, err := c.ParseId(test.NAME1)
	assert.Nil(t, err)

	err = c.StateManager.AddInfo(&Info{
		UserInfo: UserInfo{
			ChannelId:   test.NAME1,
			ChannelName: test.NAME1,
		},
	})
	assert.Nil(t, err)

	_, err = c.StateManager.AddGroupConcern(test.G1, test.NAME1, Live)
	assert.Nil(t, err)

	_, err = c.StateManager.AddGroupConcern(test.G2, test.NAME1, Live)
	assert.Nil(t, err)

	identityInfo, err := c.Get(test.NAME1)
	assert.Nil(t, err)
	assert.EqualValues(t, test.NAME1, identityInfo.GetUid())

	assert.NotNil(t, c.GetGroupConcernConfig(test.G1, test.NAME1))

	testEventChan <- &VideoInfo{
		UserInfo: UserInfo{
			ChannelId:   test.NAME1,
			ChannelName: test.NAME1,
		},
		VideoType:         VideoType_Live,
		VideoStatus:       VideoStatus_Living,
		liveStatusChanged: true,
	}

	time.Sleep(time.Millisecond * 500)

	var g1 = false
	var g2 = false

	for i := 0; i < 2; i++ {
		select {
		case notify := <-testNotifyChan:
			if notify.GetGroupCode() == test.G1 {
				g1 = true
				assert.Equal(t, test.G1, notify.GetGroupCode())
				assert.Equal(t, test.NAME1, notify.GetUid())
			}
			if notify.GetGroupCode() == test.G2 {
				g2 = true
				assert.Equal(t, test.G2, notify.GetGroupCode())
				assert.Equal(t, test.NAME1, notify.GetUid())
			}
		case <-time.After(time.Second):
			assert.Fail(t, "no notify received")
		}
	}

	assert.True(t, g1)
	assert.True(t, g2)

	select {
	case <-testNotifyChan:
		assert.Fail(t, "should no notify received")
	case <-time.After(time.Second):

	}

	_, err = c.Remove(nil, test.G1, test.NAME1, Live)
	assert.Nil(t, err)
	_, err = c.Remove(nil, test.G2, test.NAME1, Live)
	assert.Nil(t, err)
}
