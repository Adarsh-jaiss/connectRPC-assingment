package main

import (
	"context"
	"log"
	"net/http"

	connect "connectrpc.com/connect"
	__ "github.com/adarsh-jaiss/grpc-assingment/types"
	pb "github.com/adarsh-jaiss/grpc-assingment/types/__connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type TwitterServiceServer struct {
	pb.UnimplementedTwitterServiceHandler
}

func NewTwitterServiceServer() *TwitterServiceServer {
	return &TwitterServiceServer{}
}

func (s *TwitterServiceServer) GetUser(ctx context.Context, req *connect.Request[__.UserRequest]) (*connect.Response[__.UserResponse], error) {
	userResponse := &connect.Response[__.UserResponse]{
		Msg: &__.UserResponse{
			Username: "adarsh",
			Bio:      "I am a software engineer",
		},
	}
	return userResponse, nil
}

func (s *TwitterServiceServer) GetTweets(ctx context.Context, req *connect.Request[__.TweetsRequest]) (*connect.Response[__.TweetsResponse], error) {
	tweets := []string{"Hello", "World!"}
	tweetsResponse := &connect.Response[__.TweetsResponse]{
		Msg: &__.TweetsResponse{
			Tweets: tweets,
		},
	}
	return tweetsResponse, nil
}

func (s *TwitterServiceServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for handling HTTP requests here
}

func main() {
	server := &TwitterServiceServer{} // Assuming you have defined your server struct as Server
	mux := http.NewServeMux()
	mux.Handle("/connectrpc.com/connect.TwitterService/", server)
	log.Fatal(http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{})))
}
