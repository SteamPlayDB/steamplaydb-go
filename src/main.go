package main

import (
	"./Controllers"
	"./Utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"path"
)

var (
	publicDir     string
	indexFile     string
	autoCertCache string
	portNumber    string
)

func setupRoutes(e *echo.Echo) {
	e.Static("/", publicDir)
	e.GET("/api/games", Controllers.Games)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if err := c.File(path.Join(publicDir, indexFile)); err != nil {
		c.Logger().Error(err)
	}

	c.Logger().Error(err)
}

func main() {

	err := godotenv.Load()

	// Assign the "global" vars from the env.
	publicDir = Utils.Getenv("PUBLIC_DIR", "public")
	indexFile = Utils.Getenv("INDEX_FILE", "index.html")
	autoCertCache = Utils.Getenv("AUTOCERT_CACHE", "")
	portNumber = Utils.Getenv("PORT", "80")

	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	e := echo.New()

	if len(autoCertCache) > 0 {
		e.AutoTLSManager.Cache = autocert.DirCache(autoCertCache)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	e.HTTPErrorHandler = customHTTPErrorHandler

	setupRoutes(e)

	// Start the server on the selected port number.
	if len(autoCertCache) > 0 {
		e.Logger.Fatal(e.StartAutoTLS(":" + portNumber))
	} else {
		e.Logger.Fatal(e.Start(":" + portNumber))
	}
}
