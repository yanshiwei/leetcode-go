/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
    // https://blog.csdn.net/qq_22326601/article/details/122176506
public:
    // a = (n - 1) * r + (L -a -x)
    // 从链表起点head开始到入口点的距离a,与从slow和fast的相遇点到入口点的距离相等a = L-a-x
    ListNode *detectCycle(ListNode *head) {
        if(head==nullptr){
            return head;
        }
        ListNode*fast=head;
        ListNode*slow=head;
        while(fast&&fast->next){
            slow=slow->next;
            fast=fast->next->next;
            if(fast==slow){
                break;
            }
        }
        if(fast==nullptr||fast->next==nullptr){
            // no cycle
            return nullptr;
        }
        ListNode* inter = fast;//保存相遇节点
        //从头结点到入口的距离，等于转了(n-1)圈以后，相遇点到入口的距离
        ListNode *cur=head;
        while(cur!=slow){
            cur=cur->next;
            slow=slow->next;
        }
        return slow;
    }
};
