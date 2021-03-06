package cache

import (
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"log"
	"strings"
	"time"
)

type CacheObject struct {
	Value   string
	Expires int64
	Size    int
	Ctime   int64
}

type CacheService struct {
	app.Service
	MaxSize int
	Init    *app.InitTask
	Get     *CacheTask
	Set     *CacheSetTask
	Remove  *CacheRemoveTask
	Keys    *CacheKeysTask

	dispatch *kk.Dispatch
	objects  map[string]*CacheObject
	size     int
}

func (S *CacheService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *CacheService) HandleInitTask(a app.IApp, task *app.InitTask) error {

	S.objects = map[string]*CacheObject{}
	S.dispatch = kk.NewDispatch()
	S.size = 0

	var fn func() = nil

	fn = func() {

		var keys = []string{}
		var now = time.Now().Unix()
		var size int = 0

		for key, value := range S.objects {

			if value.Expires != 0 && value.Expires+value.Ctime < now {
				keys = append(keys, key)
				size = size + value.Size
			}

		}

		for _, key := range keys {
			delete(S.objects, key)
		}

		S.size = S.size - size

		log.Println("[CacheService][clean]", keys)

		S.dispatch.AsyncDelay(fn, time.Second*6)
	}

	S.dispatch.AsyncDelay(fn, time.Second*6)

	return nil
}

func (S *CacheService) HandleCacheTask(a app.IApp, task *CacheTask) error {

	S.dispatch.Sync(func() {

		var v, ok = S.objects[task.Key]

		if ok {
			v.Ctime = time.Now().Unix()
			task.Result.Value = v.Value
			task.Result.Expires = v.Expires
		} else {
			task.Result.Errno = ERROR_CACHE
			task.Result.Errmsg = "Not found value"
		}

	})

	return nil
}

func (S *CacheService) HandleCacheSetTask(a app.IApp, task *CacheSetTask) error {

	S.dispatch.Sync(func() {

		var v, ok = S.objects[task.Key]
		var size int = len(task.Value)

		if ok {

			if S.size+size-v.Size > S.MaxSize {
				task.Result.Errno = ERROR_CACHE_SIZE
				task.Result.Errmsg = "not enough space"
				return
			}

			S.size = S.size + size - v.Size

			v.Size = size
			v.Expires = task.Expires
			v.Ctime = time.Now().Unix()
			v.Value = task.Value

		} else {

			if S.size+size > S.MaxSize {
				task.Result.Errno = ERROR_CACHE_SIZE
				task.Result.Errmsg = "not enough space"
				return
			}

			S.size = S.size + size

			S.objects[task.Key] = &CacheObject{task.Value, task.Expires, size, time.Now().Unix()}

		}

	})

	return nil
}

func (S *CacheService) HandleCacheRemoveTask(a app.IApp, task *CacheRemoveTask) error {

	S.dispatch.Sync(func() {

		if task.Prefix != "" {

			var keys = []string{}
			var size int = 0

			for key, value := range S.objects {

				if strings.HasPrefix(key, task.Prefix) {
					keys = append(keys, key)
					size = size + value.Size
				}

			}

			for _, key := range keys {
				delete(S.objects, key)
			}

			S.size = S.size - size

		} else if task.Key != "" {

			var v, ok = S.objects[task.Key]

			if ok {
				delete(S.objects, task.Key)
				S.size = S.size - v.Size
			}
		}

	})

	return nil
}

func (S *CacheService) HandleCacheKeysTask(a app.IApp, task *CacheKeysTask) error {

	S.dispatch.Sync(func() {

		var keys = []string{}

		for key, _ := range S.objects {

			if strings.HasPrefix(key, task.Prefix) {
				keys = append(keys, key)
			}

		}

		task.Result.Keys = keys

	})

	return nil
}
