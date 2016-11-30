package cache

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type CacheSetTaskResult struct {
	app.Result
}

type CacheSetTask struct {
	app.Task
	Key     string `json:"key"`
	Value   string `json:"value"`
	Expires int64  `json:"expires"`
	Result  CacheSetTaskResult
}

func (task *CacheSetTask) GetResult() interface{} {
	return &task.Result
}

func (task *CacheSetTask) GetInhertType() string {
	return "cache"
}

func (task *CacheSetTask) GetClientName() string {
	return "Cache.Set"
}
