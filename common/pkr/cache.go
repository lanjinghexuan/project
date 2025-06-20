package pkr

import (
	"crypto/md5"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/lanjinghexuan/project/common/gload"
	"time"
)

type CacheData struct {
	Prefix     string
	Params     interface{}
	Expire     time.Duration
	ClearCache bool
}

func GetCaches(resData interface{}, opts CacheData, query func() (interface{}, error)) error {
	ParamsStr, _ := json.Marshal(opts.Params)
	ParamsHash := fmt.Sprintf("%x", md5.Sum(ParamsStr))
	newCacheKey := fmt.Sprintf("%s_%s", opts.Prefix, ParamsHash)

	if !opts.ClearCache {
		res, err := gload.REDIS.Get(gload.Ctx, newCacheKey).Result()
		if err == nil && res != "" {
			err := json.Unmarshal([]byte(res), resData)
			if err == nil {
				return nil
			}
		}
	}

	res1, err := query()
	if err != nil {
		return err
	}
	strRes1, _ := json.Marshal(res1)
	err = gload.REDIS.Set(gload.Ctx, newCacheKey, string(strRes1), opts.Expire).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(strRes1, resData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

/*
//模拟调用


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUser() ([]User, error) {
	time.Sleep(10 * time.Second)
	return []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}, nil
}

func GetUserCache(c *gin.Context) {
	var res []User
	err := pkr.GetCaches(&res, pkr.CacheData{
		Prefix: "getUserCache：v2",
		Params: nil,
		Expire: 86400 * time.Second,
	}, func() (interface{}, error) {
		return getUser()
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}



*/
