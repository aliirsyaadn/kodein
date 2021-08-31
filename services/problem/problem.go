package problem

import (
	"context"

	"github.com/google/uuid"

	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/internal/response"
	"github.com/aliirsyaadn/kodein/model"
)

const problemTag = "ProblemService"

type Service interface {
	GetProblems(ctx context.Context) (res entity.ListProblemResponse, err error)
	GetProblemByID(ctx context.Context, id string) (res entity.GetProblemResponse, err error)
	CreateProblem(ctx context.Context, arg entity.CreateProblemRequest) (res entity.CreateProblemResponse, err error)
	UpdateProblem(ctx context.Context, arg entity.UpdateProblemRequest, id string) (res entity.UpdateProblemResponse, err error)
	DeleteProblem(ctx context.Context, id string) (res entity.DeleteProblemResponse, err error)
}

type Repository interface {
	GetProblems(ctx context.Context) ([]model.Problem, error)
	GetProblemByID(ctx context.Context, id uuid.UUID) (model.Problem, error)
	InsertProblem(ctx context.Context, arg model.InsertProblemParams) (model.Problem, error)
	UpdateProblem(ctx context.Context, arg model.UpdateProblemParams) (model.Problem, error)
	DeleteProblem(ctx context.Context, id uuid.UUID) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetProblems(ctx context.Context) (res entity.ListProblemResponse, err error) {

	data, err := s.r.GetProblems(ctx)
	if err != nil {
		log.ErrorDetail(problemTag, "error GetAllProblem from DB: %v", err)
		return
	}

	res = entity.ListProblemResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) GetProblemByID(ctx context.Context, id string) (res entity.GetProblemResponse, err error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(problemTag, "error parse uuid: %v", err)
		return
	}

	data, err := s.r.GetProblemByID(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(problemTag, "error GetProblemByID from DB: %v", err)
		return
	}

	res = entity.GetProblemResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) CreateProblem(ctx context.Context, arg entity.CreateProblemRequest) (res entity.CreateProblemResponse, err error) {
	dataInsert := model.InsertProblemParams{
		Name:        arg.Data.Name,
		Description: arg.Data.Description,
		Category: arg.Data.Category,
		Difficulty: arg.Data.Difficulty,
		GraderCode: arg.Data.GraderCode,
	}

	data, err := s.r.InsertProblem(ctx, dataInsert)
	if err != nil {
		log.ErrorDetail(problemTag, "error InsertProblem from DB: %v", err)
		return
	}

	res = entity.CreateProblemResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) UpdateProblem(ctx context.Context, arg entity.UpdateProblemRequest, id string) (res entity.UpdateProblemResponse, err error) {

	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(problemTag, "error parse uuid: %v", err)
		return
	}

	dataUpdate := model.UpdateProblemParams{
		ID: idParsed,
		Name:        arg.Data.Name.String,
		Description: arg.Data.Description.String,
		Category: arg.Data.Category.String,
		Difficulty: model.DifficultyTypeInsane,
		GraderCode: arg.Data.GraderCode.String,
	}

	data, err := s.r.UpdateProblem(ctx, dataUpdate)
	if err != nil {
		log.ErrorDetail(problemTag, "error UpdateProblem from DB: %v", err)
		return
	}

	res = entity.UpdateProblemResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) DeleteProblem(ctx context.Context, id string) (res entity.DeleteProblemResponse, err error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(problemTag, "error parse uuid: %v", err)
		return
	}

	err = s.r.DeleteProblem(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(problemTag, "error DeleteProblem from DB: %v", err)
		return
	}

	res = entity.DeleteProblemResponse{
		ID:       id,
		Response: response.OK,
	}

	return res, nil
}
