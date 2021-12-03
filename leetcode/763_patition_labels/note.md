763 Partition Labels

給定一字串，在每個字元只能出現於同一個 partition 的前提下，將字串分割為盡可能多的 partition
盡可能多 partition 等同於每個 partition 盡可能的短

--- 

Approach:

先檢視整個 string ，取得每個字元最後出現的位置

接下來從第一個字元開始檢視，預設 partition 的結束位置就在這個字元最後一次出現的位置(盡可能短)
如果在 pointer curr 到達 end 之前，發現中間的字元結束位置比現在的假設更後面，就將結束位置移往該位置
直到 curr 達到 end，記錄這個 partition ，並繼續處理下一個區段。
