class Solution {
    /*
    给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
例如：
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...
    */
public:
    string convertToTitle(int columnNumber) {
        string res="";
        while(columnNumber){
            res+=('A'+(columnNumber-1)%26);
            columnNumber=(columnNumber-1)/26;
        }
        reverse(res.begin(),res.end());
        return res;
    }
};
