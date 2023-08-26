package unit

import (
	"context"
	"database/sql"
	"testing"
	"time"

	domainmock "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/mocks"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestProfileGetByIDUSECASE(t *testing.T) {
	profileRepoMock := &domainmock.FakeProfileRepo{}
	profileUsecase := usecase.NewProfileUsecaseImpl(profileRepoMock, 5*time.Second)

	profileMockData := model.Profile{}
	profileMockData = *profileMockData.DefaultValue("userid1")

	profileRepoMock.GetProfileByID(context.Background(), profileMockData.UserID)
	profileRepoMock.GetProfileByIDReturns(&profileMockData, nil)
	ctxMock, idMock := profileRepoMock.GetProfileByIDArgsForCall(0)
	assert.Equal(t, 1, profileRepoMock.GetProfileByIDCallCount())
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, profileMockData.UserID, idMock)

	profile, err := profileUsecase.GetProfileByID(context.Background(), profileMockData.UserID)
	t.Log(profile)
	assert.NotNil(t, profile)
	assert.NoError(t, err)
}

func TestProfileGetByUserIDUSECASE(t *testing.T) {
	profileRepoMock := &domainmock.FakeProfileRepo{}
	profileUsecase := usecase.NewProfileUsecaseImpl(profileRepoMock, 5*time.Second)

	profileMockData := model.Profile{}
	profileMockData = *profileMockData.DefaultValue("userid1")

	profileRepoMock.GetProfileByID(context.Background(), profileMockData.UserID)
	profileRepoMock.GetProfileByIDReturns(nil, sql.ErrNoRows)
	ctxMock, idMock := profileRepoMock.GetProfileByIDArgsForCall(0)
	assert.Equal(t, 1, profileRepoMock.GetProfileByIDCallCount())
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, profileMockData.UserID, idMock)

	profileRepoMock.GetProfileByUserID(context.Background(), profileMockData.UserID)
	profileRepoMock.GetProfileByUserIDReturns(&profileMockData, nil)
	ctxMock, idMock = profileRepoMock.GetProfileByUserIDArgsForCall(0)
	assert.Equal(t, 1, profileRepoMock.GetProfileByUserIDCallCount())
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, profileMockData.UserID, idMock)

	profile, err := profileUsecase.GetProfileByID(context.Background(), profileMockData.UserID)
	t.Log(profile)
	assert.NotNil(t, profile)
	assert.NoError(t, err)
}

func TestProfileGetByIDWithStoreUSECASE(t *testing.T) {
	profileRepoMock := &domainmock.FakeProfileRepo{}
	profileUsecase := usecase.NewProfileUsecaseImpl(profileRepoMock, 5*time.Second)

	profileMockData := model.Profile{}
	profileMockData = *profileMockData.DefaultValue("userid1")

	profileRepoMock.GetProfileByID(context.Background(), "userid2")
	profileRepoMock.GetProfileByIDReturns(nil, sql.ErrNoRows)
	ctxMock, idMock := profileRepoMock.GetProfileByIDArgsForCall(0)
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, "userid2", idMock)

	profileRepoMock.GetProfileByUserID(context.Background(), "userid2")
	profileRepoMock.GetProfileByUserIDReturnsOnCall(1, &profileMockData, nil)
	profileMockData = *profileMockData.DefaultValue("userid2")
	ctxMock, idMock = profileRepoMock.GetProfileByUserIDArgsForCall(0)
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, "userid2", idMock)

	profileRepoMock.StoreProfile(context.Background(), profileMockData)
	profileRepoMock.StoreProfileReturnsOnCall(1, profileMockData, nil)
	ctxMock, profileMock := profileRepoMock.StoreProfileArgsForCall(0)
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, profileMockData, profileMock)

	profile, err := profileUsecase.GetProfileByID(context.Background(), "userid2")
	t.Log(profile)
	assert.NoError(t, err)
	assert.NotNil(t, profile)
}
