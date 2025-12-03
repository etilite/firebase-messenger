package firebase

import (
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/etilite/firebase-messenger/internal/model"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMulticastSender_Push(t *testing.T) {
	t.Parallel()

	// TODO: error cases, empty results etc
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		mc := minimock.NewController(t)
		clientMock := NewClientMock(mc)

		clientMock.SendEachForMulticastMock.Expect(t.Context(), &messaging.MulticastMessage{
			Tokens: []string{
				"token_1",
				"token_2",
			},
			Notification: &messaging.Notification{
				Title: "Hello!",
				Body:  "This is a multicast test",
			},
			Data: map[string]string{
				"foo": "bar",
			},
		}).Return(&messaging.BatchResponse{
			SuccessCount: 1,
			FailureCount: 1,
			Responses: []*messaging.SendResponse{{
				Success:   true,
				MessageID: "some-id1",
				Error:     nil,
			},
				{
					Success:   false,
					MessageID: "",
					Error:     assert.AnError,
				}},
		}, nil)

		sender := &MulticastSender{
			client: clientMock,
		}
		req := model.SendRequest{
			Tokens: []string{
				"token_1",
				"token_2",
			},
			Notification: model.Notification{
				Title: "Hello!",
				Body:  "This is a multicast test",
			},
			Data: map[string]string{
				"foo": "bar",
			},
		}

		want := model.SendResponse{
			SuccessCount: 1,
			FailureCount: 1,
			Responses: []model.TokenSendResult{
				{
					Success:   true,
					MessageID: "some-id1",
					Error:     nil,
				},
				{
					Success:   false,
					MessageID: "",
					Error: &model.SendError{
						Message: assert.AnError.Error(),
					},
				},
			},
		}
		got, err := sender.Push(t.Context(), req)
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}
