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
// dfs，自顶向下路径和，https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
    void dfs(TreeNode* root, int targetSum,vector<int>tmp){
        if(root==nullptr){
            return;
        }
        targetSum-=root->val;
        tmp.push_back(root->val);
        if(root->left==nullptr&&root->right==nullptr&&targetSum==0){
            res.push_back(tmp);
            return;
        }
        dfs(root->left,targetSum,tmp);
        dfs(root->right,targetSum,tmp);
    }
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        if(root==nullptr){
            return {};
        }
        vector<int>tmp;
        dfs(root,targetSum,tmp);
        return res;
    }
private:
    vector<vector<int>> res;
};
