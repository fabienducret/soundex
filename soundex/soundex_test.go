package soundex_test

import (
	"soundex/soundex"
	"testing"
)

func TestSoundex(t *testing.T) {
	t.Run("Retains sole letter of one letter word", func(t *testing.T) {
		assertEqual(t, soundex.Encode("A"), "A000")
	})

	t.Run("Pads with zeros to ensure three digits", func(t *testing.T) {
		assertEqual(t, soundex.Encode("I"), "I000")
	})

	t.Run("Replace consonants with appropriate digits for Ab", func(t *testing.T) {
		assertEqual(t, soundex.Encode("Ax"), "A200")
	})

	t.Run("Ignore non alphabetics", func(t *testing.T) {
		assertEqual(t, soundex.Encode("A#"), "A000")
	})

	t.Run("Replace multiple consonants with digits", func(t *testing.T) {
		assertEqual(t, soundex.Encode("Acdl"), "A234")
	})

	t.Run("Limits length to four characters", func(t *testing.T) {
		assertEqual(t, len(soundex.Encode("Dcdlb")), 4)
	})

	t.Run("Ignores vowel like letter", func(t *testing.T) {
		assertEqual(t, soundex.Encode("BaAeEiIoOuUhHyYcdl"), "B234")
	})

	t.Run("Combine duplicate encodings", func(t *testing.T) {
		assertEqual(t, soundex.Encode("Abfcgdt"), "A123")
	})

	t.Run("Uppercases first letter", func(t *testing.T) {
		assertEqual(t, soundex.Encode("abcd"), "A123")
	})

	t.Run("Ignores case when encoding consonants", func(t *testing.T) {
		assertEqual(t, soundex.Encode("Bcdl"), soundex.Encode("BCDL"))
	})
}

func assertEqual[T comparable](t *testing.T, got T, want T) {
	if got != want {
		t.Errorf("error, got %v, want %v", got, want)
	}
}
