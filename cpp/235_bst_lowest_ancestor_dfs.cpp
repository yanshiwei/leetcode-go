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
    // p 比r且q比r大，则r为最近祖先
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root==NULL){
            return NULL;
        }
        // p q比root都小，左边找
        if(p->val<root->val&&q->val<root->val){
            return lowestCommonAncestor(root->left,p,q);
        }
        // p q比root都大，右边找
        if(p->val>root->val&&q->val>root->val){
            return lowestCommonAncestor(root->right,p,q);
        } 
        // p 比r且q比r大，则r为最近祖先
        // p or q is root，则r为最近祖先
        return root;   
    }
};
