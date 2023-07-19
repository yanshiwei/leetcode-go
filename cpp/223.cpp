class Solution {
    /*
    给你 二维 平面上两个 由直线构成且边与坐标轴平行/垂直 的矩形，请你计算并返回两个矩形覆盖的总面积。
每个矩形由其 左下 顶点和 右上 顶点坐标表示：
第一个矩形由其左下顶点 (ax1, ay1) 和右上顶点 (ax2, ay2) 定义。
第二个矩形由其左下顶点 (bx1, by1) 和右上顶点 (bx2, by2) 定义。
    */
public:
    int computeArea(int ax1, int ay1, int ax2, int ay2, int bx1, int by1, int bx2, int by2) {
        // 求覆盖区域关键点
        int x1=max(ax1,bx1);
        int y1=max(ay1,by1);
        int x2=min(ax2,bx2);
        int y2=min(ay2,by2);
        return area(ax1,ay1,ax2,ay2)+area(bx1,by1,bx2,by2)-area(x1,y1,x2,y2);
    }
private:
    int area(int x1,int y1,int x2,int y2){
        // 如果没有覆盖则为负数
        int l=x2-x1;
        int h=y2-y1;
        if(l<=0||h<=0){
            return 0;
        }
        return l*h;
    }
};
