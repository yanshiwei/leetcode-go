/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
    /*
    题目：
    给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
    题解：k个链表，每次需要拿出最小值作为新链表一个结点。最小堆
    自定义比较的三种方案：
    1 类内重载操作符
    bool operator<(const node&, const node&b){
        return a.val < b.val;// 按照value从大到小排列
    }
    priority_queue<node>qu;
    2 自定义仿函数
    struct cmp{
        bool operator()(const node&, const node&b){
            return a.val >b.val;// 按照value从小到大排列
        }
    }
    priority_queue<node,vector<node>,cmp>qu;
    3 类内定义友元操作类重载函数
    struct node{
        int val;
        friend bool operator(const node&, const node&b){
            return a.val < b.val;// 按照value从大到小排列
        }
    }
    priority_queue<node>qu;
    */
public:
    struct greater{
        bool operator()(ListNode*a,ListNode*b){
            return a->val>b->val;
        }
    };
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if(lists.empty()){
            return nullptr;
        }
        int k=lists.size();
        // 构建大小k的最小堆 将k个链表头节点加入优先队列
        priority_queue<ListNode*,vector<ListNode*>,greater>minHeap;//base vector deque
        for(int i=0;i<k;i++){
            if(lists[i]){
                minHeap.push(lists[i]);
            }
        }
        ListNode*dummy=new ListNode();//构造dummy，方便返回后续节点
        ListNode*pre=dummy;
        while(!minHeap.empty()){
            // 先获得最小值
            ListNode*node=minHeap.top();
            minHeap.pop();
            pre->next=node;
            pre=pre->next;
            if(node->next){
                minHeap.push(node->next);
            }
        }
        return dummy->next;
    }
};
