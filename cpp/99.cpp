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
    给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
    */
public:
    int idx=0;
    void inOrderGet(TreeNode*root,vector<int>&res){
        if(root==nullptr){
            return;
        }
        inOrderGet(root->left,res);
        res.push_back(root->val);
        inOrderGet(root->right,res);
    }
    void inOrderSet(TreeNode*root,vector<int>&res){
        if(root==nullptr){
            return;
        }
        inOrderSet(root->left,res);
        root->val=res[idx];
        idx++;
        inOrderSet(root->right,res);
    }
    void recoverTree(TreeNode* root) {
        vector<int>res;
        inOrderGet(root,res);
        sort(res.begin(),res.end());
        inOrderSet(root,res);
    }
};
