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
public:
    ListNode* sortList(ListNode* head) {
        /*
        if (head==nullptr){
            return head;
        }
        ListNode*dummyHead=new ListNode(0,head);
        quick_sort_list(dummyHead->next,nullptr);
        return dummyHead->next;
        */
        if (head==nullptr){
            return head;
        }
        int length=0;
        ListNode* node=head;
        while(node){
            length++;
            node=node->next;
        }
        // 子数组大小分别为1、2、4、8···，刚开始是1然后2（自底向上，即先按照一个数组两个个元素排序然后四个）
        ListNode*dummyHead=new ListNode(0,head);
        for(int seg=1;seg<length;seg+=seg){
            // 从最左侧开始
            ListNode*prev=dummyHead,*cur=dummyHead->next;
            while(cur){
                // 子序列长度seg，合并2个这样的子序列
                ListNode*head1=cur;
                for(int i=1;i<seg&&cur->next;i++){
                    cur=cur->next;
                }
                ListNode*head2=cur->next;
                cur->next=nullptr;//独立子序列
                cur=head2;
                for(int i=1;i<seg&&cur&&cur->next;i++){
                    cur=cur->next;
                }
                ListNode*next=nullptr;
                if(cur){
                    next=cur->next;
                    cur->next=nullptr;//独立子序列
                }
                ListNode*merged=merge(head1,head2);
                prev->next=merged;
                // 遍历到已排序的末尾
                while(prev->next){
                    prev=prev->next;
                }
                cur=next;
            }
        }
        return dummyHead->next;
    }
    ListNode*merge(ListNode*head1,ListNode*head2){
        ListNode*dummyHead=new ListNode(0,nullptr);
        ListNode *tmp=dummyHead,*tmp1=head1,*tmp2=head2;
        while(tmp1&&tmp2){
            if(tmp1->val<=tmp2->val){
                tmp->next=tmp1;
                tmp1=tmp1->next;
            }else{
                tmp->next=tmp2;
                tmp2=tmp2->next;
            }
            tmp=tmp->next;
        }
        if(tmp1){
            tmp->next=tmp1;
        }
        if(tmp2){
            tmp->next=tmp2;
        }
        return dummyHead->next;
    }
    // 基准 小 小 小 大 大 大 未知 未知
    ListNode *GetPart(ListNode *begin, ListNode *end) {
    int baseVal = begin->val;
    ListNode *p = begin; // 前面指针找比基数小的
    ListNode *q = p->next; // 后面指针找比基数大的
    // 前后指针遍历
    while (q != end) {
        if (q->val < baseVal) {
            p = p->next;
            if(p!=q){
              int tmp = p->val;
              p->val = q->val;
              q->val = tmp;
            }
        }
        q = q->next;
    }
    int tmp = p->val;
    p->val = begin->val;
    begin->val = tmp;
    return p;
}

void quick_sort_list(ListNode *begin, ListNode *end) {
    if(begin==nullptr||begin->next==nullptr||begin==end||begin->next==end){
        return;
    }
        ListNode *idx = GetPart(begin, end);
        quick_sort_list(begin, idx);
        quick_sort_list(idx->next, end);
}
};
