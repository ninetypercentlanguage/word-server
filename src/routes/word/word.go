package word

import (
	"audio-language/words/server/rediscli"
	"audio-language/words/server/words"

	"github.com/go-chi/chi"
)

func GetWordSubRoute(wordList []words.WordString, cli *rediscli.WordRedisCli) func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/", getSearchWordsRoute(wordList, cli))
		router.Get("/{partOfSpeechId}", getFindPartOfSpeechRoute(cli))
	}
}
