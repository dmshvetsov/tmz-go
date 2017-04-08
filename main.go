package main

import (
  "fmt"
  "time"
  "os"
)

var Tformat = time.Kitchen
var SecInHour = 3600

func toutput(time time.Time) {
  fmt.Printf("Your time is: %v\n", time.Format(Tformat))
}

func doutput(ltime, rtime time.Time) {
  _, loffset := ltime.Zone()
  _, roffset := rtime.Zone()
  if loffset < roffset {
    fmt.Printf("Difference is: %v hour(s)\n", (loffset - roffset) / SecInHour)
  } else {
    fmt.Printf("Difference is: %v hour(s)\n", (roffset - loffset) / SecInHour)
  }
}

func main() {
  utime := time.Now()

  userInput := os.Args[1]
  othertimezone, err := time.LoadLocation(userInput)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  toutput(utime)
  toutput(utime.In(othertimezone))
  doutput(utime, utime.In(othertimezone))
}
