package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {
	// Create library service
	libraryService := services.NewLibrary()
	
	// Create controller with the service
	controller := controllers.NewLibraryController(libraryService)
	
	// Start the console interface
	controller.Start()
}