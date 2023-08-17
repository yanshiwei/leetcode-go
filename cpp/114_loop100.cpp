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
    给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
题解：
1. 将左子树插入到右子树的地方
2. 将原来的右子树接到左子树的最右边节点
3. 考虑新的右子树的根节点，一直重复上边的过程，直到新的右子树为 null
    */
public:
    void flatten(TreeNode* root) {
        if(root==nullptr){
            return;
        }
        while(root){
            if(root->left==nullptr){
                // 左子树为空 直接考虑下一个节点
                root=root->right;
            }else{
                // 左子树最右侧节点
                TreeNode*pre=root->left;
                while(pre->right){
                    pre=pre->right;
                }
                // 原来的右子树接到左子树的最右边节点
                pre->right=root->right;
                // 将左子树插入到右子树旧位置
                root->right=root->left;
                root->left=nullptr;// 原来左子树旧位置置空
                // 考虑下一个节点
                root=root->right;
            }
        }
    }
};
