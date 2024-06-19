package forum

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// on handle l'integralité des dossier, templates et des differents composants pour notre serveur
func Setup() {
	http.Handle("/PageCSS/", http.StripPrefix("/PageCSS/", http.FileServer(http.Dir("./PageCSS/"))))
	http.Handle("/JS/", http.StripPrefix("/JS/", http.FileServer(http.Dir("./JS/"))))

	http.HandleFunc("/pages/login", Login)
	http.HandleFunc("/pages/about", Aboutus)
	http.HandleFunc("/pages/profile", Profile)
	http.HandleFunc("/pages/register", Register)
	http.HandleFunc("/pages/logout", Logout)
	http.HandleFunc("/index.html", Home)
	http.HandleFunc("/", Home)
	http.HandleFunc("/pages/mkpost", Addposts)
	http.HandleFunc("/pages/post", Postfunc)
	http.HandleFunc("/pages/admin", Admin)
	fmt.Printf("Started server successfully on http://localhost:1709/\n")
	http.ListenAndServe(":1709", nil)
}
