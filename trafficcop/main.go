package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/traffic-cop"
)

// CustomContext ...
type CustomContext struct {
	trafficcop.Context
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

	tcop := trafficcop.New()
	tcop.Use(Logger())
	tcop.Use(func(h trafficcop.HandlerFunc) trafficcop.HandlerFunc {
		return func(ctx trafficcop.Context) {
			h(&CustomContext{ctx})
		}
	})
	tcop.Get("/", root)
	tcop.Get("/assets/*", func(c trafficcop.Context) {
		fs.ServeHTTP(c.Response().Writer(), c.Request())
	})

	g1 := tcop.Group("/route1/")
	// g1.Use(Logger2())
	g1.Get("", route1)
	g1.Get("a/", route1a)

	g2 := tcop.Group("/route2/")
	// g2.Use(Logger2())
	g2.Get("", route2)
	g2.Get(":name/", route2a)

	http.ListenAndServe(":3006", tcop)
}

func root(ctx trafficcop.Context) {

	err := t.ExecuteTemplate(ctx.Response(), "T", "Dean Karn")
	if err != nil {
		http.Error(ctx.Response(), "Error Executing Template", http.StatusInternalServerError)
	}

	// ctx.Response().Write([]byte(fmt.Sprint(reflect.TypeOf(ctx))))
	// ctx.Response().Write([]byte(fmt.Sprint(reflect.TypeOf(ctx), ctx.(*CustomContext).Test())))
}

func route1(ctx trafficcop.Context) {
	ctx.Response().Write([]byte("Route 1"))
}

func route2(ctx trafficcop.Context) {
	ctx.Response().Write([]byte("Route 2"))
}

func route1a(ctx trafficcop.Context) {
	ctx.Response().Write([]byte("Route 1a"))
}

func route2a(ctx trafficcop.Context) {
	ctx.Response().Write([]byte("Route 2a " + ctx.Param("name")))
}

func Logger() trafficcop.MiddlewareFunc {

	return func(h trafficcop.HandlerFunc) trafficcop.HandlerFunc {
		return func(c trafficcop.Context) {
			req := c.Request()
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

func Logger2() trafficcop.MiddlewareFunc {

	return func(h trafficcop.HandlerFunc) trafficcop.HandlerFunc {
		return func(c trafficcop.Context) {
			req := c.Request()
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
