class LRUCache {
    /* 双端链表+map
    map:insert erase find count empty/size
    list:push_back/push_front; pop_back pop_front
         back. front
         clear erase
         size empty
         sort
    请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类
    */
public:
    LRUCache(int capacity):cap(capacity) {

    }
    // 注意访问一次 就要更新到最新
    int get(int key) {
        if(key2node.find(key)==key2node.end()){
            return -1;
        }
        pair<int,int> node=*key2node[key];//*p is value
        cache.erase(key2node[key]);//将节点移到链表头部并更新map
        cache.push_front(node);
        key2node[key]=cache.begin();
        return node.second;
    }
    // 满了就删除最旧的数据
    void put(int key, int value) {
        pair<int,int> newNode=make_pair(key,value);
        if(key2node.count(key)){
            // 若该节点已存在，则删除旧的节点
            cache.erase(key2node[key]);
        }else{
            // 节点不存在
            if(cap==cache.size()){
                // full 删除链表最后一个数据
                key2node.erase(cache.back().first);
                cache.pop_back();
            }
        }
        cache.push_front(newNode);//插入新的节点到头部
        key2node[key]=cache.begin();
    }
private:
    list<pair<int,int>>cache;// 双向链表，{pair.value}
    unordered_map<int,list<pair<int,int>>::iterator>key2node;// key to list node
    int cap;
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
