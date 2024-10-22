# Go Boilerplate API

> Attempt to implement an todo list API follow [The 12 Factor App](https://12factor.net/)

## Dependencies

1. [Cobra - Command-line Interface Tools](https://github.com/spf13/cobra): build the admin command
   * Use [cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md) to generate cobra command is more efficient
2. [Viper - Configuration](https://github.com/spf13/viper): manage different format of configuration
3. [Gin_Gonic - HTTP Server](https://github.com/gin-gonic/gin): the web framework
4. [GORM - ORM](https://gorm.io/gorm): abstract different type of relational database
5. [Logrus - Logging](https://github.com/sirupsen/logrus): logging tool
6. [Testify - Testing](https://github.com/stretchr/testify): assertion and mocking framework

:bulb: the order to `go get` those packages matters

## Command

### Local Development

```bash
docker-compose up -d
go run main.go server
```

### Build and Run

```bash
go build -o api
api --help # check the admin command for this server
```

## Notes

* Fix go.mod with `go mod tidy` command