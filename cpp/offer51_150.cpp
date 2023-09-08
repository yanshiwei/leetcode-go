class Solution {
    /*
    https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/solutions/622496/jian-zhi-offer-51-shu-zu-zhong-de-ni-xu-pvn2h/
    参考：归并排序：https://leetcode.cn/problems/sort-an-array/submissions/
    */
public:
    void merge(vector<int>&nums,int left,int mid, int right,vector<int>& tmp){
        // 2段合并到tmp
        int i=left;
        int j=mid+1;
        int k=0;
        while(i<=mid&&j<=right){
            if(nums[i]<=nums[j]){
                tmp[k++]=nums[i++];
            }else{
                //当左边数组的大与右边数组的元素时，就对当前元素以及后面的元素的个数进行统计，
                // 此时这个数就是，逆序数
                cnt+=(mid-i+1);
                tmp[k++]=nums[j++];
            }
        }
        while(i<=mid){
            tmp[k++]=nums[i++];
        }
        while(j<=right){
            tmp[k++]=nums[j++];
        }
        //将新数组中的元素，覆盖nums旧数组中的元素。
        //此时数组的元素已经是有序的
        for(i=left;i<=right;i++){
            nums[i]=tmp[i-left];
        }
    }
    void mergeSort(vector<int>&nums,int left,int right, vector<int>& res){
        //当只有一个节点的时候，直接返回，退出递归
        if(left>=right){
            return;
        }
        int mid=(right-left)/2+left;
        // 左拆分
        mergeSort(nums,left,mid,res);
        // 左拆分
        mergeSort(nums,mid+1,right,res);
        // 合并
        merge(nums,left,mid,right,res);
    }
    int cnt=0;
    int reversePairs(vector<int>& nums) {
        cnt=0;
        vector <int> ans(nums.size());
        mergeSort(nums,0,nums.size()-1,ans);
        return cnt;
    }
};
