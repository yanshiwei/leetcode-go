class LRUCache {
    struct Node{
        int key;
        int val;
        Node*pre;//双端链表，方便增删节点
        Node*next;
        Node():key(0),val(0),pre(nullptr),next(nullptr){}
        Node(int k,int v):key(k),val(v),pre(nullptr),next(nullptr){}
    };
private:
    unordered_map<int,Node*>cache;// dict方便快速查询
    Node*head;//方便头节点插入
    Node*tail;//方便尾节点删除
    int cur_size;
    int cap;
private:
    void deleteNode(Node*node){
        node->pre->next=node->next;
        node->next->pre=node->pre;
    }
    void addTohead(Node*node){
        node->pre=head;
        node->next=head->next;

        head->next->pre=node;
        head->next=node;
    }
    // node in list
    void movoToHead(Node*node){
        // 删除原始位置，节点还需要用故不删除节点
        deleteNode(node);
        // 节点放入头部
        addTohead(node);
    }
    Node*removeFromTail(){
        Node*node=tail->pre;
        deleteNode(node);
        return node;
    }


public:
    LRUCache(int capacity) {
        head=new Node();
        tail=new Node();
        head->next=tail;
        tail->pre=head;

        cur_size=0;
        cap=capacity;
    }
    
    int get(int key) {
        if(cache.count(key)<1){
            return -1;
        }
        Node*node=cache[key];
        // 更新到头节点
        movoToHead(node);
        return node->val;
    }
    
    void put(int key, int value) {
        if(cache.count(key)<1){
            // 不存在
            // 1 构建新节点
            Node*node=new Node(key, value);
            // 2 更新字典
            cache[key]=node;
            // 3 插入链表头部，表示最新
            addTohead(node);
            //4 更新当前大小并判断是否满
            cur_size++;
            if(cur_size>cap){
                // 5.0 满了删除最久的也就是尾部数据
                Node*deleteNode=removeFromTail();
                // 5.1 删除字典中数据
                cache.erase(deleteNode->key);
                // 5.2 避免内存泄漏
                delete deleteNode;
                // 5.3 更新当前大小
                cur_size--;
            }
        }else{
            // 存在
            // 1 更新数据
            Node*node=cache[key];
            node->val=value;
            // 2 插入链表头部，表示最新
            movoToHead(node);
        }
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
