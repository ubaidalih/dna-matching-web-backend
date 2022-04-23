package algorithm

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(fmt.Sprintf(".\\src\\algorithm\\%s", filename))
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}
	return string(data)
}

func ValidateInput(s string) bool {
	regex, err := regexp.Compile("^[ACGT]+$")
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	return regex.MatchString(s)
}

func ValidateQuery(s string) bool {
	regex1, err := regexp.Compile(`^(31(\/|-|\.|\s)(0[13578]|1[02]|(Januari|Maret|Mei|Juli|Agustus|Oktober|Desember))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	if regex1.MatchString(s) {
		return true
	}
	regex2, err := regexp.Compile(`^((29|30)(\/|-|\.|\s)(0[1,3-9]|1[0-2]|(Januari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember))(\/|-|\.|\s)(([6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	if regex2.MatchString(s) {
		return true
	}
	regex3, err := regexp.Compile(`^(29(\/|-|\.|\s)(02|(Februari))(\/|-|\.|\s)(((1[6-9]|[2-9]\d)(0[48]|[2468][048]|[13579][26]))))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	if regex3.MatchString(s) {
		return true
	}
	regex4, err := regexp.Compile(`^(29(\/|-|\.|\s)(02|(Februari))(\/|-|\.|\s)((16|[2468][048]|[3579][26])00))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	if regex4.MatchString(s) {
		return true
	}
	regex5, err := regexp.Compile(`^((0[1-9]|1\d|2[0-8])(\/|-|\.|\s)(0[1-9]|(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	if regex5.MatchString(s) {
		return true
	}
	regex6, err := regexp.Compile(`^((0[1-9]|1\d|2[0-8])(\/|-|\.|\s)(1[0-2]|(Oktober|November|Desember))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	if regex6.MatchString(s) {
		return true
	}
	return false
}
