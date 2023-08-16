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
    给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树

          1
         / \
        2   3
       / \     
      4   5    
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
    题解：
    每一条二叉树的「直径」长度，就是一个节点的左右子树的最大深度之和。 求整棵树中的最长「直径」，那直截了当的思路就是遍历整棵树中的每个节点，然后通过每个节点的左右子树的最大深度算出每个节点的「直径」，最后把所有「直径」求个最大值即可。
    */
public:
    // 任意点到任意点路径，DFS，https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
    int maxD=0;
    // root为根结点的最大深度
    int helper(TreeNode*root){
        if(root==nullptr){
            return 0;
        }
        int left=helper(root->left);//左儿子为根的子树的深度
        int right=helper(root->right);//右儿子为根的子树的深度
        maxD=max(maxD,left+right);//当前结点的直径为左右子树深度之和，更新全局最大值
        //返回该节点为根的子树的深度
        return max(left,right)+1;
    }
    int diameterOfBinaryTree(TreeNode* root) {
        helper(root);
        return maxD;
    }
};
