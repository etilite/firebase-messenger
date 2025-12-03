package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/etilite/firebase-messenger/internal/app"
	"github.com/etilite/firebase-messenger/internal/model"
	"google.golang.org/api/option"
)

type client interface {
	SendEachForMulticast(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error)
}

type MulticastSender struct {
	client client
}

func New(ctx context.Context, cfg app.Config) (*MulticastSender, error) {
	firebaseApp, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(cfg.FirebaseCredentials))
	if err != nil {
		return nil, err
	}

	firebaseClient, err := firebaseApp.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return &MulticastSender{
		client: firebaseClient,
	}, nil
}

func (s *MulticastSender) Push(ctx context.Context, req model.SendRequest) (model.SendResponse, error) {
	message := &messaging.MulticastMessage{
		Tokens: req.Tokens,
		Notification: &messaging.Notification{
			Title: req.Notification.Title,
			Body:  req.Notification.Body,
		},
		Data: req.Data,
	}

	resp, err := s.client.SendEachForMulticast(ctx, message)
	if err != nil {
		return model.SendResponse{}, nil
	}

	responses := make([]model.TokenSendResult, 0, len(resp.Responses))
	for _, r := range resp.Responses {
		var sendErr *model.SendError
		if r.Error != nil {
			sendErr = &model.SendError{
				// firebase-go has no any codes in errors.
				Code:    "",
				Message: r.Error.Error(),
			}
		}

		responses = append(responses, model.TokenSendResult{
			Success:   r.Success,
			MessageID: r.MessageID,
			Error:     sendErr,
		})
	}

	return model.SendResponse{
		SuccessCount: resp.SuccessCount,
		FailureCount: resp.FailureCount,
		Responses:    responses,
	}, nil
}
