package main

import (
	"log"
	"smartcalc_v3_wadina/internal/model"
	"smartcalc_v3_wadina/internal/view"
	"smartcalc_v3_wadina/internal/viewmodel"
)

func main() {

	m := model.NewModel()
	vm := viewmodel.NewViewModel(m)
	v := view.NewView(vm)
	v.Fire()
	log.Fatalf("Privet: %v, %v, %v", m, vm, v)
}
