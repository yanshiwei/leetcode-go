/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int pathSum(TreeNode* root, int sum) {
        if(root==nullptr){
            return 0;
        }
        //以root为起始点查找路径
        dfs(root,sum);
        //题目不要求从根节点开始，故双重递归
        pathSum(root->left,sum);
        pathSum(root->right,sum);
        return res;
    }
private:
    // long 避免溢出
    void dfs(TreeNode* root, long sum){
        if(root==nullptr){
            return;
        }
        sum-=root->val;
        // https://leetcode.cn/problems/path-sum-ii/description/
        // 相比较113，这里不需要在叶子节点结束所以直接判断剩余数值
        // 注意不要return,因为不要求到叶节点结束,所以一条路径下面还可能有另一条
        if(sum==0){
            res++;
        }
        dfs(root->left,sum);
        dfs(root->right,sum);
    }
    int res=0;
};
