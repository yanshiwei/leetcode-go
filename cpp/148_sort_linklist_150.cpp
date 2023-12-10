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
    1. 快速排序
    思想：
    快速排序的思想大家应该都知道，就是找一个pivot，把小于pivot分在一边，大于等于pivot的分在另外一边
这个过程也叫做partition分区，数组的分区很好做，左右两个指针i/j，不断交换就好了
链表因为只能单向遍历，所以要换一种partition方法。
    做法：
    目的是使得左边的值都小于pivot，右边的值都不小于pivot 所以用一个索引记录左边的坐标，遍历过程中，
    每次碰到比pivot小的，都要交换一下，放到左边。 遍历完成后，再把pivot交换到中间来，这样就达成了目的
    题解：
    https://blog.csdn.net/u010429424/article/details/77776731
    https://leetcode.cn/problems/sort-list/solutions/9318/gui-bing-pai-xu-he-kuai-su-pai-xu-by-a380922457/
    2. 归并排序
    思想：
    鉴于快速排序超时，引入归并排序，归并排序是一种分而治之的算法。
    其思想是将原始数组切割分成较小的数组，直到每个数组中只有一个元素，接着将小数组归并成较大的数组，
    直到最后只有一个排序完毕的大数组。
    做法：
    归并排序有自顶向下的递归方法，和自底向上的迭代方法
    （1）自顶向下的归并排序，递归到只剩一个结点或者为空链表，返回的时候返回合并后的链表头结点。
     时间复杂度 O(nlogn) 空间复杂度 O(logn)。首先需要找到中间点将两个链表一分为二，再分两部分进行递归排序，最后使用合并两个有序链表的方法合并两部分。一分为二时使用快慢指针找到中点区分。
     (2) 自底向上的归并排序,先两个两个的 merge，完成一趟后，再 4 个4个的 merge，直到结束
     题解：
     https://leetcode.cn/problems/sort-list/solutions/1525677/by-lafiteee-7vi7/
    */
public:
    ListNode* sortList(ListNode* head) {
        // quick_sort_list(head,nullptr);
        // return head;
        return merge_sort_undfs(head);
    }
private:
    // 归并排序
    // 2.1 基于快慢指针找链表中点实现1
    ListNode*findMid(ListNode*head){
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
    // 2.2 基于快慢指针找链表中点实现1
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
    // 2.3 两个链表合并
    ListNode*merge(ListNode*head1,ListNode*head2){
        ListNode*dummy=new ListNode();
        ListNode*pre=dummy;
        while(head1&&head2){
            if(head1->val<head2->val){
                pre->next=head1;
                head1=head1->next;
            }else{
                pre->next=head2;
                head2=head2->next;
            }
            pre=pre->next;
        }
        pre->next=head1?head1:head2;
        return dummy->next;
    }
    // 2.3 合并，递归实现
    ListNode*merge_sort_dfs(ListNode*head){
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        ListNode*mid=findMid(head);
        ListNode*head2=mid->next;
        //断开链接单独成一个链表
        mid->next=nullptr;
        // 递归知道单个节点或者空
        ListNode*left=merge_sort_dfs(head);
        ListNode*right=merge_sort_dfs(head2);
        // 合并
        return merge(left,right);
    }
    // 2.4 合并，非递归实现
    ListNode*merge_sort_undfs(ListNode*head){
        if(head==nullptr||head->next==nullptr){
            return head;
        }
        // 求链表长度
        int n=0;
        ListNode*cur=head;
        while(cur){
            n++;
            cur=cur->next;
        }
        ListNode*dummy= new ListNode();
        dummy->next=head;
        // 子链表长度，从 1 开始，每次翻倍
        for(int seg=1;seg<n;seg*=2){
            // 每次合并都从最左侧开始
            ListNode*pre=dummy;
            ListNode*cur=dummy->next;
            while(cur){
                // 子序列长度seg，合并2个这样的子序列
                // 先找到第一段[head,cur]长度最多为seg
                ListNode*head=cur;
                for(int i=1;i<seg&&cur&&cur->next;i++){
                    cur=cur->next;
                }
                ListNode*head2=cur->next;//第二个链表的头结点
                cur->next=nullptr;//独立子序列1
                // 再找到第一段[head2,cur]长度最多为seg
                cur=head2;
                for(int i=1;i<seg&&cur&&cur->next;i++){
                    cur=cur->next;
                }
                // 设置2个子段后的下一个节点
                ListNode*next=nullptr;
                if(cur){
                    next=cur->next;
                    cur->next=nullptr;//独立子序列2
                }   
                // 合并子序列1和2
                pre->next=merge(head,head2);// head2有可能空
                // 更新pre
                while(pre->next){
                    pre=pre->next;
                }
                // 从 next 开始继续归并
                cur=next;            
            }
        }
        return dummy->next;
    }
private:
    // 快速排序
    void quick_sort_list(ListNode*head,ListNode*tail){
        if(head==tail||head->next==tail){
            return;
        }
        ListNode*i=head;//初始化i，为链表第一个元素（最左边的元素），i及其i之前的都是小于等于
        ListNode*j=i->next;//初始化j = i + 1
        int pivot=head->val;//基准数字
        while(j!=tail){
            // 大循环条件，j不能超过链表长度
            // 如果 j 指向的值大于等于基准数字（如果比基准大，直接跳过）
            // 否则，j 指向的值小于基准，则交换
            if((j->val)<pivot){
                // i当前指向<=的故先更新
                i=i->next;
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
