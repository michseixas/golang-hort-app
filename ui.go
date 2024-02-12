package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// obtener los datos de la API (Probabilitat de precipitacions)
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//insertar los datos/info de la Aemet dentro del contenedor
	contenedorClimaDades := container.NewGridWithColumns(4,
		precipitacio,
		tempMax,
		tempMin,
		humitat,
	)
	//preparamos el contenedor vertical
	app.ClimaDadesContainer = contenidorClimaDades
	contenidorFinal := container.NewVBox(contenidorClimaDades)

	//agregar el contenedor al main

}
