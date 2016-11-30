package cache

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CacheRemoveTaskResult struct {
	app.Result
}

type CacheRemoveTask struct {
	app.Task
	Key    string `json:"key"`
	Prefix string `json:"prefix"`
	Result CacheRemoveTaskResult
}

func (task *CacheRemoveTask) GetResult() interface{} {
	return &task.Result
}

func (task *CacheRemoveTask) GetInhertType() string {
	return "cache"
}

func (task *CacheRemoveTask) GetClientName() string {
	return "Cache.Remove"
}
