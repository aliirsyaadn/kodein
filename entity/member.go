package entity

import (
	"github.com/aliirsyaadn/kodein/model"
	"gopkg.in/guregu/null.v3"
)

type ListMemberResponse struct{
	Data []model.Member `json:"data"`
	Response
}


type CreateMemberRequest struct {
	Data model.InsertMemberParams `json:"data"`
}

type CreateMemberResponse struct {
	Data model.Member `json:"data"`
	Response
}

type UpdateMember struct {
	Name null.String `json:"name"`
	Username null.String `json:"username"`
	Email null.String `json:"email"`
}

type UpdateMemberRequest struct {
	Data UpdateMember `json:"data"`
}

type UpdateMemberResponse struct {
	Data model.Member `json:"data"`
	Response
}