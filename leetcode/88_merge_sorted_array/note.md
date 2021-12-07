合併兩個排序過的array，特別的是，不需return 新的 array，而是直接加在 nums1 的後面

用三個 pointer 分別代表兩個 array 中目前處理到的值，以及新的值要插入的位置，並妥善處理 edge cases 即可。