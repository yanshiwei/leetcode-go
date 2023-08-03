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
    // dfs，自顶到下求路径和，https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
public:
    bool dfs(TreeNode*root,int targetSum){
        if(root==nullptr){
            return false;
        }
        if(root->left==nullptr&&root->right==nullptr){
            return targetSum==root->val;
        }
        bool left=dfs(root->left,targetSum-root->val);
        if (left){
            return true;
        }
        bool right=dfs(root->right,targetSum-root->val);
        return right;
    }
    bool hasPathSum(TreeNode* root, int targetSum) {
        if(root==nullptr){
            return false;
        }
        return dfs(root,targetSum);
    }
};
