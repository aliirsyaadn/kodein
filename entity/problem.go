package entity

import (
	"github.com/aliirsyaadn/kodein/model"
	"gopkg.in/guregu/null.v3"
)

type ListProblemResponse struct {
	Data []model.Problem `json:"data"`
	Response
}

type GetProblemResponse struct {
	Data model.Problem `json:"data"`
	Response
}

type CreateProblemRequest struct {
	Data model.InsertProblemParams `json:"data"`
}

type CreateProblemResponse struct {
	Data model.Problem `json:"data"`
	Response
}

type UpdateProblem struct {
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	Category 	null.String `json:"category"`
	Difficulty  null.String `json:"technology"`
	Url         null.String `json:"url"`
	GraderCode  null.String `json:"grader_code"`
}

type UpdateProblemRequest struct {
	Data UpdateProblem `json:"data"`
}

type UpdateProblemResponse struct {
	Data model.Problem `json:"data"`
	Response
}

type DeleteProblemResponse struct {
	ID string `json:"id"`
	Response
}
