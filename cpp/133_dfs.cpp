/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;
    Node() {
        val = 0;
        neighbors = vector<Node*>();
    }
    Node(int _val) {
        val = _val;
        neighbors = vector<Node*>();
    }
    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};
*/

class Solution {
public:
    unordered_map<Node*,Node*>hash;//源结点到克隆节点之间的映射关系
    Node* cloneGraph(Node* node){
        if(node==nullptr){
            return nullptr;
        }
        return dfs(node);
    }
private:
    Node*dfs(Node*node){
        if(hash[node]){
            // node对应的clone已经存在，无需再处理
            return hash[node];
        }
        // clone本结点
        Node*clone=new Node(node->val);
        hash[node]=clone;
        // clone边
        for(Node*n:node->neighbors){
            // 递归该结点对应
            clone->neighbors.push_back(dfs(n));
        }
        return clone;
    }
};
