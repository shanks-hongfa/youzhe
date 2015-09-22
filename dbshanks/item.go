package dbshanks


type Item struct {
	name string
	url  string
	pic  []string
}

type Bag struct {
	base     Item
	color    string
	size     Size
	weight   uint32
	material string
}

/**
尺寸 长宽高
 */
type Size struct {
	x uint32
	y uint32
	z uint32
}

func (bag *Bag)QueryAll() {

}