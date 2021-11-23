package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "database/sql"
    "strconv"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "Tillet";
    dbname = "tillet_coin"
)
  
// User represents a users data
type userstruct struct {
    ID string `json:"ID"`
    FName string `json:"FName"`
    LName string `json:"LName"`
    Wealth int `json:"Wealth"`
    Class1 string `json: Class1`
    Class2 string `json: Class2`
    Class3 string `json: Class3`
    Class4 string `json: Class4`
    IsTeacher bool `json: IsTeacher`

}

// Gets user data from the database when given an ID 
func GetUserData(ID string) [9]string{
    var usrdata[9] string
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }  
    err = db.Ping()
    if err != nil {
        panic(err)
    }

    var sqlStatement string = `SELECT id, first_name, last_name, wealth, class1, class2, class3, class4, isteacher FROM users WHERE id=$1;`
    var id string
    var first_name string
    var last_name string
    var wealth int
    var class1 sql.NullString
    var class2 sql.NullString
    var class3 sql.NullString
    var class4 sql.NullString
    var isteacher bool
    row := db.QueryRow(sqlStatement, ID)
    switch err := row.Scan(&id, &first_name, &last_name, &wealth, &class1, &class2, &class3, &class4, &isteacher); err {
        case sql.ErrNoRows:
            fmt.Println("No rows found")
        case nil:
            usrdata[0] = id
            usrdata[1] = first_name
            usrdata[2] = last_name
            usrdata[3] = strconv.Itoa(wealth)
            if class1.Valid {usrdata[4] = class1.String} else {usrdata[4] = "null"}
            if class2.Valid {usrdata[5] = class1.String} else {usrdata[5] = "null"}
            if class3.Valid {usrdata[6] = class1.String} else {usrdata[6] = "null"}
            if class4.Valid {usrdata[7] = class1.String} else {usrdata[7] = "null"}
            if isteacher {usrdata[8] = "true"} else {usrdata[8] = "false"}
        
        default:
            panic(err)
    }
    return usrdata
}

func main() {
	router := gin.Default()

    //Get info for a user given their ID
    router.GET("/user/:id", func(c *gin.Context) {
        ID := c.Param("id")
        UserData := GetUserData(ID)
        Wealthint, err := strconv.Atoi(UserData[3])
        fmt.Println(err)
            if UserData[8] == "false" {
                c.IndentedJSON(http.StatusOK, []userstruct {
                    {ID: UserData[0], FName: UserData[1], LName: UserData[2], Wealth: Wealthint, Class1: UserData[4],Class2: UserData[5],Class3: UserData[6],Class4: UserData[7], IsTeacher: false},
                })} else {
                c.IndentedJSON(http.StatusOK, []userstruct {
                    {ID: UserData[0], FName: UserData[1], LName: UserData[2], Wealth: Wealthint, Class1: UserData[4],Class2: UserData[5],Class3: UserData[6],Class4: UserData[7], IsTeacher: true},
                })
            }
    })

    router.GET("/user", func(c *gin.Context) {
        c.String(http.StatusOK, "You forget the ID")
    })

    fmt.Println("Starting...")
    // TODO: Use net.InterfaceAddrs() func from "net" to automatically change IP
    router.Run("192.168.0.71:8080")
}