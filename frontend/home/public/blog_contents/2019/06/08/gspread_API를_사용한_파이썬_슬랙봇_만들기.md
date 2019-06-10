# [Python] gspread API를 활용한 파이썬 슬랙봇 만들기

## 시작은 포비님의 한 마디 말로부터
___

현재 우아한 형제들의 개발자 양성학교인 [우아한 테크코스](https://woowacourse.github.io/) 1기로 교육을 받고 있습니다. 벌써 이 글을 작성하는 시점 기준으로는 교육을 받기 시작한지 5주 정도 지났네요. 매주 한 번 교육생들과 포비(a.k.a. 자바지기)님의 질답 시간이 있는데, 이 토이 프로젝트의 시작은 포비님의 말 한 마디로부터 시작되었습니다.  

잠실 주변 맛집들을 알려주실 수 있느냐는 질문에 포비님은 '현재 맛집 리스트를 모으고 있는 중이다. 조만간 공개할 예정이다. 혹시 아는가, 누가 점심 추천 **슬랙봇**이라도 만들어줄지...'

## 점심 추천 슬랙봇
___

**슬랙봇...!** 코스 내 소통공간으로 [슬랙(Slack)](https://slack.com/intl/en-kr/)을 사용하고 있는데, 예전에 광일공방에서 텔레그램 봇을 파이썬으로 만들어보았던 가물가물한 감을 살려서 한 번 도전해볼 수 있겠다 싶었습니다.  

준비된 식당 리스트를 구글 스프레드 시트에 넣어놓고, 서버에서 내용을 주기적으로 긁어와서 슬랙봇 API로 식당 4개 정도를 랜덤으로 뿌려주면 되겠다 싶었습니다. 그리고 어떤 식당에 대한 사용자의 reaction emoji를 받을 수 있다면 추천/비추천 시스템도 만들 수 있을 것 같다 생각했습니다.

*하지만 탁상공론 입코딩의 결과는 늘 커다란 댓가를...*  

같이 공부하고 있는 덕수 선생님과 함께 둘이서 개발을 진행하였습니다.  

## gspread
___

먼저 식당들을 정리하기 위해 구글 스프레드 시트를 생성합니다. ~~*갑자기 컴활1급 5수의 추억이*~~  
저 같은 경우는 식당 이름, 종류(한식 등), 대표 메뉴, 대표 메뉴 가격, 네이버 플레이스 url, 좋아요, 별로에요 등의 column들을 두었습니다. 사용하고자 하는 내용에 맞게 table을 구성해줍니다.  

찾아보니 [gspread](https://github.com/burnash/gspread)라는 파이썬 라이브러리가 있습니다. 해당 구글 시트에 접근하기 위해서 OAuth 2.0 권한을 받아와야 하는데, [gspread docs](https://gspread.readthedocs.io/en/latest/oauth2.html)에 자세한 방법이 나와있습니다. [Google Developers Console](https://console.developers.google.com/)에서 프로젝트를 생성하고 Service Account Key를 만들면 JSON 파일로 secret key를 줍니다. (당연한 말이지만 공개된 github 등의 장소에 업로드되지 않도록 주의하세요! `.gitignore` 설정은 생활화!)  

```javascript
{
    "private_key_id": "2cd … ba4",
    "private_key": "-----BEGIN PRIVATE KEY-----\nNrDyLw … jINQh/9\n-----END PRIVATE KEY-----\n",
    "client_email": "473000000000-yoursisdifferent@developer.gserviceaccount.com",
    "client_id": "473 … hd.apps.googleusercontent.com",
    "type": "service_account"
}
```

JSON 파일을 열어보면 `client_email`이라는 이름으로 어떤 이메일 주소가 들어가 있습니다. 그 것이 바로 슬랙봇이 입고 구글시트에 접근할 주소입니다. 스프레드 시트에서 해당 email에게 공유를 걸어 접근 권한을 줍니다.  

그리고 한 두 시간동안 헤맸던 부분인데, [Google Drive Console](https://console.cloud.google.com/)에서 Google Drive API를 enable 해주셔야 합니다. 처음에 Sheet API만 enable 해주면 되는줄 알고... 삽질 ٩( ᐛ )و  

이제 터미널을 열어 다음 세 가지 라이브러리를 설치합니다.

```bash
pip install gspread
pip install --upgrade oauth2client
pip install PyOpenSSL
```

그리고 JSON key file로 구글 시트에 접근할 수 있는 소스를 다음과 같이 작성합니다.  

```python
import gspread
from oauth2client.service_account import ServiceAccountCredentials

scope = ['https://spreadsheets.google.com/feeds',
         'https://www.googleapis.com/auth/drive']

credentials = ServiceAccountCredentials.from_json_keyfile_name(JSON_KEY_FILE, scope)

gc = gspread.authorize(credentials)

wks = gc.open(SHEET_NAME).sheet1
```

worksheet를 이름 혹은 sheet number로 받아오고 나면, 필요한 함수들을 사용하여 데이터를 읽고 쓸 수 있습니다. 대표적인 함수들로는 다음과 같은 것들이 있습니다.  

```python 
# 1행에 있는 모든 값들 가져오기 (1행의 index는 1)
values_list = worksheet.row_values(1)

# 1열에 있는 모든 값들 가져오기 (1열의 index는 1)
values_list = worksheet.col_values(1)

# list의 list로 모든 데이터 가져오기
list_of_lists = worksheet.get_all_values()

# string 값으로 검색
cell = worksheet.find("Dough")

# 'B1' cell update하기
worksheet.update_acell('B1', 'Bingo!')
worksheet.update_cell(1, 2, 'Bingo!')

# 범위 지정해서 가져오기
cell_list = worksheet.range('A1:C7')
```

지금 소개해드린 것들은 많은 함수들 중 매우 일부분이기 때문에 [gspread github](https://github.com/burnash/gspread), 또는 더 상세한 [gspread docs](https://gspread.readthedocs.io/en/latest/index.html)를 참고하시는 것이 좋습니다.  

## 슬랙봇 생성하기
___

![슬랙앱 생성](http://cdn.wbluke.com/lunch_bot_image/slackBot1.png){width=600px}

먼저 Workspace에서 사용할 슬랙봇을 만들어봅시다. [Slack API](https://api.slack.com/)에서 `Start Building`을 누르면 새로운 봇을 만들 수 있습니다. 어플리케이션의 이름과 사용할 Workspace를 지정해 줍니다.  

![Bot User 추가](http://cdn.wbluke.com/lunch_bot_image/slackBot2.png){width=600px}

초기 페이지에서는 봇의 대표 이미지와 간단한 설정들을 지정해 줄 수 있습니다. 왼쪽에 Bot User로 가서 `Add Bot User`를 해줍니다. 이름을 지정할 수 있고, 항상 온라인으로 보여줄 것인지 정할 수 있습니다. 제 봇은 언제든지 사용할 점심 추천봇이기 때문에 설정을 On으로 해 두었습니다.  

![Workspace에 설치](http://cdn.wbluke.com/lunch_bot_image/slackBot3.png){width=600px}

이제 Install App으로 가서 우리의 Workspace에 만든 봇을 설치합니다. 성공적으로 설치가 되면, `xoxb`로 시작하는 Bot User OAuth Access Token을 보여줍니다. 당연히 유출은 조심하셔야 하고, 사용을 위해 복사해서 잘 가지고 있도록 합시다.  


## slack API
___

이번엔 Slack과 소통하기 위해 Slack API를 알아봅니다. 터미널에서 다음과 같이 설치합니다.  

```bash
pip3 install slackclient
```

개발 과정에서 알게 된 사실인데, slackclient가 2.0.0 버전으로 업그레이드 된 이후에는 사용법이 완전히 달라졌습니다. 검색했을 때 이전 버전의 블로그 글이나 문서가 꽤 많아서 고생했던 기억이 있네요.  

2.0.0 대 버전에서는 사용법이 전보다 훨씬 간단합니다. 데코레이터 기반으로 모든 동작들을 처리할 수가 있는데, 블랙박스 처리가 되어서 단순히 사용하기엔 좋지만, 내부 동작 과정을 모르기 때문에 깊게 활용하기가 어려운 점은 있습니다. 기회가 되면 내부 코드를 한 번 따라가면서 공부해보고 싶습니다.  

스크립트를 실행시키기 전에 터미널에서 다음과 같이 환경변수로 아까 획득한 봇의 토큰을 세팅합니다.  

```bash
export SLACK_BOT_TOKEN='xoxb-어쩌구-저쩌구-조쩌구'
```

스크립트의 기본적인 구조는 다음과 같습니다.  

```python
import slack

@slack.RTMClient.run_on(event='message')
def recommend(**payload):
    data = payload['data']
    web_client = payload['web_client']
    channel_id = data['channel']
    
    if data['text'].strip() == '?':
        # Do something with web_client and channel_id
        return

    if '밥!' in data['text']:
        # Do something with web_client and channel_id
        return
        
if __name__ == "__main__":
    slack_token = os.environ["SLACK_BOT_TOKEN"]
    rtm_client = slack.RTMClient(token=slack_token)
    rtm_client.start()
```

main에서 환경변수에 먹인 토큰을 받아와서, `rtm_client`에 세팅한 뒤 start()를 시킵니다. 그 이후에는 각 함수에 달려있는 데코레이터의 인자에 따라서 각 동작이 처리가 됩니다. `run_on(event='message')`에 있는 event 인자에 따라서 처리하는 동작을 지정할 수 있는데요. 위 경우는 사용자가 보낸 메시지를 감지해서 어떤 일을 처리하는 함수입니다. 저희는 `event='message'` 이외에도 사용자의 이모지 반응을 캐치해서 식당 추천기능으로 사용하기 위해 `event='reaction_added'`, `event='reaction_removed'` 등을 사용하였습니다. 더 많은 기능은 [Real Time Messaging API](https://api.slack.com/rtm)에서 확인하실 수 있습니다.  

함수가 실행되면, payload를 통해 각종 정보들이 들어오게 되는데, 필요한 정보들을 가공하여 하고 싶은 일들을 처리할 수 있습니다. 점심봇인 만큼 '?'를 통해 사용설명서를 보거나, '밥!'이라는 키워드를 통해 식당 메뉴 등을 추천받을 수 있도록 했습니다. 이 때 다시 메시지를 사용자에게 보내기 위해 `web_client`와 `channel_id`를 사용합니다.  

```python
web_client.chat_postMessage(
        channel=channel_id,
        attachments=[
            {
                'text': '<' + restaurant.get_naver_place_addr() + '|' + restaurant.get_name() + '>',
                'fields': [
                    {
                        'title': '대표 메뉴',
                        'value': restaurant.get_popular_menu() + ' ' + str(restaurant.get_price_of_popular_menu()) + '원',
                        'short': True
                    },
                    {
                        'title': '추천 정보',
                        'value': ':thumbsup: '+ str(restaurant.get_good()) + '   :thumbsdown: ' + str(restaurant.get_bad()),
                        'short': True
                    }
                ],
                'color': restaurant_color,
                'thumb_url': restaurant_thumb_url
            }
        ]
    )

```

`web_client`의 `chat_postMessage` 함수를 사용하면 사용자에게 메시지를 보낼 수 있습니다. 이 때 slack에서 지정해놓은 attachments의 옵션들을 통해서 메시지를 예쁘게 단장할 수 있습니다. 예쁜거 좋아요. 실제로 위 설정을 통해서 다음과 같이 UI를 구성할 수 있습니다. (restaurant는 식당 정보를 들고 있는 객체인데, 코드 보시기에 헷갈리지 않게 restaurant를 받아오는 과정을 위 코드에서는 생략했습니다.)  

![예뻐야 합니다](http://cdn.wbluke.com/lunch_bot_image/slackBot4.png){width=600px}

실제로 덕수 선생님이 제가 촵촵촵촵 만들어 온 위 UI를 보더니 `프론트 개발자 지망인데 여기 잘못 온거 아닙니까?` 라고 했습니다. 저는 프론트 백 둘 다 잘하고 싶어요 쫒아내지 마세요...  

attachments의 사용은 [slack API message-attachments](https://api.slack.com/docs/message-attachments) 페이지에서 자세하게 보실 수 있습니다. 실제로 프로덕션 코드에 사용하기 전에 웹에서 미리 적용시켜보고 사용할 수 있습니다.  

## 초기 모델의 구성과 문제점
___

![초기 슬랙봇 모델](http://cdn.wbluke.com/lunch_bot_image/slackBot5.png){width=600px}

기존에 머릿속으로 생각한 모델은 위와 같았습니다. gspread로 시트에서 가져와서, 봇이 slack에 뿌려주고, 사용자 반응을 이모지로 받아서 다시 gspread에 저장하고... 그런데 실제로 개발을 진행해보니 여러가지 문제가 있었습니다.  

가장 큰 이슈는 속도 문제였습니다. 서버와 구글 스프레드가 통신하는 시간도 길고, 서버와 슬랙이 통신하는 시간도 길었습니다. 점심봇은 사용자 명령에 최대한 빠르게 응답을 주어야 하는데, 생각 이상으로 데이터가 오고가는 시간이 길었습니다. 그래서 여러가지 고민과 실험 끝에 다음과 같이 모델을 변경하였습니다.  

## 개선한 모델
___

![개선한 슬랙봇 모델](http://cdn.wbluke.com/lunch_bot_image/slackBot6.png){width=600px}

초기에 서버를 시작하면 RestaurantRepo(이하 Repo) 내부에 있는 Gspread Client가 시트에서 모든 식당정보를 긁어와 Repo 안에 저장합니다. 그리고 그 이후에는 봇이 슬랙과만 소통하면서 최신 정보를 Repo에 반영합니다. 그리고 시트와 Repo 간의 싱크는 3분 정도 주기로 update 및 fetch를 하여 데이터 정합성이 깨질 우려를 최대한 줄여줍니다. 이런 구조로 접근하니 속도 측면에서 나은 성과를 얻을 수 있었습니다.  

## 아직 진행중인 프로젝트
___

덕수 선생님과 둘이 아무것도 모르는 상태에서 뚝딱뚝딱 만들다 보니 아직 시도해봐야 할 것도, 수정해야 할 것도 넘쳐납니다. 기능적으로도 다양한 추천 로직을 추가해야하고, 서버가 최대한 죽지 않도록 여러 Thread를 관리하고 예외 처리를 해야 합니다. 또한 예외가 발생했을 때의 상황 기록을 위한 Logger도 공부하고 적용해야 합니다. 공부하면 공부할수록 공부해야할 것들이 `exponential`하게 증가합니다. ~~*개발하면서 제일 많이 한 말 : 살려주세요..*~~  

그래도 이렇게 자발적으로 즐거움 속에서 개발하는 것만큼 빠른 성장을 이루어내는 것은 없는 것 같습니다. 목표는 남은 코스 과정 7개월 동안 무사히 서비스를 진행해보는 것입니다. 서버 개발자가 되기 위한 작은 발걸음을 이제 한 걸음 뗀 느낌입니다. :)

