package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the API Home Page!")
}

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    // name := vars["name"]
    age, exist := vars["age"]
	if exist {
		fmt.Fprintln(w, "Age is: " ,exist, age)
	} else {
		fmt.Fprintln(w, "Age is: " ,exist)
		fmt.Fprintln(w, "Age is not exist")
	}
    // fmt.Fprintf(w, "You requested information for: %s\n", name)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w ,"Hello world!")
}


func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if userIsLoggedIn(r) {
            next.ServeHTTP(w, r)
        } else {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
        }
    })
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is a secure page. You can only access it if you are logged in.")
}

func userIsLoggedIn(r *http.Request) bool {
    return false
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login page")
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/info/{name}", GetInfoHandler)
	r.HandleFunc("/hello", HelloWorld)
	r.HandleFunc("/login", Login)

	r.Handle("/secure", AuthMiddleware(http.HandlerFunc(SecureHandler)))

    http.Handle("/", r)

    fmt.Println("Server is running on :8080...")
    http.ListenAndServe(":8080", nil)
}
