package utilities

import (
	"errors"
	"fmt"
	"strconv"
)

func ConvertStringToInt64InPennies(s string) (int64, error) {
	valAsFloat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	valAsFloat *= 100
	valAsInt := int(valAsFloat)
	result, err := strconv.ParseInt(fmt.Sprintf("%v", valAsInt), 10, 64)

	if err != nil {
		return 0, err
	}
	return result, nil
}

func FormatStringIntoDollarsAndCents(amount string) (string, error) {
	if amount == "" {
		return "", errors.New("invalid (empty) amount received")
	}

	if len(amount) > 2 {
		cents := amount[len(amount) - 2:]
		dollars := amount[:len(amount) - 2]

		return dollars + "." + cents, nil
	}

	return "0." + amount, nil
}