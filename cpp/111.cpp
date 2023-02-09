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
/*
题目：
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
*/
    int minDepthHelper(TreeNode*root){
        // 叶子结点
        if(root->left==nullptr&&root->right==nullptr){
            return 1;
        }
        // 只考虑存在叶子结点
        if(root->left==nullptr){
            return minDepthHelper(root->right)+1;
        }
        if(root->right==nullptr){
            return minDepthHelper(root->left)+1;
        }
        return min(minDepthHelper(root->left),minDepthHelper(root->right))+1;
    }
    int minDepth(TreeNode* root) {
        if(root==nullptr){
            return 0;
        }
        return minDepthHelper(root);
    }
};
