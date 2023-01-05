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
    给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
    题解：
    对于[1,2,3...,n],构造BST，可以有以下方式
    以1为根，[]为左子树，[2,3,...,n]构成右子树
    以2为根，[1]为左子树，[3,4,...,n]构成右子树
    。。。
    以n为根，[1,2,...,n-1]为左子树，[]构成右子树
    同样的对于[2,3,...,n]或者[3,4,...,n]或者[1,2,...,n-1]相同的做法，继续遍历其中每个数字作为根，然后切分成左右子树
    知道只有一个数字或者为空则返回
    */
public:
    vector<TreeNode*> dfs(int start,int end){
        vector<TreeNode*> res;
        if(start>end){
            //没有数字
            res.push_back(nullptr);
            return res;
        }
        if(start==end){
            //只有一个数字
            TreeNode*root=new TreeNode(start);
            res.push_back(root);
            return res;
        }
        // 多个数字，遍历其中每个数字作为根，然后切分成左右子树
        for(int i=start;i<=end;i++){
            vector<TreeNode*> leftList=dfs(start,i-1);
            vector<TreeNode*> rightList=dfs(i+1,end);
            // 左子树集合和右子树结合两两结合
            for(auto &l:leftList){
                for(auto &r:rightList){
                    TreeNode*root=new TreeNode(i);
                    root->left=l;
                    root->right=r;
                    res.push_back(root);
                }
            }
        }
        return res;
    }
    vector<TreeNode*> generateTrees(int n) {
        if(n<1){
            return {};
        }
        return dfs(1,n);
    }
};
