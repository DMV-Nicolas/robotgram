package util

import (
	"fmt"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RandomString(size int) (str string) {
	digits := "abcdefghijklmnopqrstuvwxyz1234567890"
	for i := 0; i < size; i++ {
		str += string(digits[rand.Intn(len(digits))])
	}
	return str
}

// RandomID generates a random mongo ObjectID
func RandomID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// RandomUsername generates a random username.
func RandomUsername() string {
	monsters := []string{
		"Caballo", "Omnitorrinco", "Avion", "HijoDePuta", "Avi",
		"Luis", "Rodrigo", "Andres", "Santiago", "Diego", "Gustavo",
		"Juan", "Nicolas", "Cristiansito", "Juliansito", "Valentinita",
		"Sebastian", "David",
	}
	actions := []string{
		"VioladorDe", "AbusadorDe", "TerapeutaDe", "SimpDe", "DesarmadorDe",
		"DominadoPor", "EsclavizadoPor", "SexualmenteAbusadoPor", "AtraidoPor",
		"AmanteDelSexoCon", "AsesinadoPor", "TraumadoPor", "LiderDe", "CreadorDe",
	}
	victims := []string{
		"Abuelas", "Feministas", "Comunistas", "Capitalistas", "Langostas",
		"Hombres", "Jirafas", "Penes", "Duendes",
	}
	str := monsters[rand.Intn(len(monsters))]
	str += actions[rand.Intn(len(actions))]
	str += victims[rand.Intn(len(victims))]
	str += fmt.Sprint(rand.Intn(1000))
	return str
}

// RandomEmail generates a random valid email.
func RandomEmail() string {
	names := []string{
		"abuela", "cristian", "santiago",
		"feminista", "nicolas", "juan",
		"comunista", "julian", "diego",
		"capitalista", "pepito", "valentina",
		"langosta", "avi", "maria",
		"hombre", "rodolfo", "fernando",
		"jirafa", "gustavo", "proplayer",
		"pene", "rodrigo", "noob",
		"duende", "luis", "hacker",
	}
	business := []string{
		"gmail", "google", "yt",
		"unal.edu", "colsubsidio.edu", "hotmail",
		"outlook", "sakura",
	}
	countries := []string{
		"com", "co", "ar", "es", "us",
		"br", "cl", "pe", "mx", "uy",
	}
	str := fmt.Sprint(rand.Intn(100))
	str += names[rand.Intn(len(names))]
	str += fmt.Sprint(rand.Intn(10000))
	str += "@"
	str += business[rand.Intn(len(business))] + "."
	str += countries[rand.Intn(len(countries))]
	return str
}

// RandomPassword generates a random password with the given size.
func RandomPassword(size int) string {
	return RandomString(size)
}

// RandomImage generates a random image.
func RandomImage() string {
	n1 := rand.Intn(100) + 400
	n2 := rand.Intn(100) + 400
	image := fmt.Sprintf("https://random.imagecdn.app/%d/%d", n1, n2)
	return image
}

// RandomImages generates random images.
func RandomImages(n int) []string {
	images := make([]string, n)
	for i := 0; i < n; i++ {
		images[i] = RandomImage()
	}
	return images
}

// RandomDescription generates a random description with the given size.
func RandomDescription(size int) string {
	return RandomString(size)
}
