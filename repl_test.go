package main 

import "testing"


func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		output []string 
	}{
		{
			input: "hello world",
			output: []string{
				"hello",
				"world",
			},
		},
	}

	for _,case1 := range cases {
		actual_val := cleanInput(case1.input)
		if len(actual_val) != len(case1.output) {
			t.Errorf("The lenghts are not same")
		}
		
		for i,_ := range actual_val {
			actual_word := actual_val[i]
			expected_word := case1.output[i]
			
			
			if actual_word != expected_word {
				t.Errorf("words are not equal ")
			}
		}
	}
	 
	


}