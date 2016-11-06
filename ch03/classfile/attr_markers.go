package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {}

func (self *MarkerAttribute) readInfo (readre *ClassReader) {
	// read nothing
}