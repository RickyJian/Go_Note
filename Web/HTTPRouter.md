# HTTPRouter

* 第三方 Router 處理器
* 解決 ServeMux URL 無法動態擷取的問題

[連結](https://github.com/julienschmidt/httprouter)

```go

package main

import (
	"net/http"

	"github.com/RickyJian/gorouterpractice/controllers"
	"github.com/julienschmidt/httprouter"
)

var userController = controllers.NewUserController()

func main() {
	// 開啟 HTTPRouter
	mux := httprouter.New()
	// GET 方法
	mux.GET("/user/:name", userController.Get)
	// POST 方法
	mux.POST("/user", userController.Create)
	// POST 方法
	mux.DELETE("/user/:name", userController.Remove)
	http.ListenAndServe("127.0.0.1:8080", mux)
}


```

## New()

開起 HTTPRouter

```go

package main 

import "github.com/julienschmidt/httprouter"

func main (){
    httprouter.New()
}

```

## Named Parameter

* URL 動態參數
* 使用 httprouter.Params


```go

// url: http://localhost:8080/uesr/:name


func (uc *UserController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // ByName: 字串對映 url (:name)
    name := p.ByName("name")

}

```

## GET()

HTTP GET 方法

```go

// package main

mux.GET("/user/:name", userController.Get)

// package controllers

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   p.ByName("name"),
		Gender: "male",
		Age:    18,
	}
	// JSON 編碼
	userJSON, _ := json.Marshal(u)
	// 設定回傳 Header
	w.Header().Set("Content-Type", "application/json")
	// 設定 HTTP Status
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", userJSON)
}

```

## POST()

HTTP POST 方法


```go

// package main

mux.POST("/user/:name", userController.Create)

// package controllers

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	// JSON解碼 body
	json.NewDecoder(r.Body).Decode(&u)

	u.Name = "Ricky"

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

```


## DELETE()

HTTP DELETE 方法

```go

// package main

mux.DELETE("/user", userController.Remove)

// package controllers

func (uc *UserController) Remove(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 刪除 邏輯
	w.WriteHeader(200)
}


```