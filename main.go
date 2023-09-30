package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const api string = "https://jsonplaceholder.typicode.com/todos"

var todos []Todo

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	// Faz a requisição HTTP GET e armazena o Body da resposta
	response, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Deserializa os dados da resposta em formato JSON baseado na struct criada
	json.Unmarshal(responseData, &todos)

	// Cria o arquivo CSV para armazenamento dos dados
	csvFile, _ := os.Create("./teste.csv")
	writer := csv.NewWriter(csvFile)

	// Loop para iterar sobre os dados da resposta dentro do slice "todos"
	// Converter os dados a API para string e fazer o append dentro do slice row
	// Escrever no arquivo CSV
	for _, todo := range todos {
		var row []string
		row = append(row, strconv.Itoa(todo.UserID))
		row = append(row, strconv.Itoa(todo.ID))
		row = append(row, todo.Title)
		row = append(row, strconv.FormatBool(todo.Completed))

		writer.Write(row)
	}
	writer.Flush()

}
