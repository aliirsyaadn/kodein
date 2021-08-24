package member

import (
	"context"

	"github.com/google/uuid"

	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/model"
)

const adminTag = "MemberService"

type Service interface {
	GetMembers(ctx context.Context) (data []model.Member, err error)
	GetMemberByID(ctx context.Context, id uuid.UUID) (data *model.Member, err error)
	GetMemberByUsername(ctx context.Context, username string) (data *model.Member, err error)
	CreateMember(ctx context.Context, username, password string) (data *model.Member, err error)
	UpdateMember(ctx context.Context, arg *model.UpdateMemberParams, ID uuid.UUID) (data *model.Member, err error)
	DeleteMember(ctx context.Context, ID uuid.UUID) (status bool, err error)
}

type Repository interface {
	GetMembers(ctx context.Context) ([]model.Member, error)
	GetMemberByID(ctx context.Context, id uuid.UUID) (model.Member, error)
	GetMemberByUsername(ctx context.Context, username string) (model.Member, error)
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

func (s *service) GetMembers(ctx context.Context) ([]model.Member, error) {
	data, err := s.r.GetMembers(ctx)
	if err != nil {
		log.ErrorDetail(adminTag, "error GetAllMember from DB: %v", err)
		return nil, err
	}

	return data, nil
}

func (s *service) GetMemberByID(ctx context.Context, ID uuid.UUID) (*model.Member, error) {
	data, err := s.r.GetMemberByID(ctx, ID)
	if err != nil {
		log.ErrorDetail(adminTag, "error GetMemberByID from DB: %v", err)
		return nil, err
	}

	return &data, nil
}

func (s *service) GetMemberByUsername(ctx context.Context, username string) (*model.Member, error) {
	data, err := s.r.GetMemberByUsername(ctx, username)
	if err != nil {
		log.ErrorDetail(adminTag, "error GetMemberByUsername from DB: %v", err)
		return nil, err
	}

	return &data, nil
}

func (s *service) CreateMember(ctx context.Context, username, password string) (*model.Member, error) {
	dataInsert := model.InsertMemberParams{
		Username: username,
		Password: password,
	}

	data, err := s.r.InsertMember(ctx, dataInsert)
	if err != nil {
		log.ErrorDetail(adminTag, "error InsertMember from DB: %v", err)
		return nil, err
	}

	return &data, nil
}

func (s *service) UpdateMember(ctx context.Context, arg *model.UpdateMemberParams, ID uuid.UUID) (*model.Member, error) {
	data, err := s.r.UpdateMember(ctx, *arg)
	if err != nil {
		log.ErrorDetail(adminTag, "error UpdateMember from DB: %v", err)
		return nil, err
	}

	return &data, nil
}

func (s *service) DeleteMember(ctx context.Context, ID uuid.UUID) (bool, error) {
	err := s.r.DeleteMember(ctx, ID)
	if err != nil {
		log.ErrorDetail(adminTag, "error DeleteMember from DB: %v", err)
		return false, err
	}

	return true, nil
}
