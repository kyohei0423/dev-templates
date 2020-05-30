package controller

import (
	"net/http"
)

type SampleController struct{}

func NewSampleController() *SampleController {
	return &SampleController{}
}

func (controller *SampleController) Hello(c Context) error {
	return c.JSON(http.StatusCreated, "Hello world")
}
