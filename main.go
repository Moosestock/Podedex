package main

import ("fmt"
    "strings"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
         
        scanner.Scan()
        text := scanner.Text()
        if text != "" {
            words := cleanInput(text)
            fmt.Printf("Your command was: %s\n", words[0])
        }        
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error: %w", err)
    }
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}
