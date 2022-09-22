package field
type _Option uint8
const (
	_AI _Option = iota
	_Nullable
	_Unique
)

func (o *_Option) CheckOp(option _Option) bool {
	return ((*o) & (1 << option)) != 0
}

func (o *_Option) SetOp(option _Option) {
	if !o.CheckOp(option) {
		*o |= 1 << option
	}
}

func (o *_Option) ClearOp(option _Option) {
	if o.CheckOp(option) {
		*o &= ^(1 << option)
	}
}