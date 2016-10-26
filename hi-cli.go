package main

import (
    "bufio"
    "fmt"
    "os"
    "flag"
)

func main() {
    wordFlag := flag.String("wordly", "the-word", "A string value.")
    numberFlag := flag.Int("numeraly", 42012, "An integer value.")
    booleanFlag := flag.Bool("booly", true, "A boolean value.")
    nameFlag := flag.String("yourname", "Person With No Name", "Put your name here.")

    flag.Parse()

    var reader = bufio.NewReader(os.Stdin)
    fmt.Print("Hello, " + *nameFlag + ", please enter some text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)

    fmt.Println("word: ", *wordFlag)
    fmt.Println("numb: ", *numberFlag)
    fmt.Println("fork: ", *booleanFlag)
}