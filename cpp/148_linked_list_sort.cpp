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
    题解：
    快速排序的思想大家应该都知道，就是找一个pivot，把小于pivot分在一边，大于等于pivot的分在另外一边
这个过程也叫做partition分区，数组的分区很好做，左右两个指针i/j，不断交换就好了
链表因为只能单向遍历，所以要换一种partition方法
目的是使得左边的值都小于pivot，右边的值都不小于pivot 所以用一个索引记录左边的坐标，遍历过程中，每次碰到比pivot小的，都要交换一下，放到左边。 遍历完成后，再把pivot交换到中间来，这样就达成了目的
    题解0:https://blog.csdn.net/u010429424/article/details/77776731
    题解1:https://leetcode.cn/problems/sort-list/solutions/9318/gui-bing-pai-xu-he-kuai-su-pai-xu-by-a380922457/
    鉴于快速排序超时，引入归并排序，归并排序是一种分而治之的算法。其思想是将原始数组切割分成较小的数组，直到每个数组中只有一个元素，接着将小数组归并成较大的数组，直到最后只有一个排序完毕的大数组。
    归并排序有自顶向下的递归方法，和自底向上的迭代方法。
    1.自顶向下的归并排序，递归到只剩一个结点或者为空链表，返回的时候返回合并后的链表头结点。 时间复杂度 O(nlogn) 空间复杂度 O(logn)。首先需要找到中间点将两个链表一分为二，再分两部分进行递归排序，最后使用合并两个有序链表的方法合并两部分。一分为二时使用快慢指针找到中点区分。
    2.自底向上的归并排序,先两个两个的 merge，完成一趟后，再 4 个4个的 merge，直到结束
    举个简单的例子：[4,3,1,7,8,9,2,11,5,6]. 
step=1: (3->4)->(1->7)->(8->9)->(2->11)->(5->6)
step=2: (1->3->4->7)->(2->8->9->11)->(5->6)
step=4: (1->2->3->4->7->8->9->11)->5->6
step=8: (1->2->3->4->5->6->7->8->9->11)
    问题关键是断开链接单独成一个链表
    题解：https://leetcode.cn/problems/sort-list/solutions/1525677/by-lafiteee-7vi7/
    */
public:
    ListNode* sortList(ListNode* head) {
        return merge_sort_from_buttom(head);
    }
/* 快速排序超时
    ListNode* sortList(ListNode* head) {
        quick_sort_list(head,nullptr);
        return head;
    }
    */
private:
    // 归并排序
    ListNode*merge(ListNode*head1,ListNode*head2){
        ListNode*dummy=new ListNode();
        ListNode*pre=dummy;
        while(head1!=nullptr&&head2!=nullptr){
            if(head1->val<=head2->val){
                pre->next=head1;
                head1=head1->next;
            }else{
                pre->next=head2;
                head2=head2->next;
            }
            pre=pre->next;
        }
        pre->next=head1!=nullptr?head1:head2;
        return dummy->next;
    }
    ListNode*findMid1(ListNode*head){
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*slow=head;
        ListNode*fast=head->next;
        while(fast&&fast->next){
            // fast is 下一个slow
            slow=slow->next;
            fast=fast->next->next;
        }
        return slow;
    }
    ListNode*findMid2(ListNode*head){
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*slow=head;
        ListNode*fast=head;
        while(fast->next&&fast->next->next){
            // fast->next is 下一个slow
            slow=slow->next;
            fast=fast->next->next;
        }
        return slow;
    }
    ListNode*merge_sort_from_top(ListNode*head){
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*mid=findMid2(head);
        ListNode*head2=mid->next;
        mid->next=nullptr;//断开链接单独成一个链表
        //递归分开直到单结点或者空
        ListNode*left=merge_sort_from_top(head);
        ListNode*right=merge_sort_from_top(head2);
        return merge(left,right);
    }
    ListNode*merge_sort_from_buttom(ListNode*head){
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        int n=0;
        ListNode*cur=head;
        while(cur){
            n++;
            cur=cur->next;
        }
        ListNode*dummy=new ListNode(0,head);
        for(int seg=1;seg<n;seg*=2){
            // 子链表长度，从 1 开始，每次翻倍
            // 从最左侧开始
            ListNode*pre=dummy;
            ListNode*cur=dummy->next;
            while(cur){
                // 子序列长度seg，合并2个这样的子序列
                ListNode*head1=cur;// 第一个链表的头结点
                for(int i=1;i<seg&&cur&&cur->next;i++){
                    cur=cur->next;
                }
                ListNode*head2=cur->next;//第二个链表的头结点
                cur->next=nullptr;//独立子序列1
                cur=head2;
                for(int i=1;i<seg&&cur&&cur->next;i++){
                    cur=cur->next;
                }
                ListNode*next=nullptr;
                if(cur){
                    // head2 的尾结点有可能为空，不为空时才能更新 next 指针，需要判断
                    next=cur->next;
                    cur->next=nullptr;//独立子序列2
                }
                // 合并子序列1和2
                pre->next=merge(head1,head2);
                // 移动 pre 指针到有序链表的尾结点
                while(pre->next){
                    pre=pre->next;
                }
                cur=next;// 从 next 开始继续归并
            }
        }
        return dummy->next;
    }
    // 快速排序
    void quick_sort_list(ListNode*head,ListNode*tail){
        if(head==tail||head->next==tail){
            return;
        }
        ListNode*i=head;// 初始化i，为链表第一个元素（最左边的元素），i及其i之前的都是小于等于
        ListNode*j=i->next;// 初始化j = i + 1
        int pivot=head->val;// 基准数字
        while(j!=tail){
            // 大循环条件，j不能超过链表长度
            // 如果 j 指向的值大于等于基准数字（如果比基准大，直接跳过）
            // 否则，j 指向的值小于基准，则交换
            // i++; swap;j++;
            if((j->val)<pivot){
                i=i->next;//i之后第一个大的
                swap(i->val,j->val);
            }
            j=j->next;
        }
        // 将pivot放到最终位置也就是i
        // i的左边都是小于nums[i]的数字，右边都是大于nums[i]的数字
        swap(head->val,i->val);
        // 接下来，对左边和右边分别排序，不断递归下去，直到元素全部有序
        quick_sort_list(head,i);
        quick_sort_list(i->next,tail);
    }
};
