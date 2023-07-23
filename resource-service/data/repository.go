package data

type Repository interface {
	ResourceRepository
}

type ResourceRepository interface {
	CreateResource(resource Resource) error
	UpdateResource(resource Resource) error
	GetResource(id int64, name string) (*Resource, error)
	GetAllResources(jobTitle, workgroup, employmentType, manager string) ([]*Resource, error)
	DeleteResource(resource Resource)
}
