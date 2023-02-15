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
    给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
    */
public:
    void pathSumHelper(TreeNode* root, int targetSum,vector<int>tmp,vector<vector<int>> &res){
        if(root==nullptr){
            return;
        }
        tmp.push_back(root->val);
        if(root->left==nullptr&&root->right==nullptr&&root->val==targetSum){
            res.push_back(tmp);
            return;
        }
        pathSumHelper(root->right,targetSum-root->val,tmp,res);
        pathSumHelper(root->left,targetSum-root->val,tmp,res);    
    } 
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        if(root==nullptr){
            return {};
        }
        vector<vector<int>>res;
        vector<int>tmp;
        pathSumHelper(root,targetSum,tmp,res);
        return res;
    }
};
