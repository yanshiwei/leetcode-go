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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        int length=0;
        ListNode*p=head;
        while(p){
            length++;
            p=p->next;
        }
        int realN=length-n;
        if(realN==0){
            // detele old head
            return head->next;
        }
        p=head;
        int curLen=0;
        while(p){
            curLen++;
            if(curLen==realN){
                p->next=p->next->next;
            }
            p=p->next;
        }
        return head;
    }
};
