package api

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strconv"
	"strings"
)

func (z ZimplerPageContext) getAll() ([]Person, error) {
	result := []Person{}
	var error *Error

	c := colly.NewCollector()

	c.OnHTML("*[id=\"top.customers\"]", func(e *colly.HTMLElement) {

		// Prepare date
		whitespaces := regexp.MustCompile(`\s+`)
		read_line := strings.ReplaceAll(e.Text, "\n", "")
		read_line = whitespaces.ReplaceAllString(read_line, " ")[18:]
		words := strings.Fields(read_line)

		// Parse string to Person array structure
		person := Person{}
		for index, element := range words {
			if index%3 == 0 && index != 0 {
				result = append(result, person)
				person = Person{}
			}

			switch index % 3 {
			case 0:
				person.Name = element
			case 1:
				person.Candy = element
			case 2:
				number, err := strconv.Atoi(element)
				if err != nil {
					error = &Error{Err: errors.New(fmt.Sprintf("Element is not a number - %v", element))}
					break
				}
				person.Eaten = number
			default:
				error = &Error{Err: errors.New("Unknown structure")}
				break
			}
		}
		result = append(result, person)
	})

	if err := c.Visit(z.Url); err != nil {
		return nil, err
	}

	if error != nil {
		return nil, error
	}

	return result, nil
}
