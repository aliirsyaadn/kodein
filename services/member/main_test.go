package member

import (
	"context"
	"reflect"
	"testing"

	"github.com/aliirsyaadn/kodein/entity"
)

func Test_service_GetMemberByID(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		id  *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.GetMemberResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
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

func Test_service_CreateMember(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		arg *entity.CreateMemberRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.CreateMemberResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
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

func Test_service_UpdateMember(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		arg *entity.UpdateMemberRequest
		id  *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.UpdateMemberResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
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

func Test_service_DeleteMember(t *testing.T) {
	type fields struct {
		r Repository
	}
	type args struct {
		ctx context.Context
		id  *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.DeleteMemberResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				r: tt.fields.r,
			}
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
