package pkr

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/lanjinghexuan/project/common/gload"
	"time"
)

type Cache struct {
	Prefix       string
	Params       interface{}
	Expire       time.Duration
	ForceRefresh bool
}

func GetCache(resultPtr interface{}, opts Cache, query func() (interface{}, error)) error {
	fmt.Println(opts)
	//把参数转为json格式
	keyBytes, _ := json.Marshal(opts.Params)
	//参数进行md5
	keyHash := fmt.Sprintf("%x", md5.Sum(keyBytes))
	//生成最后的key
	cacheKey := fmt.Sprintf("%s:%s", opts.Prefix, keyHash)
	//如果不刷新

	if !opts.ForceRefresh {
		//获取缓存是否存在
		cached, err := gload.REDIS.Get(gload.Ctx, cacheKey).Result()
		fmt.Println(cached, err)
		//无报错解析json
		if err == nil && cached != "" {
			if err := json.Unmarshal([]byte(cached), resultPtr); err == nil {
				return nil
			}
		}
	}
	//无缓存或者清除缓存调用方法2
	dbResult, err := query()
	if err != nil {
		return err
	}
	dbBytes, err := json.Marshal(dbResult)
	if err == nil {
		gload.REDIS.Set(gload.Ctx, cacheKey, dbBytes, opts.Expire)
	}
	json.Unmarshal(dbBytes, resultPtr)
	return nil
}

// FetchWithCache 封装的列表查询缓存逻辑
// resultPtr 是一个指向结果的指针（例如 *[]User），queryFunc 是未命中缓存时的查询函数
func FetchWithCache(resultPtr interface{}, opts Cache, queryFunc func() (interface{}, error)) error {
	// 1. 构建缓存键
	keyBytes, _ := json.Marshal(opts.Params)
	keyHash := fmt.Sprintf("%x", md5.Sum(keyBytes))
	cacheKey := fmt.Sprintf("%s:%s", opts.Prefix, keyHash)

	// 2. 如果不是强制刷新，尝试读取缓存
	if !opts.ForceRefresh {
		cached, err := gload.REDIS.Get(gload.Ctx, cacheKey).Result()
		if err == nil {
			// 命中缓存
			if jsonErr := json.Unmarshal([]byte(cached), resultPtr); jsonErr == nil {
				return nil
			}
			// 如果缓存格式错误，继续查库并重建缓存
		}
	}

	// 3. 未命中缓存或强制刷新，查询数据库
	dbResult, err := queryFunc()
	if err != nil {
		return err
	}

	// 4. 序列化并写入缓存
	dbBytes, err := json.Marshal(dbResult)
	if err != nil {
		gload.REDIS.Set(gload.Ctx, cacheKey, dbBytes, opts.Expire)
	}

	// 5. 写入调用方传入的结果指针
	// resultPtr 是 interface{}，dbResult 是 interface{}，我们先转类型
	jsonBytes, err := json.Marshal(dbResult)
	if err == nil {
		json.Unmarshal(jsonBytes, resultPtr)
	}
	return nil
}
