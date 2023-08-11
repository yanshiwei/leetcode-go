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
    给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。
    */
public:
    // same tree algo，判断两个树是否相同
    bool helper(TreeNode* A, TreeNode* B){
        if(A==nullptr&&B==nullptr){
            return true;
        }
        if(A==nullptr||B==nullptr||A->val!=B->val){
            return false;
        }
        // A & B, and A==B
        return helper(A->left,B->left)&&helper(A->right,B->right);
    }
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        if(root==nullptr||subRoot==nullptr){
            return false;
        }
        // 子结点subRoot可能是与root种任意结点相同，故需要双重DFS
        return helper(root,subRoot)||isSubtree(root->left,subRoot)||isSubtree(root->right,subRoot);
    }
};
