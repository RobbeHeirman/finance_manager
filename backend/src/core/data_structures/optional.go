package data_structures

type Optional[U any] struct {
	field *U
}

func CreateOptional[U any](field *U) *Optional[U] {
	return &Optional[U]{field: field}
}

func CreateEmptyOptional[U any]() *Optional[U] {
	return &Optional[U]{field: nil}
}

func (opt *Optional[U]) Get() (*U, bool) {
	return opt.field, opt.field != nil
}

func (opt *Optional[U]) GetUnchecked() *U {
	return opt.field
}

func (opt *Optional[U]) IsEmpty() bool {
	return opt.field == nil
}
