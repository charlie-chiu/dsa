判斷一個 Binary Tree 是否 balanced

balance 的定義為：對於任一個 node，左右 subtree 的 height 差異不超過 1

# first try
因為需要知道 subtree 的高度，用了 104 max depth 的程式
簡單的粗暴的對每個節點判斷一次

# leetcode 101
依照題目，所有的node 都必須要是 balanced 才行，原本的方式會對已經不平衡的 subtree 重複判斷，
故修改 func maxDepth，不平衡時直接 return -1  
理論上應該會比較快才對，但在 leetcode 上的速度卻比較慢，不曉得為何