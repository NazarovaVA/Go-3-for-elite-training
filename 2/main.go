package main

func convert(input int64) uint16 {
	if input < 0 || input > 65535 {
		return 0
	}
	return uint16(input)
}
