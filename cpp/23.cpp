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
    给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。
    */
public:
    // 利用优先级队列（堆）优化每次从k个节点中寻找最小节点
    //  priority_queue push pop top size empty swap
    struct cmp{
        bool operator()(ListNode*node1,ListNode*node2){
            return node1->val>node2->val;
        }// 小顶堆选择greater
    };
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if(lists.size()<1){
            return nullptr;
        }
        priority_queue<ListNode*,vector<ListNode*>,cmp>que;// 最小堆 求K个的最小值
        for(auto node:lists){
            if(node){
                que.push(node);//将k个链表头节点加入优先队列
            }
        }
        ListNode*head=new ListNode(0);
        ListNode*p=head;
        while(que.empty()==false){
            auto cur=que.top();// 堆顶是K个里最小节点
            p->next=cur;
            que.pop();
            p=p->next;// 更新q
            if(p->next){
                // 如果该链表后面还有节点
                que.push(p->next);
            }
        }
        return head->next;
    }
};
