/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Pokemon random",
	Long:  `Descripción larga de pokemon random`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomPokemon()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getRandomPokemon() {
	rand.Seed(time.Now().UnixNano())
	randomNum := fmt.Sprintf("%d", rand.Intn(905))
	endpoint := "https://pokeapi.co/api/v2/pokemon/" + randomNum
	responseBytes := getPokemonData(endpoint)
	pokemon := Pokemon{}

	if err := json.Unmarshal(responseBytes, &pokemon); err != nil {
		log.Printf("No se ha podido coger un pokemon - %v", err)
	}

	fmt.Println(pokemon.Name)
}

func getPokemonData(baseAPI string) []byte {
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
