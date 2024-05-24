package controllers

import (
	"github.com/gin-gonic/gin"
)

type Employee interface {
	GetEmployeeDetails(*gin.Context)
}
