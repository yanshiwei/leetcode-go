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
    给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
    */
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode dummy=ListNode();
        dummy.next=head;
        ListNode*pre=&dummy;
        ListNode*cur=head;
        while(cur&&cur->next){
            bool dep=false;
            while(cur->next&&cur->val==cur->next->val){
                cur=cur->next;
                dep=true;
            }
            if(dep){
                pre->next=cur->next;
            }else{
                pre=cur;
            }
            cur=cur->next;
        }
        return dummy.next;
    }
};
