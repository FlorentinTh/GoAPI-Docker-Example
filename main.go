package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/go-redis/redis"
)

// This example shows the minimal code needed to get a restful.WebService working.
//
// GET http://localhost:8080/hello

var client *redis.Client

func main() {

	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	err = client.Set("hit", "0", 0).Err()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(root))
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "To see the hello world go to /hello")
}

func hello(req *restful.Request, resp *restful.Response) {
	val, _ := client.Get("hit").Result()

	if val != "" {
		intVal, _ := strconv.Atoi(val)
		io.WriteString(resp, "world " + val)
		client.Set("hit", strconv.Itoa(intVal + 1), 0).Err()
	} else {
		io.WriteString(resp, "world " + val)
	}
}
