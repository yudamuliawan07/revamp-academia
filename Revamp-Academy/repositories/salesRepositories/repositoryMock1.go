package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepositoryMock struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewRepositoryMock(dbHandler *sql.DB) *RepositoryMock {
	return &RepositoryMock{
		dbHandler: dbHandler,
	}
}

func (rm RepositoryMock) GetMockup(ctx *gin.Context, nama string) (*dbContext.CreateprogramEntityParams, *models.ResponseError) {

	store := dbContext.New(rm.dbHandler)
	mockup, err := store.GetProgramEntity(ctx, nama)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}

func (rm RepositoryMock) GetListProgram(ctx *gin.Context, nama string) ([]*dbContext.CreateprogramEntityParams, *models.ResponseError) {

	store := dbContext.New(rm.dbHandler)
	program_entity, err := store.Listprogram_entity(ctx, nama)

	listProgramEntity := make([]*dbContext.CreateprogramEntityParams, 0)

	for _, v := range program_entity {
		sales := &dbContext.CreateprogramEntityParams{
			ProgTitle:        v.ProgTitle,
			ProgHeadline:     v.ProgHeadline,
			ProgLearningType: v.ProgLearningType,
			ProgImage:        v.ProgImage,
			ProgPrice:        v.ProgPrice,
			ProgDuration:     v.ProgDuration,
		}
		listProgramEntity = append(listProgramEntity, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramEntity, nil
}
