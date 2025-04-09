package graph

import (
	"context"
	"feed-service/graph/model"
	"log"
	"os"
	"sort"
	"sync"

	pb "github.com/maniteja0126/paper-social/post-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userFollows = map[string][]string{
	"user1": {"user2", "user3", "user4"},
	"user2": {"user1"},
	"user3": {"user1", "user2"},
	"user4": {"user1", "user2", "user3"},
	"user5": {"user1", "user2", "user3", "user4"},
}

func (r *queryResolver) GetTimeline(ctx context.Context, userID string) ([]*model.Post, error) {
	following, ok := userFollows[userID]
	if !ok {
		return []*model.Post{}, nil
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var allPosts []*model.Post
	errChan := make(chan error, len(following))

	for _, followedUser := range following {
		wg.Add(1)

		go func(user string) {
			defer wg.Done()

			addr := os.Getenv("POST_SERVICE_ADDR")

			conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("failed to connect: %v", err)
				errChan <- err
				return
			}
			defer conn.Close()
			client := pb.NewPostServiceClient(conn)

			resp, err := client.ListPostByUser(context.Background(), &pb.ListPostRequest{
				UserId: user,
			})

			if err != nil {
				errChan <- err
				return
			}

			var posts []*model.Post
			for _, p := range resp.Post {
				posts = append(posts, &model.Post{
					ID:        p.Id,
					Content:   p.Content,
					Author:    p.Author,
					Timestamp: int32(p.Timestamp),
				})
			}

			mu.Lock()
			allPosts = append(allPosts, posts...)
			mu.Unlock()
		}(followedUser)
	}
	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		for err := range errChan {
			log.Printf("error fetching posts: %v", err)
		}
	}

	sort.Slice(allPosts, func(i, j int) bool {
		return allPosts[i].Timestamp > allPosts[j].Timestamp
	})

	if len(allPosts) > 20 {
		allPosts = allPosts[:20]
	}

	return allPosts, nil

}
