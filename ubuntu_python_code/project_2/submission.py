import cv2
import numpy as np
import math

#global variables - we're using a single global to contain the upper left and lower right points
boxPoints=[]
isActive = False
tmpImage = None
source = np.zeros((1, 1, 3), dtype = "uint8")

# function for annotation drawing a square
def drawSquare(action, x, y, flags, userdata):
    global boxPoints
    global isActive
    global tmpImage
    global source
    if action == cv2.EVENT_LBUTTONDOWN:
        boxPoints = [(x,y)] #Set boxPoints[0] to upper left point
        #draw origination point to give user feedback
        cv2.circle(source, boxPoints[0], 1, (255,255,0), 2, cv2.LINE_AA )
        isActive = True
        tmpImage = source.copy()
    #
    elif (action == cv2.EVENT_MOUSEMOVE) & isActive:
        #actively draw box
        if len(boxPoints) == 2:
            boxPoints[1] = (x,y)
        else:
            boxPoints.append((x,y))
        source = tmpImage.copy()
        cv2.rectangle(source, boxPoints[0], boxPoints[1], (255, 255, 0), 2, cv2.LINE_AA)
        #cv2.imshow("Window", source)
    elif (action == cv2.EVENT_LBUTTONUP) & isActive:
        if (len(boxPoints) == 2):
            #cv2.rectangle(source, boxPoints[0], boxPoints[1], (255, 255, 255), 2, cv2.LINE_AA)
            boxPoints[1] = (x, y)
        else:
            boxPoints.append((x, y))
        source = tmpImage.copy()
        cv2.circle(source, boxPoints[0], 1, (255, 255, 255), 2, cv2.LINE_AA)
        cv2.rectangle(source, boxPoints[0], boxPoints[1], (255, 255, 0), 2, cv2.LINE_AA)
        #cv2.imshow("Window", source)
        isActive = False

source = cv2.imread("./data/images/sample.jpg",1)
cv2.imshow("Window", source)

dummy = source.copy()
cv2.namedWindow("Window")

cv2.setMouseCallback("Window", drawSquare)
k = 0
botOff = source.shape[1] - 30
while k != 27:

    cv2.imshow("Window", source)
    cv2.putText(source, 'Choose upper left corner and drag.',
                (10, 30), cv2.FONT_HERSHEY_SIMPLEX,
                0.7, (255, 255, 255), 2);
    cv2.putText(source, 'Press ESC to exit and c to clear.',
                (10, botOff), cv2.FONT_HERSHEY_SIMPLEX,
                0.7, (255, 255, 255), 2);

    k = cv2.waitKey(1) & 0xFF
    # Another way of cloning
    if k == 99:
        source = dummy.copy()

#Save cropped image
if len(boxPoints) < 2:
    print("No area was selected!")
else:
    face = dummy[boxPoints[0][1]:boxPoints[1][1],boxPoints[0][0]:boxPoints[1][0]]
    cv2.imwrite("./face.png", face)
cv2.destroyAllWindows()



