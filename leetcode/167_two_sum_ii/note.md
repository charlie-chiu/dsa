這題有一個很關鍵的條件是 numbers 是排序過的，並保證一定會有，也只有一個結果
使用兩個 pointer 分別從最左及最右判斷和，直到找到 target 即可

奇怪的點是題目給的是 1-indexed array ，但是 submit 後發現測資給的好像還是 0-indexed 。