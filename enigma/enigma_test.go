package enigma

import (
	"testing"
)

func TestEnigma_EncryptDecrypt(t *testing.T) {
	// Build a historical Enigma I configuration
	rotorI, _ := NewHistoricalRotor("I")
	rotorII, _ := NewHistoricalRotor("II")
	rotorIII, _ := NewHistoricalRotor("III")
	reflectorB, _ := NewHistoricalReflector("UKW-B")
	plugboard, _ := NewPlugboard("AB CD EF")

	enigma := NewEnigma([]*Rotor{rotorI, rotorII, rotorIII}, reflectorB, plugboard)

	// set initial rotor positions and ring settings
	_ = enigma.SetRotorPositions(0, 0, 0)
	rotorI.SetRingSetting(0)
	rotorII.SetRingSetting(0)
	rotorIII.SetRingSetting(0)

	message := "HELLO WORLD"
	encrypted, err := enigma.Encrypt(message)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	if encrypted == message {
		t.Fatalf("Encryption did not change message: got %s", encrypted)
	}

	// Reset rotor positions before decryption
	_ = enigma.SetRotorPositions(0, 0, 0)
	decrypted, err := enigma.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if decrypted != message {
		t.Errorf("Decrypted message does not match original\nOriginal: %s\nDecrypted: %s", message, decrypted)
	}
}

func TestEnigma_LowercaseAndNonAlpha(t *testing.T) {
	rotorI, _ := NewHistoricalRotor("I")
	rotorII, _ := NewHistoricalRotor("II")
	rotorIII, _ := NewHistoricalRotor("III")
	reflectorB, _ := NewHistoricalReflector("UKW-B")

	enigma := NewEnigma([]*Rotor{rotorI, rotorII, rotorIII}, reflectorB, nil)
	_ = enigma.SetRotorPositions(0, 0, 0)

	input := "Hello, World! 123"

	_, err := enigma.Encrypt(input)
	if err != nil {
		t.Fatalf("Encrypt failed on input with lowercase/non-alpha chars: %v", err)
	}
}

func TestEnigma_EncryptConsistency(t *testing.T) {
	// Tests that same input with same initial positions produces same output

	// First Enigma
	rotorI1, _ := NewHistoricalRotor("I")
	rotorII1, _ := NewHistoricalRotor("II")
	rotorIII1, _ := NewHistoricalRotor("III")
	reflectorB1, _ := NewHistoricalReflector("UKW-B")
	enigma1 := NewEnigma([]*Rotor{rotorI1, rotorII1, rotorIII1}, reflectorB1, nil)
	_ = enigma1.SetRotorPositions(0, 0, 0)

	// Second Enigma (new rotor instances)
	rotorI2, _ := NewHistoricalRotor("I")
	rotorII2, _ := NewHistoricalRotor("II")
	rotorIII2, _ := NewHistoricalRotor("III")
	reflectorB2, _ := NewHistoricalReflector("UKW-B")
	enigma2 := NewEnigma([]*Rotor{rotorI2, rotorII2, rotorIII2}, reflectorB2, nil)
	_ = enigma2.SetRotorPositions(0, 0, 0)

	message := "TESTMESSAGE"

	enc1, err := enigma1.Encrypt(message)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	enc2, err := enigma2.Encrypt(message)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	if enc1 != enc2 {
		t.Errorf("Encryption is not consistent: got %s vs %s", enc1, enc2)
	}
}
