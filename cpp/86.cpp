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
    给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。
    */
public:
    ListNode* partition(ListNode* head, int x) {
        //虚拟头结点 dummy 比较原链表分为两个子链表，一个小于 x，一个不小于 x
        ListNode dummy1=ListNode(0);
        dummy1.next=nullptr;
        ListNode dummy2=ListNode(0);
        dummy2.next=nullptr;    
        ListNode*p1=&dummy1;
        ListNode*p2=&dummy2;
        ListNode*p=head;
        while(p){
            if(p->val<x){
                p1->next=p;
                p1=p1->next;
            }else{
                p2->next=p;
                p2=p2->next;
            }
            p=p->next;
        }
        p1->next=dummy2.next;
        p2->next=nullptr;
        return dummy1.next;
    }
};
