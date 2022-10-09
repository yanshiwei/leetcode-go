class Solution {
    /*
    题目：
    给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。
如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。
请注意，答案不一定是 arr 中的数字。
    思路：二分法。严格大于 value 的元素变成 value,value 越大，「转变后的数组」的和越大，这是「单调性」故可以引入二分法。value 取值从0到max of arr 
    问题关键是怎么衡量最接近，最接近的情况是：选定了一个 value 求和以后，恰恰好等于 target。不过更有可能出现的情况是：value 选得小了，数组和小于target，而 value 再+1后，数组和又不小于target。解决方案是：把边界的上下方的可能的 value 值（一共就两个）都拿出来进行一次比较即可。
    完整思路：https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target/solution/er-fen-cha-zhao-by-liweiwei1419-2/
    */
public:
    int findBestValue(vector<int>& arr, int target) {
        // 二分法 value from 0 to max of arr to try
        int left=0;
        int right=getMax(arr);
        while(left<right){
            int mid=left+(right-left)/2;
            int sum=calSum(arr,mid);
            if(sum<target){
                // 严格小于
                left=mid+1;
            }else{
                // 大于等于
                right=mid;
            }
        }
        //把边界的上下方的可能的 value 值（一共就两个）都拿出来进行一次比较
        int sum1=calSum(arr,left-1);// <
        int sum2=calSum(arr,left);//>=
        if(target-sum1<=sum2-target){
            //如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。 故这里要保留等于号，这样同样差距时选择更小的
            return left-1;
        }
        return left;
    }
private:
    int getMax(vector<int>& arr){
        if(arr.size()<1){
            return 0;
        }
        int res=arr[0];
        for(int i=1;i<arr.size();i++){
            if(res<arr[i]){
                res=arr[i];
            }
        }
        return res;
    }
    int calSum(vector<int>& arr,int value){
        int sum=0;
        for(auto &d:arr){
            if(d>=value){
                sum+=value;
            }else{
                sum+=d;
            }
        }
        return sum;
    }
};
