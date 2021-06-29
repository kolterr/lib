package math

// TernaryExpression 三元表达式
func TernaryExpression(condition bool, val1, val2 interface{}) interface{} {
	if condition {
		return val1
	}
	return val2
}
