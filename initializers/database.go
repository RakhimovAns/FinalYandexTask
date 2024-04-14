package initializers

import (
	"encoding/hex"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

type expression struct {
	ID           int            `gorm:"primaryKey;autoIncrement"`
	Expression   string         `gorm:"type:text"`
	AddTime      int64          `gorm:"type:integer"`
	SubTime      int64          `gorm:"type:integer"`
	MultiplyTime int64          `gorm:"type:integer"`
	DivideTime   int64          `gorm:"type:integer"`
	Created      time.Time      `gorm:"type:timestamp"`
	Result       int64          `gorm:"type:integer;default:null"`
	IsCounted    bool           `gorm:"type:boolean;default:false"`
	DeletedAt    gorm.DeletedAt `gorm:"index;"`
}
type user struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:text"`
	Password string `gorm:"type:text"`
}

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")
	dsn = "host=postgres user=postgres password=postgres dbname=yandex port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error with connection to database")
	}
}

func CreateTable() {
	err := DB.AutoMigrate(&expression{})
	if err != nil {
		log.Fatal("failed to migrate expression")
	}
	err = DB.AutoMigrate(&user{})
	if err != nil {
		log.Fatal("failed to migrate user")
	}
}
func CreateModel(expression models.Expression) int64 {
	DB.Create(&models.Expression{Expression: expression.Expression, AddTime: expression.AddTime, SubTime: expression.SubTime, DivideTime: expression.DivideTime, MultiplyTime: expression.MultiplyTime})
	DB.Table("expressions").Where("expression=? AND add_time=? AND sub_time=? AND multiply_time=? AND divide_time=?", expression.Expression, expression.AddTime, expression.SubTime, expression.MultiplyTime, expression.DivideTime).Find(&expression)
	return expression.ID
}

func GetByID(ID int64) models.Expression {
	var expression models.Expression
	DB.Table("expressions").Where("id=?", ID).Find(&expression)
	return expression
}

func SetResult(id, result interface{}) {
	DB.Model(&expression{}).Where("id = ?", id).Updates(map[string]interface{}{"result": result, "is_counted": true})
}

func Register(customer models.User) error {
	var existingUser models.User
	log.Println(customer.Name)
	if err := DB.Where("name = ?", customer.Name).First(&existingUser).Error; err == nil {
		log.Println(existingUser.ID, existingUser.Name, existingUser.Password)
		return models.ErrUserExist
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(customer.Password))
	if err != nil {
		log.Println(err)
		return models.ErrInvalidPassword
	}
	DB.Create(&models.User{Password: hex.EncodeToString(hash), Name: customer.Name})
	return nil
}

func Login(customer models.User) (string, error) {
	var existingUser models.User
	err := DB.Where("name = ?", customer.Name).First(&existingUser).Error
	if err != nil {
		return "", models.ErrUserNotExist
	}
	hashed, err := hex.DecodeString(existingUser.Password)
	if err != nil {
		log.Println(err)
	}
	err = bcrypt.CompareHashAndPassword(hashed, []byte(customer.Password))
	if err != nil {
		return "", models.ErrInvalidPassword
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
		existingUser.ID,
	})
	TokenStr, err := token.SignedString([]byte("My Key"))
	return TokenStr, nil
}
