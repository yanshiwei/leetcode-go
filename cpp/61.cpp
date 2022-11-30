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
    题目：
    给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
    k <= 2 * 109 直接操作超时
    */
public:
    ListNode* rotateRight(ListNode* head, int k) {
        if(head==nullptr||head->next==nullptr||k<1){
            return head;
        }
        int len=1;
        ListNode*p=head;
        while(p->next){
            len++;
            p=p->next;
        }
        // 首尾相连
        p->next=head;
        // 再步行len-k%len就到达下一个链表头的前节点
        for(int i=0;i<len-k%len;i++){
            p=p->next;
        }
        ListNode*newHead=p->next;
        // 断开
        p->next=nullptr;
        return newHead;
    }
};
