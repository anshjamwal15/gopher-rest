
# Authorization+Authentication service in Golang

Created this project for HousewareHQ-github-octernship. My design decision was pretty simple as I created golang project for the first time. In the beginning I was pretty confused while thinking about folder structure for this project. So I just tried my best. Thanks. Ps :- Sorry for the Github Flow avoidation, was working on local-repo and commited the whole project on last day in this repo. 
+ See [PERSONAL REPO](https://github.com/aimbot1526)

### The kit provides the following features right out of the box:

* User Login
* Admin User adds a new User account(by providing the username & password)
* Admin User deletes an existing User account from their organization
* List all Users in their organization

## Used Technologies
* Go Fiber.
* GORM.
* PostgreSQL.

## Folder Structure

* [app/](./app)
  * [controllers/](./app/controllers)
    * [OrganizationController.go](./app/controllers/OrganizationController.go)
    * [UserController.go](./app/controllers/UserController.go)
  * [models/](./app/models)
    * [Organization.go](./app/models/Organization.go)
    * [User.go](./app/models/User.go)
    * [base.go](./app/models/base.go)
* [build/](./build)
  * [apiserver](./build/apiserver)
* [cmd/](./cmd)
  * [server/](./cmd/server)
    * [server.go](./cmd/server/server.go)
* [configs/](./configs)
  * [dbConfig.go](./configs/dbConfig.go)
  * [fiberConfig.go](./configs/fiberConfig.go)
* [docs/](./docs)
  * [docs.go](./docs/docs.go)
  * [swagger.json](./docs/swagger.json)
  * [swagger.yaml](./docs/swagger.yaml)
* [pkg/](./pkg)
  * [middleware/](./pkg/middleware)
    * [FIberMiddleware.go](./pkg/middleware/FIberMiddleware.go)
    * [JwtMiddleware.go](./pkg/middleware/JwtMiddleware.go)
  * [payload/](./pkg/payload)
    * [request/](./pkg/payload/request)
      * [CreateOrgRequest.go](./pkg/payload/request/CreateOrgRequest.go)
      * [CreateUserRequest.go](./pkg/payload/request/CreateUserRequest.go)
    * [response/](./pkg/payload/response)
      * [CreateOrgResponse.go](./pkg/payload/response/CreateOrgResponse.go)
      * [CreateUserResponse.go](./pkg/payload/response/CreateUserResponse.go)
      * [UserResponse.go](./pkg/payload/response/UserResponse.go)
  * [routes/](./pkg/routes)
    * [privateRoutes.go](./pkg/routes/privateRoutes.go)
    * [publicRoutes.go](./pkg/routes/publicRoutes.go)
    * [swaggerRoute.go](./pkg/routes/swaggerRoute.go)
  * [utils/](./pkg/utils)
    * [jwtGenerator.go](./pkg/utils/jwtGenerator.go)
    * [jwtParser.go](./pkg/utils/jwtParser.go)
    * [payload.go](./pkg/utils/payload.go)
    * [validator.go](./pkg/utils/validator.go)
  * [README.md](./pkg/README.md)
* [platform/](./platform)
  * [migrations/](./platform/migrations)
* [test/](./test)
  * [README.md](./test/README.md)
  * [create_org_test.go](./test/create_org_test.go)
  * [org_add_user_test.go](./test/org_add_user_test.go)
  * [org_delete_user_test.go](./test/org_delete_user_test.go)
  * [org_view_all_user_test.go](./test/org_view_all_user_test.go)
  * [org_view_user_test.go](./test/org_view_user_test.go)
  * [public_route_test.go](./test/public_route_test.go)
* [.env](./.env)
* [.gitignore](./.gitignore)
* [Makefile](./Makefile)
* [README.md](./README.md)
* [go.mod](./go.mod)
* [go.sum](./go.sum)
* [main.go](./main.go)


# Installation

## 1. Create Database and tables.
```bash
su -u postgres psql

CREATE DATABASE test
```
### Now execute code from sql files in platforms/migrations in psql shell.

## 2. Install packages

```bash
go mod download
```
## 3. Install make.

```bash
sudo apt install make -y
```
## 4. Run build with test cases.

```bash
make build
```

## 5. Run binary.

```bash
make run
```
## Getting Started

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `POST /api/v1/user`: Register a user in database.
* `POST /api/v1/login`: authenticates a user and generates a JWT.
* `GET /api/v1/view/:userid`: view user 
* `GET /api/v1/all/:userid`: get all user list in org.
* `POST /api/v1/create`: creates new org.
* `POST /api/v1/add`: creates new user.
* `DELETE /api/v1/delete/:userid/:orgid`: delete user.

## Test
Test APIs with Swagger.
* `GET /swagger/`: get all user list in org.

# Database design
## Organizations
```
| id            | name          | created_by | created_at | updated_at |
| ------------- |:-------------:| ----------:|-----------:|-----------:|
```
## Users
```
| id            | username      | password   | created_at | updated_at | role |
| ------------- |:-------------:| ----------:|-----------:|-----------:|-----:|
```
## Org_users
```
| organization_id | user_id |
| ----------------|:-------:|
```
