package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/delivery/restapi/response"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils/message"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
)

type ProfileUsecaseImpl struct {
	profileRepo repository.ProfileRepo
	userRepo    repository.UserRepo
	ctxTimeout  time.Duration
}

func NewProfileUsecaseImpl(
	profileRepo repository.ProfileRepo,
	userRepo repository.UserRepo,
	timeout time.Duration,
) usecase.ProfileUsecase {
	return &ProfileUsecaseImpl{
		profileRepo: profileRepo,
		userRepo:    userRepo,
		ctxTimeout:  timeout,
	}
}

func (u *ProfileUsecaseImpl) GetProfileByID(c context.Context, req *dto.GetProfileReq) (resp *dto.ProfileResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err = u.profileRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		u.profileRepo.CloseConn()
	}()

	res, err := u.profileRepo.GetProfileByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	resp = res.ToResp()
	return resp, nil
}

func (u *ProfileUsecaseImpl) StoreProfile(c context.Context, req *dto.StoreProfileReq) (profileResp *dto.ProfileResp, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	err = u.profileRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		u.profileRepo.CloseConn()
	}()

	_, err = u.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrForbidden
		}
		return nil, err
	}
	var profile *model.Profile
	profile = profile.DefaultValue(req.UserID)

	err = u.profileRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		if errEndTx := u.profileRepo.EndTx(err); errEndTx != nil {
			err = errEndTx
			profile = nil
		}
	}()

	profileRes, err := u.profileRepo.StoreProfile(ctx, *profile)
	if err != nil {
		if errors.Is(err, model.ErrConflict) {
			return nil, response.Err409(map[string][]string{
				"profile": {
					message.ProfileIsAlvailable,
				},
			}, err)
		}
		return nil, err
	}

	profileResp = profileRes.ToResp()

	return profileResp, nil
}
