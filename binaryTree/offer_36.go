/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
public:
    Node* treeToDoublyList(Node* root) {
        if (root==nullptr){
            return nullptr;
        }
        dfs(root);
        head->left=pre;//head的上一个就是最后一个
        pre->right=head;//最后一个的下一个就是第一个
        return head;
    }
    private:
    Node *pre, *head;
    void dfs(Node *cur){
        if (cur==nullptr){
            return;
        }
        dfs(cur->left);
        if (pre!=nullptr){
            pre->right=cur;
        }else{
            ///init
            head=cur;
        }
        cur->left=pre;
        pre=cur;
        dfs(cur->right);
    }
};
