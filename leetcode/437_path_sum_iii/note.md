暴力法，先用 dfs (理論上任何 traversal 都可以)

對於每個 node 再以 dfs 的方式找出總如 = target 的 path

遇到一個 wrong answer 出現在以下：
         1
       /   \
     -2    -3
     / \   /
    1   3  2
   /
 -1

target = -1
expected = 4
演算法的輸出為 3 ，經檢查是最左側的path，計算到 1 + -2 = -1 就停止了，但 1 + -2 + 1 + -1 也等於目標的 -1 
修改讓每個 node 的 sum 計算都算到底來修正這個問題