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
    unordered_map<int,TreeNode*>parents;//用一个哈希表记录每个结点的父结点，由于结点值唯一，所以key可以使用结点的值
    unordered_set<int>hset;
    // 由于Tree没有记录父结点，故先DFS获取
    void dfsForParents(TreeNode*root){
        if(root==nullptr){
            return;
        }
        if(root->left){
            parents[root->left->val]=root;
            dfsForParents(root->left);
        }
        if(root->right){
            parents[root->right->val]=root;
            dfsForParents(root->right);
        }
    }
    // 从 target 出发，使用深度优先搜索遍历整棵树，除了搜索左右儿子外，还可以顺着父结点向上搜索
    // 由于可以从左右和父结点3个方向DFS，故可能存在重复访问问题，需要去重
    void dfsForDistance(TreeNode*target, int k){
        if(target==nullptr){
            return;
        }
        if(hset.count(target->val)){
            // 已经访问过
            return;
        }
        hset.insert(target->val);
        if(k==0){
            res.push_back(target->val);
            return;
        }
        // 3个方向遍历
        dfsForDistance(target->left,k-1);
        dfsForDistance(target->right,k-1);
        dfsForDistance(parents[target->val],k-1);//新增方向
    }
    vector<int> distanceK(TreeNode* root, TreeNode* target, int k) {
        parents[root->val]=nullptr;
        dfsForParents(root);//基于root DFS获取父结点
        dfsForDistance(target,k);// 基于target开始不同方向尝试
        return res;
    }
};
