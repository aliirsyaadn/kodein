package member

import (
	"context"

	"github.com/google/uuid"

	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/model"
)

const adminTag = "MemberService"

type Service interface {
	GetMembers(ctx context.Context) (data []model.Member, err error)
	GetMemberByID(ctx context.Context, ID string) (data *model.Member, err error)
	GetMemberByUsername(ctx context.Context, username string) (data *model.Member, err error)
	CreateMember(ctx context.Context, arg *model.InsertMemberParams) (data *model.Member, err error)
	UpdateMember(ctx context.Context, arg entity.UpdateMemberRequest, ID string) (data *model.Member, err error)
	DeleteMember(ctx context.Context, ID string) (err error)
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

func (s *service) GetMemberByID(ctx context.Context, ID string) (*model.Member, error) {
	idParsed, err := uuid.Parse(ID)
	if err != nil {
		log.ErrorDetail(adminTag, "error parse uuid: %v", err)
		return nil, err
	}
	data, err := s.r.GetMemberByID(ctx, idParsed)
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

func (s *service) CreateMember(ctx context.Context, arg *model.InsertMemberParams) (*model.Member, error) {
	data, err := s.r.InsertMember(ctx, *arg)
	if err != nil {
		log.ErrorDetail(adminTag, "error InsertMember from DB: %v", err)
		return nil, err
	}

	return &data, nil
}

func (s *service) UpdateMember(ctx context.Context, arg entity.UpdateMemberRequest, ID string) (*model.Member, error) {
	idParsed, err := uuid.Parse(ID)
	if err != nil {
		log.ErrorDetail(adminTag, "error parse uuid: %v", err)
		return nil, err
	}

	dataUpdate := model.UpdateMemberParams{
		ID : idParsed,
		Name: arg.Data.Name.String,
		Username: arg.Data.Username.String,
		Email: arg.Data.Email.String,
	}

	
	data, err := s.r.UpdateMember(ctx, dataUpdate)
	if err != nil {
		log.ErrorDetail(adminTag, "error UpdateMember from DB: %v", err)
		return nil, err
	}

	return &data, nil
}

func (s *service) DeleteMember(ctx context.Context, ID string) (error) {
	idParsed, err := uuid.Parse(ID)
	if err != nil {
		log.ErrorDetail(adminTag, "error parse uuid: %v", err)
		return err
	}

	err = s.r.DeleteMember(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(adminTag, "error DeleteMember from DB: %v", err)
		return err
	}

	return nil
}
