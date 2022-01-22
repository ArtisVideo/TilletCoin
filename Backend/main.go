package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	externalip "github.com/glendc/go-external-ip"
	_ "github.com/lib/pq"
)

// https://gobyexample.com/command-line-arguments
const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "Tillet"
	dbname   = "tillet_coin"
)

// User represents a users data
type userstruct struct {
	ID        string `json:"ID"`
	FName     string `json:"FName"`
	LName     string `json:"LName"`
	Wealth    int    `json:"Wealth"`
	Class1    string `json: Class1`
	Class2    string `json: Class2`
	Class3    string `json: Class3`
	Class4    string `json: Class4`
	IsTeacher bool   `json: IsTeacher`
}

func GetInteralIP(port string) string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String() + ":" + port
}

func GetExternalIP(port string) string {
	consensus := externalip.DefaultConsensus(nil, nil)
	ip, err := consensus.ExternalIP()
	if err != nil {
		panic(err)
	}

	return ip.String() + ":" + port
}

// Gets user data from the database when given an ID
func GetUserData(ID string) [9]string {
	var usrdata [9]string
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
		if class1.Valid {
			usrdata[4] = class1.String
		} else {
			usrdata[4] = "null"
		}
		if class2.Valid {
			usrdata[5] = class1.String
		} else {
			usrdata[5] = "null"
		}
		if class3.Valid {
			usrdata[6] = class1.String
		} else {
			usrdata[6] = "null"
		}
		if class4.Valid {
			usrdata[7] = class1.String
		} else {
			usrdata[7] = "null"
		}
		if isteacher {
			usrdata[8] = "true"
		} else {
			usrdata[8] = "false"
		}

	default:
		panic(err)
	}
	return usrdata
}

func main() {
	router := gin.Default()

	//Get info for a user given their ID
	router.GET("/user/getdata/:id", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		ID := c.Param("id")
		UserData := GetUserData(ID)
		Wealthint, err := strconv.Atoi(UserData[3])
		if err != nil {
			panic(err)
		}
		if UserData[8] == "false" {
			c.IndentedJSON(http.StatusOK, []userstruct{
				{ID: UserData[0], FName: UserData[1], LName: UserData[2], Wealth: Wealthint, Class1: UserData[4], Class2: UserData[5], Class3: UserData[6], Class4: UserData[7], IsTeacher: false},
			})
			return
		} else {
			c.IndentedJSON(http.StatusOK, []userstruct{
				{ID: UserData[0], FName: UserData[1], LName: UserData[2], Wealth: Wealthint, Class1: UserData[4], Class2: UserData[5], Class3: UserData[6], Class4: UserData[7], IsTeacher: true},
			})
			return
		}
	})

	router.POST("/user/gatherdata/:JWT", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		JWT := c.Param("JWT")
		// TODO: Verify JWT
		fmt.Println(JWT)
		// s := strings.SplitN(JWT, ".", 3)
		//fmt.Println(s[2])
	})

	//Check if a ID is in the DB
	router.GET("/user/isvalid/:id", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		ID := c.Param("id")
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
		row := db.QueryRow(`SELECT ID FROM users WHERE id=$1;`, ID)
		switch err := row.Scan(&ID); err {
		case sql.ErrNoRows:
			c.Status(http.StatusNoContent)
			return
		case nil:
			c.Status(http.StatusOK)
			return
		}
	})

	router.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "You forget the ID")
	})
	fmt.Println("Starting...")

	var IsReserved string
	var port string

	fmt.Println("Do you want to use an internal or external address? (I/E)")
	fmt.Scanln(&IsReserved)

	fmt.Println("Please enter a port number (Deafult: 8080)")
	fmt.Scanln(&port)
	if port == "" {
		port = "8080"
	}

	switch strings.ToUpper(IsReserved) {
	case "I":
		router.Run(GetInteralIP(port))
	case "E":
		router.Run(GetExternalIP(port))
	default:
		fmt.Println("Please enter I/E")
	}
}
