class Solution {
    /*
     题目：
    此题求两个有序数组的中位数，并且限制时间复杂度为O(log (m+n))，所以自然想到要用二分法求解
    题解：
    参考：
    1 https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/485550/zong-he-bai-jia-ti-jie-zong-jie-zui-qing-xi-yi-don/?envType=study-plan-v2&envId=top-100-liked
    2 is key， https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/8999/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-2/?envType=study-plan-v2&envId=top-100-liked
    中位数：如果某个有序数组长度是奇数，那么其中位数就是最中间那个，如果是偶数，那么就是最中间两个数字的平均值。这里对于两个有序数组也是一样的，假设两个有序数组的长度分别为m和n，由于两个数组长度之和m+n的奇偶不确定，因此需要分情况来讨论，对于奇数的情况，直接找到最中间的数即可，偶数的话需要求最中间两个数的平均值。
使用一个小trick，可以避免讨论奇偶：
我们分别找第 (m+n+1)/2个数，和(m+n+2)/2个数，然后求其平均值即可，这对奇偶数均适用。假如 m+n 为奇数的话，那么其实 (m+n+1) / 2 和 (m+n+2) / 2 的值相等，相当于两个相同的数字相加再除以2，还是其本身。
题目是求中位数，其实就是求第 k 小数的一种特殊情况，
    所以我们先广泛一点，求两个有序数组的第k个小元素，最终中位数就是k=(m+n)/2时候的特殊情况。
    取两个数组中的第k/2个元素（midVal1和midVal2）进行比较，
    如果midVal1 < midVal2，则说明数组1中的前k/2个元素不可能成为第k个元素的候选，所以将数组1中的前k/2个元素去掉，去掉后的nums1作为新数组和数组2求第k-k/2小的元素，注意这里因为我们把前k/2个元素去掉了，所以相应的k值也应该减少k/2。midVal1 > midVal2的情况亦然。
    不断进行处理，直到两种情况：
    1 当k=1时候，相当于求最小值，我们只要比较nums1和nums2的起始位置
    2 一个数组为空，则直接蜕变成在另一个数组寻找
    注意的问题：由于两个数组的长度不定，所以有可能某个数组元素数不足k/2，所以我们需要先检查一下，数组中到底存不存在第k/2个数字，若没有，则直接去掉这个数组全部元素，下一次比较时候这个数组为空
    */
public:
    // 求两个有序数组的第k个数
    double getKth(vector<int>& nums1,int start1,int end1, vector<int>& nums2,int start2,int end2, int k){
        int len1=end1-start1+1;
        int len2=end2-start2+1;
        //让 len1 的长度小于 len2，这样就能保证如果有数组空了，一定是 len1 方便后续处理
        if(len1>len2){
            return getKth(nums2,start2,end2,nums1,start1,end1,k);
        }
        if(len1==0){
            // 一个数组为空，则kth为另一个数组的第k个
            return nums2[start2+k-1];
        }
        if(k==1){
            // 两个有序数组的最小1个数，直接看两个数组的开头两个大小
            return min(nums1[start1],nums2[start2]);
        }
        // 为了防止数组长度小于 k/2，所以每次比较 min(k/2，len(数组) 对应的数字，把小的那个对应的数组的数字排除
        int i=start1+min(len1,k/2)-1;
        int j=start2+min(len2,k/2)-1;
        // 去除小的部分，只看大的部分
        if(nums1[i]>nums2[j]){
            // 数组2中的前k/2个元素不可能成为第k个元素的候选，所以将数组2中的前k/2个元素去掉
            // 去除的元素个数，也就是nums2[i]部分去除的个数
            int cnt=j-start2+1;
            return getKth(nums1,start1,end1,nums2,j+1,end2,k-cnt);
        }else{
            // 数组1中的前k/2个元素不可能成为第k个元素的候选，所以将数组1中的前k/2个元素去掉
            // 去除的元素个数，也就是nums1[i]部分去除的个数
            int cnt=i-start1+1;
            return getKth(nums1,i+1,end1,nums2,start2,end2,k-cnt);
        }
    }  
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int n=nums1.size();
        int m=nums2.size();
        int left=(n+m+1)/2;
        int right=(n+m+2)/2;
        //将偶数和奇数的情况合并，如果是奇数，会求两次同样的 k
        return (getKth(nums1,0,n-1,nums2,0,m-1,left)+getKth(nums1,0,n-1,nums2,0,m-1,right))*0.5;
    }
};
