package main

type Weapon interface {
	Attack()
	Repair()
}

type Bow struct{}

func (b *Bow) Attack() {
	println("bow Attack")
}
func (b *Bow) Repair() {
	println("bow Repair")
}

type Sword struct{}

func (s *Sword) Attack() {
	println("sword Attack")
}
func (s *Sword) Repair() {
	println("sword Repair")
}

type WeaponHandler interface {
	Handle()
}

type Warrior struct {
	weapon Weapon
}

func (w *Warrior) Handle() {
	w.weapon.Attack()
}

type Smith struct {
	weapon Weapon
}

func (s *Smith) Handle() {
	s.weapon.Repair()
}

// WeaponHandler 인터페이스를 인자로 받는 함수를 만들어봅시다!
func doit(w WeaponHandler) {
	w.Handle()
}

func main() {
	sword := &Sword{}      // 칼을만들고
	smith := &Smith{sword} // 스미스에게 칼을 줍니다

	warrior := &Warrior{sword} // 워리어에게 칼을 줍니다

	doit(warrior) // "sword Attack"
	doit(smith)   // "sword Repair"
}
