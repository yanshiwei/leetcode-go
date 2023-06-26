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
    数组版本插入排序
    假设a[0]-a[i-1]已经是排序序列，则从未排序序列a[i]-a[n-1]拿出a[i]插入排序序列
    从a[i-1]开始向前遍历，当大于a[i]的向后挪，直到找到第一个<=a[i]的插入
    链表类似，只不过已排序序列不能向前遍历，只能向后，故需要一个指针指向已排序好的链表
    */
public:
    ListNode* insertionSortList(ListNode* head) {
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*newHead=new ListNode(0);//指向已排序好的链表
        ListNode*pre=newHead;
        ListNode*cur=head;
        while(cur){
            // 已排序好的链表找到第一个大于cur
            ListNode*next=cur->next;
            while(pre->next&&pre->next->val<cur->val){
                pre=pre->next;
            }
            // pre->next->val >= cur->next
            cur->next=pre->next;
            pre->next=cur;
            pre=newHead;
            cur=next;
        }
        return newHead->next;
    }
};
