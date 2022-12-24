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
    /*
    给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。
返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。
    */
public:
    vector<int>res;
    unordered_map<TreeNode*,TreeNode*>myHash;//存储每个节点的父亲节点
    unordered_set<TreeNode*>mySet;//在遍历k范围内的节点时要确定这个节点是否已经来过
    void dfsForParent(TreeNode*root){
        if(root==NULL){
            return;
        }
        myHash[root->left]=root;
        myHash[root->right]=root;
        dfsForParent(root->left);
        dfsForParent(root->right);
    }
    void dfsForSearch(TreeNode*target,int k){
        if(target==NULL){
            return;
        }
        if(mySet.count(target)){
            return;
        }
        mySet.insert(target);// 记录来过该节点
        if(k==0){
            res.push_back(target->val);
            return;
        }
        k--;
        //继续搜索左右节点和父亲节点
        dfsForSearch(target->left,k);
        dfsForSearch(target->right,k);
        dfsForSearch(myHash[target],k);
    }
    vector<int> distanceK(TreeNode* root, TreeNode* target, int k) {
        myHash[root]=NULL;
        dfsForParent(root);// 为父节点建立链接
        dfsForSearch(target,k);
        return res;
    }
};
