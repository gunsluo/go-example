package main

import (
	"fmt"
	"time"

	"github.com/karlseguin/ccache"
)

type User struct {
	ID   string
	Name string
}

func main() {
	//test()
	test2()
}

func test() {
	var cache = ccache.New(ccache.Configure())

	user := &User{ID: "user:4", Name: "luoji"}
	cache.Set(user.ID, user, time.Second)

	item := cache.Get(user.ID)
	if item == nil {
		//handle
	} else {
		nu := item.Value().(*User)
		fmt.Println("-->", user == nu)
	}

	time.Sleep(time.Second * 2)
	item = cache.Get(user.ID)
	if item == nil {
		//handle
	} else {
		nu := item.Value().(*User)
		fmt.Println("-->", user == nu, " expire:", item.Expired())
		if item.Expired() {
			ok := cache.Delete(user.ID)
			fmt.Println("delete:", ok)
		}
	}

	item, err := cache.Fetch("user:4", time.Minute*10, func() (interface{}, error) {
		//code to fetch the data incase of a miss
		//should return the data to cache and the error, if any
		return user, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("fetch:")
}

func test2() {
	var cache = ccache.New(ccache.Configure().MaxSize(1).Buckets(1))

	user := &User{ID: "user:4", Name: "luoji"}
	user2 := &User{ID: "user:5", Name: "luoji2"}
	cache.Set(user.ID, user, time.Second)
	cache.Set(user2.ID, user2, time.Second)

	item := cache.Get(user.ID)
	if item == nil {
		//handle
	} else {
		nu := item.Value().(*User)
		fmt.Println("-->", user == nu, nu.ID, nu.Name)
	}

	item = cache.Get(user2.ID)
	if item == nil {
		//handle
	} else {
		nu := item.Value().(*User)
		fmt.Println("-->", user2 == nu, nu.ID, nu.Name)
	}

	time.Sleep(time.Second * 2)
	item = cache.Get(user.ID)
	if item == nil {
		//handle
	} else {
		nu := item.Value().(*User)
		fmt.Println("-->", user == nu, " expire:", item.Expired())
		if item.Expired() {
			ok := cache.Delete(user.ID)
			fmt.Println("delete:", ok)
		}
	}
}
