package entity

import (
	"github.com/aliirsyaadn/kodein/model"
	"gopkg.in/guregu/null.v3"
)

type ListProjectResponse struct {
	Data []model.Project `json:"data"`
	Response
}

type GetProjectResponse struct {
	Data model.Project `json:"data"`
	Response
}

type CreateProjectRequest struct {
	Data model.InsertProjectParams `json:"data"`
}

type CreateProjectResponse struct {
	Data model.Project `json:"data"`
	Response
}

type UpdateProject struct {
	MemberID    string      `json:"member_id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	Technology  null.String `json:"technology"`
	Url         null.String `json:"url"`
}

type UpdateProjectRequest struct {
	Data UpdateProject `json:"data"`
}

type UpdateProjectResponse struct {
	Data model.Project `json:"data"`
	Response
}

type DeleteProjectResponse struct {
	ID string `json:"id"`
	Response
}
