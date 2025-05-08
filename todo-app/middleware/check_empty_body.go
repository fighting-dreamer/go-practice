package middleware

import "net/http"

func CheckEmptyBody(w http.ResponseWriter, r *http.Request)  {
	if r.ContentLength == 0 {
		return 
	}
}
