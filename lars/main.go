package main

import (
	"bufio"
	"compress/gzip"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/lars"
)

// // CustomContext ...
// type CustomContext struct {
// 	rum.Context
// }

// func (ctx *CustomContext) Test() string {
// 	return "\n\nI Am HERE!"
// }

var (
	t *template.Template
)

type Globals struct {
	Name string
	DB   int
}

func (g *Globals) Reset() {
	g.DB++
}

func main() {

	var err error

	t, err = template.New("foo").Parse(`{{define "T"}}<html><body><p>Hello, {{.}}!</p></body></html>{{end}}`)
	if err != nil {
		log.Fatal("Error Creating Template")
	}

	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))

	l := lars.New()
	l.RegisterGlobalsFunc(func() lars.IGlobals { return new(Globals) })
	// nc := func() l.Context {
	// 	return &CustomContext{rum.NewContext(nil, rum.NewResponse(nil), r)}
	// }

	// r.SetContextNew(nc)

	l.Use(Logger())
	l.Use(Gzip())
	// l.Use(func(h rum.HandlerFunc) rum.HandlerFunc {
	// 	return func(ctx rum.Context) {
	// 		h(&CustomContext{ctx})
	// 	}
	// })
	l.Get("/", root)
	l.Get("/teachers/list", route1)
	l.Get("/teachers/:id/profile", route2)
	// l.Get("/assets/*filepath", func(c *lars.Context) {
	// 	fs.ServeHTTP(c.Response, c.Request)
	// })

	l.Get("/assets/*filepath", fs)

	// g1 := l.Group("/teachers/")
	// // g1.Use(Logger2())
	// g1.Get("list/", route1)
	// // g1.Get("a/", route1a)
	// g1.Get(":id/profile", route2)

	// g2 := l.Group("/route2/")
	// // g2.Use(Logger2())
	// g2.Get("", route2)
	// g2.Get(":name/", route2a)

	http.ListenAndServe(":3006", l)
}

func root(ctx *lars.Context) {

	err := t.ExecuteTemplate(ctx.Response, "T", "Dean Karn")
	if err != nil {
		http.Error(ctx.Response, "Error Executing Template", http.StatusInternalServerError)
	}
	//
	// ctx.Response.Write([]byte(strconv.Itoa(ctx.Globals.(*Globals).DB)))
	// ctx.Response.Write([]byte(fmt.Sprint(reflect.TypeOf(ctx.Globals))))
	// ctx.Response.Write([]byte(fmt.Sprint(reflect.TypeOf(ctx))))
	// ctx.Response.Write([]byte(fmt.Sprint(reflect.TypeOf(ctx), ctx.(*CustomContext).Test())))
	// ctx.Response.Write([]byte("Home"))
}

func route1(ctx *lars.Context) {
	ctx.Response.Write([]byte("Teachers List"))
}

func route2(ctx *lars.Context) {
	ctx.Response.Write([]byte("teacher profile"))
}

func route1a(ctx *lars.Context) {
	ctx.Response.Write([]byte("Route 1a"))
}

func route2a(ctx *lars.Context) {
	ctx.Response.Write([]byte("Route 2a " + ctx.Param("name")))
}

func Logger() lars.MiddlewareFunc {

	return func(h lars.HandlerFunc) lars.HandlerFunc {
		return func(c *lars.Context) {
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
			// c.Next()
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

type (
	gzipWriter struct {
		io.Writer
		http.ResponseWriter
	}
)

func (w gzipWriter) Write(b []byte) (int, error) {
	if w.Header().Get(lars.ContentType) == "" {
		w.Header().Set(lars.ContentType, http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

func (w gzipWriter) Flush() error {
	return w.Writer.(*gzip.Writer).Flush()
}

func (w gzipWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (w *gzipWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

var writerPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(ioutil.Discard)
	},
}

const (
	scheme = "gzip"
)

// Gzip returns a middleware which compresses HTTP response using gzip compression
// scheme.
func Gzip() lars.MiddlewareFunc {
	// scheme := "gzip"

	return func(h lars.HandlerFunc) lars.HandlerFunc {
		return func(c *lars.Context) {
			c.Response.Header().Add(lars.Vary, lars.AcceptEncoding)

			if strings.Contains(c.Request.Header.Get(lars.AcceptEncoding), scheme) {

				w := writerPool.Get().(*gzip.Writer)
				w.Reset(c.Response)

				defer func() {
					w.Close()
					writerPool.Put(w)
				}()

				gw := gzipWriter{Writer: w, ResponseWriter: c.Response}
				c.Response.Header().Set(lars.ContentEncoding, scheme)
				c.Response = &lars.Response{ResponseWriter: gw}
				// c.Response().SetWriter(gw)
			}

			h(c)
			// if err := h(c); err != nil {
			// 	c.Error(err)
			// }
			// return nil
		}
	}
}
