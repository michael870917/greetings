# Grettings GO
Este paquete proporrciona una forma simple de saludos

## Instalacion
Ejecuta el siguiente comando para installar el paquete

go get -u github.com/michael870917/greetings

## Uso
```go
package main

import (
	"fmt"
	otherWord "github.com/michael870917/greetings"
	"log"
)

func main() {
	fmt.Println("=======")
	log.SetPrefix("grettings: ")
	log.SetFlags(0)

	names := []string{
		"Alex",
		"Juan",
	}

	messages, err := otherWord.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println("=======")
}

```
