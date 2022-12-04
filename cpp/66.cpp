class Solution {
    //0 <= digits[i] <= 9
public:
    vector<int> plusOne(vector<int>& digits) {
        if(digits.size()<1){
            return {};
        }
        for(int i=digits.size()-1;i>=0;i--){
            digits[i]++;
            if(digits[i]/10==0){
                // no carry
                return digits;
            }
            digits[i]=0;
        }
        // 还能走到此处 说明到最高位还有进位
        digits[0]=1;
        digits.push_back(0);
        return digits;
    }
};
