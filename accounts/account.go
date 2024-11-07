package accounts

import (
	"reflect"
	"strconv"

	"github.com/gagliardetto/solana-go"
)

// ParseAccountsIntoStruct is a generic function that parses accounts into any struct with account tags
func ParseAccountsIntoStruct[T any](accounts []solana.PublicKey) (result T) {
	resultValue := reflect.ValueOf(&result).Elem()
	resultType := resultValue.Type()

	for i := 0; i < resultType.NumField(); i++ {
		field := resultType.Field(i)
		if accountIdx, ok := field.Tag.Lookup("idx"); ok {
			idx, err := strconv.Atoi(accountIdx)
			if err != nil {
				continue
			}
			if idx < len(accounts) {
				resultValue.Field(i).Set(reflect.ValueOf(accounts[idx]))
			}
		}
	}
	return result
}
