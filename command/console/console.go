package console

type Console struct {
	Signature string
	Description string
	Handle func()
}
