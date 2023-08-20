class Solution {
    /*
    题目：
    题解：
    https://leetcode.cn/problems/generate-parentheses/solutions/35947/hui-su-suan-fa-by-liweiwei1419/?envType=study-plan-v2&envId=top-100-liked
    DFS，结论：
    1 当前左右括号都有大于 0个可以使用的时候，才产生分支
    2 产生左分支的时候，只看当前是否还有左括号可以使用
    3 产生右分支的时候，还受到左分支的限制，右边剩余可以使用的括号数量一定得在严格大于左边剩余的数量的时候，才可以产生分支
    4 在左边和右边剩余的括号数都等于 0的时候结束
    */
public:
    void dfs(int left, int right,string curStr,vector<string>&res){
        // 在左边和右边剩余的括号数都等于 0的时候结束
        if(left==0&&right==0){
            res.push_back(curStr);
            return;
        }
        // 剪枝（如图，左括号可以使用的个数严格大于右括号可以使用的个数，才剪枝，注意这个细节）
        if(left>right){
            return;
        }
        // 产生左分支的时候，只看当前是否还有左括号可以使用
        if(left>0){
            dfs(left-1, right, curStr+"(", res);
        }
        if(right>0){
            dfs(left, right-1, curStr+")", res);
        }
    }
    vector<string> generateParenthesis(int n) {
        vector<string>res;
        if(n<1){
            return res;
        }
        dfs(n,n,"",res);
        return res;
    }
};
