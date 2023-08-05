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
    我们可以为二叉树 T 定义一个 翻转操作 ，如下所示：选择任意节点，然后交换它的左子树和右子树。
只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转 等价 于二叉树 Y。
这些树由根节点 root1 和 root2 给出。如果两个二叉树是否是翻转 等价 的函数，则返回 true ，否则返回 false 。
    */
public:
    bool flipEquiv(TreeNode* root1, TreeNode* root2) {
        if(root1==nullptr&&root2==nullptr){
            return true;
        }
        if(root1==nullptr||root2==nullptr){
            return false;
        }     
        if(root1->val!=root2->val){
            return false;
        }   
        //当 root1 和 root2 的值相等的情况下，需要继续判断 root1 的孩子节点是不是跟 root2 的孩子节点相当
        //case 1节点本来就一样； case2 节点对称一样
        return (flipEquiv(root1->left,root2->left)&&flipEquiv(root1->right,root2->right))||(flipEquiv(root1->left,root2->right)&&flipEquiv(root1->right,root2->left));
    }
};
