package entity

import (
	"github.com/aliirsyaadn/kodein/model"
	"gopkg.in/guregu/null.v3"
)

type ListAttemptResponse struct {
	Data []model.Attempt `json:"data"`
	Response
}

type GetAttemptResponse struct {
	Data model.Attempt `json:"data"`
	Response
}

type CreateAttemptRequest struct {
	Data model.InsertAttemptParams `json:"data"`
}

type CreateAttemptResponse struct {
	Data model.Attempt `json:"data"`
	Response
}

type UpdateAttempt struct {
	Language  null.String `json:"language"`
	IsSolved  null.Bool   `json:"is_solved"`
	Score     null.Int    `json:"score"`
	Code      null.String `json:"code"`
}

type UpdateAttemptRequest struct {
	Data UpdateAttempt `json:"data"`
}

type UpdateAttemptResponse struct {
	Data model.Attempt `json:"data"`
	Response
}

type DeleteAttemptResponse struct {
	ID string `json:"id"`
	Response
}
