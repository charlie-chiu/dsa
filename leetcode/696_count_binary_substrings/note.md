
寫完 647-迴文子字串後接著寫，直覺的用同一思考模式實作，
iterate 字串，對於每個位置往左右檢查
但是有一個不同的狀況，依題目要求，0/1 必須是連續的，故 1100 可以計入， 0101 即不算
所以在每個位置的迴圈中加入了額外的檢查，可以 work 但不太優雅

--- 

leetcode 101 的解法：
計算連續的 0/1 數量，不同的字元只需於前一段的數量比對即可
複雜度下降為 O(N)