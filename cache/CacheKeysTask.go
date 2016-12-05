package cache

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CacheKeysTaskResult struct {
	app.Result
	Keys []string `json:"keys,omitempty"`
}

type CacheKeysTask struct {
	app.Task
	Prefix string `json:"prefix"`
	Result CacheKeysTaskResult
}

func (task *CacheKeysTask) GetResult() interface{} {
	return &task.Result
}

func (task *CacheKeysTask) GetInhertType() string {
	return "cache"
}

func (task *CacheKeysTask) GetClientName() string {
	return "Cache.Keys"
}
