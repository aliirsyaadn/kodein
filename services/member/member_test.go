package member

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"

	"github.com/aliirsyaadn/kodein/entity"
	"github.com/aliirsyaadn/kodein/internal/response"
	"github.com/aliirsyaadn/kodein/model"
	mock "github.com/aliirsyaadn/kodein/services/member/mock"
)

func TestGetMembers(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockRepository(mockCtrl)
	id1 := uuid.New()
	id2 := uuid.New()

	mockRepo.EXPECT().GetMembers(context.Background()).Return([]model.Member{
		{
			ID: id1,
			Name: "Ali Irsyaad",
			Username: "aliirsyaadn",
			Password: "alipassword",
			Email: "ali@gmail.com",
		},
		{
			ID: id2,
			Name: "John Doe",
			Username: "johnDoe",
			Password: "johnpassword",
			Email: "john@gmail.com",
		},
	}, nil)

	memberService := NewService(mockRepo)

	members, err := memberService.GetMembers(context.Background())
	
	assert.Nil(t, err)
	assert.NotNil(t, members)
	assert.Equal(t, len(members.Data), 2)
}

func TestGetMemberByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockRepository(mockCtrl)
	id1 := uuid.New()
	id2 := uuid.New()

	rtr1 := model.Member{
		ID: id1,
		Name: "Ali Irsyaad",
		Username: "aliirsyaadn",
		Password: "alipassword",
		Email: "ali@gmail.com",
	}

	rtr2 := model.Member{
		ID: id2,
		Name: "John Doe",
		Username: "johnDoe",
		Password: "johnpassword",
		Email: "john@gmail.com",	
	}

	mockRepo.EXPECT().GetMemberByID(context.Background(), id1).Return(rtr1, nil)

	mockRepo.EXPECT().GetMemberByID(context.Background(), id2).Return(rtr2, nil)

	
	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.GetMemberResponse
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id: id1.String(),
			},
			want: entity.GetMemberResponse{
				Data: rtr1,
				Response: response.OK,
			},
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id: id2.String(),
			},
			want: entity.GetMemberResponse{
				Data: rtr2,
				Response: response.OK,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.r)
	
			got, err := s.GetMemberByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetMemberByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetMemberByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateMember(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockRepository(mockCtrl)
	id1 := uuid.New()
	id2 := uuid.New()

	dataInsert1 := model.InsertMemberParams{
		Name: "Ali Irsyaad",
		Username: "aliirsyaadn",
		Password: "alipassword",
		Email: "ali@gmail.com",
	}

	dataInsert2 := model.InsertMemberParams{
		Name: "John Doe",
		Username: "johnDoe",
		Password: "johnpassword",
		Email: "john@gmail.com",
	}

	rtr1 := model.Member{
		ID: id1,
		Name: "Ali Irsyaad",
		Username: "aliirsyaadn",
		Password: "alipassword",
		Email: "ali@gmail.com",
	}
	rtr2 := model.Member{
		ID: id2,
		Name: "John Doe",
		Username: "johnDoe",
		Password: "johnpassword",
		Email: "john@gmail.com",
	}
	mockRepo.EXPECT().InsertMember(context.Background(), dataInsert1).Return(rtr1, nil)

	mockRepo.EXPECT().InsertMember(context.Background(), dataInsert2).Return(rtr2, nil)

	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		arg entity.CreateMemberRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.CreateMemberResponse
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				arg: entity.CreateMemberRequest{
					Data: dataInsert1,
				},
			},
			want: entity.CreateMemberResponse{
				Data: rtr1,
				Response: response.OK,
			},
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				arg: entity.CreateMemberRequest{
					Data: dataInsert2,
				},
			},
			want: entity.CreateMemberResponse{
				Data: rtr2,
				Response: response.OK,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.r)
			got, err := s.CreateMember(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMember(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockRepository(mockCtrl)
	id1 := uuid.New()
	id2 := uuid.New()

	dataUpdate1 := model.UpdateMemberParams{
		ID: id1,
		Name: "Ali Irsyaad",
		Username: "aliirsyaadn",
		Email: "ali@gmail.com",
	}

	dataUpdate2 := model.UpdateMemberParams{
		ID: id2,
		Name: "John Doe",
		Username: "johnDoe",
		Email: "john@gmail.com",
	}

	rtr1 := model.Member{
		ID: id1,
		Name: "Ali Irsyaad",
		Username: "aliirsyaadn",
		Password: "alipassword",
		Email: "ali@gmail.com",
	}

	rtr2 := model.Member{
		ID: id2,
		Name: "John Doe",
		Username: "johnDoe",
		Password: "johnpassword",
		Email: "john@gmail.com",	
	}

	mockRepo.EXPECT().UpdateMember(context.Background(), dataUpdate1).Return(rtr1, nil)

	mockRepo.EXPECT().UpdateMember(context.Background(), dataUpdate2).Return(rtr2, nil)

	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		arg entity.UpdateMemberRequest
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.UpdateMemberResponse
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				arg: entity.UpdateMemberRequest{
					Data: entity.UpdateMember{
						Name: null.NewString(dataUpdate1.Name, true),
						Username: null.NewString(dataUpdate1.Username, true),
						Email: null.NewString(dataUpdate1.Email, true),
					},
				},
				id: id1.String(),
			},
			want: entity.UpdateMemberResponse{
				Data: rtr1,
				Response: response.OK,
			},
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				arg: entity.UpdateMemberRequest{
					Data: entity.UpdateMember{
						Name: null.NewString(dataUpdate2.Name, true),
						Username: null.NewString(dataUpdate2.Username, true),
						Email: null.NewString(dataUpdate2.Email, true),
					},
				},
				id: id2.String(),
			},
			want: entity.UpdateMemberResponse{
				Data: rtr2,
				Response: response.OK,
			},
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.r)
			got, err := s.UpdateMember(tt.args.ctx, tt.args.arg, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.UpdateMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMember(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock.NewMockRepository(mockCtrl)
	id1 := uuid.New()
	id2 := uuid.New()

	mockRepo.EXPECT().DeleteMember(context.Background(), id1).Return(nil)
	mockRepo.EXPECT().DeleteMember(context.Background(), id2).Return(nil)

	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.DeleteMemberResponse
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id: id1.String(),
			},
			want: entity.DeleteMemberResponse{
				ID: id1.String(),
				Response: response.OK,
			},
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				r: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				id: id2.String(),
			},
			want: entity.DeleteMemberResponse{
				ID: id2.String(),
				Response: response.OK,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.r)
			got, err := s.DeleteMember(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.DeleteMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

