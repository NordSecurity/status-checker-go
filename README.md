# Status checker

## Description
Simple status checker which allows you to check if your services are healthy.

## Usage

### Basic

Create status checker pool with your service checkers:

```Go
package main

import (
	"database/sql"
	"fmt"

	"github.com/NordSec/status-checker-go"
	"github.com/NordSec/status-checker-go/checker"
	"github.com/NordSec/status-checker-go/maintenance"
	"github.com/go-redis/redis/v7"
)

func check(mysqlDB *sql.DB, redisClient redis.UniversalClient, maintenanceClient maintenance.Client) {
	statusCheckerPool := status.NewCheckerPool(
		checker.NewMysqlChecker(mysqlDB),
		checker.NewRedisChecker(redisClient),
		checker.NewMaintenanceChecker(maintenanceClient),
	)

	fmt.Println("Status:", statusCheckerPool.Status())
	fmt.Println("Details:", statusCheckerPool.Details())
}
```

### Handlers

#### Echo handler

```Go
package main

import (
	"github.com/NordSec/status-checker-go"
	"github.com/NordSec/status-checker-go/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	statusCheckerPool := status.NewCheckerPool(
		// checkers
	)

	echoEngine := echo.New()

	echoHandler := handler.NewEchoHandler(statusCheckerPool)

	echoEngine.GET("/status", echoHandler.Status)
	echoEngine.GET("/status/details", echoHandler.Details)

	echoEngine.Logger.Fatal(echoEngine.Start(":8080"))
}
```

#### Gin handler

```Go
package main

import (
	"log"
	
	"github.com/NordSec/status-checker-go"
	"github.com/NordSec/status-checker-go/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	statusCheckerPool := status.NewCheckerPool(
		// checkers
	)

	ginEngine := gin.Default()

	ginHandler := handler.NewGinHandler(statusCheckerPool)

	ginEngine.GET("/status", ginHandler.Status)
	ginEngine.GET("/status/details", ginHandler.Details)

	log.Fatal(ginEngine.Run())
}
```

## Custom checker
Custom checker can be created by implementing `status.Checker` interface:

```Go
package checker

import (
	"github.com/NordSec/status-checker-go"
)

func NewExampleChecker() status.Checker {
	return &exampleChecker{}
}

type exampleChecker struct {}

func (ec *exampleChecker) Name() string {
	return "example"
}

func (ec *exampleChecker) Status() status.Status {
	return status.OK
}
```
