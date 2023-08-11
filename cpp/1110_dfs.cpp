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
返回森林中的每棵树。你可以按任意顺序组织答案。
    */
public:
    set<int>mySet;
    vector<TreeNode*>res;
    TreeNode*delNodesHelper(TreeNode* root,bool has_parent){
        if(root==nullptr){
            return nullptr;
        }
        bool del_flag=false;
        if(mySet.count(root->val)>0){
            del_flag =true;
        }
        if(del_flag==false&&has_parent==false){
            // 没有删除且没有父节点，就是一个新的根节点
            // 不退出，继续左右子树操作
            res.push_back(root);
        }
        // 去左右子树进行删除
        root->left=delNodesHelper(root->left, !del_flag);
        root->right=delNodesHelper(root->right, !del_flag);
        return del_flag?nullptr:root;
    }
    vector<TreeNode*> delNodes(TreeNode* root, vector<int>& to_delete) {
        for (auto it = to_delete.begin(); it != to_delete.end(); it++){
            mySet.insert(*it);
        }
        delNodesHelper(root,false);
        return res;
    }
};
