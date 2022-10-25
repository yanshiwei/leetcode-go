class Solution {
    /*
    给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
    */
public:
    int longestSubstring(string s, int k) {
        int res=0;
        //滑动窗口
        for(int i=1;i<=26;i++){
            //窗口内 恰好有 i 个不同字母 如ababcb i=3
            int differentCnt=0;//统计当前窗口内不同的字母数量,如ababcb,3
            int beyondCnt=0;//统计当前窗口内已 ≥k 的字母数量,如ababcb,2
            vector<int>visited(26,0);//26个字符出现的次数
            int left=0,right=0;
            while(right<s.size()){
                // 右移动逻辑，只需进行值的统计和更新。 
                // 当 visited[tmp] == 1 表示，某个字母从没有出现到出现，则增加了一个字母，所以differCnt加1;
                // 而当 visited[tmp] == k 则表示刚好有某个字母达到要求的k值，次数beyondCnt则加1
                int tmp=s[right]-'a';
                visited[tmp]++;
                if(visited[tmp]==1){
                    differentCnt++;
                }
                if(visited[tmp]==k){
                    beyondCnt++;
                }
                right++;
                // 左移动逻辑，要使得当前窗口内不同数量的字母不能超过当前遍历的个数i，也就是当 differentCnt > i 要进行左移操作
                while(differentCnt>i){
                    //  左移动处理的过程中则是上面右移动操作的反过程
                    int tmp=s[left]-'a';
                    if(visited[tmp]==1){
                        differentCnt--;
                    }
                    if(visited[tmp]==k){
                        beyondCnt--;
                    }
                    visited[tmp]--;
                    left++;
                }
                // 更新逻辑， 每次我们要的是恰好 i 个不同的字母存在，所以要判断 differCnt == i 
                // 然后这i个不同字母同时也要超过k（才能算连续字串），所以判断 beyondCnt == i 满足这两个条件，我们再进行更新的操作
                if(differentCnt==i&&beyondCnt==i){
                    res=max(res,right-left);//此时right其实已经是在可以满足条件的情况向前多走了一步的，即前面的right++故这里不是right - left + 1
                }
            }
        }
        return res;
    }
};
