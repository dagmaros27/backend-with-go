package main

import (
	c "example.com/library-managment/controllers"
	s "example.com/library-managment/services"
	m "example.com/library-managment/models"

)

func main() {

	library := s.NewLibrary()
	library.AddMember(m.Member{ID: 1, Name: "Alice"}) 
	library.AddMember(m.Member{ID: 2, Name: "Bob"})  

	controller := c.NewLibraryController(&library)
	controller.Run()
	}
	