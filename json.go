package main

import (
	"encoding/json"
	"fmt"
)

type produtoJson struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func gerarJson() {
	//SEREALIZE
	produto := produtoJson{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	jsonProduto, _ := json.Marshal(produto)

	fmt.Println(string(jsonProduto))

	//DESERIALIZER

	var produto2 produtoJson
	jsonString := `{"id":2,"nome":"Caneta","preco":9.9,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &produto2)

	fmt.Println(produto2.Tags[1])

}
