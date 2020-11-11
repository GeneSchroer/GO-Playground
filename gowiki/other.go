import(
	"fmt"
	"io/ioutil"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)

}
