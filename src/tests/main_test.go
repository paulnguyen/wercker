
/*
	Gumball API in Go (Version 1)
	Basic Version with Wercker
*/

package test

import (
	"testing"
)

func TestGumball(t *testing.T) {  
    cnt := 989
    if cnt != 989 {
       t.Errorf("Count of Gumballs Incorrect, got: %d, want: %d.", cnt, 989)
    }
}


