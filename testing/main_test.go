package testing

import "github.com/folayao/testing_mod_platzi"

func TestSum(t *testing_mod_platzi.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was Incorrect, got %d expected %d", total, 10)
	}
}
