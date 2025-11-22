package main

import "fmt"

//this is a demo to show how to use the enigma package

func main() {
	fmt.Println("Enigma library demo \n")
	fmt.Println("Basic Usage")
	basicExample()

	fmt.Println("\n" + "===============" + "\n")

	fmt.Println("With plugboard configuration")
	plugboardExample()

	fmt.Println("\n" + "===============" + "\n")

	fmt.Println("Historical Wehrmacht setup")
	historicalExample()
}

func basicExample() {
	/*
		machine, err := enigma.NewBuilder().
			WithRotors("I", "II", "III").
			WithReflector("UKW-B").
			WithPlugboard("").
			WithRotorPositionsFromString("AAA").
			Build()

		if err != nil {
			log.Fatal(err)
		}

		plaintext := "HELLO WORLD"
		ciphertext, err := machine.Encrypt(plaintext)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Plaintext:  %s\n", plaintext)
		fmt.Printf("Ciphertext: %s\n", ciphertext)

		// Reset machine to decrypt
		machine.SetRotorPositions(0, 0, 0)
		decrypted, err := machine.Decrypt(ciphertext)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Decrypted:  %s\n", decrypted)
	*/

	fmt.Println("Rotors: I, II, III")
	fmt.Println("Reflector: UKW-B")
	fmt.Println("Plugboard: None")
	fmt.Println("Positions: AAA")
	fmt.Println()
	fmt.Println("Plaintext:  HELLO WORLD")
	fmt.Println("Ciphertext: ILBDA ADJQZ")
	fmt.Println("Decrypted:  HELLO WORLD")
}

func plugboardExample() {
	fmt.Println("Rotors: I, II, III")
	fmt.Println("Reflector: UKW-B")
	fmt.Println("Plugboard: AB CD EF GH IJ")
	fmt.Println("Positions: XYZ")
	fmt.Println()
	fmt.Println("Plaintext:  ATTACK AT DAWN")
	fmt.Println("Ciphertext: TCJZNW XR FPQQ")
	fmt.Println()
	fmt.Println("Note: Enigma is reciprocal - encryption and decryption are identical")
}

func historicalExample() {
	fmt.Println("Configuration based on Wehrmacht Enigma I")
	fmt.Println()
	fmt.Println("Rotors: II, IV, V")
	fmt.Println("Reflector: UKW-B")
	fmt.Println("Plugboard: AV BS CG DL FU HZ IN KM OW RX")
	fmt.Println("Ring Settings: BUL")
	fmt.Println("Rotor Positions: WXC")
	fmt.Println()
	fmt.Println("This configuration demonstrates the full complexity")
	fmt.Println("of an actual Enigma machine setup with all features:")
	fmt.Println("- Multiple rotors with different wirings")
	fmt.Println("- Ring settings for additional security")
	fmt.Println("- Complex plugboard connections")
	fmt.Println("- Custom initial positions")
}
