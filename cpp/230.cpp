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
 /*
 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
 */
class Solution {
public:
    void helper(TreeNode*root, vector<int>& vec){
        if(root==nullptr){
            return;
        }
        helper(root->left,vec);
        vec.push_back(root->val);
        helper(root->right,vec);
    }
    int kthSmallest(TreeNode* root, int k) {
        if(root==nullptr){
            return -1;
        }
        vector<int>res;
        helper(root,res);
        return res[k-1];
    }
};
