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
    给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 
    */
public:
    ListNode* reverseBetween(ListNode* head, int left, int right) {
        if (head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*dummyHead=new ListNode();
        dummyHead->next=head;
        ListNode*leftPre=dummyHead;
        ListNode*p=dummyHead->next;
        int idx=1;
        while(p&&idx!=left){
            idx++;
            leftPre=p;
            p=p->next;
        }
        ListNode*leftBegin=p;
        while(p&&idx!=right){
            idx++;
            p=p->next;
        }
        ListNode*rightTail=p->next;
        // reverse, head is pre,
        ListNode*pre=nullptr;
        p=leftPre->next;
        while(p&&p!=rightTail){
            ListNode*next=p->next;
            p->next=pre;
            pre=p;
            p=next;
        }
        leftPre->next=pre;
        leftBegin->next=rightTail;
        return dummyHead->next;
    }
};
