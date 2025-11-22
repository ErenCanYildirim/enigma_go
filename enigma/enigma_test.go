package enigma

import (
	"testing"
)

func TestRotorForwardBackward(t *testing.T) {
	rotor, err := NewRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q")
	if err != nil {
		t.Fatalf("failed to create rotor: %v", err)
	}

	// Test forward/backward mapping with no rotation
	for i := 0; i < AlphabetSize; i++ {
		fwd := rotor.Forward(i)
		bwd := rotor.Backward(fwd)
		if bwd != i {
			t.Errorf("Rotor forward/backward mismatch: input %d got %d", i, bwd)
		}
	}

	// Test with ring setting and position
	rotor.SetRingSetting(1)
	rotor.SetPosition(5)
	for i := 0; i < AlphabetSize; i++ {
		fwd := rotor.Forward(i)
		bwd := rotor.Backward(fwd)
		if bwd != i {
			t.Errorf("Rotor with ring/pos mismatch: input %d got %d", i, bwd)
		}
	}
}

func TestReflector(t *testing.T) {
	ref, err := NewReflector("B", "YRUHQSLDPXNGOKMIEBFZCWVJAT")
	if err != nil {
		t.Fatalf("failed to create reflector: %v", err)
	}

	for i := 0; i < AlphabetSize; i++ {
		out := ref.Reflect(i)
		// reflector is reciprocal
		if ref.Reflect(out) != i {
			t.Errorf("Reflector not reciprocal for %d", i)
		}
	}
}

func TestPlugboard(t *testing.T) {
	pb, err := NewPlugboard("AB CD EF")
	if err != nil {
		t.Fatalf("failed to create plugboard: %v", err)
	}

	tests := map[int]int{
		int('A' - 'A'): int('B' - 'A'),
		int('B' - 'A'): int('A' - 'A'),
		int('C' - 'A'): int('D' - 'A'),
		int('D' - 'A'): int('C' - 'A'),
		int('E' - 'A'): int('F' - 'A'),
		int('F' - 'A'): int('E' - 'A'),
		int('G' - 'A'): int('G' - 'A'),
	}

	for input, expected := range tests {
		out := pb.Forward(input)
		if out != expected {
			t.Errorf("Plugboard mismatch for %d: got %d, want %d", input, out, expected)
		}
	}
}

func TestEnigmaEncryptDecrypt(t *testing.T) {
	rotorI, _ := NewRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q")
	rotorII, _ := NewRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", "E")
	rotorIII, _ := NewRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", "V")
	ref, _ := NewReflector("B", "YRUHQSLDPXNGOKMIEBFZCWVJAT")
	pb, _ := NewPlugboard("AQ EP")

	enigma := NewEnigma([]*Rotor{rotorIII, rotorII, rotorI}, ref, pb)
	enigma.SetRotorPositions(0, 0, 0)

	plaintext := "HELLO WORLD"
	ciphertext, err := enigma.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	// reset positions to decrypt
	enigma.SetRotorPositions(0, 0, 0)
	decrypted, err := enigma.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	if decrypted != "HELLO WORLD" {
		t.Errorf("decryption mismatch: got %s, want %s", decrypted, plaintext)
	}
}

func TestRotorStepping(t *testing.T) {
	rotorI, _ := NewRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q")
	rotorII, _ := NewRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", "E")
	rotorIII, _ := NewRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", "V")
	ref, _ := NewReflector("B", "YRUHQSLDPXNGOKMIEBFZCWVJAT")

	enigma := NewEnigma([]*Rotor{rotorIII, rotorII, rotorI}, ref, nil)

	// Step the rotors 26 times and check wrap-around
	for i := 0; i < 26; i++ {
		enigma.stepRotors()
	}

	positions := enigma.GetRotorPositions()
	for i, pos := range positions {
		if pos < 0 || pos >= AlphabetSize {
			t.Errorf("Rotor %d has invalid position: %d", i, pos)
		}
	}
}
