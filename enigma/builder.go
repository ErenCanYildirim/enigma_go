package enigma

import "fmt"

//provides interface for constructing an Enigma machine
type Builder struct {
	rotors         []*Rotor
	reflector      *Reflector
	plugboard      *Plugboard
	rotorPositions []int
	ringSettings   []int
	err            error
}

func NewBuilder() *Builder {
	return &Builder{}
}

// set the rotors using historical rotor types, rotorType should be named like 'I', 'II' etc.
func (b *Builder) WithRotors(rotorTypes ...string) *Builder {
	if b.err != nil {
		return b
	}

	b.rotors = make([]*Rotor, 0, len(rotorTypes))
	for _, rotorType := range rotorTypes {
		rotor, err := NewHistoricalRotor(rotorType)
		if err != nil {
			b.err = err
			return b
		}
		b.rotors = append(b.rotors, rotor)
	}
	return b
}

// set custom rotors
func (b *Builder) WithCustomRotors(rotors ...*Rotor) *Builder {
	if b.err != nil {
		return b
	}

	b.rotors = rotors
	return b
}

// sets the reflector using a historical reflector type
// reflectorType should be like "UKW-B"
func (b *Builder) WithReflector(reflectorType string) *Builder {
	if b.err != nil {
		return b
	}

	reflector, err := NewHistoricalReflector(reflectorType)
	if err != nil {
		b.err = err
		return b
	}

	b.reflector = reflector
	return b
}

// sets a custom reflector
func (b *Builder) WithCustomReflector(reflector *Reflector) *Builder {
	if b.err != nil {
		return b
	}

	b.reflector = reflector
	return b
}

// sets the plugboard connections, conns are strings like "AB CD EF GH"
func (b *Builder) WithPlugboard(connections string) *Builder {
	if b.err != nil {
		return b
	}

	plugboard, err := NewPlugboard(connections)
	if err != nil {
		b.err = err
		return b
	}

	b.plugboard = plugboard
	return b
}

// sets the initial rotor positions, specified as integers (0-25) or converted from letters
func (b *Builder) WithRotorPositions(positions ...int) *Builder {
	if b.err != nil {
		return b
	}

	b.rotorPositions = positions
	return b
}

//sets the initial rotor positions from a string
// positionStr should be like "AAA" or "XYZ" where each letter represents a rotor position
func (b *Builder) WithRotorPositionsFromString(positionStr string) *Builder {
	if b.err != nil {
		return b
	}

	positions := make([]int, len(positionStr))
	for i, char := range positionStr {
		if char < 'A' || char > 'Z' {
			if char >= 'a' && char <= 'z' {
				char = char - 'a' + 'A'
			} else {
				b.err = fmt.Errorf("invalid position character: %c", char)
				return b
			}
		}
		positions[i] = int(char - 'A')
	}

	b.rotorPositions = positions
	return b
}

//sets the ring settings for the rotors
func (b *Builder) WithRingSettings(settings ...int) *Builder {
	if b.err != nil {
		return b
	}

	b.ringSettings = settings
	return b
}

//sets the ring settings from a string
// settingsStr should be like "AAA" or "XYZ" where each letter represents a ring setting
func (b *Builder) WithRingSettingsFromString(settingsStr string) *Builder {
	if b.err != nil {
		return b
	}

	settings := make([]int, len(settingsStr))
	for i, char := range settingsStr {
		if char < 'A' || char > 'Z' {
			if char >= 'a' && char <= 'z' {
				char = char - 'a' + 'A'
			} else {
				b.err = fmt.Errorf("invalid ring setting character: %c", char)
				return b
			}
		}
		settings[i] = int(char - 'A')
	}

	b.ringSettings = settings
	return b
}

// construct the Enigma machine with specified configuration
func (b *Builder) Build() (*Enigma, error) {
	if b.err != nil {
		return nil, b.err
	}

	if len(b.rotors) == 0 {
		return nil, fmt.Errorf("at least one rotor must be specified")
	}

	if b.reflector == nil {
		return nil, fmt.Errorf("reflector must be specified")
	}

	if b.plugboard == nil {
		//create empty plugboard
		b.plugboard, _ = NewPlugboard("")
	}

	//apply the ring settings
	if len(b.ringSettings) > 0 {
		if len(b.ringSettings) != len(b.rotors) {
			return nil, fmt.Errorf("number of ring settings (%d) must match number of rotors (%d)",
				len(b.ringSettings), len(b.rotors))
		}
		for i, setting := range b.ringSettings {
			b.rotors[i].SetRingSetting(setting)
		}
	}

	enigma := NewEnigma(b.rotors, b.reflector, b.plugboard)

	// apply the rotor positions
	if len(b.rotorPositions) > 0 {
		if err := enigma.SetRotorPositions(b.rotorPositions...); err != nil {
			return nil, err
		}
	}
	return enigma, nil
}
