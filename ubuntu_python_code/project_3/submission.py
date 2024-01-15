import cv2
import numpy as np

maxScaleUp = 100
scaleValue = 1
scaleType = 0
maxType = 1
scaleFactor = 1.0
myShape = (200, 300, 3)
stp = np.zeros(myShape, dtype=np.uint8)
stp[:] = (128, 128, 128)
origstp = stp.copy()

windowName = "Resize Image"
trackbarValue = "Scale"
trackbarType = "Type: \n 0: Scale Up \n 1: Scale Down"


def show_stype(sftype):
    global stp
    global origstp
    stp = origstp.copy()
    if sftype == False:
        text = "Scale UP"
    else:
        text = "Scale Down"
    fontScale = 1.5
    fontFace = cv2.FONT_HERSHEY_COMPLEX
    fontColor = (255, 255, 255)
    fontThickness = 2
    cv2.putText(stp, text, (5, 100), fontFace, fontScale, fontColor, fontThickness, cv2.LINE_AA);



# load an image
im = cv2.imread("./data/images/truth.png")
scale_window = "ScalingFactor"
#stp = cv2.cvtColor(stp, cv2.COLOR_BGR2GRAY)
cv2.namedWindow(
        scale_window,
        flags=( cv2.WINDOW_GUI_NORMAL | cv2.WINDOW_FREERATIO | cv2.WINDOW_AUTOSIZE))
cv2.moveWindow(scale_window, 640,100)
# Create a window to display results
cv2.namedWindow(windowName, cv2.WINDOW_AUTOSIZE)
show_stype(False)
cv2.imshow(scale_window, stp)

# Callback functions
def scaleTypeImage(*args):
    global scaleType
    global scaleValue
    global scaleFactor
    global scale_window
    scaleType = args[0]

    if scaleType == 1:
        scaleFactor = 1 - scaleValue/100.0
    else:
        scaleFactor = 1 + scaleValue/100.0
    if scaleFactor ==0 :
        scaleFactor = 1
    scaledImage = cv2.resize(im, None, fx=scaleFactor,\
            fy = scaleFactor, interpolation = cv2.INTER_LINEAR)
    show_stype(scaleFactor == 0)
    cv2.imshow(scale_window, stp)
    cv2.imshow(windowName, scaledImage)

def scaleImage(*args):
    global scaleValue
    global scaleType
    global scaleFactor
    scaleValue = args[0]
    if scaleType == 1:
        scaleFactor = 1 - scaleValue/100.0
    else:
        scaleFactor = 1 + scaleValue/100.0
    if scaleFactor == 0:
        scaleFactor = 1
    scaledImage = cv2.resize(im, None, fx=scaleFactor,\
            fy = scaleFactor, interpolation = cv2.INTER_LINEAR)
    cv2.imshow(windowName, scaledImage)

cv2.createTrackbar(trackbarValue, windowName, scaleValue, maxScaleUp, scaleImage)
cv2.createTrackbar(trackbarType, windowName, scaleType, maxType, scaleTypeImage)

scaleImage(10)

while True:
    c = cv2.waitKey(10)
    if c==27:
        break

cv2.destroyAllWindows()