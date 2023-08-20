package unit

import (
	"context"
	"database/sql"
	"testing"
	"time"

	domainprofile "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/domain-profile"
	domainmock "github.com/DueIt-Jasanya-Aturuang/spongebob/domain/mocks"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/db/mocks"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestProfileUsecaseGetById(t *testing.T) {
	// log.Logger = log.Output(zerolog.Nop())
	profileRepoMock := &domainmock.FakeProfileRepo{}
	sqlMock := &mocks.FakeSQL{}
	profileUsecase := usecase.NewProfileUsecaseImpl(profileRepoMock, sqlMock, 5*time.Second)

	profileMockData := domainprofile.Profile{}
	profileMockData.UserId = "userid1"
	profileMockData = profileMockData.DefaultValue()

	profileRepoMock.GetProfileById(context.Background(), sqlMock.SqlDB(), profileMockData.UserId)
	profileRepoMock.GetProfileByIdReturns(&profileMockData, nil)
	ctxMock, dbMock, idMock := profileRepoMock.GetProfileByIdArgsForCall(0)
	assert.Equal(t, 1, profileRepoMock.GetProfileByIdCallCount())
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, sqlMock.SqlDB(), dbMock)
	assert.Equal(t, profileMockData.UserId, idMock)

	profile, err := profileUsecase.GetProfileById(context.Background(), profileMockData.UserId)
	t.Log(profile)
	assert.NotNil(t, profile)
	assert.NoError(t, err)
}

func TestProfileUsecaseGetByUserId(t *testing.T) {
	// log.Logger = log.Output(zerolog.Nop())
	profileRepoMock := &domainmock.FakeProfileRepo{}
	sqlMock := &mocks.FakeSQL{}
	profileUsecase := usecase.NewProfileUsecaseImpl(profileRepoMock, sqlMock, 5*time.Second)

	profileMockData := domainprofile.Profile{}
	profileMockData.UserId = "userid1"
	profileMockData = profileMockData.DefaultValue()

	profileRepoMock.GetProfileById(context.Background(), sqlMock.SqlDB(), profileMockData.UserId)
	profileRepoMock.GetProfileByIdReturns(nil, sql.ErrNoRows)
	ctxMock, dbMock, idMock := profileRepoMock.GetProfileByIdArgsForCall(0)
	assert.Equal(t, 1, profileRepoMock.GetProfileByIdCallCount())
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, sqlMock.SqlDB(), dbMock)
	assert.Equal(t, profileMockData.UserId, idMock)

	profileRepoMock.GetProfileByUserId(context.Background(), sqlMock.SqlDB(), profileMockData.UserId)
	profileRepoMock.GetProfileByUserIdReturns(&profileMockData, nil)
	ctxMock, dbMock, idMock = profileRepoMock.GetProfileByUserIdArgsForCall(0)
	assert.Equal(t, 1, profileRepoMock.GetProfileByUserIdCallCount())
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, sqlMock.SqlDB(), dbMock)
	assert.Equal(t, profileMockData.UserId, idMock)

	profile, err := profileUsecase.GetProfileById(context.Background(), profileMockData.UserId)
	t.Log(profile)
	assert.NotNil(t, profile)
	assert.NoError(t, err)
}

func TestProfileUsecaseGetByIdWithStore(t *testing.T) {
	// log.Logger = log.Output(zerolog.Nop())
	profileRepoMock := &domainmock.FakeProfileRepo{}
	sqlMock := &mocks.FakeSQL{}
	profileUsecase := usecase.NewProfileUsecaseImpl(profileRepoMock, sqlMock, 5*time.Second)

	profileMockData := domainprofile.Profile{}
	profileMockData.UserId = "userid1"
	profileMockData = profileMockData.DefaultValue()

	profileRepoMock.GetProfileById(context.Background(), sqlMock.SqlDB(), "userid2")
	profileRepoMock.GetProfileByIdReturns(nil, sql.ErrNoRows)
	ctxMock, dbMock, idMock := profileRepoMock.GetProfileByIdArgsForCall(0)
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, sqlMock.SqlDB(), dbMock)
	assert.Equal(t, "userid2", idMock)

	profileRepoMock.GetProfileByUserId(context.Background(), sqlMock.SqlDB(), "userid2")
	profileRepoMock.GetProfileByUserIdReturnsOnCall(1, &profileMockData, nil)
	profileMockData.UserId = "userid2"
	profileMockData = profileMockData.DefaultValue()
	ctxMock, dbMock, idMock = profileRepoMock.GetProfileByUserIdArgsForCall(0)
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, sqlMock.SqlDB(), dbMock)
	assert.Equal(t, "userid2", idMock)

	profileRepoMock.StoreProfile(context.Background(), &sql.Tx{}, profileMockData)
	profileRepoMock.StoreProfileReturnsOnCall(1, profileMockData, nil)
	ctxMock, txMock, profileMock := profileRepoMock.StoreProfileArgsForCall(0)
	assert.Equal(t, context.Background(), ctxMock)
	assert.Equal(t, &sql.Tx{}, txMock)
	assert.Equal(t, profileMockData, profileMock)

	profile, err := profileUsecase.GetProfileById(context.Background(), "userid2")
	t.Log(profile)
	assert.NoError(t, err)
	assert.NotNil(t, profile)
}
