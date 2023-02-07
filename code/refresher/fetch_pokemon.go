package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func isItemInList(item string, list []string) bool {
	for _, key := range list {
		if key == item {
			return true
		}
	}
	return false
}

type Pokemon struct {
	name   string
	order  string
	height string
	weight string
}

func main() {
	// var baseEndpoint string = "https://pokeapi.co/api/v2/pokemon/%d"
	// var parsedPokemon = []Pokemon{}
	// for idx := 0; idx < 1000; idx++ {
	// 	currentEndpoint := fmt.Sprintf(baseEndpoint, idx+1)
	// 	response, err := http.Get(currentEndpoint)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(1)
	// 	}
	// 	responseData, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	var responseJson map[string]interface{}

	// 	err = json.Unmarshal([]byte(responseData), &responseJson)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	var currentPokemon = Pokemon{}
	// 	keysOfInterest := []string{"name", "order", "height", "weight"}
	// 	for key, value := range responseJson {
	// 		if isItemInList(key, keysOfInterest) {
	// 			switch key {
	// 			case "name":
	// 				currentPokemon.name = fmt.Sprintf("%v", value)
	// 			case "order":
	// 				currentPokemon.order = fmt.Sprintf("%v", value)
	// 			case "height":
	// 				currentPokemon.height = fmt.Sprintf("%v", value)
	// 			case "weight":
	// 				currentPokemon.weight = fmt.Sprintf("%v", value)
	// 			}

	// 		}
	// 	}
	// 	parsedPokemon = append(parsedPokemon, currentPokemon)
	// }

	// // Iterate over the created pokemon List and write to CSV file
	const outputCSVFile string = "data/pokemon_details.csv"
	// csvFile, err := os.Create(outputCSVFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer csvFile.Close()
	// csvWriter := csv.NewWriter(csvFile)

	// csvWriter.Write([]string{"Name", "Order", "Height", "Weight"})
	// for i := 0; i < len(parsedPokemon); i++ {
	// 	var row = []string{parsedPokemon[i].name, parsedPokemon[i].order,
	// 		parsedPokemon[i].weight, parsedPokemon[i].height}
	// 	_ = csvWriter.Write(row)
	// }
	// csvWriter.Flush()

	// Read input CSV File and filter where height < 10 and weight between 100 and 175
	csvFile, err := os.Open(outputCSVFile)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		pokemon := Pokemon{record[0], record[1], record[2], record[3]}
		//name, order, weight, height := record
		fmt.Println(pokemon)

	}
}
