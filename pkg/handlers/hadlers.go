package handlers

import (
	"context"
	desc "github.com/RakhimovAns/FinalYandexTask/api/gen/api/service"
	"github.com/RakhimovAns/FinalYandexTask/initializers"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"github.com/RakhimovAns/FinalYandexTask/pkg/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	grpcServerURL = "localhost:50051"
)

func PostExpression(c *gin.Context) {
	var expression models.Expression
	if err := c.ShouldBindJSON(&expression); err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON data: " + err.Error()})
		return
	}
	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed connect: " + err.Error()})
		return
	}
	defer conn.Close()
	client := desc.NewCalculusClient(conn)
	resp, err := client.Calculate(context.Background(), &desc.Expression{Expression: expression.Expression, AddTime: expression.AddTime, DivideTime: expression.DivideTime, SubTime: expression.SubTime, MultiplyTime: expression.MultiplyTime})
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info from gRPC server " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": resp.Id})
}

func GetStatus(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed connect: " + err.Error()})
		return
	}
	defer conn.Close()
	client := desc.NewCalculusClient(conn)
	_, err = client.Status(context.Background(), &desc.ID{Id: int64(ID)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info from gRPC server " + err.Error()})
		return
	}
	expression := initializers.GetByID(int64(ID))
	if expression.IsCounted {
		c.JSON(http.StatusOK, gin.H{"result": expression.Result})
	} else {
		// Отправляем сообщение о том, что подсчет начат
		c.JSON(http.StatusOK, gin.H{"result": "counting"})
		// Запускаем горутину для асинхронного подсчета выражения
		go func() {
			_, err := service.CountExpression(expression)
			if err != nil {
				log.Printf("Error counting expression: %v", err)
				return
			}
		}()
	}
}

func RegisterHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data: " + err.Error()})
		return
	}
	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed connect: " + err.Error()})
		return
	}
	defer conn.Close()
	client := desc.NewCalculusClient(conn)
	_, err = client.Register(context.Background(), &desc.User{Name: user.Name, Password: user.Password})
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: already exists!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "registered successfully, try to log in"})
}
func Logout(c *gin.Context) {
	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed connect: " + err.Error()})
		return
	}
	defer conn.Close()
	client := desc.NewCalculusClient(conn)
	_, err = client.Logout(context.Background(), &desc.Empty{})
	// Получаем список куки из запроса
	cookies := c.Request.Cookies()

	for _, cookie := range cookies {
		newCookie := &http.Cookie{
			Name:   cookie.Name,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
		c.SetCookie(newCookie.Name, newCookie.Value, newCookie.MaxAge, newCookie.Path, newCookie.Domain, newCookie.Secure, newCookie.HttpOnly)
	}
	// Выполняем перенаправление на страницу входа
	c.Redirect(http.StatusSeeOther, "/log")
}
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data: " + err.Error()})
		return
	}
	conn, err := grpc.Dial(grpcServerURL, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed connect: " + err.Error()})
		return
	}
	defer conn.Close()
	client := desc.NewCalculusClient(conn)
	token, err := client.Login(context.Background(), &desc.User{Name: user.Name, Password: user.Password})
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	livingTime := 60 * time.Minute
	expiration := time.Now().Add(livingTime)
	cookie := http.Cookie{Name: "token", Value: url.QueryEscape(token.Token), Expires: expiration}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, gin.H{"token": token.Token})
}
