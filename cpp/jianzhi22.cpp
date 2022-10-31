/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* getKthFromEnd(ListNode* head, int k) {
        if(head==NULL){
            return head;
        }
        if(k<1){
            return NULL;
        }
        ListNode*p=head;
        for(int i=0;i<k-1;i++){
            if(p){
                p=p->next;
            }else{
                return NULL;
            }
        }
        //cout<<p->val;
        ListNode*q=head;
        while(p->next){
            q=q->next;
            p=p->next;
        }
        return q;
    }
};
