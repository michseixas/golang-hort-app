package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) makeUI() {
	// obtener los datos de la API (Probabilitat de precipitacions)
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//insertar los datos/info de la Aemet dentro del contenedor
	climaDadesContent := container.NewGridWithColumns(4,
		precipitacio,
		tempMax,
		tempMin,
		humitat,
	) //definimos un contenedor con 4 columnas
	//preparamos el contenedor vertical

	app.ClimaDadesContainer = climaDadesContent


	//obtenim la barra d'eines o toolbar
	toolBar := app.getToolBar(app.MainWindow)

	//obtenim les pestanyes de l'aplicació
	tabs := container.NewAppTabs( //Definim un contenidor per les pestanyes i dins afegim cada una de les pestanyes amb icones
		container.NewTabItemWithIcon("Pronòstic", theme.HomeIcon(), canvas.NewText("El contingut dels pronostics dels propers 4 dies els situarem aqui", nil)),
		container.NewTabItemWithIcon("Diari Meteorològic", theme.InfoIcon(), canvas.NewText("El contingut dels valors enregistrats els situarem aqui", nil)),
	)

	//afegir el contenidor a la finestra
	finalContent := container.NewVBox(climaDadesContent, toolBar, tabs) //Definim un nou contenidor i que afegirem al canvas general.

	//Invoquem la pàgina principal i fem servir el mètode SetContent per afegir el contenidor
	app.MainWindow.SetContent(finalContent)
}
