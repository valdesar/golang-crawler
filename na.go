package main

import( 
    "os"
    "bufio"
    "fmt"
    "log"
     "io/ioutil"
    "net/http"
//    "reflect"
)

func getreq(siteNameQ string) {

fmt.Println(siteNameQ)

resp, err := http.Get(siteNameQ)
if err != nil {
    log.Fatalln(err)
}

defer resp.Body.Close()


body, err := ioutil.ReadAll(resp.Body)
if err != nil {
log.Fatalln(err)
}

log.Println(string(body))

}



func main() {

    pathToFile := os.Args[1]
    numberOfThreads := os.Args[2]


    fmt.Println(pathToFile)
    fmt.Println(numberOfThreads)



   file, err := os.Open(pathToFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        siteNameQ := scanner.Text()
        fmt.Println(scanner.Text())
        getreq(siteNameQ)

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

}