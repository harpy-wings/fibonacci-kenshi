package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/harpy-wings/fibonacci-kenshi/internal/codec"
	"github.com/harpy-wings/fibonacci-kenshi/pkg/constants"
	"github.com/harpy-wings/fibonacci-kenshi/pkg/fibonacci"
	"github.com/labstack/echo/v4"
)

type defaultController struct {
	FB fibonacci.IFibonacci
}

func NewDefaultController(FB fibonacci.IFibonacci) Controller {
	return &defaultController{FB: FB}
}

func (c *defaultController) Post(ctx echo.Context) error {
	req := ctx.Request()
	reqCodec, err := codec.New(codec.OptionDescriber(req.Header.Get(constants.ContentTypeHeader)))
	if err != nil {
		return errors.Join(err, fmt.Errorf("invalid content-type header"))
	}
	bs, err := io.ReadAll(req.Body)
	if err != nil {
		return errors.Join(err, fmt.Errorf("read request body failed"))
	}
	FNum, err := reqCodec.Decode(bs)
	if err != nil {
		return errors.Join(err, fmt.Errorf("decode failed"))
	}
	next, err := c.FB.Next(FNum)
	if err != nil {
		return errors.Join(err, fmt.Errorf("fibonacci calculation failed"))
	}
	resCodec, err := codec.New(codec.OptionDescriber(req.Header.Get(constants.AcceptHeader)))
	if err != nil {
		return errors.Join(err, fmt.Errorf("invalid accept header"))
	}
	bs, err = resCodec.Encode(next)
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to encode response"))
	}
	err = ctx.String(http.StatusOK, string(bs))
	if err != nil {
		return errors.Join(err, fmt.Errorf("writing response failed"))
	}
	return nil
}

func (c *defaultController) Register(e *echo.Echo) error {
	e.POST("/", c.Post)
	return nil
}
