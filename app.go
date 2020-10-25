package main

import (
	"io"
	"log"
	"os"
	"sync"

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

type Tester struct {
	data   []byte
	locker *sync.Mutex
	eof    bool
}

func NewTester() (tester *Tester, err error) {
	tester = new(Tester)
	var file *os.File
	if file, err = os.Open("app.go"); err != nil {
		return
	}
	tester.locker = new(sync.Mutex)

	go func() {
		defer func() {
			if err = file.Close(); err != nil {
				log.Println(err)
			}
		}()
		for {
			var data = make([]byte, 512)
			var n int
			n, err = file.Read(data)
			tester.locker.Lock()
			tester.data = append(tester.data, data[:n]...)
			tester.locker.Unlock()
			if err == io.EOF {
				tester.eof = true
				break
			}
			if err != nil {
				log.Panic(err)
			}
		}
	}()
	return
}

func (t *Tester) Read(p []byte) (n int, err error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	n = copy(p, t.data[:])
	t.data = t.data[n:]
	if n == 0 && t.eof {
		err = io.EOF
	}
	return
}

func (t *Tester) Close() (err error) {
	return
}
