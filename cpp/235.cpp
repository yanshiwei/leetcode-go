/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
    //二叉搜索树的最近公共祖先
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root==NULL){
            return NULL;
        }
        if(root->val>p->val&&root->val>q->val){
            // root's left
            return lowestCommonAncestor(root->left,p,q);
        }else if(root->val<p->val&&root->val<q->val){
            // root's right
            return lowestCommonAncestor(root->right,p,q);
        }else{
            // two sides
            return root;
        }
    }
};
