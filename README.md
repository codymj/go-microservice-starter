# go-microservice-starter

My simple implementation of what a Go microservice should consist of.

## `/cmd/app`

This directory contains the application configuration logic as well as logic for
setting up routes.

#### `/main.go`

* initializes the application's configuration parameters
* initializes a database connection
* initializes any internal services
* initializes http routes
* starts the application

#### `/config/config.go`

* sets the application's configuration parameters
* sets logger configuration parameters
* sets database configuration parameters
* creates service instances

#### `/routes/routes.go`

* sets up http router
* connects route handlers to service logic

#### `/routes/posthello.go`

* an example of a route handler
* file name follows the format (note, `/api/v1` is ignored):

| file name        | http endpoint                   |
|------------------|---------------------------------|
| getusers         | `GET /api/v1/users`             |
| getusersid       | `GET /api/v1/users/{id}`        |
| postusers        | `POST /api/v1/users`            |
| putusersiddevice | `PUT /api/v1/users/{id}/device` |
| deleteusersid    | `DELETE /api/v1/users/{id}`     |

#### `/routes/structs.go`

* structs used in the `routes` package

#### `/routes/util.go`

* utility logic used in the `routes` package
