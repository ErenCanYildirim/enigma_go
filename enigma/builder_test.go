package enigma

import "testing"

func TestBuilderEndToEnd(t *testing.T) {
	builder := NewBuilder().
		WithRotors("I", "II", "III").
		WithReflector("UKW-B").
		WithPlugboard("AB CD EF").
		WithRotorPositionsFromString("AAA").
		WithRingSettingsFromString("AAA")

	machine, err := builder.Build()
	if err != nil {
		t.Fatalf("failed to build enigma: %v", err)
	}

	plaintext := "HELLO"
	encrypted, err := machine.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("encrypt failed: %v", err)
	}

	// Reset to starting configuration to decrypt
	builder = NewBuilder().
		WithRotors("I", "II", "III").
		WithReflector("UKW-B").
		WithPlugboard("AB CD EF").
		WithRotorPositionsFromString("AAA").
		WithRingSettingsFromString("AAA")

	machine2, err := builder.Build()
	if err != nil {
		t.Fatalf("failed to rebuild enigma: %v", err)
	}

	decrypted, err := machine2.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("decrypt failed: %v", err)
	}

	if decrypted != plaintext {
		t.Fatalf("decryption mismatch: got %s, want %s", decrypted, plaintext)
	}
}
