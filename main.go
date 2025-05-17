package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/MatheusBentoP/PrimeiroProjetoEmGO.git/models"
	"github.com/gin-gonic/gin"
)

var pessoa = []models.Pessoa{}

func main() {
	carregaPessoa()
	router := gin.Default()
	router.GET("/pessoa", getPessoa)
	router.POST("/pessoa", postPessoa)
	router.GET("/pessoa/:ID", getPessoaById)
	//router.DELETE("/pessoa", deletePessoa)
	router.Run()
}

func carregaPessoa() {
	file, err := os.Open("dados/pessoa.json")
	if err != nil {
		fmt.Println("Error file", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(pessoa); err != nil {
		fmt.Println("Decoding Error")
	}

}

func getPessoa(c *gin.Context) {
	c.JSON(200, gin.H{
		"Pessoa": pessoa,
	})

}

func postPessoa(c *gin.Context) {
	var newPessoa models.Pessoa
	if err := c.ShouldBindJSON(&newPessoa); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	newPessoa.ID = len(pessoa) + 1
	pessoa = append(pessoa, newPessoa)
	savePessoa()
	c.JSON(201, newPessoa)
}

func savePessoa() {
	file, err := os.Create("dados/pessoa.json")
	if err != nil {
		fmt.Println("Error file", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pessoa); err != nil {
		fmt.Println("Encondig Error")
	}

}

func getPessoaById(c *gin.Context) {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	for _, p := range pessoa {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pessoa não encontrada"})

}
