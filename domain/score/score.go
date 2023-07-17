package score

import (
	"github.com/opensourceways/xihe-script/infrastructure/message"
)

type CalculateScore interface {
	Calculate(*message.MatchFields) ([]byte, error)
}

type EvaluateScore interface {
	Evaluate(*message.MatchFields) ([]byte, error)
}
