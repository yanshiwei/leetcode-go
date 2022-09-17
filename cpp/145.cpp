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
    /*
    给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
    */
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int>res;
        if(root==nullptr){
            return res;
        }
        postorderTraversalCore(root,res);
        return res;
    }
private:
    void postorderTraversalCore(TreeNode*root,vector<int>&res){
        if(root==nullptr){
            return;
        }
        postorderTraversalCore(root->left,res);
        postorderTraversalCore(root->right,res);
        res.push_back(root->val);
    }
};
