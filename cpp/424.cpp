class Solution {
    /*
    给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
在执行上述操作后，返回包含相同字母的最长子字符串的长度。
    */
public:
    int characterReplacement(string s, int k) {
        // 枚举字符串，每一个位置作为右侧节点，然后找到最远的左侧节点位置
        // 该区间满足：除了次数最的那一类字符外，剩余字符出现的次数不超过k次（这样可以k次修改）
        // 双指针。每次右指针右移，如果区间仍然满足条件，那么左指针不移动，否则左指针至多右移一格，保证区间长度不减小。
        // 向右移动这样的操作会导致部分区间不符合条件，即该区间内非最长重复字符超过了 k 个
        vector<int>num(26);//字符串中仅包含大写字母，我们可以使用一个长度为 262626 的数组维护每一个字符的出现次数
        int left=0;
        int right=0;
        int maxCnt=0;// 重复字符出现次数的
        while(right<s.size()){
            int tmp=s[right]-'A';
            num[tmp]++;
            if(maxCnt<num[tmp]){
                maxCnt=num[tmp];
            }
            //使用该最大值计算出区间内非最长重复字符的数量,以此判断左指针是否需要右移即可
            if(right-left+1-maxCnt>k){
                // 区间不满足条件
                num[s[left]-'A']--;
                left++;
            }
            right++;
        }
        //右指针移动到尽头，right - left在整个过程是非递减的。只要right 的值加进去不满足条件，left和right就一起右滑，因为长度小于right - left的区间就没必要考虑了，所以right - left一直保持为当前的最大值
        return right-left;//因为right多走一步，结果为(right-1)-left+1==right-left
    }
};
