package component

import (
	"magic_info_engine/domain/domainModel"
	"magic_info_engine/port/domainPort/componentPort"
)

type createComponentUseCase struct {
}

func NewCreateComponentUseCase(
) componentPort.ICreateComponentPort {
	return &createComponentUseCase{}
}
func (thisCR *createComponentUseCase)	CreateComponent(component *domainModel.Component) (string, error){
	
	return "savfdbg ", nil
}
