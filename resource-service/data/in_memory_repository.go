package data

type InMemoryRepository struct {
	Resources []Resource
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		Resources: []Resource{},
	}
}

func (r *InMemoryRepository) CreateResource(resource Resource) error {
	r.Resources = append(r.Resources, resource)
	return nil
}

func (r *InMemoryRepository) UpdateResource(resource Resource) error {
	for i, res := range r.Resources {
		if res.Email == resource.Email {
			r.Resources[i] = resource
		}
	}
	return nil
}

func (r *InMemoryRepository) GetResource(id int64, name string) (*Resource, error) {
	for _, res := range r.Resources {
		if res.ID == id || res.Name == name {
			return &res, nil
		}
	}
	return &Resource{}, nil
}

func (r *InMemoryRepository) GetAllResources(jobTitle, workgroup, employmentType, manager string) ([]*Resource, error) {
	return []*Resource{}, nil
}

func (r *InMemoryRepository) DeleteResource(resource Resource) {
	for i, res := range r.Resources {
		if res.ID == resource.ID {
			r.Resources = append(r.Resources[:i], r.Resources[i+1:]...)
		}
	}
}
