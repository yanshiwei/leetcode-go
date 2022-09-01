class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int cnt=0;//当前位置前面有几个与val值相同的元素
        for(int i=0;i<nums.size();i++){
            if (nums[i]==val){
                cnt++;
            }else{
                nums[i-cnt]=nums[i];//当前位置元素需要向前移动多少位
            }
        }
        return nums.size()-cnt;
    }
};
