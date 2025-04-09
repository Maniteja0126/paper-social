# 📰 Paper.Social - Timeline Microservice

This is a simplified backend microservice for a social feed system like Paper.Social, designed to showcase a real-world use of **Go**, **GraphQL**, and **gRPC**.

---

## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/Maniteja0126/paper-social.git
cd paper-social
```

### 2. Docker Setup

All Docker-related files are located in the docker/ folder.

```bash
cd docker
docker-compose up --build
```
This starts both:

- Feed Service (GraphQL) at ```http://localhost:8080```

- Post Service (gRPC) used internally by the feed service 

##
## 🚀 How to Run the Service (without Docker)

### 1. Run the gRPC Post Service
```bash
cd post-service
go run server.go
```

### 2. In a new terminal, run the Feed Service
```bash
cd feed-service
go run server.go
```


## Sample GraphQL Query

Visit ```http://localhost:8080``` and try this query:

```graphql
query {
  getTimeline(userId: "user1") {
    id
    author
    content
    timestamp
  }
}
```
This will return the most recent posts from users followed by "user1", sorted in reverse chronological order.

###  Expected Output Example (GraphQL response)

```bash
{
  "data": {
    "getTimeline": [
      {
        "id": "9",
        "author": "user3",
        "content": "Behind give enough away ability hand for several.",
        "timestamp": 1744184065
      }
    ]
  }
}
```

##  Description of Approach

### Objective
Deliver a timeline by aggregating posts from users that a given user follows.

###  Implementation Highlights
- GraphQL API with ```getTimeline(userId: ID!)``` query
- gRPC Post Service with:
 ```bash
     rpc ListPostsByUser(ListPostsRequest) returns (ListPostsResponse)
```
- In-Memory Data Models simulate:
    - Users & their follower relationships
    - Posts with unique IDs, timestamps, content, and authors

## 
### 📁 Folder Structure
```bash
paper-social/
├── docker/              # Dockerfiles and docker-compose.yml
├── feed-service/        # GraphQL feed service
├── post-service/        # gRPC post service
```
