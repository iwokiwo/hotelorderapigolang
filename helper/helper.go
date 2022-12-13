package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"os"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseV2 struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type WebResponsePaginationV2 struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Total   int64       `json:"total"`
	Data    interface{} `json:"data"`
}
type WebResponsePagination struct {
	Meta PaginationMeta `json:"meta"`
	Data interface{}    `json:"data"`
}

type PaginationMeta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Total   int64  `json:"total"`
}

func APIPagination(message string, code int, status string, total int64, data interface{}) WebResponsePaginationV2 {
	// meta := PaginationMeta{
	// 	Message: message,
	// 	Code:    code,
	// 	Status:  status,
	// 	Total:   total,
	// }

	jsonResponsePagination := WebResponsePaginationV2{
		//Meta: meta,
		Message: message,
		Code:    code,
		Status:  status,
		Total:   total,
		Data:    data,
	}

	return jsonResponsePagination
}

func LoggerFile(message string, types string, idUser int, errors interface{}) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	//--------------------------setting file log from .env---------------------------
	path := os.Getenv("LOGFILE") + "application.log"
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	//-------------------------switch type logger------------------------------------
	switch types {
	case "info":
		logger.WithField("userId", idUser).WithField("Errors", errors).Info(message)
	case "Warn":
		logger.WithField("userId", idUser).WithField("Errors", errors).Warnf(message)
	case "Error":
		logger.WithField("userId", idUser).WithField("Errors", errors).Error(message)
	}
}

func APIResponse(message string, code int, status string, data interface{}) ResponseV2 {
	// meta := Meta{
	// 	Message: message,
	// 	Code:    code,
	// 	Status:  status,
	// }

	jsonResponse := ResponseV2{
		//Meta: meta,
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

type ResponsePagination struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
