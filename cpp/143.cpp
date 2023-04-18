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
    给定一个单链表 L 的头节点 head ，单链表 L 表示为：
L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
    题解：
    1->2-3->4->5->6
    // 第一步，找到中点，划分两个list
    1->2->3
    4->5->6
    // 第二步，逆序第二个list
    1->2->3
    6->5->4
    // 第三步，合并两个list
    1->6->2->5->3->4
    */
public:
    ListNode*reverseList(ListNode*head){
        ListNode*tail=head;
        head=head->next;
        tail->next=nullptr;
        while(head){
            ListNode*tmp=head->next;
            head->next=tail;
            tail=head;
            head=tmp;
        }
        return tail;
    }
    void reorderList(ListNode* head) {
        if(head==nullptr||head->next==nullptr||head->next->next==nullptr){
            return;
        }
        // 第一步，找到中点，划分两个list
        ListNode*slow=head;
        ListNode*fast=head;
        while(fast->next&&fast->next->next){
            slow=slow->next;
            fast=fast->next->next;
        }
        ListNode*newHead=slow->next;
        slow->next=nullptr;
        // 第二步，逆序第二个list
        newHead=reverseList(newHead);
        // 第三步，合并两个list
        ListNode*cur=head;
        while(cur&&newHead){
            ListNode*tmp=newHead->next;
            newHead->next=cur->next;
            cur->next=newHead;
            cur=newHead->next;
            newHead=tmp;
        }
    }
};
