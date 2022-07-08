package mygoframework

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
)

type Context struct {
	request    Request
	response   Response
	statusCode int
}

func (ctx *Context) SetStatus(statusCode int) *Context {
	if statusCode == 0 {
		statusCode = 200
	}

	ctx.statusCode = statusCode

	return ctx
}

func (ctx *Context) SendString(text string) error {
	ctx.response.Write([]byte(text))

	return nil
}

func (ctx *Context) SendJson(data interface{}) error {

	jsonData, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	ctx.response.Header().Set("Content-Type", "application/json")

	ctx.response.WriteHeader(ctx.getStatusCode())

	ctx.response.Write(jsonData)

	return nil
}

func (ctx *Context) SendXml(data interface{}) error {

	xmlData, err := xml.MarshalIndent(data, "", " ")

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	ctx.response.Header().Set("Content-Type", "application/xml")

	ctx.response.WriteHeader(ctx.getStatusCode())

	ctx.response.Write(xmlData)

	return nil
}

func (ctx *Context) Param(key string) (string, error) {

	value, ok := ctx.request.Params().Get(key)

	if !ok {
		return "", errors.New("params couldn't found")
	}

	return value, nil
}
