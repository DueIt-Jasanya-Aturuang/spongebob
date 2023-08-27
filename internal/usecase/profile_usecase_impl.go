package usecase

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
)

type ProfileUsecaseImpl struct {
	profileRepo repository.ProfileRepo
	ctxTimeout  time.Duration
}

func NewProfileUsecaseImpl(
	profileRepo repository.ProfileRepo,
	timeout time.Duration,
) usecase.ProfileUsecase {
	return &ProfileUsecaseImpl{
		profileRepo: profileRepo,
		ctxTimeout:  timeout,
	}
}

func (u *ProfileUsecaseImpl) GetProfileByID(c context.Context, req dto.GetProfileReq) (resp *dto.ProfileResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	res, err := u.profileRepo.GetProfileByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	resp = res.ToResp()
	return resp, nil
}

func (u *ProfileUsecaseImpl) StoreProfile(c context.Context, req dto.StoreProfileReq) (profile *model.Profile, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	profile = profile.DefaultValue(req.UserID)

	profileRepoUOW := u.profileRepo.UoW()

	err = profileRepoUOW.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		if errEndTx := profileRepoUOW.EndTx(err); errEndTx != nil {
			err = errEndTx
			profile = nil
		}
	}()

	profileRes, err := u.profileRepo.StoreProfile(ctx, *profile)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("%v", profileRes)
	profile = &profileRes
	return profile, nil
}
