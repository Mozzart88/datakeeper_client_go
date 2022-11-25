# Data Keeper Client for DiliApi

Create new DataKeeper Client

```go
dkUrl := "http://dkc/"
apiVersion := "0.0"
dkc := dkc.NewClient(dkUrl, apiVersion)
dkc.Secret := "very_strong_secret"
dkc.UserAgent := "dkc_allowed_user_agent"

```
Creates singleton

To get singleton
```go
dkc := dkc.Get()
```

## Send requests to DataKeeper
### Add data
```go
import "diliapi.com/api"

entity := "User"
data := api.JsonMap{
	"fields" {
        "name": "tom",
        "email": "hanks@ex.com",
        "role": "actor"
    }
}

client := dkc.Get()
dkc := client.Add(entity, data)
```
#### Returns
__Success__
```json
{
    "status": 200,
    "body": {
        "msg": "success"
    }
}
```

### Get data

```go
import "diliapi.com/api"

entity := "User"
data := api.JsonMap{
    "fields" [
        "name",
        "email",
        "role",
    ],
    "where": {
        "uid": "3edfr45tyh"
    }
}

client := dkc.Get()
dkc := client.Get(entity, data)
```
#### Returns
__Success__
```json
{
    "status": 200,
    "body": {
        "name": "tom",
        "email": "hanks@ex.com",
        "role": "actor"
    }
}
```


### Update data

```go
import "diliapi.com/api"

entity := "User"
data := api.JsonMap{
	"new" {
        "role": "legend",
    },
	"where": {
		"uid": "3edfr45tyh"
	}
}

client := dkc.Get()
dkc := client.Update(entity, data)
```
#### Returns
__Success__
```json
{
    "status": 200,
    "body": {
        "msg": "success"
    }
}
```

### Delete data
```go
import "diliapi.com/api"

entity := "User"
data := api.JsonMap{
	"where": {
        "uid": "5tgbhu8ik"
    }
}

client := dkc.Get()
dkc := client.Delete(entity, data)
```
#### Returns
__Success__
```json
{
    "status": 200,
    "body": {
        "msg": "success"
    }
}
```



## Error responses
```json
{
    "status": 401,
    "body": {
        "msg": "unauthorized"
    }
}
```
```json
{
    "status": 403,
    "body": {
        "msg": "forbidden"
    }
}
```
```json
{
    "status": 404,
    "body": {
        "msg": "not found"
    }
}
```
```json
{
    "status": 500,
    "body": {
        "msg": "internal error"
    }
}
```
