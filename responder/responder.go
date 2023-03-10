package responder

import (
	"bot/config"

	openai "github.com/sashabaranov/go-openai"
)

type Responsder struct {
	client *openai.Client
	alive  bool
}

func NewResponsder() *Responsder {
	return &Responsder{}
}

func (ai *Responsder) Init(cfg *config.Config) {
	if cfg.Token == "" {
		panic("token为空")
	}

	ai.client = openai.NewClient(cfg.Token)
}

func (ai *Responsder) Chat(string, string) (string, error) {

	return "", nil
}

func (ai *Responsder) Alive() bool {
	return ai.alive
}
