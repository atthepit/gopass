package pass_generator

import (
	"regexp"
	"strings"
)

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

const minimun_pass_length int = 8

func matches(exp, pass string) []string {
	re := regexp.MustCompile(exp)
	return re.FindAllString(pass, -1)
}

func number_of_matches(exp, pass string) int {
	return len(matches(exp, pass))
}

func number_of_characters(pass string) int {
	return len(pass)
}

func number_of_upper_letters(pass string) int {
	return number_of_matches("[A-Z]", pass)
}

func number_of_lower_letters(pass string) int {
	return number_of_matches("[a-z]", pass)
}

func number_of_numbers(pass string) int {
	return number_of_matches("[0-9]", pass)
}

func number_of_symbols(pass string) int {
	return number_of_matches(`[!"·$%&/()=?¿¡'¬#@|\\\[\]{}\-_.:,;<>ºª]`, pass)
}

func number_of_middle_nums_or_symbols(pass string) int {
	symbols := `[!"·$%&/()=?¿¡'¬#@|\\\[\]{}\-_.:,;<>ºª]`
	return number_of_matches(symbols, pass) - number_of_matches(`^`+symbols, pass) - number_of_matches(symbols+`$`, pass)
}

func contains_upper_letters(pass string) bool {
	return number_of_upper_letters(pass) > 0
}

func contains_lower_letters(pass string) bool {
	return number_of_lower_letters(pass) > 0
}

func contains_numbers(pass string) bool {
	return number_of_numbers(pass) > 0
}

func contains_symbols(pass string) bool {
	return number_of_symbols(pass) > 0
}

func number_of_requiremets_passed(pass string) int {
	count := 0

	if contains_upper_letters(pass) {
		count += 1
	}
	if contains_lower_letters(pass) {
		count += 1
	}
	if contains_numbers(pass) {
		count += 1
	}
	if contains_symbols(pass) {
		count += 1
	}
	if len(pass) >= minimun_pass_length {
		count += 1
	}
	if count == 5 {
		count += 1
	}

	return count
}

func letters_only(pass string) int {
	return number_of_matches(`^[A-z]*$`, pass) * len(pass)
}

func numbers_only(pass string) int {
	return number_of_matches(`^[0-9]*$`, pass) * len(pass)
}

func consecutive_characters(exp, pass string) int {
	groups := matches(exp+`*`, pass)
	count := 0
	for i := 0; i < len(groups); i++ {
		group := groups[i]
		count += len(group) - 1
	}
	return count
}

func consecutive_lower_letters(pass string) int {
	return consecutive_characters(`[a-z]`, pass)
}

func consecutive_upper_letters(pass string) int {
	return consecutive_characters(`[A-Z]`, pass)
}

func consecutive_numbers(pass string) int {
	return consecutive_characters(`[0-9]`, pass)
}

func possible_sequential_groups(exp string, length int) []string {
	var groups []string
	for i := length; i >= 3; i-- {
		for j := 0; j <= len(exp)-i; j++ {
			aux := exp[j : i+j]
			groups = append(groups, aux)
		}
	}
	return groups
}

func possible_sequential_letters(length int) []string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	return possible_sequential_groups(letters, length)
}

func possible_sequential_numbers(length int) []string {
	numbers := "0123456789"
	return possible_sequential_groups(numbers, length)
}

func sequential_groups(exp, pass string) int {
	groups := matches(exp+`*`, pass)
	count := 0
	for i := 0; i < len(groups); i++ {
		group := groups[i]
		if len(group) >= 3 {
			var seq []string
			if exp == `[A-z]` {
				seq = possible_sequential_letters(len(group))
			} else {
				seq = possible_sequential_numbers(len(group))
			}
			exp := strings.Join(seq, "|")
			sequence := matches(exp, strings.ToLower(group))
			for j := 0; j < len(sequence); j++ {
				count += len(sequence[j])
			}
		}
	}
	return count
}

func sequential_letters(pass string) int {
	return sequential_groups(`[A-z]`, pass)
}

func sequential_numbers(pass string) int {
	return sequential_groups(`[0-9]`, pass)
}

func Check_password(pass string) (bool, int) {

	char_num := number_of_characters(pass)

	adds := (char_num * 4) + (char_num - (number_of_upper_letters(pass))*2) + ((char_num - number_of_lower_letters(pass)) * 2) + (number_of_numbers(pass) * 4) + (number_of_symbols(pass) * 6) + (number_of_middle_nums_or_symbols(pass) * 2) + (number_of_requiremets_passed(pass) * 2)

	cons := letters_only(pass) + numbers_only(pass) + (consecutive_lower_letters(pass) * 2) + (consecutive_upper_letters(pass) * 2) + (consecutive_numbers(pass) * 2) + (sequential_letters(pass) * 2) + (sequential_numbers(pass) * 2)

	final_score := adds - cons
	return (final_score >= 90), final_score
}
