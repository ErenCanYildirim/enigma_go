# Enigma Go

<img src="https://blog.golang.org/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="100" alt="Go Logo">

A full **Go implementation of the WWII German Enigma I machine**, including rotors, reflectors, ring settings, and plugboard configuration. This simulator follows the Enigma I configuration used by the German Wehrmacht.

## Features

- Historical rotor and reflector configurations
- Custom rotor and reflector support
- Configurable ring settings
- Plugboard (Steckerbrett) support
- Double-stepping rotor mechanism
- Encrypts and decrypts messages (reciprocal encryption)
- Preserves space and ignores non-alphabetic characters

## Installation

```bash
git clone https://github.com/ErenCanYildirim/enigma_go.git
cd enigma
go test ./...
```

## Usage

```go 

package <your_package>

import (
    "fmt"
    "https://github.com/ErenCanYildirim/enigma_go"
)

func main() {
    //Build the enigma machine
    enigmaMachine, err := enigma.NewBuilder().
        WithRotors("I","II","III").
        WithReflector("UKW-B").
        WithPlugboard("AB CD EF").
        WithRotorPositionsFromString("AAA").
        WithRingSettingsFromString("AAA").
        Build()

    if err != nil {
        panic(err)
    }

    message := "HELLO WORLD"
    encrypted, _ := enigmaMachine.Encrypt(message)
    fmt.Println("Encrypted", encrypted)

    // Reset rotor position before decryption
    enigmaMachine.SetRotorPositions(0, 0, 0)
    decrypted, _ := enigmaMachine.Decrypt(encrypted)
    fmt.Println("Decrypted:", decrypted)
}
```

## Contributing

Contributions are welcome. Please create your features and add them for pull requests!

## Acknowledgements

Inspired by the historical Enigma machine in WWII.