package main
import "net/http"
import "github.com/labstack/echo/v4"
import "math/rand"
import "time"
import "os"

type quote struct {
	Title string
	Author string
}

//var quotes []quote = make([]quote, 0)

var quotes []quote = []quote {
	{"Learn to Lead","Sai Vidya Campus"},
	{"cat","dog"},
	{"rabbit","parrot"},
	{"ur great","unknown"},
	{"u have potential","unknown"},
}
 
func main() {

	rand.Seed(time.Now().Unix())
	api := echo.New()

	api.GET("/quotes", getQuotes)
	api.GET("/quotes/random", getRandomQuote)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}
	api.Start(":"+ port)
	// localhost:32445

}

func getQuotes(c echo.Context) error {
	return c.JSON(http.StatusOK, quotes)
}

func getRandomQuote(c echo.Context) error {
	index := rand.Intn(len(quotes))

	return c.JSON(http.StatusOK, quotes[index])
