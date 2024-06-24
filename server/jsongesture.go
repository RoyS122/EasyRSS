package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile(filename string, content string) error {
	// Ouvrir le fichier en mode écriture. Si le fichier n'existe pas, créez-le.
	// Le fichier est ouvert avec les droits en écriture seulement pour l'utilisateur.
	dir, _ := os.Getwd()
	file, err := os.OpenFile(dir+filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		return fmt.Errorf("erreur à l'ouverture du fichier: %v", err)
	}
	defer file.Close() // S'assurer que le fichier sera fermé à la fin.

	// Écrire le contenu dans le fichier.
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier: %v", err)
	}

	return nil // Aucune erreur, retourner nil.
}

func getFluxFromJson(filename string) (fluxlist []Flux) {
	// Emplacement du fichier JSON contenant les profils des joueurs.
	dir, _ := os.Getwd()
	// Ouvrir le fichier.
	file, err := os.Open(dir + filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	// Lire le contenu du fichier.
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	// Décoder le contenu JSON dans une slice de Player.
	if err := json.Unmarshal(bytes, &fluxlist); err != nil {
		return nil
	}
	// fmt.Println(players, bytes)
	return fluxlist
}

func getCategoriesFromJson(filename string) (categories []Categorie) {

	dir, _ := os.Getwd()
	// Ouvrir le fichier.
	file, err := os.Open(dir + filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	// Lire le contenu du fichier.
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	// Décoder le contenu JSON dans une slice de Player.
	if err := json.Unmarshal(bytes, &categories); err != nil {
		return nil
	}
	// fmt.Println(players, bytes)
	
	return categories
}