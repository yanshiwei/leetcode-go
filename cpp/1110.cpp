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
    给出二叉树的根节点 root，树上每个节点都有一个不同的值。
如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
返回森林中的每棵树。你可以按任意顺序组织答案
    */
public:
    unordered_set<int>mySet;
    TreeNode* dfs(TreeNode*root,vector<TreeNode*>&res,bool parent){
        if(root==nullptr){
            return nullptr;
        }
        bool del=false;
        if(mySet.count(root->val)){
            del=true;
        }
        if(del==false&&parent==false){
            // 没有删除且没有父节点，就是一个新的根节点
            res.push_back(root);
        }
        // 去左右子树进行删除
        root->left=dfs(root->left,res,!del);
        root->right=dfs(root->right,res,!del);
        // 如果需要被删除，返回 null 给父节点
        return del?nullptr:root;
    }
    vector<TreeNode*> delNodes(TreeNode* root, vector<int>& to_delete) {
        for (auto it = to_delete.begin(); it != to_delete.end(); it++){
            mySet.insert(*it);
        }
        vector<TreeNode*> res;
        dfs(root,res,false);
        return res;
    }
};
