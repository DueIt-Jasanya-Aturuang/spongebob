package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/dtoconv"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helpers/format"
	"github.com/rs/zerolog/log"
)

type AccountUsecaseImpl struct {
	profileRepo repository.ProfileRepo
	userRepo    repository.UserRepo
	tsx         repository.SqlTransactionRepo
	minioClient repository.MinioRepo
	ctxTimeout  time.Duration
}

func NewAccountUsecaseImpl(
	profileRepo repository.ProfileRepo,
	userRepo repository.UserRepo,
	tsx repository.SqlTransactionRepo,
	minioClient repository.MinioRepo,
	ctxTimeout time.Duration,
) usecase.AccountUsecase {
	return &AccountUsecaseImpl{
		profileRepo: profileRepo,
		userRepo:    userRepo,
		tsx:         tsx,
		minioClient: minioClient,
		ctxTimeout:  ctxTimeout,
	}
}

func (u *AccountUsecaseImpl) AccountUpdate(c context.Context, req dto.UpdateAccountReq) (*dto.UserResp, *dto.ProfileResp, error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	profile, err := u.profileRepo.GetProfileByUserId(ctx, req.UserID)
	if err != nil {
		return nil, nil, err
	}

	user, err := u.userRepo.GetUserById(ctx, req.UserID)
	if err != nil {
		return nil, nil, err
	}

	oldImage := user.Image
	reqImageBool := req.Image != nil && req.Image.Size > 0
	delImageBool := !strings.Contains(oldImage, "default-male") && !strings.Contains(oldImage, "google")

	log.Info().Msgf("%v", reqImageBool)
	err = u.tsx.Transaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		log.Info().Msg("start tx")
		filename := u.minioClient.GenerateFileName(req.Image, "user-images/public/", "")
		if reqImageBool {
			user.Image = fmt.Sprintf("/%s/%s", config.MinIoBucket, filename)
		}

		profileConv, userConv := dtoconv.UpdateAccountToModel(req, profile.ProfileId, user.Image)

		profile, err = u.profileRepo.UpdateProfile(ctx, tx, profileConv)
		if err != nil {
			return err
		}

		user, err = u.userRepo.UpdateUser(ctx, tx, userConv)
		if err != nil {
			return err
		}

		if reqImageBool {
			err = u.minioClient.UploadFile(ctx, req.Image, filename, config.MinIoBucket)
			if err != nil {
				return err
			}

			if delImageBool {
				oldImageArr := strings.Split(oldImage, "/")
				filename = fmt.Sprintf("/%s/%s/%s", oldImageArr[2], oldImageArr[3], oldImageArr[4])
				err = u.minioClient.DeleteFile(ctx, filename, config.MinIoBucket)
			}
		}
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	userResp := user.ToResp(format.EmailFormat(user.Email))
	profileResp := profile.ToResp()
	return &userResp, &profileResp, nil
}
