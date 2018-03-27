package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

var jwtSecret = []byte("secret")

type user struct {
	Name string `json:"name"`
}

type userClaims struct {
	user
	jwt.StandardClaims
}

func main() {
	app := iris.New()

	/*
		jwtHandler := jwtmw.New(jwtmw.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			},
			// When set, the middleware verifies that tokens are signed with the specific signing algorithm
			// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
			// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
			SigningMethod: jwt.SigningMethodHS256,
		})
	*/

	// Load all templates from the "./views" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./views", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})

	app.Get("/login", func(ctx iris.Context) {
		userParam := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		userParam.Username = "admin"
		userParam.Password = "admin"
		if userParam.Username != "admin" || userParam.Password != "admin" {
			//http.Error(w, "invalid login", http.StatusUnauthorized)
			ctx.StatusCode(http.StatusUnauthorized)
			ctx.WriteString("invalid login")
			return
		}

		//generate token
		expire := time.Now().Add(time.Hour * 1).Unix()
		// Create the Claims
		claims := userClaims{
			user: user{
				Name: userParam.Username,
			},
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expire,
				Issuer:    "login",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, _ := token.SignedString(jwtSecret)

		ctx.WriteString("siginedToken: " + signedToken)
		ctx.Values().Set("jwt", token)
		//WithHeader("Authorization", "Bearer "+tokenString).
		//Expect().Status(iris.StatusOK).Body().Contains("Iauthenticated").Contains("bar")
	})

	// Method:    GET
	// Resource:  http://localhost:8080/user/42
	//
	// Need to use a custom regexp instead?
	// Easy,
	// just mark the parameter's type to 'string'
	// which accepts anything and make use of
	// its `regexp` macro function, i.e:
	// app.Get("/user/{id:string regexp(^[0-9]+$)}")
	apiRoutes := app.Party("/api", apiMiddleware)
	//apiRoutes.Use(jwtHandler.Serve)
	{
		apiRoutes.Get("/user/{id:long}", func(ctx iris.Context) {
			userID, _ := ctx.Params().GetInt64("id")
			ctx.Writef("User ID: %d", userID)
		})
		apiRoutes.Get("/role/{id:long}", func(ctx iris.Context) {
			roleID, _ := ctx.Params().GetInt64("id")
			ctx.Writef("Role ID: %d", roleID)
		})

	}

	// Start the server using a network address.
	app.Run(iris.Addr(":8080"))
}

func apiMiddleware(ctx iris.Context) {
	fmt.Println("-->", ctx.Path())
	// [...]
	ctx.Next() // to move to the next handler, or don't that if you have any auth logic.
}
