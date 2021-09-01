package attempt

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/aliirsyaadn/kodein/internal/redis"
	"github.com/aliirsyaadn/kodein/internal/response"
	"github.com/aliirsyaadn/kodein/model"
)

const attemptTag = "AttemptService"
const attempt = "attempt"
const attemptMember = "attempt_member"
const expiredCache time.Duration = 60

type Service interface {
	GetAttemptsByMemberID(ctx context.Context, memberID string) (res entity.ListAttemptResponse, err error)
	GetAttemptByID(ctx context.Context, id string) (res entity.GetAttemptResponse, err error)
	CreateAttempt(ctx context.Context, arg entity.CreateAttemptRequest) (res entity.CreateAttemptResponse, err error)
	UpdateAttempt(ctx context.Context, arg entity.UpdateAttemptRequest, id string) (res entity.UpdateAttemptResponse, err error)
	DeleteAttempt(ctx context.Context, id string) (res entity.DeleteAttemptResponse, err error)
}

type Repository interface {
	GetAttemptsByMemberID(ctx context.Context, memberID uuid.UUID) ([]model.Attempt, error)
	GetAttemptByID(ctx context.Context, id uuid.UUID) (model.Attempt, error)
	InsertAttempt(ctx context.Context, arg model.InsertAttemptParams) (model.Attempt, error)
	UpdateAttempt(ctx context.Context, arg model.UpdateAttemptParams) (model.Attempt, error)
	DeleteAttempt(ctx context.Context, id uuid.UUID) error
}

type service struct {
	r   Repository
	rc  redis.RedisCache
}

func NewService(r Repository, rc redis.RedisCache) Service {
	return &service{
		r,
		rc,
	}
}

func (s *service) GetAttemptsByMemberID(ctx context.Context, memberID string) (entity.ListAttemptResponse, error) {
	var res entity.ListAttemptResponse
	var data []model.Attempt

	// Get Data from cache
	dataRedis := s.rc.Get(ctx, attemptMember, memberID)
	if dataRedis == "" {
		memberIDParsed, err := uuid.Parse(memberID)
		if err != nil {
			log.ErrorDetail(attemptTag, "error parse uuid: %v", err)
			return res, err
		}

		data, err = s.r.GetAttemptsByMemberID(ctx, memberIDParsed)
		if err != nil {
			log.ErrorDetail(attemptTag, "error GetAllAttempt from DB: %v", err)
			return res, err
		}
		ok := s.rc.SetJSON(ctx, attemptMember, memberID, data, expiredCache)
		if !ok {
			log.WarnDetail(attemptTag, "error set redis key:%s:%s", attemptMember, memberID)
		}
	} else {
		err := json.Unmarshal([]byte(dataRedis), &data)
		if err != nil {
			log.ErrorDetail(attemptTag, "error parse json: %v", err)
			return res, err
		}
		log.DebugDetail(attemptTag, "%v", data)
	}
	
	res = entity.ListAttemptResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) GetAttemptByID(ctx context.Context, id string) (entity.GetAttemptResponse, error) {
	var res entity.GetAttemptResponse
	var data model.Attempt

	// Get data from cache
	dataRedis := s.rc.Get(ctx, attempt, id)
	if dataRedis == "" {
		idParsed, err := uuid.Parse(id)
		if err != nil {
			log.ErrorDetail(attemptTag, "error parse uuid: %v", err)
			return res, err
		}

		data, err = s.r.GetAttemptByID(ctx, idParsed)
		if err != nil {
			log.ErrorDetail(attemptTag, "error GetAttemptByID from DB: %v", err)
			return res, err
		}

		// Set data to cache
		ok := s.rc.SetJSON(ctx, attempt, id, data, expiredCache)
		if !ok {
			log.WarnDetail(attemptTag, "error set redis key:%s:%s", attempt, id)
		}

	} else {
		err := json.Unmarshal([]byte(dataRedis), &data)
		if err != nil {
			log.ErrorDetail(attemptTag, "error parse json: %v", err)
			return res, err
		}
		log.DebugDetail(attemptTag, "%v", data)
	}

	res = entity.GetAttemptResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) CreateAttempt(ctx context.Context, arg entity.CreateAttemptRequest) (res entity.CreateAttemptResponse, err error) {
	dataInsert := model.InsertAttemptParams{
		MemberID:  arg.Data.MemberID,
		ProblemID: arg.Data.ProblemID,
		Language:  model.LanguageTypePython,
		IsSolved:  arg.Data.IsSolved,
		Score:     arg.Data.Score,
		Code:      arg.Data.Code,
	}

	data, err := s.r.InsertAttempt(ctx, dataInsert)
	if err != nil {
		log.ErrorDetail(attemptTag, "error InsertAttempt from DB: %v", err)
		return
	}

	ok := s.rc.SetJSON(ctx, attempt, data.ID.String(), data, expiredCache)
	if !ok {
		log.WarnDetail(attemptTag, "error set redis key:%s:%s", attempt, data.ID.String())
	}

	res = entity.CreateAttemptResponse{
		Data:     data,
		Response: response.OK,
	}

	return res, nil
}

func (s *service) UpdateAttempt(ctx context.Context, arg entity.UpdateAttemptRequest, id string) (res entity.UpdateAttemptResponse, err error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(attemptTag, "error parse uuid: %v", err)
		return
	}

	dataUpdate := model.UpdateAttemptParams{
		ID:       idParsed,
		Language: model.LanguageTypePython,
		IsSolved: arg.Data.IsSolved.Bool,
		Score:    int16(arg.Data.Score.Int64),
		Code:     arg.Data.Code.String,
	}

	data, err := s.r.UpdateAttempt(ctx, dataUpdate)
	if err != nil {
		log.ErrorDetail(attemptTag, "error UpdateAttempt from DB: %v", err)
		return
	}

	ok := s.rc.SetJSON(ctx, attempt, data.ID.String(), data, expiredCache)
	if !ok {
		log.WarnDetail(attemptTag, "error set redis key:%s:%s", attempt, data.ID.String())
	}

	res = entity.UpdateAttemptResponse{
		Data:     data,
		Response: response.OK,
	}

	return
}

func (s *service) DeleteAttempt(ctx context.Context, id string) (res entity.DeleteAttemptResponse, err error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.ErrorDetail(attemptTag, "error parse uuid: %v", err)
		return
	}

	err = s.r.DeleteAttempt(ctx, idParsed)
	if err != nil {
		log.ErrorDetail(attemptTag, "error DeleteAttempt from DB: %v", err)
		return
	}

	// Delete data in cache
	s.rc.Del(ctx, attempt, id)

	res = entity.DeleteAttemptResponse{
		ID:       id,
		Response: response.OK,
	}

	return res, nil
}
