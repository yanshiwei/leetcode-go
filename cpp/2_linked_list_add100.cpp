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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode*p1=l1;
        ListNode*p2=l2;
        int carry=0;
        ListNode*head=nullptr;
        ListNode*tail=head;
        while(p1||p2){
            int v1=0;
            if(p1){
                v1=p1->val;
            }
            int v2=0;
            if(p2){
                v2=p2->val;
            }
            int sum=v1+v2+carry;
            if(head==nullptr){
                head=new ListNode(sum%10);
                tail=head;
            }else{
                tail->next=new ListNode(sum%10);
                tail=tail->next;
            }
            carry=sum/10;
            if(p1) p1=p1->next;
            if(p2) p2=p2->next;
        }
        if(carry){
            tail->next=new ListNode(carry);
            tail=tail->next;
        }
        return head;
    }
};
