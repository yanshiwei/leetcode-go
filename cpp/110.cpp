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
    int depth(TreeNode*root){
        if(root==nullptr){
            return 0;
        }
        return max(depth(root->left),depth(root->right))+1;
    }
    int abs(int x){
        if(x<0){
            return -x;
        }
        return x;
    }
    bool isBalanced(TreeNode* root) {
        if(root==nullptr){
            return true;
        }
        int left=depth(root->left);
        int right=depth(root->right);
        int diff=abs(left-right);
        if(diff>1){
            return false;
        }
        return isBalanced(root->left)&&isBalanced(root->right);
    }
};
