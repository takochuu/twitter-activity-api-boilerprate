package usecase

type TwitterCRCCheckUseCase struct{}

type TwitterCRCCheckInterface interface {
	Check() error
}

func NewTwitterCRCCheckUseCase() *TwitterCRCCheckUseCase {
	return &TwitterCRCCheckUseCase{}
}

func (u TwitterCRCCheckUseCase) Check() error {
	return nil
}
