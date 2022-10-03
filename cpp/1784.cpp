class Solution {
public:
    bool checkOnesSegment(string s) {
        bool hasOne=false;
        for(auto &w:s){
            if(w=='0'){
                hasOne=true;
            }else if(hasOne==false&&w=='1'){
                continue;
            }else if(hasOne&&w=='1'){
                return false;
            }
        }
        return true;
    }
};
