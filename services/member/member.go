package member

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/internal/redis"
	"github.com/aliirsyaadn/kodein/internal/response"
	"github.com/aliirsyaadn/kodein/model"
)

const memberTag = "MemberService"
const member = "member"
const expiredCache time.Duration = 60

type Service interface {
	GetMembers(ctx context.Context) (res entity.ListMemberResponse, err error)
	GetMemberByID(ctx context.Context, id string) (res entity.GetMemberResponse, err error)
	CreateMember(ctx context.Context, arg entity.CreateMemberRequest) (res entity.CreateMemberResponse, err error)
	UpdateMember(ctx context.Context, arg entity.UpdateMemberRequest, id string) (res entity.UpdateMemberResponse, err error)
	DeleteMember(ctx context.Context, id string) (res entity.DeleteMemberResponse, err error)
}

type Repository interface {
	GetMembers(ctx context.Context) ([]model.Member, error)
	GetMemberByID(ctx context.Context, id uuid.UUID) (model.Member, error)
	InsertMember(ctx context.Context, arg model.InsertMemberParams) (model.Member, error)
	UpdateMember(ctx context.Context, arg model.UpdateMemberParams) (model.Member, error)
	DeleteMember(ctx context.Context, id uuid.UUID) error
}

type service struct {
	r  Repository
	rc redis.RedisCache
}

func NewService(r Repository, rc redis.RedisCache) Service {
	return &service{r, rc}
}

func (s *service) GetMembers(ctx context.Context) (entity.ListMemberResponse, error) {
	var res entity.ListMemberResponse
	var data []model.Member

	// Get Data from cache
	err := s.rc.GetJSON(ctx, member, "s", &data)
	if err != nil {
		data, err := s.r.GetMembers(ctx)
		if err != nil {
			log.ErrorDetail(memberTag, "error GetAllMember from DB: %v", err)
			return res, err
		}
		ok := s.rc.SetJSON(ctx, member, "s", data, expiredCache)
		if !ok {
			log.WarnDetail(memberTag, "error set redis key:%s:%s", member, "s")
		}
	} else {
		log.DebugDetail(memberTag, "%v", data)
	}

	res = entity.ListMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) GetMemberByID(ctx context.Context, id string) (entity.GetMemberResponse, error) {
	var res entity.GetMemberResponse
	var data model.Member

	// Get data from cache
	err := s.rc.GetJSON(ctx, member, id, &data)
	if err != nil {
		idParsed, err := uuid.Parse(id)
		if err != nil {
			log.ErrorDetail(memberTag, "error parse uuid: %v", err)
			return res, err
		}
		data, err := s.r.GetMemberByID(ctx, idParsed)
		if err != nil {
			log.ErrorDetail(memberTag, "error GetMemberByID from DB: %v", err)
			return res, err
		}

		// Set data to cache
		ok := s.rc.SetJSON(ctx, member, id, data, expiredCache)
		if !ok {
			log.WarnDetail(memberTag, "error set redis key:%s:%s", member, id)
		}

	} else {
		log.DebugDetail(memberTag, "%v", data)
	}

	res = entity.GetMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) CreateMember(ctx context.Context, arg entity.CreateMemberRequest) (res entity.CreateMemberResponse, err error) {
	dataInsert := model.InsertMemberParams{
		Name:     arg.Data.Name,
		Username: arg.Data.Username,
		Password: arg.Data.Password,
		Email:    arg.Data.Email,
	}

	data, err := s.r.InsertMember(ctx, dataInsert)
	if err != nil {
		log.ErrorDetail(memberTag, "error InsertMember from DB: %v", err)
		return
	}

	ok := s.rc.SetJSON(ctx, member, data.ID.String(), data, expiredCache)
	if !ok {
		log.WarnDetail(memberTag, "error set redis key:%s:%s", member, data.ID.String())
	}

	res = entity.CreateMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) UpdateMember(ctx context.Context, arg entity.UpdateMemberRequest, id string) (res entity.UpdateMemberResponse, err error) {

	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(memberTag, "error parse uuid: %v", err)
		return
	}

	dataUpdate := model.UpdateMemberParams{
		ID:       idParsed,
		Name:     arg.Data.Name.String,
		Username: arg.Data.Username.String,
		Email:    arg.Data.Email.String,
	}

	data, err := s.r.UpdateMember(ctx, dataUpdate)
	if err != nil {
		log.ErrorDetail(memberTag, "error UpdateMember from DB: %v", err)
		return
	}

	ok := s.rc.SetJSON(ctx, member, data.ID.String(), data, expiredCache)
	if !ok {
		log.WarnDetail(memberTag, "error set redis key:%s:%s", member, data.ID.String())
	}

	res = entity.UpdateMemberResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) DeleteMember(ctx context.Context, id string) (res entity.DeleteMemberResponse, err error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(memberTag, "error parse uuid: %v", err)
		return
	}

	err = s.r.DeleteMember(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(memberTag, "error DeleteMember from DB: %v", err)
		return
	}

	// Delete data in cache
	s.rc.Del(ctx, member, id)

	res = entity.DeleteMemberResponse{
		ID:       id,
		Response: response.OK,
	}

	return res, nil
}
