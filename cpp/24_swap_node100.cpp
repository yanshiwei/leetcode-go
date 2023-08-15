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
        while(p&&p->next){
            ListNode* p1=p->next;
            ListNode* p2=p->next->next;
            ListNode* p3=nullptr;
            if(p2){
                p3=p2->next;
            }else{
                // 剩余只有一个结点
                break;
            }
            p->next=p2;
            p2->next=p1;
            p1->next=p3;
            p=p1;
        }
        return dummy->next;
    }
};
