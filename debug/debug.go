package debug

import (
	"runtime"
	"net/http"
	_ "net/http/pprof"
)

func Debug() {
	runtime.SetBlockProfileRate(1)
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
}
