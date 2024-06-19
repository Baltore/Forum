package forum

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Ratelimit(w, r)
	Ifregistered(w, r)
	Reports(w, r)
	Getposts(w, r)
	Addlike(w, r)
	Notifications(w, r)
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "../PageHtml/Home.html"))).Execute(w, templ)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	Ratelimit(w, r)
	if Ifregistered(w, r) {
		Getprofileinfo(w, r)
		template.Must(template.ParseFiles(filepath.Join(templatesDir, "../PageHtml/Profile.html"))).Execute(w, templ)
	}
}

func Aboutus(w http.ResponseWriter, r *http.Request) {
	Ratelimit(w, r)
	Ifregistered(w, r)
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "../PageHtml/About.html"))).Execute(w, templ)
}
