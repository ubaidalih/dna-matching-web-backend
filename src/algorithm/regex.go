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

func ValidateQuery(s string) int {
	regex1, err := regexp.Compile(`^(31(\/|-|\.|\s)(0[13578]|1[02]|(Januari|Maret|Mei|Juli|Agustus|Oktober|Desember))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex1.MatchString(s) {
		return 1
	}
	regex1, err = regexp.Compile(`^(31(\/|-|\.|\s)(0[13578]|1[02]|(Januari|Maret|Mei|Juli|Agustus|Oktober|Desember))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex1.MatchString(s) {
		return 2
	}
	regex2, err := regexp.Compile(`^((29|30)(\/|-|\.|\s)(0[1,3-9]|1[0-2]|(Januari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember))(\/|-|\.|\s)(([6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex2.MatchString(s) {
		return 1
	}
	regex2, err = regexp.Compile(`^((29|30)(\/|-|\.|\s)(0[1,3-9]|1[0-2]|(Januari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember))(\/|-|\.|\s)(([6-9]|[2-9]\d)\d{2}))$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex2.MatchString(s) {
		return 2
	}
	regex3, err := regexp.Compile(`^(29(\/|-|\.|\s)(02|(Februari))(\/|-|\.|\s)(((1[6-9]|[2-9]\d)(0[48]|[2468][048]|[13579][26]))))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex3.MatchString(s) {
		return 1
	}
	regex3, err = regexp.Compile(`^(29(\/|-|\.|\s)(02|(Februari))(\/|-|\.|\s)(((1[6-9]|[2-9]\d)(0[48]|[2468][048]|[13579][26]))))$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex3.MatchString(s) {
		return 2
	}
	regex4, err := regexp.Compile(`^(29(\/|-|\.|\s)(02|(Februari))(\/|-|\.|\s)((16|[2468][048]|[3579][26])00))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex4.MatchString(s) {
		return 1
	}
	regex4, err = regexp.Compile(`^(29(\/|-|\.|\s)(02|(Februari))(\/|-|\.|\s)((16|[2468][048]|[3579][26])00))$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex4.MatchString(s) {
		return 2
	}
	regex5, err := regexp.Compile(`^((0[1-9]|1\d|2[0-8])(\/|-|\.|\s)(0[1-9]|(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex5.MatchString(s) {
		return 1
	}
	regex5, err = regexp.Compile(`^((0[1-9]|1\d|2[0-8])(\/|-|\.|\s)(0[1-9]|(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex5.MatchString(s) {
		return 2
	}
	regex6, err := regexp.Compile(`^((0[1-9]|1\d|2[0-8])(\/|-|\.|\s)(1[0-2]|(Oktober|November|Desember))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))(\s)(.*)$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex6.MatchString(s) {
		return 1
	}
	regex6, err = regexp.Compile(`^((0[1-9]|1\d|2[0-8])(\/|-|\.|\s)(1[0-2]|(Oktober|November|Desember))(\/|-|\.|\s)((1[6-9]|[2-9]\d)\d{2}))$`)
	if err != nil {
		fmt.Println("Error: ", err)
		return -1
	}
	if regex6.MatchString(s) {
		return 2
	}
	return 3
}

func ParseQuery(s string, n int) []string {
	var result []string
	i := 0
	if n == 3 {
		query := []string{"", s}
		return query
	}
	for j := 0; j < len(s); j++ {
		if j == len(s)-1 {
			result = append(result, s[i:])
			i = j + 1
		}
		if s[j] == ' ' || s[j] == '/' || s[j] == '-' {
			result = append(result, s[i:j])
			i = j + 1
			if len(result) == 3 {
				break
			}
		}
	}
	if i == len(s) {
		result = append(result, "")
	} else {
		result = append(result, s[i:])
	}
	if len(result[0]) == 1 {
		result[0] = "0" + result[0]
	}
	bulan := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	for i := 0; i < len(bulan); i++ {
		if result[1] == bulan[i] {
			if i+1 < 10 {
				result[1] = fmt.Sprintf("0%d", i+1)
			} else {
				result[1] = fmt.Sprintf("%d", i+1)
			}
			break
		}
	}
	query := []string{fmt.Sprintf("%s-%s-%s", result[2], result[1], result[0]), result[3]}
	return query
}
