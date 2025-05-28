package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"1337b04rd/internal/adapters/driven/redis"
	"1337b04rd/internal/usecase"
	"1337b04rd/pkg/config"
)

func main() {
	conf := config.InitConfig()

	redisConf := conf.GetRedisConfig()
	red := redis.InitRickRedis(redisConf)

	use := usecase.InitRickAndMortyCase(red)

	mainCtx := context.Background()
	ctx, cancel := context.WithTimeout(mainCtx, 10*time.Second)
	defer cancel()
	ans, err := use.GetCharacter(ctx)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans.ID)
		fmt.Println(ans.Name)
		fmt.Println(ans.Image)
	}
	fmt.Println("finish")
	// sessionConf := conf.GetSessionConfig()
	// redSession := redis.InitSessionRedis(redisConf, sessionConf.GetDuration())
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize a new cookie containing the string "Hello world!" and some
	// non-default attributes.
	cookie := http.Cookie{
		Name:     "i_love_session",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	w.Write([]byte("cookie set!"))
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:

			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	// Echo out the cookie value in the response body.
	w.Write([]byte(cookie.Value))
}
