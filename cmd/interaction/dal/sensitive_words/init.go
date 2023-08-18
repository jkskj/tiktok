package sensitive_words

import (
	"bufio"
	"io"
	"os"

	"github.com/ozline/tiktok/pkg/utils"
)

var St *utils.SensitiveTrie

func Init() {
	fileHandle, err := os.OpenFile("cmd/interaction/dal/sensitive_words/words.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	reader := bufio.NewReader(fileHandle)

	var words []string
	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		words = append(words, string(line))
	}
	St = utils.NewSensitiveTrie()
	St.AddWords(words)
}