package alias

import (
	"context"

	"go.uber.org/zap"
)

type usecase struct {
	repo Repository
	log  *zap.SugaredLogger
}

func NewUseCase(repo Repository, logger *zap.SugaredLogger) Usecase {
	return &usecase{
		repo: repo,
		log:  logger.Named("alias.usecase"),
	}
}

func (uc *usecase) Save(ctx context.Context, alias *Alias) (*Alias, error) {
	err := uc.repo.Create(ctx, alias)
	if err != nil {
		uc.log.Errorw("Failed get by key",
			"key", alias.Key,
			"url", alias.URL,
			"err", err,
		)
		return nil, err
	}
	return alias, nil
}

func (uc *usecase) Get(ctx context.Context, key string) (*Alias, error) {
	alias, err := uc.repo.Get(ctx, key)
	if err != nil {
		uc.log.Errorw("Failed get by key",
			"key", key,
			"err", err,
		)
		return nil, err
	}
	return alias, nil
}

func (uc usecase) IncrPV(ctx context.Context, id int) error {
	err := uc.repo.IncrPV(ctx, id)
	if err != nil {
		uc.log.Errorw("increase pv error",
			"id", id,
			"err", err,
		)
	}
	return err
}
