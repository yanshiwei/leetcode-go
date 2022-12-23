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
    给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
    */
public:
    TreeNode* helper(vector<int>& preorder,int p_start,int p_end,vector<int>& inorder,int i_start,int i_end){
        // preorder 为空，直接返回 nullptr
        if(p_start==p_end){
            return nullptr;
        }
        int root_val=preorder[p_start];
        TreeNode* root=new TreeNode(root_val);
        //在中序遍历中找到根节点的位置
        int idx=-1;
        for(int i=i_start;i<i_end;i++){
            if(root_val==inorder[i]){
                idx=i;
                break;
            }
        }
        if(idx<0){
            return nullptr;
        }
        int leftNum=idx-i_start;// 左子树节点个数
        //递归的构造左子树
        root->left=helper(preorder,p_start+1,p_start+1+leftNum,inorder,i_start,idx);
        //递归的构造右子树
        root->right=helper(preorder,p_start+1+leftNum,p_end,inorder,idx+1,i_end);
        return root;
    }
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return helper(preorder,0,preorder.size(),inorder,0,inorder.size());
    }
};
