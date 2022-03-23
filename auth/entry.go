package auth

import (
	"net/http"
)

func AuthCheck(w http.ResponseWriter, r *http.Request){
	_,err := r.Cookie("p")
	if err != nil{
		http.Redirect(w,r,"/create",301)
	}else{
		http.Redirect(w,r,"/main",301)
	}
}
