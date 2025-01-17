package main

import "testing"


func TestRemoveSpecialCharacters(t *testing.T) {
    output := removeSpecialCharacters("¡Hola! Mi nombre es A@ron González, ¿cual es el tuyo?")

    if output != "HolaMinombreesAronGonzálezcualeseltuyo" {
        t.Fatalf("got %s\n", output)
    }
}
