package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
	Plot     string
}

func main() {
	fmt.Println("Running Server... \nAt port http://localhost:8000")
	handle1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Blade Runner", Director: "Ridley Scott", Plot: "Blade Runner is a 1982 science fiction film directed by Ridley Scott. The film is set in a dystopian future Los Angeles of 2019, in which synthetic humans known as replicants are bio-engineered to work on off-world colonies. The plot follows a fugitive replicant police officer who is hunting down and trying to eliminate a group of rogue replicants."},
				{Title: "The Matrix", Director: "The Wachowskis", Plot: "The Matrix is a 1999 science fiction action film written and directed by the Wachowskis. It depicts a dystopian future in which humanity is unknowingly trapped inside a simulated reality, the Matrix, created by intelligent machines to distract humans while using their bodies as an energy source."},
				{Title: "Star Wars: Episode IV - A New Hope", Director: "George Lucas", Plot: "Star Wars: Episode IV - A New Hope is a 1977 space opera film written and directed by George Lucas. It is the first film in the original Star Wars trilogy and follows Luke Skywalker's journey to become a Jedi and join the Rebel Alliance in their struggle against the evil Galactic Empire."},
				{Title: "Inception", Director: "Christopher Nolan", Plot: "Inception is a 2010 science fiction action film written and directed by Christopher Nolan. The film follows a professional thief who steals information by infiltrating the subconscious of his targets. He is offered a chance to have his criminal history erased as payment for the implantation of another person's idea into a target's subconscious."},
				{Title: "E.T. the Extra-Terrestrial", Director: "Steven Spielberg", Plot: "E.T. the Extra-Terrestrial is a 1982 science fiction film directed by Steven Spielberg. It tells the story of Elliott, a boy who befriends an extraterrestrial being stranded on Earth and helps him return to his home planet."},
				{Title: "Back to the Future", Director: "Robert Zemeckis", Plot: "Back to the Future is a 1985 science fiction film directed by Robert Zemeckis. It follows the adventures of Marty McFly, who is accidentally sent back in time to 1955 in a time-traveling DeLorean invented by his eccentric scientist friend Doc Brown."},
				{Title: "2001: A Space Odyssey", Director: "Stanley Kubrick", Plot: "2001: A Space Odyssey is a 1968 science fiction film directed by Stanley Kubrick. The film follows a voyage to Jupiter with the sentient computer HAL-9000 on board the spacecraft Discovery One."},
				{Title: "Interstellar", Director: "Christopher Nolan", Plot: "Interstellar is a 2014 science fiction film directed by Christopher Nolan. The story revolves around a group of astronauts who travel through a wormhole near Saturn in search of a new home for humanity."},
				{Title: "The Terminator", Director: "James Cameron", Plot: "The Terminator is a 1984 science fiction film directed by James Cameron. It follows a cyborg assassin, played by Arnold Schwarzenegger, who is sent back in time to kill Sarah Connor, whose son will one day become a savior against machines."},
				{Title: "Jurassic Park", Director: "Steven Spielberg", Plot: "Jurassic Park is a 1993 science fiction adventure film directed by Steven Spielberg. The film is set on the fictional island of Isla Nublar, where a billionaire philanthropist and a team of genetic scientists have created a wildlife park of cloned dinosaurs."},
				{Title: "The Martian", Director: "Ridley Scott", Plot: "The Martian is a 2015 science fiction film directed by Ridley Scott. It follows an astronaut stranded on Mars who must figure out how to survive while awaiting rescue."},
				{Title: "Star Trek: The Motion Picture", Director: "Robert Wise", Plot: "Star Trek: The Motion Picture is a 1979 science fiction film directed by Robert Wise. The film is based on the television series Star Trek and follows the crew of the starship USS Enterprise as they investigate a mysterious and powerful alien entity."},
				{Title: "Avatar", Director: "James Cameron", Plot: "Avatar is a 2009 science fiction film directed by James Cameron. The story is set in the mid-22nd century when humans are colonizing Pandora, a lush habitable moon of a gas giant in the Alpha Centauri star system, in order to mine the mineral unobtanium."},
				{Title: "The Day the Earth Stood Still", Director: "Robert Wise", Plot: "The Day the Earth Stood Still is a 1951 science fiction film directed by Robert Wise. The film depicts the arrival of an extraterrestrial being, Klaatu, who comes to Earth with a powerful robot, Gort, to deliver an important message to humanity."},
				{Title: "Close Encounters of the Third Kind", Director: "Steven Spielberg", Plot: "Close Encounters of the Third Kind is a 1977 science fiction drama film written and directed by Steven Spielberg. The story follows a man who experiences a close encounter with an unidentified flying object and becomes obsessed with uncovering the truth."},
				{Title: "The Fifth Element", Director: "Luc Besson", Plot: "The Fifth Element is a 1997 science fiction action film directed by Luc Besson. The film is set in the 23rd century and follows the quest to find the Fifth Element, a being with the power to save the world from an impending evil."},
				{Title: "A.I. Artificial Intelligence", Director: "Steven Spielberg", Plot: "A.I. Artificial Intelligence is a 2001 science fiction drama film directed by Steven Spielberg. The story is set in a future where highly advanced humanoid robots with emotions are created to serve humans."},
				{Title: "The War of the Worlds", Director: "Byron Haskin", Plot: "The War of the Worlds is a 1953 science fiction film directed by Byron Haskin. It is an adaptation of H.G. Wells' classic novel and depicts an alien invasion of Earth."},
				{Title: "The Time Machine", Director: "George Pal", Plot: "The Time Machine is a 1960 science fiction film directed by George Pal. It follows the story of a Victorian-era inventor who creates a machine that enables him to travel through time."},
				{Title: "District 9", Director: "Neill Blomkamp", Plot: "District 9 is a 2009 science fiction action film directed by Neill Blomkamp. The story unfolds in an alternative reality where extraterrestrial refugees are forced to live in slum-like conditions in Johannesburg, South Africa."},
			},
		}
		tmpl.Execute(w, films)
	}

	handle2 := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX request received")
		log.Print(r.Header.Get("HX-Request"))

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		plot := r.PostFormValue("plot")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-card-element", Film{Title: title, Director: director, Plot: plot})
	}

	/* routers */
	http.HandleFunc("/", handle1)
	http.HandleFunc("/add-film/", handle2)

	/* web server error handler */
	log.Fatal(http.ListenAndServe(":8000", nil))
}
