package main

import( 
    "os"
    "bufio"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "time"
    "net"
)

func getreq(siteNameQ string) {

siteNameFull := "http://" + siteNameQ

client := &http.Client{ Timeout: 5 * time.Second,
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
  } }

resp, err := client.Get(siteNameFull)
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode == 301 {
	resp, err := http.Get(siteNameFull)
	    if err != nil {
    	    log.Fatal(err)
	    }  
	    fmt.Println(siteNameFull)
	    fmt.Println("301 ", resp.StatusCode)

}


    if resp.StatusCode == 302 {
	resp, err := http.Get(siteNameFull)
	    if err != nil {
    	    log.Fatal(err)
	    }  
	    fmt.Println(siteNameFull)
	    fmt.Println("302 ", resp.StatusCode)

}


if resp.StatusCode != 301 && resp.StatusCode != 302  {

    // Print the HTTP Status Code and Status Name
    fmt.Println(siteNameFull)
    fmt.Println("", resp.StatusCode)
//, http.StatusText(resp.StatusCode))
}
}


func askDns(siteNameQ string) {


ips, err := net.LookupIP(siteNameQ)
    if err != nil {
	fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
	os.Exit(1)
    }
    for _, ip := range ips {
    fmt.Printf(" IN A %s\n", ip.String())
	
    }

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
//       fmt.Println(scanner.Text())
       askDns(siteNameQ)
       go getreq(siteNameQ)
       time.Sleep(300 * time.Millisecond)
}
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}