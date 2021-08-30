package member

import (
	"context"

	"github.com/google/uuid"

	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/internal/response"
	"github.com/aliirsyaadn/kodein/model"
)

const adminTag = "MemberService"

type Service interface {
	GetMembers(ctx context.Context) (res *entity.ListMemberResponse, err error)
	GetMemberByID(ctx context.Context, id *string) (res *entity.GetMemberResponse, err error)
	CreateMember(ctx context.Context, arg *entity.CreateMemberRequest) (res *entity.CreateMemberResponse, err error)
	UpdateMember(ctx context.Context, arg *entity.UpdateMemberRequest, id *string) (res *entity.UpdateMemberResponse, err error)
	DeleteMember(ctx context.Context, id *string) (res *entity.DeleteMemberResponse, err error)
}

type Repository interface {
	GetMembers(ctx context.Context) ([]model.Member, error)
	GetMemberByID(ctx context.Context, id uuid.UUID) (model.Member, error)
	InsertMember(ctx context.Context, arg model.InsertMemberParams) (model.Member, error)
	UpdateMember(ctx context.Context, arg model.UpdateMemberParams) (model.Member, error)
	DeleteMember(ctx context.Context, id uuid.UUID) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetMembers(ctx context.Context) (*entity.ListMemberResponse, error) {
	data, err := s.r.GetMembers(ctx)
	if err != nil {
		log.ErrorDetail(adminTag, "error GetAllMember from DB: %v", err)
		return nil, err
	}

	res := &entity.ListMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) GetMemberByID(ctx context.Context, id *string) (*entity.GetMemberResponse, error) {
	idParsed, err := uuid.Parse(*id)
	if err != nil {
		log.ErrorDetail(adminTag, "error parse uuid: %v", err)
		return nil, err
	}

	data, err := s.r.GetMemberByID(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(adminTag, "error GetMemberByID from DB: %v", err)
		return nil, err
	}

	res := &entity.GetMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) CreateMember(ctx context.Context, arg *entity.CreateMemberRequest) (*entity.CreateMemberResponse, error) {
	dataInsert := model.InsertMemberParams{
		Name:     arg.Data.Name,
		Username: arg.Data.Username,
		Password: arg.Data.Password,
		Email:    arg.Data.Email,
	}

	data, err := s.r.InsertMember(ctx, dataInsert)
	if err != nil {
		log.ErrorDetail(adminTag, "error InsertMember from DB: %v", err)
		return nil, err
	}

	res := &entity.CreateMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) UpdateMember(ctx context.Context, arg *entity.UpdateMemberRequest, id *string) (*entity.UpdateMemberResponse, error) {
	idParsed, err := uuid.Parse(*id)
	if err != nil {
		log.ErrorDetail(adminTag, "error parse uuid: %v", err)
		return nil, err
	}

	dataUpdate := model.UpdateMemberParams{
		ID:       idParsed,
		Name:     arg.Data.Name.String,
		Username: arg.Data.Username.String,
		Email:    arg.Data.Email.String,
	}

	data, err := s.r.UpdateMember(ctx, dataUpdate)
	if err != nil {
		log.ErrorDetail(adminTag, "error UpdateMember from DB: %v", err)
		return nil, err
	}

	res := &entity.UpdateMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) DeleteMember(ctx context.Context, id *string) (*entity.DeleteMemberResponse, error) {
	idParsed, err := uuid.Parse(*id)
	if err != nil {
		log.ErrorDetail(adminTag, "error parse uuid: %v", err)
		return nil, err
	}

	err = s.r.DeleteMember(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(adminTag, "error DeleteMember from DB: %v", err)
		return nil, err
	}

	res := &entity.DeleteMemberResponse{
		ID:       *id,
		Response: response.OK,
	}

	return res, nil
}
