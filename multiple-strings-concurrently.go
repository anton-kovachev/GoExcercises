/*
    Write a function that concurrently multiplies strings by a given factor.
    This function needs to accept a slice of strings and a positive (possibly
    zero), integer factor and return a new slice containing the resulting
    multiplied strings.
  
    The multiplication of a string is simply the repeated concatenation of
    itself. For example, multiplying "tim" by
    3 results in "timtimtim".
  
    The slice that is returned must contain the multiplied strings in the order
    in which they appeared in the original input slice. This function must also
    perform each string multiplication concurrently, using a go routine.
  
  	Note: Multiplying a string by zero results in an empty string: "".
	  
	Sample Input
	strings = []string{"bird", "plane", "superman"}  
	factor = 3  

	Sample Output
	[birdbirdbird planeplaneplane supermansupermansuperman]
*/

package main

type MultipliedString struct {
    str string
    idx uint
}

func MultiplyStringsConcurrently(strings []string, factor uint) []string {
	// Write your code here.
    stringResults := make([]string,len(strings))
    resultChannel := make(chan MultipliedString, len(strings))
    
    for i, str := range strings {
        go MultiplyString(str, uint(i), factor, resultChannel)
    }

    for i := 0; i < len(strings); i++ {
        multipledString := <- resultChannel
        stringResults[multipledString.idx] = multipledString.str
    }
    
    
    return stringResults
}

func MultiplyString(str string, idx uint, factor uint, rc chan <- MultipliedString) {
    resultString := ""
    
    for i := uint(0); i < factor; i++ {
        resultString += str
    }

    rc <- MultipliedString{resultString, idx}
}
