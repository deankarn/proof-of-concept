package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/go-playground/rum"
)

// CustomContext ...
type CustomContext struct {
	rum.Context
}

func (ctx *CustomContext) Test() string {
	return "\n\nI Am HERE!"
}

var (
	t *template.Template
)

func main() {

	var err error

	t, err = template.New("foo").Parse(`{{define "T"}}<html><body><p>Hello, {{.}}!</p></body></html>{{end}}`)
	if err != nil {
		log.Fatal("Error Creating Template")
	}

	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))

	r := rum.New()

	// nc := func() rum.Context {
	// 	return &CustomContext{rum.NewContext(nil, rum.NewResponse(nil), r)}
	// }

	// r.SetContextNew(nc)

	r.Use(Logger())
	// r.Use(func(h rum.HandlerFunc) rum.HandlerFunc {
	// 	return func(ctx rum.Context) {
	// 		h(&CustomContext{ctx})
	// 	}
	// })
	r.Get("/", root)
	r.Get("/assets/*", func(c *rum.Context) {
		fs.ServeHTTP(c.Response, c.Request)
	})

	r.Get("/teachers/list", route1)
	r.Get("/teachers/:id/profile", route2)

	// g1 := r.Group("/route1/")
	// // g1.Use(Logger2())
	// g1.Get("", route1)
	// g1.Get("a/", route1a)

	// g2 := r.Group("/route2/")
	// // g2.Use(Logger2())
	// g2.Get("", route2)
	// g2.Get(":name/", route2a)

	http.ListenAndServe(":3006", r)
}

func root(ctx *rum.Context) {

	// err := t.ExecuteTemplate(ctx.Response(), "T", "Dean Karn")
	// if err != nil {
	// 	http.Error(ctx.Response(), "Error Executing Template", http.StatusInternalServerError)
	// }

	ctx.Response.Write([]byte(fmt.Sprint(reflect.TypeOf(ctx))))
	// ctx.Response().Write([]byte(fmt.Sprint(reflect.TypeOf(ctx), ctx.(*CustomContext).Test())))
}

func route1(ctx *rum.Context) {
	ctx.Response.Write([]byte("Route 1"))
}

func route2(ctx *rum.Context) {
	ctx.Response.Write([]byte("Route 2"))
}

func route1a(ctx *rum.Context) {
	ctx.Response.Write([]byte("Route 1a"))
}

func route2a(ctx *rum.Context) {
	ctx.Response.Write([]byte("Route 2a " + ctx.Param("name")))
}

func Logger() rum.MiddlewareFunc {

	return func(h rum.HandlerFunc) rum.HandlerFunc {
		return func(c *rum.Context) {
			req := c.Request
			// res := c.Response()
			// logger := c.Echo().Logger()

			// remoteAddr := req.RemoteAddr
			// if ip := req.Header.Get(trafficcop.XRealIP); ip != "" {
			// 	remoteAddr = ip
			// } else if ip = req.Header.Get(trafficcop.XForwardedFor); ip != "" {
			// 	remoteAddr = ip
			// } else {
			// 	remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
			// }

			start := time.Now()
			h(c)
			// if err := h(c); err != nil {
			// 	c.Error(err)
			// }
			stop := time.Now()
			method := req.Method
			path := req.URL.Path

			if path == "" {
				path = "/"
			}

			// size := res.Size()

			// n := res.Status()
			// code := color.Green(n)
			// switch {
			// case n >= 500:
			// 	code = color.Red(n)
			// case n >= 400:
			// 	code = color.Yellow(n)
			// case n >= 300:
			// 	code = color.Cyan(n)
			// }

			log.Printf("%s %s %s\n", method, path, stop.Sub(start))
			// logger.Info("%s %s %s %s %s %d", remoteAddr, method, path, code, stop.Sub(start), size)
		}
	}
}

func Logger2() rum.MiddlewareFunc {

	return func(h rum.HandlerFunc) rum.HandlerFunc {
		return func(c *rum.Context) {
			req := c.Request
			// res := c.Response()
			// logger := c.Echo().Logger()

			// remoteAddr := req.RemoteAddr
			// if ip := req.Header.Get(trafficcop.XRealIP); ip != "" {
			// 	remoteAddr = ip
			// } else if ip = req.Header.Get(trafficcop.XForwardedFor); ip != "" {
			// 	remoteAddr = ip
			// } else {
			// 	remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
			// }

			start := time.Now()
			h(c)
			// if err := h(c); err != nil {
			// 	c.Error(err)
			// }
			stop := time.Now()
			method := req.Method
			path := req.URL.Path

			if path == "" {
				path = "/"
			}

			// size := res.Size()

			// n := res.Status()
			// code := color.Green(n)
			// switch {
			// case n >= 500:
			// 	code = color.Red(n)
			// case n >= 400:
			// 	code = color.Yellow(n)
			// case n >= 300:
			// 	code = color.Cyan(n)
			// }

			log.Printf("L2: %s %s %s\n", method, path, stop.Sub(start))
			// logger.Info("%s %s %s %s %s %d", remoteAddr, method, path, code, stop.Sub(start), size)
			// return nil
		}
	}
}
