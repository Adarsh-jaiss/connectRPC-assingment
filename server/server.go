package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

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
    switch r.URL.Path {
    case "/user":
        s.handleGetUser(w, r)
    case "/tweets":
        s.handleGetTweets(w, r)
    default:
        http.NotFound(w, r)
    }
}


func (s *TwitterServiceServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
    userResponse, err := s.GetUser(r.Context(), nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse, err := json.Marshal(userResponse)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func (s *TwitterServiceServer) handleGetTweets(w http.ResponseWriter, r *http.Request) {
    tweetsResponse, err := s.GetTweets(r.Context(), nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jsonResponse, err := json.Marshal(tweetsResponse)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func main() {
	server := &TwitterServiceServer{} 
	mux := http.NewServeMux()
	path, handler := pb.NewTwitterServiceHandler(server)
	mux.Handle(path, handler)
	mux.Handle("/user", http.HandlerFunc(server.handleGetUser))
	mux.Handle("/tweets", http.HandlerFunc(server.handleGetTweets))
	listenAddr := flag.String("listenAddr", ":8080", "The address to listen on for gRPC requests.")
	flag.Parse()
	fmt.Printf("running a new server at port : %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, h2c.NewHandler(mux, &http2.Server{})))
}
