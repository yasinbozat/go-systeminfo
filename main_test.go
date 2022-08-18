package main

import "testing"

func TestAddUser(t *testing.T) {

	t.Run("add user", func(t *testing.T) {
		got, _ := SystemInfo("SerialNumber")
		want := "PC27EQPR"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
