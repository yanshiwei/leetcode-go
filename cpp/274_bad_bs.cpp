class Solution {
    // 二分https://leetcode.cn/problems/h-index/solutions/870465/er-fen-cai-lun-wen-pian-shu-java-by-liwe-zoh7/
    /*
    类似287
    我们知道，一个数组在某个属性上有单调性，就可以进行二分查找
（最明显就是排序数组的数组值）。对于本题来说还有什么单调性呢。我们尝试探索一下
    例如arr=[1 3 5 7 10 13]
    引用次数>=1，文章篇数=6>1
    引用次数>=3，文章篇数=5>3
    引用次数>=5，文章篇数=4<5
    引用次数>=7，文章篇数=3<7
    引用次数>=10，文章篇数=2<10
    引用次数>=13，文章篇数=1<13
    所以引用次数越大越大，超过该引用的文章篇数就越少，我们记这个属性为cnt,则现在我们
知道，对于【0, 1，2，3,...n】这个数组值升序数组来说，数组的cnt属性值是降序的
    对于目标target，它的cnt属性有什么特点呢。引用次数>=target，文章篇数>=target
    if 引用次数<=target,则引用次数的cnt属性值也就是文章篇数>=引用次数
    if 引用次数>target,则引用次数的cnt属性值也就是文章篇数<引用次数
    根据这个结论，我们就可以进行二分查找
    */
public:
    int hIndex(vector<int>& citations) {
        //1 3 | 5 7 10 13
        int n=citations.size();
        int left=0;//最差情况下，所有的论文被引用次数都为 0
        int right=n;//最好情况下，所有的论文的引用次数 >= 总共论文篇数
        while(left<right){
            // 若等于，left==right。走到1步骤时候会循环
            int mid=left+(right-left)/2+1;//这里+1，否则[1]会循环
            if(check(citations,mid)>=mid){
                // target在 in the right [mid,right]
                left=mid;
            }else{
                // target在 in the left [left,mid-1]
                right=mid-1;
            }
        }
        return left;
    }
private:
    int check(vector<int>&citations,int mid){
        int cnt=0;
        for(auto &n:citations){
            if(n>=mid){
                cnt++;
            }
        }
        return cnt;
    }
};
