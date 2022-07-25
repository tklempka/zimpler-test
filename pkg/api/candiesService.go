package api

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

func GetFavouriteCandiesService(ctx CandyStoreRepository) ([]PersonFavouriteCandies, error) {
	if ctx == nil {
		return nil, &Error{Err: errors.New(fmt.Sprintf("Nil reference context - %v", reflect.TypeOf(ctx)))}
	}

	data, err := ctx.getAll()
	if err != nil {
		return nil, err
	}

	result, err := selectFavouriveSnacks(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func selectFavouriveSnacks(p []Person) ([]PersonFavouriteCandies, error) {
	if p == nil {
		return nil, &Error{Err: errors.New(fmt.Sprintf("Nil reference - %v", reflect.TypeOf(p)))}
	}

	personFavouriteSnacks := []PersonFavouriteCandies{}

	// sort people by Name firstly and then by Candy
	sort.Slice(p, func(i, j int) bool {
		if p[i].Name != p[j].Name {
			return p[i].Name < p[j].Name
		}
		return p[i].Candy < p[j].Candy
	})

	lastPerson := p[0]
	totalCandyEaten := 0
	currentEatenSpecificCandy := 0
	eatenFavouriteCandy := 0
	currentFavouriteCandyName := p[0].Candy

	for _, person := range p {
		// when we have a new person, we put it into array andreset values
		if lastPerson.Name != person.Name {
			favCandy := PersonFavouriteCandies{
				Name:           lastPerson.Name,
				FavouriteSnack: currentFavouriteCandyName,
				TotalSnacks:    totalCandyEaten,
			}
			personFavouriteSnacks = append(personFavouriteSnacks, favCandy)
			totalCandyEaten = 0
			currentEatenSpecificCandy = 0
			eatenFavouriteCandy = 0
			currentFavouriteCandyName = person.Candy
		}

		// next row is the same candy, otherwise I check if the
		// currentEatenSpecificCandy is greater and assign new favourite candy
		if currentFavouriteCandyName == person.Candy {
			currentEatenSpecificCandy = currentEatenSpecificCandy + person.Eaten
		} else if currentEatenSpecificCandy > eatenFavouriteCandy {
			eatenFavouriteCandy = currentEatenSpecificCandy
			currentFavouriteCandyName = lastPerson.Candy
		}

		totalCandyEaten = totalCandyEaten + person.Eaten
		lastPerson = person
	}

	// append last person
	favCandy := PersonFavouriteCandies{Name: lastPerson.Name, FavouriteSnack: currentFavouriteCandyName, TotalSnacks: totalCandyEaten}
	personFavouriteSnacks = append(personFavouriteSnacks, favCandy)

	return personFavouriteSnacks, nil
}
