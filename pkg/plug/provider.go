package plug

type ProviderInterface interface {
	Info() Info
	DescribeCard(arg DescribeCardArg) Card
	DescribePanel(arg DescribePanelArg) Panel
	Register(arg RegisterArg) error
	Get(arg GetArg) Record
	Set(arg SetArg) error
	Del(arg DelArg) error
}

type DescribeCardArg struct {
	Name string
}
type DescribePanelArg struct {
	Name string
}
type RegisterArg struct {
	Model string
	Name  string
}
type GetArg struct {
	Model string
	Name  string
}
type SetArg struct {
	Model  string
	Name   string
	Record Record
}
type DelArg struct {
	Model string
	Name  string
}
