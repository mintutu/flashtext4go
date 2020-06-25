package flashtext

import "strings"

type KeywordTrieNode struct {
	keyword *string
	cleanName *string
	children map[rune]*KeywordTrieNode
}

func (k *KeywordTrieNode) ContainsChar(char rune) bool{
	_, ok := k.children[char]
	return ok
}

func (k *KeywordTrieNode) ContainsWord(word string) bool {
	return k.keyword != nil && *k.keyword == word
}

func (k *KeywordTrieNode) IsEmpty() bool {
	return len(k.children) == 0
}

func (k *KeywordTrieNode) GetChar(char rune) *KeywordTrieNode {
	return k.children[char]
}

func (k *KeywordTrieNode) Get() *string {
	if k.cleanName != nil {
		return k.cleanName
	}
	if k.keyword != nil {
		return k.keyword
	}
	return nil
}

func (k *KeywordTrieNode) Add(word string, cleanName string, characters []rune) *KeywordTrieNode {
	if len(characters) == 0 {
		return &KeywordTrieNode{
			keyword: &word,
			cleanName: &cleanName,
			children: make(map[rune]*KeywordTrieNode),
		}
	}
	head := characters[0]
	node := k.GetChar(head)
	if node == nil {
		node = &KeywordTrieNode{}
	}
	if k.children == nil {
		k.children = make(map[rune]*KeywordTrieNode)
	}
	k.children[head] = node.Add(word, cleanName, characters[1:])
	return k
}

func (k *KeywordTrieNode) ToString() string {
	return k.toStringHelper("")
}

func (k *KeywordTrieNode) toStringHelper(pad string) string {
	var b strings.Builder
	wordPtr := k.Get()
	if wordPtr != nil {
		b.WriteString(*wordPtr)
	}
	b.WriteString("\n")
	if len(k.children) > 0 {
		for key, value := range k.children {
			b.WriteString(pad)
			b.WriteRune(key)
			b.WriteString(":")
			b.WriteString(value.toStringHelper(pad + " "))
		}
	}
	return b.String()
}




