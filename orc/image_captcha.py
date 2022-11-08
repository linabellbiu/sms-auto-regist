# 图片验证码
import ddddocr

ocr = ddddocr.DdddOcr(show_ad=False)
with open('/Users/xudong/code/my/githup/sms-auto-regist/test/imgcode.jpeg', 'rb') as f:
    img_bytes = f.read()
res = ocr.classification(img_bytes)

print(res)