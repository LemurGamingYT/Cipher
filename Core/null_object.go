package core

type NullObject struct{}

func (n *NullObject) Repr(context any) string {
	return "null"
}

func (n *NullObject) Add(other any) any {
	ReportOperatorError("+")
	return nil
}

func (n *NullObject) Sub(other any) any {
	ReportOperatorError("-")
	return nil
}

func (n *NullObject) Mul(other any) any {
	ReportOperatorError("*")
	return nil
}

func (n *NullObject) Div(other any) any {
	ReportOperatorError("/")
	return nil
}

func (n *NullObject) Mod(other any) any {
	ReportOperatorError("%")
	return nil
}

func (n *NullObject) Pow(other any) any {
	ReportOperatorError("**")
	return nil
}

func (n *NullObject) Eq(other any) any {
	ReportOperatorError("==")
	return nil
}

func (n *NullObject) Neq(other any) any {
	ReportOperatorError("!=")
	return nil
}

func (n *NullObject) Gt(other any) any {
	ReportOperatorError(">")
	return nil
}

func (n *NullObject) Lt(other any) any {
	ReportOperatorError("<")
	return nil
}

func (n *NullObject) Gte(other any) any {
	ReportOperatorError(">=")
	return nil
}

func (n *NullObject) Lte(other any) any {
	ReportOperatorError("<=")
	return nil
}

func (n *NullObject) And(other any) any {
	ReportOperatorError("&&")
	return nil
}

func (n *NullObject) Or(other any) any {
	ReportOperatorError("||")
	return nil
}

func NewNullObject() *NullObject {
	return &NullObject{}
}
