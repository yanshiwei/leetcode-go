class Solution {
    /*
    下一个要更大，则很显然要增大数字；又要求原地交换，故有两种思想：
    1 从左比较寻找大数字，如 4312 先从个位开始，先比较十位的 若更大则交换 否则继续找这里变成了4213 直接变小了不符合题意 抛弃
    2 从右比较寻找大数字，如4312 个位数右边空故从十位开始比较。若更大则交换如这里变成了4321 变大了 符合题意 继续分析
    假如一直向左边寻找到底都没有找到更大的 说明数字已经是递增如1234 此时按照题意直接逆序即可；假如找到了为了保证交换后数值最小化，需找到第一个大于该值的数字
    比如1548751，第一个存在右边大值的是4，比4大的第一个数字是5。交换后1558741较1548751还不是最小的大值，最小的是1551487 也就是5后面的数字逆序
    注意此时5后面的顺序天然逆序 直接reverse
    */
public:
    void nextPermutation(vector<int>& nums) {
        if(nums.size()<1){
            return;
        }
        // 从十位开始寻找右侧大于该值的
        int i=nums.size()-2;
        while(i>=0&&nums[i+1]<=nums[i]){
            i--;
        }
        // 到左侧 还是没有找到比它大的 说明序列是987654321这种类型 按照题意 直接reverse
        if(i<0){
            reverse(nums,0);
            return;
        }
        //若找到了这样的数字，（此时i右侧都是逆序排序好的，有大于nums[i]也有小于的）则需要找到第一个大于的以保证最终交换后不会超过很大
        int j=nums.size()-1;
        while(j>=0&&nums[j]<=nums[i]){
            j--;
        }
        // 交换
        swap(nums,i,j);
        // 交换后还不一定是最小的超过当前值的数字比如1548751，第一个存在右边大值的是4，比4大的第一个数字是5交换后
        // 1558741较1548751还不是最小的大值，最小的是1551487 也就是5后面的数字逆序
        // 注意此时5后面的顺序天然逆序 直接reverse
        reverse(nums,i+1);
        return;
    }
    void swap(vector<int>& nums,int i,int j){
        int tmp=nums[i];
        nums[i]=nums[j];
        nums[j]=tmp;
    }
    void reverse(vector<int>& nums,int start){
        int i=start;
        int j=nums.size()-1;
        while(i<j){
            swap(nums,i,j);
            i++;
            j--;
        }
    }
};
