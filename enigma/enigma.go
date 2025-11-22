package enigma

/*
	Enigma provides a complete implementation of the Enigma machine used by the German Wehrmacht in WWII
	This implementation is based on the Enigma I config with configurable rotors, reflectors, ring settings
	and plugboard connections.
*/

import (
	"fmt"
	"strings"
)

const (
	AlphabetSize = 26
)

//-------------------- Rotor -----------------------------

// Rotor represents a single rotor with its wiring, position, ring setting and turnover notches
type Rotor struct {
	wiring      [AlphabetSize]int //mapping of the forward wiring
	wiringRev   [AlphabetSize]int //mapping of the reverse wiring
	position    int               // current rotor pos (0-25)
	ringSetting int               // offset of the ring setting (0-25)
	notches     []int             // turnover notch position
	Name        string            // rotor identifier
}

func NewRotor(name string, wiring string, notches string) (*Rotor, error) {
	if len(wiring) != AlphabetSize {
		return nil, fmt.Errorf("invalid wiring length: expected %d, got %d", AlphabetSize, len(wiring))
	}

	r := &Rotor{
		Name:        name,
		position:    0,
		ringSetting: 0,
	}

	// building the forward and reverse wiring
	for i, char := range strings.ToUpper(wiring) {
		if char < 'A' || char > 'Z' {
			return nil, fmt.Errorf("invalid character in wiring: %c", char)
		}
		output := int(char - 'A')
		r.wiring[i] = output
		r.wiringRev[output] = i
	}

	// parsing the notches
	for _, char := range strings.ToUpper(notches) {
		if char < 'A' || char > 'Z' {
			return nil, fmt.Errorf("invalid notch character: %c", char)
		}
		r.notches = append(r.notches, int(char-'A'))
	}

	return r, nil
}

// Set current rotor position (A=0, B=1, ... , Z = 25)
func (r *Rotor) SetPosition(pos int) {
	r.position = pos % AlphabetSize
}

// set the ring settings for this rotor (A=0, B=1, ..., Z=25)
func (r *Rotor) SetRingSetting(setting int) {
	r.ringSetting = setting % AlphabetSize
}

func (r *Rotor) Position() int {
	return r.position
}

// return true if the rotor is at a turnover notch position
func (r *Rotor) AtNotch() bool {
	for _, notch := range r.notches {
		if r.position == notch {
			return true
		}
	}
	return false
}

// advances the rotor by one position
func (r *Rotor) Step() {
	r.position = (r.position + 1) % AlphabetSize
}

// pass a signal through the rotor in the forward direction
func (r *Rotor) Forward(input int) int {
	//adjust for rotor pos and ring setting
	shift := r.position - r.ringSetting
	input = (input + shift + AlphabetSize) % AlphabetSize

	// pass through the wiring
	output := r.wiring[input]

	// adjust back
	output = (output - shift + AlphabetSize) % AlphabetSize
	return output
}

// passes a signal through the rotor in the reverse direction
func (r *Rotor) Backward(input int) int {
	//adjusting for rotor pos and ring setting
	shift := r.position - r.ringSetting
	input = (input + shift + AlphabetSize) % AlphabetSize

	// pass through reverse wiring
	output := r.wiringRev[input]

	// adjust back
	output = (output - shift + AlphabetSize) % AlphabetSize
	return output
}

// -------------- Reflector --------------------

// represents the reflector component that bounces back through the rotors
type Reflector struct {
	wiring [AlphabetSize]int
	name   string
}

// creates a new reflector with the specified wiring
func NewReflector(name string, wiring string) (*Reflector, error) {
	if len(wiring) != AlphabetSize {
		return nil, fmt.Errorf("invalid wiring length: expected %d, got %d", AlphabetSize, len(wiring))
	}

	ref := &Reflector{name: name}

	for i, char := range strings.ToUpper(wiring) {
		if char < 'A' || char > 'Z' {
			return nil, fmt.Errorf("invalid character in wiring: %c", char)
		}
		ref.wiring[i] = int(char - 'A')
	}

	return ref, nil
}

// passes a signal through the reflector
func (ref *Reflector) Reflect(input int) int {
	return ref.wiring[input]
}

//----------------------- Plugboard ----------------------------------------

// Plugboard represents the plugboard (Steckerbrett) that swaps letter pairs
type Plugboard struct {
	wiring [AlphabetSize]int
}

// create a plugboard with the specified connections
// connections should be a string like "AB CD EF" where each pair is swapped
func NewPlugboard(connections string) (*Plugboard, error) {
	pb := &Plugboard{}

	//init identity mapping
	for i := 0; i < AlphabetSize; i++ {
		pb.wiring[i] = i
	}

	if connections == "" {
		return pb, nil
	}

	//parse connections
	pairs := strings.Fields(strings.ToUpper(connections))
	used := make(map[rune]bool)

	for _, pair := range pairs {
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid plugboard pair: %s", pair)
		}

		a, b := rune(pair[0]), rune(pair[1])

		if a < 'A' || a > 'Z' || b < 'A' || b > 'Z' {
			return nil, fmt.Errorf("invalid characters in plugboard pair: %s", pair)
		}

		if used[a] || used[b] {
			return nil, fmt.Errorf("letter used multiple times in plugboard: %s", pair)
		}

		aIdx, bIdx := int(a-'A'), int(b-'A')
		pb.wiring[aIdx] = bIdx
		pb.wiring[bIdx] = aIdx

		used[a] = true
		used[b] = true
	}

	return pb, nil
}

// passes signal forward through the plugboard
func (pb *Plugboard) Forward(input int) int {
	return pb.wiring[input]
}

//------------------- ENIGMA ----------------------------

// struct for the entire Enigma machine
type Enigma struct {
	rotors    []*Rotor
	reflector *Reflector
	plugboard *Plugboard
}

func NewEnigma(rotors []*Rotor, reflector *Reflector, plugboard *Plugboard) *Enigma {
	if plugboard == nil {
		plugboard, _ = NewPlugboard("")
	}
	return &Enigma{
		rotors:    rotors,
		reflector: reflector,
		plugboard: plugboard,
	}
}

// implement stepping mechanism with double-stepping
/*
	Enigma logic:
		rightmost rotor always steps on every press
		middle rotor stops when either:
			- right rotor is at notch before stepping -> middle steps
			- middle rotor itself is at notch -> left and middle step
		left rotor stops only, when the middle rotor is at its notch
*/
func (e *Enigma) stepRotors() {
	right := e.rotors[0]
	middle := e.rotors[1]
	left := e.rotors[2]

	if middle.AtNotch() {
		middle.Step()
		left.Step()
	} else {
		middle.Step()
	}
	right.Step()
}

// encrypts a signle char
func (e *Enigma) EncryptChar(char rune) (rune, error) {
	char = rune(strings.ToUpper(string(char))[0])

	if char < 'A' || char > 'Z' {
		return char, fmt.Errorf("invalid character: %c (only A-Z supported)", char)
	}

	// step rotors before encryption
	e.stepRotors()

	//convert to index
	signal := int(char - 'A')

	//through plugboard
	signal = e.plugboard.Forward(signal)

	//through rotors (right to left)
	for i := 0; i < len(e.rotors); i++ {
		signal = e.rotors[i].Forward(signal)
	}

	// through reflector
	signal = e.reflector.Reflect(signal)

	// back through the rotors (left to right)
	for i := len(e.rotors) - 1; i >= 0; i-- {
		signal = e.rotors[i].Backward(signal)
	}

	// through plugboard again
	signal = e.plugboard.Forward(signal)

	return rune(signal + 'A'), nil
}

// encrypts a message, preserves spaces, ignores non-alphabetic chars
func (e *Enigma) Encrypt(message string) (string, error) {
	var result strings.Builder
	result.Grow(len(message))

	for _, char := range message {
		// preserve spaces
		if char == ' ' {
			result.WriteRune(' ')
			continue
		}

		// convert lowercase to uppercase
		if char >= 'a' && char <= 'z' {
			char = char - 'a' + 'A'
		}

		// skip non-alphabetic characters
		if char < 'A' || char > 'Z' {
			continue
		}

		encrypted, err := e.EncryptChar(char)
		if err != nil {
			return "", err
		}
		result.WriteRune(encrypted)
	}
	return result.String(), nil
}

// decrypt is identical to Encrypt due to the reciprocal nature of the Enigma
func (e *Enigma) Decrypt(message string) (string, error) {
	return e.Encrypt(message)
}

// sets the starting positions of all rotors
func (e *Enigma) SetRotorPositions(positions ...int) error {
	if len(positions) != len(e.rotors) {
		return fmt.Errorf("expected %d positions, got %d", len(e.rotors), len(positions))
	}

	for i, pos := range positions {
		e.rotors[i].SetPosition(pos)
	}

	return nil
}

// returns the current positions of all rotors
func (e *Enigma) GetRotorPositions() []int {
	positions := make([]int, len(e.rotors))
	for i, rotor := range e.rotors {
		positions[i] = rotor.Position()
	}
	return positions
}
