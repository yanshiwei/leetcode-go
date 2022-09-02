class Solution {
public:
    /*
   给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
*/
    string largestNumber(vector<int>& nums) {
        vector<string>cpy;
        for(int i=0;i<nums.size();i++){
            cpy.push_back(to_string(nums[i]));
        }
        //自定义排序 A+B < B+A
        sort(cpy.begin(),cpy.end(),[](const string&a,const string&b){
            return a+b>b+a;
        });
        // 根据排序不会出现比0还小的在0后面，所以第一个为零就是全零
        if (cpy[0]=="0"){
            return "0";
        }
        string res="";
        // 拼接
        for(int i=0;i<cpy.size();i++){
            res+=cpy[i];
        }
        return res;
    }
};
