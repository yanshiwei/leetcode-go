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
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        return dfs(cloned,target);
    }
    TreeNode* dfs(TreeNode* cloned, TreeNode* target){
        if (cloned==nullptr){
            return nullptr;
        }
        if (cloned->val==target->val){
            return cloned;
        }
        auto res1=dfs(cloned->left,target);
        auto res2=dfs(cloned->right,target);
        if (res1!=nullptr){
            return res1;
        }else{
            return res2;
        }
    }
};
