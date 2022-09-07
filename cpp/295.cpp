class MedianFinder {
    /*
    中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
    */
public:
    MedianFinder() {
        // 左部分小值 大顶堆存储，右部分大值 小顶堆存储，两个堆大小不超过1，暂定小堆大小最多超过1个
        // 总数是偶数时，取两个堆顶元素平均；总数是奇数时 取小顶堆顶部元素
    }
    
    void addNum(int num) {
        if(((min.size() + max.size()) & 1) == 0){
            // 总数是偶数时,优先放入小堆，但不是直接放，而是保证顺序，也就是当前值小于大堆堆顶值 先放入大堆重新堆化
            // 确保给小堆的是更大的
            if(max.size()>0&&num<max[0]){
                max.push_back(num);
                push_heap(max.begin(),max.end(),less<int>());
                num=max[0];
                // 去掉堆顶值
                pop_heap(max.begin(),max.end(),less<int>());
                max.pop_back();
            }
            min.push_back(num);
            push_heap(min.begin(),min.end(),greater<int>());
        }else{
            // 总数是奇数时,优先放入大堆，但不是直接放，而是保证顺序，也就是当前值大于小堆堆顶值 先放入小堆重新堆化
            // 确保给大堆的是更小的
            if (min.size()>0&&num>min[0]){
                min.push_back(num);
                push_heap(min.begin(),min.end(),greater<int>());
                num=min[0];
                pop_heap(min.begin(),min.end(),greater<int>());
                min.pop_back();
            }
            max.push_back(num);
            push_heap(max.begin(),max.end(),less<int>());
        }
    }
    
    double findMedian() {
        double median=0.0;
        int totalSize=max.size()+min.size();
        if(totalSize<1){
            return median;
        }
        
        if((totalSize&1)==1){
            median=double(min[0]);
        }else{
            median=double(min[0]-max[0])/2+max[0];
        }
        return median;
    }
private:
    vector<int>min;
    vector<int>max;
};

/**
 * Your MedianFinder object will be instantiated and called as such:
 * MedianFinder* obj = new MedianFinder();
 * obj->addNum(num);
 * double param_2 = obj->findMedian();
 */
