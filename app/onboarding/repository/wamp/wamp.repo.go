package wamp_repo

import (
	"context"
	"web_server/domain"

	"github.com/gammazero/nexus/v3/client"
)

type wampRepo struct {
	client *client.Client
}

func NewSessionRepository(client *client.Client) domain.SessionRepository {
	return &wampRepo{client: client}
}

func (w *wampRepo) GetSessionMeta(ctx context.Context, authId string) {
	// TODO collect session in
	// any login with that session, change its role
	// w.client.Publish()
	panic("not implemented") // TODO: Implement
}
