class Solution {
    /*
    题目：
    假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
    这道题给了我们一个队列，队列中的每个元素是一个 pair，分别为身高和前面身高不低于当前身高的人的个数，让我们重新排列队列，使得每个 pair 的第二个参数都满足题意
    题解：
    首先将身高按从大到小排列，同身高的按k值大小，从小到大排列
    例如： people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
    先按照身高排序，同身高的按k值大小
    1 people = [[7,0],[7,1],[6,1],[5,0],[5,2],[4,4]]
    2. 定一个新数组
        插入第一个，[[7,0]]
        插入第二个，[[7,0],[7,1]]
        插入第三个，第三个[6,1]要求前面有一个不比它小的因此直接append肯定不行，按照要求应该插入到第1个位置:[[7,0],[6,1],[7,1]]
        插入第四个，类似按照要求，[[5,0],[7,0],[6,1],[7,1]]
        ...
        可以看到这样排序后，可以保证身高矮的加入到前面也不会影响前面的k值，例如[6,1]插入时依然不影响[7,1],也就是高个子眼里是看不到矮个子的，不受矮个子影响。
    */
public:
    vector<vector<int>> reconstructQueue(vector<vector<int>>& people) {
        if(people.size()<1){
            return {};
        }
        sort(people.begin(),people.end(),[](vector<int>&a,vector<int>&b){
            return a[0]>b[0]||(a[0]==b[0]&&a[1]<b[1]);
        });
        vector<vector<int>> res;
        for(auto p:people){
            if(res.size()>p[1]){
                // 不满足ki这个要求
                res.insert(res.begin()+p[1],p);
            }else{
                // 满足ki这个要求，至今插入
                res.push_back(p);
            }
        }
        return res;
    }
};
