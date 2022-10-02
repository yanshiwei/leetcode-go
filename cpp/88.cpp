class Solution {
    /*
    给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列
    */
public:
    // 从后往前合并数组，每一次找最大的元素，不会造成数组元素被覆盖的问题
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int i=m-1;
        int j=n-1;
        int k=nums1.size()-1;
        while(i>=0||j>=0){
            if(j<0&&i>=0){
                nums1[k--]=nums1[i--];
            }
            if(i<0&&j>=0){
                nums1[k--]=nums2[j--];
            }
            if(i>=0&&j>=0){
                nums1[k--]=(nums1[i]>nums2[j])?nums1[i--]:nums2[j--];
            }
        }
    }
};
