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
    // 若p/q在root两侧，则最近祖先就是root；
    // 若p/q其中就是root，则最近祖先就是root;
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root==NULL){
            return NULL;
        }
        // 若p/q其中就是root，则最近祖先就是root
        if(p==root||q==root){
            return root;
        }
        TreeNode*left=lowestCommonAncestor(root->left,p,q);
        TreeNode*right=lowestCommonAncestor(root->right,p,q);
        // left && right不为空，说明 p q在root两侧，root为最近祖先
        if(left&&right){
            return root;
        }
        //left为空，说明p/q不在左边子树
        if(left==NULL){
            return right;
        }
        //right为空，说明p/q不在右边子树
        if(right==NULL){
            return left;
        }
        return root;
    }
};
