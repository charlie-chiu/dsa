204 count prime

很基礎的質數題，題目要求給定一個整數 n ，輸出的結果為小於 n 的質數數量。

最直覺的解法大概就是遍歷所有小於 n 的數，個別判斷是否為質數並記錄質數的數量；
毫不意外的這個解法在 leetcode 上會是 TLE

這裡選用 

Sieve of Eratosthenes (埃氏篩)
https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
據 WIKI 的說明，是一個古老的演算法了，2021很有可能有更快的解法可以使用

todo: 試著以其他演算法來完成這一題