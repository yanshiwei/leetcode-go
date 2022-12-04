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
    https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/297276/shou-hui-tu-jie-hen-you-ya-de-yi-dao-dfsti-by-hyj8/
    */
public:
    int dfs(TreeNode*root){
        if(root==nullptr){
            return 0;
        }
        int left=dfs(root->left);
        int right=dfs(root->right);
        // 每递归一个子树，都求一下当前子树内部的最大路径和并更新全局最大路径和
        // 以root为根节点的子树中的内部路径，要包含根节点
        int innerSum=left+right+root->val;
        res=max(res,innerSum);
        // 以root为根节点的子树 对外提供的最大收益 分三种情况
        // 只包含root，root+left,root+right
        // 如果某个子树 dfs 结果为负，走入它，收益不增反减，该子树就没用，需杜绝走入，像对待 null 一样让它返回 0
        int sunSum=max(left,right);
        int outSum=root->val+max(0,sunSum);
        return max(0,outSum);
    }
    int res=0x80000000;
    int maxPathSum(TreeNode* root) {
        dfs(root);
        return res;
    }
};
