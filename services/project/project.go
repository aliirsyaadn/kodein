package project

import (
	"context"

	"github.com/google/uuid"

	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/internal/response"
	"github.com/aliirsyaadn/kodein/model"
)

const projectTag = "ProjectService"

type Service interface {
	GetProjectsByMemberID(ctx context.Context, memberID string) (res entity.ListProjectResponse, err error)
	GetProjectByID(ctx context.Context, id string) (res entity.GetProjectResponse, err error)
	CreateProject(ctx context.Context, arg entity.CreateProjectRequest) (res entity.CreateProjectResponse, err error)
	UpdateProject(ctx context.Context, arg entity.UpdateProjectRequest, id string) (res entity.UpdateProjectResponse, err error)
	DeleteProject(ctx context.Context, id string) (res entity.DeleteProjectResponse, err error)
}

type Repository interface {
	GetProjectsByMemberID(ctx context.Context, memberID uuid.UUID) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id uuid.UUID) (model.Project, error)
	InsertProject(ctx context.Context, arg model.InsertProjectParams) (model.Project, error)
	UpdateProject(ctx context.Context, arg model.UpdateProjectParams) (model.Project, error)
	DeleteProject(ctx context.Context, id uuid.UUID) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetProjectsByMemberID(ctx context.Context, memberID string) (res entity.ListProjectResponse, err error) {
	memberIDParsed, err := uuid.Parse(memberID)
	if err != nil {
		log.ErrorDetail(projectTag, "error parse uuid: %v", err)
		return
	}

	data, err := s.r.GetProjectsByMemberID(ctx, memberIDParsed)
	if err != nil {
		log.ErrorDetail(projectTag, "error GetAllProject from DB: %v", err)
		return
	}

	res = entity.ListProjectResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) GetProjectByID(ctx context.Context, id string) (res entity.GetProjectResponse, err error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(projectTag, "error parse uuid: %v", err)
		return
	}

	data, err := s.r.GetProjectByID(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(projectTag, "error GetProjectByID from DB: %v", err)
		return
	}

	res = entity.GetProjectResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) CreateProject(ctx context.Context, arg entity.CreateProjectRequest) (res entity.CreateProjectResponse, err error) {
	dataInsert := model.InsertProjectParams{
		MemberID:    arg.Data.MemberID,
		Name:        arg.Data.Name,
		Description: arg.Data.Description,
		Technology:  arg.Data.Technology,
		Url:         arg.Data.Url,
	}

	data, err := s.r.InsertProject(ctx, dataInsert)
	if err != nil {
		log.ErrorDetail(projectTag, "error InsertProject from DB: %v", err)
		return
	}

	res = entity.CreateProjectResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) UpdateProject(ctx context.Context, arg entity.UpdateProjectRequest, id string) (res entity.UpdateProjectResponse, err error) {

	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(projectTag, "error parse uuid: %v", err)
		return
	}

	dataUpdate := model.UpdateProjectParams{
		ID: idParsed,
		Name:        arg.Data.Name.String,
		Description: arg.Data.Description.NullString,
		Technology:  arg.Data.Technology.String,
		Url:         arg.Data.Url.String,
	}

	data, err := s.r.UpdateProject(ctx, dataUpdate)
	if err != nil {
		log.ErrorDetail(projectTag, "error UpdateProject from DB: %v", err)
		return
	}

	res = entity.UpdateProjectResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) DeleteProject(ctx context.Context, id string) (res entity.DeleteProjectResponse, err error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(projectTag, "error parse uuid: %v", err)
		return
	}

	err = s.r.DeleteProject(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(projectTag, "error DeleteProject from DB: %v", err)
		return
	}

	res = entity.DeleteProjectResponse{
		ID:       id,
		Response: response.OK,
	}

	return res, nil
}
