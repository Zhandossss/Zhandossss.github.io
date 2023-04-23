package main

import (
    "database/sql"
    "html/template"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Age  int    `db:"age"`
}

var db *sqlx.DB

func main() {
    var err error
    db, err = sqlx.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
    if err != nil {
        panic(err)
    }

    router := gin.Default()
    router.GET("/", showUsers)
    router.GET("/edit/:id", editUserForm)
    router.POST("/edit/:id", editUser)
    router.Run(":8080")
}

func showUsers(c *gin.Context) {
    var users []User
    err := db.Select(&users, "SELECT * FROM users")
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    c.HTML(http.StatusOK, "users.tmpl", gin.H{
        "users": users,
    })
}

func editUserForm(c *gin.Context) {
    id := c.Param("id")

    var user User
    err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    c.HTML(http.StatusOK, "edit_user.tmpl", gin.H{
        "user": user,
    })
}

func editUser(c *gin.Context) {
    id := c.Param("id")
    name := c.PostForm("name")
    age := c.PostForm("age")

    _, err := db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", name, age, id)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    c.Redirect(http.StatusFound, "/")
}

func init() {
    var err error
    db, err = sqlx.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    template.Must(template.ParseFiles("users.tmpl", "edit_user.tmpl"))
}
