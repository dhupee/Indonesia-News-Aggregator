package main

import (
	"fmt"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"
)


var kompasCategoryList = []string{
	"all",
	"nasional",
	"regional",
	"megapolitan",
	"global",
	"tren",
	"health",
	"food",
	"edukasi",
	"money",
	"properti",
	"bola",
	"travel",
	"otomotif",
	"sains",
	"hype",
	"jeo",
	"skola",
	"stori",
	"konsultasihukum",
	"wiken",
	"headline",
	"terpopuler",
	"sorotan",
	"topik",
	"advertorial",
}

func KompasGetNewsList(rawHtml string) []string{
	newsList := []string{}

	return newsList
}

func KompasCategoryCheck(category []string, subcategory string) bool {
	var results bool
	if utils.IsInSlice(subcategory, category){
		results = true
	} else{
		results = false
	}
	return results
}

func main(){
	results := KompasCategoryCheck(bola, "liga-italia")
	fmt.Println(results)

	results2 := KompasCategoryCheck(bola, "liga-belanda")
	fmt.Println(results2)
}
