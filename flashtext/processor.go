package flashtext

import (
	"strings"
	"unicode"
)

type KeywordProcessor struct {
	caseSensitive bool
	rootNode      *KeywordTrieNode
}

func NewKeyWordProcessor(caseSensitive bool) *KeywordProcessor {
	return &KeywordProcessor{
		caseSensitive: caseSensitive,
		rootNode:      &KeywordTrieNode{},
	}
}

func (k *KeywordProcessor) AddKeyWord(word string) {
	if k.caseSensitive {
		word = strings.ToLower(word)
	}
	k.AddKeyWords(word, word)
}

func (k *KeywordProcessor) AddKeyWords(word, cleanName string) {
	if k.caseSensitive {
		word = strings.ToLower(word)
	}
	k.rootNode.Add(word, cleanName, []rune(word))
}

func (k *KeywordProcessor) FindKeyWords(input string) []string {
	keywordsSet := make(map[string]struct{})
	var keywords []string
	curNode := k.rootNode
	for _, c := range input {
		if k.caseSensitive {
			c = unicode.ToLower(c)
		}
		trie := curNode.GetChar(c)
		if trie != nil {
			curNode = trie
			word := curNode.Get()
			if word != nil {
				keywordsSet[*word] = struct{}{}
			}
		} else {
			curNode = k.rootNode
		}
	}
	for w := range keywordsSet {
		keywords = append(keywords, w)
	}
	return keywords
}

func (k *KeywordProcessor) ReplaceKeyWords(input string) string {
	output := strings.Builder{}
	buffer := strings.Builder{}
	curNode := k.rootNode
	for _, c := range input {
		char := c
		if k.caseSensitive {
			char = unicode.ToLower(c)
		}
		node := curNode.GetChar(char)
		if node != nil {
			curNode = node
			wordPtr := curNode.Get()
			if wordPtr != nil {
				buffer.Reset()
				buffer.WriteString(*wordPtr)
			} else {
				buffer.WriteRune(c)
			}
		} else {
			output.WriteString(buffer.String())
			output.WriteRune(c)
			buffer.Reset()
			curNode = k.rootNode
		}
	}
	if buffer.Len() > 0 {
		output.WriteString(buffer.String())
	}
	return output.String()
}
