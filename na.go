package main

import( 
    "os"
    "bufio"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "time"
)

func getreq(siteNameQ string) {

resp, err := http.Get(siteNameQ)
    if err != nil {
        log.Fatal(err)
    }

    // Print the HTTP Status Code and Status Name
    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

}



func main() {

    pathToFile := os.Args[1]
    numberOfThreads := os.Args[2]


    fmt.Println(pathToFile)
    fmt.Println(numberOfThreads)
    numberInt,_ := strconv.Atoi(numberOfThreads)
    fmt.Println(numberInt)


   file, err := os.Open(pathToFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
       siteNameQ := scanner.Text()
       fmt.Println(scanner.Text())
       go getreq(siteNameQ)
       time.Sleep(200 * time.Millisecond)
}
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}