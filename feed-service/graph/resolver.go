package graph

import "feed-service/graph/model"

var localUserFollows = map[string][]string{
	"user1": {"user2", "user3"},
	"user2": {"user1"},
	"user3": {"user1", "user2"},
}

func (r *Resolver) GetUserFollows(userID string) []string {
	return localUserFollows[userID]
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type Resolver struct {
	followMap  map[string][]string
	grpcClient interface {
		ListPostByUser(userID string) []*model.Post
	}
}

type queryResolver struct{ *Resolver }
