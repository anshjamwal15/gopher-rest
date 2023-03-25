
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

* [app/](./backend/app)
  * [controllers/](./backend/app/controllers)
    * [OrganizationController.go](./backend/app/controllers/OrganizationController.go)
    * [UserController.go](./backend/app/controllers/UserController.go)
  * [models/](./backend/app/models)
    * [Organization.go](./backend/app/models/Organization.go)
    * [User.go](./backend/app/models/User.go)
    * [base.go](./backend/app/models/base.go)
* [build/](./backend/build)
  * [apiserver](./backend/build/apiserver)
* [cmd/](./backend/cmd)
  * [server/](./backend/cmd/server)
    * [server.go](./backend/cmd/server/server.go)
* [configs/](./backend/configs)
  * [dbConfig.go](./backend/configs/dbConfig.go)
  * [fiberConfig.go](./backend/configs/fiberConfig.go)
* [docs/](./backend/docs)
  * [docs.go](./backend/docs/docs.go)
  * [swagger.json](./backend/docs/swagger.json)
  * [swagger.yaml](./backend/docs/swagger.yaml)
* [pkg/](./backend/pkg)
  * [middleware/](./backend/pkg/middleware)
    * [FIberMiddleware.go](./backend/pkg/middleware/FIberMiddleware.go)
    * [JwtMiddleware.go](./backend/pkg/middleware/JwtMiddleware.go)
  * [payload/](./backend/pkg/payload)
    * [request/](./backend/pkg/payload/request)
      * [CreateOrgRequest.go](./backend/pkg/payload/request/CreateOrgRequest.go)
      * [CreateUserRequest.go](./backend/pkg/payload/request/CreateUserRequest.go)
    * [response/](./backend/pkg/payload/response)
      * [CreateOrgResponse.go](./backend/pkg/payload/response/CreateOrgResponse.go)
      * [CreateUserResponse.go](./backend/pkg/payload/response/CreateUserResponse.go)
      * [UserResponse.go](./backend/pkg/payload/response/UserResponse.go)
  * [routes/](./backend/pkg/routes)
    * [privateRoutes.go](./backend/pkg/routes/privateRoutes.go)
    * [publicRoutes.go](./backend/pkg/routes/publicRoutes.go)
    * [swaggerRoute.go](./backend/pkg/routes/swaggerRoute.go)
  * [utils/](./backend/pkg/utils)
    * [jwtGenerator.go](./backend/pkg/utils/jwtGenerator.go)
    * [jwtParser.go](./backend/pkg/utils/jwtParser.go)
    * [payload.go](./backend/pkg/utils/payload.go)
    * [validator.go](./backend/pkg/utils/validator.go)
  * [README.md](./backend/pkg/README.md)
* [platform/](./backend/platform)
  * [migrations/](./backend/platform/migrations)
* [test/](./backend/test)
  * [README.md](./backend/test/README.md)
  * [create_org_test.go](./backend/test/create_org_test.go)
  * [org_add_user_test.go](./backend/test/org_add_user_test.go)
  * [org_delete_user_test.go](./backend/test/org_delete_user_test.go)
  * [org_view_all_user_test.go](./backend/test/org_view_all_user_test.go)
  * [org_view_user_test.go](./backend/test/org_view_user_test.go)
  * [public_route_test.go](./backend/test/public_route_test.go)
* [.env](./backend/.env)
* [.gitignore](./backend/.gitignore)
* [Makefile](./backend/Makefile)
* [README.md](./backend/README.md)
* [go.mod](./backend/go.mod)
* [go.sum](./backend/go.sum)
* [main.go](./backend/main.go)


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
## 3. Run build with test cases.

```bash
make build
```

## 4. Run binary.

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
