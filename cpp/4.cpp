class Solution {
    /*
    给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
    */
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int l1=nums1.size();
        int l2=nums2.size();
        vector<int>tmp(l1+l2);
        int idx1=0;
        int idx2=0;
        int idx=0;
        while(idx1<l1&&idx2<l2){
            if(nums1[idx1]<=nums2[idx2]){
                tmp[idx++]=nums1[idx1++];
            }else{
                tmp[idx++]=nums2[idx2++];
            }
        }
        while(idx1<l1){
            tmp[idx++]=nums1[idx1++];
        }
        while(idx2<l2){
            tmp[idx++]=nums2[idx2++];
        }
        if((tmp.size()&1)==1){
            return double(tmp[tmp.size()/2]);
        }else{
            return double(tmp[tmp.size()/2-1]+tmp[tmp.size()/2])/2;
        }
    }
};
