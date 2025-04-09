package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	pb "github.com/maniteja0126/paper-social/post-service/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPostServiceServer
	posts map[string][]*pb.Post
}

func (s *server) ListPostByUser(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostResponse, error) {
	userPosts, ok := s.posts[req.UserId]
	if !ok {
		return &pb.ListPostResponse{Post: []*pb.Post{}}, nil
	}
	return &pb.ListPostResponse{Post: userPosts}, nil
}

func generateFakePosts() map[string][]*pb.Post {
	gofakeit.Seed(0)
	users := []string{"user1", "user2", "user3", "user4", "user5"}
	posts := make(map[string][]*pb.Post)
	var globalID int64 = 1
	now := time.Now().Unix()

	for _, user := range users {
		n := rand.Intn(3) + 3
		var userPosts []*pb.Post
		for i := 0; i < n; i++ {
			userPosts = append(userPosts, &pb.Post{
				Id:        fmt.Sprintf("%d", globalID),
				Author:    user,
				Content:   gofakeit.Sentence(8),
				Timestamp: now - int64(rand.Intn(10000)),
			})
			globalID++
		}
		posts[user] = userPosts
	}

	return posts
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{
		posts: generateFakePosts(),
	})
	log.Println("Post service is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
