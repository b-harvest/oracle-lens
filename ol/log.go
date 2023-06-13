package ol

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func Info(msg any) {
	message := fmt.Sprint(msg)
	log.Info().Msg(message)
}

func Error(msg any) {
	message := fmt.Sprint(msg)
	log.Error().Msg(message)
}

func Debug(msg any) {
	message := fmt.Sprint(msg)
	log.Debug().Msg(message)
}
