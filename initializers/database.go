package initializers

import (
	"github.com/RakhimovAns/FinalYandexTask/models"
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
	dsn = "host=localhost user=postgres password=postgres dbname=yandex port=5432 sslmode=disable"
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
