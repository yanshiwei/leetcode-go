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
public:
    ListNode* swapPairs(ListNode* head) {
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*dummy=new ListNode();
        dummy->next=head;
        ListNode*p=dummy;
        while(p&&p->next&&p->next->next){
            auto p1=p->next;
            auto p2=p->next->next;
            p->next=p2;
            p1->next=p2->next;
            p2->next=p1;
            p=p1;
        }
        return dummy->next;
    }
};
