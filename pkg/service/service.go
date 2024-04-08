package service

import (
	"errors"
	"github.com/RakhimovAns/FinalYandexTask/govaluate"
	"github.com/RakhimovAns/FinalYandexTask/models"
	"time"
)

func IsValidate(expression models.Expression) error {
	size := len(expression.Expression)
	if expression.SubTime < 0 || expression.AddTime < 0 || expression.MultiplyTime < 0 || expression.DivideTime < 0 {
		return errors.New("invalid format of expression")
	}
	if size == 0 || !(expression.Expression[0] >= '0' && expression.Expression[0] <= '9') {
		return errors.New("invalid format of expression")
	}
	if !(expression.Expression[size-1] >= '0' && expression.Expression[size-1] <= '9') {
		return errors.New("invalid format of expression")
	}
	for i := 1; i < size; i++ {
		if !(expression.Expression[i] >= '0' && expression.Expression[i] <= '9') && !(expression.Expression[i-1] >= '0' && expression.Expression[i-1] <= '9') {
			return errors.New("invalid format of expression")
		}
	}
	return nil
}

func CountExpression(expression models.Expression) (interface{}, error) {
	expr, err := govaluate.NewEvaluableExpression(expression.Expression)
	if err != nil {
		return -1, errors.New("error with parsing")
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return -1, errors.New("error evaluating expression")
	}
	Time := FindTime(expression)
	time.Sleep(time.Duration(Time) * time.Second)
	return result, nil
}

func FindTime(expression models.Expression) int64 {
	result := int64(0)
	for _, char := range expression.Expression {
		switch char {
		case '+':
			result += expression.AddTime
		case '-':
			result += expression.SubTime
		case '/':
			result += expression.DivideTime
		case '*':
			result += expression.MultiplyTime
		}
	}
	return result
}
