package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nivelHeroi string

	fmt.Print("Informe o nome do herói: ")
	scanner.Scan()
	var nomeHeroi string = scanner.Text()

	fmt.Print("Informe o xp do herói: ")
	scanner.Scan()
	xpHeroiString := scanner.Text()
	xpHeroi, erro := strconv.ParseUint(xpHeroiString, 10, 64)

	if erro != nil {
		log.Fatal("Ocorreu um erro ao tentar converter o xp do herói!")
	}

	switch {
	case xpHeroi <= 1000:
		nivelHeroi = "Ferro"
	case xpHeroi >= 1001 && xpHeroi <= 2000:
		nivelHeroi = "Bronze"
	case xpHeroi >= 2001 && xpHeroi <= 5000:
		nivelHeroi = "Prata"
	case xpHeroi >= 5001 && xpHeroi <= 7000:
		nivelHeroi = "Ouro"
	case xpHeroi >= 7001 && xpHeroi <= 8000:
		nivelHeroi = "Platina"
	case xpHeroi >= 8001 && xpHeroi <= 9000:
		nivelHeroi = "Ascendente"
	case xpHeroi >= 9001 && xpHeroi <= 10000:
		nivelHeroi = "Imortal"
	case xpHeroi >= 10001:
		nivelHeroi = "Radiante"

	}

	fmt.Printf("\n O Herói de nome **{%s}** está no nível de **{%s}** \n", nomeHeroi, nivelHeroi)
}
