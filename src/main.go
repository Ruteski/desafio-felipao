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

	fmt.Print("Informe a quantidade de vitórias: ")
	scanner.Scan()
	_vitorias := scanner.Text()
	vitorias, erro := strconv.ParseUint(_vitorias, 10, 64)
	if erro != nil {
		log.Fatal("Ocorreu um erro ao tentar converter as vitórias!")
	}

	fmt.Print("Informe a quantidade de derrotas: ")
	scanner.Scan()
	_derrotas := scanner.Text()
	derrotas, erro := strconv.ParseUint(_derrotas, 10, 64)
	if erro != nil {
		log.Fatal("Ocorreu um erro ao tentar converter as derrotas!")
	}

	totalVitorias := calculaRank(uint32(vitorias), uint32(derrotas))
	rank := retornaRank(totalVitorias)

	fmt.Printf("\n O Herói tem de saldo de **{%d}** está no nível de **{%s}** \n", totalVitorias, rank)
}

func calculaRank(vitorias, derrotas uint32) uint32 {
	return vitorias - derrotas
}

func retornaRank(vitorias uint32) string {
	switch {
	case vitorias <= 10:
		return "Ferro"
	case vitorias >= 11 && vitorias <= 20:
		return "Bronze"
	case vitorias >= 21 && vitorias <= 50:
		return "Prata"
	case vitorias >= 51 && vitorias <= 80:
		return "Ouro"
	case vitorias >= 81 && vitorias <= 90:
		return "Diamante"
	case vitorias >= 91 && vitorias <= 100:
		return "Lendário"
	default:
		return "Imortal"
	}
}
