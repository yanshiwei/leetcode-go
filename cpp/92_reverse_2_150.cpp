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
    给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
    */
public:
    ListNode* reverseBetween(ListNode* head, int left, int right) {
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*dummy=new ListNode(0);
        dummy->next=head;
        ListNode*pre=dummy;// 要反转的序列前一个节点
        int count = 1;  // 节点编号
        while(count<left){
            pre=pre->next;
            count++;
        }
        ListNode*reverseHead=pre->next;// 获取反转区间的头节点
        // 反转区间[left, right]的节点
        ListNode*cur=reverseHead;
        ListNode*last=nullptr;
        ListNode*next;
        while(count<=right){
            next=cur->next;
            cur->next=last;
            last=cur;
            cur=next;
            count++;
        }
        // 重新拼接反转后的节点
        pre->next=last;// 反转区间前一个节点应该连接到反转区间的最后一个节点，即当前的last
        reverseHead->next=cur;// 反转区间的头节点应该连接到反转区间的下一个节点，即当前的next
        return dummy->next;
    }
};
