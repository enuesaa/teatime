package routes

import (
    "api/src/controller"
)

func Routes(url string) string {
    routesDict := getRoutes()
    v, isExist := routesDict[url]
    if !isExist {
        return "error"
    }
    return v()
}

func getRoutes() map[string]func() string {
    return map[string]func() string {
        "/": func() string {
            return "a"
        },
        "/books": func() string {
            return books.GetIndex()
        },
    }
}
