package main

import ("testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
    } {
	    {
    		input:    "  hElLo  wOrLd  ",
	    	expected: []string{"hello", "world"},
    	},
        {
            input: " dO the  moNkey danC3   f0r   aLl tHe  moNk3ys of thE wOrlD!     ",
            expected: []string{"do", "the", "monkey", "danc3", "f0r", "all", "the", "monk3ys", "of", "the", "world!"},
        },
    }

    for _, c := range cases {  // for each case in cases
    	actual := cleanInput(c.input)
    	if len(c.expected) != len(actual) {
            t.Errorf("Length of slices don't match: %d != %d, %v %v", len(c.expected), len(actual), c.expected, actual)
            continue
        }
        // Check the length of the actual slice against the expected slice
	    // if they don't match, use t.Errorf to print an error message
    	// and fail the test
	    for i := range actual {
    		word := actual[i]
	    	expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Word mismatch: %q != %q", expectedWord, word)
            }
   		    // Check each word in the slice
    		// if they don't match, use t.Errorf to print an error message
	    	// and fail the test
	    }
    }
}
