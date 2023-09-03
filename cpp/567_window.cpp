class Solution {
    /*
    明显的滑动窗口算法，相当给你一个S和一个T，请问你S中是否存在一个子串，包含T中所有字符且不包含其他字符？
    本题移动left缩小窗口的时机是窗口大小大于t.size()时，因为排列嘛，显然长度应该是一样的
    当发现valid == need.size()时，就说明窗口中就是一个合法的排列，所以立即返回true
    类似76题。
    */
public:
    bool checkInclusion(string s1, string s2) {
        unordered_map<char,int>need,windows;
        for(char c : s1){
            need[c]++;
        }
        int left=0;
        int right=0;
        int cnt=0;//s2字符串[left,right]区间中满足s1字符串的元素的个数
        while(right<s2.size()){
            char c=s2[right];
            right++;
            // 进行窗口内数据的一系列更新
            if(need.count(c)){
                windows[c]++;
                if(windows[c]==need[c]){
                    cnt++;
                }
            }
            // 关键点：判断左侧窗口是否要收缩,保证窗口的大小始终和需要查找的字串一致
            while(right-left>=s1.size()){
                if(cnt==need.size()){
                    return true;
                }
                char d=s2[left];
                left++;
                // 进行窗口内数据的一系列更新
                if(need.count(d)){
                    if(windows[d]==need[d]){
                        cnt--;
                    }
                    windows[d]--;
                }
            }
        }
        return false;
    }
};
