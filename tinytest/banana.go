package tinytest

type Banana struct {
	count int
}

func (b *Banana) Inc() {
	b.count++
}
