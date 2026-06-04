// https://leetcode.com/problems/delete-duplicate-folders-in-system/description/

type trie struct {
    root *node
    directories map[string]int
}

func newTrie() trie {
    return trie{
        root: newNode(),
        directories: make(map[string]int),
    }
}

type node struct {
    next map[string]*node
    hash string
}

func newNode() *node {
    return &node{
        next: make(map[string]*node),
        hash: "",
    }
}

func (t *trie) dedup() {
    var dfs func(*node)
    dfs = func(root *node) {
        if len(root.next) == 0 {
            return
        }

        for n := range root.next {
            if k := t.directories[root.next[n].hash]; k > 1 {
                delete(root.next, n)
            } else {
                dfs(root.next[n])
            }
        }
    }

    dfs(t.root)
}

func (t *trie) insert(path []string) {
    root := t.root
    for _, p := range path {
        if root.next[p] == nil {
            root.next[p] = newNode()
        }
        root = root.next[p]
    }
}

func (t trie) serialize() {
    var dfs func(*node, string) string
    dfs = func(root *node, n string) string {
        if len(root.next) == 0 {
            return n
        }

        leaves := []string{}
        for n := range root.next {
            leaves = append(leaves, n)
        }

        sort.Strings(leaves)

        hash := ""
        for _, l := range leaves {
            hash += fmt.Sprintf("(%s)", dfs(root.next[l], l))
        }

        root.hash = hash
        t.directories[hash]++
        
        return n + "/" + hash 
    }

    t.root.hash = dfs(t.root, "")
}

func (t trie) traverse() [][]string {
    paths := [][]string{}

    var dfs func(root *node, current []string)
    dfs = func(root *node, current []string) {
        if len(root.next) == 0 {
            if len(current) > 0 {
                paths = append(paths, slices.Clone(current))
            }
            return
        }

        if len(current) > 0 {
            paths = append(paths, slices.Clone(current))
        }

        for n := range root.next {
            dfs(root.next[n], append(current, n))
        }
    }

    dfs(t.root, []string{})

    return paths
}

func deleteDuplicateFolder(paths [][]string) [][]string {
    trie := newTrie()
    for _, p := range paths {
        trie.insert(p)
    }

    trie.serialize()
    trie.dedup()
    
    return trie.traverse()
}
