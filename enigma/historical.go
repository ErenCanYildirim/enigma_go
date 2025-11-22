package enigma

//Historical rotor wirings and config that was used by the Enigma of the German Wehrmacht

var (
	//the historical wiring config for standard rotors
	RotorWirings = map[string]string{
		"I":    "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"II":   "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"III":  "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"IV":   "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		"V":    "VZBRGITYUPSDNHLXAWMJQOFECK",
		"VI":   "JPGVOUMFYQBENHZRDKASXLICTW",
		"VII":  "NZJHGRCXMYSWBOUFAIVLPEKQDT",
		"VIII": "FKQHTLXOCBJSPDZRAMEWNIUYGV",
	}

	//RotorNotches -> turnover notch positions for each rotor
	RotorNotches = map[string]string{
		"I":    "Q",  // Turnover from Q to R
		"II":   "E",  // Turnover from E to F
		"III":  "V",  // Turnover from V to W
		"IV":   "J",  // Turnover from J to K
		"V":    "Z",  // Turnover from Z to A
		"VI":   "ZM", // Two turnover positions
		"VII":  "ZM", // Two turnover positions
		"VIII": "ZM", // Two turnover positions
	}

	// ReflectorWirings contains the historical reflector configurations
	ReflectorWirings = map[string]string{
		"UKW-A": "EJMZALYXVBWFCRQUONTSPIKHGD",
		"UKW-B": "YRUHQSLDPXNGOKMIEBFZCWVJAT",
		"UKW-C": "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	}
)

//creates a new rotor using hisotrical configuration
func NewHistoricalRotor(rotorType string) (*Rotor, error) {
	wiring, ok := RotorWirings[rotorType]
	if !ok {
		return nil, ErrInvalidRotorType(rotorType)
	}

	notches, ok := RotorNotches[rotorType]
	if !ok {
		return nil, ErrInvalidRotorType(rotorType)
	}

	return NewRotor(rotorType, wiring, notches)
}

// creates a reflector using historical configurations
func NewHistoricalReflector(reflectorType string) (*Reflector, error) {
	wiring, ok := ReflectorWirings[reflectorType]
	if !ok {
		return nil, ErrInvalidReflectorType(reflectorType)
	}

	return NewReflector(reflectorType, wiring)
}
