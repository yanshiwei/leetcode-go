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
public:
    TreeNode* buildTreeHelper(vector<int>& inorder,int i_start,int i_end,vector<int>& postorder,int p_start,int p_end,unordered_map<int,int>&hash){
        if(p_start>=p_end){
            return nullptr;
        }
        int root_val=postorder[p_end-1];
        TreeNode*root=new TreeNode(root_val);
        unordered_map<int,int>::iterator it=hash.find(root_val);
        int idx=it->second;
        int leftNum=idx-i_start;
        // left
        root->left=buildTreeHelper(inorder,i_start,idx,postorder,p_start,p_start+leftNum,hash);
        // right
        root->right=buildTreeHelper(inorder,idx+1,i_end,postorder,p_start+leftNum,p_end-1,hash);
        return root;
    }
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        unordered_map<int,int>hash;//每次for循环求根节点耗时，空间换时间
        for(int i=0;i<inorder.size();i++){
            hash.insert(pair<int,int>(inorder[i],i));
        }
        return buildTreeHelper(inorder,0,inorder.size(),postorder,0,postorder.size(),hash);
    }
};
