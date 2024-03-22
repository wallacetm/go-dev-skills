package task

// Task1 demonstrates the use of the CatFeline adapter
// CatFeline adapt the Cat.Meow() so it can be used as a Feline.Roar()
// The RoarService receives only felines, so the CatFeline is used to adapt the Cat to a Feline
func (t *Task) Task1() {
	t.felineService.RoarAll()
}
