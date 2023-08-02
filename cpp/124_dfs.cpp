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
    题目：
    路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和
    题解：
    https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/262792/li-jie-lu-jing-gai-nian-jian-dan-si-lu-jie-ti-fu-t/
    */
public:
    // 二叉树中的路径可以分为以下四种：
    // 1 单一结点；2 结点+左子树；3 结点+右子树；4 左子树+结点+右子树
    // 要注意其中类型4的路径是无法作为子路径返回给上一级结点的，否则形成的就不是路径（只有惟一的起点和终点）而是树
    // dfs返回root树的最大路径和
    int dfs(TreeNode*root){
        if(root==nullptr){
            return 0;
        }
        int leftpath=dfs(root->left)+root->val;//for 2
        int rightpath=dfs(root->right)+root->val;// for 3
        int rootpath=leftpath+rightpath-root->val;// for 4
        res=max(res,max(root->val,max(rootpath,max(leftpath,rightpath))));
        return max(root->val,max(leftpath,rightpath));// 4 not return
    }
    int res=0x80000000;
    int maxPathSum(TreeNode* root) {
        dfs(root);
        return res;
    }
};
