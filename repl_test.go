package main

import "testing" 

func TestCleanInput(t *testing.T) {
  tests := []struct {
    input string
    expected []string
  } {
    // structure
    {
      input: "Hello World", 
      expected: []string{
        "hello",
        "world",
      },
    },
  }

  for _, test := range tests {
    actual := cleanInput(test.input)     
    if(len(actual) != len(test.expected)) {
      t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(test.expected))
    }

    continue

    for i := range actual {
     actualWord := actual[i]
     expectedWord := test.expected[i]

     if(actualWord != expectedWord) {
         t.Errorf("%v does not equal %v", actualWord, expectedWord)
     }
    }
  
  }
}

