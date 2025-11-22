# How the Enigma implementation works

Checkout out: [Enigma](https://www.cryptomuseum.com/crypto/enigma/i/index.htm) for more

## Overview

The Enigma machine consists of three main components:

1. **Rotors** - Cipher wheels that scramble letters
2. **Reflector** - Bounces signals back through the rotors
3. **Plugboard** - Swaps pairs of letters before and after rotor processing 

Encryption and decryption are reciprocal, meaning the same process is used for both.

## 1. Rotors

A rotor is represented by the `Rotor` struct:

```go
type Rotor struct {
    wiring      [26]int // forward mapping
    wiringRev   [26]int // reverse mapping
    position    int     // current rotor position
    ringSetting int     // ring setting offset
    notches     []int   // positions that trigger rotor stepping
    Name        string
}
```

Features:
    - Forward mapping: Encodes a letter going from right to left
    - Reverse mapping: Encodes a letter coming back from the reflector
    - Ring settings: Adjusts the wiring offset (historical Enigma feature)
    - Rotor stepping: Handles rotation, double-stepping is implemented
    - Turnover notches: Certain positions trigger the next rotor to step

Methods:
    Forward(input int) int -> passes a letter forward through the rotor
    Backward(input int) int – Passes a letter backward through the rotor.
    Step() – Advances the rotor by one position.
    AtNotch() – Checks if the rotor is at a turnover position.


## 2. Reflector

The reflector is represented via :

```go
type Reflector struct {
    wiring [26]int 
    name string
}
```

Features:
    - reflects a signal back through the rotors
    - ensures that encryption is reciprocal
    - cannot be configured dynamically in historical mode

## 3. Plugboard:

type Plugboard struct {
    wiring [26]int 
}

Features:
    - maps pairs of letters before and after rotor encryption
    - connections are defined using strings like "AB CD EF"
    - letters not in pairs remain unchanged

## Enigma machine

The Enigma struct ties the parts together:
```go 
    type Enigma struct {
        rotors []*Rotor
        reflector *Reflector
        plugboard *Plugboard
    }
```

Features:

    - Handles rotor stepping, including double-stepping.
    - Encrypts single characters and full messages.
    - Preserves spaces and ignores non-alphabetic characters.
    - Supports historical rotor configurations or custom rotors.
    
Stepping logic:
    1. rightmost rotor always steps on every key press
    2. middle rotor steps if the right rotor is at its notch or if it's at a notch itself (double-stepping)
    3. Left rotor only steps when the middle rotor is at its notch 

## Builder pattern

To create the machine yourself, the **Builder** pattern is provided:

```go 
NewBuilder().
    WithRotors("I", "II", "III").
    WithReflector("UKW-B").
    WithPlugboard("AB CD EF").
    WithRotorPositionsFromString("AAA").
    WithRingSettingsFromString("AAA").
    Build()
```

Features:
    - choose historical rotors and reflectors
    - set custom rotors and reflectors
    - define rotor starting positions
    - set ring settings
    - configure plugboard connections
    - automatically validates configuration

## Encryption flow

1. step rotors according to the stepping mechanism
2. convert letter to uppercase index (A=0,B=1,...,Z=25)
3. Pass through plugboard
4. Pass through rotors (right to left)
5. Pass through reflector
6. Pass back through rotors (left to right)
7. Pass through pluboard again
8. Convert index back to letter

The decryption uses the same process, due to the reciprocal wiring.

## Testing

Tests are located in **enigma_test.go**:

**TestEnigma_EncryptDecrypt** ensures that decryption recovers the original message.
**TestEnigma_LowercaseAndNonAlpha** confirms handling of lowercase letters and ignoring non-alphabetic characters.
**TestEnigma_EncryptConsistency** checks that the same rotor positions produce the same output.

## Historical Configurations
    Rotors: I–VIII, each with unique wiring and notches.

    Reflectors: UKW-A, UKW-B, UKW-C.

    Ring settings: Adjustable per rotor.

    Plugboard: Optional letter swaps.

## Notes

- non-alphabetic characters are skipped during encryption
- spaces are preserved
- ring settings allow simulation of real Enigma machines
- this implementation is suitable for educational or simulation purposes, not secure encryption, as Enigma is not secure today anymore.