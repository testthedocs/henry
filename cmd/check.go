package main

import (
   "fmt"
   "os"
)

func main() {

	// check if docker is installed
	// print message telling this
	fmt.Println("Checking for dependecies")

   // get FileInfo structure describing file
   _, err := os.Stat("/usr/bin/docker")

   // check if there is an error
   if err != nil {

      // check if error is file does not exist
      if os.IsNotExist(err) {
         fmt.Println("Can't find docker")
      }

   } else {
    fmt.Println("Yeah, looks good !")
   }
}
