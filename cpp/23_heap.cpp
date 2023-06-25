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
    */
public:
    struct greater{
        bool operator()(ListNode*node1,ListNode*node2){
            return node1->val>node2->val;
        }
    };
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if(lists.empty()){
            return nullptr;
        }
        int k=lists.size();
        // 构建大小k的最小堆 将k个链表头节点加入优先队列
        priority_queue<ListNode*,vector<ListNode*>,greater>minHeap;
        for(int i=0;i<k;i++){
            if(lists[i]){
                minHeap.push(lists[i]);
            }
        }
        ListNode*head=new ListNode(0);
        ListNode*p=head;
        while(!minHeap.empty()){
            // pop最小值
            ListNode*node=minHeap.top();
            p->next=node;
            minHeap.pop();
            if(node->next){
                // 重新堆构造保持堆序列
                minHeap.push(node->next);
            }
            p=p->next;
        }
        return head->next;
    }
};
