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
    // dfs,自顶向下递归获得路径，https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
public:
    void dfs(TreeNode*root, string tmp){
        if(root==nullptr){
            return;
        }
        tmp+=(to_string(root->val));
        if(root->left==nullptr&&root->right==nullptr){
            res.push_back(tmp);
            return;
        }else{
           tmp+="->" ;
        }
        dfs(root->left,tmp);
        dfs(root->right,tmp);
    }
    vector<string> binaryTreePaths(TreeNode* root) {
        string tmp="";
        dfs(root,tmp);
        return res;
    }
private:
    vector<string>res;
};
