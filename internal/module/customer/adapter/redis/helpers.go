package redis

import "fmt"

func GetCustomerKey(customerId int) string {
	return fmt.Sprintf("customer-%d", customerId)
}
