package service

import (
	"fmt"
	"log"
	"net/http"
	"redispractice/pojo"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var RC *redis.Client

// 連接Redis資料庫
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	log.Println(pong)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

//post user
// @Summary add new user
// @Tags User
// @version 1.0
// @produce application/json
// @Param param body pojo.User true "user information in json"
// @Success 200 string string "Successfully posted"
// @Router /v1/users/ [post]
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "Error:" + err.Error()})
		return
	}
	RC = newClient()
	defer RC.Close()
	hkey := fmt.Sprintf("user%d", user.Id)
	RC.HSet(hkey, "id", user.Id)
	RC.HSet(hkey, "name", user.Name)
	RC.HSet(hkey, "password", user.Password)
	RC.HSet(hkey, "email", user.Email)
	c.JSON(http.StatusOK, gin.H{"msg": "user added sucessful"})
}

// Find all
// @Summary find all user infomation
// @Tags User
// @version 1.0
// @produce text/plain
// @Success 200 string string pojo.User
// @Router /v1/users/ [get]
func FindAllUser(c *gin.Context) {
	RC = newClient()
	defer RC.Close()
	userList := []pojo.User{}
	users, _, _ := RC.Scan(0, "user*", 1000).Result()
	log.Println("user keys: ", users)
	for i := 0; i < len(users); i++ {
		user := pojo.User{}
		id, _ := RC.HGet(users[i], "id").Result()
		name, _ := RC.HGet(users[i], "name").Result()
		password, _ := RC.HGet(users[i], "password").Result()
		email, _ := RC.HGet(users[i], "email").Result()
		user.Id, _ = strconv.Atoi(id)
		user.Name = name
		user.Password = password
		user.Email = email
		userList = append(userList, user)
	}
	c.JSON(http.StatusOK, userList)
}

// put User (amend user)
// @Summary amend user
// @Tags User
// @version 1.0
// @produce application/json
// @Param id path integer true "id of user"
// @Param user body pojo.User true "amended user information in json"
// @Success 200 string string "Successfully amended"
// @Router /v1/users/{id} [put]
func PutUser(c *gin.Context) {
	beforeUser := pojo.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	RC = newClient()
	defer RC.Close()
	users, _, _ := RC.Scan(0, "user*", 1000).Result()
	for i := 0; i < len(users); i++ {
		id, _ := RC.HGet(users[i], "id").Result()
		if intid, _ := strconv.Atoi(id); intid == userId {
			RC.HSet(users[i], "name", beforeUser.Name)
			RC.HSet(users[i], "password", beforeUser.Password)
			RC.HSet(users[i], "email", beforeUser.Email)
			c.JSON(http.StatusOK, gin.H{"msg": "user information updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "Error"})
}

// delete User
// @Summary delete user
// @Tags User
// @version 1.0
// @produce text/plain
// @Param id path integer true "id of user"
// @Success 200 string string "Successfully deleted"
// @Router /v1/users/{id} [delete]
func Deleteuser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	RC = newClient()
	defer RC.Close()
	users, _, _ := RC.Scan(0, "user*", 1000).Result()
	for i := 0; i < len(users); i++ {
		id, _ := RC.HGet(users[i], "id").Result()
		if intid, _ := strconv.Atoi(id); intid == userId {
			RC.HDel(users[i], "id", "name", "password", "email")
			c.JSON(http.StatusOK, gin.H{"msg": "user deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "Error"})
}
