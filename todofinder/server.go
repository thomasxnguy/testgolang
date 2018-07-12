package todofinder

import (
	"github.com/valyala/fasthttp"
	"net"
	. "todofinder/todofinder/error"
	. "todofinder/todofinder/http"
)

//Name of the service
const (
	ServerName  = "todofinder"
)

type Server struct {
	config *Configuration
}

func (s *Server) Init(conf *Configuration) error {
	s.config = conf
	return nil
}

func (s *Server) getListener() (net.Listener, error) {
	ln, err := net.Listen(s.config.Network, s.config.ListenOn)
	//log port
	if err != nil {
		//log
		return nil, err
	}
	return ln, nil
}

func (s *Server) Run() error {
	fhttps := &fasthttp.Server{
		Handler: s.handler,
		Name:    ServerName,
	}
	ln, err := s.getListener()
	if err != nil {
		return err
	}

	if s.config.EnableTls {
		if err := fhttps.ServeTLS(ln, s.config.CertFile, s.config.KeyFile); err != nil {
			//Log.Log.WithField("func", "Run").Fatal("Error when serving incoming connections")
			return err
		}
	} else {
		if err := fhttps.Serve(ln); err != nil {
			//Log.Log.WithField("func", "Run").Fatal("Error when serving incoming connections")
			return err
		}
	}
	//log info
	return nil
}

//Run requestHandler with error handling
func (s *Server) handler(ctx *fasthttp.RequestCtx) {
	//log := &LoggerContext{}
	//log.init("todofinder", ctx)
	err := Router(ctx)
	if err != nil {
		errorHandler(ctx, err)
	}
}

//Handle the error, return an http response error according to the error code
func errorHandler(ctx *fasthttp.RequestCtx, e *Error) {
	if e.Error != nil {
		//LogDebugWithFields(ctx, logrus.Fields{"func": "ErrorHandler"}, "Return Error", e.Error)
	}
	if e.ErrorCode == "" {
		//LogDebugWithFields(ctx, logrus.Fields{"func": "ErrorHandler"}, "Return Error", e.Error)
		e.ErrorCode = SERVER_ERROR
	}
	errorMessage := ErrorMessages[e.ErrorCode]
	errorHttpResponse := &ErrorHttpResponse{errorMessage.ErrorMessage, e.ErrorCode, errorMessage.HttpStatus}
	ctx.Error(errorHttpResponse.ToJson(), errorHttpResponse.HttpStatus)
	//Need to set the content type and headers which have been reset by the previous method
	ctx.SetContentType(CONTENT_TYPE_JSON)
	ctx.Response.Header.Add(ACCESS_CONTROL_ALLOW_ORIGIN, "*")
	ctx.Response.Header.Add(ACCESS_CONTROL_ALLOW_METHOD, "*")
	ctx.Response.Header.Add(ACCESS_CONTROL_ALLOW_HEADERS, "*")
}
