package main

import (
	"fmt"
	"io/ioutil" // per leggere il corpo della risposta
	"net/http" // pacchetto necessario per creare la richiesta GET
	"os" //pacchetto della libreria standard di Go che permette di interagire con il Sistema Operativo
)

func main() {
	// URL dell'API JSONPlaceholder
	url := "https://jsonplaceholder.typicode.com/posts"

	// Effettua una richiesta GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Errore nella richiesta:", err)
		return
	}
	defer resp.Body.Close()

	// Legge il corpo della risposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Errore nella lettura del corpo della risposta:", err)
		return
	}

	// Stampa la risposta
	//fmt.Println("Risposta dell'API:")
	//fmt.Println(string(body))

	// Crea un file per salvare i dati
	fileName := "posts.json"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println ("Errore nella creazione del file: ", err)
		return
	}
	defer file.Close()

	//Scrive i dati nel file
	_, err = file.Write(body)
	if err != nil {
		fmt.Println ("Errore nella scrittura del file: ", err)
		return
	}

	fmt.Printf("I dati sono stati scritti con successo nel file %s\n", fileName)

}
