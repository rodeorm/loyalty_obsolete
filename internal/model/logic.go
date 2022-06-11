package model

func CheckOrderNum(num string) bool {
	idx := len(num) - 1
	total := 0
	pos := 0
	for i := idx; i > -1; i-- {
		char := num[i]
		if char == ' ' {
			continue
		}
		if char < 48 || char > 57 {
			return false
		}
		digit := int(char - 48)
		if pos%2 == 0 {
			total += digit
		} else {
			switch digit {
			case 1, 2, 3, 4:
				total += digit << 1
			case 5, 6, 7, 8:
				total += (digit << 1) - 9
			case 9:
				total += digit
			}
		}
		pos++
	}
	return pos > 1 && total%10 == 0
}
