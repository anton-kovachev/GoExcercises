package main
import "fmt"

/*
 	Write a function that accepts a pointer to a 2-dimensional slice of
 	slices. Your function should modify this slice in-place
	 so, each nested slice only contains their longest and shortest word. In
 	other words, remove all elements from each nested slice that is

    the longest or shortest word contained within it. Your function should not
    return a value, it should simply modify the input slice.

    You may assume each nested slice will only contain words with unique lengths
    (i.e., no two words have the same length). If a slice only contains one
    word, that word is considered both the longest and the shortest word and
    should remain in the slice. If a slice contains no words, then no change is
    necessary. 
  
  
    Note: The remaining words should maintain the order in which they appear in
    the original slice.
*/


func KeepLongestAndShortestWord(wordSlices *[][]string) {
    for i := 0; i < len(*wordSlices); i++ {
        short, long := findShortestAndLongest(&((*wordSlices)[i]))
        newArr := []string{}

        for _, word := range (*wordSlices)[i] {
           if word == short {
               newArr = append(newArr, word)
           } else if word == long && long != short {
               newArr = append(newArr, word)
           }
        }

        (*wordSlices)[i] = newArr
    }

    fmt.Println(*wordSlices)
}

func findShortestAndLongest(words *[]string) (string, string) {
        var short string = ""
        var long string = ""

        if len(*words) > 0 {
            short = (*words)[0]
            long = (*words)[0]
        }

        for i := 1; i < len(*words); i++ {
           if len(short) > len((*words)[i]) {
                short = (*words)[i]
            } else if len(long) < len((*words)[i]) {
                long = (*words)[i]
            }
        }

        return short, long
}
