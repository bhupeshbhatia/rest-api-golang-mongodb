package utility

import "fmt"

//CheckError functions lists the error. Should be converted into a error trace package
func CheckError(err) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
