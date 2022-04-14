package main
import (
	"encoding/json"
	"net/http"
)
type Profile struct {
	Name    string
	Address []string
	Marks   []int
}
func main() {
	http.HandleFunc("/", studentInfo)
	http.ListenAndServe(":9000", nil)
}

func studentInfo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Harshal Lobhesh Bharat ", []string{"Wardha", "pune", "mumbai  "}, []int{89, 60, 85}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
