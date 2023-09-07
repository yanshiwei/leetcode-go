class Solution {
    /*
    1 假如只有一类任务如AAA,n=2，则处理流程：A 休闲 A 休闲 A
    除了最后一个A，其他除了都是任务处理时间1+间隔时间，最后一个A只需要处理时间不需要冷却故
    t=(任务数-1)*（1+n）+1
    2 对于有两种任务的，只要任务A的数量少于任务B，那么A总可以穿擦在B的间隔中完成，故只需要看出现次数最多的任务的耗时
    3 若存在多个任务最多，如AAABBB，n=2，
    则处理流程：A->B-> ->A->B-> ->A->B 结果是8
    t=(最多的那个任务的任务数-1)*（1+n）+任务数量等于最大的任务数的数量
    4 异常情况
    AAABBBCCCDD n=2
     按照公式算出来结果是9，但实际上数组的长度都大于9，所以是特例，这种情况只需要返回数组长度就行了
    至于为什么会出现这样的情况，我们研究一下：
    他的一种可能结果是
    ABCDABCDABC,他不需要任何等待，就可以执行下一个任务，因此这种情况下需要返回数组长度
    */
public:
    int leastInterval(vector<char>& tasks, int n) {
        unordered_map<char,int>fre;
        for(int i=0;i<tasks.size();i++){
            fre[tasks[i]]++;
        }
        int maxCnt=0;
        for(auto &kv:fre){
            maxCnt=max(maxCnt,kv.second);
        }
        int res=(maxCnt-1)*(1+n);
        for(auto &kv:fre){
            if(kv.second==maxCnt){
                res++;
            }
        }
        // 异常处理
        return res<tasks.size()?tasks.size():res;
    }
};
