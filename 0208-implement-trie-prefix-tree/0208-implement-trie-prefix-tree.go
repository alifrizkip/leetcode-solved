type TrieNode struct {
    IsEndOfWord bool
    Children map[rune]*TrieNode
}

type Trie struct {
    Root *TrieNode
}


func Constructor() Trie {
    return Trie{
        Root: &TrieNode{
            Children: make(map[rune]*TrieNode),
        },
    }
}


func (this *Trie) Insert(word string)  {
    node := this.Root
    for _, char := range word {
        if _, ok := node.Children[char]; !ok {
            node.Children[char] = &TrieNode{
                Children: make(map[rune]*TrieNode),
            }
        }
        node = node.Children[char]
    }
    node.IsEndOfWord = true
}


func (this *Trie) Search(word string) bool {
    node := this.Root
    for _, char := range word {
        if _, ok := node.Children[char]; !ok {
            return false
        }
        node = node.Children[char]
    }
    return node.IsEndOfWord
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this.Root
    for _, char := range prefix {
        if _, ok := node.Children[char]; !ok {
            return false
        }
        node = node.Children[char]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */