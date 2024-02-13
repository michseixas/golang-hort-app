package main

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"io"
	"os"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pronosticTab() *fyne.Container {
	//invocar el metodo inferior
	grafic := app.obtenirGrafic()
	contenidorGrafic := container.NewVBox(grafic)
	app.PronosticGraficContainer = contenidorGrafic
	return contenidorGrafic
}

func (app *Config) obtenirGrafic() *canvas.Image {
	//Definirem la url
	url := "https://my.meteoblue.com/images/meteogram?temperature_units=C&wind_units=kmh&precipitation_units=mm&winddirection=3char&iso2=pt&lat=39.4667&lon=-8.2&asl=133&tz=Europe%2FLisbon&dpi=72&apikey=jhMJTOUVRNvs25m4&lang=es&location_name=Abrantes&windspeed_units=kmh&sig=621e401f6fd3a0aa6838c68f3601648c"
	var img *canvas.Image

	err := app.descarregaArxiu(url, "pronostic.png")
	if err != nil {
		//codigo para error de conexion
		canvas.NewImageFromResource(resourceNodisponiblePng)
	}else{
		img = canvas.NewImageFromFile("pronostic.png")
	}

	//Definimos el tamaño
	img.SetMinSize(fyne.Size{
		Width: 770,
		Height: 480,
	})

	//Definir la actuación versus canvas
	img.FillMode = canvas.ImageFillOriginal
	
	return nil
}

func (app *Config) descarregaArxiu(url string, nomArxiu string) error {
	resposta, err := app.HTTPClient.Get(url)
	if err != nil {
		return err
	}

	if resposta.StatusCode != 200 {
		return errors.New("Codigo de respuesta erroneo.")
	}

	binari, err := io.ReadAll(resposta.Body)
	if err != nil {
		return err
	}
	defer resposta.Body.Close()

	//descodificamos el flujo recibido
	img, _, err := image.Decode(bytes.NewReader(binari))
	if err != nil {
		return err
	}

	//obtenemos la sortida y creamos el archivo de cero
	out, err := os.Create(fmt.Sprintf("./%s", nomArxiu))
	if err != nil {
		return err
	}


	//Codificar el png
	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil

}

