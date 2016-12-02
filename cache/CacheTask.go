package cache

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CacheTaskResult struct {
	app.Result
	Expires int64  `json:"expires"`
	Value   string `json:"value,omitempty"`
}

type CacheTask struct {
	app.Task
	Key    string `json:"key"`
	Result CacheTaskResult
}

func (task *CacheTask) GetResult() interface{} {
	return &task.Result
}

func (task *CacheTask) GetInhertType() string {
	return "cache"
}

func (task *CacheTask) GetClientName() string {
	return "Cache.Get"
}
