/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
    /*
    给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
    */
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if(headA==NULL||headB==NULL){
            return NULL;
        }
        int lenA=0;
        int lenB=0;
        ListNode*p=headA;
        while(p){
            lenA++;
            p=p->next;
        }
        p=headB;
        while(p){
            lenB++;
            p=p->next;
        }
        int lenMore=lenA;
        p=headA;
        ListNode*lessP=headB;
        if(lenA<lenB){
            lenMore=lenB;
            p=headB;
            lessP=headA;
        }
        int diff=lenA-lenB;
        if(diff<0){
            diff=-diff;
        }
        while(p&&diff>0){
            diff--;
            p=p->next;
        }
        while(p&&lessP){
            if(p==lessP){
                break;
            }
            p=p->next;
            lessP=lessP->next;
        }
        return p;
    }
};
