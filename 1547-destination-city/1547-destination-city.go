func destCity(paths [][]string) string {
    cities := map[string]string{}

    for _, path := range paths {
        cities[path[0]] = path[1]
    }

    var res string
    for _, path := range paths {
        if _, ok := cities[path[1]]; !ok {
            res = path[1]
            break
        }
    }

    return res
}