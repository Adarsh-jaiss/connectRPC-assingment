package main

import (
	"context"
	"fmt"
	"log"
	pb "github.com/adarsh-jaiss/grpc-assingment/types"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTwitterServiceClient(conn)
	// Call GetUser RPC
	userResponse, err := client.GetUser(context.Background(), &pb.UserRequest{UserId: "1"})
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}
	fmt.Println("User Response:", userResponse)

	// Call GetTweets RPC
	tweetsResponse, err := client.GetTweets(context.Background(), &pb.TweetsRequest{Hashtag: "internship"})
	if err != nil {
		log.Fatalf("Error calling GetTweets: %v", err)
	}
	fmt.Println("Tweets Response:", tweetsResponse)

}