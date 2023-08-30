class Solution {
    /*
    如果选择站点 i 作为起点「恰好」无法走到站点 j，那么 i 和 j 中间的任意站点 k 都不可能作为起点。
    证明：
    假设 tank 记录当前油箱中的油量，如果从站点 i 出发（tank = 0），走到 j 时恰好出现 tank < 0 的情况，那说明走到 i, j 之间的任意站点 k 时都满足 tank > 0；如果把 k 作为起点的话，相当于在站点 k 时 tank = 0，那走到 j 时必然有 tank < 0，也就是说 k 肯定不能是起点。因为从 i 出发走到 k 好歹 tank > 0，都无法达到 j，现在你还让 tank = 0 了，那更不可能走到 j 了。
    https://leetcode.cn/problems/gas-station/solutions/2074636/dang-lao-si-ji-xue-hui-liao-tan-xin-suan-8s2g/?envType=study-plan-v2&envId=top-interview-150
    场景就被抽象成了一个环形数组，数组中的第i个元素就是gas[i] - cost[i] ● 有了这个环形数组，我们需要判断这个环形数组中是否能够找到一个起点start，使得从这个起点开始的累加和一直大于等于 0。
    https://mp.weixin.qq.com/s/ab8j3zCsmVn_-kmrUjcZyA
    有趣的解法：
    有一个环形路上有n个站点； 每个站点都有一个好人或一个坏人； 好人会给你钱，坏人会收你一定的过路费，如果你带的钱不够付过路费，坏人会跳起来把你砍死； 问：从哪个站点出发，能绕一圈活着回到出发点?

首先考虑一种情况：如果全部好人给你 的钱加起来 小于 坏人收的过路费之和，那么总有一次你的钱不够付过路费，你的结局注定会被砍死。

假如你随机选一点 start 出发，那么你肯定会选一个有好人的站点开始，因为开始的时候你没有钱，遇到坏人只能被砍死；

现在你在start出发，走到了某个站点end，被end站点的坏人砍死了，说明你在 [start, end) 存的钱不够付 end点坏人的过路费，因为start站点是个好人，所以在 (start, end) 里任何一点出发，你存的钱会比现在还少，还是会被end站点的坏人砍死；

于是你重新读档，聪明的选择从 end+1点出发，继续你悲壮的征程； 终于有一天，你发现自己走到了尽头（下标是n-1)的站点而没有被砍死； 此时你犹豫了一下，那我继续往前走，身上的钱够不够你继续走到出发点Start?

当然可以，因为开始已经判断过，好人给你的钱数是大于等于坏人要的过路费的，你现在攒的钱完全可以应付 [0, start) 这一段坏人向你收的过路费。所以无需再走直接结束 这时候你的嘴角微微上扬，眼眶微微湿润，因为你已经知道这个世界的终极奥秘：Start就是这个问题的答案。
    */
public:
    int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
        int n=gas.size();
        int sum=0;
        for(int i=0;i<n;i++){
            sum+=(gas[i]-cost[i]);
        }
        if(sum<0){
            // 循环数组，总油量小于总的消耗，无解
            return -1;
        }
        // 每个加油站的剩余量
        int tank=0;
         // 记录起点, 循环数组任意位置都可以，这里选择0
        int start=0;
        for(int i=0;i<n;i++){
            tank+=(gas[i]-cost[i]);
            if(tank<0){
                // 无法从 start 到达 i 
                // 则[start,i]中间任意结点都不能到达i故直接在i之后寻找
                tank=0;
                start=i+1;
            }
        }
        return start;
    }
};
