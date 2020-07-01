package tests
import (
    "testing"
    "fmt"
    "net/http"
    "net/http/httptest"

	"gintest/routes"

    "github.com/gin-gonic/gin"
)


func Server(w *httptest.ResponseRecorder, r *http.Request){
    router := routes.InitRouter()
	router.ServeHTTP(w,r)
}
func TestRoutePing(t *testing.T){
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request, _ = http.NewRequest("GET", "/auth?username=test&password=test123456", nil)
    Server(w,c.Request)

    fmt.Printf("auth 路由返回%s",w.Body.String())
    fmt.Println()
}
 