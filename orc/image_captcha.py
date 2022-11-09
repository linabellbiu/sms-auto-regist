# 图片验证码
import ddddocr
import sys
path = sys.argv[1]
ocr = ddddocr.DdddOcr(show_ad=False)
with open(path, 'rb') as f:
    img_bytes = f.read()
res = ocr.classification(img_bytes)
print(res)