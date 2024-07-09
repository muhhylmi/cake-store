# Cake-Store
Hello this is Cake Store API build with golang 

## How to Run The Project
- soon

## List Cake

Endpoint : GET /api/cake/

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
    "message": "yeay, success retrieve data"
}
```


## Get Detail Cake

Endpoint : GET /api/cake/:cakeId

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
    }
}
```

## Create Cake

Endpoint : POST /api/cake

Request Body :

```json
{
  "title": "Lemon cheesecake",
  "description": "A cheesecake made of lemon",
  "rating": 7,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}
```

Response Body :

```json
{
  "data" : {
      "id": 1,
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
}
```


## Update Cake

Endpoint : PATCH /api/cake/:cakeId

Request Body :

```json
{
  "title": "Lemon cheesecake",
  "description": "A cheesecake made of lemon",
  "rating": 7,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}
```

Response Body :

```json
{
  "data" : {
      "id": 1,
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
}
```


## Delete Cake

Endpoint : DELETE /api/cake/:cakeId

Request Body :

Response Body :

```json
{
  "data" : {
    "success": boolean
  }
}
```