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
public:
    bool hasPathSumHelper(TreeNode*root,int targetSum){
        // 叶子结点
        if(root->left==nullptr&&root->right==nullptr){
            return targetSum==root->val;
        }
        // 只有单侧子节点，只需要遍历该侧，因为叶子结点只会在该侧
        if(root->left==nullptr){
            return hasPathSumHelper(root->right,targetSum-root->val);
        }
        if(root->right==nullptr){
            return hasPathSumHelper(root->left,targetSum-root->val);
        }   
        return hasPathSumHelper(root->left,targetSum-root->val)||hasPathSumHelper(root->right,targetSum-root->val);
    }
    bool hasPathSum(TreeNode* root, int targetSum) {
        if(root==nullptr){
            return false;
        }
        return hasPathSumHelper(root,targetSum);
    }
};
