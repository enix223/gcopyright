package command

import (
	"bufio"
	"log"
	"os"
	"strings"

	arg "github.com/alexflint/go-arg"
	"github.com/gingfrederik/docx"
	"github.com/yargevad/filepathx"
)

func assertAbort() {
	if RunOptions.AbortIfError {
		os.Exit(-1)
	}
}

func Run() {
	// parse args
	arg.MustParse(&RunOptions)

	doc := docx.NewFile()

	var lines int
	for _, pattern := range RunOptions.Patterns {
		matchGlobs, err := filepathx.Glob(pattern)
		if err != nil {
			log.Printf("[ERR] failed to find with pattern: %s, err: %s", pattern, err)
			assertAbort()
		}

		for _, match := range matchGlobs {
			lines += readFile(match, doc, lines)
			if lines >= RunOptions.TotalLines {
				log.Print("[INFO] reach line limit")
				saveDoc(lines, doc)
				return
			}
		}
	}

	saveDoc(lines, doc)
}

func saveDoc(lines int, doc *docx.File) {
	log.Printf("[INFO] done, total lines: %d", lines)
	if err := doc.Save(RunOptions.Output); err != nil {
		log.Printf("[ERR] failed to save docx, path: %s, err: %s", RunOptions.Output, err)
	}
}

func readFile(path string, doc *docx.File, currentLines int) (lines int) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("[ERR] failed to read file: %s, err: %s", path, err)
		assertAbort()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	spaces := []rune{'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0}
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimRight(line, string(spaces))
		if !RunOptions.AllowEmptyLine && len(trimmed) == 0 {
			continue
		}
		if RunOptions.Verbose {
			log.Println(line)
		}
		doc.AddParagraph().AddText(trimmed)
		lines++
		currentLines++
	}
	log.Printf("[INFO] read file [%s] success, total lines: %d", path, lines)
	return
}
