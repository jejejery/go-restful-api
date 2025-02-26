package sample

type SimpleRepository struct{

}

type SimpleService struct{
	*SimpleRepository
}

func NewSimpleService(simpleRepository *SimpleRepository) *SimpleService{
	return &SimpleService{simpleRepository}
}