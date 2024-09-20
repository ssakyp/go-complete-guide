package main

import "bufio"

var reader bufio.NewReader(os.Stdin)
func main() {
  
}

func getUserChoice() (string, error) {
  fmt.Println("Please make your choice")
  fmt.Println("1) Add up all the numbers of to number X")
  fmt.Println("2) Build the factorial up to number X")
  fmt.Println("3) Sum up manually entered numbers")
  fmt.Println("4) Sum up a list of entered numbers")

  userInput, err := reader.ReadString("\n")

  if err != nil {
    return "", err
  }

  userInput = strings.Replace(userInput, "\n", "", -1)
}
