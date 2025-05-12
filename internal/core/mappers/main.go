package mappers

// MapperFactory provides centralized access to all mappers
type MapperFactory struct {
	userMapper *UserMapper
}

// NewMapperFactory creates a new MapperFactory instance
func NewMapperFactory() *MapperFactory {
	return &MapperFactory{
		userMapper: NewUserMapper(),
	}
}

// UserMapper returns the UserMapper instance
func (f *MapperFactory) UserMapper() *UserMapper {
	return f.userMapper
}
