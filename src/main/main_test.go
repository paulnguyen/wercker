
/*
	Gumball API in Go (Version 1)
	Basic Version with Wercker
*/

package main

import (
	"testing"
	"gumball"
)

func TestGumball(t *testing.T) {  
    cnt := machine.CountGumballs
    if cnt != 989 {
       t.Errorf("Count of Gumballs Incorrect, got: %d, want: %d.", cnt, 989)
    }
}


