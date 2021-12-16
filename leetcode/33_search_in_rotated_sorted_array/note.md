基本概念：
在 middle 的左右兩邊，一定有其中一邊是排好序的，雖然不知道 pivot point 的確切位置，但還是可以判斷 target 位於哪一個區間。

但很遺憾我還無法自己實作出來，致敬 Discuss 完成的。
https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14435/Clever-idea-making-it-simple

值得注意的是，一樣使用了左閉右開的方式來設計區間，在判斷上只使用第 0 ~ m 個元素，仍然沒有用到 r