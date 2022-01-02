package main

func main() {
	// list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	// list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	findRestaurant(list1, list2)
}

func findRestaurant(list1 []string, list2 []string) []string {
	listToMap := make(map[string]int, len(list1))
	// result := make(map[string]int)
	result := []string{}
	for i, v := range list1 {
		listToMap[v] = i
	}
	for i, v := range list2 {
		if _, ok := listToMap[v]; ok {
			listToMap[v] = listToMap[v] + i
			if len(result) == 0 || listToMap[v] == listToMap[result[0]] {
				result = append(result, v)
			} else if listToMap[v] < listToMap[result[0]] {
				result = []string{v}
			}
		}
	}
	return result
}
