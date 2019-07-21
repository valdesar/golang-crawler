package main

import (
      //  "reflect"
        "fmt"
        "net/http"
        "os"
        "time"
        "bufio"
        "log"
        "net"
 		"strings"
    //    "sync/atomic"
     //   "math/rand"
        )





func wordCount(str string) map[string]int {
    wordList := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range wordList {
        _, ok := counts[word]
        if ok {
            counts[word] += 1
        } else {
            counts[word] = 1
        }
    }
    return counts


}
 



func getreq(siteNameQ string) {


siteNameFull := "http://" + siteNameQ

tr := &http.Transport{ MaxConnsPerHost: 1, }

client := &http.Client{ Timeout: 5 * time.Second,
                        Transport: tr,
                        CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
  } }

//check error handling!!! may be use this 

//&httpError{
//  err:     err.Error() + " (Client.Timeout exceeded while awaiting headers)",
 // timeout: true,
//}

//https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

resp, err := client.Get(siteNameFull)
    if err != nil {
        log.Fatal(err)
    }
    if resp.StatusCode == 301 {
	resp, err := http.Get(siteNameFull)
	    if err != nil {
    	    log.Fatal(err)
	    }  
	    fmt.Println(siteNameFull, "301 ", resp.StatusCode, "\n")
//	    fmt.Println("301 ", resp.StatusCode)
}
    if resp.StatusCode == 302 {
	resp, err := http.Get(siteNameFull)
	    if err != nil {
    	    log.Fatal(err)
	    }  
	    fmt.Println(siteNameFull, "302 ", resp.StatusCode)
//	    fmt.Println("302 ", resp.StatusCode)
}
if resp.StatusCode != 301 && resp.StatusCode != 302  {
    fmt.Println(siteNameFull, "", resp.StatusCode)
//    fmt.Println("", resp.StatusCode)
}
}


func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}



func main() {
//taking command line arguments
file, err := os.Open("azaaza.txt")
 
    if err != nil {
  log.Fatalf("failed opening file: %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var txtlines []string
    for scanner.Scan() {
  txtlines = append(txtlines, scanner.Text())
    }
    file.Close()

//goroutines code
    oreChannel := make(chan string, 2)
  //  minedOreChan := make(chan string)
   // minedOreChan2 := make(chan string)
    minedOreChan3 := make(chan map[string]string, 2)


// Разведчик
go func(mine []string) { 	
 for _, item := range mine {
  if item != "" {
   oreChannel <- item //передаем данные в oreChannel
  }
 }
}(txtlines)


// Добытчик







go func() {
 for i := 0; i < 400; i++ {
  foundOre := <-oreChannel //чтение из канала oreChannel
 // fmt.Println("\n")
//  fmt.Println("Processing: ", foundOre)
  var siteNameQ = foundOre


  ips, err := net.LookupIP(siteNameQ)
    if err != nil {
	fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
//	os.Exit(1)
    }
    if len(ips) > 1 {
        fmt.Printf("Got more then 1 ip address for this website, skipping\n")
    }     else {
    	 for _, ip := range ips {
       //         fmt.Printf("%s\n", ip)   
                var ipAddressString string = ip.String()
                sitenameIpPair := make(map[string]string, 100)
                 sitenameIpPair[siteNameQ]=ipAddressString
   				 for index,element := range sitenameIpPair{
	    //         fmt.Println(index,"=>",element)
	      //       minedOreChan <-  element
	        //     minedOreChan2 <- index
                 blank1 := index
                 _= blank1
                 blank2 := element
                 _ = blank2
	             minedOreChan3 <- sitenameIpPair
       }
               }
                        }
 // go getreq(siteNameQ)
 //minedOreChan <- foundOre //передаем данные в minedOreChan
 }
}()
// Переработчик


//var ops uint64
var beenRead[]string

//var visitedSites[]string

go func() {

 var arr[]string
// var arr2[]string


ticker := time.NewTicker(10 * time.Millisecond)
    go func() {
        for t := range ticker.C {
            //fmt.Println("Tick at", t)
            var emptyvar = t
            _ = emptyvar
            arr = arr[:0]
        }
    }()




 semaphoreChan := make(map[string](chan int))
 





 for i := 0; i < 400; i++ {
 // minedOre := <-minedOreChan //чтение данных из minedOreChan
  //minedOre2 := <-minedOreChan2


  minedOre3 := <- minedOreChan3

  //fmt.Println(minedOre3)
  for inmap,inelem := range minedOre3{
     //   fmt.Println(inmap,"=>",inelem)
        minedOre2 := inmap
        minedOre := inelem


semaphoreChan[minedOre] = make(chan int,50)



 // fmt.Println("From Miner: ", minedOre)
 var siteNameQ string = minedOre2
 // fmt.Println(siteNameQ)
//  var pairSiteIp string = minedOre2 + " " + minedOre
 // arr2 = append(arr2, pairSiteIp )
 // arr2 = append(arr2, minedOre2 )
  
  arr = append(arr, minedOre )
 // fmt.Println("Adding IP: ", minedOre)
 // fmt.Println("IP: ", arr)



/*
concurrencyLimit := 2
semaphoreChan := make(map[string](chan int))
//semaphoreChan := make(chan int, concurrencyLimit)
//m := make(map[string]int)


semaphoreChan[minedOre] = make(chan int, 2)

semaphoreChan[minedOre] <- 2
    mything := <-semaphoreChan[minedOre]

      //	semaphoreChan <- 1
*/
   



   

counter := make( map[string]int, 100 )    
//counter := make( map[map[string]string]int, 100 )    
for _, row := range arr {
     counter[row]++
 //     fmt.Println(counter)
} 







//ConnsToServer := make(chan int, 2)






//for ind,elem := range counter{
	//	if elem > 2 {
		//fmt.Println("--PASS")
//		atomic.AddUint64(&ops, 1)
    //   fmt.Println(ind,"=--->",elem)
   //     empty1 := index
 //       empty2 := element
     //   _ = empty1
   //     _ = empty2
               
//} else {
//	_ = ind



/*

var m map[string]int
m = make(map[string]int)
m["oneone"] = 1
m["twotwo"] = 2
for k, v := range m {
    fmt.Printf("key[%s] value[%s]\n", k, v)
}

*/


/*
things := make(map[string](chan int))


things[minedOre] = make(chan int, 2)
    things[minedOre] <- 2
    things[minedOre] <- 2
    things[minedOre] <- 2
    mything := <-things[minedOre]
    fmt.Printf("my thing: %d \n", mything)
    fmt.Println(things)

*/





//ConnsToServer <- 2
//ConnsToServer <- 2
//ConnsToServer <- 2
//fmt.Println(<-ConnsToServer)
//CurrentConnsToServer := <-ConnsToServer

//fmt.Println(<-ConnsToServer)
//fmt.Println(<-ConnsToServer)
//if CurrentConnsToServer <= 2 {



//fmt.Println(elem)
//		atomic.AddUint64(&ops, 1)
       // fmt.Println(ind,"=--+->",elem)





   

    
 
 


 s1 := siteNameQ
//   fmt.Println(contains(beenRead, s1))
  if contains(beenRead, s1) == false {






    semaphoreChan[minedOre] <- i
 //   mything := <-semaphoreChan[minedOre]
   // fmt.Printf("%s %d \n", minedOre, mything)
   fmt.Println(minedOre)
    perIpChannel := make(chan int, 4)
    
   for elemIP := range semaphoreChan {
  //  fmt.Printf("%s  \n", elemIP)
  _ = elemIP
     perIpChannel <- 1

//if elemIP == minedOre {


//	fmt.Println("rrr")

}


   //     fmt.Println(elemIP)
//    if contains(elemIP, minedOre) == false {
    
    //		getreq(s1)
  //  }

//for krr, vrr := range semaphoreChan { 
  //  fmt.Printf("key[%s] value[%d]\n", krr, vrr)
//}
 //   }




  	 go getreq(s1)
for len(perIpChannel) > 0 {
  <-perIpChannel
}
//  	<-semaphoreChan
  


  }



   beenRead = append(beenRead, siteNameQ)
//}

/*
  beenRead <- siteNameQ  
  readBeenRead := <- beenRead
fmt.Println(readBeenRead, "fuxk")
*/



  // for _,tttt := range beenRead{
		//fmt.Println("--PASS")

	//	if tttt == siteNameQ {
      //  fmt.Println(tttt, "already visited")
		// } else { fmt.Println(tttt, "first visit") }


	//	}







/*

beenReadst := strings.Join(beenRead, " ")

    for index,element := range wordCount(beenReadst){
        fmt.Println(index,"=>",element)
        if element == 1 {

         //    for _,ttrt := range visitedSites{
        
     //        if ttrt == siteNameQ {

       //      	fmt.Println(ttrt, "visited")
         //    } 
        	getreq(index)


//}
        }
        
    }


*/
 





  
 // visitedSites = append(visitedSites, siteNameQ)
   //fmt.Println(beenRead)
 




//  i := 0
 //beenRead[i] = beenRead[len(beenRead)-1] // Copy last element to index i.
 //beenRead[len(beenRead)-1] = ""   // Erase last element (write zero value).
 //beenRead = beenRead[:len(beenRead)-1]   // Truncate slice.


  // for tttt := range beenRead{
		//fmt.Println("--PASS")
    //    fmt.Println(tttt)
	//	 }

  
   //	getreq(siteNameQ)






/*
for _, rrrrr := range arr2 {
     counter[rrrrr]++
      fmt.Println(counter)
} 
*/



	
//}
       // var siteNameQ = foundOre 



  //      opsFinal := atomic.LoadUint64(&ops)
     //   fmt.Println("ops:", opsFinal)

   // if ops > 1 {
   // break
//}// else {

 
        
   // }
  //  }
    }
 }
}() 
<-time.After(time.Second * 1) // Все еще можете игнорировать

}
