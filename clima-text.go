package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

//hacemos una funcion que retorna 4 elementos de texto con fyne y que realizará una interferencia en la estructura Config
func (app *Config) getClimaText() (*canvas.Text,*canvas.Text,*canvas.Text,*canvas.Text ){

	//Definimos las variables
	var g Diaria //variable que contienen toda la información de la climatologia de tipo Diaria
	var precipitacio, tempMax, tempMin, humitat *canvas.Text //4 variables de tipo canvas.Text

	prediccio, err := g.GetPrediccions()
	//controlar error
	if err != nil {
		//se definen los colores de los textos segundo los valores mostrados. 
		//logica por error de lectura api (en el caso que se haya generado un error al invocar la función de obtener predicciones)
		gris := color.NRGBA{R:155, G:155, B:155, A:255}
		precipitacio = canvas.NewText("Precipitació: No visible", gris)
		tempMax = canvas.NewText("Temp. Max: No visible", gris)
		tempMin = canvas.NewText("Temp. Min: No visible", gris)
		humitat = canvas.NewText("Humitat: No visible", gris)
	}else{
		//Si funciona 
		displayColor := color.NRGBA{R:155, G:155, B:155, A:255}

		if prediccio.ProbPrecipitacio < 50 {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255} //Color per quan el valor és inferior al 50%
		}

		//Preparem els textos
		precipitacioTxt := fmt.Sprintf("Precipitació: %d%%", prediccio.ProbPrecipitacio)
		tempMaxTxt := fmt.Sprintf("Precipitació: %d%", prediccio.TemperaturaMax)
		tempMinTxt := fmt.Sprintf("Precipitació: %d%", prediccio.TemperaturaMin)
		humitatTxt := fmt.Sprintf("Precipitació: %d%%", prediccio.HumitatRelativa)

		//Muntage dels canvas text
		precipitacio = canvas.NewText(precipitacioTxt, displayColor) 
		tempMax = canvas.NewText(tempMaxTxt, nil)
		tempMin = canvas.NewText(tempMinTxt, nil)
		humitat = canvas.NewText(humitatTxt, nil)

	}

	precipitacio.Alignment = fyne.TextAlignLeading
	tempMax.Alignment =fyne.TextAlignCenter
	tempMin.Alignment = fyne.TextAlignCenter
	humitat.Alignment = fyne.TextAlignTrailing

	return precipitacio, tempMax, tempMin, humitat
}
