// saludo_test
package main

import (
	"testing"
)

func TestSaludo(t *testing.T) {
	expectativa := "esto es una prueba!"
	resultado := saludo()
	if resultado != expectativa {
		t.Fatalf("expectativa %s,go %s", expectativa, resultado)
		//fmt.Println(t, expectativa, resultado, "no se ejecuto correctamente")
	}
}
