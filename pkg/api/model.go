package api

type Error struct {
	Err error
}

type Person struct {
	Name  string
	Candy string
	Eaten int
}

type PersonFavouriteCandies struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

type CandyStoreRepository interface {
	getAll() ([]Person, error)
}

type ZimplerPageContext struct {
	Url string
}
