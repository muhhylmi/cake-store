# Cake-Store
Hello this is Cake Store API build with golang 

## How to Run The Project
- install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/database/postgres) using go install -tags 'database1,database2' github.com/golang-migrate/migrate/v4/cmd/migrate@latest 
    - example: `go install -tags 'postgres,mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` 

- ### Run Database Migrations
    - Create Database
    - run `migrate -database "mysql://user:password@tcp(host:port)/db-name" -path migration_path up`
- run `go mod tidy`
- run `go run main.go`

## Run Project Using Docker Container
- Create mysql docker container as the db `docker pull (mysql_image)`
- Create network `docker network create {network_name}`
- Assignt db container to network `docker connect {network_name} {container_name}`
- create golang app iamge `docker build -t {image_name} .`
- create golang app container with env and network `docker run -d -p 3000:3000 --name {container_name} --env-file {env_file_name} --network {network_name}{image_name}`
- access the API using [http://localhost:3000](http://localhost:3000/)

## List Cake

Endpoint : GET /api/cake/

Headers :
- api-key: token

Response Body :

```json
{
    "data": [
        {
            "id": 1,
            "title": "Lemon cheesecake",
            "description": "A cheesecake made of lemon",
            "rating": 7,
            "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
            "created_at": "2020-02-01 10:56:31",
            "updated_at": "2020-02-13 09:30:23"
        }
    ],
    "status": "OK",
    "code": 200
}
```


## Get Detail Cake

Endpoint : GET /api/cake/:cakeId

Headers :
- api-key: token

Response Body :

```json
{
    "data" : {
        "id": 1,
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2020-02-01 10:56:31",
        "updated_at": "2020-02-13 09:30:23"
    },
    "status": "OK",
    "code": 200
}
```

## Create Cake

Endpoint : POST /api/cake

Headers :
- api-key: token

Request Body :

```json
{
    "data": {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
}
```

Response Body :

```json
{
    "code": 201,
    "status": "OK",
    "data": {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
}
```


## Update Cake

Endpoint : PATCH /api/cake/:cakeId

Headers :
- api-key: token

Request Body :

```json
{
    "data": {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
}
```

Response Body :

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
}
```


## Delete Cake

Endpoint : DELETE /api/cake/:cakeId

Headers :
- api-key: token

Request Body :

Response Body :

```json
{
    "code": 200,
    "status": "OK",
    "data" : "boolean"
}
```