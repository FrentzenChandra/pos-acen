package helper

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func EnabledCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mengambil header
		header := w.Header()
		// Mengubah Access Control Allow Origin menjadi boleh semua
		// yang dimana dampaknya adalah kita bisa berturkar fetch data dari domain ini ke domain manapun
		header.Set("Access-Control-Allow-Origin", "*")
		// Mengubah Access Control Allow Methods menjadi boleh semua method
		//dampaknya adalah kita bisa melakukan request method yang diminta
		header.Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT, PATCH")
		// mengubah access control allow headers menjadi boleh semua
		header.Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func URLRewriter(router *mux.Router, baseURLPath string) func(w http.ResponseWriter, r *http.Request) {
	// mengembalikan function http
	return func(w http.ResponseWriter, r *http.Request) {
		// mengambil r.URL.Path masukkan string yang dimana kita menambahkan
		// url yang biasa misalnya /users/signin menjadi
		// BaseUrl/users/signin hanya jika base url tidak ada pada
		//url dalam sebuah path pertama yang dituju sebelumnya
		r.URL.Path = func(url string) string {
			if strings.Index(url, baseURLPath) != 1 {
				url = url[len(baseURLPath):]
			}
			return url
		}(r.URL.Path)

		router.ServeHTTP(w, r)
	}
}

func LoggerMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "notifications") {
				next.ServeHTTP(w, r)
				return
			}

			start := time.Now()

			recorder := httptest.NewRecorder()
			next.ServeHTTP(recorder, r)

			for k, v := range recorder.Header() {
				w.Header()[k] = v
			}
			w.WriteHeader(recorder.Code)
			recorder.Body.WriteTo(w)

			responseTime := time.Since(start).Seconds()
			formattedResponseTime := fmt.Sprintf("%.9f", responseTime)
			formattedResponseTime = fmt.Sprintf("%sµs", formattedResponseTime)

			log.Printf("%s - [%s] - [%s] \"%s %s %s\" %d %s\n",
				r.RemoteAddr,
				time.Now().Format(time.RFC1123),
				formattedResponseTime,
				r.Method,
				r.URL.Path,
				r.Proto,
				recorder.Code,
				r.UserAgent(),
			)
		})
	}
}
