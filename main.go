package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/widget"
)

//Struct per enmagatzemar les configuracions de l' app

type Config struct { //donde guardaremos todos los ajustes de nuestra aplicacion
	App      fyne.App    //App Fyne para la GUI
	InfoLog  *log.Logger //Log de acciones ordinarias, para tener un detalle de lo que ha hecho el usuario (muy bueno para tener un seguimiento de lo que ha pasado)
	ErrorLog *log.Logger //Log de errors
	MainWindow fyne.Window //Enmagatzemar la finestra principal. Si se cierra la ventana principal, se cierran tambien todas las demás ventanas que pueda haber abiertas en el programa
}

//Nuestra intencion es que tendremo spestañas en nuestra app, y cuando cerremos una la app recordará la info que habñiamos dejado alli y al regresar 
//podremos seguir viendo la pestaña tal y como la dejamos

// Crear una variable de tipo Config que representará el canvas donde se construirá la app, y así situar la configuración de la app
var myApp Config

func main() {
	// Crearem la App
	fyneApp := app.NewWithID("cat.cibernarium.ecohortapp")
	myApp.App = fyneApp

	//Els logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Lshortfile)


	//Conexión am la BBDD

	//Check i creación del repositorio

	//Definirem el tamany de la pantall
	//Decir todo lo que quiero poner en la ventana, todas las etapas. 
	myApp.MainWindow = fyneApp.NewWindow("Eco Hort App")
	myApp.MainWindow.Resize(fyne.NewSize(800,500)) //definimos el tamaño de la finestra
	myApp.MainWindow.SetFixedSize(true) //fijo, no se podrá redimensionar
	myApp.MainWindow.SetMaster() //Se establece el rol final: es la ventana principal

	myApp.makeUI()

	// Mostrarem i executarem l' app
	myApp.MainWindow.ShowAndRun()
}





// func main() {
// 	a := app.New() //Estamos creando una aplicación y la guardamos en la variable a
// 	w := a.NewWindow("Hola Mundo")

// 	w.SetContent(widget.NewLabel("Texte exemple")) //widget sale del segundo import)
// 	w.ShowAndRun()
// }
