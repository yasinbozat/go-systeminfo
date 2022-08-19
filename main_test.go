package main

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {

	t.Run("add user", func(t *testing.T) {
		got, _ := SystemInfo("SerialNumber")
		want := "PC27EQ9S" //PC27EQ9S
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		fmt.Println("\"" + got + "\"")
		fmt.Println("\"" + want + "\"")
		t.Errorf("got %v want %v", got, want)
	}
}
