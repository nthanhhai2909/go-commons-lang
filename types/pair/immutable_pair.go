package pair

type ImmutablePair struct {
	pair
}

func NewImmutablePair(left, right interface{}) *MutablePair {
	return &MutablePair{pair{left: left, right: right}}
}

func (pair *ImmutablePair) GetLeft() interface{} {
	return pair.left
}

func (pair *ImmutablePair) GetRight() interface{} {
	return pair.right
}
