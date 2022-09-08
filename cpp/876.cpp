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
    给定一个头结点为 head 的非空单链表，返回链表的中间结点。如果有两个中间结点，则返回第二个中间结点。
    */
public:
    ListNode* middleNode(ListNode* head) {
        if (head==nullptr||head->next==nullptr){
            return head;
        }
        int cnt=0;
        ListNode*p=head;
        while(p){
            cnt++;
            p=p->next;
        }
        int mid=cnt/2+1;
        p=head;
        ListNode*pre=nullptr;
        int idx=0;
        while(p&&idx!=mid){
            idx++;
            pre=p;
            p=p->next;
        }
        return pre;
    }
};
