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
    给你一个二叉树的根节点 root ， 检查它是否轴对称
    */
public:
    bool isSymmetric(TreeNode* root) {
        if(root==nullptr){
            return false;
        }
        return isSymmetricCore(root,root);
    }
private:
   bool isSymmetricCore(TreeNode* root1,TreeNode*root2){
       if(root1==nullptr&&root2==nullptr){
           return true;
       }
       if(root1==nullptr||root2==nullptr){
           return false;
       }
       if(root1->val!=root2->val){
           return false;
       }
       return isSymmetricCore(root1->left,root2->right)&&isSymmetricCore(root1->right,root2->left);
   }
};
