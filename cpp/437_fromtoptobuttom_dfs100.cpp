/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
    // dfs，自顶向下路径和，https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
    // 是否要双重递归(即调用根节点的dfs函数后，继续调用根左右节点的pathsum函数)：看题目要不要求从根节点开始的，还是从任意节点开始
public:
    int pathSum(TreeNode* root, int targetSum) {
        if(root==nullptr){
            return 0;
        }
        dfs(root,targetSum);//以root为起始点查找路径
        //题目不要求从根节点开始，故双重递归
        pathSum(root->left,targetSum);
        pathSum(root->right,targetSum);
        return res;
    }
private:
    // long target, use int target may be overflow
    void dfs(TreeNode*root,long targetSum){
        if(root==nullptr){
            return;
        }
        targetSum-=root->val;
        // https://leetcode.cn/problems/path-sum-ii/description/
        // 相比较113，这里不需要在叶子节点结束所以直接判断剩余数值
        // 注意不要return,因为不要求到叶节点结束,所以一条路径下面还可能有另一条
        if(targetSum==0){
            res++;
        }
        dfs(root->left,targetSum);
        dfs(root->right,targetSum);
    }
    int res=0;
};
