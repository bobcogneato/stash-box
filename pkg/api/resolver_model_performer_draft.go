package api

import (
	"context"

	"github.com/stashapp/stash-box/pkg/models"
)

type performerDraftResolver struct{ *Resolver }

func (r *performerDraftResolver) ID(ctx context.Context, obj *models.PerformerDraft) (*string, error) {
	if obj.ID != nil {
		val := obj.ID.String()
		return &val, nil
	}
	return nil, nil
}

func (r *performerDraftResolver) Image(ctx context.Context, obj *models.PerformerDraft) (*models.Image, error) {
	if obj.Image == nil {
		return nil, nil
	}

	fac := r.getRepoFactory(ctx)
	qb := fac.Image()
	return qb.Find(*obj.Image)
}
