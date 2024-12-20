# Go Boilerplate API

> Attempt to implement an todo list API follow [The 12 Factor App](https://12factor.net/)

## Dependencies

1. [Cobra - Command-line Interface Tools](https://github.com/spf13/cobra): build the admin command
   * [User Guideline](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)
   * Use [cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md) to generate cobra command is more efficient
2. [Viper - Configuration](https://github.com/spf13/viper): manage different format of configuration
3. [Gin_Gonic - HTTP Server](https://github.com/gin-gonic/gin): the web framework
4. [GORM - ORM](https://gorm.io/gorm): abstract different type of relational database
5. [zap - Logging](https://github.com/uber-go/zap): logging tool
6. [Testify - Testing](https://github.com/stretchr/testify): assertion and mocking framework
7. [swag - OpenAPI](https://github.com/swaggo/gin-swagger): generate swagger document
   * [How to comment API document to generate document](https://github.com/swaggo/swag?tab=readme-ov-file#general-api-info)
8. [Resty](https://github.com/go-resty/resty): http client package. much better than `net/http`
9. [Hot Reload](https://github.com/air-verse/air): auto reload gin app code change to make developer's live easier
10. [Graceful Shutdown - Github Example](https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go)

:bulb: the order to `go get` those packages matters

## Command

### Local Development

```bash
docker-compose up -d

# please install `air` first to hot reload while code change
go install github.com/air-verse/air@latest
air
```

### Build and Run

```bash
go build -o api
api --help # check the admin command for this server
```

```bash
docker build -t my-golang-webapp:latest .
docker run --env GIN_MODE=release \
   --name api \
   --mount type=bind,src=/Users/henry.chou/sources/go-boilerplate-api/configs/local.yml,dst=/app/configs/local.yml \
   --network=host \
   -p 8081:8081 \
   my-golang-webapp:latest
docker rm $(docker container ls --filter "ancestor=my-golang-webapp" -aq | head -n 1)
```

### Build open api document

```bash
# generate swagger json and yaml automatically
swag init -o api

# format swagger comment
swag fmt
```

## Notes

* Fix go.mod with `go mod tidy` command