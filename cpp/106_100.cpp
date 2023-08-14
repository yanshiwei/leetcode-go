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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if(headA==nullptr){
            return nullptr;
        }
        if(headB==nullptr){
            return nullptr;
        }
        int lenA=0;
        ListNode*cur=headA;
        while(cur){
            lenA++;
            cur=cur->next;
        }
        int lenB=0;
        cur=headB;
        while(cur){
            lenB++;
            cur=cur->next;
        }
        int lenBig = lenA;
        cur = headA;
        ListNode*lessCur=headB;
        if(lenA<lenB){
            lenBig=lenB;
            cur=headB;
            lessCur=headA;
        }
        int diff=lenA-lenB;
        if(diff<0){
            diff=-diff;
        }
        while(cur&&diff>0){
            diff--;
            cur=cur->next;
        }
        while(cur&&lessCur){
            if(cur==lessCur){
                break;
            }
            cur=cur->next;
            lessCur=lessCur->next;
        }
        return cur;
    }
};
