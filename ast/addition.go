package ast

type Addition struct{}

func (a Addition) Calculate(exp *BinaryExpression) (interface{}, error) {
	if exp.A.GetType() == Number && exp.B.GetType() == Number {
		valA, _ := exp.A.Get()
		valB, _ := exp.B.Get()
		return valA.(float64) + valB.(float64), nil
	}
	return exp.A.ToString() + exp.B.ToString(), nil
}
