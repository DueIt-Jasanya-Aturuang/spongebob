package profileConfig_usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/repository"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/util"
)

func (p *ProfileConfigUsecaseImpl) SchedulerDailyNotify(ctx context.Context, minuteSecond string, day string) error {
	profileConfigs, err := p.profileCfgRepo.GetBySchedulerDailyNotify(ctx, minuteSecond, day)
	if err != nil {
		return err
	}

	if len(*profileConfigs) < 1 {
		return nil
	}

	notificationHelper, err := p.notifRepo.GetNotifHelperByName(ctx, util.DailyNotify)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		notificationHelper.Message = util.DefaultMessageDailyNotify
		notificationHelper.Title = util.DefaultTitleDailyNotify
		notificationHelper.Icon = util.DefaultIconDailyNotify
	}

	for _, profileConfig := range *profileConfigs {
		formatConfigValue := map[string]any{}
		err = json.Unmarshal([]byte(profileConfig.ConfigValue), &formatConfigValue)
		if err != nil {
			return err
		}

		err = p.profileCfgRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
			err = p.notifRepo.Create(ctx, &repository.Notification{
				ID:           util.NewUlid,
				ProfileID:    profileConfig.ProfileID,
				UserConfigID: profileConfig.ID,
				Message:      notificationHelper.Message,
				Status:       "unread",
				Title:        notificationHelper.Title,
				Icon:         notificationHelper.Icon,
				AuditInfo: repository.AuditInfo{
					CreatedAt: time.Now().Unix(),
					CreatedBy: profileConfig.ProfileID,
					UpdatedAt: time.Now().Unix(),
				},
			})
			if err != nil {
				return err
			}

			// 	push to fcm

			return nil
		})

		if err != nil {
			return err
		}

	}

	return nil
}

func (p *ProfileConfigUsecaseImpl) SchedulerMonthlyPeriode(ctx context.Context, tgl int, id *string) (*string, error) {
	var profileConfigs *[]repository.ProfileConfig
	var err error

	if id != nil {
		profileConfigs, err = p.profileCfgRepo.GetBySchedulerMonthlyPeriode(ctx, tgl, *id)
	} else {
		profileConfigs, err = p.profileCfgRepo.GetBySchedulerMonthlyPeriode(ctx, tgl, "")
	}
	if err != nil {
		return nil, err
	}

	if len(*profileConfigs) < 1 {
		return nil, nil
	}

	notificationHelper, err := p.notifRepo.GetNotifHelperByName(ctx, util.MonthlyPeriode)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		notificationHelper.Message = util.DefaultMessageMonthlyPeriod
		notificationHelper.Title = util.DefaultTitleMonthlyPeriod
		notificationHelper.Icon = util.DefaultIconMonthlyPeriod
	}

	for _, profileConfig := range *profileConfigs {
		formatConfigValue := map[string]any{}
		err = json.Unmarshal([]byte(profileConfig.ConfigValue), &formatConfigValue)
		if err != nil {
			return nil, err
		}

		err = p.profileCfgRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
			err = p.notifRepo.Create(ctx, &repository.Notification{
				ID:           ulid.Make().String(),
				ProfileID:    profileConfig.ProfileID,
				UserConfigID: profileConfig.ID,
				Message:      notificationHelper.Message,
				Status:       "unread",
				Title:        notificationHelper.Title,
				Icon:         notificationHelper.Icon,
				AuditInfo: repository.AuditInfo{
					CreatedAt: time.Now().Unix(),
					CreatedBy: profileConfig.ProfileID,
					UpdatedAt: time.Now().Unix(),
				},
			})
			if err != nil {
				return err
			}

			// 	push to fmc

			return nil
		})

		if err != nil {
			return nil, err
		}

	}

	cursor := (*profileConfigs)[len(*profileConfigs)-1].ID
	return &cursor, nil
}
