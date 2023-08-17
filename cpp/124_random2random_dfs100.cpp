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
    // 任意点到任意点，dfs https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
    // dfs返回root树的最大路径和
    int dfs(TreeNode*root){
        if(root==nullptr){
            return 0;
        }
        int leftpath=max(0,dfs(root->left));//如果最大路径和<0,意味着该路径和对总路径和做负贡献，因此不要计入到总路径中，将它设置为0
        int rightpath=max(0,dfs(root->right));
        res=max(res,root->val+leftpath+rightpath);//root能提供的最大路径和,可能root/root+left/root+right/root+left+right
        return max(leftpath,rightpath)+root->val;
    }
    int res=0x80000000;
    int maxPathSum(TreeNode* root) {
        dfs(root);
        return res;
    }
};
