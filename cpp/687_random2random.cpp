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
    int longestUnivaluePath(TreeNode* root) {
        if(root==nullptr){
            return 0;
        }
        dfs(root);
        return res;
    }
private:
    //root结点能提供的最长路径
    //任意点到任意点，https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
    int dfs(TreeNode*root){
        if(root==nullptr){
            return 0;
        }
        int left=dfs(root->left);
        int right=dfs(root->right);
        // 如果存在左子节点和根节点同值，更新左最长路径;
        // 否则左最长路径为0
        if(root->left&&root->left->val==root->val){
            left++;
        }else{
            left=0;
        }
        if(root->right&&root->right->val==root->val){
            right++;
        }else{
            right=0;
        }
        res=max(res,left+right);
        return max(left,right);
    }
    int res=0;
};
