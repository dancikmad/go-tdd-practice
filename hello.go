package main 

import "fmt"

const spanish = "Spanish"
const french = "French"
const russian = "Russian"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const russianHelloPrefix = "Привет, "


func Hello(name string, language string) string {
  if name == "" {
      name = "World"
  }
  if language == spanish {
      return spanishHelloPrefix + name 
  }
  if language == french {
      return frenchHelloPrefix + name 
  }
  if language == russian {
      return russianHelloPrefix + name 
  }
  return englishHelloPrefix + name
}

func main() {
  fmt.Println(Hello("world", ""))
}
