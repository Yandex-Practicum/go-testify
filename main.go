package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var cafes = map[string][]string{
	"moscow": {"CoffeeShop1", "CoffeeShop2", "CoffeeShop3", "CoffeeShop4"},
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")
	city := r.URL.Query().Get("city")

	// Проверка на наличие города
	cityCafes, ok := cafes[city]
	if !ok {
		http.Error(w, "wrong city value", http.StatusBadRequest)
		return
	}

	// Преобразование count в число
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 0 {
		http.Error(w, "invalid count", http.StatusBadRequest)
		return
	}

	// Ограничение количества
	if count > len(cityCafes) {
		count = len(cityCafes)
	}

	selected := cityCafes[:count]
	fmt.Fprint(w, strings.Join(selected, ","))
}
