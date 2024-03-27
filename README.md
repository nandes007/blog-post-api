# Blog Post RESTful API
## Overview

This project provides a RESTful API for managing blog posts. It is built using Go for the backend, Docker for containerization, Docker Compose for orchestration, Go Migrate for database migrations, and PostgreSQL as the database.

## Installation

### Prerequisites

Before installing the Image Converter, ensure you have the following prerequisites installed:

Make sure you have the following installed:
- [Install Go](https://go.dev/dl/)
- [Install Docker](https://docs.docker.com/desktop/)
- Docker Compose
- Git

### Setting Up the Application

1. Clone the repository:
```
git clone https://github.com/nandes007/blog-post-api.git
```

2. Build and run project using docker compose
```
cd blog-post-api       # Enter to project

docker compose up -d     # Run project in background using docker compose
```

### Usage
### API Endpoints
- GET /posts: Get all blog posts
- GET /posts/{id}: Get a single blog post by ID
- POST /posts: Create a new blog post
- PUT /posts/{id}: Update an existing blog post by ID
- DELETE /posts/{id}: Delete a blog post by ID

### Authentication
The API uses token-based authentication. To authenticate, include a valid token in the Authorization header of your requests.

