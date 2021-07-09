package GitExercise

import "testing"

func TestHelloworld(t *testing.T) {
	p := sayHello()
	if p != "Hello world!" {
		t.Errorf("jsad")
	}
}
