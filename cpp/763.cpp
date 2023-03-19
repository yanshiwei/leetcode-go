class Solution {
    /*
    题目：
    给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
返回一个表示每个字符串片段的长度的列表。
    题解：贪心。
    1.同一个字母只能出现在同一个片段，故同一个字母第一次出现和最后一次出现的下标肯定也在同一个位置。
    故需遍历字符串，得到每个字符最后一次出现的下标；
    2.基于贪心思想将字符串划分尽可能多，也就是每个片段尽可能短：
    2.1 从左到右，维护当前片段的开始/结束的下标start/end，初始化0
    2.2 对于每个字母c，得当该字母最后一次出现的下标endc，则该字母最终的片段的结束下标一定不小于endc，故end=max(end,endc)
    2.3 当遍历到下标end时，当前片段结束。当前片段的下标范围是 [start,end]长度为 end−start+1，将当前片段的长度添加到返回值，然后令start=end+1，继续寻找下一个片段。
    3.由于每次取的片段都是符合要求的最短的片段，因此得到的片段数也是最多的。
    */
public:
    vector<int> partitionLabels(string s) {
        // 统计每一个字符最后出现的位置
        int pos[27]={0};
        for(int i=0;i<s.size();i++){
            pos[s[i]-'a']=i;
        }
        vector<int>res;
        int start=0;
        int end=0;
        for(int i=0;i<s.size();i++){
            end=max(end,pos[s[i]-'a']);
            if(i==end){
                res.push_back(end-start+1);
                start=end+1;
            }
        }
        return res;
    }
};
