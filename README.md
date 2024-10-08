
# Authorization+Authentication service in Golang

Created this project for HousewareHQ-github-octernship. My design decision was pretty simple as I created golang project for the first time. In the beginning I was pretty confused while thinking about folder structure for this project. So I just tried my best. Thanks. Ps :- Sorry for the Github Flow avoidation, was working on local-repo and commited the whole project on last day in this repo. 
+ See [PERSONAL REPO](https://github.com/aimbot1526/gopher-rest)

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

* [app/](./gopher-rest/app)
  * [controllers/](./gopher-rest/app/controllers)
    * [OrganizationController.go](./gopher-rest/app/controllers/OrganizationController.go)
    * [UserController.go](./gopher-rest/app/controllers/UserController.go)
  * [models/](./gopher-rest/app/models)
    * [Organization.go](./gopher-rest/app/models/Organization.go)
    * [User.go](./gopher-rest/app/models/User.go)
    * [base.go](./gopher-rest/app/models/base.go)
* [build/](./gopher-rest/build)
  * [apiserver](./gopher-rest/build/apiserver)
* [cmd/](./gopher-rest/cmd)
  * [server/](./gopher-rest/cmd/server)
    * [server.go](./cmd/server/server.go)
* [configs/](./gopher-rest/configs)
  * [dbConfig.go](./gopher-rest/configs/dbConfig.go)
  * [fiberConfig.go](./gopher-rest/configs/fiberConfig.go)
* [docs/](./gopher-rest/docs)
  * [docs.go](./gopher-rest/docs/docs.go)
  * [swagger.json](./gopher-rest/docs/swagger.json)
  * [swagger.yaml](./gopher-rest/docs/swagger.yaml)
* [pkg/](./gopher-rest/pkg)
  * [middleware/](./gopher-rest/pkg/middleware)
    * [FIberMiddleware.go](./gopher-rest/pkg/middleware/FIberMiddleware.go)
    * [JwtMiddleware.go](./gopher-rest/pkg/middleware/JwtMiddleware.go)
  * [payload/](./gopher-rest/pkg/payload)
    * [request/](./gopher-rest/pkg/payload/request)
      * [CreateOrgRequest.go](./gopher-rest/pkg/payload/request/CreateOrgRequest.go)
      * [CreateUserRequest.go](./gopher-rest/pkg/payload/request/CreateUserRequest.go)
    * [response/](./gopher-rest/pkg/payload/response)
      * [CreateOrgResponse.go](./gopher-rest/pkg/payload/response/CreateOrgResponse.go)
      * [CreateUserResponse.go](./gopher-rest/pkg/payload/response/CreateUserResponse.go)
      * [UserResponse.go](./gopher-rest/pkg/payload/response/UserResponse.go)
  * [routes/](./gopher-rest/pkg/routes)
    * [privateRoutes.go](./gopher-rest/pkg/routes/privateRoutes.go)
    * [publicRoutes.go](./gopher-rest/pkg/routes/publicRoutes.go)
    * [swaggerRoute.go](./gopher-rest/pkg/routes/swaggerRoute.go)
  * [utils/](./gopher-rest/pkg/utils)
    * [jwtGenerator.go](./gopher-rest/pkg/utils/jwtGenerator.go)
    * [jwtParser.go](./gopher-rest/pkg/utils/jwtParser.go)
    * [payload.go](./gopher-rest/pkg/utils/payload.go)
    * [validator.go](./gopher-rest/pkg/utils/validator.go)
  * [README.md](./gopher-rest/pkg/README.md)
* [platform/](./gopher-rest/platform)
  * [migrations/](./gopher-rest/platform/migrations)
* [test/](./gopher-rest/test)
  * [README.md](./gopher-rest/test/README.md)
  * [create_org_test.go](./gopher-rest/test/create_org_test.go)
  * [org_add_user_test.go](./gopher-rest/test/org_add_user_test.go)
  * [org_delete_user_test.go](./gopher-rest/test/org_delete_user_test.go)
  * [org_view_all_user_test.go](./gopher-rest/test/org_view_all_user_test.go)
  * [org_view_user_test.go](./gopher-rest/test/org_view_user_test.go)
  * [public_route_test.go](./gopher-rest/test/public_route_test.go)
* [.env](./gopher-rest/.env)
* [.gitignore](./gopher-rest/.gitignore)
* [Makefile](./gopher-rest/Makefile)
* [README.md](./gopher-rest/README.md)
* [go.mod](./gopher-rest/go.mod)
* [go.sum](./gopher-rest/go.sum)
* [main.go](./gopher-rest/main.go)


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
