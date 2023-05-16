class Solution {
    /*
    给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。
例如：
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
    */
public:
    int titleToNumber(string columnTitle) {
        int res=0;
        for(int i=0;i<columnTitle.size();i++){
            res=res*26+(columnTitle[i]-'A')+1;
        }
        return res;
    }
};
