package main

import (
	"encoding/json"
	"fmt"
	"zimpler-test/pkg/api"
)

func main() {
	// would be cobra and viper for CLI and coniguration file for URL, like
	// if url := viper.GetString("https://candystore.zimpler.net/"); url == "" {
	//	 panic("No ZIMPLERURL field in configuration")
	// }

	zimpler := api.ZimplerPageContext{Url: "https://candystore.zimpler.net/"}
	res, _ := api.GetFavouriteCandiesService(zimpler)

	resJson, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resJson))
}
