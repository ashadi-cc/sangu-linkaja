package linkaja

import "fmt"

func GenerateItems(items []PublicTokenItemRequest) string {
	var is string
	for i, v := range items {
		if i > 0 {
			is = is + ","
		}
		is = is + fmt.Sprintf("[\"%v\", \"%v\", \"%v\"]", v.Name, v.Price, v.Quantity)
	}

	return fmt.Sprintf("[%v]", is)
}
