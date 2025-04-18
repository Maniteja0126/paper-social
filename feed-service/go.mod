module feed-service

go 1.24.1

require (
	github.com/99designs/gqlgen v0.17.70
	github.com/vektah/gqlparser/v2 v2.5.23
	google.golang.org/grpc v1.71.1
	github.com/maniteja0126/paper-social/post-service v0.0.0
)

replace github.com/maniteja0126/paper-social/post-service => ../post-service

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
