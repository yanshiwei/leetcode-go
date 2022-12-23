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
    // same tree algo
    bool helper(TreeNode* A, TreeNode* B){
        //这里如果B为空并且A为空，表示A和B已经遍历完成，并且经过了A.val!=B.val的考验
        if(A==nullptr&&B==nullptr){
            return true;
        }
        if(A==nullptr||B==nullptr||A->val!=B->val){
            return false;
        }
        //当前节点比较完之后还要继续判断左右子节点
        return helper(A->left,B->left)&&helper(A->right,B->right);
    }
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        if(root==nullptr||subRoot==nullptr){
            return false;
        }
        //subTree不光有可能是root的子树，也有可能是root左子树的子树或者右子树的子树，所以如果从根节点判断subTree不是root的子树，还需要判断root的left和A的right
        return helper(root,subRoot)||isSubtree(root->left,subRoot)||isSubtree(root->right,subRoot);
    }
};
