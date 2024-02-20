package utils 
import(
	"encodeing/json"
	"io"
	"net/http"
)


func ParseBody(r *http.Request , X interface{}){
	if body , err :=io.ReadAll(r.Body);
}