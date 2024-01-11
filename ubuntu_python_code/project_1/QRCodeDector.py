#Import libraries
import cv2

import numpy as np
import matplotlib.pyplot as plt
from dataPath import DATA_PATH

# Read image in
img = cv2.imread(DATA_PATH+"IDCard-Satya.png", cv2.IMREAD_COLOR)

# Create QRCodeReader instance
qrDecoder = cv2.QRCodeDetector()
opencvData, bbox, rval = qrDecoder.detectAndDecode(img)
print(opencvData)

cv2.imshow("Original Image",img)
newImage = img.copy()

n = len(bbox)
for i in range(n):
    dlen = len(bbox[i])
    endPt = 0
    for j in range(dlen):
        endPt = j + 1
        if endPt < dlen:
            print(bbox[i][j], bbox[i][endPt])
            newImage = cv2.line(newImage,(bbox[i][j][0].astype(int), bbox[i][j][1].astype(int)),
                           (bbox[i][endPt][0].astype(int), bbox[i][endPt][1].astype(int)),
                           (255, 0, 0),5, cv2.LINE_AA)
        else:
            newImage = cv2.line(newImage,(bbox[i][j][0].astype(int), bbox[i][j][1].astype(int)),
                           (bbox[i][0][0].astype(int), bbox[i][0][1].astype(int)),
                           (255, 0, 0),5, cv2.LINE_AA)

#lt_tuple = (np.int32(points[0][0][0]), np.int32(points[0][0][1]))
#br_tuple = (np.int32(points[0][2][0]), np.int32(points[0][2][1]))

#cv2.rectangle(newImage, lt_tuple, br_tuple, (255,0,0), thickness=5,
#              lineType=cv2.LINE_8)

cv2.imshow("Copied Image", newImage)
cv2.imwrite(DATA_PATH+"QRCode-Output.png", newImage)

# Now try and write my own badge

while True:
    c = cv2.waitKey(20)
    if c == 27:
        break

cv2.destroyAllWindows()