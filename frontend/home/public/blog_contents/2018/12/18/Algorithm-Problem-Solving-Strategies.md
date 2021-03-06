# [알고리즘] Algorithm Problem Solving Strategies

### 0. 알고리즘 문제 해결 전략
___
![프로그래밍 대회에서 배우는 알고리즘 문제해결전략](https://cdn.wbluke.com/Algorithmic_PS_Book.jpg){width=300px height=300px}  

도서 사이트 : <http://book.algospot.com/>  
채점받을 수 있는 알고스팟 사이트 : <https://algospot.com/>

### 1. 알고리즘 공부를 시작합니다.
___
꽤나 유명하고 많은 분들이 이미 공부하셨던 교재를 선택하였습니다.  
구종만 저자님의 **'프로그래밍 대회에서 배우는 알고리즘 문제해결전략'** ~~일명 종만북~~

사실 처음 책을 샀을 때는 C++을 접해보지 않았던 때여서, 책의 초반부를 유심히 읽다가 책 전체가 C++로 쓰였다는 것을 깨닫고는 바로 ~~기쁨으로 가득차~~ C++ 교재를 따로 사서 1회독 독학 후 공부하기 시작했습니다.  
(나름 Java로 OOP를 처음 접했던 상태여서 다행히 생각보다는 재미도 있고 수월했습니다.)

사실 이 책이 C++의 이해를 상당수준 요구한다기 보다는, 책에 나와있는 `Skeleton code`를 볼 수 있을 정도의 문법 이해도와, 많이 사용하는 STL의 사용법을 알고 있다면 문법적으로는 크게 이해하는 데에 문제되지 않는 책 같습니다. 정작 어려운 것은 내용이기에...

저도 제가 공부했던 C++책에는 STL설명이 없었기에 알고리즘 공부하면서 그때그때 STL이 나올 때마다 정리하면서 보고 있습니다.

### 2. 개발자를 지망하게 된 계기: 알고리즘
___
저는 학부전공이 수학입니다.  
수학과를 간 이유는 사실 중고등학교 학창시절 어려운 수학문제, 올림피아드 문제와 밤새 씨름하는 것을 즐거워했던 경험이 있어서였습니다.  
다만 수학과에 진학해보니 아시는 분들은 아시겠지만 제가 재미를 느꼈던 그런 `문제를 푼다는 것`의 내용이나 과정들은 전혀 보이지 않았고... 쉽게 흥미를 잃게 되었습니다.

그 재미를 일깨워준 것이 우연히 친구들과 나간 교내 프로그래밍 경시대회였습니다.  
코딩의 ㅋ도 모는 상태의 저와, 전공자 친구 1명, (지금은 멋지게 개발자로 살아가고 있는)컴공 지망자인 국어교육과 친구 1명.  
이렇게 셋이 나가게 된 대회에서 나름 중상위권의 성적을 내었고, 아직도 신선한 기억으로 남아있습니다.  
저는 나름대로 수학적으로 알고리즘만 짜서 나머지 둘에게 바쁘게 건네주었던 기억이 있네요.

개발자를 하기로 마음을 먹은 것은 그 후로 몇 년 뒤에 졸업할 때가 다 되어서이지만, 공부를 시작하고 보니 그 당시 접했던 이 **알고리즘** 이라는 분야는 확실히 한참 잊고있었던 `문제를 푼다는 것`에 대한 재미를 일깨워주었습니다.

### 3. 1회독을 목표로!
___
개발자의 기초체력이라 불리는 알고리즘이기도 하고, 또 흥미로운 분야인 만큼 기대를 갖고 뛰어들었으나,  
막상 제대로 파고드니 처음이라 내용이 정말 어렵더군요. 많은 분들의 기록들을 보니 1회독도 어렵더라...는 후문이.  
그래서 지금은 취업이 우선순위이긴 하지만, 취업 여부와 관련없이 꼭 1회독을 하기로 마음을 먹어봅니다.  
무사히 1회독을 하면서 또 좋은 취미로 남을 수 있지 않을까 하는 생각도 가져봅니다.  

### 4. 공부하는 순서
___
책 p11에 나와있는 `표 1.1 입문자를 위한 커리큘럼`에 따라 공부를 진행하려 합니다.  
표는 다음과 같습니다.  

| 장 번호 | 장 제목 |
|:---|:---|
|2    |문제 해결 전략|
|3|코딩과 디버깅|
|4|알고리즘의 시간 복잡도 분석|
|6|무식하게 풀기|
|7|분할 정복|
|8|동적 계획법|
|18|선형 자료 구조|
|19|큐와 스택, 데크|
|21|트리의 구현과 순회|
|22|이진 검색 트리|
|23|우선순위 큐와 힙|
|27|그래프의 표현과 정의|
|28|그래프의 깊이 우선 탐색|
|29|그래프의 너비 우선 탐색|
|30|최단 경로 알고리즘|
: 표 1.1 입문자를 위한 커리큘럼

위 표에 따라 순차적으로 공부한 후에 나머지 장들을 공부할 계획입니다.  

블로그 포스팅은 장 단위가 아닌 문제 단위로 포스팅 하려 합니다. 설명하는 장인 2,3,4장을 제외하고 6장부터 문제 별로 중점적인 아이디어나 필요한 문법들, 제가 생각한 풀이 방식들을 나름 기록하는 느낌으로 남기려고 합니다.  

공부 방식은 먼저 `Github`에서 `Issue Tracker`방식으로 매일매일 `commit`기록을 남기면서 한 문제를 여러 날에 걸쳐서 공부하는 방식을 택했습니다.  
~~어차피 못풀고 답을 보게 될테니~~ 풀이를 한번 보고 이해했다고 기억에 남지는 않을테니까요. 반복학습이 중요할 듯 싶어서-.
풀이가 완료된 문제들에 대해서 `Issue` 기록을 참고하여 포스팅을 하려고 합니다.  
부디 꾸준히 공부할 수 있기를-!
