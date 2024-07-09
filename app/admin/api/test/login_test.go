package test

import (
	"fmt"
	"testing"
)

var address = "127.0.0.1:6010"

func TestHello(t *testing.T) {
	fmt.Println(address)
}

func TestAdd(t *testing.T) {
	fmt.Println("result:", 1+1)
}

func main() {

}
