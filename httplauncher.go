package main


import (
	"fmt"
  "flag"
  "os"
  "net"
  "io/ioutil"
  "strings"
  "net/http"
  "bufio"
)

func main() {
  verbool, urlstring, wordlstring, timefloat, concurrint, refstring, datastring := parseargs()
  fmt.Println(verbool, urlstring, wordlstring, timefloat, concurrint, refstring, datastring)
  res :=  request(urlstring, wordlstring, datastring)
  fmt.Println(res)
}

func validate(urlstring string, datastring string, wordlstring string) {
  _, URL := net.Dial("tcp", fmt.Sprintf("%v:80", urlstring))
  _, Datapath := os.ReadFile(datastring)
  _, Wordlistpath := os.ReadFile(wordlstring)
  if URL != nil || Datapath != nil || Wordlistpath!= nil {
    fmt.Println("Error: Check ", Wordlistpath)
    os.Exit(1)
  }
}

func parseargs() (bool, string, string, float64, int, string, string) {
  verbflag := flag.Bool("v", false, "Displays verbose output")
  urlflag := flag.String("u", "", "URL to target")
  wordlflag := flag.String("w", "", "Wordlist locations seperated by colons")
  timeflag := flag.Float64("t", 0, "Time in between packets sent")
  concurrflag := flag.Int("c", 16, "Number of oncurrently running workers")
  refflag := flag.String("r", "", "String to check the absence of every response")
  dataflag := flag.String("d", "", "Location of request data to be sent")
  
  flag.Parse()

  if (*wordlflag == "" || *refflag == "" || *dataflag == "") {
    fmt.Println(*wordlflag, *refflag, *dataflag)
    fmt.Println("Not enough info!")
    os.Exit(1)
  }

  return *verbflag, *urlflag, *wordlflag, *timeflag, *concurrflag, *refflag, *dataflag
}

func getdatastring(datastring string) string {
  b, err := ioutil.ReadFile(datastring)
    if err != nil {
        fmt.Print(err)
    }
    return string(b)
}

func request(urlstring string, wordlstring string, datastring string) string {
  data := getdatastring(datastring)
  readFile, err := os.Open(wordlstring)
  if err != nil {
          os.Exit(1)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  for fileScanner.Scan() {
    newdata := strings.Replace(data, "^P1^", fileScanner.Text(), -1)
    _, err := http.NewRequest("POST", urlstring, strings.NewReader(newdata))
    if err != nil {
      os.Exit(1)
    }
  }

  readFile.Close()
  return "tmp"
}
