package usecase

import (
	"os"

	"github.com/takochuu/twitter-activity-api-boilerprate/lib"
)

type TwitterCRCCheckUseCase struct{}

type TwitterCRCCheckInteractor interface {
	Check(string) string
}

func NewTwitterCRCCheckUseCase() TwitterCRCCheckInteractor {
	return &TwitterCRCCheckUseCase{}
}

func (u *TwitterCRCCheckUseCase) Check(crcToken string) string {
	cs := os.Getenv("PAIRS_TWITTER_BOT_CONSUMER_SECRET")
	token := lib.CreateCRCToken(crcToken, cs)
	return token
}
