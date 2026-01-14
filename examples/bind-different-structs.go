package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type fromA struct {
	Foo string `json:"foo" xl:"foo" binding:"required"`
}

type fromB struct {
	Foo string `json:"bar" xl:"bar" binding:"required"`
}

func BindDifferentStructs(c *gin.Context) {
	objA := fromA{}
	objB := fromB{}

	if errA := c.ShouldBind(&objA); errA != nil {
		c.String(http.StatusOK, `the body should be fromA`)
	} else if errB := c.ShouldBind(&objB); errB != nil {
		c.String(http.StatusOK, `the body should be fromB`)
	}
}

func BindingBodyWith(c *gin.Context) {
	objA := fromA{}
	objB := fromB{}

	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA != nil {
		c.String(http.StatusOK, `the body should be fromA`)
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB != nil {
		c.String(http.StatusOK, `the body should be fromB`)
	}
}
