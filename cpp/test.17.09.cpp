class Solution {
public:
/*
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21
    暴力
    bool isMagic(int num){
        while(num%3==0){
            num/=3;
        }
        while(num%5==0){
            num/=5;
        }
        while(num%7==0){
            num/=7;
        }
        return (num==1)?true:false;
    }
    int getKthMagicNumber(int k) {
        if(k<1){
            return 0;
        }
        int idx=0;
        int num=0;
        while(idx<k){
            num++;
            if(isMagic(num)){
                idx++;
            }
        }
        return num;
    }
    */
    // 从小到大的第 k 个数，可以使用最小堆,第 k次从最小堆中取出的元素即为第 k 个数
    // 每次取出堆顶元素 x，则 x是堆中最小的数，由于 3x,5x,7x 也是符合要求的数，因此将 3x,5x,7x 加入堆。
    // 上述做法会导致堆中出现重复元素的情况。为了避免重复元素，可以使用哈希集合去重，避免相同元素多次加入堆
    // 注意中间结果long 避免溢出
    int getKthMagicNumber(int k){
        unordered_set<int>se;// erase insert empty size find count
        priority_queue<long,vector<long>,greater<long>>q;// small heap, push pop top empty size swap
        // 1 init by 1
        se.insert(1L);// 1 of long int
        vector<int> lists={3,5,7};
        q.push(1L);
        int res=0;
        for(int i=0;i<k;i++){
            auto cur=q.top();
            q.pop(); // 每次pop当前最小值cur
            res=int(cur);
            // 3x,5x,7x 也是符合要求的数，因此将 3x,5x,7x 加入堆。
            for(int j=0;j<lists.size();j++){
                long next=long(cur)*lists[j];
                // 没有重复
                if(se.count(next)==false){
                    se.insert(next);
                    q.push(next);
                }
            }
        }
        return res;
    }
};
