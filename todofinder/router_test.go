package todofinder_test

import (
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/error"
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/app"
	"github.com/valyala/fasthttp"
)

var _ = Describe("Router", func() {
	var (
		ctx *fasthttp.RequestCtx
	)

	BeforeEach(func() {
		ctx = &fasthttp.RequestCtx{}
		s := "conf/todofinder_test.yaml"
		config, _ := LoadConfiguration(&s)
		log, _ := LoadAppLogger(config)
		ctx.Init(&fasthttp.Request{}, nil, log)
	})

	Describe("Routing request", func() {
		Context("when calling bad request", func() {
			It("should return bad request error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/notexist"))
				error := Router(ctx)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(NOT_FOUND))
			})
		})
	})

	Describe("Routing request", func() {
		Context("when calling not allowed method", func() {
			It("should return method not allowed error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/search"))
				ctx.Request.Header.SetMethodBytes([]byte("PUT"))
				error := Router(ctx)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(METHOD_NOT_ALLOWED))
			})
		})
	})

	Describe("Routing request", func() {
		Context("when using bad parameters", func() {
			It("should return bad parameters error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/search"))
				ctx.Request.Header.SetMethodBytes([]byte("GET"))
				error := Router(ctx)
				Expect(error).NotTo(BeNil())
				Expect(error.ErrorCode).To(Equal(BAD_PARAMETER))
			})
		})
	})

	Describe("Routing request", func() {
		Context("using happy path", func() {
			It("should not return error", func() {
				ctx.Request.SetRequestURIBytes([]byte("/search"))
				ctx.Request.Header.SetMethodBytes([]byte("GET"))
				ctx.QueryArgs().Add("package", "fmt")
				ctx.QueryArgs().Add("pattern", "TODO")
				error := Router(ctx)
				Expect(error).To(BeNil())
			})
		})
	})
})
