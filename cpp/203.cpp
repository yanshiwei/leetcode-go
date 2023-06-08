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
    // 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 
public:
    ListNode* removeElements(ListNode* head, int val) {
        if(head==nullptr){
            return nullptr;
        }
        ListNode*newHead=new ListNode(0);
        newHead->next=head;
        ListNode*pre=newHead;
        ListNode*cur=head;
        while(cur){
            if(cur->val==val){
                pre->next=cur->next;
            }else{
                pre=cur;
            }
            cur=cur->next;
        }
        return newHead->next;
    }
};
