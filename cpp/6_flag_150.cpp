class Solution {
public:
    string convert(string s, int numRows) {
        if(numRows==1){
            return s;
        }
        vector<string>rows(numRows,"");
        int curRow=0;
        bool goingDown=false;
        for(char c : s){
            rows[curRow].push_back(c);
            // change direct
            if(curRow==0||curRow==numRows-1){
                // 最开始和最末，肯定转向
                goingDown=!goingDown;
            }
            if(goingDown){
                curRow++;
            }else{
                curRow--;
            }
        }
        string res="";
        for(int i=0;i<rows.size();i++){
            res+=rows[i];
        }
        return res;
    }
};
