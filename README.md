# triadmoko-be-golang

Documentation API 

**Stack App**:

- Golang
- Gorm
- Gin
- Postgres

## Global Response

```json
{
  "message": "Register Succesfully ",
  "code": 200,
  "status": "success",
  "data": {....}
}
```

## User End Point

### Register 
Method `POST` Register 

#### Endpoint
```
https://test-golang-triadmoko.herokuapp.com/api/v1/user/register
```

#### Body
```json
{
    "email":"triadmoko@gmail.com",
    "firstname":"asdfsf",
    "lastname":"sdffgad",
    "username":"sdgdfg",
    "address":"jhgaiugrpu",
    "phone": 5676345,
    "password":"qwerty123",
    "confirm_password":"qwerty123"
}
```
#### Response
```json
{
    "message": "Register Succesfully ",
    "code": 200,
    "status": "success",
    "data": {
        "email": "triadmoko12@gmail.com",
        "firstname": "triadmoko",
        "lastname": "denny fatrosa",
        "username": "triadmoko",
        "address": "batam",
        "phone": 821238741
    }
}
```

#### Example Request Go Native
```go
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://test-golang-triadmoko.herokuapp.com/api/v1/user/register"
  method := "POST"

  payload := strings.NewReader(`{
    "email":"triadmoko@gmail.com",
    "firstname":"asdfsf",
    "lastname":"sdffgad",
    "username":"sdgdfg",
    "address":"jhgaiugrpu",
    "phone": 5676345,
    "password":"qwerty123",
    "confirm_password":"qwerty123"
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```


### Login
Method `Post` Login

Time Expire Login 30 Minute

**EndPoint Login**
```
https://test-golang-triadmoko.herokuapp.com/api/v1/user/login
```

#### Body
```json
{
    "email":"triadmoko@gmail.com",
    "password":"qwerty123"
}
```

#### Response
```json
{
    "message": "Login Success",
    "code": 200,
    "status": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.GpWcp44Gk6-R5edBZD7aEmk1ZjvmoCLs8PId057RS58"
    }
}
```

#### Example Request Go Native

```go
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://test-golang-triadmoko.herokuapp.com/api/v1/user/login"
  method := "POST"

  payload := strings.NewReader(`{
    "email":"triadmoko@gmail.com",
    "password":"qwerty123"
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```

## Faskes Endpoint

### Create Faskes
Method `Post`

**Endpoit Create Faskes**
```
https://test-golang-triadmoko.herokuapp.com/api/v1/faskes/create
```

**Request Headers**
```text
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgzOTUwMzksInVzZXJfaWQiOjJ9.vYSBM_sk0bvmmPsVnvN4gTmh8-v6nwr4ccNfcV2MiVA
```

**Body**
```json
{
    "name":"Klinik"
}
```

**Response**
```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": {
        "id": 4,
        "name": "Klinik"
    }
}
```
### Get Data Faskes
Method Get

**Endpoint Get Data Faskes**

```
https://test-golang-triadmoko.herokuapp.com/api/v1/faskes/
```

**Request Headers**
```text
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgzOTUwMzksInVzZXJfaWQiOjJ9.vYSBM_sk0bvmmPsVnvN4gTmh8-v6nwr4ccNfcV2MiVA
```

**Body**
```json
{
    "name":"Klinik"
}
```

**Response **
```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": {
        "id": 4,
        "name": "Klinik"
    }
}
```

## Nakes 
### Create Nakes
Method `POST` Nakes

```
https://test-golang-triadmoko.herokuapp.com/api/v1/nakes/create
```

**Request Headers**
```text
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgzOTUwMzksInVzZXJfaWQiOjJ9.vYSBM_sk0bvmmPsVnvN4gTmh8-v6nwr4ccNfcV2MiVA
```

**Body**
```json
{
    "jumlah_nakes": 1223,
    "name": "triadmoko",
    "address":"dfas4wefwef",
    "faskes_id":1
}
```
**Response**
```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": {
        "jumlah_nakes": 1223,
        "name": "triadmoko",
        "address": "dfas4wefwef",
        "faskes_id": 1
    }
}
```
### Update Nakes
Method `PUT` 
Endpoint 

```text
https://test-golang-triadmoko.herokuapp.com/api/v1/nakes/update/1
```

**Request Headers**
```text
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgzOTUwMzksInVzZXJfaWQiOjJ9.vYSBM_sk0bvmmPsVnvN4gTmh8-v6nwr4ccNfcV2MiVA
```

**Body**
```json
{
    "jumlah_nakes": 122,
    "name": "triadqmokao",
    "address":"dfas4wefwef",
    "faskes_id":1
}
```
**Response**
```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": {
        "jumlah_nakes": 122,
        "name": "triadqmokao",
        "address": "dfas4wefwef",
        "faskes_id": 1
    }
}
```
### Get All Nakes
Method `GET`

Endpoint Get All Nakes

```text
https://test-golang-triadmoko.herokuapp.com/api/v1/nakes/
```


**Request Headers**

```text
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgzOTUwMzksInVzZXJfaWQiOjJ9.vYSBM_sk0bvmmPsVnvN4gTmh8-v6nwr4ccNfcV2MiVA
```

**Response**
```json
{
    "message": "Get Nakes Success",
    "code": 200,
    "status": "success",
    "data": [
        {
            "jumlah_nakes": 122,
            "name": "triadqmokao",
            "address": "dfas4wefwef",
            "faskes_id": 1
        },
        {
            "jumlah_nakes": 1223,
            "name": "triadmoko",
            "address": "dfas4wefwef",
            "faskes_id": 1
        }
    ]
}
```

### Delete Nakes
Method `Delete`

Endpoint 
```
https://test-golang-triadmoko.herokuapp.com/api/v1/nakes/delete/1
```

```text
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgzOTUwMzksInVzZXJfaWQiOjJ9.vYSBM_sk0bvmmPsVnvN4gTmh8-v6nwr4ccNfcV2MiVA
```
**Response**
```json
{
    "message": "Delete Success",
    "code": 200,
    "status": "success",
    "data": "Data Success Delete"
}
```
## Get PDF

Method `Get`
Endpoint
```
localhost:8080/api/v1/nakes/pdf
```

```text
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgzOTUwMzksInVzZXJfaWQiOjJ9.vYSBM_sk0bvmmPsVnvN4gTmh8-v6nwr4ccNfcV2MiVA
```
**Response**
```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": "Download pdf di link berikut : 'localhost:8080/pdf/file.pdf'"
}
```

### Show PDF
Method `Get`
```
localhost:8080/pdf/file.pdf
```

**Response**

![Response Show PDF](pdf.png)
