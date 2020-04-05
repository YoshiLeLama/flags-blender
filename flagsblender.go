package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"time"

	"flagsblender.fr/yoshi/flags/utils"
)

func main() {
	// On récupère les noms des pays dont l'utilisateur veut mélanger les drapeaux
	var firstFlagName string
	var secondFlagName string

	for firstFlagName == "" {
		fmt.Println("Veuillez rentrer le nom du premier pays :")
		fmt.Scanln(&firstFlagName)
	}

	for secondFlagName == "" {
		fmt.Println("Veuillez rentrer le nom du second pays :")
		fmt.Scanln(&secondFlagName)
	}

	// On récupère les images correspondant aux drapeaux
	flagReader, err := os.Open("flags/" + strings.ToLower(firstFlagName) + ".png")
	if err != nil {
		utils.LogFatal(err)
	}

	// On prévoit la fermeture du reader
	defer flagReader.Close()

	firstFlag, err := png.Decode(flagReader)
	if err != nil {
		utils.LogFatal(err)
	}

	flagReader, err = os.Open("flags/" + strings.ToLower(secondFlagName) + ".png")
	if err != nil {
		utils.LogFatal(err)
	}

	secondFlag, err := png.Decode(flagReader)
	if err != nil {
		utils.LogFatal(err)
	}

	beginTime := time.Now()

	minWidth, minHeight := utils.AdaptFlags(&firstFlag, &secondFlag)

	// On crée une nouvelle image RGBA pour stocker le nouveau drapeau
	newFlag := image.NewRGBA(image.Rectangle{Max: image.Point{
		minWidth, minHeight,
	}})

	// On initialise la variable qui permet de stocker la couleur de chaque pixel
	var color = color.RGBA{A: 255}

	for x := 0; x < minWidth; x++ {
		for y := 0; y < minWidth; y++ {
			// On récupère la couleur du pixel de coordonnées (x, y) de chaque drapeau
			firstR, firstG, firstB, _ := firstFlag.At(x, y).RGBA()
			secondR, secondG, secondB, _ := secondFlag.At(x, y).RGBA()

			// On fait la moyenne de chacune des composantes de la couleur (rouge, vert et bleu)
			color.R = uint8((firstR + secondR) / 2)
			color.G = uint8((firstG + secondG) / 2)
			color.B = uint8((firstB + secondB) / 2)

			// On stocke la couleur moyenne dans le nouveau drapeau
			newFlag.SetRGBA(x, y, color)
		}
	}

	// On crée le fichier destiné à stocker le nouveau drapeau (ici, il sera créé à la racine de l'application)
	newFlagFile, err := os.Create("NewFlag.png")
	if err != nil {
		utils.LogFatal(err)
	}
	// On prévoit la fermeture du fichier
	defer newFlagFile.Close()

	err = png.Encode(newFlagFile, newFlag)
	if err != nil {
		utils.LogFatal(err)
	}

	// On confirme à l'utilisateur la réussite de l'opération
	fmt.Println("Successfully create a mix of the flags of " + firstFlagName + " and " + secondFlagName)
	fmt.Printf("It takes %d ms", time.Now().Unix()-beginTime.Unix())
}
