package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password", // no password set
		DB:       0,          // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	user := &User{ID: "user:4", Name: "luoji"}
	user2 := &User{ID: "user:5", Name: "luoji2"}

	buffer, err := msgpack.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = client.Set(user.ID, buffer, time.Second).Err()
	if err != nil {
		panic(err)
	}

	buffer, err = msgpack.Marshal(user2)
	if err != nil {
		panic(err)
	}

	err = client.Set(user2.ID, buffer, time.Second).Err()
	if err != nil {
		panic(err)
	}

	//item, err := client.Get(user.ID).Result()
	buffer, err = client.Get(user.ID).Bytes()
	if err != nil {
		panic(err)
	}

	nuser := &User{}
	err = msgpack.Unmarshal(buffer, nuser)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", nuser.ID, nuser.Name)

	buffer, err = client.Get(user2.ID).Bytes()
	if err != nil {
		panic(err)
	}
	nuser2 := &User{}
	err = msgpack.Unmarshal(buffer, nuser2)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", nuser2.ID, nuser2.Name)

	// delete
	n, err := client.Del(user2.ID).Result()
	fmt.Println("delete", n, err)

	_, err = client.Get(user2.ID).Bytes()
	if err != nil {
		if err != redis.Nil {
			panic(err)
		}
		fmt.Println("not found")
	}

	time.Sleep(time.Second * 2)
	_, err = client.Get(user.ID).Bytes()
	if err != nil {
		if err != redis.Nil {
			panic(err)
		}
		fmt.Println("not found")
	}

	_, err = client.Get(user2.ID).Bytes()
	if err != nil {
		if err != redis.Nil {
			panic(err)
		}
		fmt.Println("not found")
	}

	client.Close()
}

type User struct {
	ID   string
	Name string
}

/*
func (u *User) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, u)
}
*/
