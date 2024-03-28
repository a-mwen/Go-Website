package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define handler function for the homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> Soccer Fans Hub </title>
			</head>	
			<body>
				<h1> Welcome to Soccer Fans Hub </h1>
				<p> Explore the exciting world of soccer with us!</p>
				<img src="https://i.ytimg.com/vi/NQ7w-aNZpXE/maxresdefault.jpg" height="250" width="250">
			</body>
		</html>
		`)
	})

	// Define handler function for the "About" page
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> About Soccer Fans Hub </title>
			</head>	
			<body>
				<h1> About Soccer Fans Hub </h1>
				<p> Learn more about our passion for soccer. Maybe one day you can be a famous soccer player, making all the big bucks!!!</p>
				<img src="https://th.bing.com/th/id/OIP.ynYySJ5RZ2-hACmJ7TgwBwHaFc?w=257&h=189&c=7&r=0&o=5&dpr=1.3&pid=1.7" height="250" width="250">
			</body> 
		</html>
		`)
	})

	// Define handler function for the "Hobbies" page
	http.HandleFunc("/hobbies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title> Soccer Hobbies </title>
			</head>	
			<body>
				<h1>  Hobbies </h1>
				<p>Some of my hoobies as soccer player is:  Watching Movies, Listening to Music and  Occassionally gaming</p>  
				<img src="https://th.bing.com/th/id/OIP.w78nkjr2dz2QsFkC-ZbZAwHaEo?rs=1&pid=ImgDetMain" height="250" width="250">
			</body>
		</html>
		`)
	})

	// Start the server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
