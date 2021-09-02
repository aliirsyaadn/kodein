package redis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func Test_redisCache_SetJSON(t *testing.T) {
	type fields struct {
		client *redis.Client
	}
	type args struct {
		ctx     context.Context
		service string
		id      string
		data    interface{}
		expired time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisCache{
				client: tt.fields.client,
			}
			if got := r.SetJSON(tt.args.ctx, tt.args.service, tt.args.id, tt.args.data, tt.args.expired); got != tt.want {
				t.Errorf("redisCache.SetJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisCache_Get(t *testing.T) {
	type fields struct {
		client *redis.Client
	}
	type args struct {
		ctx     context.Context
		service string
		id      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisCache{
				client: tt.fields.client,
			}
			if got := r.Get(tt.args.ctx, tt.args.service, tt.args.id); got != tt.want {
				t.Errorf("redisCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisCache_GetJSON(t *testing.T) {
	type fields struct {
		client *redis.Client
	}
	type args struct {
		ctx         context.Context
		service     string
		id          string
		destination interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisCache{
				client: tt.fields.client,
			}
			if err := r.GetJSON(tt.args.ctx, tt.args.service, tt.args.id, tt.args.destination); (err != nil) != tt.wantErr {
				t.Errorf("redisCache.GetJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisCache_Del(t *testing.T) {
	type fields struct {
		client *redis.Client
	}
	type args struct {
		ctx     context.Context
		service string
		id      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisCache{
				client: tt.fields.client,
			}
			r.Del(tt.args.ctx, tt.args.service, tt.args.id)
		})
	}
}
