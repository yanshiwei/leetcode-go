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
    int res=0;
    void inorderReverse(TreeNode* root){
        if(root==nullptr){
            return;
        }
        inorderReverse(root->right);
        res+=root->val;
        root->val=res;
        inorderReverse(root->left);
    }
    TreeNode* convertBST(TreeNode* root) {
        // 逆序中序遍历，之前是左中右，现在是右中左
        res=0;
        inorderReverse(root);
        return root;
    }
};
