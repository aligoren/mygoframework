package mygoframework

import "encoding/json"

type Context struct {
	request  Request
	response Response
}

func (ctx *Context) SetStatus(statusCode int) error {
	ctx.response.WriteHeader(statusCode)

	return nil
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

	ctx.response.Write(jsonData)

	return nil
}
