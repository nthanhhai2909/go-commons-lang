package pair

type MutablePair struct {
	pair
}

func NewMutablePair(left, right interface{}) *MutablePair {
	return &MutablePair{pair{left: left, right: right}}
}

func (pair *MutablePair) GetLeft() interface{} {
	return pair.left
}

func (pair *MutablePair) GetRight() interface{} {
	return pair.right
}

func (pair *MutablePair) SetLeft(left interface{}) {
	pair.left = left
}

func (pair *MutablePair) SetRight(right interface{}) {
	pair.right = right
}
