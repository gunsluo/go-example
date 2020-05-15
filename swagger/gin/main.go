package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

const (
	addr = ":9000"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	engine.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	apiRoutes := engine.Group("/api")
	apiRoutes.Use(handlerAuth())

	apiRoutes.GET("/regions", allRegions)
	apiRoutes.GET("/user/:sub/regions", userRegions)
	apiRoutes.GET("/user/:sub/organizations", userOrganizations)

	logger := logrus.New()
	logger.WithField("addr", addr).Println("Starting HTTP RESTFul server")

	err := engine.Run(addr)
	if err != nil {
		logger.WithError(err).Fatalf("Server exceptions: %s", err)
	}
}

type regionInfo struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	RedirectURL string `json:"redirectUrl"`
}

func allRegions(ctx *gin.Context) {
	regions := []regionInfo{}
	regions = append(regions, regionInfo{Name: "dev", Label: "test", RedirectURL: "http://abc.com"})

	ctx.JSON(http.StatusOK, gin.H{
		"regions": regions})
}

func userRegions(ctx *gin.Context) {
	subject := ctx.Param("sub")
	if subject == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": codes.InvalidArgument,
			"err":  "missing subject"})
		return
	}

	regions := []regionInfo{}
	regions = append(regions, regionInfo{Name: "dev", Label: "test", RedirectURL: "http://abc.com"})
	ctx.JSON(http.StatusOK, gin.H{
		"regions": regions})
}

type organizationInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type regionWithOrganizationInfo struct {
	regionInfo
	Organizations []organizationInfo `json:"organization"`
}

func userOrganizations(ctx *gin.Context) {
	subject := ctx.Param("sub")
	if subject == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": codes.InvalidArgument,
			"err":  "missing subject"})
		return
	}

	regions := []*regionWithOrganizationInfo{}
	ro := &regionWithOrganizationInfo{
		regionInfo: regionInfo{Name: "dev", Label: "test", RedirectURL: "http://abc.com"}}
	ro.Organizations = append(ro.Organizations,
		organizationInfo{Id: 1, Name: "test"})
	regions = append(regions, ro)

	ctx.JSON(http.StatusOK, gin.H{
		"regions": regions})
}

func handlerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		authHeaders, ok := r.Header["Authorization"]
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Authorization header is empty"))
			return
		}
		if len(authHeaders) != 1 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("More than one Authorization headers sent"))
			return
		}

		parts := strings.SplitN(authHeaders[0], " ", 2)
		if len(parts) != 2 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Bad Authorization header"))
			return
		}
		if !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Only Bearer tokens accepted"))
			return
		}

		// TODO: check parts[1]
		fmt.Println("token:", parts[1])

		c.Next()
	}
}
