package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var app = gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.Header("Content-Disposition", `attachment; filename="backup.db"`)
		_ = ctx.Stream(func(w io.Writer) bool {
			var err error
			var tester *Tester

			if tester, err = NewTester(); err != nil {
				_ = ctx.Error(err)
				return false
			}
			if _, err = io.Copy(w, tester); err != nil {
				_ = ctx.Error(err)
				return false
			}
			return false
		})
	})
	_ = app.Run()
}

// Tester ..
type Tester struct {
	data []byte
}

// NewTester ..
func NewTester() (tester *Tester, err error) {
	var file *os.File
	if file, err = os.Open("app.go"); err != nil {
		return
	}

	go func() {
		defer func() {
			if err = file.Close(); err != nil {
				log.Println(err)
			}
		}()
		for {
			var data = make([]byte, 512)

			_, err = file.Read(data)
		}
	}()
	return
}

func (t *Tester) Read(p []byte) (n int, err error) {

	n = copy(p, t.data[:])

	return
}

// Close ..
func (t *Tester) Close() (err error) {
	return
}
