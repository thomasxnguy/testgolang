package jwt

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"strings"
)

var conf *Configuration

func RequestHandler(ctx *fasthttp.RequestCtx) *error {

	path := strings.SplitN(string(ctx.Path()), "/", 3)

	if len(path) == 2 && path[0] == ".well-known" {
		return keysHandler(ctx, &path[1], conf)
	} else {
		return nil //Bad request
	}
}

//KeysHandler is handler, which return App public key to clients
func keysHandler(ctx *fasthttp.RequestCtx, filename *string, conf *Configuration) *error {
	if *filename == "jwks.json" {
		jsonMarshal(ctx, conf.PublicKeys)
	} else {
		return nil //Return Error
	}
	return nil
}


//Set the response body content with the interface v converted into JSON
func jsonMarshal(ctx *fasthttp.RequestCtx, v interface{}) {
	ctx.SetContentType("application/json; charset=utf-8")

	err := json.NewEncoder(ctx).Encode(v)
	if err != nil {
		//Log Error
	}
}
