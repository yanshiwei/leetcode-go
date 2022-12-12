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
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root==NULL){
            return NULL;
        }
        if(p==root||q==root){
            //当遇到节点 p或 q 时返回。从底至顶回溯，
            return root;
        }
        auto left=lowestCommonAncestor(root->left,p,q);// root左子树寻找
        auto right=lowestCommonAncestor(root->right,p,q);// root右子树寻找
        //left为空，说明p/q不在左边子树
        if(left==NULL){
            return right;
        }
        //right为空，说明p/q不在右边子树
        if(right==NULL){
            return left;
        }
        // left && right不为空，说明 p q在root两侧，root为最近祖先
        //考虑通过递归对二叉树进行先序遍历，当节点p,q 在节点 root 的异侧时，节点 rootr即为最近公共祖先，则向上返回 root。
        return root;
    }
};
