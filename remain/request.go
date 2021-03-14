package remain

import "net/http"

func CheckMethod(met string,  r *http.Request) bool{
	if r.Method == met{
		return true
	} else{
		return false
	}
}