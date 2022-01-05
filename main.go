package main

import (
  "DEMOPROJECT/person"
  "DEMOPROJECT/setdatabase"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm/dialects/mysql"
  "log"
  "net/http"
  "os"
  "time"
)

func GetDetails(c *gin.Context) {
  router.POST("/login", Login)
  var users []user
    database.DB.Find(&users)
    c.JSON(http.StatusOK, gin.H{"data": users})
  }

func Login(c *gin.Context) {
	var u user
	if err := c.ShouldBindJSON(&u); err != nil {
	   c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
	   return
	}
  
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
	   c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	   return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
	   c.JSON(http.StatusUnprocessableEntity, err.Error())
	   return
	} 
	c.JSON(http.StatusOK, token)
 
  func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
  }
  type CreateUser struct {
    FirstName string `json:"firstName" binding:"required"`
	  LastName  string `json:"lastName" binding:"required"`
	  Age       int    `json:"age" binding:"required"`
	  Address	  string `json:"address binding:"required"`
	  JobInfo   map[string]string  `json:"joninfo" binding:"required"`
	  username    int    `json:"logid" binding:"required"`
	  Password  int	 `json:"password" binding:"required"`
    }


  func Createuser(c *gin.Context) {
    // Validate input
    var input CreateUser
    if err := c.ShouldBindJSON(&input); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
      // Create user
    user := person.user{FirstName: input.FirstName, LastName: input.LastName, Age: input.Age, Address: input.Address,JobInfo:input.JobInfo,username:input.username,password:input.password}
    databse.DB.Create(&user)
  
    c.JSON(http.StatusOK, gin.H{"data": book})
  }


func main() {
  router = gin.Default()//initialize router
  database.ConnectDatabase() // connecting database
  router.post("/createuser",createuser)// create new user
  router.GET("/GetDetails", Getdetails)  // get details
  log.Fatal(router.Run(":8080"))// listen report
  r.Run() // start router

  var user = User{
    FirstName: "Ramesh",
    LastName: :"Patil",
    Age: 30,
    Address:"Mumbai", 
    JobInfo: {"comapny": XYZ, Totexp:"15"},
    usename:"username",
    Password:"password"
  }
}


  






