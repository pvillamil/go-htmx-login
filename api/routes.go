package api

import (
	"log"
	"net/http"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received at:", r.URL.Path)

	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "public/index.html")
	case "/public/assets/main.css":
		http.ServeFile(w, r, "public/assets/main.css")
	case "/initial":
		w.Write([]byte(`
			<div class="card-container">
				<h2>Login</h2>
				<form hx-post="/login" hx-target="this" hx-swap="outerHTML">
					<input type="text" name="username" placeholder="Username">
					<input type="password" name="password" placeholder="Password">
					<button type="submit">Login</button>
				</form>
				<a href="/createuser" hx-get="/createuser" hx-target="closest div" hx-swap="outerHTML">Create User</a>
			</div>
		`))
	case "/createuser":
		w.Write([]byte(`
			<div class="card-container">
				<h2>Create User</h2>
				<form hx-post="/login" hx-target="this" hx-swap="outerHTML">
					<input type="text" name="username" placeholder="Username">
					<input type="password" name="password" placeholder="Password">
					<button type="submit">Create</button>
				</form>
				<a href="/initial" hx-get="/initial" hx-target="closest div" hx-swap="outerHTML">Back to Login</a>
			</div>
		`))
	case "/login":
		// Handle login logic
		// After successful login, display the app view
		w.Write([]byte(`
			<div>
				<h2>Welcome, [Username]!</h2>
				<!-- App view content here -->
			</div>
		`))
	}
}
