/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
    /*
    题目：
    给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
    */
public:
    Node* connect(Node* root) {
        if(root==NULL){
            return NULL;
        }
        queue<Node*>qu;
        qu.push(root);
        while(qu.empty()==false){
            int size=qu.size();
            Node*pre=NULL;
            for(int i=0;i<size;i++){
                Node*cur=qu.front();
                qu.pop();
                if(i>0){
                    pre->next=cur;
                }
                if(cur->left){
                    qu.push(cur->left);
                }
                if(cur->right){
                    qu.push(cur->right);
                }
                pre=cur;
            }
        }
        return root;
    }
};
