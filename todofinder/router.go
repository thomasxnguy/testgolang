package todofinder

import (
	"github.com/valyala/fasthttp"
	"strings"
	"encoding/json"
	. "todofinder/todofinder/error"
	. "todofinder/todofinder/http"
	"os"
)

//SearchResult HTTP Result
type SearchResultHttpResponse struct {
	Result []SearchResult `json:"result"`
}

func Router(ctx *fasthttp.RequestCtx) *Error {
	path := strings.Split(string(ctx.Path()), "/")
	switch len(path) {
	case 2:
		switch path[1] {
		case "search":
			return searchHandler(ctx)
		}
	}
	return &Error{NOT_FOUND, nil}
}

func searchHandler(ctx *fasthttp.RequestCtx) *Error {
	//SetHeaderGenerator(ctx)
	switch {
	case ctx.IsGet():
		{
			if ctx.QueryArgs().Has("package") && ctx.QueryArgs().Has("pattern") {
				searchResponse := SearchResultHttpResponse{}
				packageName := string(ctx.QueryArgs().Peek("package"))
				pattern := string(ctx.QueryArgs().Peek("package"))
				//TODO remove
				dir, err := os.Getwd()
				if err != nil {
					panic(err)
				}
				if p, error := ImportPkg(packageName, dir); error != nil {
					rch := make(chan *SearchResult, 10)
					go ExtractPattern(p, pattern, rch)
					for {
						searchResult := <-rch
						if searchResult == nil {
							return nil
						} else if searchResult.Error != nil {
							return searchResult.Error
						} else {
							searchResponse.Result = append(searchResponse.Result, *searchResult)
						}
					}
				} else {
					return error
				}

				JsonMarshal(ctx, searchResponse)
				return nil
			} else {
				//LogErrorWithFields(ctx, logrus.Fields{"func": "mediaHandler"}, "Wrong parameters", nil)
				return &Error{BAD_PARAMETER, nil}
			}
		}
	default:
		{
			return &Error{METHOD_NOT_ALLOWED, nil}
		}
	}

}

//Set the response body content with the interface v converted into JSON
func JsonMarshal(ctx *fasthttp.RequestCtx, v interface{}) {
	ctx.SetContentType(CONTENT_TYPE_JSON)
	err := json.NewEncoder(ctx).Encode(v)
	if err != nil {
		//LogErrorWithFields(ctx, logrus.Fields{"err": err, "func": "jsonMarshal"}, "Couldn't encode JSON", err)
	}
}
