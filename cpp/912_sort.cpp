class Solution {
public:
    vector<int> sortArray(vector<int>& nums) {
        // solve 1.1 headSort by hand
        // heapSort(nums);
        // solve 1.2 headSort by c++'s algorithms
        // make_heap(nums.begin(),nums.end());// default is big heap
        // sort_heap(nums.begin(),nums.end());
        // solve 2 insert sort, may be time out
        // insertSort(nums);
        // solve 3 shell sort
        // shellSort(nums);
        // solve 4 select sort
        //selectSort(nums);
        // solve 5 swap sort
        // swapSort(nums);
        // solve 6 quick sort
        quickSort(nums);
        return nums;
    }
private:
    // 1 堆排序
    void constructHeap(vector<int>&nums){
        // 从最后一个非叶子节点len/2-1开始调整，从下到上
        for(int i=nums.size()/2-1;i>=0;i--){
            adjustHeap(nums,i,nums.size());
        }
    }
    void adjustHeap(vector<int>&nums,int start,int end){
        // 每次调整使得满足堆定义
        int child;
        int i=start;
        while(1){
            child=2*i+1;
            if(child>end-1){
                break;
            }
            if(child+1<=end-1&&nums[child]<nums[child+1]){
                // find bigger child
                child+=1;
            }
            if(nums[i]<nums[child]){
                swap(nums[i],nums[child]);
                i=child; //交换后，以child为根的子树不一定满足堆定义，所以从child处开始调整
            }else{
                break;
            }
        }
    }
    void heapSort(vector<int>&nums){
        // 构建大顶堆
        constructHeap(nums);
        // 把根节点跟最后一个元素交换位置，调整剩下的n-1个节点，即可排好序
        for(int i=nums.size()-1;i>=0;i--){
            swap(nums[0],nums[i]);
            adjustHeap(nums,0,i);
        }
    }
    // 2 插入排序。在已排序的序列基础上从未排序序列选出一个插入排序序列
    // 一直往前找直到找到它该插入的位置。如果遇见一个与插入元素相等的，那么把待插入的元素放在相等元素的后面。因此是稳定排序
    void insertSort(vector<int>&nums){
        int n=nums.size();
        int i,j;
        for(i=1;i<n;i++){
            //刚开始这个有序的小序列只有1个元素，也就是第一个元素
            //a[0]...a[i-1]已经是排序好
            //a[i]...a[n-1]是未排序，先选择a[i]插入排序好序列
            //然后依次从未排序序列拿一个
            int tmp=nums[i];
            j=i-1;
            //从排序序列找到插入的位置，否则往后挪位置
            while(j>=0&&nums[j]>tmp){
                //a[j]>tmp 比较大 往后挪，给待插入值腾位置
                nums[j+1]=nums[j];
                j--;
            }
            nums[j+1]=tmp;
        }
    }
    // 3 shell排序，升级版插入排序，核心是选择一个增量d（如原始数组的一半），对相距d距离的所有元素进行插入排序，完成一次排序后只是说明间隔d的元素有序，缩小增量d，直至d=1才说明完全有序
    // 在不同的插入排序过程中，相同的元素可能在各自的插入排序中移动，最后其稳定性就会被打乱，故不稳定
    void shellSort(vector<int>&nums){
        int n=nums.size();
        int d=n/2;
        int i,j;
        while(d>0){
            //对相距d距离的所有元素进行插入排序
            for(i=d;i<n;i++){
                //刚开始这个有序的小序列只有1个元素，也就是第一个元素
                //a[0]...a[i-d]已经是排序好，间隔d
                //a[i]...a[n-1]是未排序，间隔d，先选择a[i]插入排序好序列
                int tmp=nums[i];
                j=i-d;
                while(j>=0&&nums[j]>tmp){
                    nums[j+d]=nums[j];
                    j-=d;
                }
                nums[j+d]=tmp;
            }
            d=d/2;
        }
    }
    // 4 选择排序，a[0]...a[i-1]是已排序数组，从a[i]...a[n-1]未排序中选出最小值作为a[i]，直接放在数组a[0]-a[i-1]最后位置
    // 在一趟选择中，当前元素a1比后面元素b要大，且当前元素一样的值a2也出现在了b的前面，那样交换位置后a1/a2位置也交换了所以不稳定
    void selectSort(vector<int>&nums){
        int n=nums.size();
        int i,j;
        //n-1轮排序，每一轮选出一个最小值，最后一个元素就是最大值故不需要n轮
        for(i=0;i<n-1;i++){
            //某一轮i，从a[i]开始选择，从a[i]...a[n-1]中选出最小值nums[min_k]
            int min_k=i;
            for(j=i;j<n;j++){
                if(nums[j]<nums[min_k]){
                    min_k=j;
                }
            }
            //最小值nums[min_k]与当前值nums[i]交换
            int tmp=nums[i];
            nums[i]=nums[min_k];
            nums[min_k]=tmp;
        }
    }
    //5 冒泡排序，直接对整个无序数组进行处理，每趟对相邻关键字进行比较和位置交换，一趟之后使得最值如气泡一般浮到最后位置
    //相邻元素交换，不会改变两个相等元素的相对位置故稳定
    void swapSort(vector<int>&nums){
        int n=nums.size();
        int i,j;
        //n-1轮排序，每一轮最大值都会漂浮到最后，最后一个元素已经到合适位置故不需要n轮
        for(i=0;i<n-1;i++){
            //第i轮，需比较交换到n-1-i
            for(j=0;j<n-1-i;j++){
                if(nums[j]>nums[j+1]){
                    int tmp=nums[j];
                    nums[j]=nums[j+1];
                    nums[j+1]=tmp;
                }
            }
        }
    }
    // 6快速排序。基本思想是n个关键字中选取一个作为pivot，然后对剩下n-1个元素进行分类，所有小的放在前面区间A，
    // 大的放在后面区间B，此时pivot其实已经完成最终位置的排序。接下来就是对A和B区间也采用类似思路，直到每个区间长为1.
    // 选择pivot算法：
    // 1 选第一个元素，若输入是顺序了，则达不到1分2效果，A区间长的0，B区间长度N,效能差，pass
    // 2 随机选择，这里特别的设计了3数中值法，
    // 2.1 选择pivot。选择序列头尾和中间位置的3个元素的中值作为pivot，且在选择后，把这3个位置数值进行排序，避免后续再排序
    // 2.2 基于该pivot分割得A和B。i为序列开始，j为序列结束:
    //   i不断右移动，跳过小于pivot元素，同时j左移，跳过大于pivot元素
    //   当i<j时，i指向大于pivot而j指向小于pivot元素，交换。
    //   继续上述操作，直到i>=j
    //   将pivot与i最后到达的位置进行交换，这个位置也就是pivot最终排序的位置
    // 2.3 对剩下的A和B区间递归调用
    // 在进行pivot与i最后到达的位置交换时很可能就破坏了相等元素顺序性故不稳定
    int getPivot(vector<int>&nums,int low,int high){
        // 3数中值法
        int mid=(low+high)/2;
        int tmp;
        // 3个元素的中值作为pivot，且在选择后，把这3个位置数值进行排
        if(nums[low]>nums[mid]){
            // low<mid
            tmp=nums[low];
            nums[low]=nums[mid];
            nums[mid]=tmp;
        }
        if(nums[low]>nums[high]){
            // low<high
            tmp=nums[low];
            nums[low]=nums[high];
            nums[high]=tmp;
        }
        if(nums[mid]>nums[high]){
            // mid<high
            tmp=nums[mid];
            nums[mid]=nums[high];
            nums[high]=tmp;
        }
        //将选择中值此时也即中间位置值作为基元并交换到序列最后第二位置方便后续排序操作
        tmp=nums[mid];
        nums[mid]=nums[high-1];
        nums[high-1]=tmp;
        return nums[high-1];
    }
    void part(vector<int>&nums,int low,int high){
        int i,j;
        int pivot;
        if(low<high){
            // 选择pivot
            pivot=getPivot(nums,low,high);
            // 基于该pivot分割得A和B。i为序列开始，j为序列结束
            i=low;//low位置已经在getPivot提前设为较小值故从low+1开始
            j=high-1;//high已经在getPivot提前设为较大值，high-1位置已经是pivot故从high-2开始
            while(1){
                while(j>i&&nums[++i]<pivot){
                }
                while(j>i&&nums[--j]>pivot){
                    
                }
                if(i<j){
                    //当i<j时，说明i指向大于pivot而j指向小于pivot元素，交换
                    int tmp=nums[i];
                    nums[i]=nums[j];
                    nums[j]=tmp;
                }else{
                    break;
                }
            }
            //将pivot与i最后到达的位置进行交换，这个位置也就是pivot最终排序的位置
            int tmp=nums[i];
            nums[i]=nums[high-1];
            nums[high-1]=tmp;
            // 对剩下的A和B区间递归调用
            part(nums,low,i-1);
            part(nums,i+1,high);
        }
    }
    void quickSort(vector<int>&nums){
        part(nums,0,nums.size()-1);
    }
};
