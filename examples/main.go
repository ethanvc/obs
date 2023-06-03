package main

import (
	"context"
	"time"
)

func main() {

}

type QueryUserReq struct {
}

type QueryUserResp struct {
}

type User struct {
}

func QueryUser(c context.Context, req *QueryUserReq) (*QueryUserResp, error) {
	c, cancel := context.WithTimeout(c, 1*time.Second)
	defer cancel()
	if GetConfig().EnableCache {

	}
	return nil, nil
}

type Config struct {
	EnableCache bool
}

func GetConfig() *Config {
	c := &Config{}
	return c
}

func GetUserFromCache(c context.Context, uid int64) (*User, error) {
	return nil, nil
}
