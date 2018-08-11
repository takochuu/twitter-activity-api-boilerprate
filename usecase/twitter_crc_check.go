package usecase

type TwitterCRCCheckUseCase struct{}

type TwitterCRCCheckInterface interface {
	Check(string) error
}

func NewTwitterCRCCheckUseCase() TwitterCRCCheckInterface {
	return &TwitterCRCCheckUseCase{}
}

func (u *TwitterCRCCheckUseCase) Check(crcToken string) error {
	return nil
}
