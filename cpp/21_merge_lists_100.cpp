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
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode*dummy=new ListNode();
        ListNode*cur=dummy;
        ListNode*p1=list1;
        ListNode*p2=list2;
        while(p1&&p2){
            if(p1->val<p2->val){
                cur->next=p1;
                p1=p1->next;
            }else{
                cur->next=p2;
                p2=p2->next;
            }
            cur=cur->next;
        }
        if(p1){
            cur->next=p1;
        }
        if(p2){
            cur->next=p2;
        }
        return dummy->next;
    }
};
