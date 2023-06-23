package viewmodel

import (
	"smartcalc_v3_wadina/internal/model"
)

type ViewModel struct {
	model *model.Model
}

func NewViewModel(m *model.Model) *ViewModel {
	return &ViewModel{
		model: m,
	}
}

func (vm *ViewModel) Calculate(line string) string {
	return vm.model.Calculate(line)
}
