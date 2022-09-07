/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// randomCmd represents the random command
var pokemonInfo = &cobra.Command{
	Use:   "pokemon",
	Short: "Pokemon Info",
	Long:  `Devuelve información básica a cerca del pokémon`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			getPokemonInfo(args[0])
		} else {
			fmt.Println(errors.New("no se ha especificado ningún pokémon"))
		}
	},
}

func init() {
	rootCmd.AddCommand(pokemonInfo)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getPokemonInfo(pokemonName string) {
	endpoint := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	responseBytes := getPokemonInfoTest(endpoint)
	pokemon := Pokemon{}

	if err := json.Unmarshal(responseBytes, &pokemon); err != nil {
		log.Printf("No se ha podido coger un pokemon - %v", err)
	}

	fmt.Println("ID: ", pokemon.ID)
	fmt.Println("Nombre: ", cases.Title(language.Und).String(pokemon.Name))
	fmt.Println("Altura: ", pokemon.Altura)
	fmt.Println("Peso: ", pokemon.Peso)
}

func getPokemonInfoTest(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("No se ha podido coger un pokemon - %v", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("No se ha podido coger un pokemon - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("No se ha podido coger un pokemon - %v", err)
	}

	return responseBytes
}
