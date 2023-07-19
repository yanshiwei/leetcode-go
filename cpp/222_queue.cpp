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
    // 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
public:
    int countNodes(TreeNode* root) {
        if(root==nullptr){
            return 0;
        }
        queue<TreeNode*>qu;
        int cnt=0;
        qu.push(root);
        while(!qu.empty()){
            TreeNode*cur=qu.front();
            qu.pop();
            cnt++;
            if(cur->left){
                qu.push(cur->left);
            }
            if(cur->right){
                qu.push(cur->right);
            }            
        }
        return cnt;
    }
};
