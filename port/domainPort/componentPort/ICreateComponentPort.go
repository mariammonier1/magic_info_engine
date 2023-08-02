package componentPort

import "magic_info_engine/domain/domainModel"

type ICreateComponentPort interface {
	CreateComponent(component *domainModel.Component) (string, error)
}
