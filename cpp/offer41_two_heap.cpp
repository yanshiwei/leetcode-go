class MedianFinder {
    /*
    小于中位数的TOP K1个数字存储在大顶堆，大于中位数的TOP K2个数字存储在小顶堆；两个堆元素大小差距最大为1，当两个堆大小为差距1时，元素个数多的那个堆的堆顶元素即为中位数；两个堆元素个数相同时，两个堆顶元素的值取平均。
    关键是数据流入一个数字x时操作：
    1.当大顶堆maxHeap为空时候，直接入大堆(即将第一个元素放入到最大堆)；
    2.maxHeap.size()==minHeap.size()，判断大堆堆顶元素和x大小，若x更大，则x直接入小堆；若x更小，则直接入大堆；
    3.maxHeap.size()>minHeap.size()，判断大堆堆顶元素和x大小，若x更大，则x直接入小堆；若x更小，为了保证两个堆大小差距不大于1，先pop出大堆里最大元素放入小堆，然后入大堆；
    4.maxHeap.size()<minHeap.size()，判断小堆堆顶元素和x大小，若x更小，则x直接入大堆；若x更大，为了保证两个堆大小差距不大于1，先pop出小堆里最小元素放入大堆，然后入小堆；
    */
public:
    /** initialize your data structure here. */
    MedianFinder() {

    }
    
    void addNum(int num) {
        if(maxHeap.size()<1){
            maxHeap.push(num);
            return;
        }
        if(maxHeap.size()==minHeap.size()){
            if(maxHeap.top()<num){
                minHeap.push(num);
            }else{
                maxHeap.push(num);
            }
            return;
        }
        if(maxHeap.size()>minHeap.size()){
            if(maxHeap.top()<num){
                minHeap.push(num);
            }else{
                int tmp=maxHeap.top();
                maxHeap.pop();
                minHeap.push(tmp);
                maxHeap.push(num);
            }
            return;
        }   
        if(maxHeap.size()<minHeap.size()){
            if(minHeap.top()<num){
                int tmp=minHeap.top();
                minHeap.pop();
                maxHeap.push(tmp);
                minHeap.push(num);
            }else{
                maxHeap.push(num);
            }
            return;
        }      
    }
    
    double findMedian() {
        double res=0.0;
        int cnt=maxHeap.size()+minHeap.size();
        if(cnt<1){
            return res;
        }
        if((cnt&1)==1){
            // 奇数
            if(maxHeap.size()>minHeap.size()){
                res=double(maxHeap.top());
            }else{
                res=double(minHeap.top());
            }
        }else{
            // 偶数
            res=double(minHeap.top()-maxHeap.top())/2+maxHeap.top();
        }
        return res;
    }
private:
    priority_queue<int,vector<int>,less<int>>maxHeap;
    priority_queue<int,vector<int>,greater<int>>minHeap;
};

/**
 * Your MedianFinder object will be instantiated and called as such:
 * MedianFinder* obj = new MedianFinder();
 * obj->addNum(num);
 * double param_2 = obj->findMedian();
 */
