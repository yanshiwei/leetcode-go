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
class BSTIterator {
    /*
    实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。
    */
public:
    vector<int>ve;
    int i;
    BSTIterator(TreeNode* root) {
        i=0;
        inorder(root);
    }
    
    int next() {
        return ve[i++];
    }
    
    bool hasNext() {
        return i<ve.size();
    }
private:
    void inorder(TreeNode*root){
        TreeNode*cur=root;
        stack<TreeNode*>st;
        while(cur||st.empty()==false){
            while(cur){
                st.push(cur);
                cur=cur->left;
            }
            //存储最左侧节点
            if(st.empty()==false){
                TreeNode*r=st.top();
                st.pop();
                ve.push_back(r->val);
                cur=r->right;
            }
        }
    }
};

/**
 * Your BSTIterator object will be instantiated and called as such:
 * BSTIterator* obj = new BSTIterator(root);
 * int param_1 = obj->next();
 * bool param_2 = obj->hasNext();
 */
