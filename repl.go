package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func startRepl() {

  for {
   scanner := bufio.NewScanner(os.Stdin)
   fmt.Print("Pokedex >")

   scanner.Scan()
   text := scanner.Text()

   cleaned := cleanInput(text)

   fmt.Print("echoing: ", cleaned)
  }


}

func cleanInput(str string) []string {
  lowered := strings.ToLower(str)
  words := strings.Fields(lowered)
  return words
}



