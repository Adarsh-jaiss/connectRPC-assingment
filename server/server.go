package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	connect "connectrpc.com/connect"
	__ "github.com/adarsh-jaiss/grpc-assingment/types"
	pb "github.com/adarsh-jaiss/grpc-assingment/types/typesconnect"
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
	mux.Handle("/getUser", server)
	mux.Handle("/getTweets", server)
	listenAddr := flag.String("listenAddr", ":8080", "The address to listen on for gRPC requests.")
	flag.Parse()
	fmt.Printf("running a new server at port : %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, h2c.NewHandler(mux, &http2.Server{})))
}
